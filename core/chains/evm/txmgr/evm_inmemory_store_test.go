package txmgr_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	commontxmgr "github.com/smartcontractkit/chainlink/v2/common/txmgr"

	evmgas "github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	evmtxmgr "github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/evmtest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
)

func TestInMemoryStore_SetBroadcastBeforeBlockNum(t *testing.T) {
	t.Parallel()

	db := pgtest.NewSqlxDB(t)
	_, dbcfg, evmcfg := evmtxmgr.MakeTestConfigs(t)
	persistentStore := cltest.NewTestTxStore(t, db, dbcfg)
	kst := cltest.NewKeyStore(t, db, dbcfg)
	_, fromAddress := cltest.MustInsertRandomKey(t, kst.Eth())

	ethClient := evmtest.NewEthClientMockWithDefaultChain(t)
	lggr := logger.TestSugared(t)
	chainID := ethClient.ConfiguredChainID()
	ctx := context.Background()

	inMemoryStore, err := commontxmgr.NewInMemoryStore[
		*big.Int,
		common.Address, common.Hash, common.Hash,
		*evmtypes.Receipt,
		evmtypes.Nonce,
		evmgas.EvmFee,
	](ctx, lggr, chainID, kst.Eth(), persistentStore, evmcfg.Transactions())
	require.NoError(t, err)

	t.Run("saves block num to unconfirmed evm.tx_attempts without one", func(t *testing.T) {
		// Insert a transaction into persistent store
		inTx := cltest.MustInsertUnconfirmedEthTxWithBroadcastLegacyAttempt(t, persistentStore, 1, fromAddress)
		// Insert the transaction into the in-memory store
		require.NoError(t, inMemoryStore.XXXTestInsertTx(fromAddress, &inTx))

		headNum := int64(9000)
		err := inMemoryStore.SetBroadcastBeforeBlockNum(testutils.Context(t), headNum, chainID)
		require.NoError(t, err)

		expTx, err := persistentStore.FindTxWithAttempts(inTx.ID)
		require.NoError(t, err)
		require.Equal(t, 1, len(expTx.TxAttempts))
		assert.Equal(t, headNum, *expTx.TxAttempts[0].BroadcastBeforeBlockNum)
		fn := func(tx *evmtxmgr.Tx) bool { return true }
		actTxs := inMemoryStore.XXXTestFindTxs(nil, fn, inTx.ID)
		require.Equal(t, 1, len(actTxs))
		actTx := actTxs[0]
		assertTxEqual(t, expTx, actTx)

		// wrong chain ID
		wrongChainID := big.NewInt(123)
		actErr := inMemoryStore.SetBroadcastBeforeBlockNum(testutils.Context(t), headNum, wrongChainID)
		expErr := persistentStore.SetBroadcastBeforeBlockNum(testutils.Context(t), headNum, wrongChainID)
		assert.NoError(t, actErr)
		assert.NoError(t, expErr)
	})

	t.Run("does not change evm.tx_attempts that already have BroadcastBeforeBlockNum set", func(t *testing.T) {
		n := int64(42)
		// Insert a transaction into persistent store
		inTx := cltest.MustInsertUnconfirmedEthTxWithBroadcastLegacyAttempt(t, persistentStore, 11, fromAddress)
		inTxAttempt := newBroadcastLegacyEthTxAttempt(t, inTx.ID, 2)
		inTxAttempt.BroadcastBeforeBlockNum = &n
		require.NoError(t, persistentStore.InsertTxAttempt(&inTxAttempt))
		// Insert the transaction into the in-memory store
		inTx.TxAttempts = append([]evmtxmgr.TxAttempt{inTxAttempt}, inTx.TxAttempts...)
		require.NoError(t, inMemoryStore.XXXTestInsertTx(fromAddress, &inTx))

		headNum := int64(9000)
		err := inMemoryStore.SetBroadcastBeforeBlockNum(testutils.Context(t), headNum, chainID)
		require.NoError(t, err)

		expTx, err := persistentStore.FindTxWithAttempts(inTx.ID)
		require.NoError(t, err)
		require.Equal(t, 2, len(expTx.TxAttempts))
		assert.Equal(t, n, *expTx.TxAttempts[0].BroadcastBeforeBlockNum)
		fn := func(tx *evmtxmgr.Tx) bool { return true }
		actTxs := inMemoryStore.XXXTestFindTxs(nil, fn, inTx.ID)
		require.Equal(t, 1, len(actTxs))
		actTx := actTxs[0]
		assertTxEqual(t, expTx, actTx)
	})
}

// assertTxEqual asserts that two transactions are equal
func assertTxEqual(t *testing.T, exp, act evmtxmgr.Tx) {
	assert.Equal(t, exp.ID, act.ID)
	assert.Equal(t, exp.IdempotencyKey, act.IdempotencyKey)
	assert.Equal(t, exp.Sequence, act.Sequence)
	assert.Equal(t, exp.FromAddress, act.FromAddress)
	assert.Equal(t, exp.ToAddress, act.ToAddress)
	assert.Equal(t, exp.EncodedPayload, act.EncodedPayload)
	assert.Equal(t, exp.Value, act.Value)
	assert.Equal(t, exp.FeeLimit, act.FeeLimit)
	assert.Equal(t, exp.Error, act.Error)
	assert.Equal(t, exp.BroadcastAt, act.BroadcastAt)
	assert.Equal(t, exp.InitialBroadcastAt, act.InitialBroadcastAt)
	assert.Equal(t, exp.CreatedAt, act.CreatedAt)
	assert.Equal(t, exp.State, act.State)
	assert.Equal(t, exp.Meta, act.Meta)
	assert.Equal(t, exp.Subject, act.Subject)
	assert.Equal(t, exp.ChainID, act.ChainID)
	assert.Equal(t, exp.PipelineTaskRunID, act.PipelineTaskRunID)
	assert.Equal(t, exp.MinConfirmations, act.MinConfirmations)
	assert.Equal(t, exp.TransmitChecker, act.TransmitChecker)
	assert.Equal(t, exp.SignalCallback, act.SignalCallback)
	assert.Equal(t, exp.CallbackCompleted, act.CallbackCompleted)

	require.Len(t, exp.TxAttempts, len(act.TxAttempts))
	for i := 0; i < len(exp.TxAttempts); i++ {
		assertTxAttemptEqual(t, exp.TxAttempts[i], act.TxAttempts[i])
	}
}

func assertTxAttemptEqual(t *testing.T, exp, act evmtxmgr.TxAttempt) {
	assert.Equal(t, exp.ID, act.ID)
	assert.Equal(t, exp.TxID, act.TxID)
	assert.Equal(t, exp.Tx, act.Tx)
	assert.Equal(t, exp.TxFee, act.TxFee)
	assert.Equal(t, exp.ChainSpecificFeeLimit, act.ChainSpecificFeeLimit)
	assert.Equal(t, exp.SignedRawTx, act.SignedRawTx)
	assert.Equal(t, exp.Hash, act.Hash)
	assert.Equal(t, exp.CreatedAt, act.CreatedAt)
	assert.Equal(t, exp.BroadcastBeforeBlockNum, act.BroadcastBeforeBlockNum)
	assert.Equal(t, exp.State, act.State)
	assert.Equal(t, exp.TxType, act.TxType)

	require.Equal(t, len(exp.Receipts), len(act.Receipts))
	for i := 0; i < len(exp.Receipts); i++ {
		assertChainReceiptEqual(t, exp.Receipts[i], act.Receipts[i])
	}
}

func assertChainReceiptEqual(t *testing.T, exp, act evmtxmgr.ChainReceipt) {
	assert.Equal(t, exp.GetStatus(), act.GetStatus())
	assert.Equal(t, exp.GetTxHash(), act.GetTxHash())
	assert.Equal(t, exp.GetBlockNumber(), act.GetBlockNumber())
	assert.Equal(t, exp.IsZero(), act.IsZero())
	assert.Equal(t, exp.IsUnmined(), act.IsUnmined())
	assert.Equal(t, exp.GetFeeUsed(), act.GetFeeUsed())
	assert.Equal(t, exp.GetTransactionIndex(), act.GetTransactionIndex())
	assert.Equal(t, exp.GetBlockHash(), act.GetBlockHash())
}
