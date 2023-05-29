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

// L1ERC20BridgeMetaData contains all meta data concerning the L1ERC20Bridge contract.
var L1ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIMailbox\",\"name\":\"_mailbox\",\"type\":\"address\"},{\"internalType\":\"contractIAllowList\",\"name\":\"_allowList\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedFailedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"__DEPRECATED_lastWithdrawalLimitReset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"__DEPRECATED_withdrawnAmountInWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBlock\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2TxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2TxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBlock\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_l2TokenBeacon\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deployBridgeImplementationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deployBridgeProxyFee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isWithdrawalFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"l2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBeacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenProxyBytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalDepositedAmountPerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// L1ERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L1ERC20BridgeMetaData.ABI instead.
var L1ERC20BridgeABI = L1ERC20BridgeMetaData.ABI

// L1ERC20Bridge is an auto generated Go binding around an Ethereum contract.
type L1ERC20Bridge struct {
	L1ERC20BridgeCaller     // Read-only binding to the contract
	L1ERC20BridgeTransactor // Write-only binding to the contract
	L1ERC20BridgeFilterer   // Log filterer for contract events
}

// L1ERC20BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ERC20BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ERC20BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1ERC20BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1ERC20BridgeSession struct {
	Contract     *L1ERC20Bridge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1ERC20BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1ERC20BridgeCallerSession struct {
	Contract *L1ERC20BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L1ERC20BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1ERC20BridgeTransactorSession struct {
	Contract     *L1ERC20BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L1ERC20BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1ERC20BridgeRaw struct {
	Contract *L1ERC20Bridge // Generic contract binding to access the raw methods on
}

// L1ERC20BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1ERC20BridgeCallerRaw struct {
	Contract *L1ERC20BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L1ERC20BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1ERC20BridgeTransactorRaw struct {
	Contract *L1ERC20BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1ERC20Bridge creates a new instance of L1ERC20Bridge, bound to a specific deployed contract.
func NewL1ERC20Bridge(address common.Address, backend bind.ContractBackend) (*L1ERC20Bridge, error) {
	contract, err := bindL1ERC20Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1ERC20Bridge{L1ERC20BridgeCaller: L1ERC20BridgeCaller{contract: contract}, L1ERC20BridgeTransactor: L1ERC20BridgeTransactor{contract: contract}, L1ERC20BridgeFilterer: L1ERC20BridgeFilterer{contract: contract}}, nil
}

// NewL1ERC20BridgeCaller creates a new read-only instance of L1ERC20Bridge, bound to a specific deployed contract.
func NewL1ERC20BridgeCaller(address common.Address, caller bind.ContractCaller) (*L1ERC20BridgeCaller, error) {
	contract, err := bindL1ERC20Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC20BridgeCaller{contract: contract}, nil
}

// NewL1ERC20BridgeTransactor creates a new write-only instance of L1ERC20Bridge, bound to a specific deployed contract.
func NewL1ERC20BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ERC20BridgeTransactor, error) {
	contract, err := bindL1ERC20Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC20BridgeTransactor{contract: contract}, nil
}

// NewL1ERC20BridgeFilterer creates a new log filterer instance of L1ERC20Bridge, bound to a specific deployed contract.
func NewL1ERC20BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L1ERC20BridgeFilterer, error) {
	contract, err := bindL1ERC20Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1ERC20BridgeFilterer{contract: contract}, nil
}

// bindL1ERC20Bridge binds a generic wrapper to an already deployed contract.
func bindL1ERC20Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1ERC20BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC20Bridge *L1ERC20BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC20Bridge.Contract.L1ERC20BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC20Bridge *L1ERC20BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.L1ERC20BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC20Bridge *L1ERC20BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.L1ERC20BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC20Bridge *L1ERC20BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC20Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC20Bridge *L1ERC20BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC20Bridge *L1ERC20BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.contract.Transact(opts, method, params...)
}

// DEPRECATEDLastWithdrawalLimitReset is a free data retrieval call binding the contract method 0x39e05cb2.
//
// Solidity: function __DEPRECATED_lastWithdrawalLimitReset(address ) view returns(uint256)
func (_L1ERC20Bridge *L1ERC20BridgeCaller) DEPRECATEDLastWithdrawalLimitReset(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1ERC20Bridge.contract.Call(opts, &out, "__DEPRECATED_lastWithdrawalLimitReset", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPRECATEDLastWithdrawalLimitReset is a free data retrieval call binding the contract method 0x39e05cb2.
//
// Solidity: function __DEPRECATED_lastWithdrawalLimitReset(address ) view returns(uint256)
func (_L1ERC20Bridge *L1ERC20BridgeSession) DEPRECATEDLastWithdrawalLimitReset(arg0 common.Address) (*big.Int, error) {
	return _L1ERC20Bridge.Contract.DEPRECATEDLastWithdrawalLimitReset(&_L1ERC20Bridge.CallOpts, arg0)
}

// DEPRECATEDLastWithdrawalLimitReset is a free data retrieval call binding the contract method 0x39e05cb2.
//
// Solidity: function __DEPRECATED_lastWithdrawalLimitReset(address ) view returns(uint256)
func (_L1ERC20Bridge *L1ERC20BridgeCallerSession) DEPRECATEDLastWithdrawalLimitReset(arg0 common.Address) (*big.Int, error) {
	return _L1ERC20Bridge.Contract.DEPRECATEDLastWithdrawalLimitReset(&_L1ERC20Bridge.CallOpts, arg0)
}

// DEPRECATEDWithdrawnAmountInWindow is a free data retrieval call binding the contract method 0x9da08054.
//
// Solidity: function __DEPRECATED_withdrawnAmountInWindow(address ) view returns(uint256)
func (_L1ERC20Bridge *L1ERC20BridgeCaller) DEPRECATEDWithdrawnAmountInWindow(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1ERC20Bridge.contract.Call(opts, &out, "__DEPRECATED_withdrawnAmountInWindow", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPRECATEDWithdrawnAmountInWindow is a free data retrieval call binding the contract method 0x9da08054.
//
// Solidity: function __DEPRECATED_withdrawnAmountInWindow(address ) view returns(uint256)
func (_L1ERC20Bridge *L1ERC20BridgeSession) DEPRECATEDWithdrawnAmountInWindow(arg0 common.Address) (*big.Int, error) {
	return _L1ERC20Bridge.Contract.DEPRECATEDWithdrawnAmountInWindow(&_L1ERC20Bridge.CallOpts, arg0)
}

// DEPRECATEDWithdrawnAmountInWindow is a free data retrieval call binding the contract method 0x9da08054.
//
// Solidity: function __DEPRECATED_withdrawnAmountInWindow(address ) view returns(uint256)
func (_L1ERC20Bridge *L1ERC20BridgeCallerSession) DEPRECATEDWithdrawnAmountInWindow(arg0 common.Address) (*big.Int, error) {
	return _L1ERC20Bridge.Contract.DEPRECATEDWithdrawnAmountInWindow(&_L1ERC20Bridge.CallOpts, arg0)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 , uint256 ) view returns(bool)
func (_L1ERC20Bridge *L1ERC20BridgeCaller) IsWithdrawalFinalized(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _L1ERC20Bridge.contract.Call(opts, &out, "isWithdrawalFinalized", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 , uint256 ) view returns(bool)
func (_L1ERC20Bridge *L1ERC20BridgeSession) IsWithdrawalFinalized(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _L1ERC20Bridge.Contract.IsWithdrawalFinalized(&_L1ERC20Bridge.CallOpts, arg0, arg1)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 , uint256 ) view returns(bool)
func (_L1ERC20Bridge *L1ERC20BridgeCallerSession) IsWithdrawalFinalized(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _L1ERC20Bridge.Contract.IsWithdrawalFinalized(&_L1ERC20Bridge.CallOpts, arg0, arg1)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeCaller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20Bridge.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeSession) L2Bridge() (common.Address, error) {
	return _L1ERC20Bridge.Contract.L2Bridge(&_L1ERC20Bridge.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeCallerSession) L2Bridge() (common.Address, error) {
	return _L1ERC20Bridge.Contract.L2Bridge(&_L1ERC20Bridge.CallOpts)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeCaller) L2TokenAddress(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20Bridge.contract.Call(opts, &out, "l2TokenAddress", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _L1ERC20Bridge.Contract.L2TokenAddress(&_L1ERC20Bridge.CallOpts, _l1Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeCallerSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _L1ERC20Bridge.Contract.L2TokenAddress(&_L1ERC20Bridge.CallOpts, _l1Token)
}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeCaller) L2TokenBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20Bridge.contract.Call(opts, &out, "l2TokenBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeSession) L2TokenBeacon() (common.Address, error) {
	return _L1ERC20Bridge.Contract.L2TokenBeacon(&_L1ERC20Bridge.CallOpts)
}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeCallerSession) L2TokenBeacon() (common.Address, error) {
	return _L1ERC20Bridge.Contract.L2TokenBeacon(&_L1ERC20Bridge.CallOpts)
}

// L2TokenProxyBytecodeHash is a free data retrieval call binding the contract method 0x823f1d96.
//
// Solidity: function l2TokenProxyBytecodeHash() view returns(bytes32)
func (_L1ERC20Bridge *L1ERC20BridgeCaller) L2TokenProxyBytecodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L1ERC20Bridge.contract.Call(opts, &out, "l2TokenProxyBytecodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2TokenProxyBytecodeHash is a free data retrieval call binding the contract method 0x823f1d96.
//
// Solidity: function l2TokenProxyBytecodeHash() view returns(bytes32)
func (_L1ERC20Bridge *L1ERC20BridgeSession) L2TokenProxyBytecodeHash() ([32]byte, error) {
	return _L1ERC20Bridge.Contract.L2TokenProxyBytecodeHash(&_L1ERC20Bridge.CallOpts)
}

// L2TokenProxyBytecodeHash is a free data retrieval call binding the contract method 0x823f1d96.
//
// Solidity: function l2TokenProxyBytecodeHash() view returns(bytes32)
func (_L1ERC20Bridge *L1ERC20BridgeCallerSession) L2TokenProxyBytecodeHash() ([32]byte, error) {
	return _L1ERC20Bridge.Contract.L2TokenProxyBytecodeHash(&_L1ERC20Bridge.CallOpts)
}

// TotalDepositedAmountPerUser is a free data retrieval call binding the contract method 0xa52afbe3.
//
// Solidity: function totalDepositedAmountPerUser(address , address ) view returns(uint256)
func (_L1ERC20Bridge *L1ERC20BridgeCaller) TotalDepositedAmountPerUser(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1ERC20Bridge.contract.Call(opts, &out, "totalDepositedAmountPerUser", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDepositedAmountPerUser is a free data retrieval call binding the contract method 0xa52afbe3.
//
// Solidity: function totalDepositedAmountPerUser(address , address ) view returns(uint256)
func (_L1ERC20Bridge *L1ERC20BridgeSession) TotalDepositedAmountPerUser(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _L1ERC20Bridge.Contract.TotalDepositedAmountPerUser(&_L1ERC20Bridge.CallOpts, arg0, arg1)
}

// TotalDepositedAmountPerUser is a free data retrieval call binding the contract method 0xa52afbe3.
//
// Solidity: function totalDepositedAmountPerUser(address , address ) view returns(uint256)
func (_L1ERC20Bridge *L1ERC20BridgeCallerSession) TotalDepositedAmountPerUser(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _L1ERC20Bridge.Contract.TotalDepositedAmountPerUser(&_L1ERC20Bridge.CallOpts, arg0, arg1)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes32[] _merkleProof) returns()
func (_L1ERC20Bridge *L1ERC20BridgeTransactor) ClaimFailedDeposit(opts *bind.TransactOpts, _depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.contract.Transact(opts, "claimFailedDeposit", _depositSender, _l1Token, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes32[] _merkleProof) returns()
func (_L1ERC20Bridge *L1ERC20BridgeSession) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.ClaimFailedDeposit(&_L1ERC20Bridge.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes32[] _merkleProof) returns()
func (_L1ERC20Bridge *L1ERC20BridgeTransactorSession) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.ClaimFailedDeposit(&_L1ERC20Bridge.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _merkleProof)
}

// Deposit is a paid mutator transaction binding the contract method 0x933999fb.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte) payable returns(bytes32 l2TxHash)
func (_L1ERC20Bridge *L1ERC20BridgeTransactor) Deposit(opts *bind.TransactOpts, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int) (*types.Transaction, error) {
	return _L1ERC20Bridge.contract.Transact(opts, "deposit", _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte)
}

// Deposit is a paid mutator transaction binding the contract method 0x933999fb.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte) payable returns(bytes32 l2TxHash)
func (_L1ERC20Bridge *L1ERC20BridgeSession) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.Deposit(&_L1ERC20Bridge.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte)
}

// Deposit is a paid mutator transaction binding the contract method 0x933999fb.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte) payable returns(bytes32 l2TxHash)
func (_L1ERC20Bridge *L1ERC20BridgeTransactorSession) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.Deposit(&_L1ERC20Bridge.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 l2TxHash)
func (_L1ERC20Bridge *L1ERC20BridgeTransactor) Deposit0(opts *bind.TransactOpts, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _L1ERC20Bridge.contract.Transact(opts, "deposit0", _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 l2TxHash)
func (_L1ERC20Bridge *L1ERC20BridgeSession) Deposit0(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.Deposit0(&_L1ERC20Bridge.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 l2TxHash)
func (_L1ERC20Bridge *L1ERC20BridgeTransactorSession) Deposit0(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.Deposit0(&_L1ERC20Bridge.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes _message, bytes32[] _merkleProof) returns()
func (_L1ERC20Bridge *L1ERC20BridgeTransactor) FinalizeWithdrawal(opts *bind.TransactOpts, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.contract.Transact(opts, "finalizeWithdrawal", _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes _message, bytes32[] _merkleProof) returns()
func (_L1ERC20Bridge *L1ERC20BridgeSession) FinalizeWithdrawal(_l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.FinalizeWithdrawal(&_L1ERC20Bridge.TransactOpts, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes _message, bytes32[] _merkleProof) returns()
func (_L1ERC20Bridge *L1ERC20BridgeTransactorSession) FinalizeWithdrawal(_l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.FinalizeWithdrawal(&_L1ERC20Bridge.TransactOpts, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _message, _merkleProof)
}

// Initialize is a paid mutator transaction binding the contract method 0xa0473785.
//
// Solidity: function initialize(bytes[] _factoryDeps, address _l2TokenBeacon, address _governor, uint256 _deployBridgeImplementationFee, uint256 _deployBridgeProxyFee) payable returns()
func (_L1ERC20Bridge *L1ERC20BridgeTransactor) Initialize(opts *bind.TransactOpts, _factoryDeps [][]byte, _l2TokenBeacon common.Address, _governor common.Address, _deployBridgeImplementationFee *big.Int, _deployBridgeProxyFee *big.Int) (*types.Transaction, error) {
	return _L1ERC20Bridge.contract.Transact(opts, "initialize", _factoryDeps, _l2TokenBeacon, _governor, _deployBridgeImplementationFee, _deployBridgeProxyFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xa0473785.
//
// Solidity: function initialize(bytes[] _factoryDeps, address _l2TokenBeacon, address _governor, uint256 _deployBridgeImplementationFee, uint256 _deployBridgeProxyFee) payable returns()
func (_L1ERC20Bridge *L1ERC20BridgeSession) Initialize(_factoryDeps [][]byte, _l2TokenBeacon common.Address, _governor common.Address, _deployBridgeImplementationFee *big.Int, _deployBridgeProxyFee *big.Int) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.Initialize(&_L1ERC20Bridge.TransactOpts, _factoryDeps, _l2TokenBeacon, _governor, _deployBridgeImplementationFee, _deployBridgeProxyFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xa0473785.
//
// Solidity: function initialize(bytes[] _factoryDeps, address _l2TokenBeacon, address _governor, uint256 _deployBridgeImplementationFee, uint256 _deployBridgeProxyFee) payable returns()
func (_L1ERC20Bridge *L1ERC20BridgeTransactorSession) Initialize(_factoryDeps [][]byte, _l2TokenBeacon common.Address, _governor common.Address, _deployBridgeImplementationFee *big.Int, _deployBridgeProxyFee *big.Int) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.Initialize(&_L1ERC20Bridge.TransactOpts, _factoryDeps, _l2TokenBeacon, _governor, _deployBridgeImplementationFee, _deployBridgeProxyFee)
}

// L1ERC20BridgeClaimedFailedDepositIterator is returned from FilterClaimedFailedDeposit and is used to iterate over the raw logs and unpacked data for ClaimedFailedDeposit events raised by the L1ERC20Bridge contract.
type L1ERC20BridgeClaimedFailedDepositIterator struct {
	Event *L1ERC20BridgeClaimedFailedDeposit // Event containing the contract specifics and raw log

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
func (it *L1ERC20BridgeClaimedFailedDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC20BridgeClaimedFailedDeposit)
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
		it.Event = new(L1ERC20BridgeClaimedFailedDeposit)
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
func (it *L1ERC20BridgeClaimedFailedDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC20BridgeClaimedFailedDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC20BridgeClaimedFailedDeposit represents a ClaimedFailedDeposit event raised by the L1ERC20Bridge contract.
type L1ERC20BridgeClaimedFailedDeposit struct {
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedFailedDeposit is a free log retrieval operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_L1ERC20Bridge *L1ERC20BridgeFilterer) FilterClaimedFailedDeposit(opts *bind.FilterOpts, to []common.Address, l1Token []common.Address) (*L1ERC20BridgeClaimedFailedDepositIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _L1ERC20Bridge.contract.FilterLogs(opts, "ClaimedFailedDeposit", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC20BridgeClaimedFailedDepositIterator{contract: _L1ERC20Bridge.contract, event: "ClaimedFailedDeposit", logs: logs, sub: sub}, nil
}

// WatchClaimedFailedDeposit is a free log subscription operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_L1ERC20Bridge *L1ERC20BridgeFilterer) WatchClaimedFailedDeposit(opts *bind.WatchOpts, sink chan<- *L1ERC20BridgeClaimedFailedDeposit, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _L1ERC20Bridge.contract.WatchLogs(opts, "ClaimedFailedDeposit", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC20BridgeClaimedFailedDeposit)
				if err := _L1ERC20Bridge.contract.UnpackLog(event, "ClaimedFailedDeposit", log); err != nil {
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

// ParseClaimedFailedDeposit is a log parse operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_L1ERC20Bridge *L1ERC20BridgeFilterer) ParseClaimedFailedDeposit(log types.Log) (*L1ERC20BridgeClaimedFailedDeposit, error) {
	event := new(L1ERC20BridgeClaimedFailedDeposit)
	if err := _L1ERC20Bridge.contract.UnpackLog(event, "ClaimedFailedDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC20BridgeDepositInitiatedIterator is returned from FilterDepositInitiated and is used to iterate over the raw logs and unpacked data for DepositInitiated events raised by the L1ERC20Bridge contract.
type L1ERC20BridgeDepositInitiatedIterator struct {
	Event *L1ERC20BridgeDepositInitiated // Event containing the contract specifics and raw log

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
func (it *L1ERC20BridgeDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC20BridgeDepositInitiated)
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
		it.Event = new(L1ERC20BridgeDepositInitiated)
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
func (it *L1ERC20BridgeDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC20BridgeDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC20BridgeDepositInitiated represents a DepositInitiated event raised by the L1ERC20Bridge contract.
type L1ERC20BridgeDepositInitiated struct {
	L2DepositTxHash [32]byte
	From            common.Address
	To              common.Address
	L1Token         common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositInitiated is a free log retrieval operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_L1ERC20Bridge *L1ERC20BridgeFilterer) FilterDepositInitiated(opts *bind.FilterOpts, l2DepositTxHash [][32]byte, from []common.Address, to []common.Address) (*L1ERC20BridgeDepositInitiatedIterator, error) {

	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1ERC20Bridge.contract.FilterLogs(opts, "DepositInitiated", l2DepositTxHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC20BridgeDepositInitiatedIterator{contract: _L1ERC20Bridge.contract, event: "DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchDepositInitiated is a free log subscription operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_L1ERC20Bridge *L1ERC20BridgeFilterer) WatchDepositInitiated(opts *bind.WatchOpts, sink chan<- *L1ERC20BridgeDepositInitiated, l2DepositTxHash [][32]byte, from []common.Address, to []common.Address) (event.Subscription, error) {

	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1ERC20Bridge.contract.WatchLogs(opts, "DepositInitiated", l2DepositTxHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC20BridgeDepositInitiated)
				if err := _L1ERC20Bridge.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
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

// ParseDepositInitiated is a log parse operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_L1ERC20Bridge *L1ERC20BridgeFilterer) ParseDepositInitiated(log types.Log) (*L1ERC20BridgeDepositInitiated, error) {
	event := new(L1ERC20BridgeDepositInitiated)
	if err := _L1ERC20Bridge.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC20BridgeWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the L1ERC20Bridge contract.
type L1ERC20BridgeWithdrawalFinalizedIterator struct {
	Event *L1ERC20BridgeWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *L1ERC20BridgeWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC20BridgeWithdrawalFinalized)
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
		it.Event = new(L1ERC20BridgeWithdrawalFinalized)
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
func (it *L1ERC20BridgeWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC20BridgeWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC20BridgeWithdrawalFinalized represents a WithdrawalFinalized event raised by the L1ERC20Bridge contract.
type L1ERC20BridgeWithdrawalFinalized struct {
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_L1ERC20Bridge *L1ERC20BridgeFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, to []common.Address, l1Token []common.Address) (*L1ERC20BridgeWithdrawalFinalizedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _L1ERC20Bridge.contract.FilterLogs(opts, "WithdrawalFinalized", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC20BridgeWithdrawalFinalizedIterator{contract: _L1ERC20Bridge.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_L1ERC20Bridge *L1ERC20BridgeFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *L1ERC20BridgeWithdrawalFinalized, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _L1ERC20Bridge.contract.WatchLogs(opts, "WithdrawalFinalized", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC20BridgeWithdrawalFinalized)
				if err := _L1ERC20Bridge.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_L1ERC20Bridge *L1ERC20BridgeFilterer) ParseWithdrawalFinalized(log types.Log) (*L1ERC20BridgeWithdrawalFinalized, error) {
	event := new(L1ERC20BridgeWithdrawalFinalized)
	if err := _L1ERC20Bridge.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
