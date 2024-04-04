package rollups

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/utils"

	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/v2/common/client"
	"github.com/smartcontractkit/chainlink/v2/common/config"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
)

// Reads L2-specific precompiles and caches the l1GasPrice set by the L2.
type optimismL1Oracle struct {
	services.StateMachine
	client     ethClient
	pollPeriod time.Duration
	logger     logger.SugaredLogger
	chainType  config.ChainType

	l1GasPriceAddress   string
	gasPriceMethod      string
	l1GasPriceMethodAbi abi.ABI
	l1GasPriceMu        sync.RWMutex
	l1GasPrice          priceEntry

	l1GasCostAddress   string
	gasCostMethod      string
	l1GasCostMethodAbi abi.ABI

	priceReader daPriceReader

	chInitialised chan struct{}
	chStop        services.StopChan
	chDone        chan struct{}
}

type opStackGasPriceReader struct {
	client ethClient
	logger logger.SugaredLogger

	oracleAddress      common.Address
	isEcotoneMethodAbi abi.ABI

	l1BaseFeeCalldata    []byte
	isEcotoneCalldata    []byte
	getL1GasUsedCalldata []byte
	getL1FeeCalldata     []byte

	isEcotone        bool
	isEcotoneCheckTs int64
}

const (
	// OPStackGasOracle_l1BaseFee fetches the l1 base fee set in the OP Stack GasPriceOracle contract
	OPStackGasOracle_l1BaseFee = "l1BaseFee"

	// OPStackGasOracle_isEcotone fetches if the OP Stack GasPriceOracle contract has upgraded to Ecotone
	OPStackGasOracle_isEcotone = "isEcotone"

	// OPStackGasOracle_getL1GasUsed fetches the l1 gas used for given tx bytes
	OPStackGasOracle_getL1GasUsed = "getL1GasUsed"

	// OPStackGasOracle_getL1Fee fetches the l1 fee for given tx bytes
	OPStackGasOracle_getL1Fee = "getL1Fee"

	// OPStackGasOracle_isEcotonePollingPeriod is the interval to poll if chain has upgraded to Ecotone
	// Set to poll every 4 hours
	OPStackGasOracle_isEcotonePollingPeriod = 14400
	// OPGasOracleAddress is the address of the precompiled contract that exists on OP stack chain.
	// This is the case for Optimism and Base.
	OPGasOracleAddress = "0x420000000000000000000000000000000000000F"
	// OPGasOracle_l1BaseFee is a hex encoded call to:
	// `function l1BaseFee() external view returns (uint256);`
	OPGasOracle_l1BaseFee = "l1BaseFee"
	// OPGasOracle_getL1Fee is a hex encoded call to:
	// `function getL1Fee(bytes) external view returns (uint256);`
	OPGasOracle_getL1Fee = "getL1Fee"

	// GasOracleAddress is the address of the precompiled contract that exists on Kroma chain.
	// This is the case for Kroma.
	KromaGasOracleAddress = "0x4200000000000000000000000000000000000005"
	// GasOracle_l1BaseFee is the a hex encoded call to:
	// `function l1BaseFee() external view returns (uint256);`
	KromaGasOracle_l1BaseFee = "l1BaseFee"
)

func NewOpStackL1GasOracle(lggr logger.Logger, ethClient ethClient, chainType config.ChainType) L1Oracle {
	var precompileAddress string
	switch chainType {
	case config.ChainOptimismBedrock:
		precompileAddress = OPGasOracleAddress
	case config.ChainKroma:
		precompileAddress = KromaGasOracleAddress
	default:
		panic(fmt.Sprintf("Received unspported chaintype %s", chainType))
	}
	priceReader := newOPPriceReader(lggr, ethClient, chainType, precompileAddress)
	return newOpStackL1GasOracle(lggr, ethClient, priceReader, chainType)
}

func newOpStackL1GasOracle(lggr logger.Logger, ethClient ethClient, priceReader daPriceReader, chainType config.ChainType) L1Oracle {
	var l1GasPriceAddress, gasPriceMethod, l1GasCostAddress, gasCostMethod string
	var l1GasPriceMethodAbi, l1GasCostMethodAbi abi.ABI
	var gasPriceErr, gasCostErr error

	l1GasPriceAddress = OPGasOracleAddress
	gasPriceMethod = OPGasOracle_l1BaseFee
	l1GasPriceMethodAbi, gasPriceErr = abi.JSON(strings.NewReader(L1BaseFeeAbiString))
	l1GasCostAddress = OPGasOracleAddress
	gasCostMethod = OPGasOracle_getL1Fee
	l1GasCostMethodAbi, gasCostErr = abi.JSON(strings.NewReader(GetL1FeeAbiString))

	if gasPriceErr != nil {
		panic(fmt.Sprintf("Failed to parse L1 gas price method ABI for chain: optimismBedrock"))
	}
	if gasCostErr != nil {
		panic(fmt.Sprintf("Failed to parse L1 gas cost method ABI for chain: optimismBedrock"))
	}

	return &optimismL1Oracle{
		client:     ethClient,
		pollPeriod: PollPeriod,
		logger:     logger.Sugared(logger.Named(lggr, fmt.Sprintf("L1GasOracle(optimismBedrock)"))),
		chainType:  chainType,

		l1GasPriceAddress:   l1GasPriceAddress,
		gasPriceMethod:      gasPriceMethod,
		l1GasPriceMethodAbi: l1GasPriceMethodAbi,
		l1GasCostAddress:    l1GasCostAddress,
		gasCostMethod:       gasCostMethod,
		l1GasCostMethodAbi:  l1GasCostMethodAbi,

		priceReader: priceReader,

		chInitialised: make(chan struct{}),
		chStop:        make(chan struct{}),
		chDone:        make(chan struct{}),
	}
}

func (o *optimismL1Oracle) Name() string {
	return o.logger.Name()
}

func (o *optimismL1Oracle) Start(ctx context.Context) error {
	return o.StartOnce(o.Name(), func() error {
		go o.run()
		<-o.chInitialised
		return nil
	})
}
func (o *optimismL1Oracle) Close() error {
	return o.StopOnce(o.Name(), func() error {
		close(o.chStop)
		<-o.chDone
		return nil
	})
}

func (o *optimismL1Oracle) HealthReport() map[string]error {
	return map[string]error{o.Name(): o.Healthy()}
}

func (o *optimismL1Oracle) run() {
	defer close(o.chDone)

	t := o.refresh()
	close(o.chInitialised)

	for {
		select {
		case <-o.chStop:
			return
		case <-t.C:
			t = o.refresh()
		}
	}
}
func (o *optimismL1Oracle) refresh() (t *time.Timer) {
	t, err := o.refreshWithError()
	if err != nil {
		o.SvcErrBuffer.Append(err)
	}
	return
}

func (o *optimismL1Oracle) refreshWithError() (t *time.Timer, err error) {
	t = time.NewTimer(utils.WithJitter(o.pollPeriod))

	ctx, cancel := o.chStop.CtxCancel(evmclient.ContextWithDefaultTimeout())
	defer cancel()

	price, err := o.fetchL1GasPrice(ctx)
	if err != nil {
		return t, err
	}

	o.l1GasPriceMu.Lock()
	defer o.l1GasPriceMu.Unlock()
	o.l1GasPrice = priceEntry{price: assets.NewWei(price), timestamp: time.Now()}
	return
}

func (o *optimismL1Oracle) fetchL1GasPrice(ctx context.Context) (price *big.Int, err error) {
	// if dedicated priceReader exists, use the reader
	if o.priceReader != nil {
		return o.priceReader.GetDAGasPrice(ctx)
	}

	var callData, b []byte
	precompile := common.HexToAddress(o.l1GasPriceAddress)
	callData, err = o.l1GasPriceMethodAbi.Pack(o.gasPriceMethod)
	if err != nil {
		errMsg := fmt.Sprintf("failed to pack calldata for %s L1 gas price method", o.chainType)
		o.logger.Errorf(errMsg)
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}
	b, err = o.client.CallContract(ctx, ethereum.CallMsg{
		To:   &precompile,
		Data: callData,
	}, nil)
	if err != nil {
		errMsg := "gas oracle contract call failed"
		o.logger.Errorf(errMsg)
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	if len(b) != 32 { // returns uint256;
		errMsg := fmt.Sprintf("return data length (%d) different than expected (%d)", len(b), 32)
		o.logger.Criticalf(errMsg)
		return nil, fmt.Errorf(errMsg)
	}
	price = new(big.Int).SetBytes(b)
	return price, nil
}

func (o *optimismL1Oracle) GasPrice(_ context.Context) (l1GasPrice *assets.Wei, err error) {
	var timestamp time.Time
	ok := o.IfStarted(func() {
		o.l1GasPriceMu.RLock()
		l1GasPrice = o.l1GasPrice.price
		timestamp = o.l1GasPrice.timestamp
		o.l1GasPriceMu.RUnlock()
	})
	if !ok {
		return l1GasPrice, fmt.Errorf("L1GasOracle is not started; cannot estimate gas")
	}
	if l1GasPrice == nil {
		return l1GasPrice, fmt.Errorf("failed to get l1 gas price; gas price not set")
	}
	// Validate the price has been updated within the pollPeriod * 2
	// Allowing double the poll period before declaring the price stale to give ample time for the refresh to process
	if time.Since(timestamp) > o.pollPeriod*2 {
		return l1GasPrice, fmt.Errorf("gas price is stale")
	}
	return
}

// Gets the L1 gas cost for the provided transaction at the specified block num
// If block num is not provided, the value on the latest block num is used
func (o *optimismL1Oracle) GetGasCost(ctx context.Context, tx *gethtypes.Transaction, blockNum *big.Int) (*assets.Wei, error) {
	ctx, cancel := context.WithTimeout(ctx, client.QueryTimeout)
	defer cancel()
	var callData, b []byte
	var err error
	if o.chainType == config.ChainKroma {
		return nil, fmt.Errorf("L1 gas cost not supported for this chain: %s", o.chainType)
	}
	// Append rlp-encoded tx
	var encodedtx []byte
	if encodedtx, err = tx.MarshalBinary(); err != nil {
		return nil, fmt.Errorf("failed to marshal tx for gas cost estimation: %w", err)
	}
	if callData, err = o.l1GasCostMethodAbi.Pack(o.gasCostMethod, encodedtx); err != nil {
		return nil, fmt.Errorf("failed to pack calldata for %s L1 gas cost estimation method: %w", o.chainType, err)
	}

	precompile := common.HexToAddress(o.l1GasCostAddress)
	b, err = o.client.CallContract(ctx, ethereum.CallMsg{
		To:   &precompile,
		Data: callData,
	}, blockNum)
	if err != nil {
		errorMsg := fmt.Sprintf("gas oracle contract call failed: %v", err)
		o.logger.Errorf(errorMsg)
		return nil, fmt.Errorf(errorMsg)
	}

	var l1GasCost *big.Int
	if len(b) != 32 { // returns uint256;
		errorMsg := fmt.Sprintf("return data length (%d) different than expected (%d)", len(b), 32)
		o.logger.Critical(errorMsg)
		return nil, fmt.Errorf(errorMsg)
	}
	l1GasCost = new(big.Int).SetBytes(b)

	return assets.NewWei(l1GasCost), nil
}

func newOPPriceReader(lggr logger.Logger, ethClient ethClient, chainType config.ChainType, oracleAddress string) daPriceReader {
	// encode calldata for each method; these calldata will remain the same for each call, we can encode them just once
	l1BaseFeeMethodAbi, err := abi.JSON(strings.NewReader(L1BaseFeeAbiString))
	if err != nil {
		panic(fmt.Errorf("failed to parse GasPriceOracle %s() method ABI for chain: %s; %w", OPStackGasOracle_l1BaseFee, chainType, err))
	}
	l1BaseFeeCalldata, err := l1BaseFeeMethodAbi.Pack(OPStackGasOracle_l1BaseFee)
	if err != nil {
		panic(fmt.Errorf("failed to parse GasPriceOracle %s() calldata for chain: %s; %w", OPStackGasOracle_l1BaseFee, chainType, err))
	}

	isEcotoneMethodAbi, err := abi.JSON(strings.NewReader(OPIsEcotoneAbiString))
	if err != nil {
		panic(fmt.Errorf("failed to parse GasPriceOracle %s() method ABI for chain: %s; %w", OPStackGasOracle_isEcotone, chainType, err))
	}
	isEcotoneCalldata, err := isEcotoneMethodAbi.Pack(OPStackGasOracle_isEcotone)
	if err != nil {
		panic(fmt.Errorf("failed to parse GasPriceOracle %s() calldata for chain: %s; %w", OPStackGasOracle_isEcotone, chainType, err))
	}

	getL1GasUsedMethodAbi, err := abi.JSON(strings.NewReader(OPGetL1GasUsedAbiString))
	if err != nil {
		panic(fmt.Errorf("failed to parse GasPriceOracle %s() method ABI for chain: %s; %w", OPStackGasOracle_getL1GasUsed, chainType, err))
	}
	getL1GasUsedCalldata, err := getL1GasUsedMethodAbi.Pack(OPStackGasOracle_getL1GasUsed, []byte{0x1})
	if err != nil {
		panic(fmt.Errorf("failed to parse GasPriceOracle %s() calldata for chain: %s; %w", OPStackGasOracle_getL1GasUsed, chainType, err))
	}

	getL1FeeMethodAbi, err := abi.JSON(strings.NewReader(GetL1FeeAbiString))
	if err != nil {
		panic(fmt.Errorf("failed to parse GasPriceOracle %s() method ABI for chain: %s; %w", OPStackGasOracle_getL1Fee, chainType, err))
	}
	getL1FeeCalldata, err := getL1FeeMethodAbi.Pack(OPStackGasOracle_getL1Fee, []byte{0x1})
	if err != nil {
		panic(fmt.Errorf("failed to parse GasPriceOracle %s() calldata for chain: %s; %w", OPStackGasOracle_getL1Fee, chainType, err))
	}

	return &opStackGasPriceReader{
		client: ethClient,
		logger: logger.Sugared(logger.Named(lggr, fmt.Sprintf("OPStackGasOracle(%s)", chainType))),

		oracleAddress:      common.HexToAddress(oracleAddress),
		isEcotoneMethodAbi: isEcotoneMethodAbi,

		l1BaseFeeCalldata:    l1BaseFeeCalldata,
		isEcotoneCalldata:    isEcotoneCalldata,
		getL1GasUsedCalldata: getL1GasUsedCalldata,
		getL1FeeCalldata:     getL1FeeCalldata,

		isEcotone:        false,
		isEcotoneCheckTs: 0,
	}
}

func (o *opStackGasPriceReader) GetDAGasPrice(ctx context.Context) (*big.Int, error) {
	isEcotone, err := o.checkIsEcotone(ctx)
	if err != nil {
		return nil, err
	}

	o.logger.Infof("Chain isEcotone result: %t", isEcotone)

	if isEcotone {
		return o.getEcotoneGasPrice(ctx)
	}

	return o.getV1GasPrice(ctx)
}

func (o *opStackGasPriceReader) checkIsEcotone(ctx context.Context) (bool, error) {
	// if chain is already Ecotone, NOOP
	if o.isEcotone {
		return true, nil
	}
	// if time since last check has not exceeded polling period, NOOP
	if time.Now().Unix()-o.isEcotoneCheckTs < OPStackGasOracle_isEcotonePollingPeriod {
		return false, nil
	}
	o.isEcotoneCheckTs = time.Now().Unix()

	// confirmed with OP team that isEcotone() is the canonical way to check if the chain has upgraded
	b, err := o.client.CallContract(ctx, ethereum.CallMsg{
		To:   &o.oracleAddress,
		Data: o.isEcotoneCalldata,
	}, nil)

	// if the chain has not upgraded to Ecotone, the isEcotone call will revert, this would be expected
	if err != nil {
		o.logger.Infof("isEcotone() call failed, this can happen if chain has not upgraded: %w", err)
		return false, nil
	}

	res, err := o.isEcotoneMethodAbi.Unpack(OPStackGasOracle_isEcotone, b)
	if err != nil {
		return false, fmt.Errorf("failed to unpack isEcotone() return data: %w", err)
	}
	o.isEcotone = res[0].(bool)
	return o.isEcotone, nil
}

func (o *opStackGasPriceReader) getV1GasPrice(ctx context.Context) (*big.Int, error) {
	b, err := o.client.CallContract(ctx, ethereum.CallMsg{
		To:   &o.oracleAddress,
		Data: o.l1BaseFeeCalldata,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("l1BaseFee() call failed: %w", err)
	}

	if len(b) != 32 {
		return nil, fmt.Errorf("l1BaseFee() return data length (%d) different than expected (%d)", len(b), 32)
	}
	return new(big.Int).SetBytes(b), nil
}

func (o *opStackGasPriceReader) getEcotoneGasPrice(ctx context.Context) (*big.Int, error) {
	rpcBatchCalls := []rpc.BatchElem{
		{
			Method: "eth_call",
			Args: []any{
				map[string]interface{}{
					"from": common.Address{},
					"to":   o.oracleAddress,
					"data": hexutil.Bytes(o.getL1GasUsedCalldata),
				},
				"latest",
			},
			Result: new(string),
		},
		{
			Method: "eth_call",
			Args: []any{
				map[string]interface{}{
					"from": common.Address{},
					"to":   o.oracleAddress,
					"data": hexutil.Bytes(o.getL1FeeCalldata),
				},
				"latest",
			},
			Result: new(string),
		},
	}

	err := o.client.BatchCallContext(ctx, rpcBatchCalls)
	if err != nil {
		return nil, fmt.Errorf("getEcotoneGasPrice batch call failed: %w", err)
	}
	if rpcBatchCalls[0].Error != nil {
		return nil, fmt.Errorf("%s call failed in a batch: %w", OPStackGasOracle_getL1GasUsed, err)
	}
	if rpcBatchCalls[1].Error != nil {
		return nil, fmt.Errorf("%s call failed in a batch: %w", OPStackGasOracle_getL1Fee, err)
	}

	l1GasUsedResult := *(rpcBatchCalls[0].Result.(*string))
	l1FeeResult := *(rpcBatchCalls[1].Result.(*string))

	l1GasUsedBytes, err := hexutil.Decode(l1GasUsedResult)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s rpc result: %w", OPStackGasOracle_getL1GasUsed, err)
	}
	l1FeeBytes, err := hexutil.Decode(l1FeeResult)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s rpc result: %w", OPStackGasOracle_getL1Fee, err)
	}

	l1GasUsed := new(big.Int).SetBytes(l1GasUsedBytes)
	l1Fee := new(big.Int).SetBytes(l1FeeBytes)

	// for the same tx byte, l1Fee / l1GasUsed will give the l1 gas price
	// note this price is per l1 gas, not l1 data byte
	return new(big.Int).Div(l1Fee, l1GasUsed), nil
}
