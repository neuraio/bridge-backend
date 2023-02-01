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

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcNft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Callback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChId\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcNft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"reciver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstNft\",\"type\":\"address\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_srcNft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstChId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_srcNft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bridgeCallBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_srcNft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChId\",\"type\":\"uint256\"}],\"name\":\"delDestNftAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"destNFTAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"destTxFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcChid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dstNft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"sendTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_auth\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_srcNft\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_dstChIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_dstNfts\",\"type\":\"address[]\"}],\"name\":\"setDestNftAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_transfer\",\"type\":\"bool\"}],\"name\":\"setOrigNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_chids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"}],\"name\":\"setTxFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Bridge *BridgeCaller) Admin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "admin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Bridge *BridgeSession) Admin(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Admin(&_Bridge.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) Admin(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Admin(&_Bridge.CallOpts, arg0)
}

// DestNFTAddr is a free data retrieval call binding the contract method 0xb1f2b1c9.
//
// Solidity: function destNFTAddr(address , uint256 ) view returns(address)
func (_Bridge *BridgeCaller) DestNFTAddr(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "destNFTAddr", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DestNFTAddr is a free data retrieval call binding the contract method 0xb1f2b1c9.
//
// Solidity: function destNFTAddr(address , uint256 ) view returns(address)
func (_Bridge *BridgeSession) DestNFTAddr(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Bridge.Contract.DestNFTAddr(&_Bridge.CallOpts, arg0, arg1)
}

// DestNFTAddr is a free data retrieval call binding the contract method 0xb1f2b1c9.
//
// Solidity: function destNFTAddr(address , uint256 ) view returns(address)
func (_Bridge *BridgeCallerSession) DestNFTAddr(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Bridge.Contract.DestNFTAddr(&_Bridge.CallOpts, arg0, arg1)
}

// DestTxFee is a free data retrieval call binding the contract method 0x0035797a.
//
// Solidity: function destTxFee(uint256 ) view returns(uint256)
func (_Bridge *BridgeCaller) DestTxFee(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "destTxFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DestTxFee is a free data retrieval call binding the contract method 0x0035797a.
//
// Solidity: function destTxFee(uint256 ) view returns(uint256)
func (_Bridge *BridgeSession) DestTxFee(arg0 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.DestTxFee(&_Bridge.CallOpts, arg0)
}

// DestTxFee is a free data retrieval call binding the contract method 0x0035797a.
//
// Solidity: function destTxFee(uint256 ) view returns(uint256)
func (_Bridge *BridgeCallerSession) DestTxFee(arg0 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.DestTxFee(&_Bridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Bridge is a paid mutator transaction binding the contract method 0x5e583a5a.
//
// Solidity: function bridge(address _srcNft, uint256 _tokenId, uint256 _dstChId, address _receiver) payable returns()
func (_Bridge *BridgeTransactor) Bridge(opts *bind.TransactOpts, _srcNft common.Address, _tokenId *big.Int, _dstChId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "bridge", _srcNft, _tokenId, _dstChId, _receiver)
}

// Bridge is a paid mutator transaction binding the contract method 0x5e583a5a.
//
// Solidity: function bridge(address _srcNft, uint256 _tokenId, uint256 _dstChId, address _receiver) payable returns()
func (_Bridge *BridgeSession) Bridge(_srcNft common.Address, _tokenId *big.Int, _dstChId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Bridge(&_Bridge.TransactOpts, _srcNft, _tokenId, _dstChId, _receiver)
}

// Bridge is a paid mutator transaction binding the contract method 0x5e583a5a.
//
// Solidity: function bridge(address _srcNft, uint256 _tokenId, uint256 _dstChId, address _receiver) payable returns()
func (_Bridge *BridgeTransactorSession) Bridge(_srcNft common.Address, _tokenId *big.Int, _dstChId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Bridge(&_Bridge.TransactOpts, _srcNft, _tokenId, _dstChId, _receiver)
}

// BridgeCallBack is a paid mutator transaction binding the contract method 0x021412bf.
//
// Solidity: function bridgeCallBack(address _sender, address _srcNft, uint256 _tokenId) returns()
func (_Bridge *BridgeTransactor) BridgeCallBack(opts *bind.TransactOpts, _sender common.Address, _srcNft common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "bridgeCallBack", _sender, _srcNft, _tokenId)
}

// BridgeCallBack is a paid mutator transaction binding the contract method 0x021412bf.
//
// Solidity: function bridgeCallBack(address _sender, address _srcNft, uint256 _tokenId) returns()
func (_Bridge *BridgeSession) BridgeCallBack(_sender common.Address, _srcNft common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeCallBack(&_Bridge.TransactOpts, _sender, _srcNft, _tokenId)
}

// BridgeCallBack is a paid mutator transaction binding the contract method 0x021412bf.
//
// Solidity: function bridgeCallBack(address _sender, address _srcNft, uint256 _tokenId) returns()
func (_Bridge *BridgeTransactorSession) BridgeCallBack(_sender common.Address, _srcNft common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeCallBack(&_Bridge.TransactOpts, _sender, _srcNft, _tokenId)
}

// DelDestNftAddr is a paid mutator transaction binding the contract method 0x39b3d3d2.
//
// Solidity: function delDestNftAddr(address _srcNft, uint256 _dstChId) returns()
func (_Bridge *BridgeTransactor) DelDestNftAddr(opts *bind.TransactOpts, _srcNft common.Address, _dstChId *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "delDestNftAddr", _srcNft, _dstChId)
}

// DelDestNftAddr is a paid mutator transaction binding the contract method 0x39b3d3d2.
//
// Solidity: function delDestNftAddr(address _srcNft, uint256 _dstChId) returns()
func (_Bridge *BridgeSession) DelDestNftAddr(_srcNft common.Address, _dstChId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DelDestNftAddr(&_Bridge.TransactOpts, _srcNft, _dstChId)
}

// DelDestNftAddr is a paid mutator transaction binding the contract method 0x39b3d3d2.
//
// Solidity: function delDestNftAddr(address _srcNft, uint256 _dstChId) returns()
func (_Bridge *BridgeTransactorSession) DelDestNftAddr(_srcNft common.Address, _dstChId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DelDestNftAddr(&_Bridge.TransactOpts, _srcNft, _dstChId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Bridge *BridgeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Bridge *BridgeSession) Initialize() (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Bridge *BridgeTransactorSession) Initialize() (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// SendTo is a paid mutator transaction binding the contract method 0x9d53468d.
//
// Solidity: function sendTo(uint256 _srcChid, address _dstNft, uint256 _tokenId, address _receiver) returns()
func (_Bridge *BridgeTransactor) SendTo(opts *bind.TransactOpts, _srcChid *big.Int, _dstNft common.Address, _tokenId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "sendTo", _srcChid, _dstNft, _tokenId, _receiver)
}

// SendTo is a paid mutator transaction binding the contract method 0x9d53468d.
//
// Solidity: function sendTo(uint256 _srcChid, address _dstNft, uint256 _tokenId, address _receiver) returns()
func (_Bridge *BridgeSession) SendTo(_srcChid *big.Int, _dstNft common.Address, _tokenId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SendTo(&_Bridge.TransactOpts, _srcChid, _dstNft, _tokenId, _receiver)
}

// SendTo is a paid mutator transaction binding the contract method 0x9d53468d.
//
// Solidity: function sendTo(uint256 _srcChid, address _dstNft, uint256 _tokenId, address _receiver) returns()
func (_Bridge *BridgeTransactorSession) SendTo(_srcChid *big.Int, _dstNft common.Address, _tokenId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SendTo(&_Bridge.TransactOpts, _srcChid, _dstNft, _tokenId, _receiver)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _user, bool _auth) returns()
func (_Bridge *BridgeTransactor) SetAdmin(opts *bind.TransactOpts, _user common.Address, _auth bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setAdmin", _user, _auth)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _user, bool _auth) returns()
func (_Bridge *BridgeSession) SetAdmin(_user common.Address, _auth bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetAdmin(&_Bridge.TransactOpts, _user, _auth)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _user, bool _auth) returns()
func (_Bridge *BridgeTransactorSession) SetAdmin(_user common.Address, _auth bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetAdmin(&_Bridge.TransactOpts, _user, _auth)
}

// SetDestNftAddr is a paid mutator transaction binding the contract method 0xfded01f0.
//
// Solidity: function setDestNftAddr(address _srcNft, uint256[] _dstChIds, address[] _dstNfts) returns()
func (_Bridge *BridgeTransactor) SetDestNftAddr(opts *bind.TransactOpts, _srcNft common.Address, _dstChIds []*big.Int, _dstNfts []common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setDestNftAddr", _srcNft, _dstChIds, _dstNfts)
}

// SetDestNftAddr is a paid mutator transaction binding the contract method 0xfded01f0.
//
// Solidity: function setDestNftAddr(address _srcNft, uint256[] _dstChIds, address[] _dstNfts) returns()
func (_Bridge *BridgeSession) SetDestNftAddr(_srcNft common.Address, _dstChIds []*big.Int, _dstNfts []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetDestNftAddr(&_Bridge.TransactOpts, _srcNft, _dstChIds, _dstNfts)
}

// SetDestNftAddr is a paid mutator transaction binding the contract method 0xfded01f0.
//
// Solidity: function setDestNftAddr(address _srcNft, uint256[] _dstChIds, address[] _dstNfts) returns()
func (_Bridge *BridgeTransactorSession) SetDestNftAddr(_srcNft common.Address, _dstChIds []*big.Int, _dstNfts []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetDestNftAddr(&_Bridge.TransactOpts, _srcNft, _dstChIds, _dstNfts)
}

// SetOrigNFT is a paid mutator transaction binding the contract method 0x3c299323.
//
// Solidity: function setOrigNFT(address _nft, bool _transfer) returns()
func (_Bridge *BridgeTransactor) SetOrigNFT(opts *bind.TransactOpts, _nft common.Address, _transfer bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setOrigNFT", _nft, _transfer)
}

// SetOrigNFT is a paid mutator transaction binding the contract method 0x3c299323.
//
// Solidity: function setOrigNFT(address _nft, bool _transfer) returns()
func (_Bridge *BridgeSession) SetOrigNFT(_nft common.Address, _transfer bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetOrigNFT(&_Bridge.TransactOpts, _nft, _transfer)
}

// SetOrigNFT is a paid mutator transaction binding the contract method 0x3c299323.
//
// Solidity: function setOrigNFT(address _nft, bool _transfer) returns()
func (_Bridge *BridgeTransactorSession) SetOrigNFT(_nft common.Address, _transfer bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetOrigNFT(&_Bridge.TransactOpts, _nft, _transfer)
}

// SetTxFee is a paid mutator transaction binding the contract method 0x11939376.
//
// Solidity: function setTxFee(uint256[] _chids, uint256[] _fees) returns()
func (_Bridge *BridgeTransactor) SetTxFee(opts *bind.TransactOpts, _chids []*big.Int, _fees []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setTxFee", _chids, _fees)
}

// SetTxFee is a paid mutator transaction binding the contract method 0x11939376.
//
// Solidity: function setTxFee(uint256[] _chids, uint256[] _fees) returns()
func (_Bridge *BridgeSession) SetTxFee(_chids []*big.Int, _fees []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetTxFee(&_Bridge.TransactOpts, _chids, _fees)
}

// SetTxFee is a paid mutator transaction binding the contract method 0x11939376.
//
// Solidity: function setTxFee(uint256[] _chids, uint256[] _fees) returns()
func (_Bridge *BridgeTransactorSession) SetTxFee(_chids []*big.Int, _fees []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetTxFee(&_Bridge.TransactOpts, _chids, _fees)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_Bridge *BridgeTransactor) Withdraw(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdraw", _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_Bridge *BridgeSession) Withdraw(_to common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_Bridge *BridgeTransactorSession) Withdraw(_to common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, _to)
}

// BridgeCallbackIterator is returned from FilterCallback and is used to iterate over the raw logs and unpacked data for Callback events raised by the Bridge contract.
type BridgeCallbackIterator struct {
	Event *BridgeCallback // Event containing the contract specifics and raw log

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
func (it *BridgeCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCallback)
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
		it.Event = new(BridgeCallback)
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
func (it *BridgeCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCallback represents a Callback event raised by the Bridge contract.
type BridgeCallback struct {
	Sender  common.Address
	SrcNft  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCallback is a free log retrieval operation binding the contract event 0xf8230cf0826ce9fb1fa42b50b6e67817f00af3545417652d78da6106004d8a62.
//
// Solidity: event Callback(address sender, address srcNft, uint256 tokenId)
func (_Bridge *BridgeFilterer) FilterCallback(opts *bind.FilterOpts) (*BridgeCallbackIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Callback")
	if err != nil {
		return nil, err
	}
	return &BridgeCallbackIterator{contract: _Bridge.contract, event: "Callback", logs: logs, sub: sub}, nil
}

// WatchCallback is a free log subscription operation binding the contract event 0xf8230cf0826ce9fb1fa42b50b6e67817f00af3545417652d78da6106004d8a62.
//
// Solidity: event Callback(address sender, address srcNft, uint256 tokenId)
func (_Bridge *BridgeFilterer) WatchCallback(opts *bind.WatchOpts, sink chan<- *BridgeCallback) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Callback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCallback)
				if err := _Bridge.contract.UnpackLog(event, "Callback", log); err != nil {
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

// ParseCallback is a log parse operation binding the contract event 0xf8230cf0826ce9fb1fa42b50b6e67817f00af3545417652d78da6106004d8a62.
//
// Solidity: event Callback(address sender, address srcNft, uint256 tokenId)
func (_Bridge *BridgeFilterer) ParseCallback(log types.Log) (*BridgeCallback, error) {
	event := new(BridgeCallback)
	if err := _Bridge.contract.UnpackLog(event, "Callback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridge contract.
type BridgeInitializedIterator struct {
	Event *BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeInitialized)
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
		it.Event = new(BridgeInitialized)
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
func (it *BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeInitialized represents a Initialized event raised by the Bridge contract.
type BridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeInitializedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeInitializedIterator{contract: _Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeInitialized)
				if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseInitialized(log types.Log) (*BridgeInitialized, error) {
	event := new(BridgeInitialized)
	if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Bridge contract.
type BridgeReceivedIterator struct {
	Event *BridgeReceived // Event containing the contract specifics and raw log

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
func (it *BridgeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeReceived)
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
		it.Event = new(BridgeReceived)
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
func (it *BridgeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeReceived represents a Received event raised by the Bridge contract.
type BridgeReceived struct {
	Receiver common.Address
	Nft      common.Address
	Id       *big.Int
	SrcChId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x14432f6e1dc0e8c1f4c0d81c69cecc80c0bea817a74482492b0211392478ab9b.
//
// Solidity: event Received(address receiver, address nft, uint256 id, uint256 srcChId)
func (_Bridge *BridgeFilterer) FilterReceived(opts *bind.FilterOpts) (*BridgeReceivedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &BridgeReceivedIterator{contract: _Bridge.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x14432f6e1dc0e8c1f4c0d81c69cecc80c0bea817a74482492b0211392478ab9b.
//
// Solidity: event Received(address receiver, address nft, uint256 id, uint256 srcChId)
func (_Bridge *BridgeFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *BridgeReceived) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeReceived)
				if err := _Bridge.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x14432f6e1dc0e8c1f4c0d81c69cecc80c0bea817a74482492b0211392478ab9b.
//
// Solidity: event Received(address receiver, address nft, uint256 id, uint256 srcChId)
func (_Bridge *BridgeFilterer) ParseReceived(log types.Log) (*BridgeReceived, error) {
	event := new(BridgeReceived)
	if err := _Bridge.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the Bridge contract.
type BridgeSentIterator struct {
	Event *BridgeSent // Event containing the contract specifics and raw log

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
func (it *BridgeSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSent)
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
		it.Event = new(BridgeSent)
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
func (it *BridgeSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSent represents a Sent event raised by the Bridge contract.
type BridgeSent struct {
	Sender  common.Address
	SrcNft  common.Address
	TokenId *big.Int
	DstChId *big.Int
	Reciver common.Address
	DstNft  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0xa32c8d97b98adc135b448bc36f8fd4fbdc09c4a7bd832e5e9a0510051a75f89c.
//
// Solidity: event Sent(address sender, address srcNft, uint256 tokenId, uint256 dstChId, address reciver, address dstNft)
func (_Bridge *BridgeFilterer) FilterSent(opts *bind.FilterOpts) (*BridgeSentIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return &BridgeSentIterator{contract: _Bridge.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0xa32c8d97b98adc135b448bc36f8fd4fbdc09c4a7bd832e5e9a0510051a75f89c.
//
// Solidity: event Sent(address sender, address srcNft, uint256 tokenId, uint256 dstChId, address reciver, address dstNft)
func (_Bridge *BridgeFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *BridgeSent) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSent)
				if err := _Bridge.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0xa32c8d97b98adc135b448bc36f8fd4fbdc09c4a7bd832e5e9a0510051a75f89c.
//
// Solidity: event Sent(address sender, address srcNft, uint256 tokenId, uint256 dstChId, address reciver, address dstNft)
func (_Bridge *BridgeFilterer) ParseSent(log types.Log) (*BridgeSent, error) {
	event := new(BridgeSent)
	if err := _Bridge.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
