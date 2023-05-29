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

// L1MessageSenderMetaData contains all meta data concerning the L1MessageSender contract.
var L1MessageSenderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"L1MessageSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendToL1\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L1MessageSenderABI is the input ABI used to generate the binding from.
// Deprecated: Use L1MessageSenderMetaData.ABI instead.
var L1MessageSenderABI = L1MessageSenderMetaData.ABI

// L1MessageSender is an auto generated Go binding around an Ethereum contract.
type L1MessageSender struct {
	L1MessageSenderCaller     // Read-only binding to the contract
	L1MessageSenderTransactor // Write-only binding to the contract
	L1MessageSenderFilterer   // Log filterer for contract events
}

// L1MessageSenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1MessageSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageSenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1MessageSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageSenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1MessageSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageSenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1MessageSenderSession struct {
	Contract     *L1MessageSender  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1MessageSenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1MessageSenderCallerSession struct {
	Contract *L1MessageSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L1MessageSenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1MessageSenderTransactorSession struct {
	Contract     *L1MessageSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L1MessageSenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1MessageSenderRaw struct {
	Contract *L1MessageSender // Generic contract binding to access the raw methods on
}

// L1MessageSenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1MessageSenderCallerRaw struct {
	Contract *L1MessageSenderCaller // Generic read-only contract binding to access the raw methods on
}

// L1MessageSenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1MessageSenderTransactorRaw struct {
	Contract *L1MessageSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1MessageSender creates a new instance of L1MessageSender, bound to a specific deployed contract.
func NewL1MessageSender(address common.Address, backend bind.ContractBackend) (*L1MessageSender, error) {
	contract, err := bindL1MessageSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1MessageSender{L1MessageSenderCaller: L1MessageSenderCaller{contract: contract}, L1MessageSenderTransactor: L1MessageSenderTransactor{contract: contract}, L1MessageSenderFilterer: L1MessageSenderFilterer{contract: contract}}, nil
}

// NewL1MessageSenderCaller creates a new read-only instance of L1MessageSender, bound to a specific deployed contract.
func NewL1MessageSenderCaller(address common.Address, caller bind.ContractCaller) (*L1MessageSenderCaller, error) {
	contract, err := bindL1MessageSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1MessageSenderCaller{contract: contract}, nil
}

// NewL1MessageSenderTransactor creates a new write-only instance of L1MessageSender, bound to a specific deployed contract.
func NewL1MessageSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*L1MessageSenderTransactor, error) {
	contract, err := bindL1MessageSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1MessageSenderTransactor{contract: contract}, nil
}

// NewL1MessageSenderFilterer creates a new log filterer instance of L1MessageSender, bound to a specific deployed contract.
func NewL1MessageSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*L1MessageSenderFilterer, error) {
	contract, err := bindL1MessageSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1MessageSenderFilterer{contract: contract}, nil
}

// bindL1MessageSender binds a generic wrapper to an already deployed contract.
func bindL1MessageSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1MessageSenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1MessageSender *L1MessageSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1MessageSender.Contract.L1MessageSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1MessageSender *L1MessageSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1MessageSender.Contract.L1MessageSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1MessageSender *L1MessageSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1MessageSender.Contract.L1MessageSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1MessageSender *L1MessageSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1MessageSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1MessageSender *L1MessageSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1MessageSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1MessageSender *L1MessageSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1MessageSender.Contract.contract.Transact(opts, method, params...)
}

// SendToL1 is a paid mutator transaction binding the contract method 0x62f84b24.
//
// Solidity: function sendToL1(bytes _message) returns(bytes32)
func (_L1MessageSender *L1MessageSenderTransactor) SendToL1(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1MessageSender.contract.Transact(opts, "sendToL1", _message)
}

// SendToL1 is a paid mutator transaction binding the contract method 0x62f84b24.
//
// Solidity: function sendToL1(bytes _message) returns(bytes32)
func (_L1MessageSender *L1MessageSenderSession) SendToL1(_message []byte) (*types.Transaction, error) {
	return _L1MessageSender.Contract.SendToL1(&_L1MessageSender.TransactOpts, _message)
}

// SendToL1 is a paid mutator transaction binding the contract method 0x62f84b24.
//
// Solidity: function sendToL1(bytes _message) returns(bytes32)
func (_L1MessageSender *L1MessageSenderTransactorSession) SendToL1(_message []byte) (*types.Transaction, error) {
	return _L1MessageSender.Contract.SendToL1(&_L1MessageSender.TransactOpts, _message)
}

// L1MessageSenderL1MessageSentIterator is returned from FilterL1MessageSent and is used to iterate over the raw logs and unpacked data for L1MessageSent events raised by the L1MessageSender contract.
type L1MessageSenderL1MessageSentIterator struct {
	Event *L1MessageSenderL1MessageSent // Event containing the contract specifics and raw log

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
func (it *L1MessageSenderL1MessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageSenderL1MessageSent)
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
		it.Event = new(L1MessageSenderL1MessageSent)
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
func (it *L1MessageSenderL1MessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageSenderL1MessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageSenderL1MessageSent represents a L1MessageSent event raised by the L1MessageSender contract.
type L1MessageSenderL1MessageSent struct {
	Sender  common.Address
	Hash    [32]byte
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterL1MessageSent is a free log retrieval operation binding the contract event 0x3a36e47291f4201faf137fab081d92295bce2d53be2c6ca68ba82c7faa9ce241.
//
// Solidity: event L1MessageSent(address indexed _sender, bytes32 indexed _hash, bytes _message)
func (_L1MessageSender *L1MessageSenderFilterer) FilterL1MessageSent(opts *bind.FilterOpts, _sender []common.Address, _hash [][32]byte) (*L1MessageSenderL1MessageSentIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _L1MessageSender.contract.FilterLogs(opts, "L1MessageSent", _senderRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return &L1MessageSenderL1MessageSentIterator{contract: _L1MessageSender.contract, event: "L1MessageSent", logs: logs, sub: sub}, nil
}

// WatchL1MessageSent is a free log subscription operation binding the contract event 0x3a36e47291f4201faf137fab081d92295bce2d53be2c6ca68ba82c7faa9ce241.
//
// Solidity: event L1MessageSent(address indexed _sender, bytes32 indexed _hash, bytes _message)
func (_L1MessageSender *L1MessageSenderFilterer) WatchL1MessageSent(opts *bind.WatchOpts, sink chan<- *L1MessageSenderL1MessageSent, _sender []common.Address, _hash [][32]byte) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _L1MessageSender.contract.WatchLogs(opts, "L1MessageSent", _senderRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageSenderL1MessageSent)
				if err := _L1MessageSender.contract.UnpackLog(event, "L1MessageSent", log); err != nil {
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

// ParseL1MessageSent is a log parse operation binding the contract event 0x3a36e47291f4201faf137fab081d92295bce2d53be2c6ca68ba82c7faa9ce241.
//
// Solidity: event L1MessageSent(address indexed _sender, bytes32 indexed _hash, bytes _message)
func (_L1MessageSender *L1MessageSenderFilterer) ParseL1MessageSent(log types.Log) (*L1MessageSenderL1MessageSent, error) {
	event := new(L1MessageSenderL1MessageSent)
	if err := _L1MessageSender.contract.UnpackLog(event, "L1MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
