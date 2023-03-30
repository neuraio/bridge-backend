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

// BridgeMintReq is an auto generated low-level Go binding around an user-defined struct.
type BridgeMintReq struct {
	Sender     common.Address
	Receiver   common.Address
	Token      common.Address
	Amount     *big.Int
	Fee        *big.Int
	RefChainId *big.Int
	BurnId     [32]byte
}

// Bridgeerc20MetaData contains all meta data concerning the Bridgeerc20 contract.
var Bridgeerc20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"burnId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumPool.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"Revoker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumPool.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MaxBurnUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinBurnUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mintId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"burnId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refChainId\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"auth\",\"type\":\"bool\"}],\"name\":\"SetAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addNativeLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"burnErc20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"burnNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pubkey\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"burnId\",\"type\":\"bytes32\"}],\"internalType\":\"structBridge.MintReq\",\"name\":\"_req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_sign\",\"type\":\"bytes\"}],\"name\":\"mintToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"burnId\",\"type\":\"bytes32\"}],\"internalType\":\"structBridge.MintReq[]\",\"name\":\"_reqs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"}],\"name\":\"mintTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_sign\",\"type\":\"bytes\"}],\"name\":\"revokeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_auth\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMaxBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMinBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_key\",\"type\":\"address\"}],\"name\":\"setPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Bridgeerc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Bridgeerc20MetaData.ABI instead.
var Bridgeerc20ABI = Bridgeerc20MetaData.ABI

// Bridgeerc20 is an auto generated Go binding around an Ethereum contract.
type Bridgeerc20 struct {
	Bridgeerc20Caller     // Read-only binding to the contract
	Bridgeerc20Transactor // Write-only binding to the contract
	Bridgeerc20Filterer   // Log filterer for contract events
}

// Bridgeerc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Bridgeerc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgeerc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Bridgeerc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgeerc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Bridgeerc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgeerc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Bridgeerc20Session struct {
	Contract     *Bridgeerc20      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Bridgeerc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Bridgeerc20CallerSession struct {
	Contract *Bridgeerc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Bridgeerc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Bridgeerc20TransactorSession struct {
	Contract     *Bridgeerc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Bridgeerc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Bridgeerc20Raw struct {
	Contract *Bridgeerc20 // Generic contract binding to access the raw methods on
}

// Bridgeerc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Bridgeerc20CallerRaw struct {
	Contract *Bridgeerc20Caller // Generic read-only contract binding to access the raw methods on
}

// Bridgeerc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Bridgeerc20TransactorRaw struct {
	Contract *Bridgeerc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeerc20 creates a new instance of Bridgeerc20, bound to a specific deployed contract.
func NewBridgeerc20(address common.Address, backend bind.ContractBackend) (*Bridgeerc20, error) {
	contract, err := bindBridgeerc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20{Bridgeerc20Caller: Bridgeerc20Caller{contract: contract}, Bridgeerc20Transactor: Bridgeerc20Transactor{contract: contract}, Bridgeerc20Filterer: Bridgeerc20Filterer{contract: contract}}, nil
}

// NewBridgeerc20Caller creates a new read-only instance of Bridgeerc20, bound to a specific deployed contract.
func NewBridgeerc20Caller(address common.Address, caller bind.ContractCaller) (*Bridgeerc20Caller, error) {
	contract, err := bindBridgeerc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20Caller{contract: contract}, nil
}

// NewBridgeerc20Transactor creates a new write-only instance of Bridgeerc20, bound to a specific deployed contract.
func NewBridgeerc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Bridgeerc20Transactor, error) {
	contract, err := bindBridgeerc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20Transactor{contract: contract}, nil
}

// NewBridgeerc20Filterer creates a new log filterer instance of Bridgeerc20, bound to a specific deployed contract.
func NewBridgeerc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Bridgeerc20Filterer, error) {
	contract, err := bindBridgeerc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20Filterer{contract: contract}, nil
}

// bindBridgeerc20 binds a generic wrapper to an already deployed contract.
func bindBridgeerc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Bridgeerc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgeerc20 *Bridgeerc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgeerc20.Contract.Bridgeerc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgeerc20 *Bridgeerc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.Bridgeerc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgeerc20 *Bridgeerc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.Bridgeerc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgeerc20 *Bridgeerc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgeerc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgeerc20 *Bridgeerc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgeerc20 *Bridgeerc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Bridgeerc20 *Bridgeerc20Caller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridgeerc20.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Bridgeerc20 *Bridgeerc20Session) Admins(arg0 common.Address) (bool, error) {
	return _Bridgeerc20.Contract.Admins(&_Bridgeerc20.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Bridgeerc20 *Bridgeerc20CallerSession) Admins(arg0 common.Address) (bool, error) {
	return _Bridgeerc20.Contract.Admins(&_Bridgeerc20.CallOpts, arg0)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Bridgeerc20 *Bridgeerc20Caller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgeerc20.contract.Call(opts, &out, "feeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Bridgeerc20 *Bridgeerc20Session) FeeReceiver() (common.Address, error) {
	return _Bridgeerc20.Contract.FeeReceiver(&_Bridgeerc20.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Bridgeerc20 *Bridgeerc20CallerSession) FeeReceiver() (common.Address, error) {
	return _Bridgeerc20.Contract.FeeReceiver(&_Bridgeerc20.CallOpts)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x2e334452.
//
// Solidity: function getPublicKey() view returns(address)
func (_Bridgeerc20 *Bridgeerc20Caller) GetPublicKey(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgeerc20.contract.Call(opts, &out, "getPublicKey")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPublicKey is a free data retrieval call binding the contract method 0x2e334452.
//
// Solidity: function getPublicKey() view returns(address)
func (_Bridgeerc20 *Bridgeerc20Session) GetPublicKey() (common.Address, error) {
	return _Bridgeerc20.Contract.GetPublicKey(&_Bridgeerc20.CallOpts)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x2e334452.
//
// Solidity: function getPublicKey() view returns(address)
func (_Bridgeerc20 *Bridgeerc20CallerSession) GetPublicKey() (common.Address, error) {
	return _Bridgeerc20.Contract.GetPublicKey(&_Bridgeerc20.CallOpts)
}

// MaxBurn is a free data retrieval call binding the contract method 0x497bf3b2.
//
// Solidity: function maxBurn(address ) view returns(uint256)
func (_Bridgeerc20 *Bridgeerc20Caller) MaxBurn(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridgeerc20.contract.Call(opts, &out, "maxBurn", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBurn is a free data retrieval call binding the contract method 0x497bf3b2.
//
// Solidity: function maxBurn(address ) view returns(uint256)
func (_Bridgeerc20 *Bridgeerc20Session) MaxBurn(arg0 common.Address) (*big.Int, error) {
	return _Bridgeerc20.Contract.MaxBurn(&_Bridgeerc20.CallOpts, arg0)
}

// MaxBurn is a free data retrieval call binding the contract method 0x497bf3b2.
//
// Solidity: function maxBurn(address ) view returns(uint256)
func (_Bridgeerc20 *Bridgeerc20CallerSession) MaxBurn(arg0 common.Address) (*big.Int, error) {
	return _Bridgeerc20.Contract.MaxBurn(&_Bridgeerc20.CallOpts, arg0)
}

// MinBurn is a free data retrieval call binding the contract method 0x7f856013.
//
// Solidity: function minBurn(address ) view returns(uint256)
func (_Bridgeerc20 *Bridgeerc20Caller) MinBurn(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridgeerc20.contract.Call(opts, &out, "minBurn", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBurn is a free data retrieval call binding the contract method 0x7f856013.
//
// Solidity: function minBurn(address ) view returns(uint256)
func (_Bridgeerc20 *Bridgeerc20Session) MinBurn(arg0 common.Address) (*big.Int, error) {
	return _Bridgeerc20.Contract.MinBurn(&_Bridgeerc20.CallOpts, arg0)
}

// MinBurn is a free data retrieval call binding the contract method 0x7f856013.
//
// Solidity: function minBurn(address ) view returns(uint256)
func (_Bridgeerc20 *Bridgeerc20CallerSession) MinBurn(arg0 common.Address) (*big.Int, error) {
	return _Bridgeerc20.Contract.MinBurn(&_Bridgeerc20.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridgeerc20 *Bridgeerc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgeerc20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridgeerc20 *Bridgeerc20Session) Owner() (common.Address, error) {
	return _Bridgeerc20.Contract.Owner(&_Bridgeerc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridgeerc20 *Bridgeerc20CallerSession) Owner() (common.Address, error) {
	return _Bridgeerc20.Contract.Owner(&_Bridgeerc20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridgeerc20 *Bridgeerc20Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridgeerc20.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridgeerc20 *Bridgeerc20Session) Paused() (bool, error) {
	return _Bridgeerc20.Contract.Paused(&_Bridgeerc20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridgeerc20 *Bridgeerc20CallerSession) Paused() (bool, error) {
	return _Bridgeerc20.Contract.Paused(&_Bridgeerc20.CallOpts)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(bool)
func (_Bridgeerc20 *Bridgeerc20Caller) Records(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridgeerc20.contract.Call(opts, &out, "records", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(bool)
func (_Bridgeerc20 *Bridgeerc20Session) Records(arg0 [32]byte) (bool, error) {
	return _Bridgeerc20.Contract.Records(&_Bridgeerc20.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(bool)
func (_Bridgeerc20 *Bridgeerc20CallerSession) Records(arg0 [32]byte) (bool, error) {
	return _Bridgeerc20.Contract.Records(&_Bridgeerc20.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "addLiquidity", _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_Bridgeerc20 *Bridgeerc20Session) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.AddLiquidity(&_Bridgeerc20.TransactOpts, _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.AddLiquidity(&_Bridgeerc20.TransactOpts, _token, _amount)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0xb238b533.
//
// Solidity: function addNativeLiquidity() payable returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) AddNativeLiquidity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "addNativeLiquidity")
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0xb238b533.
//
// Solidity: function addNativeLiquidity() payable returns()
func (_Bridgeerc20 *Bridgeerc20Session) AddNativeLiquidity() (*types.Transaction, error) {
	return _Bridgeerc20.Contract.AddNativeLiquidity(&_Bridgeerc20.TransactOpts)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0xb238b533.
//
// Solidity: function addNativeLiquidity() payable returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) AddNativeLiquidity() (*types.Transaction, error) {
	return _Bridgeerc20.Contract.AddNativeLiquidity(&_Bridgeerc20.TransactOpts)
}

// BurnErc20 is a paid mutator transaction binding the contract method 0x848ee9fa.
//
// Solidity: function burnErc20(address _receiver, address _token, uint256 _amount, uint256 _dstChainId, uint256 _nonce) returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) BurnErc20(opts *bind.TransactOpts, _receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "burnErc20", _receiver, _token, _amount, _dstChainId, _nonce)
}

// BurnErc20 is a paid mutator transaction binding the contract method 0x848ee9fa.
//
// Solidity: function burnErc20(address _receiver, address _token, uint256 _amount, uint256 _dstChainId, uint256 _nonce) returns()
func (_Bridgeerc20 *Bridgeerc20Session) BurnErc20(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.BurnErc20(&_Bridgeerc20.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce)
}

// BurnErc20 is a paid mutator transaction binding the contract method 0x848ee9fa.
//
// Solidity: function burnErc20(address _receiver, address _token, uint256 _amount, uint256 _dstChainId, uint256 _nonce) returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) BurnErc20(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.BurnErc20(&_Bridgeerc20.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce)
}

// BurnNative is a paid mutator transaction binding the contract method 0xd713bb90.
//
// Solidity: function burnNative(address _receiver, uint256 _dstChainId, uint256 _nonce) payable returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) BurnNative(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "burnNative", _receiver, _dstChainId, _nonce)
}

// BurnNative is a paid mutator transaction binding the contract method 0xd713bb90.
//
// Solidity: function burnNative(address _receiver, uint256 _dstChainId, uint256 _nonce) payable returns()
func (_Bridgeerc20 *Bridgeerc20Session) BurnNative(_receiver common.Address, _dstChainId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.BurnNative(&_Bridgeerc20.TransactOpts, _receiver, _dstChainId, _nonce)
}

// BurnNative is a paid mutator transaction binding the contract method 0xd713bb90.
//
// Solidity: function burnNative(address _receiver, uint256 _dstChainId, uint256 _nonce) payable returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) BurnNative(_receiver common.Address, _dstChainId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.BurnNative(&_Bridgeerc20.TransactOpts, _receiver, _dstChainId, _nonce)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _feeReceiver, address _pubkey) returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) Initialize(opts *bind.TransactOpts, _feeReceiver common.Address, _pubkey common.Address) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "initialize", _feeReceiver, _pubkey)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _feeReceiver, address _pubkey) returns()
func (_Bridgeerc20 *Bridgeerc20Session) Initialize(_feeReceiver common.Address, _pubkey common.Address) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.Initialize(&_Bridgeerc20.TransactOpts, _feeReceiver, _pubkey)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _feeReceiver, address _pubkey) returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) Initialize(_feeReceiver common.Address, _pubkey common.Address) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.Initialize(&_Bridgeerc20.TransactOpts, _feeReceiver, _pubkey)
}

// MintToken is a paid mutator transaction binding the contract method 0x3a857ee6.
//
// Solidity: function mintToken((address,address,address,uint256,uint256,uint256,bytes32) _req, bytes _sign) returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) MintToken(opts *bind.TransactOpts, _req BridgeMintReq, _sign []byte) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "mintToken", _req, _sign)
}

// MintToken is a paid mutator transaction binding the contract method 0x3a857ee6.
//
// Solidity: function mintToken((address,address,address,uint256,uint256,uint256,bytes32) _req, bytes _sign) returns()
func (_Bridgeerc20 *Bridgeerc20Session) MintToken(_req BridgeMintReq, _sign []byte) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.MintToken(&_Bridgeerc20.TransactOpts, _req, _sign)
}

// MintToken is a paid mutator transaction binding the contract method 0x3a857ee6.
//
// Solidity: function mintToken((address,address,address,uint256,uint256,uint256,bytes32) _req, bytes _sign) returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) MintToken(_req BridgeMintReq, _sign []byte) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.MintToken(&_Bridgeerc20.TransactOpts, _req, _sign)
}

// MintTokens is a paid mutator transaction binding the contract method 0x5714020e.
//
// Solidity: function mintTokens((address,address,address,uint256,uint256,uint256,bytes32)[] _reqs, bytes[] _signs, address[] _tokens, uint256[] _fees) returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) MintTokens(opts *bind.TransactOpts, _reqs []BridgeMintReq, _signs [][]byte, _tokens []common.Address, _fees []*big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "mintTokens", _reqs, _signs, _tokens, _fees)
}

// MintTokens is a paid mutator transaction binding the contract method 0x5714020e.
//
// Solidity: function mintTokens((address,address,address,uint256,uint256,uint256,bytes32)[] _reqs, bytes[] _signs, address[] _tokens, uint256[] _fees) returns()
func (_Bridgeerc20 *Bridgeerc20Session) MintTokens(_reqs []BridgeMintReq, _signs [][]byte, _tokens []common.Address, _fees []*big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.MintTokens(&_Bridgeerc20.TransactOpts, _reqs, _signs, _tokens, _fees)
}

// MintTokens is a paid mutator transaction binding the contract method 0x5714020e.
//
// Solidity: function mintTokens((address,address,address,uint256,uint256,uint256,bytes32)[] _reqs, bytes[] _signs, address[] _tokens, uint256[] _fees) returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) MintTokens(_reqs []BridgeMintReq, _signs [][]byte, _tokens []common.Address, _fees []*big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.MintTokens(&_Bridgeerc20.TransactOpts, _reqs, _signs, _tokens, _fees)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridgeerc20 *Bridgeerc20Session) Pause() (*types.Transaction, error) {
	return _Bridgeerc20.Contract.Pause(&_Bridgeerc20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) Pause() (*types.Transaction, error) {
	return _Bridgeerc20.Contract.Pause(&_Bridgeerc20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridgeerc20 *Bridgeerc20Session) RenounceOwnership() (*types.Transaction, error) {
	return _Bridgeerc20.Contract.RenounceOwnership(&_Bridgeerc20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridgeerc20.Contract.RenounceOwnership(&_Bridgeerc20.TransactOpts)
}

// RevokeLiquidity is a paid mutator transaction binding the contract method 0x389c57cf.
//
// Solidity: function revokeLiquidity(address _token, address _to, bytes _sign) returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) RevokeLiquidity(opts *bind.TransactOpts, _token common.Address, _to common.Address, _sign []byte) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "revokeLiquidity", _token, _to, _sign)
}

// RevokeLiquidity is a paid mutator transaction binding the contract method 0x389c57cf.
//
// Solidity: function revokeLiquidity(address _token, address _to, bytes _sign) returns()
func (_Bridgeerc20 *Bridgeerc20Session) RevokeLiquidity(_token common.Address, _to common.Address, _sign []byte) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.RevokeLiquidity(&_Bridgeerc20.TransactOpts, _token, _to, _sign)
}

// RevokeLiquidity is a paid mutator transaction binding the contract method 0x389c57cf.
//
// Solidity: function revokeLiquidity(address _token, address _to, bytes _sign) returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) RevokeLiquidity(_token common.Address, _to common.Address, _sign []byte) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.RevokeLiquidity(&_Bridgeerc20.TransactOpts, _token, _to, _sign)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _user, bool _auth) returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) SetAdmin(opts *bind.TransactOpts, _user common.Address, _auth bool) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "setAdmin", _user, _auth)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _user, bool _auth) returns()
func (_Bridgeerc20 *Bridgeerc20Session) SetAdmin(_user common.Address, _auth bool) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.SetAdmin(&_Bridgeerc20.TransactOpts, _user, _auth)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _user, bool _auth) returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) SetAdmin(_user common.Address, _auth bool) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.SetAdmin(&_Bridgeerc20.TransactOpts, _user, _auth)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) SetFeeReceiver(opts *bind.TransactOpts, _feeReceiver common.Address) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "setFeeReceiver", _feeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_Bridgeerc20 *Bridgeerc20Session) SetFeeReceiver(_feeReceiver common.Address) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.SetFeeReceiver(&_Bridgeerc20.TransactOpts, _feeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) SetFeeReceiver(_feeReceiver common.Address) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.SetFeeReceiver(&_Bridgeerc20.TransactOpts, _feeReceiver)
}

// SetMaxBurn is a paid mutator transaction binding the contract method 0xf9a8ea08.
//
// Solidity: function setMaxBurn(address[] _tokens, uint256[] _amounts) returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) SetMaxBurn(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "setMaxBurn", _tokens, _amounts)
}

// SetMaxBurn is a paid mutator transaction binding the contract method 0xf9a8ea08.
//
// Solidity: function setMaxBurn(address[] _tokens, uint256[] _amounts) returns()
func (_Bridgeerc20 *Bridgeerc20Session) SetMaxBurn(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.SetMaxBurn(&_Bridgeerc20.TransactOpts, _tokens, _amounts)
}

// SetMaxBurn is a paid mutator transaction binding the contract method 0xf9a8ea08.
//
// Solidity: function setMaxBurn(address[] _tokens, uint256[] _amounts) returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) SetMaxBurn(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.SetMaxBurn(&_Bridgeerc20.TransactOpts, _tokens, _amounts)
}

// SetMinBurn is a paid mutator transaction binding the contract method 0xbf4816f0.
//
// Solidity: function setMinBurn(address[] _tokens, uint256[] _amounts) returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) SetMinBurn(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "setMinBurn", _tokens, _amounts)
}

// SetMinBurn is a paid mutator transaction binding the contract method 0xbf4816f0.
//
// Solidity: function setMinBurn(address[] _tokens, uint256[] _amounts) returns()
func (_Bridgeerc20 *Bridgeerc20Session) SetMinBurn(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.SetMinBurn(&_Bridgeerc20.TransactOpts, _tokens, _amounts)
}

// SetMinBurn is a paid mutator transaction binding the contract method 0xbf4816f0.
//
// Solidity: function setMinBurn(address[] _tokens, uint256[] _amounts) returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) SetMinBurn(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.SetMinBurn(&_Bridgeerc20.TransactOpts, _tokens, _amounts)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x097eb759.
//
// Solidity: function setPublicKey(address _key) returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) SetPublicKey(opts *bind.TransactOpts, _key common.Address) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "setPublicKey", _key)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x097eb759.
//
// Solidity: function setPublicKey(address _key) returns()
func (_Bridgeerc20 *Bridgeerc20Session) SetPublicKey(_key common.Address) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.SetPublicKey(&_Bridgeerc20.TransactOpts, _key)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x097eb759.
//
// Solidity: function setPublicKey(address _key) returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) SetPublicKey(_key common.Address) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.SetPublicKey(&_Bridgeerc20.TransactOpts, _key)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridgeerc20 *Bridgeerc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.TransferOwnership(&_Bridgeerc20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridgeerc20.Contract.TransferOwnership(&_Bridgeerc20.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridgeerc20 *Bridgeerc20Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgeerc20.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridgeerc20 *Bridgeerc20Session) Unpause() (*types.Transaction, error) {
	return _Bridgeerc20.Contract.Unpause(&_Bridgeerc20.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridgeerc20 *Bridgeerc20TransactorSession) Unpause() (*types.Transaction, error) {
	return _Bridgeerc20.Contract.Unpause(&_Bridgeerc20.TransactOpts)
}

// Bridgeerc20BurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the Bridgeerc20 contract.
type Bridgeerc20BurnedIterator struct {
	Event *Bridgeerc20Burned // Event containing the contract specifics and raw log

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
func (it *Bridgeerc20BurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgeerc20Burned)
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
		it.Event = new(Bridgeerc20Burned)
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
func (it *Bridgeerc20BurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgeerc20BurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgeerc20Burned represents a Burned event raised by the Bridgeerc20 contract.
type Bridgeerc20Burned struct {
	BurnId     [32]byte
	Sender     common.Address
	Receiver   common.Address
	Token      common.Address
	Amount     *big.Int
	DstChainId *big.Int
	Nonce      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x50e22ad3fc6c213f811692f757b38468af04f08c4377a69004aaf21c7f04485b.
//
// Solidity: event Burned(bytes32 indexed burnId, address sender, address receiver, address token, uint256 amount, uint256 dstChainId, uint256 nonce)
func (_Bridgeerc20 *Bridgeerc20Filterer) FilterBurned(opts *bind.FilterOpts, burnId [][32]byte) (*Bridgeerc20BurnedIterator, error) {

	var burnIdRule []interface{}
	for _, burnIdItem := range burnId {
		burnIdRule = append(burnIdRule, burnIdItem)
	}

	logs, sub, err := _Bridgeerc20.contract.FilterLogs(opts, "Burned", burnIdRule)
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20BurnedIterator{contract: _Bridgeerc20.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x50e22ad3fc6c213f811692f757b38468af04f08c4377a69004aaf21c7f04485b.
//
// Solidity: event Burned(bytes32 indexed burnId, address sender, address receiver, address token, uint256 amount, uint256 dstChainId, uint256 nonce)
func (_Bridgeerc20 *Bridgeerc20Filterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *Bridgeerc20Burned, burnId [][32]byte) (event.Subscription, error) {

	var burnIdRule []interface{}
	for _, burnIdItem := range burnId {
		burnIdRule = append(burnIdRule, burnIdItem)
	}

	logs, sub, err := _Bridgeerc20.contract.WatchLogs(opts, "Burned", burnIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgeerc20Burned)
				if err := _Bridgeerc20.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0x50e22ad3fc6c213f811692f757b38468af04f08c4377a69004aaf21c7f04485b.
//
// Solidity: event Burned(bytes32 indexed burnId, address sender, address receiver, address token, uint256 amount, uint256 dstChainId, uint256 nonce)
func (_Bridgeerc20 *Bridgeerc20Filterer) ParseBurned(log types.Log) (*Bridgeerc20Burned, error) {
	event := new(Bridgeerc20Burned)
	if err := _Bridgeerc20.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgeerc20InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridgeerc20 contract.
type Bridgeerc20InitializedIterator struct {
	Event *Bridgeerc20Initialized // Event containing the contract specifics and raw log

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
func (it *Bridgeerc20InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgeerc20Initialized)
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
		it.Event = new(Bridgeerc20Initialized)
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
func (it *Bridgeerc20InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgeerc20InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgeerc20Initialized represents a Initialized event raised by the Bridgeerc20 contract.
type Bridgeerc20Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgeerc20 *Bridgeerc20Filterer) FilterInitialized(opts *bind.FilterOpts) (*Bridgeerc20InitializedIterator, error) {

	logs, sub, err := _Bridgeerc20.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20InitializedIterator{contract: _Bridgeerc20.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgeerc20 *Bridgeerc20Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Bridgeerc20Initialized) (event.Subscription, error) {

	logs, sub, err := _Bridgeerc20.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgeerc20Initialized)
				if err := _Bridgeerc20.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bridgeerc20 *Bridgeerc20Filterer) ParseInitialized(log types.Log) (*Bridgeerc20Initialized, error) {
	event := new(Bridgeerc20Initialized)
	if err := _Bridgeerc20.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgeerc20LiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the Bridgeerc20 contract.
type Bridgeerc20LiquidityAddedIterator struct {
	Event *Bridgeerc20LiquidityAdded // Event containing the contract specifics and raw log

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
func (it *Bridgeerc20LiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgeerc20LiquidityAdded)
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
		it.Event = new(Bridgeerc20LiquidityAdded)
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
func (it *Bridgeerc20LiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgeerc20LiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgeerc20LiquidityAdded represents a LiquidityAdded event raised by the Bridgeerc20 contract.
type Bridgeerc20LiquidityAdded struct {
	Provider  common.Address
	TokenType uint8
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0x9c1e05a78b09e8a473137cd68bb0266d2c13d87daca427d8a74d316907a48679.
//
// Solidity: event LiquidityAdded(address provider, uint8 tokenType, address token, uint256 amount)
func (_Bridgeerc20 *Bridgeerc20Filterer) FilterLiquidityAdded(opts *bind.FilterOpts) (*Bridgeerc20LiquidityAddedIterator, error) {

	logs, sub, err := _Bridgeerc20.contract.FilterLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20LiquidityAddedIterator{contract: _Bridgeerc20.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0x9c1e05a78b09e8a473137cd68bb0266d2c13d87daca427d8a74d316907a48679.
//
// Solidity: event LiquidityAdded(address provider, uint8 tokenType, address token, uint256 amount)
func (_Bridgeerc20 *Bridgeerc20Filterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *Bridgeerc20LiquidityAdded) (event.Subscription, error) {

	logs, sub, err := _Bridgeerc20.contract.WatchLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgeerc20LiquidityAdded)
				if err := _Bridgeerc20.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0x9c1e05a78b09e8a473137cd68bb0266d2c13d87daca427d8a74d316907a48679.
//
// Solidity: event LiquidityAdded(address provider, uint8 tokenType, address token, uint256 amount)
func (_Bridgeerc20 *Bridgeerc20Filterer) ParseLiquidityAdded(log types.Log) (*Bridgeerc20LiquidityAdded, error) {
	event := new(Bridgeerc20LiquidityAdded)
	if err := _Bridgeerc20.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgeerc20LiquidityRevokedIterator is returned from FilterLiquidityRevoked and is used to iterate over the raw logs and unpacked data for LiquidityRevoked events raised by the Bridgeerc20 contract.
type Bridgeerc20LiquidityRevokedIterator struct {
	Event *Bridgeerc20LiquidityRevoked // Event containing the contract specifics and raw log

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
func (it *Bridgeerc20LiquidityRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgeerc20LiquidityRevoked)
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
		it.Event = new(Bridgeerc20LiquidityRevoked)
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
func (it *Bridgeerc20LiquidityRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgeerc20LiquidityRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgeerc20LiquidityRevoked represents a LiquidityRevoked event raised by the Bridgeerc20 contract.
type Bridgeerc20LiquidityRevoked struct {
	Revoker   common.Address
	TokenType uint8
	Token     common.Address
	To        common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRevoked is a free log retrieval operation binding the contract event 0x7fc9a37c96d59ee5d68dfc2637c318e2913afb0bf875060fe10eaf57eb023804.
//
// Solidity: event LiquidityRevoked(address Revoker, uint8 tokenType, address token, address to, uint256 amount)
func (_Bridgeerc20 *Bridgeerc20Filterer) FilterLiquidityRevoked(opts *bind.FilterOpts) (*Bridgeerc20LiquidityRevokedIterator, error) {

	logs, sub, err := _Bridgeerc20.contract.FilterLogs(opts, "LiquidityRevoked")
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20LiquidityRevokedIterator{contract: _Bridgeerc20.contract, event: "LiquidityRevoked", logs: logs, sub: sub}, nil
}

// WatchLiquidityRevoked is a free log subscription operation binding the contract event 0x7fc9a37c96d59ee5d68dfc2637c318e2913afb0bf875060fe10eaf57eb023804.
//
// Solidity: event LiquidityRevoked(address Revoker, uint8 tokenType, address token, address to, uint256 amount)
func (_Bridgeerc20 *Bridgeerc20Filterer) WatchLiquidityRevoked(opts *bind.WatchOpts, sink chan<- *Bridgeerc20LiquidityRevoked) (event.Subscription, error) {

	logs, sub, err := _Bridgeerc20.contract.WatchLogs(opts, "LiquidityRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgeerc20LiquidityRevoked)
				if err := _Bridgeerc20.contract.UnpackLog(event, "LiquidityRevoked", log); err != nil {
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

// ParseLiquidityRevoked is a log parse operation binding the contract event 0x7fc9a37c96d59ee5d68dfc2637c318e2913afb0bf875060fe10eaf57eb023804.
//
// Solidity: event LiquidityRevoked(address Revoker, uint8 tokenType, address token, address to, uint256 amount)
func (_Bridgeerc20 *Bridgeerc20Filterer) ParseLiquidityRevoked(log types.Log) (*Bridgeerc20LiquidityRevoked, error) {
	event := new(Bridgeerc20LiquidityRevoked)
	if err := _Bridgeerc20.contract.UnpackLog(event, "LiquidityRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgeerc20MaxBurnUpdatedIterator is returned from FilterMaxBurnUpdated and is used to iterate over the raw logs and unpacked data for MaxBurnUpdated events raised by the Bridgeerc20 contract.
type Bridgeerc20MaxBurnUpdatedIterator struct {
	Event *Bridgeerc20MaxBurnUpdated // Event containing the contract specifics and raw log

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
func (it *Bridgeerc20MaxBurnUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgeerc20MaxBurnUpdated)
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
		it.Event = new(Bridgeerc20MaxBurnUpdated)
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
func (it *Bridgeerc20MaxBurnUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgeerc20MaxBurnUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgeerc20MaxBurnUpdated represents a MaxBurnUpdated event raised by the Bridgeerc20 contract.
type Bridgeerc20MaxBurnUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxBurnUpdated is a free log retrieval operation binding the contract event 0xa3181379f6db47d9037efc6b6e8e3efe8c55ddb090b4f0512c152f97c4e47da5.
//
// Solidity: event MaxBurnUpdated(address token, uint256 amount)
func (_Bridgeerc20 *Bridgeerc20Filterer) FilterMaxBurnUpdated(opts *bind.FilterOpts) (*Bridgeerc20MaxBurnUpdatedIterator, error) {

	logs, sub, err := _Bridgeerc20.contract.FilterLogs(opts, "MaxBurnUpdated")
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20MaxBurnUpdatedIterator{contract: _Bridgeerc20.contract, event: "MaxBurnUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxBurnUpdated is a free log subscription operation binding the contract event 0xa3181379f6db47d9037efc6b6e8e3efe8c55ddb090b4f0512c152f97c4e47da5.
//
// Solidity: event MaxBurnUpdated(address token, uint256 amount)
func (_Bridgeerc20 *Bridgeerc20Filterer) WatchMaxBurnUpdated(opts *bind.WatchOpts, sink chan<- *Bridgeerc20MaxBurnUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridgeerc20.contract.WatchLogs(opts, "MaxBurnUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgeerc20MaxBurnUpdated)
				if err := _Bridgeerc20.contract.UnpackLog(event, "MaxBurnUpdated", log); err != nil {
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

// ParseMaxBurnUpdated is a log parse operation binding the contract event 0xa3181379f6db47d9037efc6b6e8e3efe8c55ddb090b4f0512c152f97c4e47da5.
//
// Solidity: event MaxBurnUpdated(address token, uint256 amount)
func (_Bridgeerc20 *Bridgeerc20Filterer) ParseMaxBurnUpdated(log types.Log) (*Bridgeerc20MaxBurnUpdated, error) {
	event := new(Bridgeerc20MaxBurnUpdated)
	if err := _Bridgeerc20.contract.UnpackLog(event, "MaxBurnUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgeerc20MinBurnUpdatedIterator is returned from FilterMinBurnUpdated and is used to iterate over the raw logs and unpacked data for MinBurnUpdated events raised by the Bridgeerc20 contract.
type Bridgeerc20MinBurnUpdatedIterator struct {
	Event *Bridgeerc20MinBurnUpdated // Event containing the contract specifics and raw log

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
func (it *Bridgeerc20MinBurnUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgeerc20MinBurnUpdated)
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
		it.Event = new(Bridgeerc20MinBurnUpdated)
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
func (it *Bridgeerc20MinBurnUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgeerc20MinBurnUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgeerc20MinBurnUpdated represents a MinBurnUpdated event raised by the Bridgeerc20 contract.
type Bridgeerc20MinBurnUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinBurnUpdated is a free log retrieval operation binding the contract event 0x3796cd0b17a8734f8da819920625598e9a18be490f686725282e5383f1d06683.
//
// Solidity: event MinBurnUpdated(address token, uint256 amount)
func (_Bridgeerc20 *Bridgeerc20Filterer) FilterMinBurnUpdated(opts *bind.FilterOpts) (*Bridgeerc20MinBurnUpdatedIterator, error) {

	logs, sub, err := _Bridgeerc20.contract.FilterLogs(opts, "MinBurnUpdated")
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20MinBurnUpdatedIterator{contract: _Bridgeerc20.contract, event: "MinBurnUpdated", logs: logs, sub: sub}, nil
}

// WatchMinBurnUpdated is a free log subscription operation binding the contract event 0x3796cd0b17a8734f8da819920625598e9a18be490f686725282e5383f1d06683.
//
// Solidity: event MinBurnUpdated(address token, uint256 amount)
func (_Bridgeerc20 *Bridgeerc20Filterer) WatchMinBurnUpdated(opts *bind.WatchOpts, sink chan<- *Bridgeerc20MinBurnUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridgeerc20.contract.WatchLogs(opts, "MinBurnUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgeerc20MinBurnUpdated)
				if err := _Bridgeerc20.contract.UnpackLog(event, "MinBurnUpdated", log); err != nil {
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

// ParseMinBurnUpdated is a log parse operation binding the contract event 0x3796cd0b17a8734f8da819920625598e9a18be490f686725282e5383f1d06683.
//
// Solidity: event MinBurnUpdated(address token, uint256 amount)
func (_Bridgeerc20 *Bridgeerc20Filterer) ParseMinBurnUpdated(log types.Log) (*Bridgeerc20MinBurnUpdated, error) {
	event := new(Bridgeerc20MinBurnUpdated)
	if err := _Bridgeerc20.contract.UnpackLog(event, "MinBurnUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgeerc20MintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Bridgeerc20 contract.
type Bridgeerc20MintedIterator struct {
	Event *Bridgeerc20Minted // Event containing the contract specifics and raw log

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
func (it *Bridgeerc20MintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgeerc20Minted)
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
		it.Event = new(Bridgeerc20Minted)
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
func (it *Bridgeerc20MintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgeerc20MintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgeerc20Minted represents a Minted event raised by the Bridgeerc20 contract.
type Bridgeerc20Minted struct {
	MintId     [32]byte
	BurnId     [32]byte
	Sender     common.Address
	Receiver   common.Address
	Token      common.Address
	Amount     *big.Int
	Fee        *big.Int
	RefChainId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x42a59c71a311ec1a2cbacd2fab15803a860d84419a5349269fac707e07071ffa.
//
// Solidity: event Minted(bytes32 indexed mintId, bytes32 indexed burnId, address sender, address receiver, address token, uint256 amount, uint256 fee, uint256 refChainId)
func (_Bridgeerc20 *Bridgeerc20Filterer) FilterMinted(opts *bind.FilterOpts, mintId [][32]byte, burnId [][32]byte) (*Bridgeerc20MintedIterator, error) {

	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}
	var burnIdRule []interface{}
	for _, burnIdItem := range burnId {
		burnIdRule = append(burnIdRule, burnIdItem)
	}

	logs, sub, err := _Bridgeerc20.contract.FilterLogs(opts, "Minted", mintIdRule, burnIdRule)
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20MintedIterator{contract: _Bridgeerc20.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x42a59c71a311ec1a2cbacd2fab15803a860d84419a5349269fac707e07071ffa.
//
// Solidity: event Minted(bytes32 indexed mintId, bytes32 indexed burnId, address sender, address receiver, address token, uint256 amount, uint256 fee, uint256 refChainId)
func (_Bridgeerc20 *Bridgeerc20Filterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *Bridgeerc20Minted, mintId [][32]byte, burnId [][32]byte) (event.Subscription, error) {

	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}
	var burnIdRule []interface{}
	for _, burnIdItem := range burnId {
		burnIdRule = append(burnIdRule, burnIdItem)
	}

	logs, sub, err := _Bridgeerc20.contract.WatchLogs(opts, "Minted", mintIdRule, burnIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgeerc20Minted)
				if err := _Bridgeerc20.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x42a59c71a311ec1a2cbacd2fab15803a860d84419a5349269fac707e07071ffa.
//
// Solidity: event Minted(bytes32 indexed mintId, bytes32 indexed burnId, address sender, address receiver, address token, uint256 amount, uint256 fee, uint256 refChainId)
func (_Bridgeerc20 *Bridgeerc20Filterer) ParseMinted(log types.Log) (*Bridgeerc20Minted, error) {
	event := new(Bridgeerc20Minted)
	if err := _Bridgeerc20.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgeerc20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridgeerc20 contract.
type Bridgeerc20OwnershipTransferredIterator struct {
	Event *Bridgeerc20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Bridgeerc20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgeerc20OwnershipTransferred)
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
		it.Event = new(Bridgeerc20OwnershipTransferred)
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
func (it *Bridgeerc20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgeerc20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgeerc20OwnershipTransferred represents a OwnershipTransferred event raised by the Bridgeerc20 contract.
type Bridgeerc20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridgeerc20 *Bridgeerc20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Bridgeerc20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridgeerc20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20OwnershipTransferredIterator{contract: _Bridgeerc20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridgeerc20 *Bridgeerc20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Bridgeerc20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridgeerc20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgeerc20OwnershipTransferred)
				if err := _Bridgeerc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridgeerc20 *Bridgeerc20Filterer) ParseOwnershipTransferred(log types.Log) (*Bridgeerc20OwnershipTransferred, error) {
	event := new(Bridgeerc20OwnershipTransferred)
	if err := _Bridgeerc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgeerc20PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridgeerc20 contract.
type Bridgeerc20PausedIterator struct {
	Event *Bridgeerc20Paused // Event containing the contract specifics and raw log

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
func (it *Bridgeerc20PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgeerc20Paused)
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
		it.Event = new(Bridgeerc20Paused)
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
func (it *Bridgeerc20PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgeerc20PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgeerc20Paused represents a Paused event raised by the Bridgeerc20 contract.
type Bridgeerc20Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridgeerc20 *Bridgeerc20Filterer) FilterPaused(opts *bind.FilterOpts) (*Bridgeerc20PausedIterator, error) {

	logs, sub, err := _Bridgeerc20.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20PausedIterator{contract: _Bridgeerc20.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridgeerc20 *Bridgeerc20Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Bridgeerc20Paused) (event.Subscription, error) {

	logs, sub, err := _Bridgeerc20.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgeerc20Paused)
				if err := _Bridgeerc20.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Bridgeerc20 *Bridgeerc20Filterer) ParsePaused(log types.Log) (*Bridgeerc20Paused, error) {
	event := new(Bridgeerc20Paused)
	if err := _Bridgeerc20.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgeerc20SetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the Bridgeerc20 contract.
type Bridgeerc20SetAdminIterator struct {
	Event *Bridgeerc20SetAdmin // Event containing the contract specifics and raw log

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
func (it *Bridgeerc20SetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgeerc20SetAdmin)
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
		it.Event = new(Bridgeerc20SetAdmin)
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
func (it *Bridgeerc20SetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgeerc20SetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgeerc20SetAdmin represents a SetAdmin event raised by the Bridgeerc20 contract.
type Bridgeerc20SetAdmin struct {
	Admin common.Address
	Auth  bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x55a5194bc0174fcaf12b2978bef43911466bf63b34db8d1dd1a0d5dcd5c41bea.
//
// Solidity: event SetAdmin(address admin, bool auth)
func (_Bridgeerc20 *Bridgeerc20Filterer) FilterSetAdmin(opts *bind.FilterOpts) (*Bridgeerc20SetAdminIterator, error) {

	logs, sub, err := _Bridgeerc20.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20SetAdminIterator{contract: _Bridgeerc20.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x55a5194bc0174fcaf12b2978bef43911466bf63b34db8d1dd1a0d5dcd5c41bea.
//
// Solidity: event SetAdmin(address admin, bool auth)
func (_Bridgeerc20 *Bridgeerc20Filterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *Bridgeerc20SetAdmin) (event.Subscription, error) {

	logs, sub, err := _Bridgeerc20.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgeerc20SetAdmin)
				if err := _Bridgeerc20.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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

// ParseSetAdmin is a log parse operation binding the contract event 0x55a5194bc0174fcaf12b2978bef43911466bf63b34db8d1dd1a0d5dcd5c41bea.
//
// Solidity: event SetAdmin(address admin, bool auth)
func (_Bridgeerc20 *Bridgeerc20Filterer) ParseSetAdmin(log types.Log) (*Bridgeerc20SetAdmin, error) {
	event := new(Bridgeerc20SetAdmin)
	if err := _Bridgeerc20.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgeerc20UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridgeerc20 contract.
type Bridgeerc20UnpausedIterator struct {
	Event *Bridgeerc20Unpaused // Event containing the contract specifics and raw log

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
func (it *Bridgeerc20UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgeerc20Unpaused)
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
		it.Event = new(Bridgeerc20Unpaused)
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
func (it *Bridgeerc20UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgeerc20UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgeerc20Unpaused represents a Unpaused event raised by the Bridgeerc20 contract.
type Bridgeerc20Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridgeerc20 *Bridgeerc20Filterer) FilterUnpaused(opts *bind.FilterOpts) (*Bridgeerc20UnpausedIterator, error) {

	logs, sub, err := _Bridgeerc20.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Bridgeerc20UnpausedIterator{contract: _Bridgeerc20.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridgeerc20 *Bridgeerc20Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Bridgeerc20Unpaused) (event.Subscription, error) {

	logs, sub, err := _Bridgeerc20.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgeerc20Unpaused)
				if err := _Bridgeerc20.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Bridgeerc20 *Bridgeerc20Filterer) ParseUnpaused(log types.Log) (*Bridgeerc20Unpaused, error) {
	event := new(Bridgeerc20Unpaused)
	if err := _Bridgeerc20.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
