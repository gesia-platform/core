// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// NotaryPublicStoreMetaData contains all meta data concerning the NotaryPublicStore contract.
var NotaryPublicStoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"prefix\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signatrue\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"Notarized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"prefix\",\"type\":\"bytes1\"},{\"internalType\":\"bytes\",\"name\":\"signatrue\",\"type\":\"bytes\"}],\"name\":\"notarize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NotaryPublicStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use NotaryPublicStoreMetaData.ABI instead.
var NotaryPublicStoreABI = NotaryPublicStoreMetaData.ABI

// NotaryPublicStore is an auto generated Go binding around an Ethereum contract.
type NotaryPublicStore struct {
	NotaryPublicStoreCaller     // Read-only binding to the contract
	NotaryPublicStoreTransactor // Write-only binding to the contract
	NotaryPublicStoreFilterer   // Log filterer for contract events
}

// NotaryPublicStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type NotaryPublicStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryPublicStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NotaryPublicStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryPublicStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NotaryPublicStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryPublicStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NotaryPublicStoreSession struct {
	Contract     *NotaryPublicStore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NotaryPublicStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NotaryPublicStoreCallerSession struct {
	Contract *NotaryPublicStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NotaryPublicStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NotaryPublicStoreTransactorSession struct {
	Contract     *NotaryPublicStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NotaryPublicStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type NotaryPublicStoreRaw struct {
	Contract *NotaryPublicStore // Generic contract binding to access the raw methods on
}

// NotaryPublicStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NotaryPublicStoreCallerRaw struct {
	Contract *NotaryPublicStoreCaller // Generic read-only contract binding to access the raw methods on
}

// NotaryPublicStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NotaryPublicStoreTransactorRaw struct {
	Contract *NotaryPublicStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNotaryPublicStore creates a new instance of NotaryPublicStore, bound to a specific deployed contract.
func NewNotaryPublicStore(address common.Address, backend bind.ContractBackend) (*NotaryPublicStore, error) {
	contract, err := bindNotaryPublicStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NotaryPublicStore{NotaryPublicStoreCaller: NotaryPublicStoreCaller{contract: contract}, NotaryPublicStoreTransactor: NotaryPublicStoreTransactor{contract: contract}, NotaryPublicStoreFilterer: NotaryPublicStoreFilterer{contract: contract}}, nil
}

// NewNotaryPublicStoreCaller creates a new read-only instance of NotaryPublicStore, bound to a specific deployed contract.
func NewNotaryPublicStoreCaller(address common.Address, caller bind.ContractCaller) (*NotaryPublicStoreCaller, error) {
	contract, err := bindNotaryPublicStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NotaryPublicStoreCaller{contract: contract}, nil
}

// NewNotaryPublicStoreTransactor creates a new write-only instance of NotaryPublicStore, bound to a specific deployed contract.
func NewNotaryPublicStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*NotaryPublicStoreTransactor, error) {
	contract, err := bindNotaryPublicStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NotaryPublicStoreTransactor{contract: contract}, nil
}

// NewNotaryPublicStoreFilterer creates a new log filterer instance of NotaryPublicStore, bound to a specific deployed contract.
func NewNotaryPublicStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*NotaryPublicStoreFilterer, error) {
	contract, err := bindNotaryPublicStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NotaryPublicStoreFilterer{contract: contract}, nil
}

// bindNotaryPublicStore binds a generic wrapper to an already deployed contract.
func bindNotaryPublicStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NotaryPublicStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NotaryPublicStore *NotaryPublicStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NotaryPublicStore.Contract.NotaryPublicStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NotaryPublicStore *NotaryPublicStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.NotaryPublicStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NotaryPublicStore *NotaryPublicStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.NotaryPublicStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NotaryPublicStore *NotaryPublicStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NotaryPublicStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NotaryPublicStore *NotaryPublicStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NotaryPublicStore *NotaryPublicStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.contract.Transact(opts, method, params...)
}

// Notarize is a paid mutator transaction binding the contract method 0xa0538ecd.
//
// Solidity: function notarize(bytes1 prefix, bytes signatrue) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactor) Notarize(opts *bind.TransactOpts, prefix [1]byte, signatrue []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.contract.Transact(opts, "notarize", prefix, signatrue)
}

// Notarize is a paid mutator transaction binding the contract method 0xa0538ecd.
//
// Solidity: function notarize(bytes1 prefix, bytes signatrue) returns()
func (_NotaryPublicStore *NotaryPublicStoreSession) Notarize(prefix [1]byte, signatrue []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.Notarize(&_NotaryPublicStore.TransactOpts, prefix, signatrue)
}

// Notarize is a paid mutator transaction binding the contract method 0xa0538ecd.
//
// Solidity: function notarize(bytes1 prefix, bytes signatrue) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactorSession) Notarize(prefix [1]byte, signatrue []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.Notarize(&_NotaryPublicStore.TransactOpts, prefix, signatrue)
}

// NotaryPublicStoreNotarizedIterator is returned from FilterNotarized and is used to iterate over the raw logs and unpacked data for Notarized events raised by the NotaryPublicStore contract.
type NotaryPublicStoreNotarizedIterator struct {
	Event *NotaryPublicStoreNotarized // Event containing the contract specifics and raw log

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
func (it *NotaryPublicStoreNotarizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryPublicStoreNotarized)
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
		it.Event = new(NotaryPublicStoreNotarized)
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
func (it *NotaryPublicStoreNotarizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryPublicStoreNotarizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryPublicStoreNotarized represents a Notarized event raised by the NotaryPublicStore contract.
type NotaryPublicStoreNotarized struct {
	Prefix    [1]byte
	Signatrue []byte
	Notary    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNotarized is a free log retrieval operation binding the contract event 0xacdb5596a58de2795a84369a204c419750a07d77b12d8743575630619249ae3f.
//
// Solidity: event Notarized(bytes1 prefix, bytes signatrue, address notary)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) FilterNotarized(opts *bind.FilterOpts) (*NotaryPublicStoreNotarizedIterator, error) {

	logs, sub, err := _NotaryPublicStore.contract.FilterLogs(opts, "Notarized")
	if err != nil {
		return nil, err
	}
	return &NotaryPublicStoreNotarizedIterator{contract: _NotaryPublicStore.contract, event: "Notarized", logs: logs, sub: sub}, nil
}

// WatchNotarized is a free log subscription operation binding the contract event 0xacdb5596a58de2795a84369a204c419750a07d77b12d8743575630619249ae3f.
//
// Solidity: event Notarized(bytes1 prefix, bytes signatrue, address notary)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) WatchNotarized(opts *bind.WatchOpts, sink chan<- *NotaryPublicStoreNotarized) (event.Subscription, error) {

	logs, sub, err := _NotaryPublicStore.contract.WatchLogs(opts, "Notarized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryPublicStoreNotarized)
				if err := _NotaryPublicStore.contract.UnpackLog(event, "Notarized", log); err != nil {
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

// ParseNotarized is a log parse operation binding the contract event 0xacdb5596a58de2795a84369a204c419750a07d77b12d8743575630619249ae3f.
//
// Solidity: event Notarized(bytes1 prefix, bytes signatrue, address notary)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) ParseNotarized(log types.Log) (*NotaryPublicStoreNotarized, error) {
	event := new(NotaryPublicStoreNotarized)
	if err := _NotaryPublicStore.contract.UnpackLog(event, "Notarized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
