// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package build

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IZkEvmV2BlockData is an auto generated low-level Go binding around an user-defined struct.
type IZkEvmV2BlockData struct {
	BlockRootHash         [32]byte
	L2BlockTimestamp      uint32
	Transactions          [][]byte
	L2ToL1MsgHashes       [][32]byte
	FromAddresses         []byte
	BatchReceptionIndices []uint16
}

// ZkEvmV2MetaData contains all meta data concerning the ZkEvmV2 contract.
var ZkEvmV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BlockTimestampError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyBlock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"FeePaymentFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"IsNotPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"IsPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"L1L2MessageNotSent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inde\",\"type\":\"uint256\"}],\"name\":\"MemoryOutOfBounds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageAlreadyReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageAlreadySent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageDoesNotExistOrHasAlreadyBeenClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"MessageSendingFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofIsEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StartingRootHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownTransactionType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueSentTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueShouldBeGreaterThanFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongBytesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"resettingAddress\",\"type\":\"address\"}],\"name\":\"AmountUsedInPeriodReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"stateRootHash\",\"type\":\"bytes32\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lastBlockFinalized\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"startingRootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"finalRootHash\",\"type\":\"bytes32\"}],\"name\":\"BlocksVerificationDone\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"messageHashes\",\"type\":\"bytes32[]\"}],\"name\":\"L1L2MessagesReceivedOnL2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"L2L1MessageHashAddedToInbox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"amountChangeBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"amountUsedLoweredToLimit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"usedAmountResetToZero\",\"type\":\"bool\"}],\"name\":\"LimitAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"UnPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proofType\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifierSetBy\",\"type\":\"address\"}],\"name\":\"VerifierAddressChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENERAL_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INBOX_STATUS_RECEIVED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INBOX_STATUS_UNKNOWN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_L2_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_L1_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUTBOX_STATUS_RECEIVED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUTBOX_STATUS_SENT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUTBOX_STATUS_UNKNOWN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVING_SYSTEM_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RATE_LIMIT_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentL2BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodAmountInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"l2BlockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"transactions\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"l2ToL1MsgHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"fromAddresses\",\"type\":\"bytes\"},{\"internalType\":\"uint16[]\",\"name\":\"batchReceptionIndices\",\"type\":\"uint16[]\"}],\"internalType\":\"structIZkEvmV2.BlockData[]\",\"name\":\"_blocksData\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_proofType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_parentStateRootHash\",\"type\":\"bytes32\"}],\"name\":\"finalizeBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"l2BlockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"transactions\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"l2ToL1MsgHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"fromAddresses\",\"type\":\"bytes\"},{\"internalType\":\"uint16[]\",\"name\":\"batchReceptionIndices\",\"type\":\"uint16[]\"}],\"internalType\":\"structIZkEvmV2.BlockData[]\",\"name\":\"_blocksData\",\"type\":\"tuple[]\"}],\"name\":\"finalizeBlocksWithoutProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"inboxL2L1MessageStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_initialStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_initialL2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_defaultVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_securityCouncil\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_rateLimitPeriodInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rateLimitAmountInWei\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextMessageNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"outboxL1L2MessageStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pauseType\",\"type\":\"bytes32\"}],\"name\":\"pauseByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pauseTypeStatuses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetAmountUsedInPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"resetRateLimitAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifierAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_proofType\",\"type\":\"uint256\"}],\"name\":\"setVerifierAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stateRootHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pauseType\",\"type\":\"bytes32\"}],\"name\":\"unPauseByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ZkEvmV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkEvmV2MetaData.ABI instead.
var ZkEvmV2ABI = ZkEvmV2MetaData.ABI

// ZkEvmV2 is an auto generated Go binding around an Ethereum contract.
type ZkEvmV2 struct {
	ZkEvmV2Caller     // Read-only binding to the contract
	ZkEvmV2Transactor // Write-only binding to the contract
	ZkEvmV2Filterer   // Log filterer for contract events
}

// ZkEvmV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ZkEvmV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkEvmV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkEvmV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkEvmV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkEvmV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkEvmV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkEvmV2Session struct {
	Contract     *ZkEvmV2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkEvmV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkEvmV2CallerSession struct {
	Contract *ZkEvmV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ZkEvmV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkEvmV2TransactorSession struct {
	Contract     *ZkEvmV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZkEvmV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ZkEvmV2Raw struct {
	Contract *ZkEvmV2 // Generic contract binding to access the raw methods on
}

// ZkEvmV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkEvmV2CallerRaw struct {
	Contract *ZkEvmV2Caller // Generic read-only contract binding to access the raw methods on
}

// ZkEvmV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkEvmV2TransactorRaw struct {
	Contract *ZkEvmV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewZkEvmV2 creates a new instance of ZkEvmV2, bound to a specific deployed contract.
func NewZkEvmV2(address common.Address, backend bind.ContractBackend) (*ZkEvmV2, error) {
	contract, err := bindZkEvmV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2{ZkEvmV2Caller: ZkEvmV2Caller{contract: contract}, ZkEvmV2Transactor: ZkEvmV2Transactor{contract: contract}, ZkEvmV2Filterer: ZkEvmV2Filterer{contract: contract}}, nil
}

// NewZkEvmV2Caller creates a new read-only instance of ZkEvmV2, bound to a specific deployed contract.
func NewZkEvmV2Caller(address common.Address, caller bind.ContractCaller) (*ZkEvmV2Caller, error) {
	contract, err := bindZkEvmV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2Caller{contract: contract}, nil
}

// NewZkEvmV2Transactor creates a new write-only instance of ZkEvmV2, bound to a specific deployed contract.
func NewZkEvmV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ZkEvmV2Transactor, error) {
	contract, err := bindZkEvmV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2Transactor{contract: contract}, nil
}

// NewZkEvmV2Filterer creates a new log filterer instance of ZkEvmV2, bound to a specific deployed contract.
func NewZkEvmV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ZkEvmV2Filterer, error) {
	contract, err := bindZkEvmV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2Filterer{contract: contract}, nil
}

// bindZkEvmV2 binds a generic wrapper to an already deployed contract.
func bindZkEvmV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZkEvmV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkEvmV2 *ZkEvmV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkEvmV2.Contract.ZkEvmV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkEvmV2 *ZkEvmV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.ZkEvmV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkEvmV2 *ZkEvmV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.ZkEvmV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkEvmV2 *ZkEvmV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkEvmV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkEvmV2 *ZkEvmV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkEvmV2 *ZkEvmV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZkEvmV2.Contract.DEFAULTADMINROLE(&_ZkEvmV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZkEvmV2.Contract.DEFAULTADMINROLE(&_ZkEvmV2.CallOpts)
}

// GENERALPAUSETYPE is a free data retrieval call binding the contract method 0x6a637967.
//
// Solidity: function GENERAL_PAUSE_TYPE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Caller) GENERALPAUSETYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "GENERAL_PAUSE_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GENERALPAUSETYPE is a free data retrieval call binding the contract method 0x6a637967.
//
// Solidity: function GENERAL_PAUSE_TYPE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Session) GENERALPAUSETYPE() ([32]byte, error) {
	return _ZkEvmV2.Contract.GENERALPAUSETYPE(&_ZkEvmV2.CallOpts)
}

// GENERALPAUSETYPE is a free data retrieval call binding the contract method 0x6a637967.
//
// Solidity: function GENERAL_PAUSE_TYPE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2CallerSession) GENERALPAUSETYPE() ([32]byte, error) {
	return _ZkEvmV2.Contract.GENERALPAUSETYPE(&_ZkEvmV2.CallOpts)
}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2Caller) INBOXSTATUSRECEIVED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "INBOX_STATUS_RECEIVED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2Session) INBOXSTATUSRECEIVED() (uint8, error) {
	return _ZkEvmV2.Contract.INBOXSTATUSRECEIVED(&_ZkEvmV2.CallOpts)
}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2CallerSession) INBOXSTATUSRECEIVED() (uint8, error) {
	return _ZkEvmV2.Contract.INBOXSTATUSRECEIVED(&_ZkEvmV2.CallOpts)
}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2Caller) INBOXSTATUSUNKNOWN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "INBOX_STATUS_UNKNOWN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2Session) INBOXSTATUSUNKNOWN() (uint8, error) {
	return _ZkEvmV2.Contract.INBOXSTATUSUNKNOWN(&_ZkEvmV2.CallOpts)
}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2CallerSession) INBOXSTATUSUNKNOWN() (uint8, error) {
	return _ZkEvmV2.Contract.INBOXSTATUSUNKNOWN(&_ZkEvmV2.CallOpts)
}

// L1L2PAUSETYPE is a free data retrieval call binding the contract method 0x11314d0f.
//
// Solidity: function L1_L2_PAUSE_TYPE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Caller) L1L2PAUSETYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "L1_L2_PAUSE_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1L2PAUSETYPE is a free data retrieval call binding the contract method 0x11314d0f.
//
// Solidity: function L1_L2_PAUSE_TYPE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Session) L1L2PAUSETYPE() ([32]byte, error) {
	return _ZkEvmV2.Contract.L1L2PAUSETYPE(&_ZkEvmV2.CallOpts)
}

// L1L2PAUSETYPE is a free data retrieval call binding the contract method 0x11314d0f.
//
// Solidity: function L1_L2_PAUSE_TYPE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2CallerSession) L1L2PAUSETYPE() ([32]byte, error) {
	return _ZkEvmV2.Contract.L1L2PAUSETYPE(&_ZkEvmV2.CallOpts)
}

// L2L1PAUSETYPE is a free data retrieval call binding the contract method 0xabd6230d.
//
// Solidity: function L2_L1_PAUSE_TYPE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Caller) L2L1PAUSETYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "L2_L1_PAUSE_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2L1PAUSETYPE is a free data retrieval call binding the contract method 0xabd6230d.
//
// Solidity: function L2_L1_PAUSE_TYPE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Session) L2L1PAUSETYPE() ([32]byte, error) {
	return _ZkEvmV2.Contract.L2L1PAUSETYPE(&_ZkEvmV2.CallOpts)
}

// L2L1PAUSETYPE is a free data retrieval call binding the contract method 0xabd6230d.
//
// Solidity: function L2_L1_PAUSE_TYPE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2CallerSession) L2L1PAUSETYPE() ([32]byte, error) {
	return _ZkEvmV2.Contract.L2L1PAUSETYPE(&_ZkEvmV2.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Caller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Session) OPERATORROLE() ([32]byte, error) {
	return _ZkEvmV2.Contract.OPERATORROLE(&_ZkEvmV2.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2CallerSession) OPERATORROLE() ([32]byte, error) {
	return _ZkEvmV2.Contract.OPERATORROLE(&_ZkEvmV2.CallOpts)
}

// OUTBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x73bd07b7.
//
// Solidity: function OUTBOX_STATUS_RECEIVED() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2Caller) OUTBOXSTATUSRECEIVED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "OUTBOX_STATUS_RECEIVED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OUTBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x73bd07b7.
//
// Solidity: function OUTBOX_STATUS_RECEIVED() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2Session) OUTBOXSTATUSRECEIVED() (uint8, error) {
	return _ZkEvmV2.Contract.OUTBOXSTATUSRECEIVED(&_ZkEvmV2.CallOpts)
}

// OUTBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x73bd07b7.
//
// Solidity: function OUTBOX_STATUS_RECEIVED() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2CallerSession) OUTBOXSTATUSRECEIVED() (uint8, error) {
	return _ZkEvmV2.Contract.OUTBOXSTATUSRECEIVED(&_ZkEvmV2.CallOpts)
}

// OUTBOXSTATUSSENT is a free data retrieval call binding the contract method 0x5b7eb4bd.
//
// Solidity: function OUTBOX_STATUS_SENT() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2Caller) OUTBOXSTATUSSENT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "OUTBOX_STATUS_SENT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OUTBOXSTATUSSENT is a free data retrieval call binding the contract method 0x5b7eb4bd.
//
// Solidity: function OUTBOX_STATUS_SENT() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2Session) OUTBOXSTATUSSENT() (uint8, error) {
	return _ZkEvmV2.Contract.OUTBOXSTATUSSENT(&_ZkEvmV2.CallOpts)
}

// OUTBOXSTATUSSENT is a free data retrieval call binding the contract method 0x5b7eb4bd.
//
// Solidity: function OUTBOX_STATUS_SENT() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2CallerSession) OUTBOXSTATUSSENT() (uint8, error) {
	return _ZkEvmV2.Contract.OUTBOXSTATUSSENT(&_ZkEvmV2.CallOpts)
}

// OUTBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x986fcddd.
//
// Solidity: function OUTBOX_STATUS_UNKNOWN() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2Caller) OUTBOXSTATUSUNKNOWN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "OUTBOX_STATUS_UNKNOWN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OUTBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x986fcddd.
//
// Solidity: function OUTBOX_STATUS_UNKNOWN() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2Session) OUTBOXSTATUSUNKNOWN() (uint8, error) {
	return _ZkEvmV2.Contract.OUTBOXSTATUSUNKNOWN(&_ZkEvmV2.CallOpts)
}

// OUTBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x986fcddd.
//
// Solidity: function OUTBOX_STATUS_UNKNOWN() view returns(uint8)
func (_ZkEvmV2 *ZkEvmV2CallerSession) OUTBOXSTATUSUNKNOWN() (uint8, error) {
	return _ZkEvmV2.Contract.OUTBOXSTATUSUNKNOWN(&_ZkEvmV2.CallOpts)
}

// PAUSEMANAGERROLE is a free data retrieval call binding the contract method 0xd84f91e8.
//
// Solidity: function PAUSE_MANAGER_ROLE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Caller) PAUSEMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "PAUSE_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEMANAGERROLE is a free data retrieval call binding the contract method 0xd84f91e8.
//
// Solidity: function PAUSE_MANAGER_ROLE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Session) PAUSEMANAGERROLE() ([32]byte, error) {
	return _ZkEvmV2.Contract.PAUSEMANAGERROLE(&_ZkEvmV2.CallOpts)
}

// PAUSEMANAGERROLE is a free data retrieval call binding the contract method 0xd84f91e8.
//
// Solidity: function PAUSE_MANAGER_ROLE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2CallerSession) PAUSEMANAGERROLE() ([32]byte, error) {
	return _ZkEvmV2.Contract.PAUSEMANAGERROLE(&_ZkEvmV2.CallOpts)
}

// PROVINGSYSTEMPAUSETYPE is a free data retrieval call binding the contract method 0xb4a5a4b7.
//
// Solidity: function PROVING_SYSTEM_PAUSE_TYPE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Caller) PROVINGSYSTEMPAUSETYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "PROVING_SYSTEM_PAUSE_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROVINGSYSTEMPAUSETYPE is a free data retrieval call binding the contract method 0xb4a5a4b7.
//
// Solidity: function PROVING_SYSTEM_PAUSE_TYPE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Session) PROVINGSYSTEMPAUSETYPE() ([32]byte, error) {
	return _ZkEvmV2.Contract.PROVINGSYSTEMPAUSETYPE(&_ZkEvmV2.CallOpts)
}

// PROVINGSYSTEMPAUSETYPE is a free data retrieval call binding the contract method 0xb4a5a4b7.
//
// Solidity: function PROVING_SYSTEM_PAUSE_TYPE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2CallerSession) PROVINGSYSTEMPAUSETYPE() ([32]byte, error) {
	return _ZkEvmV2.Contract.PROVINGSYSTEMPAUSETYPE(&_ZkEvmV2.CallOpts)
}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Caller) RATELIMITSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "RATE_LIMIT_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Session) RATELIMITSETTERROLE() ([32]byte, error) {
	return _ZkEvmV2.Contract.RATELIMITSETTERROLE(&_ZkEvmV2.CallOpts)
}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2CallerSession) RATELIMITSETTERROLE() ([32]byte, error) {
	return _ZkEvmV2.Contract.RATELIMITSETTERROLE(&_ZkEvmV2.CallOpts)
}

// CurrentL2BlockNumber is a free data retrieval call binding the contract method 0x695378f5.
//
// Solidity: function currentL2BlockNumber() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Caller) CurrentL2BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "currentL2BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentL2BlockNumber is a free data retrieval call binding the contract method 0x695378f5.
//
// Solidity: function currentL2BlockNumber() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Session) CurrentL2BlockNumber() (*big.Int, error) {
	return _ZkEvmV2.Contract.CurrentL2BlockNumber(&_ZkEvmV2.CallOpts)
}

// CurrentL2BlockNumber is a free data retrieval call binding the contract method 0x695378f5.
//
// Solidity: function currentL2BlockNumber() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2CallerSession) CurrentL2BlockNumber() (*big.Int, error) {
	return _ZkEvmV2.Contract.CurrentL2BlockNumber(&_ZkEvmV2.CallOpts)
}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Caller) CurrentPeriodAmountInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "currentPeriodAmountInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Session) CurrentPeriodAmountInWei() (*big.Int, error) {
	return _ZkEvmV2.Contract.CurrentPeriodAmountInWei(&_ZkEvmV2.CallOpts)
}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2CallerSession) CurrentPeriodAmountInWei() (*big.Int, error) {
	return _ZkEvmV2.Contract.CurrentPeriodAmountInWei(&_ZkEvmV2.CallOpts)
}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Caller) CurrentPeriodEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "currentPeriodEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Session) CurrentPeriodEnd() (*big.Int, error) {
	return _ZkEvmV2.Contract.CurrentPeriodEnd(&_ZkEvmV2.CallOpts)
}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2CallerSession) CurrentPeriodEnd() (*big.Int, error) {
	return _ZkEvmV2.Contract.CurrentPeriodEnd(&_ZkEvmV2.CallOpts)
}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Caller) CurrentTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "currentTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Session) CurrentTimestamp() (*big.Int, error) {
	return _ZkEvmV2.Contract.CurrentTimestamp(&_ZkEvmV2.CallOpts)
}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2CallerSession) CurrentTimestamp() (*big.Int, error) {
	return _ZkEvmV2.Contract.CurrentTimestamp(&_ZkEvmV2.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZkEvmV2.Contract.GetRoleAdmin(&_ZkEvmV2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZkEvmV2.Contract.GetRoleAdmin(&_ZkEvmV2.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZkEvmV2 *ZkEvmV2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZkEvmV2 *ZkEvmV2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZkEvmV2.Contract.HasRole(&_ZkEvmV2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZkEvmV2 *ZkEvmV2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZkEvmV2.Contract.HasRole(&_ZkEvmV2.CallOpts, role, account)
}

// InboxL2L1MessageStatus is a free data retrieval call binding the contract method 0x5c721a0c.
//
// Solidity: function inboxL2L1MessageStatus(bytes32 ) view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Caller) InboxL2L1MessageStatus(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "inboxL2L1MessageStatus", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InboxL2L1MessageStatus is a free data retrieval call binding the contract method 0x5c721a0c.
//
// Solidity: function inboxL2L1MessageStatus(bytes32 ) view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Session) InboxL2L1MessageStatus(arg0 [32]byte) (*big.Int, error) {
	return _ZkEvmV2.Contract.InboxL2L1MessageStatus(&_ZkEvmV2.CallOpts, arg0)
}

// InboxL2L1MessageStatus is a free data retrieval call binding the contract method 0x5c721a0c.
//
// Solidity: function inboxL2L1MessageStatus(bytes32 ) view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2CallerSession) InboxL2L1MessageStatus(arg0 [32]byte) (*big.Int, error) {
	return _ZkEvmV2.Contract.InboxL2L1MessageStatus(&_ZkEvmV2.CallOpts, arg0)
}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Caller) LimitInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "limitInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Session) LimitInWei() (*big.Int, error) {
	return _ZkEvmV2.Contract.LimitInWei(&_ZkEvmV2.CallOpts)
}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2CallerSession) LimitInWei() (*big.Int, error) {
	return _ZkEvmV2.Contract.LimitInWei(&_ZkEvmV2.CallOpts)
}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Caller) NextMessageNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "nextMessageNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Session) NextMessageNumber() (*big.Int, error) {
	return _ZkEvmV2.Contract.NextMessageNumber(&_ZkEvmV2.CallOpts)
}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2CallerSession) NextMessageNumber() (*big.Int, error) {
	return _ZkEvmV2.Contract.NextMessageNumber(&_ZkEvmV2.CallOpts)
}

// OutboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x3fc08b65.
//
// Solidity: function outboxL1L2MessageStatus(bytes32 ) view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Caller) OutboxL1L2MessageStatus(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "outboxL1L2MessageStatus", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x3fc08b65.
//
// Solidity: function outboxL1L2MessageStatus(bytes32 ) view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Session) OutboxL1L2MessageStatus(arg0 [32]byte) (*big.Int, error) {
	return _ZkEvmV2.Contract.OutboxL1L2MessageStatus(&_ZkEvmV2.CallOpts, arg0)
}

// OutboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x3fc08b65.
//
// Solidity: function outboxL1L2MessageStatus(bytes32 ) view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2CallerSession) OutboxL1L2MessageStatus(arg0 [32]byte) (*big.Int, error) {
	return _ZkEvmV2.Contract.OutboxL1L2MessageStatus(&_ZkEvmV2.CallOpts, arg0)
}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 ) view returns(bool)
func (_ZkEvmV2 *ZkEvmV2Caller) PauseTypeStatuses(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "pauseTypeStatuses", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 ) view returns(bool)
func (_ZkEvmV2 *ZkEvmV2Session) PauseTypeStatuses(arg0 [32]byte) (bool, error) {
	return _ZkEvmV2.Contract.PauseTypeStatuses(&_ZkEvmV2.CallOpts, arg0)
}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 ) view returns(bool)
func (_ZkEvmV2 *ZkEvmV2CallerSession) PauseTypeStatuses(arg0 [32]byte) (bool, error) {
	return _ZkEvmV2.Contract.PauseTypeStatuses(&_ZkEvmV2.CallOpts, arg0)
}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Caller) PeriodInSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "periodInSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2Session) PeriodInSeconds() (*big.Int, error) {
	return _ZkEvmV2.Contract.PeriodInSeconds(&_ZkEvmV2.CallOpts)
}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_ZkEvmV2 *ZkEvmV2CallerSession) PeriodInSeconds() (*big.Int, error) {
	return _ZkEvmV2.Contract.PeriodInSeconds(&_ZkEvmV2.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_ZkEvmV2 *ZkEvmV2Caller) Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_ZkEvmV2 *ZkEvmV2Session) Sender() (common.Address, error) {
	return _ZkEvmV2.Contract.Sender(&_ZkEvmV2.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_ZkEvmV2 *ZkEvmV2CallerSession) Sender() (common.Address, error) {
	return _ZkEvmV2.Contract.Sender(&_ZkEvmV2.CallOpts)
}

// StateRootHashes is a free data retrieval call binding the contract method 0x8be745d1.
//
// Solidity: function stateRootHashes(uint256 ) view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Caller) StateRootHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "stateRootHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateRootHashes is a free data retrieval call binding the contract method 0x8be745d1.
//
// Solidity: function stateRootHashes(uint256 ) view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2Session) StateRootHashes(arg0 *big.Int) ([32]byte, error) {
	return _ZkEvmV2.Contract.StateRootHashes(&_ZkEvmV2.CallOpts, arg0)
}

// StateRootHashes is a free data retrieval call binding the contract method 0x8be745d1.
//
// Solidity: function stateRootHashes(uint256 ) view returns(bytes32)
func (_ZkEvmV2 *ZkEvmV2CallerSession) StateRootHashes(arg0 *big.Int) ([32]byte, error) {
	return _ZkEvmV2.Contract.StateRootHashes(&_ZkEvmV2.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZkEvmV2 *ZkEvmV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZkEvmV2 *ZkEvmV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZkEvmV2.Contract.SupportsInterface(&_ZkEvmV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZkEvmV2 *ZkEvmV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZkEvmV2.Contract.SupportsInterface(&_ZkEvmV2.CallOpts, interfaceId)
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 ) view returns(address)
func (_ZkEvmV2 *ZkEvmV2Caller) Verifiers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ZkEvmV2.contract.Call(opts, &out, "verifiers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 ) view returns(address)
func (_ZkEvmV2 *ZkEvmV2Session) Verifiers(arg0 *big.Int) (common.Address, error) {
	return _ZkEvmV2.Contract.Verifiers(&_ZkEvmV2.CallOpts, arg0)
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 ) view returns(address)
func (_ZkEvmV2 *ZkEvmV2CallerSession) Verifiers(arg0 *big.Int) (common.Address, error) {
	return _ZkEvmV2.Contract.Verifiers(&_ZkEvmV2.CallOpts, arg0)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) ClaimMessage(opts *bind.TransactOpts, _from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "claimMessage", _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_ZkEvmV2 *ZkEvmV2Session) ClaimMessage(_from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.ClaimMessage(&_ZkEvmV2.TransactOpts, _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) ClaimMessage(_from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.ClaimMessage(&_ZkEvmV2.TransactOpts, _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x4165d6dd.
//
// Solidity: function finalizeBlocks((bytes32,uint32,bytes[],bytes32[],bytes,uint16[])[] _blocksData, bytes _proof, uint256 _proofType, bytes32 _parentStateRootHash) returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) FinalizeBlocks(opts *bind.TransactOpts, _blocksData []IZkEvmV2BlockData, _proof []byte, _proofType *big.Int, _parentStateRootHash [32]byte) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "finalizeBlocks", _blocksData, _proof, _proofType, _parentStateRootHash)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x4165d6dd.
//
// Solidity: function finalizeBlocks((bytes32,uint32,bytes[],bytes32[],bytes,uint16[])[] _blocksData, bytes _proof, uint256 _proofType, bytes32 _parentStateRootHash) returns()
func (_ZkEvmV2 *ZkEvmV2Session) FinalizeBlocks(_blocksData []IZkEvmV2BlockData, _proof []byte, _proofType *big.Int, _parentStateRootHash [32]byte) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.FinalizeBlocks(&_ZkEvmV2.TransactOpts, _blocksData, _proof, _proofType, _parentStateRootHash)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x4165d6dd.
//
// Solidity: function finalizeBlocks((bytes32,uint32,bytes[],bytes32[],bytes,uint16[])[] _blocksData, bytes _proof, uint256 _proofType, bytes32 _parentStateRootHash) returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) FinalizeBlocks(_blocksData []IZkEvmV2BlockData, _proof []byte, _proofType *big.Int, _parentStateRootHash [32]byte) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.FinalizeBlocks(&_ZkEvmV2.TransactOpts, _blocksData, _proof, _proofType, _parentStateRootHash)
}

// FinalizeBlocksWithoutProof is a paid mutator transaction binding the contract method 0x90dad3f6.
//
// Solidity: function finalizeBlocksWithoutProof((bytes32,uint32,bytes[],bytes32[],bytes,uint16[])[] _blocksData) returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) FinalizeBlocksWithoutProof(opts *bind.TransactOpts, _blocksData []IZkEvmV2BlockData) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "finalizeBlocksWithoutProof", _blocksData)
}

// FinalizeBlocksWithoutProof is a paid mutator transaction binding the contract method 0x90dad3f6.
//
// Solidity: function finalizeBlocksWithoutProof((bytes32,uint32,bytes[],bytes32[],bytes,uint16[])[] _blocksData) returns()
func (_ZkEvmV2 *ZkEvmV2Session) FinalizeBlocksWithoutProof(_blocksData []IZkEvmV2BlockData) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.FinalizeBlocksWithoutProof(&_ZkEvmV2.TransactOpts, _blocksData)
}

// FinalizeBlocksWithoutProof is a paid mutator transaction binding the contract method 0x90dad3f6.
//
// Solidity: function finalizeBlocksWithoutProof((bytes32,uint32,bytes[],bytes32[],bytes,uint16[])[] _blocksData) returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) FinalizeBlocksWithoutProof(_blocksData []IZkEvmV2BlockData) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.FinalizeBlocksWithoutProof(&_ZkEvmV2.TransactOpts, _blocksData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZkEvmV2 *ZkEvmV2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.GrantRole(&_ZkEvmV2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.GrantRole(&_ZkEvmV2.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x7973ead6.
//
// Solidity: function initialize(bytes32 _initialStateRootHash, uint256 _initialL2BlockNumber, address _defaultVerifier, address _securityCouncil, address[] _operators, uint256 _rateLimitPeriodInSeconds, uint256 _rateLimitAmountInWei) returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) Initialize(opts *bind.TransactOpts, _initialStateRootHash [32]byte, _initialL2BlockNumber *big.Int, _defaultVerifier common.Address, _securityCouncil common.Address, _operators []common.Address, _rateLimitPeriodInSeconds *big.Int, _rateLimitAmountInWei *big.Int) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "initialize", _initialStateRootHash, _initialL2BlockNumber, _defaultVerifier, _securityCouncil, _operators, _rateLimitPeriodInSeconds, _rateLimitAmountInWei)
}

// Initialize is a paid mutator transaction binding the contract method 0x7973ead6.
//
// Solidity: function initialize(bytes32 _initialStateRootHash, uint256 _initialL2BlockNumber, address _defaultVerifier, address _securityCouncil, address[] _operators, uint256 _rateLimitPeriodInSeconds, uint256 _rateLimitAmountInWei) returns()
func (_ZkEvmV2 *ZkEvmV2Session) Initialize(_initialStateRootHash [32]byte, _initialL2BlockNumber *big.Int, _defaultVerifier common.Address, _securityCouncil common.Address, _operators []common.Address, _rateLimitPeriodInSeconds *big.Int, _rateLimitAmountInWei *big.Int) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.Initialize(&_ZkEvmV2.TransactOpts, _initialStateRootHash, _initialL2BlockNumber, _defaultVerifier, _securityCouncil, _operators, _rateLimitPeriodInSeconds, _rateLimitAmountInWei)
}

// Initialize is a paid mutator transaction binding the contract method 0x7973ead6.
//
// Solidity: function initialize(bytes32 _initialStateRootHash, uint256 _initialL2BlockNumber, address _defaultVerifier, address _securityCouncil, address[] _operators, uint256 _rateLimitPeriodInSeconds, uint256 _rateLimitAmountInWei) returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) Initialize(_initialStateRootHash [32]byte, _initialL2BlockNumber *big.Int, _defaultVerifier common.Address, _securityCouncil common.Address, _operators []common.Address, _rateLimitPeriodInSeconds *big.Int, _rateLimitAmountInWei *big.Int) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.Initialize(&_ZkEvmV2.TransactOpts, _initialStateRootHash, _initialL2BlockNumber, _defaultVerifier, _securityCouncil, _operators, _rateLimitPeriodInSeconds, _rateLimitAmountInWei)
}

// PauseByType is a paid mutator transaction binding the contract method 0x8264bd82.
//
// Solidity: function pauseByType(bytes32 _pauseType) returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) PauseByType(opts *bind.TransactOpts, _pauseType [32]byte) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "pauseByType", _pauseType)
}

// PauseByType is a paid mutator transaction binding the contract method 0x8264bd82.
//
// Solidity: function pauseByType(bytes32 _pauseType) returns()
func (_ZkEvmV2 *ZkEvmV2Session) PauseByType(_pauseType [32]byte) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.PauseByType(&_ZkEvmV2.TransactOpts, _pauseType)
}

// PauseByType is a paid mutator transaction binding the contract method 0x8264bd82.
//
// Solidity: function pauseByType(bytes32 _pauseType) returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) PauseByType(_pauseType [32]byte) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.PauseByType(&_ZkEvmV2.TransactOpts, _pauseType)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZkEvmV2 *ZkEvmV2Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.RenounceRole(&_ZkEvmV2.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.RenounceRole(&_ZkEvmV2.TransactOpts, role, account)
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) ResetAmountUsedInPeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "resetAmountUsedInPeriod")
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_ZkEvmV2 *ZkEvmV2Session) ResetAmountUsedInPeriod() (*types.Transaction, error) {
	return _ZkEvmV2.Contract.ResetAmountUsedInPeriod(&_ZkEvmV2.TransactOpts)
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) ResetAmountUsedInPeriod() (*types.Transaction, error) {
	return _ZkEvmV2.Contract.ResetAmountUsedInPeriod(&_ZkEvmV2.TransactOpts)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) ResetRateLimitAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "resetRateLimitAmount", _amount)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_ZkEvmV2 *ZkEvmV2Session) ResetRateLimitAmount(_amount *big.Int) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.ResetRateLimitAmount(&_ZkEvmV2.TransactOpts, _amount)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) ResetRateLimitAmount(_amount *big.Int) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.ResetRateLimitAmount(&_ZkEvmV2.TransactOpts, _amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZkEvmV2 *ZkEvmV2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.RevokeRole(&_ZkEvmV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.RevokeRole(&_ZkEvmV2.TransactOpts, role, account)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) SendMessage(opts *bind.TransactOpts, _to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "sendMessage", _to, _fee, _calldata)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_ZkEvmV2 *ZkEvmV2Session) SendMessage(_to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.SendMessage(&_ZkEvmV2.TransactOpts, _to, _fee, _calldata)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) SendMessage(_to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.SendMessage(&_ZkEvmV2.TransactOpts, _to, _fee, _calldata)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0xc2116974.
//
// Solidity: function setVerifierAddress(address _newVerifierAddress, uint256 _proofType) returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) SetVerifierAddress(opts *bind.TransactOpts, _newVerifierAddress common.Address, _proofType *big.Int) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "setVerifierAddress", _newVerifierAddress, _proofType)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0xc2116974.
//
// Solidity: function setVerifierAddress(address _newVerifierAddress, uint256 _proofType) returns()
func (_ZkEvmV2 *ZkEvmV2Session) SetVerifierAddress(_newVerifierAddress common.Address, _proofType *big.Int) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.SetVerifierAddress(&_ZkEvmV2.TransactOpts, _newVerifierAddress, _proofType)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0xc2116974.
//
// Solidity: function setVerifierAddress(address _newVerifierAddress, uint256 _proofType) returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) SetVerifierAddress(_newVerifierAddress common.Address, _proofType *big.Int) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.SetVerifierAddress(&_ZkEvmV2.TransactOpts, _newVerifierAddress, _proofType)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0xb45a4f2c.
//
// Solidity: function unPauseByType(bytes32 _pauseType) returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) UnPauseByType(opts *bind.TransactOpts, _pauseType [32]byte) (*types.Transaction, error) {
	return _ZkEvmV2.contract.Transact(opts, "unPauseByType", _pauseType)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0xb45a4f2c.
//
// Solidity: function unPauseByType(bytes32 _pauseType) returns()
func (_ZkEvmV2 *ZkEvmV2Session) UnPauseByType(_pauseType [32]byte) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.UnPauseByType(&_ZkEvmV2.TransactOpts, _pauseType)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0xb45a4f2c.
//
// Solidity: function unPauseByType(bytes32 _pauseType) returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) UnPauseByType(_pauseType [32]byte) (*types.Transaction, error) {
	return _ZkEvmV2.Contract.UnPauseByType(&_ZkEvmV2.TransactOpts, _pauseType)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZkEvmV2 *ZkEvmV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkEvmV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZkEvmV2 *ZkEvmV2Session) Receive() (*types.Transaction, error) {
	return _ZkEvmV2.Contract.Receive(&_ZkEvmV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZkEvmV2 *ZkEvmV2TransactorSession) Receive() (*types.Transaction, error) {
	return _ZkEvmV2.Contract.Receive(&_ZkEvmV2.TransactOpts)
}

// ZkEvmV2AmountUsedInPeriodResetIterator is returned from FilterAmountUsedInPeriodReset and is used to iterate over the raw logs and unpacked data for AmountUsedInPeriodReset events raised by the ZkEvmV2 contract.
type ZkEvmV2AmountUsedInPeriodResetIterator struct {
	Event *ZkEvmV2AmountUsedInPeriodReset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2AmountUsedInPeriodResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2AmountUsedInPeriodReset)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2AmountUsedInPeriodReset)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2AmountUsedInPeriodResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2AmountUsedInPeriodResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2AmountUsedInPeriodReset represents a AmountUsedInPeriodReset event raised by the ZkEvmV2 contract.
type ZkEvmV2AmountUsedInPeriodReset struct {
	ResettingAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAmountUsedInPeriodReset is a free log retrieval operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterAmountUsedInPeriodReset(opts *bind.FilterOpts, resettingAddress []common.Address) (*ZkEvmV2AmountUsedInPeriodResetIterator, error) {

	var resettingAddressRule []interface{}
	for _, resettingAddressItem := range resettingAddress {
		resettingAddressRule = append(resettingAddressRule, resettingAddressItem)
	}

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "AmountUsedInPeriodReset", resettingAddressRule)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2AmountUsedInPeriodResetIterator{contract: _ZkEvmV2.contract, event: "AmountUsedInPeriodReset", logs: logs, sub: sub}, nil
}

// WatchAmountUsedInPeriodReset is a free log subscription operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchAmountUsedInPeriodReset(opts *bind.WatchOpts, sink chan<- *ZkEvmV2AmountUsedInPeriodReset, resettingAddress []common.Address) (event.Subscription, error) {

	var resettingAddressRule []interface{}
	for _, resettingAddressItem := range resettingAddress {
		resettingAddressRule = append(resettingAddressRule, resettingAddressItem)
	}

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "AmountUsedInPeriodReset", resettingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2AmountUsedInPeriodReset)
				if err := _ZkEvmV2.contract.UnpackLog(event, "AmountUsedInPeriodReset", log); err != nil {
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

// ParseAmountUsedInPeriodReset is a log parse operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseAmountUsedInPeriodReset(log types.Log) (*ZkEvmV2AmountUsedInPeriodReset, error) {
	event := new(ZkEvmV2AmountUsedInPeriodReset)
	if err := _ZkEvmV2.contract.UnpackLog(event, "AmountUsedInPeriodReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2BlockFinalizedIterator is returned from FilterBlockFinalized and is used to iterate over the raw logs and unpacked data for BlockFinalized events raised by the ZkEvmV2 contract.
type ZkEvmV2BlockFinalizedIterator struct {
	Event *ZkEvmV2BlockFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2BlockFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2BlockFinalized)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2BlockFinalized)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2BlockFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2BlockFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2BlockFinalized represents a BlockFinalized event raised by the ZkEvmV2 contract.
type ZkEvmV2BlockFinalized struct {
	BlockNumber   *big.Int
	StateRootHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBlockFinalized is a free log retrieval operation binding the contract event 0xf2c535759092d16e9334a11dd9b52eca543f1d9cca5ba9d16c472aef009de432.
//
// Solidity: event BlockFinalized(uint256 indexed blockNumber, bytes32 indexed stateRootHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterBlockFinalized(opts *bind.FilterOpts, blockNumber []*big.Int, stateRootHash [][32]byte) (*ZkEvmV2BlockFinalizedIterator, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var stateRootHashRule []interface{}
	for _, stateRootHashItem := range stateRootHash {
		stateRootHashRule = append(stateRootHashRule, stateRootHashItem)
	}

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "BlockFinalized", blockNumberRule, stateRootHashRule)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2BlockFinalizedIterator{contract: _ZkEvmV2.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xf2c535759092d16e9334a11dd9b52eca543f1d9cca5ba9d16c472aef009de432.
//
// Solidity: event BlockFinalized(uint256 indexed blockNumber, bytes32 indexed stateRootHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchBlockFinalized(opts *bind.WatchOpts, sink chan<- *ZkEvmV2BlockFinalized, blockNumber []*big.Int, stateRootHash [][32]byte) (event.Subscription, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var stateRootHashRule []interface{}
	for _, stateRootHashItem := range stateRootHash {
		stateRootHashRule = append(stateRootHashRule, stateRootHashItem)
	}

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "BlockFinalized", blockNumberRule, stateRootHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2BlockFinalized)
				if err := _ZkEvmV2.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
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

// ParseBlockFinalized is a log parse operation binding the contract event 0xf2c535759092d16e9334a11dd9b52eca543f1d9cca5ba9d16c472aef009de432.
//
// Solidity: event BlockFinalized(uint256 indexed blockNumber, bytes32 indexed stateRootHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseBlockFinalized(log types.Log) (*ZkEvmV2BlockFinalized, error) {
	event := new(ZkEvmV2BlockFinalized)
	if err := _ZkEvmV2.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2BlocksVerificationDoneIterator is returned from FilterBlocksVerificationDone and is used to iterate over the raw logs and unpacked data for BlocksVerificationDone events raised by the ZkEvmV2 contract.
type ZkEvmV2BlocksVerificationDoneIterator struct {
	Event *ZkEvmV2BlocksVerificationDone // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2BlocksVerificationDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2BlocksVerificationDone)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2BlocksVerificationDone)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2BlocksVerificationDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2BlocksVerificationDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2BlocksVerificationDone represents a BlocksVerificationDone event raised by the ZkEvmV2 contract.
type ZkEvmV2BlocksVerificationDone struct {
	LastBlockFinalized *big.Int
	StartingRootHash   [32]byte
	FinalRootHash      [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBlocksVerificationDone is a free log retrieval operation binding the contract event 0x5c885a794662ebe3b08ae0874fc2c88b5343b0223ba9cd2cad92b69c0d0c901f.
//
// Solidity: event BlocksVerificationDone(uint256 indexed lastBlockFinalized, bytes32 startingRootHash, bytes32 finalRootHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterBlocksVerificationDone(opts *bind.FilterOpts, lastBlockFinalized []*big.Int) (*ZkEvmV2BlocksVerificationDoneIterator, error) {

	var lastBlockFinalizedRule []interface{}
	for _, lastBlockFinalizedItem := range lastBlockFinalized {
		lastBlockFinalizedRule = append(lastBlockFinalizedRule, lastBlockFinalizedItem)
	}

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "BlocksVerificationDone", lastBlockFinalizedRule)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2BlocksVerificationDoneIterator{contract: _ZkEvmV2.contract, event: "BlocksVerificationDone", logs: logs, sub: sub}, nil
}

// WatchBlocksVerificationDone is a free log subscription operation binding the contract event 0x5c885a794662ebe3b08ae0874fc2c88b5343b0223ba9cd2cad92b69c0d0c901f.
//
// Solidity: event BlocksVerificationDone(uint256 indexed lastBlockFinalized, bytes32 startingRootHash, bytes32 finalRootHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchBlocksVerificationDone(opts *bind.WatchOpts, sink chan<- *ZkEvmV2BlocksVerificationDone, lastBlockFinalized []*big.Int) (event.Subscription, error) {

	var lastBlockFinalizedRule []interface{}
	for _, lastBlockFinalizedItem := range lastBlockFinalized {
		lastBlockFinalizedRule = append(lastBlockFinalizedRule, lastBlockFinalizedItem)
	}

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "BlocksVerificationDone", lastBlockFinalizedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2BlocksVerificationDone)
				if err := _ZkEvmV2.contract.UnpackLog(event, "BlocksVerificationDone", log); err != nil {
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

// ParseBlocksVerificationDone is a log parse operation binding the contract event 0x5c885a794662ebe3b08ae0874fc2c88b5343b0223ba9cd2cad92b69c0d0c901f.
//
// Solidity: event BlocksVerificationDone(uint256 indexed lastBlockFinalized, bytes32 startingRootHash, bytes32 finalRootHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseBlocksVerificationDone(log types.Log) (*ZkEvmV2BlocksVerificationDone, error) {
	event := new(ZkEvmV2BlocksVerificationDone)
	if err := _ZkEvmV2.contract.UnpackLog(event, "BlocksVerificationDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZkEvmV2 contract.
type ZkEvmV2InitializedIterator struct {
	Event *ZkEvmV2Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2Initialized)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2Initialized)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2Initialized represents a Initialized event raised by the ZkEvmV2 contract.
type ZkEvmV2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*ZkEvmV2InitializedIterator, error) {

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2InitializedIterator{contract: _ZkEvmV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZkEvmV2Initialized) (event.Subscription, error) {

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2Initialized)
				if err := _ZkEvmV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseInitialized(log types.Log) (*ZkEvmV2Initialized, error) {
	event := new(ZkEvmV2Initialized)
	if err := _ZkEvmV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2L1L2MessagesReceivedOnL2Iterator is returned from FilterL1L2MessagesReceivedOnL2 and is used to iterate over the raw logs and unpacked data for L1L2MessagesReceivedOnL2 events raised by the ZkEvmV2 contract.
type ZkEvmV2L1L2MessagesReceivedOnL2Iterator struct {
	Event *ZkEvmV2L1L2MessagesReceivedOnL2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2L1L2MessagesReceivedOnL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2L1L2MessagesReceivedOnL2)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2L1L2MessagesReceivedOnL2)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2L1L2MessagesReceivedOnL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2L1L2MessagesReceivedOnL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2L1L2MessagesReceivedOnL2 represents a L1L2MessagesReceivedOnL2 event raised by the ZkEvmV2 contract.
type ZkEvmV2L1L2MessagesReceivedOnL2 struct {
	MessageHashes [][32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterL1L2MessagesReceivedOnL2 is a free log retrieval operation binding the contract event 0x95e84bb4317676921a29fd1d13f8f0153508473b899c12b3cd08314348801d64.
//
// Solidity: event L1L2MessagesReceivedOnL2(bytes32[] messageHashes)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterL1L2MessagesReceivedOnL2(opts *bind.FilterOpts) (*ZkEvmV2L1L2MessagesReceivedOnL2Iterator, error) {

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "L1L2MessagesReceivedOnL2")
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2L1L2MessagesReceivedOnL2Iterator{contract: _ZkEvmV2.contract, event: "L1L2MessagesReceivedOnL2", logs: logs, sub: sub}, nil
}

// WatchL1L2MessagesReceivedOnL2 is a free log subscription operation binding the contract event 0x95e84bb4317676921a29fd1d13f8f0153508473b899c12b3cd08314348801d64.
//
// Solidity: event L1L2MessagesReceivedOnL2(bytes32[] messageHashes)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchL1L2MessagesReceivedOnL2(opts *bind.WatchOpts, sink chan<- *ZkEvmV2L1L2MessagesReceivedOnL2) (event.Subscription, error) {

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "L1L2MessagesReceivedOnL2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2L1L2MessagesReceivedOnL2)
				if err := _ZkEvmV2.contract.UnpackLog(event, "L1L2MessagesReceivedOnL2", log); err != nil {
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

// ParseL1L2MessagesReceivedOnL2 is a log parse operation binding the contract event 0x95e84bb4317676921a29fd1d13f8f0153508473b899c12b3cd08314348801d64.
//
// Solidity: event L1L2MessagesReceivedOnL2(bytes32[] messageHashes)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseL1L2MessagesReceivedOnL2(log types.Log) (*ZkEvmV2L1L2MessagesReceivedOnL2, error) {
	event := new(ZkEvmV2L1L2MessagesReceivedOnL2)
	if err := _ZkEvmV2.contract.UnpackLog(event, "L1L2MessagesReceivedOnL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2L2L1MessageHashAddedToInboxIterator is returned from FilterL2L1MessageHashAddedToInbox and is used to iterate over the raw logs and unpacked data for L2L1MessageHashAddedToInbox events raised by the ZkEvmV2 contract.
type ZkEvmV2L2L1MessageHashAddedToInboxIterator struct {
	Event *ZkEvmV2L2L1MessageHashAddedToInbox // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2L2L1MessageHashAddedToInboxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2L2L1MessageHashAddedToInbox)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2L2L1MessageHashAddedToInbox)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2L2L1MessageHashAddedToInboxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2L2L1MessageHashAddedToInboxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2L2L1MessageHashAddedToInbox represents a L2L1MessageHashAddedToInbox event raised by the ZkEvmV2 contract.
type ZkEvmV2L2L1MessageHashAddedToInbox struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterL2L1MessageHashAddedToInbox is a free log retrieval operation binding the contract event 0x810484e22f73d8f099aaee1edb851ec6be6d84d43045d0a7803e5f7b3612edce.
//
// Solidity: event L2L1MessageHashAddedToInbox(bytes32 indexed messageHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterL2L1MessageHashAddedToInbox(opts *bind.FilterOpts, messageHash [][32]byte) (*ZkEvmV2L2L1MessageHashAddedToInboxIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "L2L1MessageHashAddedToInbox", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2L2L1MessageHashAddedToInboxIterator{contract: _ZkEvmV2.contract, event: "L2L1MessageHashAddedToInbox", logs: logs, sub: sub}, nil
}

// WatchL2L1MessageHashAddedToInbox is a free log subscription operation binding the contract event 0x810484e22f73d8f099aaee1edb851ec6be6d84d43045d0a7803e5f7b3612edce.
//
// Solidity: event L2L1MessageHashAddedToInbox(bytes32 indexed messageHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchL2L1MessageHashAddedToInbox(opts *bind.WatchOpts, sink chan<- *ZkEvmV2L2L1MessageHashAddedToInbox, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "L2L1MessageHashAddedToInbox", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2L2L1MessageHashAddedToInbox)
				if err := _ZkEvmV2.contract.UnpackLog(event, "L2L1MessageHashAddedToInbox", log); err != nil {
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

// ParseL2L1MessageHashAddedToInbox is a log parse operation binding the contract event 0x810484e22f73d8f099aaee1edb851ec6be6d84d43045d0a7803e5f7b3612edce.
//
// Solidity: event L2L1MessageHashAddedToInbox(bytes32 indexed messageHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseL2L1MessageHashAddedToInbox(log types.Log) (*ZkEvmV2L2L1MessageHashAddedToInbox, error) {
	event := new(ZkEvmV2L2L1MessageHashAddedToInbox)
	if err := _ZkEvmV2.contract.UnpackLog(event, "L2L1MessageHashAddedToInbox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2LimitAmountChangedIterator is returned from FilterLimitAmountChanged and is used to iterate over the raw logs and unpacked data for LimitAmountChanged events raised by the ZkEvmV2 contract.
type ZkEvmV2LimitAmountChangedIterator struct {
	Event *ZkEvmV2LimitAmountChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2LimitAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2LimitAmountChanged)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2LimitAmountChanged)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2LimitAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2LimitAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2LimitAmountChanged represents a LimitAmountChanged event raised by the ZkEvmV2 contract.
type ZkEvmV2LimitAmountChanged struct {
	AmountChangeBy           common.Address
	Amount                   *big.Int
	AmountUsedLoweredToLimit bool
	UsedAmountResetToZero    bool
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterLimitAmountChanged is a free log retrieval operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterLimitAmountChanged(opts *bind.FilterOpts, amountChangeBy []common.Address) (*ZkEvmV2LimitAmountChangedIterator, error) {

	var amountChangeByRule []interface{}
	for _, amountChangeByItem := range amountChangeBy {
		amountChangeByRule = append(amountChangeByRule, amountChangeByItem)
	}

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "LimitAmountChanged", amountChangeByRule)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2LimitAmountChangedIterator{contract: _ZkEvmV2.contract, event: "LimitAmountChanged", logs: logs, sub: sub}, nil
}

// WatchLimitAmountChanged is a free log subscription operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchLimitAmountChanged(opts *bind.WatchOpts, sink chan<- *ZkEvmV2LimitAmountChanged, amountChangeBy []common.Address) (event.Subscription, error) {

	var amountChangeByRule []interface{}
	for _, amountChangeByItem := range amountChangeBy {
		amountChangeByRule = append(amountChangeByRule, amountChangeByItem)
	}

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "LimitAmountChanged", amountChangeByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2LimitAmountChanged)
				if err := _ZkEvmV2.contract.UnpackLog(event, "LimitAmountChanged", log); err != nil {
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

// ParseLimitAmountChanged is a log parse operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseLimitAmountChanged(log types.Log) (*ZkEvmV2LimitAmountChanged, error) {
	event := new(ZkEvmV2LimitAmountChanged)
	if err := _ZkEvmV2.contract.UnpackLog(event, "LimitAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2MessageClaimedIterator is returned from FilterMessageClaimed and is used to iterate over the raw logs and unpacked data for MessageClaimed events raised by the ZkEvmV2 contract.
type ZkEvmV2MessageClaimedIterator struct {
	Event *ZkEvmV2MessageClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2MessageClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2MessageClaimed)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2MessageClaimed)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2MessageClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2MessageClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2MessageClaimed represents a MessageClaimed event raised by the ZkEvmV2 contract.
type ZkEvmV2MessageClaimed struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageClaimed is a free log retrieval operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterMessageClaimed(opts *bind.FilterOpts, _messageHash [][32]byte) (*ZkEvmV2MessageClaimedIterator, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2MessageClaimedIterator{contract: _ZkEvmV2.contract, event: "MessageClaimed", logs: logs, sub: sub}, nil
}

// WatchMessageClaimed is a free log subscription operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchMessageClaimed(opts *bind.WatchOpts, sink chan<- *ZkEvmV2MessageClaimed, _messageHash [][32]byte) (event.Subscription, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2MessageClaimed)
				if err := _ZkEvmV2.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
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

// ParseMessageClaimed is a log parse operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseMessageClaimed(log types.Log) (*ZkEvmV2MessageClaimed, error) {
	event := new(ZkEvmV2MessageClaimed)
	if err := _ZkEvmV2.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2MessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the ZkEvmV2 contract.
type ZkEvmV2MessageSentIterator struct {
	Event *ZkEvmV2MessageSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2MessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2MessageSent)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2MessageSent)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2MessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2MessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2MessageSent represents a MessageSent event raised by the ZkEvmV2 contract.
type ZkEvmV2MessageSent struct {
	From        common.Address
	To          common.Address
	Fee         *big.Int
	Value       *big.Int
	Nonce       *big.Int
	Calldata    []byte
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterMessageSent(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (*ZkEvmV2MessageSentIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2MessageSentIterator{contract: _ZkEvmV2.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *ZkEvmV2MessageSent, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2MessageSent)
				if err := _ZkEvmV2.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseMessageSent(log types.Log) (*ZkEvmV2MessageSent, error) {
	event := new(ZkEvmV2MessageSent)
	if err := _ZkEvmV2.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ZkEvmV2 contract.
type ZkEvmV2PausedIterator struct {
	Event *ZkEvmV2Paused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2Paused)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2Paused)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2Paused represents a Paused event raised by the ZkEvmV2 contract.
type ZkEvmV2Paused struct {
	MessageSender common.Address
	PauseType     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xc343aefb875672fb1857ecda2bdf9fa822ff1e924e3714f6a3d88c5199dee261.
//
// Solidity: event Paused(address messageSender, bytes32 pauseType)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterPaused(opts *bind.FilterOpts) (*ZkEvmV2PausedIterator, error) {

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2PausedIterator{contract: _ZkEvmV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xc343aefb875672fb1857ecda2bdf9fa822ff1e924e3714f6a3d88c5199dee261.
//
// Solidity: event Paused(address messageSender, bytes32 pauseType)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ZkEvmV2Paused) (event.Subscription, error) {

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2Paused)
				if err := _ZkEvmV2.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xc343aefb875672fb1857ecda2bdf9fa822ff1e924e3714f6a3d88c5199dee261.
//
// Solidity: event Paused(address messageSender, bytes32 pauseType)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParsePaused(log types.Log) (*ZkEvmV2Paused, error) {
	event := new(ZkEvmV2Paused)
	if err := _ZkEvmV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZkEvmV2 contract.
type ZkEvmV2RoleAdminChangedIterator struct {
	Event *ZkEvmV2RoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2RoleAdminChanged)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2RoleAdminChanged)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2RoleAdminChanged represents a RoleAdminChanged event raised by the ZkEvmV2 contract.
type ZkEvmV2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZkEvmV2RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2RoleAdminChangedIterator{contract: _ZkEvmV2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZkEvmV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2RoleAdminChanged)
				if err := _ZkEvmV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseRoleAdminChanged(log types.Log) (*ZkEvmV2RoleAdminChanged, error) {
	event := new(ZkEvmV2RoleAdminChanged)
	if err := _ZkEvmV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZkEvmV2 contract.
type ZkEvmV2RoleGrantedIterator struct {
	Event *ZkEvmV2RoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2RoleGranted)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2RoleGranted)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2RoleGranted represents a RoleGranted event raised by the ZkEvmV2 contract.
type ZkEvmV2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZkEvmV2RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2RoleGrantedIterator{contract: _ZkEvmV2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZkEvmV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2RoleGranted)
				if err := _ZkEvmV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseRoleGranted(log types.Log) (*ZkEvmV2RoleGranted, error) {
	event := new(ZkEvmV2RoleGranted)
	if err := _ZkEvmV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZkEvmV2 contract.
type ZkEvmV2RoleRevokedIterator struct {
	Event *ZkEvmV2RoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2RoleRevoked)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2RoleRevoked)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2RoleRevoked represents a RoleRevoked event raised by the ZkEvmV2 contract.
type ZkEvmV2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZkEvmV2RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2RoleRevokedIterator{contract: _ZkEvmV2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZkEvmV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2RoleRevoked)
				if err := _ZkEvmV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseRoleRevoked(log types.Log) (*ZkEvmV2RoleRevoked, error) {
	event := new(ZkEvmV2RoleRevoked)
	if err := _ZkEvmV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2UnPausedIterator is returned from FilterUnPaused and is used to iterate over the raw logs and unpacked data for UnPaused events raised by the ZkEvmV2 contract.
type ZkEvmV2UnPausedIterator struct {
	Event *ZkEvmV2UnPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2UnPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2UnPaused)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2UnPaused)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2UnPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2UnPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2UnPaused represents a UnPaused event raised by the ZkEvmV2 contract.
type ZkEvmV2UnPaused struct {
	MessageSender common.Address
	PauseType     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnPaused is a free log retrieval operation binding the contract event 0xb54c82d9fabaaa460c07181bb36c08c0e72d79293e77a42ac273c81d2a54281b.
//
// Solidity: event UnPaused(address messageSender, bytes32 pauseType)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterUnPaused(opts *bind.FilterOpts) (*ZkEvmV2UnPausedIterator, error) {

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2UnPausedIterator{contract: _ZkEvmV2.contract, event: "UnPaused", logs: logs, sub: sub}, nil
}

// WatchUnPaused is a free log subscription operation binding the contract event 0xb54c82d9fabaaa460c07181bb36c08c0e72d79293e77a42ac273c81d2a54281b.
//
// Solidity: event UnPaused(address messageSender, bytes32 pauseType)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchUnPaused(opts *bind.WatchOpts, sink chan<- *ZkEvmV2UnPaused) (event.Subscription, error) {

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2UnPaused)
				if err := _ZkEvmV2.contract.UnpackLog(event, "UnPaused", log); err != nil {
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

// ParseUnPaused is a log parse operation binding the contract event 0xb54c82d9fabaaa460c07181bb36c08c0e72d79293e77a42ac273c81d2a54281b.
//
// Solidity: event UnPaused(address messageSender, bytes32 pauseType)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseUnPaused(log types.Log) (*ZkEvmV2UnPaused, error) {
	event := new(ZkEvmV2UnPaused)
	if err := _ZkEvmV2.contract.UnpackLog(event, "UnPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkEvmV2VerifierAddressChangedIterator is returned from FilterVerifierAddressChanged and is used to iterate over the raw logs and unpacked data for VerifierAddressChanged events raised by the ZkEvmV2 contract.
type ZkEvmV2VerifierAddressChangedIterator struct {
	Event *ZkEvmV2VerifierAddressChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkEvmV2VerifierAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkEvmV2VerifierAddressChanged)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkEvmV2VerifierAddressChanged)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkEvmV2VerifierAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkEvmV2VerifierAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkEvmV2VerifierAddressChanged represents a VerifierAddressChanged event raised by the ZkEvmV2 contract.
type ZkEvmV2VerifierAddressChanged struct {
	VerifierAddress common.Address
	ProofType       *big.Int
	VerifierSetBy   common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVerifierAddressChanged is a free log retrieval operation binding the contract event 0x4ea861139068e7701a770b8975bb54b6f8f446897fac206dd29424035b4a61eb.
//
// Solidity: event VerifierAddressChanged(address indexed verifierAddress, uint256 indexed proofType, address indexed verifierSetBy)
func (_ZkEvmV2 *ZkEvmV2Filterer) FilterVerifierAddressChanged(opts *bind.FilterOpts, verifierAddress []common.Address, proofType []*big.Int, verifierSetBy []common.Address) (*ZkEvmV2VerifierAddressChangedIterator, error) {

	var verifierAddressRule []interface{}
	for _, verifierAddressItem := range verifierAddress {
		verifierAddressRule = append(verifierAddressRule, verifierAddressItem)
	}
	var proofTypeRule []interface{}
	for _, proofTypeItem := range proofType {
		proofTypeRule = append(proofTypeRule, proofTypeItem)
	}
	var verifierSetByRule []interface{}
	for _, verifierSetByItem := range verifierSetBy {
		verifierSetByRule = append(verifierSetByRule, verifierSetByItem)
	}

	logs, sub, err := _ZkEvmV2.contract.FilterLogs(opts, "VerifierAddressChanged", verifierAddressRule, proofTypeRule, verifierSetByRule)
	if err != nil {
		return nil, err
	}
	return &ZkEvmV2VerifierAddressChangedIterator{contract: _ZkEvmV2.contract, event: "VerifierAddressChanged", logs: logs, sub: sub}, nil
}

// WatchVerifierAddressChanged is a free log subscription operation binding the contract event 0x4ea861139068e7701a770b8975bb54b6f8f446897fac206dd29424035b4a61eb.
//
// Solidity: event VerifierAddressChanged(address indexed verifierAddress, uint256 indexed proofType, address indexed verifierSetBy)
func (_ZkEvmV2 *ZkEvmV2Filterer) WatchVerifierAddressChanged(opts *bind.WatchOpts, sink chan<- *ZkEvmV2VerifierAddressChanged, verifierAddress []common.Address, proofType []*big.Int, verifierSetBy []common.Address) (event.Subscription, error) {

	var verifierAddressRule []interface{}
	for _, verifierAddressItem := range verifierAddress {
		verifierAddressRule = append(verifierAddressRule, verifierAddressItem)
	}
	var proofTypeRule []interface{}
	for _, proofTypeItem := range proofType {
		proofTypeRule = append(proofTypeRule, proofTypeItem)
	}
	var verifierSetByRule []interface{}
	for _, verifierSetByItem := range verifierSetBy {
		verifierSetByRule = append(verifierSetByRule, verifierSetByItem)
	}

	logs, sub, err := _ZkEvmV2.contract.WatchLogs(opts, "VerifierAddressChanged", verifierAddressRule, proofTypeRule, verifierSetByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkEvmV2VerifierAddressChanged)
				if err := _ZkEvmV2.contract.UnpackLog(event, "VerifierAddressChanged", log); err != nil {
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

// ParseVerifierAddressChanged is a log parse operation binding the contract event 0x4ea861139068e7701a770b8975bb54b6f8f446897fac206dd29424035b4a61eb.
//
// Solidity: event VerifierAddressChanged(address indexed verifierAddress, uint256 indexed proofType, address indexed verifierSetBy)
func (_ZkEvmV2 *ZkEvmV2Filterer) ParseVerifierAddressChanged(log types.Log) (*ZkEvmV2VerifierAddressChanged, error) {
	event := new(ZkEvmV2VerifierAddressChanged)
	if err := _ZkEvmV2.contract.UnpackLog(event, "VerifierAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
