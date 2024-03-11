// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf_coordinator_v2_5

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type VRFCoordinatorV25RequestCommitment struct {
	BlockNum         uint64
	SubId            *big.Int
	CallbackGasLimit uint32
	NumWords         uint32
	Sender           common.Address
	ExtraArgs        []byte
}

type VRFProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

type VRFV2PlusClientRandomWordsRequest struct {
	KeyHash              [32]byte
	SubId                *big.Int
	RequestConfirmations uint16
	CallbackGasLimit     uint32
	NumWords             uint32
	ExtraArgs            []byte
}

var VRFCoordinatorV25MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blockhashStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"internalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"externalBalance\",\"type\":\"uint256\"}],\"name\":\"BalanceInvariantViolated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"BlockhashNotInStore\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendNative\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToTransferLink\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"GasLimitTooBig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"}],\"name\":\"GasPriceExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"InsufficientGasForConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCalldata\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"InvalidConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"linkWei\",\"type\":\"int256\"}],\"name\":\"InvalidLinkWeiPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"have\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"min\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"max\",\"type\":\"uint16\"}],\"name\":\"InvalidRequestConfirmations\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"flatFeeLinkDiscountPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flatFeeNativePPM\",\"type\":\"uint32\"}],\"name\":\"LinkDiscountTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"max\",\"type\":\"uint32\"}],\"name\":\"MsgDataTooBig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedOwner\",\"type\":\"address\"}],\"name\":\"MustBeRequestedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoCorrespondingRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"NoSuchProvingKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"NumWordsTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableFromLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRequestExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"ProvingKeyAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConsumers\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeNativePPM\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkDiscountPPM\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"nativePremiumPercentage\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"linkPremiumPercentage\",\"type\":\"uint8\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"MigrationCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativeFundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxGas\",\"type\":\"uint64\"}],\"name\":\"ProvingKeyDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxGas\",\"type\":\"uint64\"}],\"name\":\"ProvingKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputSeed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"payment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"nativePayment\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"onlyPremium\",\"type\":\"bool\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preSeed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RandomWordsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountNative\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldNativeBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNativeBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFundedWithNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCKHASH_STORE\",\"outputs\":[{\"internalType\":\"contractBlockhashStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK_NATIVE_FEED\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CONSUMERS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_WORDS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REQUEST_CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"deregisterMigratableCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"deregisterProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"gamma\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uWitness\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"sHashWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structVRFCoordinatorV2_5.RequestCommitment\",\"name\":\"rc\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"onlyPremium\",\"type\":\"bool\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"payment\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"fundSubscriptionWithNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"getActiveSubscriptionIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"nativeBalance\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"reqCount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicKey\",\"type\":\"uint256[2]\"}],\"name\":\"hashOfKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newCoordinator\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"ownerCancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"pendingRequestExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverNativeFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"registerMigratableCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint64\",\"name\":\"maxGas\",\"type\":\"uint64\"}],\"name\":\"registerProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structVRFV2PlusClient.RandomWordsRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_config\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"reentrancyLock\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeNativePPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkDiscountPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"nativePremiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"linkPremiumPercentage\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_currentSubNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_fallbackWeiPerUnitLink\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_provingKeyHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"s_provingKeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"maxGas\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requestCommitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalBalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalNativeBalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeNativePPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkDiscountPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"nativePremiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"linkPremiumPercentage\",\"type\":\"uint8\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkNativeFeed\",\"type\":\"address\"}],\"name\":\"setLINKAndLINKNativeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162005de238038062005de2833981016040819052620000349162000183565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000d7565b50505060601b6001600160601b031916608052620001b5565b6001600160a01b038116331415620001325760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000602082840312156200019657600080fd5b81516001600160a01b0381168114620001ae57600080fd5b9392505050565b60805160601c615c07620001db6000396000818161055001526132280152615c076000f3fe60806040526004361061021c5760003560e01c80638402595e11610124578063b2a7cac5116100a6578063b2a7cac514610732578063bec4c08c14610752578063caf70c4a14610772578063cb63179714610792578063d98e620e146107b2578063da2f2610146107d2578063dac83d2914610831578063dc311dd314610851578063e72f6e3014610882578063ee9d2d38146108a2578063f2fde38b146108cf57600080fd5b80638402595e146105c757806386fe91c7146105e75780638da5cb5b1461060757806395b55cfc146106255780639b1c385e146106385780639d40a6fd14610658578063a21a23e414610690578063a4c0ed36146106a5578063a63e0bfb146106c5578063aa433aff146106e5578063aefb212f1461070557600080fd5b8063405b84fa116101ad578063405b84fa1461044e57806340d6bb821461046e57806341af6c871461049957806351cff8d9146104c95780635d06b4ab146104e957806364d51a2a14610509578063659827441461051e578063689c45171461053e57806372e9d5651461057257806379ba5097146105925780637a5a2aef146105a757600080fd5b806304104edb14610221578063043bd6ae14610243578063088070f51461026c57806308821d581461033a5780630ae095401461035a57806315c48b841461037a57806318e3dd27146103a25780631b6b6d23146103e15780632f622e6b1461040e578063301f42e91461042e575b600080fd5b34801561022d57600080fd5b5061024161023c366004614fd3565b6108ef565b005b34801561024f57600080fd5b5061025960105481565b6040519081526020015b60405180910390f35b34801561027857600080fd5b50600c546102dd9061ffff81169063ffffffff62010000820481169160ff600160301b8204811692600160381b8304811692600160581b8104821692600160781b8204831692600160981b83041691600160b81b8104821691600160c01b9091041689565b6040805161ffff909a168a5263ffffffff98891660208b01529615159689019690965293861660608801529185166080870152841660a08601529290921660c084015260ff91821660e08401521661010082015261012001610263565b34801561034657600080fd5b506102416103553660046150b1565b610a5c565b34801561036657600080fd5b5061024161037536600461539b565b610c06565b34801561038657600080fd5b5061038f60c881565b60405161ffff9091168152602001610263565b3480156103ae57600080fd5b50600a546103c990600160601b90046001600160601b031681565b6040516001600160601b039091168152602001610263565b3480156103ed57600080fd5b50600254610401906001600160a01b031681565b60405161026391906155cc565b34801561041a57600080fd5b50610241610429366004614fd3565b610c4e565b34801561043a57600080fd5b506103c96104493660046151b7565b610d9a565b34801561045a57600080fd5b5061024161046936600461539b565b611001565b34801561047a57600080fd5b506104846101f481565b60405163ffffffff9091168152602001610263565b3480156104a557600080fd5b506104b96104b436600461513a565b6113b3565b6040519015158152602001610263565b3480156104d557600080fd5b506102416104e4366004614fd3565b6114c3565b3480156104f557600080fd5b50610241610504366004614fd3565b611651565b34801561051557600080fd5b5061038f606481565b34801561052a57600080fd5b50610241610539366004614ff0565b611708565b34801561054a57600080fd5b506104017f000000000000000000000000000000000000000000000000000000000000000081565b34801561057e57600080fd5b50600354610401906001600160a01b031681565b34801561059e57600080fd5b50610241611768565b3480156105b357600080fd5b506102416105c23660046150cd565b611812565b3480156105d357600080fd5b506102416105e2366004614fd3565b611942565b3480156105f357600080fd5b50600a546103c9906001600160601b031681565b34801561061357600080fd5b506000546001600160a01b0316610401565b61024161063336600461513a565b611a54565b34801561064457600080fd5b506102596106533660046152a5565b611b78565b34801561066457600080fd5b50600754610678906001600160401b031681565b6040516001600160401b039091168152602001610263565b34801561069c57600080fd5b50610259611eb9565b3480156106b157600080fd5b506102416106c0366004615029565b61208c565b3480156106d157600080fd5b506102416106e03660046152fa565b612208565b3480156106f157600080fd5b5061024161070036600461513a565b612485565b34801561071157600080fd5b506107256107203660046153c0565b6124cd565b6040516102639190615643565b34801561073e57600080fd5b5061024161074d36600461513a565b6125cf565b34801561075e57600080fd5b5061024161076d36600461539b565b6126c4565b34801561077e57600080fd5b5061025961078d366004615101565b6127d0565b34801561079e57600080fd5b506102416107ad36600461539b565b612800565b3480156107be57600080fd5b506102596107cd36600461513a565b612a6f565b3480156107de57600080fd5b506108126107ed36600461513a565b600d6020526000908152604090205460ff81169061010090046001600160401b031682565b6040805192151583526001600160401b03909116602083015201610263565b34801561083d57600080fd5b5061024161084c36600461539b565b612a90565b34801561085d57600080fd5b5061087161086c36600461513a565b612b27565b604051610263959493929190615818565b34801561088e57600080fd5b5061024161089d366004614fd3565b612c15565b3480156108ae57600080fd5b506102596108bd36600461513a565b600f6020526000908152604090205481565b3480156108db57600080fd5b506102416108ea366004614fd3565b612de8565b6108f7612df9565b60115460005b81811015610a3457826001600160a01b03166011828154811061092257610922615b63565b6000918252602090912001546001600160a01b03161415610a2457601161094a600184615a13565b8154811061095a5761095a615b63565b600091825260209091200154601180546001600160a01b03909216918390811061098657610986615b63565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060118054806109c5576109c5615b4d565b600082815260209020810160001990810180546001600160a01b03191690550190556040517ff80a1a97fd42251f3c33cda98635e7399253033a6774fe37cd3f650b5282af3790610a179085906155cc565b60405180910390a1505050565b610a2d81615acb565b90506108fd565b5081604051635428d44960e01b8152600401610a5091906155cc565b60405180910390fd5b50565b610a64612df9565b604080518082018252600091610a939190849060029083908390808284376000920191909152506127d0915050565b6000818152600d602090815260409182902082518084019093525460ff811615158084526101009091046001600160401b03169183019190915291925090610af157604051631dfd6e1360e21b815260048101839052602401610a50565b6000828152600d6020526040812080546001600160481b0319169055600e54905b81811015610bc25783600e8281548110610b2e57610b2e615b63565b90600052602060002001541415610bb257600e610b4c600184615a13565b81548110610b5c57610b5c615b63565b9060005260206000200154600e8281548110610b7a57610b7a615b63565b600091825260209091200155600e805480610b9757610b97615b4d565b60019003818190600052602060002001600090559055610bc2565b610bbb81615acb565b9050610b12565b507f9b6868e0eb737bcd72205360baa6bfd0ba4e4819a33ade2db384e8a8025639a5838360200151604051610bf8929190615656565b60405180910390a150505050565b81610c1081612e4e565b610c18612eaf565b610c21836113b3565b15610c3f57604051631685ecdd60e31b815260040160405180910390fd5b610c498383612eda565b505050565b610c56612eaf565b610c5e612df9565b600b54600160601b90046001600160601b0316610c8e57604051631e9acf1760e31b815260040160405180910390fd5b600b8054600160601b90046001600160601b0316908190600c610cb18380615a4f565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600a600c8282829054906101000a90046001600160601b0316610cf99190615a4f565b92506101000a8154816001600160601b0302191690836001600160601b031602179055506000826001600160a01b0316826001600160601b031660405160006040518083038185875af1925050503d8060008114610d73576040519150601f19603f3d011682016040523d82523d6000602084013e610d78565b606091505b5050905080610c495760405163950b247960e01b815260040160405180910390fd5b6000610da4612eaf565b60005a9050610324361115610dd657604051630f28961b60e01b81523660048201526103246024820152604401610a50565b6000610de2868661308e565b90506000610df88583600001516020015161334a565b60408301516060888101519293509163ffffffff16806001600160401b03811115610e2557610e25615b79565b604051908082528060200260200182016040528015610e4e578160200160208202803683370190505b50925060005b81811015610eb65760408051602081018590529081018290526060016040516020818303038152906040528051906020012060001c848281518110610e9b57610e9b615b63565b6020908102919091010152610eaf81615acb565b9050610e54565b5050602080850180516000908152600f9092526040822082905551610edc908a85613398565b60208a8101516000908152600690915260409020805491925090601890610f1290600160c01b90046001600160401b0316615ae6565b91906101000a8154816001600160401b0302191690836001600160401b0316021790555060008960a0015160018b60a0015151610f4f9190615a13565b81518110610f5f57610f5f615b63565b60209101015160f81c6001149050610f798786838c613433565b9750610f8a88828c60200151613463565b6020808b015187820151604080518781526001600160601b038d16948101949094528415159084015284151560608401528b1515608084015290917faeb4b4786571e184246d39587f659abf0e26f41f6a3358692250382c0cdb47b79060a00160405180910390a3505050505050505b9392505050565b611009612eaf565b611012816135b6565b6110315780604051635428d44960e01b8152600401610a5091906155cc565b60008060008061104086612b27565b945094505093509350336001600160a01b0316826001600160a01b0316146110a35760405162461bcd60e51b81526020600482015260166024820152752737ba1039bab139b1b934b83a34b7b71037bbb732b960511b6044820152606401610a50565b6110ac866113b3565b156110f25760405162461bcd60e51b815260206004820152601660248201527550656e64696e6720726571756573742065786973747360501b6044820152606401610a50565b6040805160c0810182526001815260208082018990526001600160a01b03851682840152606082018490526001600160601b038088166080840152861660a08301529151909160009161114791849101615680565b604051602081830303815290604052905061116188613622565b505060405163ce3f471960e01b81526001600160a01b0388169063ce3f4719906001600160601b0388169061119a90859060040161566d565b6000604051808303818588803b1580156111b357600080fd5b505af11580156111c7573d6000803e3d6000fd5b50506002546001600160a01b0316158015935091506111f0905057506001600160601b03861615155b156112ba5760025460405163a9059cbb60e01b81526001600160a01b039091169063a9059cbb90611227908a908a90600401615613565b602060405180830381600087803b15801561124157600080fd5b505af1158015611255573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611279919061511d565b6112ba5760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b6044820152606401610a50565b600c805460ff60301b1916600160301b17905560005b8351811015611361578381815181106112eb576112eb615b63565b60200260200101516001600160a01b0316638ea98117896040518263ffffffff1660e01b815260040161131e91906155cc565b600060405180830381600087803b15801561133857600080fd5b505af115801561134c573d6000803e3d6000fd5b505050508061135a90615acb565b90506112d0565b50600c805460ff60301b191690556040517fd63ca8cb945956747ee69bfdc3ea754c24a4caf7418db70e46052f7850be4187906113a19089908b906155e0565b60405180910390a15050505050505050565b60008181526005602052604081206002018054806113d5575060009392505050565b600e5460005b828110156114b75760008482815481106113f7576113f7615b63565b60009182526020822001546001600160a01b031691505b838110156114a457600061146c600e838154811061142e5761142e615b63565b60009182526020808320909101546001600160a01b03871683526004825260408084208e855290925291205485908c906001600160401b03166137ca565b506000818152600f6020526040902054909150156114935750600198975050505050505050565b5061149d81615acb565b905061140e565b5050806114b090615acb565b90506113db565b50600095945050505050565b6114cb612eaf565b6114d3612df9565b6002546001600160a01b03166114fc5760405163c1f0c0a160e01b815260040160405180910390fd5b600b546001600160601b031661152557604051631e9acf1760e31b815260040160405180910390fd5b600b80546001600160601b031690819060006115418380615a4f565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600a60008282829054906101000a90046001600160601b03166115899190615a4f565b82546001600160601b039182166101009390930a92830291909202199091161790555060025460405163a9059cbb60e01b81526001600160a01b039091169063a9059cbb906115de9085908590600401615613565b602060405180830381600087803b1580156115f857600080fd5b505af115801561160c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611630919061511d565b61164d57604051631e9acf1760e31b815260040160405180910390fd5b5050565b611659612df9565b611662816135b6565b15611682578060405163ac8a27ef60e01b8152600401610a5091906155cc565b601180546001810182556000919091527f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c680180546001600160a01b0319166001600160a01b0383161790556040517fb7cabbfc11e66731fc77de0444614282023bcbd41d16781c753a431d0af01625906116fd9083906155cc565b60405180910390a150565b611710612df9565b6002546001600160a01b03161561173a57604051631688c53760e11b815260040160405180910390fd5b600280546001600160a01b039384166001600160a01b03199182161790915560038054929093169116179055565b6001546001600160a01b031633146117bb5760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b6044820152606401610a50565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61181a612df9565b6040805180820182526000916118499190859060029083908390808284376000920191909152506127d0915050565b6000818152600d602052604090205490915060ff161561187f57604051634a0b8fa760e01b815260048101829052602401610a50565b60408051808201825260018082526001600160401b0385811660208085019182526000878152600d9091528581209451855492516001600160481b031990931690151568ffffffffffffffff00191617610100929093169190910291909117909255600e805491820181559091527fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd01829055517f9b911b2c240bfbef3b6a8f7ed6ee321d1258bb2a3fe6becab52ac1cd3210afd390610a179083908590615656565b61194a612df9565b600a544790600160601b90046001600160601b03168181111561198a576040516354ced18160e11b81526004810182905260248101839052604401610a50565b81811015610c4957600061199e8284615a13565b90506000846001600160a01b03168260405160006040518083038185875af1925050503d80600081146119ed576040519150601f19603f3d011682016040523d82523d6000602084013e6119f2565b606091505b5050905080611a145760405163950b247960e01b815260040160405180910390fd5b7f4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c8583604051611a459291906155e0565b60405180910390a15050505050565b611a5c612eaf565b6000818152600560205260409020546001600160a01b0316611a9157604051630fb532db60e11b815260040160405180910390fd5b60008181526006602052604090208054600160601b90046001600160601b0316903490600c611ac083856159be565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555034600a600c8282829054906101000a90046001600160601b0316611b0891906159be565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550817f7603b205d03651ee812f803fccde89f1012e545a9c99f0abfea9cedd0fd8e902823484611b5b919061595f565b604080519283526020830191909152015b60405180910390a25050565b6000611b82612eaf565b602080830135600081815260059092526040909120546001600160a01b0316611bbe57604051630fb532db60e11b815260040160405180910390fd5b33600090815260046020908152604080832084845291829052909120546001600160401b031680611c065782336040516379bfd40160e01b8152600401610a509291906156f5565b600c5461ffff16611c1d60608701604088016152df565b61ffff161080611c40575060c8611c3a60608701604088016152df565b61ffff16115b15611c8657611c5560608601604087016152df565b600c5460405163539c34bb60e11b815261ffff92831660048201529116602482015260c86044820152606401610a50565b600c5462010000900463ffffffff16611ca560808701606088016153e2565b63ffffffff161115611cf557611cc160808601606087016153e2565b600c54604051637aebf00f60e11b815263ffffffff9283166004820152620100009091049091166024820152604401610a50565b6101f4611d0860a08701608088016153e2565b63ffffffff161115611d4e57611d2460a08601608087016153e2565b6040516311ce1afb60e21b815263ffffffff90911660048201526101f46024820152604401610a50565b611d5781615ae6565b90506000611d6886353386856137ca565b90955090506000611d8c611d87611d8260a08a018a61586d565b613843565b6138c0565b905085611d97613931565b86611da860808b0160608c016153e2565b611db860a08c0160808d016153e2565b3386604051602001611dd09796959493929190615778565b60405160208183030381529060405280519060200120600f600088815260200190815260200160002081905550336001600160a01b03168588600001357feb0e3652e0f44f417695e6e90f2f42c99b65cd7169074c5a654b16b9748c3a4e89868c6040016020810190611e4391906152df565b8d6060016020810190611e5691906153e2565b8e6080016020810190611e6991906153e2565b89604051611e7c96959493929190615739565b60405180910390a450506000928352602091909152604090912080546001600160401b0319166001600160401b039092169190911790555b919050565b6000611ec3612eaf565b6007546001600160401b031633611edb600143615a13565b6040516001600160601b0319606093841b81166020830152914060348201523090921b1660548201526001600160c01b031960c083901b16606882015260700160408051601f1981840301815291905280516020909101209150611f40816001615977565b600780546001600160401b0319166001600160401b03928316179055604080516000808252608082018352602080830182815283850183815260608086018581528a86526006855287862093518454935191516001600160601b039182166001600160c01b031990951694909417600160601b9190921602176001600160c01b0316600160c01b9290981691909102969096179055835194850184523385528481018281528585018481528884526005835294909220855181546001600160a01b03199081166001600160a01b03928316178355935160018301805490951691161790925592518051929493919261203e9260028501920190614ced565b5061204e915060089050846139c1565b50827f1d3015d7ba850fa198dc7b1a3f5d42779313a681035f77c8c03764c61005518d3360405161207f91906155cc565b60405180910390a2505090565b612094612eaf565b6002546001600160a01b031633146120bf576040516344b0e3c360e01b815260040160405180910390fd5b602081146120e057604051638129bbcd60e01b815260040160405180910390fd5b60006120ee8284018461513a565b6000818152600560205260409020549091506001600160a01b031661212657604051630fb532db60e11b815260040160405180910390fd5b600081815260066020526040812080546001600160601b03169186919061214d83856159be565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555084600a60008282829054906101000a90046001600160601b031661219591906159be565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550817f1ced9348ff549fceab2ac57cd3a9de38edaaab274b725ee82c23e8fc8c4eec7a8287846121e8919061595f565b6040805192835260208301919091520160405180910390a2505050505050565b612210612df9565b60c861ffff8a16111561224a5760405163539c34bb60e11b815261ffff8a1660048201819052602482015260c86044820152606401610a50565b6000851361226e576040516321ea67b360e11b815260048101869052602401610a50565b60008463ffffffff1611801561229057508363ffffffff168363ffffffff1610155b156122be576040516313c06e5960e11b815263ffffffff808516600483015285166024820152604401610a50565b604080516101208101825261ffff8b1680825263ffffffff808c16602084018190526000848601528b8216606085018190528b8316608086018190528a841660a08701819052938a1660c0870181905260ff808b1660e08901819052908a16610100909801889052600c8054600160c01b90990260ff60c01b19600160b81b9093029290921661ffff60b81b19600160981b90940263ffffffff60981b19600160781b9099029890981667ffffffffffffffff60781b19600160581b90960263ffffffff60581b19600160381b9098029790971668ffffffffffffffffff60301b196201000090990265ffffffffffff19909c16909a179a909a1796909616979097179390931791909116959095179290921793909316929092179190911790556010869055517f2c6b6b12413678366b05b145c5f00745bdd00e739131ab5de82484a50c9d78b690612472908b908b908b908b908b908b908b908b908b9061ffff99909916895263ffffffff97881660208a0152958716604089015293861660608801526080870192909252841660a086015290921660c084015260ff91821660e0840152166101008201526101200190565b60405180910390a1505050505050505050565b61248d612df9565b6000818152600560205260409020546001600160a01b0316806124c357604051630fb532db60e11b815260040160405180910390fd5b61164d8282612eda565b606060006124db60086139cd565b90508084106124fd57604051631390f2a160e01b815260040160405180910390fd5b6000612509848661595f565b905081811180612517575083155b6125215780612523565b815b905060006125318683615a13565b9050806001600160401b0381111561254b5761254b615b79565b604051908082528060200260200182016040528015612574578160200160208202803683370190505b50935060005b818110156125c45761259761258f888361595f565b6008906139d7565b8582815181106125a9576125a9615b63565b60209081029190910101526125bd81615acb565b905061257a565b505050505b92915050565b6125d7612eaf565b6000818152600560205260409020546001600160a01b03168061260d57604051630fb532db60e11b815260040160405180910390fd5b6000828152600560205260409020600101546001600160a01b03163314612664576000828152600560205260409081902060010154905163d084e97560e01b8152610a50916001600160a01b0316906004016155cc565b600082815260056020526040908190208054336001600160a01b031991821681178355600190920180549091169055905183917fd4114ab6e9af9f597c52041f32d62dc57c5c4e4c0d4427006069635e216c938691611b6c9185916155f9565b816126ce81612e4e565b6126d6612eaf565b600083815260056020526040902060020180546064141561270a576040516305a48e0f60e01b815260040160405180910390fd5b6001600160a01b038316600090815260046020908152604080832087845291829052909120546001600160401b031615612745575050505050565b600085815260208281526040808320805460016001600160401b0319909116811790915585549081018655858452919092200180546001600160a01b0319166001600160a01b0387161790555185907f1e980d04aa7648e205713e5e8ea3808672ac163d10936d36f91b2c88ac1575e1906127c19087906155cc565b60405180910390a25050505050565b6000816040516020016127e39190615635565b604051602081830303815290604052805190602001209050919050565b8161280a81612e4e565b612812612eaf565b61281b836113b3565b1561283957604051631685ecdd60e31b815260040160405180910390fd5b6001600160a01b03821660009081526004602090815260408083208684529091529020546001600160401b03166128875782826040516379bfd40160e01b8152600401610a509291906156f5565b6000838152600560209081526040808320600201805482518185028101850190935280835291929091908301828280156128ea57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116128cc575b505050505090506000600182516129019190615a13565b905060005b8251811015612a0b57846001600160a01b031683828151811061292b5761292b615b63565b60200260200101516001600160a01b031614156129fb57600083838151811061295657612956615b63565b602002602001015190508060056000898152602001908152602001600020600201838154811061298857612988615b63565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394909416939093179092558881526005909152604090206002018054806129d3576129d3615b4d565b600082815260209020810160001990810180546001600160a01b031916905501905550612a0b565b612a0481615acb565b9050612906565b506001600160a01b03841660009081526004602090815260408083208884529091529081902080546001600160401b03191690555185907f32158c6058347c1601b2d12bc696ac6901d8a9a9aa3ba10c27ab0a983e8425a7906127c19087906155cc565b600e8181548110612a7f57600080fd5b600091825260209091200154905081565b81612a9a81612e4e565b612aa2612eaf565b600083815260056020526040902060018101546001600160a01b03848116911614612b21576001810180546001600160a01b0319166001600160a01b03851617905560405184907f21a4dad170a6bf476c31bbcf4a16628295b0e450672eec25d7c93308e05344a190612b1890339087906155f9565b60405180910390a25b50505050565b600081815260056020526040812054819081906001600160a01b0316606081612b6357604051630fb532db60e11b815260040160405180910390fd5b600086815260066020908152604080832054600583529281902060020180548251818502810185019093528083526001600160601b0380861695600160601b810490911694600160c01b9091046001600160401b0316938893929091839190830182828015612bfb57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612bdd575b505050505090509450945094509450945091939590929450565b612c1d612df9565b6002546001600160a01b0316612c465760405163c1f0c0a160e01b815260040160405180910390fd5b6002546040516370a0823160e01b81526000916001600160a01b0316906370a0823190612c779030906004016155cc565b60206040518083038186803b158015612c8f57600080fd5b505afa158015612ca3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cc79190615153565b600a549091506001600160601b031681811115612d01576040516354ced18160e11b81526004810182905260248101839052604401610a50565b81811015610c49576000612d158284615a13565b60025460405163a9059cbb60e01b81529192506001600160a01b03169063a9059cbb90612d4890879085906004016155e0565b602060405180830381600087803b158015612d6257600080fd5b505af1158015612d76573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d9a919061511d565b612db757604051631f01ff1360e21b815260040160405180910390fd5b7f59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b4366008482604051610bf89291906155e0565b612df0612df9565b610a59816139e3565b6000546001600160a01b03163314612e4c5760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b6044820152606401610a50565b565b6000818152600560205260409020546001600160a01b031680612e8457604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b0382161461164d5780604051636c51fda960e11b8152600401610a5091906155cc565b600c54600160301b900460ff1615612e4c5760405163769dd35360e11b815260040160405180910390fd5b600080612ee684613622565b60025491935091506001600160a01b031615801590612f0d57506001600160601b03821615155b15612fbc5760025460405163a9059cbb60e01b81526001600160a01b039091169063a9059cbb90612f4d9086906001600160601b038716906004016155e0565b602060405180830381600087803b158015612f6757600080fd5b505af1158015612f7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f9f919061511d565b612fbc57604051631e9acf1760e31b815260040160405180910390fd5b6000836001600160a01b0316826001600160601b031660405160006040518083038185875af1925050503d8060008114613012576040519150601f19603f3d011682016040523d82523d6000602084013e613017565b606091505b50509050806130395760405163950b247960e01b815260040160405180910390fd5b604080516001600160a01b03861681526001600160601b03808616602083015284169181019190915285907f8c74ce8b8cf87f5eb001275c8be27eb34ea2b62bfab6814fcc62192bb63e81c4906060016127c1565b6040805160a081018252600060608201818152608083018290528252602082018190529181019190915260006130c784600001516127d0565b6000818152600d602090815260409182902082518084019093525460ff811615158084526101009091046001600160401b0316918301919091529192509061312557604051631dfd6e1360e21b815260048101839052602401610a50565b6000828660800151604051602001613147929190918252602082015260400190565b60408051601f1981840301815291815281516020928301206000818152600f9093529120549091508061318d57604051631b44092560e11b815260040160405180910390fd5b85516020808801516040808a015160608b015160808c015160a08d015193516131bc978a9790969591016157c4565b6040516020818303038152906040528051906020012081146131f15760405163354a450b60e21b815260040160405180910390fd5b60006132008760000151613a87565b9050806132d8578651604051631d2827a760e31b81526001600160401b0390911660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063e9413d389060240160206040518083038186803b15801561327257600080fd5b505afa158015613286573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132aa9190615153565b9050806132d857865160405163175dadad60e01b81526001600160401b039091166004820152602401610a50565b60008860800151826040516020016132fa929190918252602082015260400190565b6040516020818303038152906040528051906020012060001c905060006133218a83613b69565b604080516060810182529788526020880196909652948601949094525092979650505050505050565b6000816001600160401b03163a111561339057821561337357506001600160401b0381166125c9565b3a8260405163435e532d60e11b8152600401610a50929190615656565b503a92915050565b6000806000631fe543e360e01b86856040516024016133b8929190615720565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252600c805460ff60301b1916600160301b17905590860151608087015191925061341c9163ffffffff9091169083613bd4565b600c805460ff60301b191690559695505050505050565b6000821561344d57613446858584613c20565b905061345b565b613458858584613d31565b90505b949350505050565b600081815260066020526040902082156135225780546001600160601b03600160601b90910481169085168110156134ae57604051631e9acf1760e31b815260040160405180910390fd5b6134b88582615a4f565b8254600160601b600160c01b031916600160601b6001600160601b039283168102919091178455600b805488939192600c926134f89286929004166159be565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555050612b21565b80546001600160601b0390811690851681101561355257604051631e9acf1760e31b815260040160405180910390fd5b61355c8582615a4f565b82546001600160601b0319166001600160601b03918216178355600b8054879260009161358b918591166159be565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505050505050565b601154600090815b8181101561361857836001600160a01b0316601182815481106135e3576135e3615b63565b6000918252602090912001546001600160a01b03161415613608575060019392505050565b61361181615acb565b90506135be565b5060009392505050565b60008181526005602090815260408083206006909252822054600290910180546001600160601b0380841694600160601b90940416925b818110156136c4576004600084838154811061367757613677615b63565b60009182526020808320909101546001600160a01b031683528281019390935260409182018120898252909252902080546001600160401b03191690556136bd81615acb565b9050613659565b50600085815260056020526040812080546001600160a01b031990811682556001820180549091169055906136fc6002830182614d52565b5050600085815260066020526040812055613718600886613f1a565b506001600160601b0384161561376b57600a80548591906000906137469084906001600160601b0316615a4f565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505b6001600160601b038316156137c35782600a600c8282829054906101000a90046001600160601b031661379e9190615a4f565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505b5050915091565b6040805160208082018790526001600160a01b03959095168183015260608101939093526001600160401b03919091166080808401919091528151808403909101815260a08301825280519084012060c083019490945260e0808301859052815180840390910181526101009092019052805191012091565b6040805160208101909152600081528161386c57506040805160208101909152600081526125c9565b63125fa26760e31b61387e8385615a6f565b6001600160e01b031916146138a657604051632923fee760e11b815260040160405180910390fd5b6138b38260048186615935565b810190610ffa919061516c565b60607f92fd13387c7fe7befbc38d303d6468778fb9731bc4583f17d92989c6fcfdeaaa826040516024016138f991511515815260200190565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915292915050565b60004661393d81613f26565b156139ba5760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561397c57600080fd5b505afa158015613990573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139b49190615153565b91505090565b4391505090565b6000610ffa8383613f49565b60006125c9825490565b6000610ffa8383613f98565b6001600160a01b038116331415613a365760405162461bcd60e51b815260206004820152601760248201527621b0b73737ba103a3930b739b332b9103a379039b2b63360491b6044820152606401610a50565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600046613a9381613f26565b15613b5a57610100836001600160401b0316613aad613931565b613ab79190615a13565b1180613ad35750613ac6613931565b836001600160401b031610155b15613ae15750600092915050565b6040516315a03d4160e11b81526001600160401b0384166004820152606490632b407a82906024015b60206040518083038186803b158015613b2257600080fd5b505afa158015613b36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ffa9190615153565b50506001600160401b03164090565b6000613b9d8360000151846020015185604001518660600151868860a001518960c001518a60e001518b6101000151613fc2565b60038360200151604051602001613bb592919061570c565b60408051601f1981840301815291905280516020909101209392505050565b60005a611388811015613be657600080fd5b611388810390508460408204820311613bfe57600080fd5b50823b613c0a57600080fd5b60008083516020850160008789f1949350505050565b600080613c636000368080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506141de92505050565b905060005a600c54613c83908890600160581b900463ffffffff1661595f565b613c8d9190615a13565b613c9790866159f4565b600c54909150600090613cbc90600160781b900463ffffffff1664e8d4a510006159f4565b90508415613d0857600c548190606490600160b81b900460ff16613ce0858761595f565b613cea91906159f4565b613cf491906159e0565b613cfe919061595f565b9350505050610ffa565b600c548190606490613d2490600160b81b900460ff1682615999565b60ff16613ce0858761595f565b600080613d3c6142ac565b905060008113613d62576040516321ea67b360e11b815260048101829052602401610a50565b6000613da46000368080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506141de92505050565b9050600082825a600c54613dc6908b90600160581b900463ffffffff1661595f565b613dd09190615a13565b613dda90896159f4565b613de4919061595f565b613df690670de0b6b3a76400006159f4565b613e0091906159e0565b600c54909150600090613e299063ffffffff600160981b8204811691600160781b900416615a2a565b613e3e9063ffffffff1664e8d4a510006159f4565b9050600084613e5583670de0b6b3a76400006159f4565b613e5f91906159e0565b905060008715613ea057600c548290606490613e8590600160c01b900460ff16876159f4565b613e8f91906159e0565b613e99919061595f565b9050613ee0565b600c548290606490613ebc90600160c01b900460ff1682615999565b613ec99060ff16876159f4565b613ed391906159e0565b613edd919061595f565b90505b6b033b2e3c9fd0803ce8000000811115613f0d5760405163e80fa38160e01b815260040160405180910390fd5b9998505050505050505050565b6000610ffa838361437b565b600061a4b1821480613f3a575062066eed82145b806125c957505062066eee1490565b6000818152600183016020526040812054613f90575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556125c9565b5060006125c9565b6000826000018281548110613faf57613faf615b63565b9060005260206000200154905092915050565b613fcb8961446e565b6140145760405162461bcd60e51b815260206004820152601a6024820152797075626c6963206b6579206973206e6f74206f6e20637572766560301b6044820152606401610a50565b61401d8861446e565b6140615760405162461bcd60e51b815260206004820152601560248201527467616d6d61206973206e6f74206f6e20637572766560581b6044820152606401610a50565b61406a8361446e565b6140b65760405162461bcd60e51b815260206004820152601d60248201527f6347616d6d615769746e657373206973206e6f74206f6e2063757276650000006044820152606401610a50565b6140bf8261446e565b61410b5760405162461bcd60e51b815260206004820152601c60248201527f73486173685769746e657373206973206e6f74206f6e206375727665000000006044820152606401610a50565b614117878a8887614531565b61415f5760405162461bcd60e51b81526020600482015260196024820152786164647228632a706b2b732a6729213d5f755769746e65737360381b6044820152606401610a50565b600061416b8a87614654565b9050600061417e898b878b8689896146b8565b9050600061418f838d8d8a866147d7565b9050808a146141d05760405162461bcd60e51b815260206004820152600d60248201526c34b73b30b634b210383937b7b360991b6044820152606401610a50565b505050505050505050505050565b6000466141ea81613f26565b1561422957606c6001600160a01b031663c6f7de0e6040518163ffffffff1660e01b815260040160206040518083038186803b158015613b2257600080fd5b61423281614817565b156142a357600f602160991b016001600160a01b03166349948e0e84604051806080016040528060488152602001615bb360489139604051602001614278929190615522565b6040516020818303038152906040526040518263ffffffff1660e01b8152600401613b0a919061566d565b50600092915050565b600c5460035460408051633fabe5a360e21b81529051600093600160381b900463ffffffff169284926001600160a01b039091169163feaf968c9160048082019260a092909190829003018186803b15801561430757600080fd5b505afa15801561431b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061433f91906153fd565b50919550909250505063ffffffff82161580159061436b57506143628142615a13565b8263ffffffff16105b156143765760105492505b505090565b6000818152600183016020526040812054801561446457600061439f600183615a13565b85549091506000906143b390600190615a13565b90508181146144185760008660000182815481106143d3576143d3615b63565b90600052602060002001549050808760000184815481106143f6576143f6615b63565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061442957614429615b4d565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506125c9565b60009150506125c9565b80516000906401000003d019116144bc5760405162461bcd60e51b8152602060048201526012602482015271696e76616c696420782d6f7264696e61746560701b6044820152606401610a50565b60208201516401000003d0191161450a5760405162461bcd60e51b8152602060048201526012602482015271696e76616c696420792d6f7264696e61746560701b6044820152606401610a50565b60208201516401000003d01990800961452a8360005b6020020151614851565b1492915050565b60006001600160a01b0382166145775760405162461bcd60e51b815260206004820152600b60248201526a626164207769746e65737360a81b6044820152606401610a50565b60208401516000906001161561458e57601c614591565b601b5b9050600070014551231950b75fc4402da1732fc9bebe1985876000602002015109865170014551231950b75fc4402da1732fc9bebe19918203925060009190890987516040805160008082526020820180845287905260ff88169282019290925260608101929092526080820183905291925060019060a0016020604051602081039080840390855afa15801561462c573d6000803e3d6000fd5b5050604051601f1901516001600160a01b039081169088161495505050505050949350505050565b61465c614d70565b61468960018484604051602001614675939291906155ab565b604051602081830303815290604052614875565b90505b6146958161446e565b6125c95780516040805160208101929092526146b19101614675565b905061468c565b6146c0614d70565b825186516401000003d019908190069106141561471f5760405162461bcd60e51b815260206004820152601e60248201527f706f696e747320696e2073756d206d7573742062652064697374696e637400006044820152606401610a50565b61472a8789886148c3565b61476f5760405162461bcd60e51b8152602060048201526016602482015275119a5c9cdd081b5d5b0818da1958dac819985a5b195960521b6044820152606401610a50565b61477a8486856148c3565b6147c05760405162461bcd60e51b815260206004820152601760248201527614d958dbdb99081b5d5b0818da1958dac819985a5b1959604a1b6044820152606401610a50565b6147cb8684846149eb565b98975050505050505050565b6000600286868685876040516020016147f596959493929190615551565b60408051601f1981840301815291905280516020909101209695505050505050565b6000600a82148061482957506101a482145b80614836575062aa37dc82145b80614842575061210582145b806125c957505062014a331490565b6000806401000003d01980848509840990506401000003d019600782089392505050565b61487d614d70565b61488682614aae565b815261489b614896826000614520565b614ae9565b602082018190526002900660011415611eb4576020810180516401000003d019039052919050565b6000826149005760405162461bcd60e51b815260206004820152600b60248201526a3d32b9379039b1b0b630b960a91b6044820152606401610a50565b8351602085015160009061491690600290615b0d565b1561492257601c614925565b601b5b9050600070014551231950b75fc4402da1732fc9bebe198387096040805160008082526020820180845281905260ff86169282019290925260608101869052608081018390529192509060019060a0016020604051602081039080840390855afa158015614997573d6000803e3d6000fd5b5050506020604051035190506000866040516020016149b69190615510565b60408051601f1981840301815291905280516020909101206001600160a01b0392831692169190911498975050505050505050565b6149f3614d70565b835160208086015185519186015160009384938493614a1493909190614b09565b919450925090506401000003d019858209600114614a705760405162461bcd60e51b815260206004820152601960248201527834b73b2d1036bab9ba1031329034b73b32b939b29037b3103d60391b6044820152606401610a50565b60405180604001604052806401000003d01980614a8f57614a8f615b37565b87860981526020016401000003d0198785099052979650505050505050565b805160208201205b6401000003d0198110611eb457604080516020808201939093528151808203840181529082019091528051910120614ab6565b60006125c9826002614b026401000003d019600161595f565b901c614be9565b60008080600180826401000003d019896401000003d019038808905060006401000003d0198b6401000003d019038a0890506000614b4983838585614c80565b9098509050614b5a88828e88614ca4565b9098509050614b6b88828c87614ca4565b90985090506000614b7e8d878b85614ca4565b9098509050614b8f88828686614c80565b9098509050614ba088828e89614ca4565b9098509050818114614bd5576401000003d019818a0998506401000003d01982890997506401000003d0198183099650614bd9565b8196505b5050505050509450945094915050565b600080614bf4614d8e565b6020808252818101819052604082015260608101859052608081018490526401000003d01960a0820152614c26614dac565b60208160c0846005600019fa925082614c765760405162461bcd60e51b81526020600482015260126024820152716269674d6f64457870206661696c7572652160701b6044820152606401610a50565b5195945050505050565b6000806401000003d0198487096401000003d0198487099097909650945050505050565b600080806401000003d019878509905060006401000003d01987876401000003d019030990506401000003d0198183086401000003d01986890990999098509650505050505050565b828054828255906000526020600020908101928215614d42579160200282015b82811115614d4257825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614d0d565b50614d4e929150614dca565b5090565b5080546000825590600052602060002090810190610a599190614dca565b60405180604001604052806002906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b5b80821115614d4e5760008155600101614dcb565b8035611eb481615b8f565b80604081018310156125c957600080fd5b600082601f830112614e0c57600080fd5b604051604081018181106001600160401b0382111715614e2e57614e2e615b79565b8060405250808385604086011115614e4557600080fd5b60005b6002811015614e67578135835260209283019290910190600101614e48565b509195945050505050565b8035611eb481615ba4565b600060c08284031215614e8f57600080fd5b614e976158ba565b9050614ea282614f91565b815260208083013581830152614eba60408401614f7d565b6040830152614ecb60608401614f7d565b60608301526080830135614ede81615b8f565b608083015260a08301356001600160401b0380821115614efd57600080fd5b818501915085601f830112614f1157600080fd5b813581811115614f2357614f23615b79565b614f35601f8201601f19168501615905565b91508082528684828501011115614f4b57600080fd5b80848401858401376000848284010152508060a085015250505092915050565b803561ffff81168114611eb457600080fd5b803563ffffffff81168114611eb457600080fd5b80356001600160401b0381168114611eb457600080fd5b803560ff81168114611eb457600080fd5b805169ffffffffffffffffffff81168114611eb457600080fd5b600060208284031215614fe557600080fd5b8135610ffa81615b8f565b6000806040838503121561500357600080fd5b823561500e81615b8f565b9150602083013561501e81615b8f565b809150509250929050565b6000806000806060858703121561503f57600080fd5b843561504a81615b8f565b93506020850135925060408501356001600160401b038082111561506d57600080fd5b818701915087601f83011261508157600080fd5b81358181111561509057600080fd5b8860208285010111156150a257600080fd5b95989497505060200194505050565b6000604082840312156150c357600080fd5b610ffa8383614dea565b600080606083850312156150e057600080fd5b6150ea8484614dea565b91506150f860408401614f91565b90509250929050565b60006040828403121561511357600080fd5b610ffa8383614dfb565b60006020828403121561512f57600080fd5b8151610ffa81615ba4565b60006020828403121561514c57600080fd5b5035919050565b60006020828403121561516557600080fd5b5051919050565b60006020828403121561517e57600080fd5b604051602081018181106001600160401b03821117156151a0576151a0615b79565b60405282356151ae81615ba4565b81529392505050565b60008060008385036101e08112156151ce57600080fd5b6101a0808212156151de57600080fd5b6151e66158e2565b91506151f28787614dfb565b82526152018760408801614dfb565b60208301526080860135604083015260a0860135606083015260c0860135608083015261523060e08701614ddf565b60a083015261010061524488828901614dfb565b60c0840152615257886101408901614dfb565b60e0840152610180870135908301529093508401356001600160401b0381111561528057600080fd5b61528c86828701614e7d565b92505061529c6101c08501614e72565b90509250925092565b6000602082840312156152b757600080fd5b81356001600160401b038111156152cd57600080fd5b820160c08185031215610ffa57600080fd5b6000602082840312156152f157600080fd5b610ffa82614f6b565b60008060008060008060008060006101208a8c03121561531957600080fd5b6153228a614f6b565b985061533060208b01614f7d565b975061533e60408b01614f7d565b965061534c60608b01614f7d565b955060808a0135945061536160a08b01614f7d565b935061536f60c08b01614f7d565b925061537d60e08b01614fa8565b915061538c6101008b01614fa8565b90509295985092959850929598565b600080604083850312156153ae57600080fd5b82359150602083013561501e81615b8f565b600080604083850312156153d357600080fd5b50508035926020909101359150565b6000602082840312156153f457600080fd5b610ffa82614f7d565b600080600080600060a0868803121561541557600080fd5b61541e86614fb9565b945060208601519350604086015192506060860151915061544160808701614fb9565b90509295509295909350565b600081518084526020808501945080840160005b838110156154865781516001600160a01b031687529582019590820190600101615461565b509495945050505050565b8060005b6002811015612b21578151845260209384019390910190600101615495565b600081518084526020808501945080840160005b83811015615486578151875295820195908201906001016154c8565b600081518084526154fc816020860160208601615a9f565b601f01601f19169290920160200192915050565b61551a8183615491565b604001919050565b60008351615534818460208801615a9f565b835190830190615548818360208801615a9f565b01949350505050565b8681526155616020820187615491565b61556e6060820186615491565b61557b60a0820185615491565b61558860e0820184615491565b60609190911b6001600160601b0319166101208201526101340195945050505050565b8381526155bb6020820184615491565b606081019190915260800192915050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b039290921682526001600160601b0316602082015260400190565b604081016125c98284615491565b602081526000610ffa60208301846154b4565b9182526001600160401b0316602082015260400190565b602081526000610ffa60208301846154e4565b6020815260ff82511660208201526020820151604082015260018060a01b0360408301511660608201526000606083015160c060808401526156c560e084018261544d565b60808501516001600160601b0390811660a0868101919091529095015190941660c0909301929092525090919050565b9182526001600160a01b0316602082015260400190565b82815260608101610ffa6020830184615491565b82815260406020820152600061345b60408301846154b4565b86815285602082015261ffff85166040820152600063ffffffff808616606084015280851660808401525060c060a08301526147cb60c08301846154e4565b878152602081018790526040810186905263ffffffff8581166060830152841660808201526001600160a01b03831660a082015260e060c08201819052600090613f0d908301846154e4565b8781526001600160401b03871660208201526040810186905263ffffffff8581166060830152841660808201526001600160a01b03831660a082015260e060c08201819052600090613f0d908301846154e4565b6001600160601b038681168252851660208201526001600160401b03841660408201526001600160a01b038316606082015260a0608082018190526000906158629083018461544d565b979650505050505050565b6000808335601e1984360301811261588457600080fd5b8301803591506001600160401b0382111561589e57600080fd5b6020019150368190038213156158b357600080fd5b9250929050565b60405160c081016001600160401b03811182821017156158dc576158dc615b79565b60405290565b60405161012081016001600160401b03811182821017156158dc576158dc615b79565b604051601f8201601f191681016001600160401b038111828210171561592d5761592d615b79565b604052919050565b6000808585111561594557600080fd5b8386111561595257600080fd5b5050820193919092039150565b6000821982111561597257615972615b21565b500190565b60006001600160401b0380831681851680830382111561554857615548615b21565b600060ff821660ff84168060ff038211156159b6576159b6615b21565b019392505050565b60006001600160601b0382811684821680830382111561554857615548615b21565b6000826159ef576159ef615b37565b500490565b6000816000190483118215151615615a0e57615a0e615b21565b500290565b600082821015615a2557615a25615b21565b500390565b600063ffffffff83811690831681811015615a4757615a47615b21565b039392505050565b60006001600160601b0383811690831681811015615a4757615a47615b21565b6001600160e01b03198135818116916004851015615a975780818660040360031b1b83161692505b505092915050565b60005b83811015615aba578181015183820152602001615aa2565b83811115612b215750506000910152565b6000600019821415615adf57615adf615b21565b5060010190565b60006001600160401b0380831681811415615b0357615b03615b21565b6001019392505050565b600082615b1c57615b1c615b37565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610a5957600080fd5b8015158114610a5957600080fdfe307866666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666a164736f6c6343000806000a",
}

var VRFCoordinatorV25ABI = VRFCoordinatorV25MetaData.ABI

var VRFCoordinatorV25Bin = VRFCoordinatorV25MetaData.Bin

func DeployVRFCoordinatorV25(auth *bind.TransactOpts, backend bind.ContractBackend, blockhashStore common.Address) (common.Address, *types.Transaction, *VRFCoordinatorV25, error) {
	parsed, err := VRFCoordinatorV25MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFCoordinatorV25Bin), backend, blockhashStore)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFCoordinatorV25{address: address, abi: *parsed, VRFCoordinatorV25Caller: VRFCoordinatorV25Caller{contract: contract}, VRFCoordinatorV25Transactor: VRFCoordinatorV25Transactor{contract: contract}, VRFCoordinatorV25Filterer: VRFCoordinatorV25Filterer{contract: contract}}, nil
}

type VRFCoordinatorV25 struct {
	address common.Address
	abi     abi.ABI
	VRFCoordinatorV25Caller
	VRFCoordinatorV25Transactor
	VRFCoordinatorV25Filterer
}

type VRFCoordinatorV25Caller struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV25Transactor struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV25Filterer struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV25Session struct {
	Contract     *VRFCoordinatorV25
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorV25CallerSession struct {
	Contract *VRFCoordinatorV25Caller
	CallOpts bind.CallOpts
}

type VRFCoordinatorV25TransactorSession struct {
	Contract     *VRFCoordinatorV25Transactor
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorV25Raw struct {
	Contract *VRFCoordinatorV25
}

type VRFCoordinatorV25CallerRaw struct {
	Contract *VRFCoordinatorV25Caller
}

type VRFCoordinatorV25TransactorRaw struct {
	Contract *VRFCoordinatorV25Transactor
}

func NewVRFCoordinatorV25(address common.Address, backend bind.ContractBackend) (*VRFCoordinatorV25, error) {
	abi, err := abi.JSON(strings.NewReader(VRFCoordinatorV25ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFCoordinatorV25(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25{address: address, abi: abi, VRFCoordinatorV25Caller: VRFCoordinatorV25Caller{contract: contract}, VRFCoordinatorV25Transactor: VRFCoordinatorV25Transactor{contract: contract}, VRFCoordinatorV25Filterer: VRFCoordinatorV25Filterer{contract: contract}}, nil
}

func NewVRFCoordinatorV25Caller(address common.Address, caller bind.ContractCaller) (*VRFCoordinatorV25Caller, error) {
	contract, err := bindVRFCoordinatorV25(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25Caller{contract: contract}, nil
}

func NewVRFCoordinatorV25Transactor(address common.Address, transactor bind.ContractTransactor) (*VRFCoordinatorV25Transactor, error) {
	contract, err := bindVRFCoordinatorV25(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25Transactor{contract: contract}, nil
}

func NewVRFCoordinatorV25Filterer(address common.Address, filterer bind.ContractFilterer) (*VRFCoordinatorV25Filterer, error) {
	contract, err := bindVRFCoordinatorV25(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25Filterer{contract: contract}, nil
}

func bindVRFCoordinatorV25(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFCoordinatorV25MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV25.Contract.VRFCoordinatorV25Caller.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.VRFCoordinatorV25Transactor.contract.Transfer(opts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.VRFCoordinatorV25Transactor.contract.Transact(opts, method, params...)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV25.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.contract.Transfer(opts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.contract.Transact(opts, method, params...)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "BLOCKHASH_STORE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) BLOCKHASHSTORE() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.BLOCKHASHSTORE(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) BLOCKHASHSTORE() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.BLOCKHASHSTORE(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) LINK() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.LINK(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) LINK() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.LINK(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) LINKNATIVEFEED(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "LINK_NATIVE_FEED")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) LINKNATIVEFEED() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.LINKNATIVEFEED(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) LINKNATIVEFEED() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.LINKNATIVEFEED(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinatorV25.Contract.MAXCONSUMERS(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinatorV25.Contract.MAXCONSUMERS(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) MAXNUMWORDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "MAX_NUM_WORDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) MAXNUMWORDS() (uint32, error) {
	return _VRFCoordinatorV25.Contract.MAXNUMWORDS(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) MAXNUMWORDS() (uint32, error) {
	return _VRFCoordinatorV25.Contract.MAXNUMWORDS(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "MAX_REQUEST_CONFIRMATIONS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinatorV25.Contract.MAXREQUESTCONFIRMATIONS(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinatorV25.Contract.MAXREQUESTCONFIRMATIONS(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) GetActiveSubscriptionIds(opts *bind.CallOpts, startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "getActiveSubscriptionIds", startIndex, maxCount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) GetActiveSubscriptionIds(startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	return _VRFCoordinatorV25.Contract.GetActiveSubscriptionIds(&_VRFCoordinatorV25.CallOpts, startIndex, maxCount)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) GetActiveSubscriptionIds(startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	return _VRFCoordinatorV25.Contract.GetActiveSubscriptionIds(&_VRFCoordinatorV25.CallOpts, startIndex, maxCount)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) GetSubscription(opts *bind.CallOpts, subId *big.Int) (GetSubscription,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(GetSubscription)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NativeBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ReqCount = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Owner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[4], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) GetSubscription(subId *big.Int) (GetSubscription,

	error) {
	return _VRFCoordinatorV25.Contract.GetSubscription(&_VRFCoordinatorV25.CallOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) GetSubscription(subId *big.Int) (GetSubscription,

	error) {
	return _VRFCoordinatorV25.Contract.GetSubscription(&_VRFCoordinatorV25.CallOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "hashOfKey", publicKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinatorV25.Contract.HashOfKey(&_VRFCoordinatorV25.CallOpts, publicKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinatorV25.Contract.HashOfKey(&_VRFCoordinatorV25.CallOpts, publicKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) Owner() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.Owner(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) Owner() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.Owner(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) PendingRequestExists(opts *bind.CallOpts, subId *big.Int) (bool, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "pendingRequestExists", subId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) PendingRequestExists(subId *big.Int) (bool, error) {
	return _VRFCoordinatorV25.Contract.PendingRequestExists(&_VRFCoordinatorV25.CallOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) PendingRequestExists(subId *big.Int) (bool, error) {
	return _VRFCoordinatorV25.Contract.PendingRequestExists(&_VRFCoordinatorV25.CallOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SConfig(opts *bind.CallOpts) (SConfig,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_config")

	outstruct := new(SConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumRequestConfirmations = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.MaxGasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ReentrancyLock = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.StalenessSeconds = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.GasAfterPaymentCalculation = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeNativePPM = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeLinkDiscountPPM = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.NativePremiumPercentage = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.LinkPremiumPercentage = *abi.ConvertType(out[8], new(uint8)).(*uint8)

	return *outstruct, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SConfig() (SConfig,

	error) {
	return _VRFCoordinatorV25.Contract.SConfig(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SConfig() (SConfig,

	error) {
	return _VRFCoordinatorV25.Contract.SConfig(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SCurrentSubNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_currentSubNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SCurrentSubNonce() (uint64, error) {
	return _VRFCoordinatorV25.Contract.SCurrentSubNonce(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SCurrentSubNonce() (uint64, error) {
	return _VRFCoordinatorV25.Contract.SCurrentSubNonce(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_fallbackWeiPerUnitLink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VRFCoordinatorV25.Contract.SFallbackWeiPerUnitLink(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VRFCoordinatorV25.Contract.SFallbackWeiPerUnitLink(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SProvingKeyHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_provingKeyHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SProvingKeyHashes(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV25.Contract.SProvingKeyHashes(&_VRFCoordinatorV25.CallOpts, arg0)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SProvingKeyHashes(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV25.Contract.SProvingKeyHashes(&_VRFCoordinatorV25.CallOpts, arg0)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SProvingKeys(opts *bind.CallOpts, arg0 [32]byte) (SProvingKeys,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_provingKeys", arg0)

	outstruct := new(SProvingKeys)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.MaxGas = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SProvingKeys(arg0 [32]byte) (SProvingKeys,

	error) {
	return _VRFCoordinatorV25.Contract.SProvingKeys(&_VRFCoordinatorV25.CallOpts, arg0)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SProvingKeys(arg0 [32]byte) (SProvingKeys,

	error) {
	return _VRFCoordinatorV25.Contract.SProvingKeys(&_VRFCoordinatorV25.CallOpts, arg0)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SRequestCommitments(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_requestCommitments", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SRequestCommitments(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV25.Contract.SRequestCommitments(&_VRFCoordinatorV25.CallOpts, arg0)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SRequestCommitments(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV25.Contract.SRequestCommitments(&_VRFCoordinatorV25.CallOpts, arg0)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) STotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_totalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) STotalBalance() (*big.Int, error) {
	return _VRFCoordinatorV25.Contract.STotalBalance(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) STotalBalance() (*big.Int, error) {
	return _VRFCoordinatorV25.Contract.STotalBalance(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) STotalNativeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_totalNativeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) STotalNativeBalance() (*big.Int, error) {
	return _VRFCoordinatorV25.Contract.STotalNativeBalance(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) STotalNativeBalance() (*big.Int, error) {
	return _VRFCoordinatorV25.Contract.STotalNativeBalance(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "acceptOwnership")
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.AcceptOwnership(&_VRFCoordinatorV25.TransactOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.AcceptOwnership(&_VRFCoordinatorV25.TransactOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV25.TransactOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV25.TransactOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) AddConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "addConsumer", subId, consumer)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.AddConsumer(&_VRFCoordinatorV25.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.AddConsumer(&_VRFCoordinatorV25.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) CancelSubscription(opts *bind.TransactOpts, subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "cancelSubscription", subId, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.CancelSubscription(&_VRFCoordinatorV25.TransactOpts, subId, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.CancelSubscription(&_VRFCoordinatorV25.TransactOpts, subId, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "createSubscription")
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.CreateSubscription(&_VRFCoordinatorV25.TransactOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.CreateSubscription(&_VRFCoordinatorV25.TransactOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) DeregisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "deregisterMigratableCoordinator", target)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) DeregisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.DeregisterMigratableCoordinator(&_VRFCoordinatorV25.TransactOpts, target)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) DeregisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.DeregisterMigratableCoordinator(&_VRFCoordinatorV25.TransactOpts, target)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "deregisterProvingKey", publicProvingKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.DeregisterProvingKey(&_VRFCoordinatorV25.TransactOpts, publicProvingKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.DeregisterProvingKey(&_VRFCoordinatorV25.TransactOpts, publicProvingKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV25RequestCommitment, onlyPremium bool) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "fulfillRandomWords", proof, rc, onlyPremium)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV25RequestCommitment, onlyPremium bool) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.FulfillRandomWords(&_VRFCoordinatorV25.TransactOpts, proof, rc, onlyPremium)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV25RequestCommitment, onlyPremium bool) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.FulfillRandomWords(&_VRFCoordinatorV25.TransactOpts, proof, rc, onlyPremium)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) FundSubscriptionWithNative(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "fundSubscriptionWithNative", subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) FundSubscriptionWithNative(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.FundSubscriptionWithNative(&_VRFCoordinatorV25.TransactOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) FundSubscriptionWithNative(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.FundSubscriptionWithNative(&_VRFCoordinatorV25.TransactOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) Migrate(opts *bind.TransactOpts, subId *big.Int, newCoordinator common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "migrate", subId, newCoordinator)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) Migrate(subId *big.Int, newCoordinator common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.Migrate(&_VRFCoordinatorV25.TransactOpts, subId, newCoordinator)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) Migrate(subId *big.Int, newCoordinator common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.Migrate(&_VRFCoordinatorV25.TransactOpts, subId, newCoordinator)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "onTokenTransfer", arg0, amount, data)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.OnTokenTransfer(&_VRFCoordinatorV25.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.OnTokenTransfer(&_VRFCoordinatorV25.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) OwnerCancelSubscription(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "ownerCancelSubscription", subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) OwnerCancelSubscription(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.OwnerCancelSubscription(&_VRFCoordinatorV25.TransactOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) OwnerCancelSubscription(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.OwnerCancelSubscription(&_VRFCoordinatorV25.TransactOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "recoverFunds", to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RecoverFunds(&_VRFCoordinatorV25.TransactOpts, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RecoverFunds(&_VRFCoordinatorV25.TransactOpts, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RecoverNativeFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "recoverNativeFunds", to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RecoverNativeFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RecoverNativeFunds(&_VRFCoordinatorV25.TransactOpts, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RecoverNativeFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RecoverNativeFunds(&_VRFCoordinatorV25.TransactOpts, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RegisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "registerMigratableCoordinator", target)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RegisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RegisterMigratableCoordinator(&_VRFCoordinatorV25.TransactOpts, target)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RegisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RegisterMigratableCoordinator(&_VRFCoordinatorV25.TransactOpts, target)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RegisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int, maxGas uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "registerProvingKey", publicProvingKey, maxGas)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RegisterProvingKey(publicProvingKey [2]*big.Int, maxGas uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RegisterProvingKey(&_VRFCoordinatorV25.TransactOpts, publicProvingKey, maxGas)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RegisterProvingKey(publicProvingKey [2]*big.Int, maxGas uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RegisterProvingKey(&_VRFCoordinatorV25.TransactOpts, publicProvingKey, maxGas)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RemoveConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "removeConsumer", subId, consumer)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RemoveConsumer(&_VRFCoordinatorV25.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RemoveConsumer(&_VRFCoordinatorV25.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RequestRandomWords(opts *bind.TransactOpts, req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "requestRandomWords", req)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RequestRandomWords(req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RequestRandomWords(&_VRFCoordinatorV25.TransactOpts, req)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RequestRandomWords(req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RequestRandomWords(&_VRFCoordinatorV25.TransactOpts, req)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV25.TransactOpts, subId, newOwner)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV25.TransactOpts, subId, newOwner)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, fulfillmentFlatFeeNativePPM uint32, fulfillmentFlatFeeLinkDiscountPPM uint32, nativePremiumPercentage uint8, linkPremiumPercentage uint8) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "setConfig", minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, fulfillmentFlatFeeNativePPM, fulfillmentFlatFeeLinkDiscountPPM, nativePremiumPercentage, linkPremiumPercentage)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, fulfillmentFlatFeeNativePPM uint32, fulfillmentFlatFeeLinkDiscountPPM uint32, nativePremiumPercentage uint8, linkPremiumPercentage uint8) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.SetConfig(&_VRFCoordinatorV25.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, fulfillmentFlatFeeNativePPM, fulfillmentFlatFeeLinkDiscountPPM, nativePremiumPercentage, linkPremiumPercentage)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, fulfillmentFlatFeeNativePPM uint32, fulfillmentFlatFeeLinkDiscountPPM uint32, nativePremiumPercentage uint8, linkPremiumPercentage uint8) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.SetConfig(&_VRFCoordinatorV25.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, fulfillmentFlatFeeNativePPM, fulfillmentFlatFeeLinkDiscountPPM, nativePremiumPercentage, linkPremiumPercentage)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) SetLINKAndLINKNativeFeed(opts *bind.TransactOpts, link common.Address, linkNativeFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "setLINKAndLINKNativeFeed", link, linkNativeFeed)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SetLINKAndLINKNativeFeed(link common.Address, linkNativeFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.SetLINKAndLINKNativeFeed(&_VRFCoordinatorV25.TransactOpts, link, linkNativeFeed)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) SetLINKAndLINKNativeFeed(link common.Address, linkNativeFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.SetLINKAndLINKNativeFeed(&_VRFCoordinatorV25.TransactOpts, link, linkNativeFeed)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.TransferOwnership(&_VRFCoordinatorV25.TransactOpts, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.TransferOwnership(&_VRFCoordinatorV25.TransactOpts, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) Withdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "withdraw", recipient)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.Withdraw(&_VRFCoordinatorV25.TransactOpts, recipient)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.Withdraw(&_VRFCoordinatorV25.TransactOpts, recipient)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) WithdrawNative(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "withdrawNative", recipient)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) WithdrawNative(recipient common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.WithdrawNative(&_VRFCoordinatorV25.TransactOpts, recipient)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) WithdrawNative(recipient common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.WithdrawNative(&_VRFCoordinatorV25.TransactOpts, recipient)
}

type VRFCoordinatorV25ConfigSetIterator struct {
	Event *VRFCoordinatorV25ConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25ConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25ConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25ConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25ConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25ConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25ConfigSet struct {
	MinimumRequestConfirmations       uint16
	MaxGasLimit                       uint32
	StalenessSeconds                  uint32
	GasAfterPaymentCalculation        uint32
	FallbackWeiPerUnitLink            *big.Int
	FulfillmentFlatFeeNativePPM       uint32
	FulfillmentFlatFeeLinkDiscountPPM uint32
	NativePremiumPercentage           uint8
	LinkPremiumPercentage             uint8
	Raw                               types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorV25ConfigSetIterator, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25ConfigSetIterator{contract: _VRFCoordinatorV25.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25ConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25ConfigSet)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseConfigSet(log types.Log) (*VRFCoordinatorV25ConfigSet, error) {
	event := new(VRFCoordinatorV25ConfigSet)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25CoordinatorDeregisteredIterator struct {
	Event *VRFCoordinatorV25CoordinatorDeregistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25CoordinatorDeregisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25CoordinatorDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25CoordinatorDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25CoordinatorDeregisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25CoordinatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25CoordinatorDeregistered struct {
	CoordinatorAddress common.Address
	Raw                types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterCoordinatorDeregistered(opts *bind.FilterOpts) (*VRFCoordinatorV25CoordinatorDeregisteredIterator, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "CoordinatorDeregistered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25CoordinatorDeregisteredIterator{contract: _VRFCoordinatorV25.contract, event: "CoordinatorDeregistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchCoordinatorDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25CoordinatorDeregistered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "CoordinatorDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25CoordinatorDeregistered)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "CoordinatorDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseCoordinatorDeregistered(log types.Log) (*VRFCoordinatorV25CoordinatorDeregistered, error) {
	event := new(VRFCoordinatorV25CoordinatorDeregistered)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "CoordinatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25CoordinatorRegisteredIterator struct {
	Event *VRFCoordinatorV25CoordinatorRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25CoordinatorRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25CoordinatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25CoordinatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25CoordinatorRegisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25CoordinatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25CoordinatorRegistered struct {
	CoordinatorAddress common.Address
	Raw                types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterCoordinatorRegistered(opts *bind.FilterOpts) (*VRFCoordinatorV25CoordinatorRegisteredIterator, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "CoordinatorRegistered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25CoordinatorRegisteredIterator{contract: _VRFCoordinatorV25.contract, event: "CoordinatorRegistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchCoordinatorRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25CoordinatorRegistered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "CoordinatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25CoordinatorRegistered)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "CoordinatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseCoordinatorRegistered(log types.Log) (*VRFCoordinatorV25CoordinatorRegistered, error) {
	event := new(VRFCoordinatorV25CoordinatorRegistered)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "CoordinatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25FundsRecoveredIterator struct {
	Event *VRFCoordinatorV25FundsRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25FundsRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25FundsRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25FundsRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25FundsRecoveredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25FundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25FundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV25FundsRecoveredIterator, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25FundsRecoveredIterator{contract: _VRFCoordinatorV25.contract, event: "FundsRecovered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25FundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25FundsRecovered)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseFundsRecovered(log types.Log) (*VRFCoordinatorV25FundsRecovered, error) {
	event := new(VRFCoordinatorV25FundsRecovered)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25MigrationCompletedIterator struct {
	Event *VRFCoordinatorV25MigrationCompleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25MigrationCompletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25MigrationCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25MigrationCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25MigrationCompletedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25MigrationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25MigrationCompleted struct {
	NewCoordinator common.Address
	SubId          *big.Int
	Raw            types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterMigrationCompleted(opts *bind.FilterOpts) (*VRFCoordinatorV25MigrationCompletedIterator, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "MigrationCompleted")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25MigrationCompletedIterator{contract: _VRFCoordinatorV25.contract, event: "MigrationCompleted", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchMigrationCompleted(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25MigrationCompleted) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "MigrationCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25MigrationCompleted)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "MigrationCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseMigrationCompleted(log types.Log) (*VRFCoordinatorV25MigrationCompleted, error) {
	event := new(VRFCoordinatorV25MigrationCompleted)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "MigrationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25NativeFundsRecoveredIterator struct {
	Event *VRFCoordinatorV25NativeFundsRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25NativeFundsRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25NativeFundsRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25NativeFundsRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25NativeFundsRecoveredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25NativeFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25NativeFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterNativeFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV25NativeFundsRecoveredIterator, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25NativeFundsRecoveredIterator{contract: _VRFCoordinatorV25.contract, event: "NativeFundsRecovered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchNativeFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25NativeFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25NativeFundsRecovered)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseNativeFundsRecovered(log types.Log) (*VRFCoordinatorV25NativeFundsRecovered, error) {
	event := new(VRFCoordinatorV25NativeFundsRecovered)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25OwnershipTransferRequestedIterator struct {
	Event *VRFCoordinatorV25OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25OwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25OwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV25OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25OwnershipTransferRequestedIterator{contract: _VRFCoordinatorV25.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25OwnershipTransferRequested)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV25OwnershipTransferRequested, error) {
	event := new(VRFCoordinatorV25OwnershipTransferRequested)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25OwnershipTransferredIterator struct {
	Event *VRFCoordinatorV25OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV25OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25OwnershipTransferredIterator{contract: _VRFCoordinatorV25.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25OwnershipTransferred)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV25OwnershipTransferred, error) {
	event := new(VRFCoordinatorV25OwnershipTransferred)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25ProvingKeyDeregisteredIterator struct {
	Event *VRFCoordinatorV25ProvingKeyDeregistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25ProvingKeyDeregisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25ProvingKeyDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25ProvingKeyDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25ProvingKeyDeregisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25ProvingKeyDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25ProvingKeyDeregistered struct {
	KeyHash [32]byte
	MaxGas  uint64
	Raw     types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterProvingKeyDeregistered(opts *bind.FilterOpts) (*VRFCoordinatorV25ProvingKeyDeregisteredIterator, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "ProvingKeyDeregistered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25ProvingKeyDeregisteredIterator{contract: _VRFCoordinatorV25.contract, event: "ProvingKeyDeregistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25ProvingKeyDeregistered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "ProvingKeyDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25ProvingKeyDeregistered)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseProvingKeyDeregistered(log types.Log) (*VRFCoordinatorV25ProvingKeyDeregistered, error) {
	event := new(VRFCoordinatorV25ProvingKeyDeregistered)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25ProvingKeyRegisteredIterator struct {
	Event *VRFCoordinatorV25ProvingKeyRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25ProvingKeyRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25ProvingKeyRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25ProvingKeyRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25ProvingKeyRegisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25ProvingKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25ProvingKeyRegistered struct {
	KeyHash [32]byte
	MaxGas  uint64
	Raw     types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterProvingKeyRegistered(opts *bind.FilterOpts) (*VRFCoordinatorV25ProvingKeyRegisteredIterator, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "ProvingKeyRegistered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25ProvingKeyRegisteredIterator{contract: _VRFCoordinatorV25.contract, event: "ProvingKeyRegistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25ProvingKeyRegistered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "ProvingKeyRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25ProvingKeyRegistered)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseProvingKeyRegistered(log types.Log) (*VRFCoordinatorV25ProvingKeyRegistered, error) {
	event := new(VRFCoordinatorV25ProvingKeyRegistered)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25RandomWordsFulfilledIterator struct {
	Event *VRFCoordinatorV25RandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25RandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25RandomWordsFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25RandomWordsFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25RandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25RandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25RandomWordsFulfilled struct {
	RequestId     *big.Int
	OutputSeed    *big.Int
	SubId         *big.Int
	Payment       *big.Int
	NativePayment bool
	Success       bool
	OnlyPremium   bool
	Raw           types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int, subId []*big.Int) (*VRFCoordinatorV25RandomWordsFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "RandomWordsFulfilled", requestIdRule, subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25RandomWordsFulfilledIterator{contract: _VRFCoordinatorV25.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25RandomWordsFulfilled, requestId []*big.Int, subId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "RandomWordsFulfilled", requestIdRule, subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25RandomWordsFulfilled)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV25RandomWordsFulfilled, error) {
	event := new(VRFCoordinatorV25RandomWordsFulfilled)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25RandomWordsRequestedIterator struct {
	Event *VRFCoordinatorV25RandomWordsRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25RandomWordsRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25RandomWordsRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25RandomWordsRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25RandomWordsRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25RandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25RandomWordsRequested struct {
	KeyHash                     [32]byte
	RequestId                   *big.Int
	PreSeed                     *big.Int
	SubId                       *big.Int
	MinimumRequestConfirmations uint16
	CallbackGasLimit            uint32
	NumWords                    uint32
	ExtraArgs                   []byte
	Sender                      common.Address
	Raw                         types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (*VRFCoordinatorV25RandomWordsRequestedIterator, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25RandomWordsRequestedIterator{contract: _VRFCoordinatorV25.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25RandomWordsRequested, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25RandomWordsRequested)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV25RandomWordsRequested, error) {
	event := new(VRFCoordinatorV25RandomWordsRequested)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionCanceledIterator struct {
	Event *VRFCoordinatorV25SubscriptionCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25SubscriptionCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25SubscriptionCanceledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionCanceled struct {
	SubId        *big.Int
	To           common.Address
	AmountLink   *big.Int
	AmountNative *big.Int
	Raw          types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionCanceledIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionCanceledIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionCanceled, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionCanceled)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorV25SubscriptionCanceled, error) {
	event := new(VRFCoordinatorV25SubscriptionCanceled)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionConsumerAddedIterator struct {
	Event *VRFCoordinatorV25SubscriptionConsumerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionConsumerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionConsumerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25SubscriptionConsumerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25SubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionConsumerAdded struct {
	SubId    *big.Int
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionConsumerAddedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionConsumerAddedIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionConsumerAdded, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionConsumerAdded)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorV25SubscriptionConsumerAdded, error) {
	event := new(VRFCoordinatorV25SubscriptionConsumerAdded)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionConsumerRemovedIterator struct {
	Event *VRFCoordinatorV25SubscriptionConsumerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionConsumerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionConsumerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25SubscriptionConsumerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25SubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionConsumerRemoved struct {
	SubId    *big.Int
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionConsumerRemovedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionConsumerRemovedIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionConsumerRemoved, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionConsumerRemoved)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorV25SubscriptionConsumerRemoved, error) {
	event := new(VRFCoordinatorV25SubscriptionConsumerRemoved)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionCreatedIterator struct {
	Event *VRFCoordinatorV25SubscriptionCreated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionCreatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25SubscriptionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25SubscriptionCreatedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionCreated struct {
	SubId *big.Int
	Owner common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionCreatedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionCreatedIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionCreated, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionCreated)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorV25SubscriptionCreated, error) {
	event := new(VRFCoordinatorV25SubscriptionCreated)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionFundedIterator struct {
	Event *VRFCoordinatorV25SubscriptionFunded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionFundedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionFunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25SubscriptionFunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25SubscriptionFundedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionFunded struct {
	SubId      *big.Int
	OldBalance *big.Int
	NewBalance *big.Int
	Raw        types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionFunded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionFundedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionFundedIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionFunded, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionFunded)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorV25SubscriptionFunded, error) {
	event := new(VRFCoordinatorV25SubscriptionFunded)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionFundedWithNativeIterator struct {
	Event *VRFCoordinatorV25SubscriptionFundedWithNative

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionFundedWithNativeIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionFundedWithNative)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25SubscriptionFundedWithNative)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25SubscriptionFundedWithNativeIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionFundedWithNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionFundedWithNative struct {
	SubId            *big.Int
	OldNativeBalance *big.Int
	NewNativeBalance *big.Int
	Raw              types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionFundedWithNative(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionFundedWithNativeIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionFundedWithNative", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionFundedWithNativeIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionFundedWithNative", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionFundedWithNative(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionFundedWithNative, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionFundedWithNative", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionFundedWithNative)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionFundedWithNative", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionFundedWithNative(log types.Log) (*VRFCoordinatorV25SubscriptionFundedWithNative, error) {
	event := new(VRFCoordinatorV25SubscriptionFundedWithNative)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionFundedWithNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator struct {
	Event *VRFCoordinatorV25SubscriptionOwnerTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionOwnerTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25SubscriptionOwnerTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionOwnerTransferRequested struct {
	SubId *big.Int
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionOwnerTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionOwnerTransferRequested, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionOwnerTransferRequested)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorV25SubscriptionOwnerTransferRequested, error) {
	event := new(VRFCoordinatorV25SubscriptionOwnerTransferRequested)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionOwnerTransferredIterator struct {
	Event *VRFCoordinatorV25SubscriptionOwnerTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionOwnerTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionOwnerTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV25SubscriptionOwnerTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV25SubscriptionOwnerTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionOwnerTransferred struct {
	SubId *big.Int
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionOwnerTransferredIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionOwnerTransferredIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionOwnerTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionOwnerTransferred, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionOwnerTransferred)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorV25SubscriptionOwnerTransferred, error) {
	event := new(VRFCoordinatorV25SubscriptionOwnerTransferred)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetSubscription struct {
	Balance       *big.Int
	NativeBalance *big.Int
	ReqCount      uint64
	Owner         common.Address
	Consumers     []common.Address
}
type SConfig struct {
	MinimumRequestConfirmations       uint16
	MaxGasLimit                       uint32
	ReentrancyLock                    bool
	StalenessSeconds                  uint32
	GasAfterPaymentCalculation        uint32
	FulfillmentFlatFeeNativePPM       uint32
	FulfillmentFlatFeeLinkDiscountPPM uint32
	NativePremiumPercentage           uint8
	LinkPremiumPercentage             uint8
}
type SProvingKeys struct {
	Exists bool
	MaxGas uint64
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFCoordinatorV25.abi.Events["ConfigSet"].ID:
		return _VRFCoordinatorV25.ParseConfigSet(log)
	case _VRFCoordinatorV25.abi.Events["CoordinatorDeregistered"].ID:
		return _VRFCoordinatorV25.ParseCoordinatorDeregistered(log)
	case _VRFCoordinatorV25.abi.Events["CoordinatorRegistered"].ID:
		return _VRFCoordinatorV25.ParseCoordinatorRegistered(log)
	case _VRFCoordinatorV25.abi.Events["FundsRecovered"].ID:
		return _VRFCoordinatorV25.ParseFundsRecovered(log)
	case _VRFCoordinatorV25.abi.Events["MigrationCompleted"].ID:
		return _VRFCoordinatorV25.ParseMigrationCompleted(log)
	case _VRFCoordinatorV25.abi.Events["NativeFundsRecovered"].ID:
		return _VRFCoordinatorV25.ParseNativeFundsRecovered(log)
	case _VRFCoordinatorV25.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFCoordinatorV25.ParseOwnershipTransferRequested(log)
	case _VRFCoordinatorV25.abi.Events["OwnershipTransferred"].ID:
		return _VRFCoordinatorV25.ParseOwnershipTransferred(log)
	case _VRFCoordinatorV25.abi.Events["ProvingKeyDeregistered"].ID:
		return _VRFCoordinatorV25.ParseProvingKeyDeregistered(log)
	case _VRFCoordinatorV25.abi.Events["ProvingKeyRegistered"].ID:
		return _VRFCoordinatorV25.ParseProvingKeyRegistered(log)
	case _VRFCoordinatorV25.abi.Events["RandomWordsFulfilled"].ID:
		return _VRFCoordinatorV25.ParseRandomWordsFulfilled(log)
	case _VRFCoordinatorV25.abi.Events["RandomWordsRequested"].ID:
		return _VRFCoordinatorV25.ParseRandomWordsRequested(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionCanceled"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionCanceled(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionConsumerAdded"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionConsumerAdded(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionConsumerRemoved"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionConsumerRemoved(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionCreated"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionCreated(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionFunded"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionFunded(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionFundedWithNative"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionFundedWithNative(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionOwnerTransferRequested"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionOwnerTransferRequested(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionOwnerTransferred"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionOwnerTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFCoordinatorV25ConfigSet) Topic() common.Hash {
	return common.HexToHash("0x2c6b6b12413678366b05b145c5f00745bdd00e739131ab5de82484a50c9d78b6")
}

func (VRFCoordinatorV25CoordinatorDeregistered) Topic() common.Hash {
	return common.HexToHash("0xf80a1a97fd42251f3c33cda98635e7399253033a6774fe37cd3f650b5282af37")
}

func (VRFCoordinatorV25CoordinatorRegistered) Topic() common.Hash {
	return common.HexToHash("0xb7cabbfc11e66731fc77de0444614282023bcbd41d16781c753a431d0af01625")
}

func (VRFCoordinatorV25FundsRecovered) Topic() common.Hash {
	return common.HexToHash("0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600")
}

func (VRFCoordinatorV25MigrationCompleted) Topic() common.Hash {
	return common.HexToHash("0xd63ca8cb945956747ee69bfdc3ea754c24a4caf7418db70e46052f7850be4187")
}

func (VRFCoordinatorV25NativeFundsRecovered) Topic() common.Hash {
	return common.HexToHash("0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c")
}

func (VRFCoordinatorV25OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFCoordinatorV25OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VRFCoordinatorV25ProvingKeyDeregistered) Topic() common.Hash {
	return common.HexToHash("0x9b6868e0eb737bcd72205360baa6bfd0ba4e4819a33ade2db384e8a8025639a5")
}

func (VRFCoordinatorV25ProvingKeyRegistered) Topic() common.Hash {
	return common.HexToHash("0x9b911b2c240bfbef3b6a8f7ed6ee321d1258bb2a3fe6becab52ac1cd3210afd3")
}

func (VRFCoordinatorV25RandomWordsFulfilled) Topic() common.Hash {
	return common.HexToHash("0xaeb4b4786571e184246d39587f659abf0e26f41f6a3358692250382c0cdb47b7")
}

func (VRFCoordinatorV25RandomWordsRequested) Topic() common.Hash {
	return common.HexToHash("0xeb0e3652e0f44f417695e6e90f2f42c99b65cd7169074c5a654b16b9748c3a4e")
}

func (VRFCoordinatorV25SubscriptionCanceled) Topic() common.Hash {
	return common.HexToHash("0x8c74ce8b8cf87f5eb001275c8be27eb34ea2b62bfab6814fcc62192bb63e81c4")
}

func (VRFCoordinatorV25SubscriptionConsumerAdded) Topic() common.Hash {
	return common.HexToHash("0x1e980d04aa7648e205713e5e8ea3808672ac163d10936d36f91b2c88ac1575e1")
}

func (VRFCoordinatorV25SubscriptionConsumerRemoved) Topic() common.Hash {
	return common.HexToHash("0x32158c6058347c1601b2d12bc696ac6901d8a9a9aa3ba10c27ab0a983e8425a7")
}

func (VRFCoordinatorV25SubscriptionCreated) Topic() common.Hash {
	return common.HexToHash("0x1d3015d7ba850fa198dc7b1a3f5d42779313a681035f77c8c03764c61005518d")
}

func (VRFCoordinatorV25SubscriptionFunded) Topic() common.Hash {
	return common.HexToHash("0x1ced9348ff549fceab2ac57cd3a9de38edaaab274b725ee82c23e8fc8c4eec7a")
}

func (VRFCoordinatorV25SubscriptionFundedWithNative) Topic() common.Hash {
	return common.HexToHash("0x7603b205d03651ee812f803fccde89f1012e545a9c99f0abfea9cedd0fd8e902")
}

func (VRFCoordinatorV25SubscriptionOwnerTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x21a4dad170a6bf476c31bbcf4a16628295b0e450672eec25d7c93308e05344a1")
}

func (VRFCoordinatorV25SubscriptionOwnerTransferred) Topic() common.Hash {
	return common.HexToHash("0xd4114ab6e9af9f597c52041f32d62dc57c5c4e4c0d4427006069635e216c9386")
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25) Address() common.Address {
	return _VRFCoordinatorV25.address
}

type VRFCoordinatorV25Interface interface {
	BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error)

	LINK(opts *bind.CallOpts) (common.Address, error)

	LINKNATIVEFEED(opts *bind.CallOpts) (common.Address, error)

	MAXCONSUMERS(opts *bind.CallOpts) (uint16, error)

	MAXNUMWORDS(opts *bind.CallOpts) (uint32, error)

	MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error)

	GetActiveSubscriptionIds(opts *bind.CallOpts, startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error)

	GetSubscription(opts *bind.CallOpts, subId *big.Int) (GetSubscription,

		error)

	HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PendingRequestExists(opts *bind.CallOpts, subId *big.Int) (bool, error)

	SConfig(opts *bind.CallOpts) (SConfig,

		error)

	SCurrentSubNonce(opts *bind.CallOpts) (uint64, error)

	SFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error)

	SProvingKeyHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error)

	SProvingKeys(opts *bind.CallOpts, arg0 [32]byte) (SProvingKeys,

		error)

	SRequestCommitments(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error)

	STotalBalance(opts *bind.CallOpts) (*big.Int, error)

	STotalNativeBalance(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error)

	AddConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error)

	CancelSubscription(opts *bind.TransactOpts, subId *big.Int, to common.Address) (*types.Transaction, error)

	CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error)

	DeregisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error)

	DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error)

	FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV25RequestCommitment, onlyPremium bool) (*types.Transaction, error)

	FundSubscriptionWithNative(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error)

	Migrate(opts *bind.TransactOpts, subId *big.Int, newCoordinator common.Address) (*types.Transaction, error)

	OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error)

	OwnerCancelSubscription(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error)

	RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	RecoverNativeFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	RegisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error)

	RegisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int, maxGas uint64) (*types.Transaction, error)

	RemoveConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error)

	RequestRandomWords(opts *bind.TransactOpts, req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error)

	RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int, newOwner common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, fulfillmentFlatFeeNativePPM uint32, fulfillmentFlatFeeLinkDiscountPPM uint32, nativePremiumPercentage uint8, linkPremiumPercentage uint8) (*types.Transaction, error)

	SetLINKAndLINKNativeFeed(opts *bind.TransactOpts, link common.Address, linkNativeFeed common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	WithdrawNative(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorV25ConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25ConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*VRFCoordinatorV25ConfigSet, error)

	FilterCoordinatorDeregistered(opts *bind.FilterOpts) (*VRFCoordinatorV25CoordinatorDeregisteredIterator, error)

	WatchCoordinatorDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25CoordinatorDeregistered) (event.Subscription, error)

	ParseCoordinatorDeregistered(log types.Log) (*VRFCoordinatorV25CoordinatorDeregistered, error)

	FilterCoordinatorRegistered(opts *bind.FilterOpts) (*VRFCoordinatorV25CoordinatorRegisteredIterator, error)

	WatchCoordinatorRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25CoordinatorRegistered) (event.Subscription, error)

	ParseCoordinatorRegistered(log types.Log) (*VRFCoordinatorV25CoordinatorRegistered, error)

	FilterFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV25FundsRecoveredIterator, error)

	WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25FundsRecovered) (event.Subscription, error)

	ParseFundsRecovered(log types.Log) (*VRFCoordinatorV25FundsRecovered, error)

	FilterMigrationCompleted(opts *bind.FilterOpts) (*VRFCoordinatorV25MigrationCompletedIterator, error)

	WatchMigrationCompleted(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25MigrationCompleted) (event.Subscription, error)

	ParseMigrationCompleted(log types.Log) (*VRFCoordinatorV25MigrationCompleted, error)

	FilterNativeFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV25NativeFundsRecoveredIterator, error)

	WatchNativeFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25NativeFundsRecovered) (event.Subscription, error)

	ParseNativeFundsRecovered(log types.Log) (*VRFCoordinatorV25NativeFundsRecovered, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV25OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV25OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV25OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV25OwnershipTransferred, error)

	FilterProvingKeyDeregistered(opts *bind.FilterOpts) (*VRFCoordinatorV25ProvingKeyDeregisteredIterator, error)

	WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25ProvingKeyDeregistered) (event.Subscription, error)

	ParseProvingKeyDeregistered(log types.Log) (*VRFCoordinatorV25ProvingKeyDeregistered, error)

	FilterProvingKeyRegistered(opts *bind.FilterOpts) (*VRFCoordinatorV25ProvingKeyRegisteredIterator, error)

	WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25ProvingKeyRegistered) (event.Subscription, error)

	ParseProvingKeyRegistered(log types.Log) (*VRFCoordinatorV25ProvingKeyRegistered, error)

	FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int, subId []*big.Int) (*VRFCoordinatorV25RandomWordsFulfilledIterator, error)

	WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25RandomWordsFulfilled, requestId []*big.Int, subId []*big.Int) (event.Subscription, error)

	ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV25RandomWordsFulfilled, error)

	FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (*VRFCoordinatorV25RandomWordsRequestedIterator, error)

	WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25RandomWordsRequested, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (event.Subscription, error)

	ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV25RandomWordsRequested, error)

	FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionCanceledIterator, error)

	WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionCanceled, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorV25SubscriptionCanceled, error)

	FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionConsumerAddedIterator, error)

	WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionConsumerAdded, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorV25SubscriptionConsumerAdded, error)

	FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionConsumerRemovedIterator, error)

	WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionConsumerRemoved, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorV25SubscriptionConsumerRemoved, error)

	FilterSubscriptionCreated(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionCreatedIterator, error)

	WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionCreated, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorV25SubscriptionCreated, error)

	FilterSubscriptionFunded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionFundedIterator, error)

	WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionFunded, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorV25SubscriptionFunded, error)

	FilterSubscriptionFundedWithNative(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionFundedWithNativeIterator, error)

	WatchSubscriptionFundedWithNative(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionFundedWithNative, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionFundedWithNative(log types.Log) (*VRFCoordinatorV25SubscriptionFundedWithNative, error)

	FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator, error)

	WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionOwnerTransferRequested, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorV25SubscriptionOwnerTransferRequested, error)

	FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV25SubscriptionOwnerTransferredIterator, error)

	WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionOwnerTransferred, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorV25SubscriptionOwnerTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
