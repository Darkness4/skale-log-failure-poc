// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exampleabi

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
	_ = abi.ConvertType
)

// EventExampleMetaData contains all meta data concerning the EventExample contract.
var EventExampleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"NewEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"launchEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506101ca806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063614e216714610030575b600080fd5b61004361003e366004610095565b610045565b005b7fc49fa2d8e562e80e750f98e95e2178dc3201633e4f247b6474aa6d8fafe29dba816040516100749190610146565b60405180910390a150565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156100a757600080fd5b813567ffffffffffffffff808211156100bf57600080fd5b818401915084601f8301126100d357600080fd5b8135818111156100e5576100e561007f565b604051601f8201601f19908116603f0116810190838211818310171561010d5761010d61007f565b8160405282815287602084870101111561012657600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b8181101561017357858101830151858201604001528201610157565b506000604082860101526040601f19601f830116850101925050509291505056fea264697066735822122075263336845de257f9453262e80cafbbb47677bd280470a40da75256826bdf7664736f6c63430008150033",
}

// EventExampleABI is the input ABI used to generate the binding from.
// Deprecated: Use EventExampleMetaData.ABI instead.
var EventExampleABI = EventExampleMetaData.ABI

// EventExampleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EventExampleMetaData.Bin instead.
var EventExampleBin = EventExampleMetaData.Bin

// DeployEventExample deploys a new Ethereum contract, binding an instance of EventExample to it.
func DeployEventExample(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EventExample, error) {
	parsed, err := EventExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EventExampleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EventExample{EventExampleCaller: EventExampleCaller{contract: contract}, EventExampleTransactor: EventExampleTransactor{contract: contract}, EventExampleFilterer: EventExampleFilterer{contract: contract}}, nil
}

// EventExample is an auto generated Go binding around an Ethereum contract.
type EventExample struct {
	EventExampleCaller     // Read-only binding to the contract
	EventExampleTransactor // Write-only binding to the contract
	EventExampleFilterer   // Log filterer for contract events
}

// EventExampleCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventExampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventExampleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventExampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventExampleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventExampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventExampleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventExampleSession struct {
	Contract     *EventExample     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventExampleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventExampleCallerSession struct {
	Contract *EventExampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EventExampleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventExampleTransactorSession struct {
	Contract     *EventExampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EventExampleRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventExampleRaw struct {
	Contract *EventExample // Generic contract binding to access the raw methods on
}

// EventExampleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventExampleCallerRaw struct {
	Contract *EventExampleCaller // Generic read-only contract binding to access the raw methods on
}

// EventExampleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventExampleTransactorRaw struct {
	Contract *EventExampleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventExample creates a new instance of EventExample, bound to a specific deployed contract.
func NewEventExample(address common.Address, backend bind.ContractBackend) (*EventExample, error) {
	contract, err := bindEventExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventExample{EventExampleCaller: EventExampleCaller{contract: contract}, EventExampleTransactor: EventExampleTransactor{contract: contract}, EventExampleFilterer: EventExampleFilterer{contract: contract}}, nil
}

// NewEventExampleCaller creates a new read-only instance of EventExample, bound to a specific deployed contract.
func NewEventExampleCaller(address common.Address, caller bind.ContractCaller) (*EventExampleCaller, error) {
	contract, err := bindEventExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventExampleCaller{contract: contract}, nil
}

// NewEventExampleTransactor creates a new write-only instance of EventExample, bound to a specific deployed contract.
func NewEventExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*EventExampleTransactor, error) {
	contract, err := bindEventExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventExampleTransactor{contract: contract}, nil
}

// NewEventExampleFilterer creates a new log filterer instance of EventExample, bound to a specific deployed contract.
func NewEventExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*EventExampleFilterer, error) {
	contract, err := bindEventExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventExampleFilterer{contract: contract}, nil
}

// bindEventExample binds a generic wrapper to an already deployed contract.
func bindEventExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventExampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventExample *EventExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventExample.Contract.EventExampleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventExample *EventExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventExample.Contract.EventExampleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventExample *EventExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventExample.Contract.EventExampleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventExample *EventExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventExample.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventExample *EventExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventExample.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventExample *EventExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventExample.Contract.contract.Transact(opts, method, params...)
}

// LaunchEvent is a paid mutator transaction binding the contract method 0x614e2167.
//
// Solidity: function launchEvent(string _message) returns()
func (_EventExample *EventExampleTransactor) LaunchEvent(opts *bind.TransactOpts, _message string) (*types.Transaction, error) {
	return _EventExample.contract.Transact(opts, "launchEvent", _message)
}

// LaunchEvent is a paid mutator transaction binding the contract method 0x614e2167.
//
// Solidity: function launchEvent(string _message) returns()
func (_EventExample *EventExampleSession) LaunchEvent(_message string) (*types.Transaction, error) {
	return _EventExample.Contract.LaunchEvent(&_EventExample.TransactOpts, _message)
}

// LaunchEvent is a paid mutator transaction binding the contract method 0x614e2167.
//
// Solidity: function launchEvent(string _message) returns()
func (_EventExample *EventExampleTransactorSession) LaunchEvent(_message string) (*types.Transaction, error) {
	return _EventExample.Contract.LaunchEvent(&_EventExample.TransactOpts, _message)
}

// EventExampleNewEventIterator is returned from FilterNewEvent and is used to iterate over the raw logs and unpacked data for NewEvent events raised by the EventExample contract.
type EventExampleNewEventIterator struct {
	Event *EventExampleNewEvent // Event containing the contract specifics and raw log

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
func (it *EventExampleNewEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventExampleNewEvent)
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
		it.Event = new(EventExampleNewEvent)
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
func (it *EventExampleNewEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventExampleNewEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventExampleNewEvent represents a NewEvent event raised by the EventExample contract.
type EventExampleNewEvent struct {
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewEvent is a free log retrieval operation binding the contract event 0xc49fa2d8e562e80e750f98e95e2178dc3201633e4f247b6474aa6d8fafe29dba.
//
// Solidity: event NewEvent(string _message)
func (_EventExample *EventExampleFilterer) FilterNewEvent(opts *bind.FilterOpts) (*EventExampleNewEventIterator, error) {

	logs, sub, err := _EventExample.contract.FilterLogs(opts, "NewEvent")
	if err != nil {
		return nil, err
	}
	return &EventExampleNewEventIterator{contract: _EventExample.contract, event: "NewEvent", logs: logs, sub: sub}, nil
}

// WatchNewEvent is a free log subscription operation binding the contract event 0xc49fa2d8e562e80e750f98e95e2178dc3201633e4f247b6474aa6d8fafe29dba.
//
// Solidity: event NewEvent(string _message)
func (_EventExample *EventExampleFilterer) WatchNewEvent(opts *bind.WatchOpts, sink chan<- *EventExampleNewEvent) (event.Subscription, error) {

	logs, sub, err := _EventExample.contract.WatchLogs(opts, "NewEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventExampleNewEvent)
				if err := _EventExample.contract.UnpackLog(event, "NewEvent", log); err != nil {
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

// ParseNewEvent is a log parse operation binding the contract event 0xc49fa2d8e562e80e750f98e95e2178dc3201633e4f247b6474aa6d8fafe29dba.
//
// Solidity: event NewEvent(string _message)
func (_EventExample *EventExampleFilterer) ParseNewEvent(log types.Log) (*EventExampleNewEvent, error) {
	event := new(EventExampleNewEvent)
	if err := _EventExample.contract.UnpackLog(event, "NewEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
