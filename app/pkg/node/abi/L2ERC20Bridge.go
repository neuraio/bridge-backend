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

// L2ERC20BridgeMetaData contains all meta data concerning the L2ERC20Bridge contract.
var L2ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FinalizeDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TokenProxyBytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"l1TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"l2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBeacon\",\"outputs\":[{\"internalType\":\"contractUpgradeableBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L2ERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ERC20BridgeMetaData.ABI instead.
var L2ERC20BridgeABI = L2ERC20BridgeMetaData.ABI

// L2ERC20Bridge is an auto generated Go binding around an Ethereum contract.
type L2ERC20Bridge struct {
	L2ERC20BridgeCaller     // Read-only binding to the contract
	L2ERC20BridgeTransactor // Write-only binding to the contract
	L2ERC20BridgeFilterer   // Log filterer for contract events
}

// L2ERC20BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ERC20BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC20BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ERC20BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC20BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ERC20BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC20BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ERC20BridgeSession struct {
	Contract     *L2ERC20Bridge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2ERC20BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ERC20BridgeCallerSession struct {
	Contract *L2ERC20BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L2ERC20BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ERC20BridgeTransactorSession struct {
	Contract     *L2ERC20BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L2ERC20BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ERC20BridgeRaw struct {
	Contract *L2ERC20Bridge // Generic contract binding to access the raw methods on
}

// L2ERC20BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ERC20BridgeCallerRaw struct {
	Contract *L2ERC20BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L2ERC20BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ERC20BridgeTransactorRaw struct {
	Contract *L2ERC20BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ERC20Bridge creates a new instance of L2ERC20Bridge, bound to a specific deployed contract.
func NewL2ERC20Bridge(address common.Address, backend bind.ContractBackend) (*L2ERC20Bridge, error) {
	contract, err := bindL2ERC20Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ERC20Bridge{L2ERC20BridgeCaller: L2ERC20BridgeCaller{contract: contract}, L2ERC20BridgeTransactor: L2ERC20BridgeTransactor{contract: contract}, L2ERC20BridgeFilterer: L2ERC20BridgeFilterer{contract: contract}}, nil
}

// NewL2ERC20BridgeCaller creates a new read-only instance of L2ERC20Bridge, bound to a specific deployed contract.
func NewL2ERC20BridgeCaller(address common.Address, caller bind.ContractCaller) (*L2ERC20BridgeCaller, error) {
	contract, err := bindL2ERC20Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC20BridgeCaller{contract: contract}, nil
}

// NewL2ERC20BridgeTransactor creates a new write-only instance of L2ERC20Bridge, bound to a specific deployed contract.
func NewL2ERC20BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ERC20BridgeTransactor, error) {
	contract, err := bindL2ERC20Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC20BridgeTransactor{contract: contract}, nil
}

// NewL2ERC20BridgeFilterer creates a new log filterer instance of L2ERC20Bridge, bound to a specific deployed contract.
func NewL2ERC20BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ERC20BridgeFilterer, error) {
	contract, err := bindL2ERC20Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ERC20BridgeFilterer{contract: contract}, nil
}

// bindL2ERC20Bridge binds a generic wrapper to an already deployed contract.
func bindL2ERC20Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ERC20BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ERC20Bridge *L2ERC20BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ERC20Bridge.Contract.L2ERC20BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ERC20Bridge *L2ERC20BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ERC20Bridge.Contract.L2ERC20BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ERC20Bridge *L2ERC20BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ERC20Bridge.Contract.L2ERC20BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ERC20Bridge *L2ERC20BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ERC20Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ERC20Bridge *L2ERC20BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ERC20Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ERC20Bridge *L2ERC20BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ERC20Bridge.Contract.contract.Transact(opts, method, params...)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_L2ERC20Bridge *L2ERC20BridgeCaller) L1Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC20Bridge.contract.Call(opts, &out, "l1Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_L2ERC20Bridge *L2ERC20BridgeSession) L1Bridge() (common.Address, error) {
	return _L2ERC20Bridge.Contract.L1Bridge(&_L2ERC20Bridge.CallOpts)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_L2ERC20Bridge *L2ERC20BridgeCallerSession) L1Bridge() (common.Address, error) {
	return _L2ERC20Bridge.Contract.L1Bridge(&_L2ERC20Bridge.CallOpts)
}

// L1TokenAddress is a free data retrieval call binding the contract method 0xf54266a2.
//
// Solidity: function l1TokenAddress(address ) view returns(address)
func (_L2ERC20Bridge *L2ERC20BridgeCaller) L1TokenAddress(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ERC20Bridge.contract.Call(opts, &out, "l1TokenAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1TokenAddress is a free data retrieval call binding the contract method 0xf54266a2.
//
// Solidity: function l1TokenAddress(address ) view returns(address)
func (_L2ERC20Bridge *L2ERC20BridgeSession) L1TokenAddress(arg0 common.Address) (common.Address, error) {
	return _L2ERC20Bridge.Contract.L1TokenAddress(&_L2ERC20Bridge.CallOpts, arg0)
}

// L1TokenAddress is a free data retrieval call binding the contract method 0xf54266a2.
//
// Solidity: function l1TokenAddress(address ) view returns(address)
func (_L2ERC20Bridge *L2ERC20BridgeCallerSession) L1TokenAddress(arg0 common.Address) (common.Address, error) {
	return _L2ERC20Bridge.Contract.L1TokenAddress(&_L2ERC20Bridge.CallOpts, arg0)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L2ERC20Bridge *L2ERC20BridgeCaller) L2TokenAddress(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ERC20Bridge.contract.Call(opts, &out, "l2TokenAddress", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L2ERC20Bridge *L2ERC20BridgeSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _L2ERC20Bridge.Contract.L2TokenAddress(&_L2ERC20Bridge.CallOpts, _l1Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L2ERC20Bridge *L2ERC20BridgeCallerSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _L2ERC20Bridge.Contract.L2TokenAddress(&_L2ERC20Bridge.CallOpts, _l1Token)
}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_L2ERC20Bridge *L2ERC20BridgeCaller) L2TokenBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC20Bridge.contract.Call(opts, &out, "l2TokenBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_L2ERC20Bridge *L2ERC20BridgeSession) L2TokenBeacon() (common.Address, error) {
	return _L2ERC20Bridge.Contract.L2TokenBeacon(&_L2ERC20Bridge.CallOpts)
}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_L2ERC20Bridge *L2ERC20BridgeCallerSession) L2TokenBeacon() (common.Address, error) {
	return _L2ERC20Bridge.Contract.L2TokenBeacon(&_L2ERC20Bridge.CallOpts)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0xcfe7af7c.
//
// Solidity: function finalizeDeposit(address _l1Sender, address _l2Receiver, address _l1Token, uint256 _amount, bytes _data) returns()
func (_L2ERC20Bridge *L2ERC20BridgeTransactor) FinalizeDeposit(opts *bind.TransactOpts, _l1Sender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ERC20Bridge.contract.Transact(opts, "finalizeDeposit", _l1Sender, _l2Receiver, _l1Token, _amount, _data)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0xcfe7af7c.
//
// Solidity: function finalizeDeposit(address _l1Sender, address _l2Receiver, address _l1Token, uint256 _amount, bytes _data) returns()
func (_L2ERC20Bridge *L2ERC20BridgeSession) FinalizeDeposit(_l1Sender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ERC20Bridge.Contract.FinalizeDeposit(&_L2ERC20Bridge.TransactOpts, _l1Sender, _l2Receiver, _l1Token, _amount, _data)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0xcfe7af7c.
//
// Solidity: function finalizeDeposit(address _l1Sender, address _l2Receiver, address _l1Token, uint256 _amount, bytes _data) returns()
func (_L2ERC20Bridge *L2ERC20BridgeTransactorSession) FinalizeDeposit(_l1Sender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ERC20Bridge.Contract.FinalizeDeposit(&_L2ERC20Bridge.TransactOpts, _l1Sender, _l2Receiver, _l1Token, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xd26b3e6e.
//
// Solidity: function initialize(address _l1Bridge, bytes32 _l2TokenProxyBytecodeHash, address _governor) returns()
func (_L2ERC20Bridge *L2ERC20BridgeTransactor) Initialize(opts *bind.TransactOpts, _l1Bridge common.Address, _l2TokenProxyBytecodeHash [32]byte, _governor common.Address) (*types.Transaction, error) {
	return _L2ERC20Bridge.contract.Transact(opts, "initialize", _l1Bridge, _l2TokenProxyBytecodeHash, _governor)
}

// Initialize is a paid mutator transaction binding the contract method 0xd26b3e6e.
//
// Solidity: function initialize(address _l1Bridge, bytes32 _l2TokenProxyBytecodeHash, address _governor) returns()
func (_L2ERC20Bridge *L2ERC20BridgeSession) Initialize(_l1Bridge common.Address, _l2TokenProxyBytecodeHash [32]byte, _governor common.Address) (*types.Transaction, error) {
	return _L2ERC20Bridge.Contract.Initialize(&_L2ERC20Bridge.TransactOpts, _l1Bridge, _l2TokenProxyBytecodeHash, _governor)
}

// Initialize is a paid mutator transaction binding the contract method 0xd26b3e6e.
//
// Solidity: function initialize(address _l1Bridge, bytes32 _l2TokenProxyBytecodeHash, address _governor) returns()
func (_L2ERC20Bridge *L2ERC20BridgeTransactorSession) Initialize(_l1Bridge common.Address, _l2TokenProxyBytecodeHash [32]byte, _governor common.Address) (*types.Transaction, error) {
	return _L2ERC20Bridge.Contract.Initialize(&_L2ERC20Bridge.TransactOpts, _l1Bridge, _l2TokenProxyBytecodeHash, _governor)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _l1Receiver, address _l2Token, uint256 _amount) returns()
func (_L2ERC20Bridge *L2ERC20BridgeTransactor) Withdraw(opts *bind.TransactOpts, _l1Receiver common.Address, _l2Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2ERC20Bridge.contract.Transact(opts, "withdraw", _l1Receiver, _l2Token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _l1Receiver, address _l2Token, uint256 _amount) returns()
func (_L2ERC20Bridge *L2ERC20BridgeSession) Withdraw(_l1Receiver common.Address, _l2Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2ERC20Bridge.Contract.Withdraw(&_L2ERC20Bridge.TransactOpts, _l1Receiver, _l2Token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _l1Receiver, address _l2Token, uint256 _amount) returns()
func (_L2ERC20Bridge *L2ERC20BridgeTransactorSession) Withdraw(_l1Receiver common.Address, _l2Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2ERC20Bridge.Contract.Withdraw(&_L2ERC20Bridge.TransactOpts, _l1Receiver, _l2Token, _amount)
}

// L2ERC20BridgeFinalizeDepositIterator is returned from FilterFinalizeDeposit and is used to iterate over the raw logs and unpacked data for FinalizeDeposit events raised by the L2ERC20Bridge contract.
type L2ERC20BridgeFinalizeDepositIterator struct {
	Event *L2ERC20BridgeFinalizeDeposit // Event containing the contract specifics and raw log

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
func (it *L2ERC20BridgeFinalizeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC20BridgeFinalizeDeposit)
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
		it.Event = new(L2ERC20BridgeFinalizeDeposit)
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
func (it *L2ERC20BridgeFinalizeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC20BridgeFinalizeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC20BridgeFinalizeDeposit represents a FinalizeDeposit event raised by the L2ERC20Bridge contract.
type L2ERC20BridgeFinalizeDeposit struct {
	L1Sender   common.Address
	L2Receiver common.Address
	L2Token    common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDeposit is a free log retrieval operation binding the contract event 0xb84fba9af218da60d299dc177abd5805e7ac541d2673cbee7808c10017874f63.
//
// Solidity: event FinalizeDeposit(address indexed l1Sender, address indexed l2Receiver, address indexed l2Token, uint256 amount)
func (_L2ERC20Bridge *L2ERC20BridgeFilterer) FilterFinalizeDeposit(opts *bind.FilterOpts, l1Sender []common.Address, l2Receiver []common.Address, l2Token []common.Address) (*L2ERC20BridgeFinalizeDepositIterator, error) {

	var l1SenderRule []interface{}
	for _, l1SenderItem := range l1Sender {
		l1SenderRule = append(l1SenderRule, l1SenderItem)
	}
	var l2ReceiverRule []interface{}
	for _, l2ReceiverItem := range l2Receiver {
		l2ReceiverRule = append(l2ReceiverRule, l2ReceiverItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _L2ERC20Bridge.contract.FilterLogs(opts, "FinalizeDeposit", l1SenderRule, l2ReceiverRule, l2TokenRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC20BridgeFinalizeDepositIterator{contract: _L2ERC20Bridge.contract, event: "FinalizeDeposit", logs: logs, sub: sub}, nil
}

// WatchFinalizeDeposit is a free log subscription operation binding the contract event 0xb84fba9af218da60d299dc177abd5805e7ac541d2673cbee7808c10017874f63.
//
// Solidity: event FinalizeDeposit(address indexed l1Sender, address indexed l2Receiver, address indexed l2Token, uint256 amount)
func (_L2ERC20Bridge *L2ERC20BridgeFilterer) WatchFinalizeDeposit(opts *bind.WatchOpts, sink chan<- *L2ERC20BridgeFinalizeDeposit, l1Sender []common.Address, l2Receiver []common.Address, l2Token []common.Address) (event.Subscription, error) {

	var l1SenderRule []interface{}
	for _, l1SenderItem := range l1Sender {
		l1SenderRule = append(l1SenderRule, l1SenderItem)
	}
	var l2ReceiverRule []interface{}
	for _, l2ReceiverItem := range l2Receiver {
		l2ReceiverRule = append(l2ReceiverRule, l2ReceiverItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _L2ERC20Bridge.contract.WatchLogs(opts, "FinalizeDeposit", l1SenderRule, l2ReceiverRule, l2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC20BridgeFinalizeDeposit)
				if err := _L2ERC20Bridge.contract.UnpackLog(event, "FinalizeDeposit", log); err != nil {
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

// ParseFinalizeDeposit is a log parse operation binding the contract event 0xb84fba9af218da60d299dc177abd5805e7ac541d2673cbee7808c10017874f63.
//
// Solidity: event FinalizeDeposit(address indexed l1Sender, address indexed l2Receiver, address indexed l2Token, uint256 amount)
func (_L2ERC20Bridge *L2ERC20BridgeFilterer) ParseFinalizeDeposit(log types.Log) (*L2ERC20BridgeFinalizeDeposit, error) {
	event := new(L2ERC20BridgeFinalizeDeposit)
	if err := _L2ERC20Bridge.contract.UnpackLog(event, "FinalizeDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC20BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2ERC20Bridge contract.
type L2ERC20BridgeInitializedIterator struct {
	Event *L2ERC20BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *L2ERC20BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC20BridgeInitialized)
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
		it.Event = new(L2ERC20BridgeInitialized)
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
func (it *L2ERC20BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC20BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC20BridgeInitialized represents a Initialized event raised by the L2ERC20Bridge contract.
type L2ERC20BridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ERC20Bridge *L2ERC20BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2ERC20BridgeInitializedIterator, error) {

	logs, sub, err := _L2ERC20Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2ERC20BridgeInitializedIterator{contract: _L2ERC20Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ERC20Bridge *L2ERC20BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2ERC20BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _L2ERC20Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC20BridgeInitialized)
				if err := _L2ERC20Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2ERC20Bridge *L2ERC20BridgeFilterer) ParseInitialized(log types.Log) (*L2ERC20BridgeInitialized, error) {
	event := new(L2ERC20BridgeInitialized)
	if err := _L2ERC20Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC20BridgeWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the L2ERC20Bridge contract.
type L2ERC20BridgeWithdrawalInitiatedIterator struct {
	Event *L2ERC20BridgeWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *L2ERC20BridgeWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC20BridgeWithdrawalInitiated)
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
		it.Event = new(L2ERC20BridgeWithdrawalInitiated)
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
func (it *L2ERC20BridgeWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC20BridgeWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC20BridgeWithdrawalInitiated represents a WithdrawalInitiated event raised by the L2ERC20Bridge contract.
type L2ERC20BridgeWithdrawalInitiated struct {
	L2Sender   common.Address
	L1Receiver common.Address
	L2Token    common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x2fc3848834aac8e883a2d2a17a7514dc4f2d3dd268089df9b9f5d918259ef3b0.
//
// Solidity: event WithdrawalInitiated(address indexed l2Sender, address indexed l1Receiver, address indexed l2Token, uint256 amount)
func (_L2ERC20Bridge *L2ERC20BridgeFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, l2Sender []common.Address, l1Receiver []common.Address, l2Token []common.Address) (*L2ERC20BridgeWithdrawalInitiatedIterator, error) {

	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var l1ReceiverRule []interface{}
	for _, l1ReceiverItem := range l1Receiver {
		l1ReceiverRule = append(l1ReceiverRule, l1ReceiverItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _L2ERC20Bridge.contract.FilterLogs(opts, "WithdrawalInitiated", l2SenderRule, l1ReceiverRule, l2TokenRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC20BridgeWithdrawalInitiatedIterator{contract: _L2ERC20Bridge.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x2fc3848834aac8e883a2d2a17a7514dc4f2d3dd268089df9b9f5d918259ef3b0.
//
// Solidity: event WithdrawalInitiated(address indexed l2Sender, address indexed l1Receiver, address indexed l2Token, uint256 amount)
func (_L2ERC20Bridge *L2ERC20BridgeFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *L2ERC20BridgeWithdrawalInitiated, l2Sender []common.Address, l1Receiver []common.Address, l2Token []common.Address) (event.Subscription, error) {

	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var l1ReceiverRule []interface{}
	for _, l1ReceiverItem := range l1Receiver {
		l1ReceiverRule = append(l1ReceiverRule, l1ReceiverItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _L2ERC20Bridge.contract.WatchLogs(opts, "WithdrawalInitiated", l2SenderRule, l1ReceiverRule, l2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC20BridgeWithdrawalInitiated)
				if err := _L2ERC20Bridge.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x2fc3848834aac8e883a2d2a17a7514dc4f2d3dd268089df9b9f5d918259ef3b0.
//
// Solidity: event WithdrawalInitiated(address indexed l2Sender, address indexed l1Receiver, address indexed l2Token, uint256 amount)
func (_L2ERC20Bridge *L2ERC20BridgeFilterer) ParseWithdrawalInitiated(log types.Log) (*L2ERC20BridgeWithdrawalInitiated, error) {
	event := new(L2ERC20BridgeWithdrawalInitiated)
	if err := _L2ERC20Bridge.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
