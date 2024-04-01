package gas_test

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
)

type arbConfig struct {
	v               uint64
	bumpPercent     uint16
	bumpMin         *assets.Wei
	limitMultiplier float32
}

func (a *arbConfig) LimitMax() uint64 {
	return a.v
}

func (a *arbConfig) BumpPercent() uint16 {
	return a.bumpPercent
}

func (a *arbConfig) BumpMin() *assets.Wei {
	return a.bumpMin
}

func (a *arbConfig) LimitMultiplier() float32 {
	return a.limitMultiplier
}

func TestArbitrumEstimator(t *testing.T) {
	t.Parallel()

	maxGasPrice := assets.NewWeiI(100)
	const maxGasLimit uint64 = 500_000
	calldata := []byte{0x00, 0x00, 0x01, 0x02, 0x03}
	const gasLimit uint64 = 80000
	const gasPriceBufferPercentage = 50
	const bumpPercent = 10
	var bumpMin = assets.NewWei(big.NewInt(1))
	const limitMultiplier = 1
	const chainType = "arbitrum"

	t.Run("calling GetLegacyGas on unstarted estimator returns error", func(t *testing.T) {
		rpcClient := mocks.NewRPCClient(t)
		ethClient := mocks.NewETHClient(t)
		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, rpcClient, ethClient)
		_, _, err := o.GetLegacyGas(testutils.Context(t), calldata, gasLimit, maxGasPrice)
		assert.EqualError(t, err, "estimator is not started")
	})

	var zeros bytes.Buffer
	zeros.Write(common.BigToHash(big.NewInt(0)).Bytes())
	zeros.Write(common.BigToHash(big.NewInt(0)).Bytes())
	zeros.Write(common.BigToHash(big.NewInt(123455)).Bytes())
	t.Run("calling GetLegacyGas on started estimator returns estimates", func(t *testing.T) {
		rpcClient := mocks.NewRPCClient(t)
		ethClient := mocks.NewETHClient(t)
		rpcClient.On("CallContext", mock.Anything, mock.Anything, "eth_gasPrice").Return(nil).Run(func(args mock.Arguments) {
			res := args.Get(1).(*hexutil.Big)
			(*big.Int)(res).SetInt64(42)
		})
		ethClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			assert.Equal(t, gas.ArbGasInfoAddress, callMsg.To.String())
			assert.Equal(t, gas.ArbGasInfo_getPricesInArbGas, fmt.Sprintf("%x", callMsg.Data))
			assert.Equal(t, big.NewInt(-1), blockNumber)
		}).Return(zeros.Bytes(), nil)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{v: maxGasLimit, bumpPercent: bumpPercent, bumpMin: bumpMin, limitMultiplier: limitMultiplier}, rpcClient, ethClient)
		servicetest.RunHealthy(t, o)
		gasPrice, chainSpecificGasLimit, err := o.GetLegacyGas(testutils.Context(t), calldata, gasLimit, maxGasPrice)
		require.NoError(t, err)
		// Expected price for a standard l2_suggested_estimator would be 42, but we add a fixed gasPriceBufferPercentage.
		assert.Equal(t, assets.NewWeiI(42).AddPercentage(gasPriceBufferPercentage), gasPrice)
		assert.Equal(t, gasLimit, chainSpecificGasLimit)
	})

	t.Run("gas price is lower than user specified max gas price", func(t *testing.T) {
		client := mocks.NewRPCClient(t)
		ethClient := mocks.NewETHClient(t)
		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, client, ethClient)

		client.On("CallContext", mock.Anything, mock.Anything, "eth_gasPrice").Return(nil).Run(func(args mock.Arguments) {
			res := args.Get(1).(*hexutil.Big)
			(*big.Int)(res).SetInt64(42)
		})
		ethClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			assert.Equal(t, gas.ArbGasInfoAddress, callMsg.To.String())
			assert.Equal(t, gas.ArbGasInfo_getPricesInArbGas, fmt.Sprintf("%x", callMsg.Data))
			assert.Equal(t, big.NewInt(-1), blockNumber)
		}).Return(zeros.Bytes(), nil)

		servicetest.RunHealthy(t, o)
		gasPrice, chainSpecificGasLimit, err := o.GetLegacyGas(testutils.Context(t), calldata, gasLimit, assets.NewWeiI(40))
		require.Error(t, err)
		assert.EqualError(t, err, "estimated gas price: 42 wei is greater than the maximum gas price configured: 40 wei")
		assert.Nil(t, gasPrice)
		assert.Equal(t, uint64(0), chainSpecificGasLimit)
	})

	t.Run("gas price is lower than global max gas price", func(t *testing.T) {
		ethClient := mocks.NewETHClient(t)
		client := mocks.NewRPCClient(t)
		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, client, ethClient)

		client.On("CallContext", mock.Anything, mock.Anything, "eth_gasPrice").Return(nil).Run(func(args mock.Arguments) {
			res := args.Get(1).(*hexutil.Big)
			(*big.Int)(res).SetInt64(120)
		})
		ethClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			assert.Equal(t, gas.ArbGasInfoAddress, callMsg.To.String())
			assert.Equal(t, gas.ArbGasInfo_getPricesInArbGas, fmt.Sprintf("%x", callMsg.Data))
			assert.Equal(t, big.NewInt(-1), blockNumber)
		}).Return(zeros.Bytes(), nil)

		servicetest.RunHealthy(t, o)
		gasPrice, chainSpecificGasLimit, err := o.GetLegacyGas(testutils.Context(t), calldata, gasLimit, assets.NewWeiI(110))
		assert.EqualError(t, err, "estimated gas price: 120 wei is greater than the maximum gas price configured: 110 wei")
		assert.Nil(t, gasPrice)
		assert.Equal(t, uint64(0), chainSpecificGasLimit)
	})

	t.Run("calling BumpLegacyGas on unstarted arbitrum estimator returns error", func(t *testing.T) {
		rpcClient := mocks.NewRPCClient(t)
		ethClient := mocks.NewETHClient(t)
		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, rpcClient, ethClient)
		_, _, err := o.BumpLegacyGas(testutils.Context(t), assets.NewWeiI(42), gasLimit, assets.NewWeiI(10), nil)
		assert.EqualError(t, err, "estimator is not started")
	})

	t.Run("calling GetLegacyGas on started estimator if initial call failed returns error", func(t *testing.T) {
		client := mocks.NewRPCClient(t)
		ethClient := mocks.NewETHClient(t)
		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, client, ethClient)

		client.On("CallContext", mock.Anything, mock.Anything, "eth_gasPrice").Return(pkgerrors.New("kaboom"))
		ethClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			assert.Equal(t, gas.ArbGasInfoAddress, callMsg.To.String())
			assert.Equal(t, gas.ArbGasInfo_getPricesInArbGas, fmt.Sprintf("%x", callMsg.Data))
			assert.Equal(t, big.NewInt(-1), blockNumber)
		}).Return(zeros.Bytes(), nil)

		servicetest.RunHealthy(t, o)

		_, _, err := o.GetLegacyGas(testutils.Context(t), calldata, gasLimit, maxGasPrice)
		assert.EqualError(t, err, "failed to estimate gas; gas price not set")
	})

	t.Run("calling GetDynamicFee always returns error", func(t *testing.T) {
		rpcClient := mocks.NewRPCClient(t)
		ethClient := mocks.NewETHClient(t)
		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, rpcClient, ethClient)
		_, err := o.GetDynamicFee(testutils.Context(t), maxGasPrice)
		assert.EqualError(t, err, "dynamic fees are not implemented for this estimator")
	})

	t.Run("calling BumpDynamicFee always returns error", func(t *testing.T) {
		rpcClient := mocks.NewRPCClient(t)
		ethClient := mocks.NewETHClient(t)
		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, rpcClient, ethClient)
		fee := gas.DynamicFee{
			FeeCap: assets.NewWeiI(42),
			TipCap: assets.NewWeiI(5),
		}
		_, err := o.BumpDynamicFee(testutils.Context(t), fee, maxGasPrice, nil)
		assert.EqualError(t, err, "dynamic fees are not implemented for this estimator")
	})

	t.Run("limit computes", func(t *testing.T) {
		rpcClient := mocks.NewRPCClient(t)
		ethClient := mocks.NewETHClient(t)
		rpcClient.On("CallContext", mock.Anything, mock.Anything, "eth_gasPrice").Return(nil).Run(func(args mock.Arguments) {
			res := args.Get(1).(*hexutil.Big)
			(*big.Int)(res).SetInt64(42)
		})
		const (
			perL2Tx       = 50_000
			perL1Calldata = 10_000
		)
		var expLimit = gasLimit + perL2Tx + perL1Calldata*uint64(len(calldata))

		var b bytes.Buffer
		b.Write(common.BigToHash(big.NewInt(perL2Tx)).Bytes())
		b.Write(common.BigToHash(big.NewInt(perL1Calldata)).Bytes())
		b.Write(common.BigToHash(big.NewInt(123455)).Bytes())
		ethClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			assert.Equal(t, gas.ArbGasInfoAddress, callMsg.To.String())
			assert.Equal(t, gas.ArbGasInfo_getPricesInArbGas, fmt.Sprintf("%x", callMsg.Data))
			assert.Equal(t, big.NewInt(-1), blockNumber)
		}).Return(b.Bytes(), nil)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{v: maxGasLimit, bumpPercent: bumpPercent, bumpMin: bumpMin, limitMultiplier: limitMultiplier}, rpcClient, ethClient)
		servicetest.RunHealthy(t, o)
		gasPrice, chainSpecificGasLimit, err := o.GetLegacyGas(testutils.Context(t), calldata, gasLimit, maxGasPrice)
		require.NoError(t, err)
		require.NotNil(t, gasPrice)
		// Again, a normal l2_suggested_estimator would return 42, but arbitrum_estimator adds a buffer.
		assert.Equal(t, "63 wei", gasPrice.String())
		assert.Equal(t, expLimit, chainSpecificGasLimit, "expected %d but got %d", expLimit, chainSpecificGasLimit)
	})

	t.Run("limit exceeds max", func(t *testing.T) {
		rpcClient := mocks.NewRPCClient(t)
		ethClient := mocks.NewETHClient(t)
		rpcClient.On("CallContext", mock.Anything, mock.Anything, "eth_gasPrice").Return(nil).Run(func(args mock.Arguments) {
			res := args.Get(1).(*hexutil.Big)
			(*big.Int)(res).SetInt64(42)
		})
		const (
			perL2Tx       = 500_000
			perL1Calldata = 100_000
		)

		var b bytes.Buffer
		b.Write(common.BigToHash(big.NewInt(perL2Tx)).Bytes())
		b.Write(common.BigToHash(big.NewInt(perL1Calldata)).Bytes())
		b.Write(common.BigToHash(big.NewInt(123455)).Bytes())
		ethClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			assert.Equal(t, gas.ArbGasInfoAddress, callMsg.To.String())
			assert.Equal(t, gas.ArbGasInfo_getPricesInArbGas, fmt.Sprintf("%x", callMsg.Data))
			assert.Equal(t, big.NewInt(-1), blockNumber)
		}).Return(b.Bytes(), nil)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{v: maxGasLimit, bumpPercent: bumpPercent, bumpMin: bumpMin, limitMultiplier: limitMultiplier}, rpcClient, ethClient,)
		servicetest.RunHealthy(t, o)
		gasPrice, chainSpecificGasLimit, err := o.GetLegacyGas(testutils.Context(t), calldata, gasLimit, maxGasPrice)
		require.Error(t, err, "expected error but got (%s, %d)", gasPrice, chainSpecificGasLimit)
	})
}
