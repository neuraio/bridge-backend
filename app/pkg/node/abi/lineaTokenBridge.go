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

// LineaTokenBridgeMetaData contains all meta data concerning the LineaTokenBridge contract.
var LineaTokenBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AlreadyBridgedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"permitData\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"permitSelector\",\"type\":\"bytes32\"}],\"name\":\"InvalidPermitData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenBridge\",\"type\":\"address\"}],\"name\":\"NotFromRemoteTokenBridge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageService\",\"type\":\"address\"}],\"name\":\"NotMessagingService\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NotReserved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"PermitNotAllowingBridge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PermitNotFromSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"remoteTokenBridge\",\"type\":\"address\"}],\"name\":\"RemoteTokenBridgeAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ReservedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenNotDeployed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ZeroAmountNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nativeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"BridgingFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgingInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nativeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customContract\",\"type\":\"address\"}],\"name\":\"CustomContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"DeploymentConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NewToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgedToken\",\"type\":\"address\"}],\"name\":\"NewTokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"remoteTokenBridge\",\"type\":\"address\"}],\"name\":\"RemoteTokenBridgeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenReserved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"bridgeToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeTokenWithPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bridgedToNativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_tokenMetadata\",\"type\":\"bytes\"}],\"name\":\"completeBridging\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"confirmDeployment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_securityCouncil\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageService\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenBeacon\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_reservedTokens\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageService\",\"outputs\":[{\"internalType\":\"contractIMessageService\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nativeToBridgedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteTokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"removeReserved\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_targetContract\",\"type\":\"address\"}],\"name\":\"setCustomContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nativeTokens\",\"type\":\"address[]\"}],\"name\":\"setDeployed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageService\",\"type\":\"address\"}],\"name\":\"setMessageService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_remoteTokenBridge\",\"type\":\"address\"}],\"name\":\"setRemoteTokenBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"setReserved\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenBeacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LineaTokenBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use LineaTokenBridgeMetaData.ABI instead.
var LineaTokenBridgeABI = LineaTokenBridgeMetaData.ABI

// LineaTokenBridge is an auto generated Go binding around an Ethereum contract.
type LineaTokenBridge struct {
	LineaTokenBridgeCaller     // Read-only binding to the contract
	LineaTokenBridgeTransactor // Write-only binding to the contract
	LineaTokenBridgeFilterer   // Log filterer for contract events
}

// LineaTokenBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type LineaTokenBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LineaTokenBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LineaTokenBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LineaTokenBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LineaTokenBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LineaTokenBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LineaTokenBridgeSession struct {
	Contract     *LineaTokenBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LineaTokenBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LineaTokenBridgeCallerSession struct {
	Contract *LineaTokenBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LineaTokenBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LineaTokenBridgeTransactorSession struct {
	Contract     *LineaTokenBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LineaTokenBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type LineaTokenBridgeRaw struct {
	Contract *LineaTokenBridge // Generic contract binding to access the raw methods on
}

// LineaTokenBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LineaTokenBridgeCallerRaw struct {
	Contract *LineaTokenBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// LineaTokenBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LineaTokenBridgeTransactorRaw struct {
	Contract *LineaTokenBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLineaTokenBridge creates a new instance of LineaTokenBridge, bound to a specific deployed contract.
func NewLineaTokenBridge(address common.Address, backend bind.ContractBackend) (*LineaTokenBridge, error) {
	contract, err := bindLineaTokenBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridge{LineaTokenBridgeCaller: LineaTokenBridgeCaller{contract: contract}, LineaTokenBridgeTransactor: LineaTokenBridgeTransactor{contract: contract}, LineaTokenBridgeFilterer: LineaTokenBridgeFilterer{contract: contract}}, nil
}

// NewLineaTokenBridgeCaller creates a new read-only instance of LineaTokenBridge, bound to a specific deployed contract.
func NewLineaTokenBridgeCaller(address common.Address, caller bind.ContractCaller) (*LineaTokenBridgeCaller, error) {
	contract, err := bindLineaTokenBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeCaller{contract: contract}, nil
}

// NewLineaTokenBridgeTransactor creates a new write-only instance of LineaTokenBridge, bound to a specific deployed contract.
func NewLineaTokenBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*LineaTokenBridgeTransactor, error) {
	contract, err := bindLineaTokenBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeTransactor{contract: contract}, nil
}

// NewLineaTokenBridgeFilterer creates a new log filterer instance of LineaTokenBridge, bound to a specific deployed contract.
func NewLineaTokenBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*LineaTokenBridgeFilterer, error) {
	contract, err := bindLineaTokenBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeFilterer{contract: contract}, nil
}

// bindLineaTokenBridge binds a generic wrapper to an already deployed contract.
func bindLineaTokenBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LineaTokenBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LineaTokenBridge *LineaTokenBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LineaTokenBridge.Contract.LineaTokenBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LineaTokenBridge *LineaTokenBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.LineaTokenBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LineaTokenBridge *LineaTokenBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.LineaTokenBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LineaTokenBridge *LineaTokenBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LineaTokenBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LineaTokenBridge *LineaTokenBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LineaTokenBridge *LineaTokenBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.contract.Transact(opts, method, params...)
}

// BridgedToNativeToken is a free data retrieval call binding the contract method 0xca41a247.
//
// Solidity: function bridgedToNativeToken(address ) view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeCaller) BridgedToNativeToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _LineaTokenBridge.contract.Call(opts, &out, "bridgedToNativeToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgedToNativeToken is a free data retrieval call binding the contract method 0xca41a247.
//
// Solidity: function bridgedToNativeToken(address ) view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeSession) BridgedToNativeToken(arg0 common.Address) (common.Address, error) {
	return _LineaTokenBridge.Contract.BridgedToNativeToken(&_LineaTokenBridge.CallOpts, arg0)
}

// BridgedToNativeToken is a free data retrieval call binding the contract method 0xca41a247.
//
// Solidity: function bridgedToNativeToken(address ) view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeCallerSession) BridgedToNativeToken(arg0 common.Address) (common.Address, error) {
	return _LineaTokenBridge.Contract.BridgedToNativeToken(&_LineaTokenBridge.CallOpts, arg0)
}

// MessageService is a free data retrieval call binding the contract method 0x8dae45dd.
//
// Solidity: function messageService() view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeCaller) MessageService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LineaTokenBridge.contract.Call(opts, &out, "messageService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageService is a free data retrieval call binding the contract method 0x8dae45dd.
//
// Solidity: function messageService() view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeSession) MessageService() (common.Address, error) {
	return _LineaTokenBridge.Contract.MessageService(&_LineaTokenBridge.CallOpts)
}

// MessageService is a free data retrieval call binding the contract method 0x8dae45dd.
//
// Solidity: function messageService() view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeCallerSession) MessageService() (common.Address, error) {
	return _LineaTokenBridge.Contract.MessageService(&_LineaTokenBridge.CallOpts)
}

// NativeToBridgedToken is a free data retrieval call binding the contract method 0xff6555c5.
//
// Solidity: function nativeToBridgedToken(address ) view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeCaller) NativeToBridgedToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _LineaTokenBridge.contract.Call(opts, &out, "nativeToBridgedToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToBridgedToken is a free data retrieval call binding the contract method 0xff6555c5.
//
// Solidity: function nativeToBridgedToken(address ) view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeSession) NativeToBridgedToken(arg0 common.Address) (common.Address, error) {
	return _LineaTokenBridge.Contract.NativeToBridgedToken(&_LineaTokenBridge.CallOpts, arg0)
}

// NativeToBridgedToken is a free data retrieval call binding the contract method 0xff6555c5.
//
// Solidity: function nativeToBridgedToken(address ) view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeCallerSession) NativeToBridgedToken(arg0 common.Address) (common.Address, error) {
	return _LineaTokenBridge.Contract.NativeToBridgedToken(&_LineaTokenBridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LineaTokenBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeSession) Owner() (common.Address, error) {
	return _LineaTokenBridge.Contract.Owner(&_LineaTokenBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeCallerSession) Owner() (common.Address, error) {
	return _LineaTokenBridge.Contract.Owner(&_LineaTokenBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LineaTokenBridge *LineaTokenBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LineaTokenBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LineaTokenBridge *LineaTokenBridgeSession) Paused() (bool, error) {
	return _LineaTokenBridge.Contract.Paused(&_LineaTokenBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LineaTokenBridge *LineaTokenBridgeCallerSession) Paused() (bool, error) {
	return _LineaTokenBridge.Contract.Paused(&_LineaTokenBridge.CallOpts)
}

// RemoteTokenBridge is a free data retrieval call binding the contract method 0xd191bcf9.
//
// Solidity: function remoteTokenBridge() view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeCaller) RemoteTokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LineaTokenBridge.contract.Call(opts, &out, "remoteTokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteTokenBridge is a free data retrieval call binding the contract method 0xd191bcf9.
//
// Solidity: function remoteTokenBridge() view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeSession) RemoteTokenBridge() (common.Address, error) {
	return _LineaTokenBridge.Contract.RemoteTokenBridge(&_LineaTokenBridge.CallOpts)
}

// RemoteTokenBridge is a free data retrieval call binding the contract method 0xd191bcf9.
//
// Solidity: function remoteTokenBridge() view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeCallerSession) RemoteTokenBridge() (common.Address, error) {
	return _LineaTokenBridge.Contract.RemoteTokenBridge(&_LineaTokenBridge.CallOpts)
}

// TokenBeacon is a free data retrieval call binding the contract method 0xccf5a77c.
//
// Solidity: function tokenBeacon() view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeCaller) TokenBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LineaTokenBridge.contract.Call(opts, &out, "tokenBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenBeacon is a free data retrieval call binding the contract method 0xccf5a77c.
//
// Solidity: function tokenBeacon() view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeSession) TokenBeacon() (common.Address, error) {
	return _LineaTokenBridge.Contract.TokenBeacon(&_LineaTokenBridge.CallOpts)
}

// TokenBeacon is a free data retrieval call binding the contract method 0xccf5a77c.
//
// Solidity: function tokenBeacon() view returns(address)
func (_LineaTokenBridge *LineaTokenBridgeCallerSession) TokenBeacon() (common.Address, error) {
	return _LineaTokenBridge.Contract.TokenBeacon(&_LineaTokenBridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x522ea81a.
//
// Solidity: function bridgeToken(address _token, uint256 _amount, address _recipient) payable returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) BridgeToken(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "bridgeToken", _token, _amount, _recipient)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x522ea81a.
//
// Solidity: function bridgeToken(address _token, uint256 _amount, address _recipient) payable returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) BridgeToken(_token common.Address, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.BridgeToken(&_LineaTokenBridge.TransactOpts, _token, _amount, _recipient)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x522ea81a.
//
// Solidity: function bridgeToken(address _token, uint256 _amount, address _recipient) payable returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) BridgeToken(_token common.Address, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.BridgeToken(&_LineaTokenBridge.TransactOpts, _token, _amount, _recipient)
}

// BridgeTokenWithPermit is a paid mutator transaction binding the contract method 0xdfa96efb.
//
// Solidity: function bridgeTokenWithPermit(address _token, uint256 _amount, address _recipient, bytes _permitData) payable returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) BridgeTokenWithPermit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _recipient common.Address, _permitData []byte) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "bridgeTokenWithPermit", _token, _amount, _recipient, _permitData)
}

// BridgeTokenWithPermit is a paid mutator transaction binding the contract method 0xdfa96efb.
//
// Solidity: function bridgeTokenWithPermit(address _token, uint256 _amount, address _recipient, bytes _permitData) payable returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) BridgeTokenWithPermit(_token common.Address, _amount *big.Int, _recipient common.Address, _permitData []byte) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.BridgeTokenWithPermit(&_LineaTokenBridge.TransactOpts, _token, _amount, _recipient, _permitData)
}

// BridgeTokenWithPermit is a paid mutator transaction binding the contract method 0xdfa96efb.
//
// Solidity: function bridgeTokenWithPermit(address _token, uint256 _amount, address _recipient, bytes _permitData) payable returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) BridgeTokenWithPermit(_token common.Address, _amount *big.Int, _recipient common.Address, _permitData []byte) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.BridgeTokenWithPermit(&_LineaTokenBridge.TransactOpts, _token, _amount, _recipient, _permitData)
}

// CompleteBridging is a paid mutator transaction binding the contract method 0xc5e19a46.
//
// Solidity: function completeBridging(address _nativeToken, uint256 _amount, address _recipient, bytes _tokenMetadata) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) CompleteBridging(opts *bind.TransactOpts, _nativeToken common.Address, _amount *big.Int, _recipient common.Address, _tokenMetadata []byte) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "completeBridging", _nativeToken, _amount, _recipient, _tokenMetadata)
}

// CompleteBridging is a paid mutator transaction binding the contract method 0xc5e19a46.
//
// Solidity: function completeBridging(address _nativeToken, uint256 _amount, address _recipient, bytes _tokenMetadata) returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) CompleteBridging(_nativeToken common.Address, _amount *big.Int, _recipient common.Address, _tokenMetadata []byte) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.CompleteBridging(&_LineaTokenBridge.TransactOpts, _nativeToken, _amount, _recipient, _tokenMetadata)
}

// CompleteBridging is a paid mutator transaction binding the contract method 0xc5e19a46.
//
// Solidity: function completeBridging(address _nativeToken, uint256 _amount, address _recipient, bytes _tokenMetadata) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) CompleteBridging(_nativeToken common.Address, _amount *big.Int, _recipient common.Address, _tokenMetadata []byte) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.CompleteBridging(&_LineaTokenBridge.TransactOpts, _nativeToken, _amount, _recipient, _tokenMetadata)
}

// ConfirmDeployment is a paid mutator transaction binding the contract method 0x4bf98dce.
//
// Solidity: function confirmDeployment(address[] _tokens) payable returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) ConfirmDeployment(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "confirmDeployment", _tokens)
}

// ConfirmDeployment is a paid mutator transaction binding the contract method 0x4bf98dce.
//
// Solidity: function confirmDeployment(address[] _tokens) payable returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) ConfirmDeployment(_tokens []common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.ConfirmDeployment(&_LineaTokenBridge.TransactOpts, _tokens)
}

// ConfirmDeployment is a paid mutator transaction binding the contract method 0x4bf98dce.
//
// Solidity: function confirmDeployment(address[] _tokens) payable returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) ConfirmDeployment(_tokens []common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.ConfirmDeployment(&_LineaTokenBridge.TransactOpts, _tokens)
}

// Initialize is a paid mutator transaction binding the contract method 0xe6bfbfd8.
//
// Solidity: function initialize(address _securityCouncil, address _messageService, address _tokenBeacon, address[] _reservedTokens) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) Initialize(opts *bind.TransactOpts, _securityCouncil common.Address, _messageService common.Address, _tokenBeacon common.Address, _reservedTokens []common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "initialize", _securityCouncil, _messageService, _tokenBeacon, _reservedTokens)
}

// Initialize is a paid mutator transaction binding the contract method 0xe6bfbfd8.
//
// Solidity: function initialize(address _securityCouncil, address _messageService, address _tokenBeacon, address[] _reservedTokens) returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) Initialize(_securityCouncil common.Address, _messageService common.Address, _tokenBeacon common.Address, _reservedTokens []common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.Initialize(&_LineaTokenBridge.TransactOpts, _securityCouncil, _messageService, _tokenBeacon, _reservedTokens)
}

// Initialize is a paid mutator transaction binding the contract method 0xe6bfbfd8.
//
// Solidity: function initialize(address _securityCouncil, address _messageService, address _tokenBeacon, address[] _reservedTokens) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) Initialize(_securityCouncil common.Address, _messageService common.Address, _tokenBeacon common.Address, _reservedTokens []common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.Initialize(&_LineaTokenBridge.TransactOpts, _securityCouncil, _messageService, _tokenBeacon, _reservedTokens)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) Pause() (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.Pause(&_LineaTokenBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.Pause(&_LineaTokenBridge.TransactOpts)
}

// RemoveReserved is a paid mutator transaction binding the contract method 0xedc42a22.
//
// Solidity: function removeReserved(address _token) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) RemoveReserved(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "removeReserved", _token)
}

// RemoveReserved is a paid mutator transaction binding the contract method 0xedc42a22.
//
// Solidity: function removeReserved(address _token) returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) RemoveReserved(_token common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.RemoveReserved(&_LineaTokenBridge.TransactOpts, _token)
}

// RemoveReserved is a paid mutator transaction binding the contract method 0xedc42a22.
//
// Solidity: function removeReserved(address _token) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) RemoveReserved(_token common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.RemoveReserved(&_LineaTokenBridge.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.RenounceOwnership(&_LineaTokenBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.RenounceOwnership(&_LineaTokenBridge.TransactOpts)
}

// SetCustomContract is a paid mutator transaction binding the contract method 0x1754f301.
//
// Solidity: function setCustomContract(address _nativeToken, address _targetContract) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) SetCustomContract(opts *bind.TransactOpts, _nativeToken common.Address, _targetContract common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "setCustomContract", _nativeToken, _targetContract)
}

// SetCustomContract is a paid mutator transaction binding the contract method 0x1754f301.
//
// Solidity: function setCustomContract(address _nativeToken, address _targetContract) returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) SetCustomContract(_nativeToken common.Address, _targetContract common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.SetCustomContract(&_LineaTokenBridge.TransactOpts, _nativeToken, _targetContract)
}

// SetCustomContract is a paid mutator transaction binding the contract method 0x1754f301.
//
// Solidity: function setCustomContract(address _nativeToken, address _targetContract) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) SetCustomContract(_nativeToken common.Address, _targetContract common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.SetCustomContract(&_LineaTokenBridge.TransactOpts, _nativeToken, _targetContract)
}

// SetDeployed is a paid mutator transaction binding the contract method 0x2a564f34.
//
// Solidity: function setDeployed(address[] _nativeTokens) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) SetDeployed(opts *bind.TransactOpts, _nativeTokens []common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "setDeployed", _nativeTokens)
}

// SetDeployed is a paid mutator transaction binding the contract method 0x2a564f34.
//
// Solidity: function setDeployed(address[] _nativeTokens) returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) SetDeployed(_nativeTokens []common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.SetDeployed(&_LineaTokenBridge.TransactOpts, _nativeTokens)
}

// SetDeployed is a paid mutator transaction binding the contract method 0x2a564f34.
//
// Solidity: function setDeployed(address[] _nativeTokens) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) SetDeployed(_nativeTokens []common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.SetDeployed(&_LineaTokenBridge.TransactOpts, _nativeTokens)
}

// SetMessageService is a paid mutator transaction binding the contract method 0xbe46096f.
//
// Solidity: function setMessageService(address _messageService) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) SetMessageService(opts *bind.TransactOpts, _messageService common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "setMessageService", _messageService)
}

// SetMessageService is a paid mutator transaction binding the contract method 0xbe46096f.
//
// Solidity: function setMessageService(address _messageService) returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) SetMessageService(_messageService common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.SetMessageService(&_LineaTokenBridge.TransactOpts, _messageService)
}

// SetMessageService is a paid mutator transaction binding the contract method 0xbe46096f.
//
// Solidity: function setMessageService(address _messageService) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) SetMessageService(_messageService common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.SetMessageService(&_LineaTokenBridge.TransactOpts, _messageService)
}

// SetRemoteTokenBridge is a paid mutator transaction binding the contract method 0xa676e8ab.
//
// Solidity: function setRemoteTokenBridge(address _remoteTokenBridge) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) SetRemoteTokenBridge(opts *bind.TransactOpts, _remoteTokenBridge common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "setRemoteTokenBridge", _remoteTokenBridge)
}

// SetRemoteTokenBridge is a paid mutator transaction binding the contract method 0xa676e8ab.
//
// Solidity: function setRemoteTokenBridge(address _remoteTokenBridge) returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) SetRemoteTokenBridge(_remoteTokenBridge common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.SetRemoteTokenBridge(&_LineaTokenBridge.TransactOpts, _remoteTokenBridge)
}

// SetRemoteTokenBridge is a paid mutator transaction binding the contract method 0xa676e8ab.
//
// Solidity: function setRemoteTokenBridge(address _remoteTokenBridge) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) SetRemoteTokenBridge(_remoteTokenBridge common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.SetRemoteTokenBridge(&_LineaTokenBridge.TransactOpts, _remoteTokenBridge)
}

// SetReserved is a paid mutator transaction binding the contract method 0xcdd914c5.
//
// Solidity: function setReserved(address _token) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) SetReserved(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "setReserved", _token)
}

// SetReserved is a paid mutator transaction binding the contract method 0xcdd914c5.
//
// Solidity: function setReserved(address _token) returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) SetReserved(_token common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.SetReserved(&_LineaTokenBridge.TransactOpts, _token)
}

// SetReserved is a paid mutator transaction binding the contract method 0xcdd914c5.
//
// Solidity: function setReserved(address _token) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) SetReserved(_token common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.SetReserved(&_LineaTokenBridge.TransactOpts, _token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.TransferOwnership(&_LineaTokenBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.TransferOwnership(&_LineaTokenBridge.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LineaTokenBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LineaTokenBridge *LineaTokenBridgeSession) Unpause() (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.Unpause(&_LineaTokenBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LineaTokenBridge *LineaTokenBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _LineaTokenBridge.Contract.Unpause(&_LineaTokenBridge.TransactOpts)
}

// LineaTokenBridgeBridgingFinalizedIterator is returned from FilterBridgingFinalized and is used to iterate over the raw logs and unpacked data for BridgingFinalized events raised by the LineaTokenBridge contract.
type LineaTokenBridgeBridgingFinalizedIterator struct {
	Event *LineaTokenBridgeBridgingFinalized // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgeBridgingFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgeBridgingFinalized)
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
		it.Event = new(LineaTokenBridgeBridgingFinalized)
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
func (it *LineaTokenBridgeBridgingFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgeBridgingFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgeBridgingFinalized represents a BridgingFinalized event raised by the LineaTokenBridge contract.
type LineaTokenBridgeBridgingFinalized struct {
	NativeToken  common.Address
	BridgedToken common.Address
	Amount       *big.Int
	Recipient    common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBridgingFinalized is a free log retrieval operation binding the contract event 0xd28a2d30314c6a2f46b657c15ee4d7ffc33b2817e78f341a260e216cebfbdbef.
//
// Solidity: event BridgingFinalized(address nativeToken, address bridgedToken, uint256 amount, address recipient)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterBridgingFinalized(opts *bind.FilterOpts) (*LineaTokenBridgeBridgingFinalizedIterator, error) {

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "BridgingFinalized")
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeBridgingFinalizedIterator{contract: _LineaTokenBridge.contract, event: "BridgingFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgingFinalized is a free log subscription operation binding the contract event 0xd28a2d30314c6a2f46b657c15ee4d7ffc33b2817e78f341a260e216cebfbdbef.
//
// Solidity: event BridgingFinalized(address nativeToken, address bridgedToken, uint256 amount, address recipient)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchBridgingFinalized(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgeBridgingFinalized) (event.Subscription, error) {

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "BridgingFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgeBridgingFinalized)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "BridgingFinalized", log); err != nil {
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

// ParseBridgingFinalized is a log parse operation binding the contract event 0xd28a2d30314c6a2f46b657c15ee4d7ffc33b2817e78f341a260e216cebfbdbef.
//
// Solidity: event BridgingFinalized(address nativeToken, address bridgedToken, uint256 amount, address recipient)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParseBridgingFinalized(log types.Log) (*LineaTokenBridgeBridgingFinalized, error) {
	event := new(LineaTokenBridgeBridgingFinalized)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "BridgingFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaTokenBridgeBridgingInitiatedIterator is returned from FilterBridgingInitiated and is used to iterate over the raw logs and unpacked data for BridgingInitiated events raised by the LineaTokenBridge contract.
type LineaTokenBridgeBridgingInitiatedIterator struct {
	Event *LineaTokenBridgeBridgingInitiated // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgeBridgingInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgeBridgingInitiated)
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
		it.Event = new(LineaTokenBridgeBridgingInitiated)
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
func (it *LineaTokenBridgeBridgingInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgeBridgingInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgeBridgingInitiated represents a BridgingInitiated event raised by the LineaTokenBridge contract.
type LineaTokenBridgeBridgingInitiated struct {
	Sender    common.Address
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgingInitiated is a free log retrieval operation binding the contract event 0xde5fcf0a1aebed387067eb25655de732ccfc43fe5b5a3d91d367c26e773fcd1c.
//
// Solidity: event BridgingInitiated(address sender, address recipient, address token, uint256 amount)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterBridgingInitiated(opts *bind.FilterOpts) (*LineaTokenBridgeBridgingInitiatedIterator, error) {

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "BridgingInitiated")
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeBridgingInitiatedIterator{contract: _LineaTokenBridge.contract, event: "BridgingInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgingInitiated is a free log subscription operation binding the contract event 0xde5fcf0a1aebed387067eb25655de732ccfc43fe5b5a3d91d367c26e773fcd1c.
//
// Solidity: event BridgingInitiated(address sender, address recipient, address token, uint256 amount)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchBridgingInitiated(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgeBridgingInitiated) (event.Subscription, error) {

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "BridgingInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgeBridgingInitiated)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "BridgingInitiated", log); err != nil {
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

// ParseBridgingInitiated is a log parse operation binding the contract event 0xde5fcf0a1aebed387067eb25655de732ccfc43fe5b5a3d91d367c26e773fcd1c.
//
// Solidity: event BridgingInitiated(address sender, address recipient, address token, uint256 amount)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParseBridgingInitiated(log types.Log) (*LineaTokenBridgeBridgingInitiated, error) {
	event := new(LineaTokenBridgeBridgingInitiated)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "BridgingInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaTokenBridgeCustomContractSetIterator is returned from FilterCustomContractSet and is used to iterate over the raw logs and unpacked data for CustomContractSet events raised by the LineaTokenBridge contract.
type LineaTokenBridgeCustomContractSetIterator struct {
	Event *LineaTokenBridgeCustomContractSet // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgeCustomContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgeCustomContractSet)
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
		it.Event = new(LineaTokenBridgeCustomContractSet)
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
func (it *LineaTokenBridgeCustomContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgeCustomContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgeCustomContractSet represents a CustomContractSet event raised by the LineaTokenBridge contract.
type LineaTokenBridgeCustomContractSet struct {
	NativeToken    common.Address
	CustomContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCustomContractSet is a free log retrieval operation binding the contract event 0xf09b6cb917f99cd6d2187f1c3e172aa8481864ad015f7a4e201458b52ea7d35d.
//
// Solidity: event CustomContractSet(address nativeToken, address customContract)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterCustomContractSet(opts *bind.FilterOpts) (*LineaTokenBridgeCustomContractSetIterator, error) {

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "CustomContractSet")
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeCustomContractSetIterator{contract: _LineaTokenBridge.contract, event: "CustomContractSet", logs: logs, sub: sub}, nil
}

// WatchCustomContractSet is a free log subscription operation binding the contract event 0xf09b6cb917f99cd6d2187f1c3e172aa8481864ad015f7a4e201458b52ea7d35d.
//
// Solidity: event CustomContractSet(address nativeToken, address customContract)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchCustomContractSet(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgeCustomContractSet) (event.Subscription, error) {

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "CustomContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgeCustomContractSet)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "CustomContractSet", log); err != nil {
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

// ParseCustomContractSet is a log parse operation binding the contract event 0xf09b6cb917f99cd6d2187f1c3e172aa8481864ad015f7a4e201458b52ea7d35d.
//
// Solidity: event CustomContractSet(address nativeToken, address customContract)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParseCustomContractSet(log types.Log) (*LineaTokenBridgeCustomContractSet, error) {
	event := new(LineaTokenBridgeCustomContractSet)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "CustomContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaTokenBridgeDeploymentConfirmedIterator is returned from FilterDeploymentConfirmed and is used to iterate over the raw logs and unpacked data for DeploymentConfirmed events raised by the LineaTokenBridge contract.
type LineaTokenBridgeDeploymentConfirmedIterator struct {
	Event *LineaTokenBridgeDeploymentConfirmed // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgeDeploymentConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgeDeploymentConfirmed)
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
		it.Event = new(LineaTokenBridgeDeploymentConfirmed)
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
func (it *LineaTokenBridgeDeploymentConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgeDeploymentConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgeDeploymentConfirmed represents a DeploymentConfirmed event raised by the LineaTokenBridge contract.
type LineaTokenBridgeDeploymentConfirmed struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeploymentConfirmed is a free log retrieval operation binding the contract event 0xd46a93ecc7e5927420f6a728c7a773ebe52497f7003a1a760739879dd32fb242.
//
// Solidity: event DeploymentConfirmed(address[] tokens)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterDeploymentConfirmed(opts *bind.FilterOpts) (*LineaTokenBridgeDeploymentConfirmedIterator, error) {

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "DeploymentConfirmed")
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeDeploymentConfirmedIterator{contract: _LineaTokenBridge.contract, event: "DeploymentConfirmed", logs: logs, sub: sub}, nil
}

// WatchDeploymentConfirmed is a free log subscription operation binding the contract event 0xd46a93ecc7e5927420f6a728c7a773ebe52497f7003a1a760739879dd32fb242.
//
// Solidity: event DeploymentConfirmed(address[] tokens)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchDeploymentConfirmed(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgeDeploymentConfirmed) (event.Subscription, error) {

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "DeploymentConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgeDeploymentConfirmed)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "DeploymentConfirmed", log); err != nil {
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

// ParseDeploymentConfirmed is a log parse operation binding the contract event 0xd46a93ecc7e5927420f6a728c7a773ebe52497f7003a1a760739879dd32fb242.
//
// Solidity: event DeploymentConfirmed(address[] tokens)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParseDeploymentConfirmed(log types.Log) (*LineaTokenBridgeDeploymentConfirmed, error) {
	event := new(LineaTokenBridgeDeploymentConfirmed)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "DeploymentConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaTokenBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LineaTokenBridge contract.
type LineaTokenBridgeInitializedIterator struct {
	Event *LineaTokenBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgeInitialized)
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
		it.Event = new(LineaTokenBridgeInitialized)
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
func (it *LineaTokenBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgeInitialized represents a Initialized event raised by the LineaTokenBridge contract.
type LineaTokenBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*LineaTokenBridgeInitializedIterator, error) {

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeInitializedIterator{contract: _LineaTokenBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgeInitialized)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParseInitialized(log types.Log) (*LineaTokenBridgeInitialized, error) {
	event := new(LineaTokenBridgeInitialized)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaTokenBridgeNewTokenIterator is returned from FilterNewToken and is used to iterate over the raw logs and unpacked data for NewToken events raised by the LineaTokenBridge contract.
type LineaTokenBridgeNewTokenIterator struct {
	Event *LineaTokenBridgeNewToken // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgeNewTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgeNewToken)
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
		it.Event = new(LineaTokenBridgeNewToken)
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
func (it *LineaTokenBridgeNewTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgeNewTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgeNewToken represents a NewToken event raised by the LineaTokenBridge contract.
type LineaTokenBridgeNewToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewToken is a free log retrieval operation binding the contract event 0x0f53e2a811b6fd2d6cd965fd6c27b44fb924ca39f7a7f321115705c22366d623.
//
// Solidity: event NewToken(address token)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterNewToken(opts *bind.FilterOpts) (*LineaTokenBridgeNewTokenIterator, error) {

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "NewToken")
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeNewTokenIterator{contract: _LineaTokenBridge.contract, event: "NewToken", logs: logs, sub: sub}, nil
}

// WatchNewToken is a free log subscription operation binding the contract event 0x0f53e2a811b6fd2d6cd965fd6c27b44fb924ca39f7a7f321115705c22366d623.
//
// Solidity: event NewToken(address token)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchNewToken(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgeNewToken) (event.Subscription, error) {

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "NewToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgeNewToken)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "NewToken", log); err != nil {
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

// ParseNewToken is a log parse operation binding the contract event 0x0f53e2a811b6fd2d6cd965fd6c27b44fb924ca39f7a7f321115705c22366d623.
//
// Solidity: event NewToken(address token)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParseNewToken(log types.Log) (*LineaTokenBridgeNewToken, error) {
	event := new(LineaTokenBridgeNewToken)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "NewToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaTokenBridgeNewTokenDeployedIterator is returned from FilterNewTokenDeployed and is used to iterate over the raw logs and unpacked data for NewTokenDeployed events raised by the LineaTokenBridge contract.
type LineaTokenBridgeNewTokenDeployedIterator struct {
	Event *LineaTokenBridgeNewTokenDeployed // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgeNewTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgeNewTokenDeployed)
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
		it.Event = new(LineaTokenBridgeNewTokenDeployed)
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
func (it *LineaTokenBridgeNewTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgeNewTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgeNewTokenDeployed represents a NewTokenDeployed event raised by the LineaTokenBridge contract.
type LineaTokenBridgeNewTokenDeployed struct {
	BridgedToken common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewTokenDeployed is a free log retrieval operation binding the contract event 0x2f1571e1dab54870fe97532aecbf4758d09fc728734048b6973d64212c0e912d.
//
// Solidity: event NewTokenDeployed(address bridgedToken)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterNewTokenDeployed(opts *bind.FilterOpts) (*LineaTokenBridgeNewTokenDeployedIterator, error) {

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "NewTokenDeployed")
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeNewTokenDeployedIterator{contract: _LineaTokenBridge.contract, event: "NewTokenDeployed", logs: logs, sub: sub}, nil
}

// WatchNewTokenDeployed is a free log subscription operation binding the contract event 0x2f1571e1dab54870fe97532aecbf4758d09fc728734048b6973d64212c0e912d.
//
// Solidity: event NewTokenDeployed(address bridgedToken)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchNewTokenDeployed(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgeNewTokenDeployed) (event.Subscription, error) {

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "NewTokenDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgeNewTokenDeployed)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "NewTokenDeployed", log); err != nil {
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

// ParseNewTokenDeployed is a log parse operation binding the contract event 0x2f1571e1dab54870fe97532aecbf4758d09fc728734048b6973d64212c0e912d.
//
// Solidity: event NewTokenDeployed(address bridgedToken)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParseNewTokenDeployed(log types.Log) (*LineaTokenBridgeNewTokenDeployed, error) {
	event := new(LineaTokenBridgeNewTokenDeployed)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "NewTokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaTokenBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LineaTokenBridge contract.
type LineaTokenBridgeOwnershipTransferredIterator struct {
	Event *LineaTokenBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgeOwnershipTransferred)
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
		it.Event = new(LineaTokenBridgeOwnershipTransferred)
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
func (it *LineaTokenBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the LineaTokenBridge contract.
type LineaTokenBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LineaTokenBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeOwnershipTransferredIterator{contract: _LineaTokenBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgeOwnershipTransferred)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*LineaTokenBridgeOwnershipTransferred, error) {
	event := new(LineaTokenBridgeOwnershipTransferred)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaTokenBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LineaTokenBridge contract.
type LineaTokenBridgePausedIterator struct {
	Event *LineaTokenBridgePaused // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgePaused)
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
		it.Event = new(LineaTokenBridgePaused)
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
func (it *LineaTokenBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgePaused represents a Paused event raised by the LineaTokenBridge contract.
type LineaTokenBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*LineaTokenBridgePausedIterator, error) {

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgePausedIterator{contract: _LineaTokenBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgePaused) (event.Subscription, error) {

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgePaused)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParsePaused(log types.Log) (*LineaTokenBridgePaused, error) {
	event := new(LineaTokenBridgePaused)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaTokenBridgeRemoteTokenBridgeSetIterator is returned from FilterRemoteTokenBridgeSet and is used to iterate over the raw logs and unpacked data for RemoteTokenBridgeSet events raised by the LineaTokenBridge contract.
type LineaTokenBridgeRemoteTokenBridgeSetIterator struct {
	Event *LineaTokenBridgeRemoteTokenBridgeSet // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgeRemoteTokenBridgeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgeRemoteTokenBridgeSet)
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
		it.Event = new(LineaTokenBridgeRemoteTokenBridgeSet)
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
func (it *LineaTokenBridgeRemoteTokenBridgeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgeRemoteTokenBridgeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgeRemoteTokenBridgeSet represents a RemoteTokenBridgeSet event raised by the LineaTokenBridge contract.
type LineaTokenBridgeRemoteTokenBridgeSet struct {
	RemoteTokenBridge common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRemoteTokenBridgeSet is a free log retrieval operation binding the contract event 0x216f068d30c10753af848efd94c2820a8b6db860fd9d1b8625537ed2acc8123f.
//
// Solidity: event RemoteTokenBridgeSet(address remoteTokenBridge)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterRemoteTokenBridgeSet(opts *bind.FilterOpts) (*LineaTokenBridgeRemoteTokenBridgeSetIterator, error) {

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "RemoteTokenBridgeSet")
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeRemoteTokenBridgeSetIterator{contract: _LineaTokenBridge.contract, event: "RemoteTokenBridgeSet", logs: logs, sub: sub}, nil
}

// WatchRemoteTokenBridgeSet is a free log subscription operation binding the contract event 0x216f068d30c10753af848efd94c2820a8b6db860fd9d1b8625537ed2acc8123f.
//
// Solidity: event RemoteTokenBridgeSet(address remoteTokenBridge)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchRemoteTokenBridgeSet(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgeRemoteTokenBridgeSet) (event.Subscription, error) {

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "RemoteTokenBridgeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgeRemoteTokenBridgeSet)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "RemoteTokenBridgeSet", log); err != nil {
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

// ParseRemoteTokenBridgeSet is a log parse operation binding the contract event 0x216f068d30c10753af848efd94c2820a8b6db860fd9d1b8625537ed2acc8123f.
//
// Solidity: event RemoteTokenBridgeSet(address remoteTokenBridge)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParseRemoteTokenBridgeSet(log types.Log) (*LineaTokenBridgeRemoteTokenBridgeSet, error) {
	event := new(LineaTokenBridgeRemoteTokenBridgeSet)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "RemoteTokenBridgeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaTokenBridgeTokenDeployedIterator is returned from FilterTokenDeployed and is used to iterate over the raw logs and unpacked data for TokenDeployed events raised by the LineaTokenBridge contract.
type LineaTokenBridgeTokenDeployedIterator struct {
	Event *LineaTokenBridgeTokenDeployed // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgeTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgeTokenDeployed)
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
		it.Event = new(LineaTokenBridgeTokenDeployed)
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
func (it *LineaTokenBridgeTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgeTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgeTokenDeployed represents a TokenDeployed event raised by the LineaTokenBridge contract.
type LineaTokenBridgeTokenDeployed struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenDeployed is a free log retrieval operation binding the contract event 0x91d24864a084ab70b268a1f865e757ca12006cf298d763b6be697302ef86498c.
//
// Solidity: event TokenDeployed(address token)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterTokenDeployed(opts *bind.FilterOpts) (*LineaTokenBridgeTokenDeployedIterator, error) {

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "TokenDeployed")
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeTokenDeployedIterator{contract: _LineaTokenBridge.contract, event: "TokenDeployed", logs: logs, sub: sub}, nil
}

// WatchTokenDeployed is a free log subscription operation binding the contract event 0x91d24864a084ab70b268a1f865e757ca12006cf298d763b6be697302ef86498c.
//
// Solidity: event TokenDeployed(address token)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchTokenDeployed(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgeTokenDeployed) (event.Subscription, error) {

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "TokenDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgeTokenDeployed)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "TokenDeployed", log); err != nil {
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

// ParseTokenDeployed is a log parse operation binding the contract event 0x91d24864a084ab70b268a1f865e757ca12006cf298d763b6be697302ef86498c.
//
// Solidity: event TokenDeployed(address token)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParseTokenDeployed(log types.Log) (*LineaTokenBridgeTokenDeployed, error) {
	event := new(LineaTokenBridgeTokenDeployed)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "TokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaTokenBridgeTokenReservedIterator is returned from FilterTokenReserved and is used to iterate over the raw logs and unpacked data for TokenReserved events raised by the LineaTokenBridge contract.
type LineaTokenBridgeTokenReservedIterator struct {
	Event *LineaTokenBridgeTokenReserved // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgeTokenReservedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgeTokenReserved)
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
		it.Event = new(LineaTokenBridgeTokenReserved)
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
func (it *LineaTokenBridgeTokenReservedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgeTokenReservedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgeTokenReserved represents a TokenReserved event raised by the LineaTokenBridge contract.
type LineaTokenBridgeTokenReserved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenReserved is a free log retrieval operation binding the contract event 0x5e023c7a09fa0534ce3199f65fc3e635a5e851c5adc88ebda3b9d332ae07cbe9.
//
// Solidity: event TokenReserved(address token)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterTokenReserved(opts *bind.FilterOpts) (*LineaTokenBridgeTokenReservedIterator, error) {

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "TokenReserved")
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeTokenReservedIterator{contract: _LineaTokenBridge.contract, event: "TokenReserved", logs: logs, sub: sub}, nil
}

// WatchTokenReserved is a free log subscription operation binding the contract event 0x5e023c7a09fa0534ce3199f65fc3e635a5e851c5adc88ebda3b9d332ae07cbe9.
//
// Solidity: event TokenReserved(address token)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchTokenReserved(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgeTokenReserved) (event.Subscription, error) {

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "TokenReserved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgeTokenReserved)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "TokenReserved", log); err != nil {
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

// ParseTokenReserved is a log parse operation binding the contract event 0x5e023c7a09fa0534ce3199f65fc3e635a5e851c5adc88ebda3b9d332ae07cbe9.
//
// Solidity: event TokenReserved(address token)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParseTokenReserved(log types.Log) (*LineaTokenBridgeTokenReserved, error) {
	event := new(LineaTokenBridgeTokenReserved)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "TokenReserved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaTokenBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LineaTokenBridge contract.
type LineaTokenBridgeUnpausedIterator struct {
	Event *LineaTokenBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *LineaTokenBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaTokenBridgeUnpaused)
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
		it.Event = new(LineaTokenBridgeUnpaused)
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
func (it *LineaTokenBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaTokenBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaTokenBridgeUnpaused represents a Unpaused event raised by the LineaTokenBridge contract.
type LineaTokenBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LineaTokenBridgeUnpausedIterator, error) {

	logs, sub, err := _LineaTokenBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LineaTokenBridgeUnpausedIterator{contract: _LineaTokenBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LineaTokenBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _LineaTokenBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaTokenBridgeUnpaused)
				if err := _LineaTokenBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LineaTokenBridge *LineaTokenBridgeFilterer) ParseUnpaused(log types.Log) (*LineaTokenBridgeUnpaused, error) {
	event := new(LineaTokenBridgeUnpaused)
	if err := _LineaTokenBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
