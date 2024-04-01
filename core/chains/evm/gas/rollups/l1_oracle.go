package rollups

import (
	"context"
	"fmt"
	"math/big"
	"slices"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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

//go:generate mockery --quiet --name ethClient --output ./mocks/ --case=underscore --structname ETHClient
type ethClient interface {
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
}

//go:generate mockery --quiet --name daPriceReader --output ./mocks/ --case=underscore --structname DAPriceReader
type daPriceReader interface {
	GetDAGasPrice(ctx context.Context) (*big.Int, error)
}

type priceEntry struct {
	price     *assets.Wei
	timestamp time.Time
}

// Reads L2-specific precompiles and caches the l1GasPrice set by the L2.
type l1Oracle struct {
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

const (
	// Interval at which to poll for L1BaseFee. A good starting point is the L1 block time.
	PollPeriod = 6 * time.Second
)

var supportedChainTypes = []config.ChainType{config.ChainArbitrum, config.ChainOptimismBedrock, config.ChainKroma, config.ChainScroll}

func IsRollupWithL1Support(chainType config.ChainType) bool {
	return slices.Contains(supportedChainTypes, chainType)
}

func NewL1GasOracle(lggr logger.Logger, ethClient ethClient, chainType config.ChainType) L1Oracle {
	var l1Oracle L1Oracle
	switch chainType {
	case config.ChainOptimismBedrock:
		l1Oracle = NewOpStackL1GasOracle(lggr, ethClient, chainType)
	case config.ChainKroma:
		l1Oracle = NewOpStackL1GasOracle(lggr, ethClient, chainType)
	case config.ChainArbitrum:
		l1Oracle = NewArbitrumL1GasOracle(lggr, ethClient)
	case config.ChainScroll:
		l1Oracle = NewScrollL1GasOracle(lggr, ethClient)
	default:
		panic(fmt.Sprintf("Received unspported chaintype %s", chainType))
	}
	return l1Oracle
}

func (o *l1Oracle) Name() string {
	return o.logger.Name()
}

func (o *l1Oracle) Start(ctx context.Context) error {
	return o.StartOnce(o.Name(), func() error {
		go o.run()
		<-o.chInitialised
		return nil
	})
}
func (o *l1Oracle) Close() error {
	return o.StopOnce(o.Name(), func() error {
		close(o.chStop)
		<-o.chDone
		return nil
	})
}

func (o *l1Oracle) HealthReport() map[string]error {
	return map[string]error{o.Name(): o.Healthy()}
}

func (o *l1Oracle) run() {
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
func (o *l1Oracle) refresh() (t *time.Timer) {
	t, err := o.refreshWithError()
	if err != nil {
		o.SvcErrBuffer.Append(err)
	}
	return
}

func (o *l1Oracle) refreshWithError() (t *time.Timer, err error) {
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

func (o *l1Oracle) fetchL1GasPrice(ctx context.Context) (price *big.Int, err error) {
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

func (o *l1Oracle) GasPrice(_ context.Context) (l1GasPrice *assets.Wei, err error) {
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
func (o *l1Oracle) GetGasCost(ctx context.Context, tx *gethtypes.Transaction, blockNum *big.Int) (*assets.Wei, error) {
	ctx, cancel := context.WithTimeout(ctx, client.QueryTimeout)
	defer cancel()
	var callData, b []byte
	var err error
	if o.chainType == config.ChainOptimismBedrock || o.chainType == config.ChainScroll {
		// Append rlp-encoded tx
		var encodedtx []byte
		if encodedtx, err = tx.MarshalBinary(); err != nil {
			return nil, fmt.Errorf("failed to marshal tx for gas cost estimation: %w", err)
		}
		if callData, err = o.l1GasCostMethodAbi.Pack(o.gasCostMethod, encodedtx); err != nil {
			return nil, fmt.Errorf("failed to pack calldata for %s L1 gas cost estimation method: %w", o.chainType, err)
		}
	} else if o.chainType == config.ChainArbitrum {
		if callData, err = o.l1GasCostMethodAbi.Pack(o.gasCostMethod, tx.To(), false, tx.Data()); err != nil {
			return nil, fmt.Errorf("failed to pack calldata for %s L1 gas cost estimation method: %w", o.chainType, err)
		}
	} else {
		return nil, fmt.Errorf("L1 gas cost not supported for this chain: %s", o.chainType)
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
	if o.chainType == config.ChainOptimismBedrock || o.chainType == config.ChainScroll {
		if len(b) != 32 { // returns uint256;
			errorMsg := fmt.Sprintf("return data length (%d) different than expected (%d)", len(b), 32)
			o.logger.Critical(errorMsg)
			return nil, fmt.Errorf(errorMsg)
		}
		l1GasCost = new(big.Int).SetBytes(b)
	} else if o.chainType == config.ChainArbitrum {
		if len(b) != 8+2*32 { // returns (uint64 gasEstimateForL1, uint256 baseFee, uint256 l1BaseFeeEstimate);
			errorMsg := fmt.Sprintf("return data length (%d) different than expected (%d)", len(b), 8+2*32)
			o.logger.Critical(errorMsg)
			return nil, fmt.Errorf(errorMsg)
		}
		l1GasCost = new(big.Int).SetBytes(b[:8])
	}

	return assets.NewWei(l1GasCost), nil
}
