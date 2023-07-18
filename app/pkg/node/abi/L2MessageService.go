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

// L2MessageServiceMetaData contains all meta data concerning the L2MessageService contract.
var L2MessageServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptyMessageHashesArray\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"FeePaymentFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"IsNotPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"IsPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageDoesNotExistOrHasAlreadyBeenClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"MessageHashesListLengthHigherThanOneHundred\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"MessageSendingFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueSentTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueShouldBeGreaterThanFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"resettingAddress\",\"type\":\"address\"}],\"name\":\"AmountUsedInPeriodReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"messageHashes\",\"type\":\"bytes32[]\"}],\"name\":\"L1L2MessageHashesAddedToInbox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"amountChangeBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"amountUsedLoweredToLimit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"usedAmountResetToZero\",\"type\":\"bool\"}],\"name\":\"LimitAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"UnPaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENERAL_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INBOX_STATUS_CLAIMED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INBOX_STATUS_RECEIVED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INBOX_STATUS_UNKNOWN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_L2_MESSAGE_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_L2_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_L1_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_FEE_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVING_SYSTEM_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RATE_LIMIT_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_messageHashes\",\"type\":\"bytes32[]\"}],\"name\":\"addL1L2MessageHashes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodAmountInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"inboxL1L2MessageStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_securityCouncil\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1l2MessageSetter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rateLimitPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rateLimitAmount\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumFeeInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextMessageNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pauseType\",\"type\":\"bytes32\"}],\"name\":\"pauseByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pauseTypeStatuses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetAmountUsedInPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"resetRateLimitAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeInWei\",\"type\":\"uint256\"}],\"name\":\"setMinimumFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pauseType\",\"type\":\"bytes32\"}],\"name\":\"unPauseByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// L2MessageServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use L2MessageServiceMetaData.ABI instead.
var L2MessageServiceABI = L2MessageServiceMetaData.ABI

// L2MessageService is an auto generated Go binding around an Ethereum contract.
type L2MessageService struct {
	L2MessageServiceCaller     // Read-only binding to the contract
	L2MessageServiceTransactor // Write-only binding to the contract
	L2MessageServiceFilterer   // Log filterer for contract events
}

// L2MessageServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2MessageServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2MessageServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2MessageServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2MessageServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2MessageServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2MessageServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2MessageServiceSession struct {
	Contract     *L2MessageService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2MessageServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2MessageServiceCallerSession struct {
	Contract *L2MessageServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// L2MessageServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2MessageServiceTransactorSession struct {
	Contract     *L2MessageServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L2MessageServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2MessageServiceRaw struct {
	Contract *L2MessageService // Generic contract binding to access the raw methods on
}

// L2MessageServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2MessageServiceCallerRaw struct {
	Contract *L2MessageServiceCaller // Generic read-only contract binding to access the raw methods on
}

// L2MessageServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2MessageServiceTransactorRaw struct {
	Contract *L2MessageServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2MessageService creates a new instance of L2MessageService, bound to a specific deployed contract.
func NewL2MessageService(address common.Address, backend bind.ContractBackend) (*L2MessageService, error) {
	contract, err := bindL2MessageService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2MessageService{L2MessageServiceCaller: L2MessageServiceCaller{contract: contract}, L2MessageServiceTransactor: L2MessageServiceTransactor{contract: contract}, L2MessageServiceFilterer: L2MessageServiceFilterer{contract: contract}}, nil
}

// NewL2MessageServiceCaller creates a new read-only instance of L2MessageService, bound to a specific deployed contract.
func NewL2MessageServiceCaller(address common.Address, caller bind.ContractCaller) (*L2MessageServiceCaller, error) {
	contract, err := bindL2MessageService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceCaller{contract: contract}, nil
}

// NewL2MessageServiceTransactor creates a new write-only instance of L2MessageService, bound to a specific deployed contract.
func NewL2MessageServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*L2MessageServiceTransactor, error) {
	contract, err := bindL2MessageService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceTransactor{contract: contract}, nil
}

// NewL2MessageServiceFilterer creates a new log filterer instance of L2MessageService, bound to a specific deployed contract.
func NewL2MessageServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*L2MessageServiceFilterer, error) {
	contract, err := bindL2MessageService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceFilterer{contract: contract}, nil
}

// bindL2MessageService binds a generic wrapper to an already deployed contract.
func bindL2MessageService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2MessageServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2MessageService *L2MessageServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2MessageService.Contract.L2MessageServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2MessageService *L2MessageServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2MessageService.Contract.L2MessageServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2MessageService *L2MessageServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2MessageService.Contract.L2MessageServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2MessageService *L2MessageServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2MessageService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2MessageService *L2MessageServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2MessageService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2MessageService *L2MessageServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2MessageService.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L2MessageService.Contract.DEFAULTADMINROLE(&_L2MessageService.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L2MessageService.Contract.DEFAULTADMINROLE(&_L2MessageService.CallOpts)
}

// GENERALPAUSETYPE is a free data retrieval call binding the contract method 0x6a637967.
//
// Solidity: function GENERAL_PAUSE_TYPE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCaller) GENERALPAUSETYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "GENERAL_PAUSE_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GENERALPAUSETYPE is a free data retrieval call binding the contract method 0x6a637967.
//
// Solidity: function GENERAL_PAUSE_TYPE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceSession) GENERALPAUSETYPE() ([32]byte, error) {
	return _L2MessageService.Contract.GENERALPAUSETYPE(&_L2MessageService.CallOpts)
}

// GENERALPAUSETYPE is a free data retrieval call binding the contract method 0x6a637967.
//
// Solidity: function GENERAL_PAUSE_TYPE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCallerSession) GENERALPAUSETYPE() ([32]byte, error) {
	return _L2MessageService.Contract.GENERALPAUSETYPE(&_L2MessageService.CallOpts)
}

// INBOXSTATUSCLAIMED is a free data retrieval call binding the contract method 0x91f7b901.
//
// Solidity: function INBOX_STATUS_CLAIMED() view returns(uint8)
func (_L2MessageService *L2MessageServiceCaller) INBOXSTATUSCLAIMED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "INBOX_STATUS_CLAIMED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INBOXSTATUSCLAIMED is a free data retrieval call binding the contract method 0x91f7b901.
//
// Solidity: function INBOX_STATUS_CLAIMED() view returns(uint8)
func (_L2MessageService *L2MessageServiceSession) INBOXSTATUSCLAIMED() (uint8, error) {
	return _L2MessageService.Contract.INBOXSTATUSCLAIMED(&_L2MessageService.CallOpts)
}

// INBOXSTATUSCLAIMED is a free data retrieval call binding the contract method 0x91f7b901.
//
// Solidity: function INBOX_STATUS_CLAIMED() view returns(uint8)
func (_L2MessageService *L2MessageServiceCallerSession) INBOXSTATUSCLAIMED() (uint8, error) {
	return _L2MessageService.Contract.INBOXSTATUSCLAIMED(&_L2MessageService.CallOpts)
}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_L2MessageService *L2MessageServiceCaller) INBOXSTATUSRECEIVED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "INBOX_STATUS_RECEIVED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_L2MessageService *L2MessageServiceSession) INBOXSTATUSRECEIVED() (uint8, error) {
	return _L2MessageService.Contract.INBOXSTATUSRECEIVED(&_L2MessageService.CallOpts)
}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_L2MessageService *L2MessageServiceCallerSession) INBOXSTATUSRECEIVED() (uint8, error) {
	return _L2MessageService.Contract.INBOXSTATUSRECEIVED(&_L2MessageService.CallOpts)
}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_L2MessageService *L2MessageServiceCaller) INBOXSTATUSUNKNOWN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "INBOX_STATUS_UNKNOWN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_L2MessageService *L2MessageServiceSession) INBOXSTATUSUNKNOWN() (uint8, error) {
	return _L2MessageService.Contract.INBOXSTATUSUNKNOWN(&_L2MessageService.CallOpts)
}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_L2MessageService *L2MessageServiceCallerSession) INBOXSTATUSUNKNOWN() (uint8, error) {
	return _L2MessageService.Contract.INBOXSTATUSUNKNOWN(&_L2MessageService.CallOpts)
}

// L1L2MESSAGESETTERROLE is a free data retrieval call binding the contract method 0x74377a34.
//
// Solidity: function L1_L2_MESSAGE_SETTER_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCaller) L1L2MESSAGESETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "L1_L2_MESSAGE_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1L2MESSAGESETTERROLE is a free data retrieval call binding the contract method 0x74377a34.
//
// Solidity: function L1_L2_MESSAGE_SETTER_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceSession) L1L2MESSAGESETTERROLE() ([32]byte, error) {
	return _L2MessageService.Contract.L1L2MESSAGESETTERROLE(&_L2MessageService.CallOpts)
}

// L1L2MESSAGESETTERROLE is a free data retrieval call binding the contract method 0x74377a34.
//
// Solidity: function L1_L2_MESSAGE_SETTER_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCallerSession) L1L2MESSAGESETTERROLE() ([32]byte, error) {
	return _L2MessageService.Contract.L1L2MESSAGESETTERROLE(&_L2MessageService.CallOpts)
}

// L1L2PAUSETYPE is a free data retrieval call binding the contract method 0x11314d0f.
//
// Solidity: function L1_L2_PAUSE_TYPE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCaller) L1L2PAUSETYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "L1_L2_PAUSE_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1L2PAUSETYPE is a free data retrieval call binding the contract method 0x11314d0f.
//
// Solidity: function L1_L2_PAUSE_TYPE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceSession) L1L2PAUSETYPE() ([32]byte, error) {
	return _L2MessageService.Contract.L1L2PAUSETYPE(&_L2MessageService.CallOpts)
}

// L1L2PAUSETYPE is a free data retrieval call binding the contract method 0x11314d0f.
//
// Solidity: function L1_L2_PAUSE_TYPE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCallerSession) L1L2PAUSETYPE() ([32]byte, error) {
	return _L2MessageService.Contract.L1L2PAUSETYPE(&_L2MessageService.CallOpts)
}

// L2L1PAUSETYPE is a free data retrieval call binding the contract method 0xabd6230d.
//
// Solidity: function L2_L1_PAUSE_TYPE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCaller) L2L1PAUSETYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "L2_L1_PAUSE_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2L1PAUSETYPE is a free data retrieval call binding the contract method 0xabd6230d.
//
// Solidity: function L2_L1_PAUSE_TYPE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceSession) L2L1PAUSETYPE() ([32]byte, error) {
	return _L2MessageService.Contract.L2L1PAUSETYPE(&_L2MessageService.CallOpts)
}

// L2L1PAUSETYPE is a free data retrieval call binding the contract method 0xabd6230d.
//
// Solidity: function L2_L1_PAUSE_TYPE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCallerSession) L2L1PAUSETYPE() ([32]byte, error) {
	return _L2MessageService.Contract.L2L1PAUSETYPE(&_L2MessageService.CallOpts)
}

// MINIMUMFEESETTERROLE is a free data retrieval call binding the contract method 0xbcbd6fcd.
//
// Solidity: function MINIMUM_FEE_SETTER_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCaller) MINIMUMFEESETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "MINIMUM_FEE_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINIMUMFEESETTERROLE is a free data retrieval call binding the contract method 0xbcbd6fcd.
//
// Solidity: function MINIMUM_FEE_SETTER_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceSession) MINIMUMFEESETTERROLE() ([32]byte, error) {
	return _L2MessageService.Contract.MINIMUMFEESETTERROLE(&_L2MessageService.CallOpts)
}

// MINIMUMFEESETTERROLE is a free data retrieval call binding the contract method 0xbcbd6fcd.
//
// Solidity: function MINIMUM_FEE_SETTER_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCallerSession) MINIMUMFEESETTERROLE() ([32]byte, error) {
	return _L2MessageService.Contract.MINIMUMFEESETTERROLE(&_L2MessageService.CallOpts)
}

// PAUSEMANAGERROLE is a free data retrieval call binding the contract method 0xd84f91e8.
//
// Solidity: function PAUSE_MANAGER_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCaller) PAUSEMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "PAUSE_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEMANAGERROLE is a free data retrieval call binding the contract method 0xd84f91e8.
//
// Solidity: function PAUSE_MANAGER_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceSession) PAUSEMANAGERROLE() ([32]byte, error) {
	return _L2MessageService.Contract.PAUSEMANAGERROLE(&_L2MessageService.CallOpts)
}

// PAUSEMANAGERROLE is a free data retrieval call binding the contract method 0xd84f91e8.
//
// Solidity: function PAUSE_MANAGER_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCallerSession) PAUSEMANAGERROLE() ([32]byte, error) {
	return _L2MessageService.Contract.PAUSEMANAGERROLE(&_L2MessageService.CallOpts)
}

// PROVINGSYSTEMPAUSETYPE is a free data retrieval call binding the contract method 0xb4a5a4b7.
//
// Solidity: function PROVING_SYSTEM_PAUSE_TYPE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCaller) PROVINGSYSTEMPAUSETYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "PROVING_SYSTEM_PAUSE_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROVINGSYSTEMPAUSETYPE is a free data retrieval call binding the contract method 0xb4a5a4b7.
//
// Solidity: function PROVING_SYSTEM_PAUSE_TYPE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceSession) PROVINGSYSTEMPAUSETYPE() ([32]byte, error) {
	return _L2MessageService.Contract.PROVINGSYSTEMPAUSETYPE(&_L2MessageService.CallOpts)
}

// PROVINGSYSTEMPAUSETYPE is a free data retrieval call binding the contract method 0xb4a5a4b7.
//
// Solidity: function PROVING_SYSTEM_PAUSE_TYPE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCallerSession) PROVINGSYSTEMPAUSETYPE() ([32]byte, error) {
	return _L2MessageService.Contract.PROVINGSYSTEMPAUSETYPE(&_L2MessageService.CallOpts)
}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCaller) RATELIMITSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "RATE_LIMIT_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceSession) RATELIMITSETTERROLE() ([32]byte, error) {
	return _L2MessageService.Contract.RATELIMITSETTERROLE(&_L2MessageService.CallOpts)
}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_L2MessageService *L2MessageServiceCallerSession) RATELIMITSETTERROLE() ([32]byte, error) {
	return _L2MessageService.Contract.RATELIMITSETTERROLE(&_L2MessageService.CallOpts)
}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_L2MessageService *L2MessageServiceCaller) CurrentPeriodAmountInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "currentPeriodAmountInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_L2MessageService *L2MessageServiceSession) CurrentPeriodAmountInWei() (*big.Int, error) {
	return _L2MessageService.Contract.CurrentPeriodAmountInWei(&_L2MessageService.CallOpts)
}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_L2MessageService *L2MessageServiceCallerSession) CurrentPeriodAmountInWei() (*big.Int, error) {
	return _L2MessageService.Contract.CurrentPeriodAmountInWei(&_L2MessageService.CallOpts)
}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_L2MessageService *L2MessageServiceCaller) CurrentPeriodEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "currentPeriodEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_L2MessageService *L2MessageServiceSession) CurrentPeriodEnd() (*big.Int, error) {
	return _L2MessageService.Contract.CurrentPeriodEnd(&_L2MessageService.CallOpts)
}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_L2MessageService *L2MessageServiceCallerSession) CurrentPeriodEnd() (*big.Int, error) {
	return _L2MessageService.Contract.CurrentPeriodEnd(&_L2MessageService.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2MessageService *L2MessageServiceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2MessageService *L2MessageServiceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L2MessageService.Contract.GetRoleAdmin(&_L2MessageService.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2MessageService *L2MessageServiceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L2MessageService.Contract.GetRoleAdmin(&_L2MessageService.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2MessageService *L2MessageServiceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2MessageService *L2MessageServiceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L2MessageService.Contract.HasRole(&_L2MessageService.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2MessageService *L2MessageServiceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L2MessageService.Contract.HasRole(&_L2MessageService.CallOpts, role, account)
}

// InboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x0f6893ca.
//
// Solidity: function inboxL1L2MessageStatus(bytes32 ) view returns(uint256)
func (_L2MessageService *L2MessageServiceCaller) InboxL1L2MessageStatus(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "inboxL1L2MessageStatus", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x0f6893ca.
//
// Solidity: function inboxL1L2MessageStatus(bytes32 ) view returns(uint256)
func (_L2MessageService *L2MessageServiceSession) InboxL1L2MessageStatus(arg0 [32]byte) (*big.Int, error) {
	return _L2MessageService.Contract.InboxL1L2MessageStatus(&_L2MessageService.CallOpts, arg0)
}

// InboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x0f6893ca.
//
// Solidity: function inboxL1L2MessageStatus(bytes32 ) view returns(uint256)
func (_L2MessageService *L2MessageServiceCallerSession) InboxL1L2MessageStatus(arg0 [32]byte) (*big.Int, error) {
	return _L2MessageService.Contract.InboxL1L2MessageStatus(&_L2MessageService.CallOpts, arg0)
}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_L2MessageService *L2MessageServiceCaller) LimitInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "limitInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_L2MessageService *L2MessageServiceSession) LimitInWei() (*big.Int, error) {
	return _L2MessageService.Contract.LimitInWei(&_L2MessageService.CallOpts)
}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_L2MessageService *L2MessageServiceCallerSession) LimitInWei() (*big.Int, error) {
	return _L2MessageService.Contract.LimitInWei(&_L2MessageService.CallOpts)
}

// MinimumFeeInWei is a free data retrieval call binding the contract method 0x89945883.
//
// Solidity: function minimumFeeInWei() view returns(uint256)
func (_L2MessageService *L2MessageServiceCaller) MinimumFeeInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "minimumFeeInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumFeeInWei is a free data retrieval call binding the contract method 0x89945883.
//
// Solidity: function minimumFeeInWei() view returns(uint256)
func (_L2MessageService *L2MessageServiceSession) MinimumFeeInWei() (*big.Int, error) {
	return _L2MessageService.Contract.MinimumFeeInWei(&_L2MessageService.CallOpts)
}

// MinimumFeeInWei is a free data retrieval call binding the contract method 0x89945883.
//
// Solidity: function minimumFeeInWei() view returns(uint256)
func (_L2MessageService *L2MessageServiceCallerSession) MinimumFeeInWei() (*big.Int, error) {
	return _L2MessageService.Contract.MinimumFeeInWei(&_L2MessageService.CallOpts)
}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_L2MessageService *L2MessageServiceCaller) NextMessageNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "nextMessageNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_L2MessageService *L2MessageServiceSession) NextMessageNumber() (*big.Int, error) {
	return _L2MessageService.Contract.NextMessageNumber(&_L2MessageService.CallOpts)
}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_L2MessageService *L2MessageServiceCallerSession) NextMessageNumber() (*big.Int, error) {
	return _L2MessageService.Contract.NextMessageNumber(&_L2MessageService.CallOpts)
}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 ) view returns(bool)
func (_L2MessageService *L2MessageServiceCaller) PauseTypeStatuses(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "pauseTypeStatuses", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 ) view returns(bool)
func (_L2MessageService *L2MessageServiceSession) PauseTypeStatuses(arg0 [32]byte) (bool, error) {
	return _L2MessageService.Contract.PauseTypeStatuses(&_L2MessageService.CallOpts, arg0)
}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 ) view returns(bool)
func (_L2MessageService *L2MessageServiceCallerSession) PauseTypeStatuses(arg0 [32]byte) (bool, error) {
	return _L2MessageService.Contract.PauseTypeStatuses(&_L2MessageService.CallOpts, arg0)
}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_L2MessageService *L2MessageServiceCaller) PeriodInSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "periodInSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_L2MessageService *L2MessageServiceSession) PeriodInSeconds() (*big.Int, error) {
	return _L2MessageService.Contract.PeriodInSeconds(&_L2MessageService.CallOpts)
}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_L2MessageService *L2MessageServiceCallerSession) PeriodInSeconds() (*big.Int, error) {
	return _L2MessageService.Contract.PeriodInSeconds(&_L2MessageService.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_L2MessageService *L2MessageServiceCaller) Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_L2MessageService *L2MessageServiceSession) Sender() (common.Address, error) {
	return _L2MessageService.Contract.Sender(&_L2MessageService.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_L2MessageService *L2MessageServiceCallerSession) Sender() (common.Address, error) {
	return _L2MessageService.Contract.Sender(&_L2MessageService.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2MessageService *L2MessageServiceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L2MessageService.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2MessageService *L2MessageServiceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2MessageService.Contract.SupportsInterface(&_L2MessageService.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2MessageService *L2MessageServiceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2MessageService.Contract.SupportsInterface(&_L2MessageService.CallOpts, interfaceId)
}

// AddL1L2MessageHashes is a paid mutator transaction binding the contract method 0xf4b476e1.
//
// Solidity: function addL1L2MessageHashes(bytes32[] _messageHashes) returns()
func (_L2MessageService *L2MessageServiceTransactor) AddL1L2MessageHashes(opts *bind.TransactOpts, _messageHashes [][32]byte) (*types.Transaction, error) {
	return _L2MessageService.contract.Transact(opts, "addL1L2MessageHashes", _messageHashes)
}

// AddL1L2MessageHashes is a paid mutator transaction binding the contract method 0xf4b476e1.
//
// Solidity: function addL1L2MessageHashes(bytes32[] _messageHashes) returns()
func (_L2MessageService *L2MessageServiceSession) AddL1L2MessageHashes(_messageHashes [][32]byte) (*types.Transaction, error) {
	return _L2MessageService.Contract.AddL1L2MessageHashes(&_L2MessageService.TransactOpts, _messageHashes)
}

// AddL1L2MessageHashes is a paid mutator transaction binding the contract method 0xf4b476e1.
//
// Solidity: function addL1L2MessageHashes(bytes32[] _messageHashes) returns()
func (_L2MessageService *L2MessageServiceTransactorSession) AddL1L2MessageHashes(_messageHashes [][32]byte) (*types.Transaction, error) {
	return _L2MessageService.Contract.AddL1L2MessageHashes(&_L2MessageService.TransactOpts, _messageHashes)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_L2MessageService *L2MessageServiceTransactor) ClaimMessage(opts *bind.TransactOpts, _from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _L2MessageService.contract.Transact(opts, "claimMessage", _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_L2MessageService *L2MessageServiceSession) ClaimMessage(_from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _L2MessageService.Contract.ClaimMessage(&_L2MessageService.TransactOpts, _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_L2MessageService *L2MessageServiceTransactorSession) ClaimMessage(_from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _L2MessageService.Contract.ClaimMessage(&_L2MessageService.TransactOpts, _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2MessageService *L2MessageServiceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2MessageService.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2MessageService *L2MessageServiceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2MessageService.Contract.GrantRole(&_L2MessageService.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2MessageService *L2MessageServiceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2MessageService.Contract.GrantRole(&_L2MessageService.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _securityCouncil, address _l1l2MessageSetter, uint256 _rateLimitPeriod, uint256 _rateLimitAmount) returns()
func (_L2MessageService *L2MessageServiceTransactor) Initialize(opts *bind.TransactOpts, _securityCouncil common.Address, _l1l2MessageSetter common.Address, _rateLimitPeriod *big.Int, _rateLimitAmount *big.Int) (*types.Transaction, error) {
	return _L2MessageService.contract.Transact(opts, "initialize", _securityCouncil, _l1l2MessageSetter, _rateLimitPeriod, _rateLimitAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _securityCouncil, address _l1l2MessageSetter, uint256 _rateLimitPeriod, uint256 _rateLimitAmount) returns()
func (_L2MessageService *L2MessageServiceSession) Initialize(_securityCouncil common.Address, _l1l2MessageSetter common.Address, _rateLimitPeriod *big.Int, _rateLimitAmount *big.Int) (*types.Transaction, error) {
	return _L2MessageService.Contract.Initialize(&_L2MessageService.TransactOpts, _securityCouncil, _l1l2MessageSetter, _rateLimitPeriod, _rateLimitAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _securityCouncil, address _l1l2MessageSetter, uint256 _rateLimitPeriod, uint256 _rateLimitAmount) returns()
func (_L2MessageService *L2MessageServiceTransactorSession) Initialize(_securityCouncil common.Address, _l1l2MessageSetter common.Address, _rateLimitPeriod *big.Int, _rateLimitAmount *big.Int) (*types.Transaction, error) {
	return _L2MessageService.Contract.Initialize(&_L2MessageService.TransactOpts, _securityCouncil, _l1l2MessageSetter, _rateLimitPeriod, _rateLimitAmount)
}

// PauseByType is a paid mutator transaction binding the contract method 0x8264bd82.
//
// Solidity: function pauseByType(bytes32 _pauseType) returns()
func (_L2MessageService *L2MessageServiceTransactor) PauseByType(opts *bind.TransactOpts, _pauseType [32]byte) (*types.Transaction, error) {
	return _L2MessageService.contract.Transact(opts, "pauseByType", _pauseType)
}

// PauseByType is a paid mutator transaction binding the contract method 0x8264bd82.
//
// Solidity: function pauseByType(bytes32 _pauseType) returns()
func (_L2MessageService *L2MessageServiceSession) PauseByType(_pauseType [32]byte) (*types.Transaction, error) {
	return _L2MessageService.Contract.PauseByType(&_L2MessageService.TransactOpts, _pauseType)
}

// PauseByType is a paid mutator transaction binding the contract method 0x8264bd82.
//
// Solidity: function pauseByType(bytes32 _pauseType) returns()
func (_L2MessageService *L2MessageServiceTransactorSession) PauseByType(_pauseType [32]byte) (*types.Transaction, error) {
	return _L2MessageService.Contract.PauseByType(&_L2MessageService.TransactOpts, _pauseType)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_L2MessageService *L2MessageServiceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2MessageService.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_L2MessageService *L2MessageServiceSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2MessageService.Contract.RenounceRole(&_L2MessageService.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_L2MessageService *L2MessageServiceTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2MessageService.Contract.RenounceRole(&_L2MessageService.TransactOpts, role, account)
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_L2MessageService *L2MessageServiceTransactor) ResetAmountUsedInPeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2MessageService.contract.Transact(opts, "resetAmountUsedInPeriod")
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_L2MessageService *L2MessageServiceSession) ResetAmountUsedInPeriod() (*types.Transaction, error) {
	return _L2MessageService.Contract.ResetAmountUsedInPeriod(&_L2MessageService.TransactOpts)
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_L2MessageService *L2MessageServiceTransactorSession) ResetAmountUsedInPeriod() (*types.Transaction, error) {
	return _L2MessageService.Contract.ResetAmountUsedInPeriod(&_L2MessageService.TransactOpts)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_L2MessageService *L2MessageServiceTransactor) ResetRateLimitAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _L2MessageService.contract.Transact(opts, "resetRateLimitAmount", _amount)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_L2MessageService *L2MessageServiceSession) ResetRateLimitAmount(_amount *big.Int) (*types.Transaction, error) {
	return _L2MessageService.Contract.ResetRateLimitAmount(&_L2MessageService.TransactOpts, _amount)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_L2MessageService *L2MessageServiceTransactorSession) ResetRateLimitAmount(_amount *big.Int) (*types.Transaction, error) {
	return _L2MessageService.Contract.ResetRateLimitAmount(&_L2MessageService.TransactOpts, _amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2MessageService *L2MessageServiceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2MessageService.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2MessageService *L2MessageServiceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2MessageService.Contract.RevokeRole(&_L2MessageService.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2MessageService *L2MessageServiceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2MessageService.Contract.RevokeRole(&_L2MessageService.TransactOpts, role, account)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_L2MessageService *L2MessageServiceTransactor) SendMessage(opts *bind.TransactOpts, _to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _L2MessageService.contract.Transact(opts, "sendMessage", _to, _fee, _calldata)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_L2MessageService *L2MessageServiceSession) SendMessage(_to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _L2MessageService.Contract.SendMessage(&_L2MessageService.TransactOpts, _to, _fee, _calldata)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_L2MessageService *L2MessageServiceTransactorSession) SendMessage(_to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _L2MessageService.Contract.SendMessage(&_L2MessageService.TransactOpts, _to, _fee, _calldata)
}

// SetMinimumFee is a paid mutator transaction binding the contract method 0x182a7506.
//
// Solidity: function setMinimumFee(uint256 _feeInWei) returns()
func (_L2MessageService *L2MessageServiceTransactor) SetMinimumFee(opts *bind.TransactOpts, _feeInWei *big.Int) (*types.Transaction, error) {
	return _L2MessageService.contract.Transact(opts, "setMinimumFee", _feeInWei)
}

// SetMinimumFee is a paid mutator transaction binding the contract method 0x182a7506.
//
// Solidity: function setMinimumFee(uint256 _feeInWei) returns()
func (_L2MessageService *L2MessageServiceSession) SetMinimumFee(_feeInWei *big.Int) (*types.Transaction, error) {
	return _L2MessageService.Contract.SetMinimumFee(&_L2MessageService.TransactOpts, _feeInWei)
}

// SetMinimumFee is a paid mutator transaction binding the contract method 0x182a7506.
//
// Solidity: function setMinimumFee(uint256 _feeInWei) returns()
func (_L2MessageService *L2MessageServiceTransactorSession) SetMinimumFee(_feeInWei *big.Int) (*types.Transaction, error) {
	return _L2MessageService.Contract.SetMinimumFee(&_L2MessageService.TransactOpts, _feeInWei)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0xb45a4f2c.
//
// Solidity: function unPauseByType(bytes32 _pauseType) returns()
func (_L2MessageService *L2MessageServiceTransactor) UnPauseByType(opts *bind.TransactOpts, _pauseType [32]byte) (*types.Transaction, error) {
	return _L2MessageService.contract.Transact(opts, "unPauseByType", _pauseType)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0xb45a4f2c.
//
// Solidity: function unPauseByType(bytes32 _pauseType) returns()
func (_L2MessageService *L2MessageServiceSession) UnPauseByType(_pauseType [32]byte) (*types.Transaction, error) {
	return _L2MessageService.Contract.UnPauseByType(&_L2MessageService.TransactOpts, _pauseType)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0xb45a4f2c.
//
// Solidity: function unPauseByType(bytes32 _pauseType) returns()
func (_L2MessageService *L2MessageServiceTransactorSession) UnPauseByType(_pauseType [32]byte) (*types.Transaction, error) {
	return _L2MessageService.Contract.UnPauseByType(&_L2MessageService.TransactOpts, _pauseType)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2MessageService *L2MessageServiceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2MessageService.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2MessageService *L2MessageServiceSession) Receive() (*types.Transaction, error) {
	return _L2MessageService.Contract.Receive(&_L2MessageService.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2MessageService *L2MessageServiceTransactorSession) Receive() (*types.Transaction, error) {
	return _L2MessageService.Contract.Receive(&_L2MessageService.TransactOpts)
}

// L2MessageServiceAmountUsedInPeriodResetIterator is returned from FilterAmountUsedInPeriodReset and is used to iterate over the raw logs and unpacked data for AmountUsedInPeriodReset events raised by the L2MessageService contract.
type L2MessageServiceAmountUsedInPeriodResetIterator struct {
	Event *L2MessageServiceAmountUsedInPeriodReset // Event containing the contract specifics and raw log

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
func (it *L2MessageServiceAmountUsedInPeriodResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2MessageServiceAmountUsedInPeriodReset)
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
		it.Event = new(L2MessageServiceAmountUsedInPeriodReset)
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
func (it *L2MessageServiceAmountUsedInPeriodResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MessageServiceAmountUsedInPeriodResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2MessageServiceAmountUsedInPeriodReset represents a AmountUsedInPeriodReset event raised by the L2MessageService contract.
type L2MessageServiceAmountUsedInPeriodReset struct {
	ResettingAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAmountUsedInPeriodReset is a free log retrieval operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_L2MessageService *L2MessageServiceFilterer) FilterAmountUsedInPeriodReset(opts *bind.FilterOpts, resettingAddress []common.Address) (*L2MessageServiceAmountUsedInPeriodResetIterator, error) {

	var resettingAddressRule []interface{}
	for _, resettingAddressItem := range resettingAddress {
		resettingAddressRule = append(resettingAddressRule, resettingAddressItem)
	}

	logs, sub, err := _L2MessageService.contract.FilterLogs(opts, "AmountUsedInPeriodReset", resettingAddressRule)
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceAmountUsedInPeriodResetIterator{contract: _L2MessageService.contract, event: "AmountUsedInPeriodReset", logs: logs, sub: sub}, nil
}

// WatchAmountUsedInPeriodReset is a free log subscription operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_L2MessageService *L2MessageServiceFilterer) WatchAmountUsedInPeriodReset(opts *bind.WatchOpts, sink chan<- *L2MessageServiceAmountUsedInPeriodReset, resettingAddress []common.Address) (event.Subscription, error) {

	var resettingAddressRule []interface{}
	for _, resettingAddressItem := range resettingAddress {
		resettingAddressRule = append(resettingAddressRule, resettingAddressItem)
	}

	logs, sub, err := _L2MessageService.contract.WatchLogs(opts, "AmountUsedInPeriodReset", resettingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2MessageServiceAmountUsedInPeriodReset)
				if err := _L2MessageService.contract.UnpackLog(event, "AmountUsedInPeriodReset", log); err != nil {
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
func (_L2MessageService *L2MessageServiceFilterer) ParseAmountUsedInPeriodReset(log types.Log) (*L2MessageServiceAmountUsedInPeriodReset, error) {
	event := new(L2MessageServiceAmountUsedInPeriodReset)
	if err := _L2MessageService.contract.UnpackLog(event, "AmountUsedInPeriodReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2MessageServiceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2MessageService contract.
type L2MessageServiceInitializedIterator struct {
	Event *L2MessageServiceInitialized // Event containing the contract specifics and raw log

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
func (it *L2MessageServiceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2MessageServiceInitialized)
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
		it.Event = new(L2MessageServiceInitialized)
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
func (it *L2MessageServiceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MessageServiceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2MessageServiceInitialized represents a Initialized event raised by the L2MessageService contract.
type L2MessageServiceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2MessageService *L2MessageServiceFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2MessageServiceInitializedIterator, error) {

	logs, sub, err := _L2MessageService.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceInitializedIterator{contract: _L2MessageService.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2MessageService *L2MessageServiceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2MessageServiceInitialized) (event.Subscription, error) {

	logs, sub, err := _L2MessageService.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2MessageServiceInitialized)
				if err := _L2MessageService.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2MessageService *L2MessageServiceFilterer) ParseInitialized(log types.Log) (*L2MessageServiceInitialized, error) {
	event := new(L2MessageServiceInitialized)
	if err := _L2MessageService.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2MessageServiceL1L2MessageHashesAddedToInboxIterator is returned from FilterL1L2MessageHashesAddedToInbox and is used to iterate over the raw logs and unpacked data for L1L2MessageHashesAddedToInbox events raised by the L2MessageService contract.
type L2MessageServiceL1L2MessageHashesAddedToInboxIterator struct {
	Event *L2MessageServiceL1L2MessageHashesAddedToInbox // Event containing the contract specifics and raw log

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
func (it *L2MessageServiceL1L2MessageHashesAddedToInboxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2MessageServiceL1L2MessageHashesAddedToInbox)
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
		it.Event = new(L2MessageServiceL1L2MessageHashesAddedToInbox)
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
func (it *L2MessageServiceL1L2MessageHashesAddedToInboxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MessageServiceL1L2MessageHashesAddedToInboxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2MessageServiceL1L2MessageHashesAddedToInbox represents a L1L2MessageHashesAddedToInbox event raised by the L2MessageService contract.
type L2MessageServiceL1L2MessageHashesAddedToInbox struct {
	MessageHashes [][32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterL1L2MessageHashesAddedToInbox is a free log retrieval operation binding the contract event 0x9995fb3da0c2de4012f2b814b6fc29ce7507571dcb20b8d0bd38621a842df1eb.
//
// Solidity: event L1L2MessageHashesAddedToInbox(bytes32[] messageHashes)
func (_L2MessageService *L2MessageServiceFilterer) FilterL1L2MessageHashesAddedToInbox(opts *bind.FilterOpts) (*L2MessageServiceL1L2MessageHashesAddedToInboxIterator, error) {

	logs, sub, err := _L2MessageService.contract.FilterLogs(opts, "L1L2MessageHashesAddedToInbox")
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceL1L2MessageHashesAddedToInboxIterator{contract: _L2MessageService.contract, event: "L1L2MessageHashesAddedToInbox", logs: logs, sub: sub}, nil
}

// WatchL1L2MessageHashesAddedToInbox is a free log subscription operation binding the contract event 0x9995fb3da0c2de4012f2b814b6fc29ce7507571dcb20b8d0bd38621a842df1eb.
//
// Solidity: event L1L2MessageHashesAddedToInbox(bytes32[] messageHashes)
func (_L2MessageService *L2MessageServiceFilterer) WatchL1L2MessageHashesAddedToInbox(opts *bind.WatchOpts, sink chan<- *L2MessageServiceL1L2MessageHashesAddedToInbox) (event.Subscription, error) {

	logs, sub, err := _L2MessageService.contract.WatchLogs(opts, "L1L2MessageHashesAddedToInbox")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2MessageServiceL1L2MessageHashesAddedToInbox)
				if err := _L2MessageService.contract.UnpackLog(event, "L1L2MessageHashesAddedToInbox", log); err != nil {
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

// ParseL1L2MessageHashesAddedToInbox is a log parse operation binding the contract event 0x9995fb3da0c2de4012f2b814b6fc29ce7507571dcb20b8d0bd38621a842df1eb.
//
// Solidity: event L1L2MessageHashesAddedToInbox(bytes32[] messageHashes)
func (_L2MessageService *L2MessageServiceFilterer) ParseL1L2MessageHashesAddedToInbox(log types.Log) (*L2MessageServiceL1L2MessageHashesAddedToInbox, error) {
	event := new(L2MessageServiceL1L2MessageHashesAddedToInbox)
	if err := _L2MessageService.contract.UnpackLog(event, "L1L2MessageHashesAddedToInbox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2MessageServiceLimitAmountChangedIterator is returned from FilterLimitAmountChanged and is used to iterate over the raw logs and unpacked data for LimitAmountChanged events raised by the L2MessageService contract.
type L2MessageServiceLimitAmountChangedIterator struct {
	Event *L2MessageServiceLimitAmountChanged // Event containing the contract specifics and raw log

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
func (it *L2MessageServiceLimitAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2MessageServiceLimitAmountChanged)
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
		it.Event = new(L2MessageServiceLimitAmountChanged)
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
func (it *L2MessageServiceLimitAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MessageServiceLimitAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2MessageServiceLimitAmountChanged represents a LimitAmountChanged event raised by the L2MessageService contract.
type L2MessageServiceLimitAmountChanged struct {
	AmountChangeBy           common.Address
	Amount                   *big.Int
	AmountUsedLoweredToLimit bool
	UsedAmountResetToZero    bool
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterLimitAmountChanged is a free log retrieval operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_L2MessageService *L2MessageServiceFilterer) FilterLimitAmountChanged(opts *bind.FilterOpts, amountChangeBy []common.Address) (*L2MessageServiceLimitAmountChangedIterator, error) {

	var amountChangeByRule []interface{}
	for _, amountChangeByItem := range amountChangeBy {
		amountChangeByRule = append(amountChangeByRule, amountChangeByItem)
	}

	logs, sub, err := _L2MessageService.contract.FilterLogs(opts, "LimitAmountChanged", amountChangeByRule)
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceLimitAmountChangedIterator{contract: _L2MessageService.contract, event: "LimitAmountChanged", logs: logs, sub: sub}, nil
}

// WatchLimitAmountChanged is a free log subscription operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_L2MessageService *L2MessageServiceFilterer) WatchLimitAmountChanged(opts *bind.WatchOpts, sink chan<- *L2MessageServiceLimitAmountChanged, amountChangeBy []common.Address) (event.Subscription, error) {

	var amountChangeByRule []interface{}
	for _, amountChangeByItem := range amountChangeBy {
		amountChangeByRule = append(amountChangeByRule, amountChangeByItem)
	}

	logs, sub, err := _L2MessageService.contract.WatchLogs(opts, "LimitAmountChanged", amountChangeByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2MessageServiceLimitAmountChanged)
				if err := _L2MessageService.contract.UnpackLog(event, "LimitAmountChanged", log); err != nil {
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
func (_L2MessageService *L2MessageServiceFilterer) ParseLimitAmountChanged(log types.Log) (*L2MessageServiceLimitAmountChanged, error) {
	event := new(L2MessageServiceLimitAmountChanged)
	if err := _L2MessageService.contract.UnpackLog(event, "LimitAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2MessageServiceMessageClaimedIterator is returned from FilterMessageClaimed and is used to iterate over the raw logs and unpacked data for MessageClaimed events raised by the L2MessageService contract.
type L2MessageServiceMessageClaimedIterator struct {
	Event *L2MessageServiceMessageClaimed // Event containing the contract specifics and raw log

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
func (it *L2MessageServiceMessageClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2MessageServiceMessageClaimed)
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
		it.Event = new(L2MessageServiceMessageClaimed)
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
func (it *L2MessageServiceMessageClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MessageServiceMessageClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2MessageServiceMessageClaimed represents a MessageClaimed event raised by the L2MessageService contract.
type L2MessageServiceMessageClaimed struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageClaimed is a free log retrieval operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_L2MessageService *L2MessageServiceFilterer) FilterMessageClaimed(opts *bind.FilterOpts, _messageHash [][32]byte) (*L2MessageServiceMessageClaimedIterator, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _L2MessageService.contract.FilterLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceMessageClaimedIterator{contract: _L2MessageService.contract, event: "MessageClaimed", logs: logs, sub: sub}, nil
}

// WatchMessageClaimed is a free log subscription operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_L2MessageService *L2MessageServiceFilterer) WatchMessageClaimed(opts *bind.WatchOpts, sink chan<- *L2MessageServiceMessageClaimed, _messageHash [][32]byte) (event.Subscription, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _L2MessageService.contract.WatchLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2MessageServiceMessageClaimed)
				if err := _L2MessageService.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
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
func (_L2MessageService *L2MessageServiceFilterer) ParseMessageClaimed(log types.Log) (*L2MessageServiceMessageClaimed, error) {
	event := new(L2MessageServiceMessageClaimed)
	if err := _L2MessageService.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2MessageServiceMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the L2MessageService contract.
type L2MessageServiceMessageSentIterator struct {
	Event *L2MessageServiceMessageSent // Event containing the contract specifics and raw log

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
func (it *L2MessageServiceMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2MessageServiceMessageSent)
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
		it.Event = new(L2MessageServiceMessageSent)
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
func (it *L2MessageServiceMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MessageServiceMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2MessageServiceMessageSent represents a MessageSent event raised by the L2MessageService contract.
type L2MessageServiceMessageSent struct {
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
func (_L2MessageService *L2MessageServiceFilterer) FilterMessageSent(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (*L2MessageServiceMessageSentIterator, error) {

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

	logs, sub, err := _L2MessageService.contract.FilterLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceMessageSentIterator{contract: _L2MessageService.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_L2MessageService *L2MessageServiceFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *L2MessageServiceMessageSent, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _L2MessageService.contract.WatchLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2MessageServiceMessageSent)
				if err := _L2MessageService.contract.UnpackLog(event, "MessageSent", log); err != nil {
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
func (_L2MessageService *L2MessageServiceFilterer) ParseMessageSent(log types.Log) (*L2MessageServiceMessageSent, error) {
	event := new(L2MessageServiceMessageSent)
	if err := _L2MessageService.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2MessageServicePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L2MessageService contract.
type L2MessageServicePausedIterator struct {
	Event *L2MessageServicePaused // Event containing the contract specifics and raw log

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
func (it *L2MessageServicePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2MessageServicePaused)
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
		it.Event = new(L2MessageServicePaused)
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
func (it *L2MessageServicePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MessageServicePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2MessageServicePaused represents a Paused event raised by the L2MessageService contract.
type L2MessageServicePaused struct {
	MessageSender common.Address
	PauseType     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xc343aefb875672fb1857ecda2bdf9fa822ff1e924e3714f6a3d88c5199dee261.
//
// Solidity: event Paused(address messageSender, bytes32 pauseType)
func (_L2MessageService *L2MessageServiceFilterer) FilterPaused(opts *bind.FilterOpts) (*L2MessageServicePausedIterator, error) {

	logs, sub, err := _L2MessageService.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L2MessageServicePausedIterator{contract: _L2MessageService.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xc343aefb875672fb1857ecda2bdf9fa822ff1e924e3714f6a3d88c5199dee261.
//
// Solidity: event Paused(address messageSender, bytes32 pauseType)
func (_L2MessageService *L2MessageServiceFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L2MessageServicePaused) (event.Subscription, error) {

	logs, sub, err := _L2MessageService.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2MessageServicePaused)
				if err := _L2MessageService.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_L2MessageService *L2MessageServiceFilterer) ParsePaused(log types.Log) (*L2MessageServicePaused, error) {
	event := new(L2MessageServicePaused)
	if err := _L2MessageService.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2MessageServiceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the L2MessageService contract.
type L2MessageServiceRoleAdminChangedIterator struct {
	Event *L2MessageServiceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *L2MessageServiceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2MessageServiceRoleAdminChanged)
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
		it.Event = new(L2MessageServiceRoleAdminChanged)
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
func (it *L2MessageServiceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MessageServiceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2MessageServiceRoleAdminChanged represents a RoleAdminChanged event raised by the L2MessageService contract.
type L2MessageServiceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L2MessageService *L2MessageServiceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*L2MessageServiceRoleAdminChangedIterator, error) {

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

	logs, sub, err := _L2MessageService.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceRoleAdminChangedIterator{contract: _L2MessageService.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L2MessageService *L2MessageServiceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *L2MessageServiceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _L2MessageService.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2MessageServiceRoleAdminChanged)
				if err := _L2MessageService.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_L2MessageService *L2MessageServiceFilterer) ParseRoleAdminChanged(log types.Log) (*L2MessageServiceRoleAdminChanged, error) {
	event := new(L2MessageServiceRoleAdminChanged)
	if err := _L2MessageService.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2MessageServiceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the L2MessageService contract.
type L2MessageServiceRoleGrantedIterator struct {
	Event *L2MessageServiceRoleGranted // Event containing the contract specifics and raw log

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
func (it *L2MessageServiceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2MessageServiceRoleGranted)
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
		it.Event = new(L2MessageServiceRoleGranted)
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
func (it *L2MessageServiceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MessageServiceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2MessageServiceRoleGranted represents a RoleGranted event raised by the L2MessageService contract.
type L2MessageServiceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2MessageService *L2MessageServiceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L2MessageServiceRoleGrantedIterator, error) {

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

	logs, sub, err := _L2MessageService.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceRoleGrantedIterator{contract: _L2MessageService.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2MessageService *L2MessageServiceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *L2MessageServiceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2MessageService.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2MessageServiceRoleGranted)
				if err := _L2MessageService.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_L2MessageService *L2MessageServiceFilterer) ParseRoleGranted(log types.Log) (*L2MessageServiceRoleGranted, error) {
	event := new(L2MessageServiceRoleGranted)
	if err := _L2MessageService.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2MessageServiceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the L2MessageService contract.
type L2MessageServiceRoleRevokedIterator struct {
	Event *L2MessageServiceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *L2MessageServiceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2MessageServiceRoleRevoked)
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
		it.Event = new(L2MessageServiceRoleRevoked)
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
func (it *L2MessageServiceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MessageServiceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2MessageServiceRoleRevoked represents a RoleRevoked event raised by the L2MessageService contract.
type L2MessageServiceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2MessageService *L2MessageServiceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L2MessageServiceRoleRevokedIterator, error) {

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

	logs, sub, err := _L2MessageService.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceRoleRevokedIterator{contract: _L2MessageService.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2MessageService *L2MessageServiceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *L2MessageServiceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2MessageService.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2MessageServiceRoleRevoked)
				if err := _L2MessageService.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_L2MessageService *L2MessageServiceFilterer) ParseRoleRevoked(log types.Log) (*L2MessageServiceRoleRevoked, error) {
	event := new(L2MessageServiceRoleRevoked)
	if err := _L2MessageService.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2MessageServiceUnPausedIterator is returned from FilterUnPaused and is used to iterate over the raw logs and unpacked data for UnPaused events raised by the L2MessageService contract.
type L2MessageServiceUnPausedIterator struct {
	Event *L2MessageServiceUnPaused // Event containing the contract specifics and raw log

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
func (it *L2MessageServiceUnPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2MessageServiceUnPaused)
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
		it.Event = new(L2MessageServiceUnPaused)
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
func (it *L2MessageServiceUnPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MessageServiceUnPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2MessageServiceUnPaused represents a UnPaused event raised by the L2MessageService contract.
type L2MessageServiceUnPaused struct {
	MessageSender common.Address
	PauseType     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnPaused is a free log retrieval operation binding the contract event 0xb54c82d9fabaaa460c07181bb36c08c0e72d79293e77a42ac273c81d2a54281b.
//
// Solidity: event UnPaused(address messageSender, bytes32 pauseType)
func (_L2MessageService *L2MessageServiceFilterer) FilterUnPaused(opts *bind.FilterOpts) (*L2MessageServiceUnPausedIterator, error) {

	logs, sub, err := _L2MessageService.contract.FilterLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return &L2MessageServiceUnPausedIterator{contract: _L2MessageService.contract, event: "UnPaused", logs: logs, sub: sub}, nil
}

// WatchUnPaused is a free log subscription operation binding the contract event 0xb54c82d9fabaaa460c07181bb36c08c0e72d79293e77a42ac273c81d2a54281b.
//
// Solidity: event UnPaused(address messageSender, bytes32 pauseType)
func (_L2MessageService *L2MessageServiceFilterer) WatchUnPaused(opts *bind.WatchOpts, sink chan<- *L2MessageServiceUnPaused) (event.Subscription, error) {

	logs, sub, err := _L2MessageService.contract.WatchLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2MessageServiceUnPaused)
				if err := _L2MessageService.contract.UnpackLog(event, "UnPaused", log); err != nil {
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
func (_L2MessageService *L2MessageServiceFilterer) ParseUnPaused(log types.Log) (*L2MessageServiceUnPaused, error) {
	event := new(L2MessageServiceUnPaused)
	if err := _L2MessageService.contract.UnpackLog(event, "UnPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
