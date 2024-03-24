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

// AppStoreStoreMetaData contains all meta data concerning the AppStoreStore contract.
var AppStoreStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AppCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"applicationID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"AppOwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferAppOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"}],\"name\":\"getApp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AppStoreStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use AppStoreStoreMetaData.ABI instead.
var AppStoreStoreABI = AppStoreStoreMetaData.ABI

// AppStoreStore is an auto generated Go binding around an Ethereum contract.
type AppStoreStore struct {
	AppStoreStoreCaller     // Read-only binding to the contract
	AppStoreStoreTransactor // Write-only binding to the contract
	AppStoreStoreFilterer   // Log filterer for contract events
}

// AppStoreStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppStoreStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppStoreStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppStoreStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppStoreStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppStoreStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppStoreStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppStoreStoreSession struct {
	Contract     *AppStoreStore    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppStoreStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppStoreStoreCallerSession struct {
	Contract *AppStoreStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AppStoreStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppStoreStoreTransactorSession struct {
	Contract     *AppStoreStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AppStoreStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppStoreStoreRaw struct {
	Contract *AppStoreStore // Generic contract binding to access the raw methods on
}

// AppStoreStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppStoreStoreCallerRaw struct {
	Contract *AppStoreStoreCaller // Generic read-only contract binding to access the raw methods on
}

// AppStoreStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppStoreStoreTransactorRaw struct {
	Contract *AppStoreStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppStoreStore creates a new instance of AppStoreStore, bound to a specific deployed contract.
func NewAppStoreStore(address common.Address, backend bind.ContractBackend) (*AppStoreStore, error) {
	contract, err := bindAppStoreStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AppStoreStore{AppStoreStoreCaller: AppStoreStoreCaller{contract: contract}, AppStoreStoreTransactor: AppStoreStoreTransactor{contract: contract}, AppStoreStoreFilterer: AppStoreStoreFilterer{contract: contract}}, nil
}

// NewAppStoreStoreCaller creates a new read-only instance of AppStoreStore, bound to a specific deployed contract.
func NewAppStoreStoreCaller(address common.Address, caller bind.ContractCaller) (*AppStoreStoreCaller, error) {
	contract, err := bindAppStoreStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppStoreStoreCaller{contract: contract}, nil
}

// NewAppStoreStoreTransactor creates a new write-only instance of AppStoreStore, bound to a specific deployed contract.
func NewAppStoreStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*AppStoreStoreTransactor, error) {
	contract, err := bindAppStoreStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppStoreStoreTransactor{contract: contract}, nil
}

// NewAppStoreStoreFilterer creates a new log filterer instance of AppStoreStore, bound to a specific deployed contract.
func NewAppStoreStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*AppStoreStoreFilterer, error) {
	contract, err := bindAppStoreStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppStoreStoreFilterer{contract: contract}, nil
}

// bindAppStoreStore binds a generic wrapper to an already deployed contract.
func bindAppStoreStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AppStoreStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppStoreStore *AppStoreStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppStoreStore.Contract.AppStoreStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppStoreStore *AppStoreStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppStoreStore.Contract.AppStoreStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppStoreStore *AppStoreStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppStoreStore.Contract.AppStoreStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppStoreStore *AppStoreStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppStoreStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppStoreStore *AppStoreStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppStoreStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppStoreStore *AppStoreStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppStoreStore.Contract.contract.Transact(opts, method, params...)
}

// GetApp is a free data retrieval call binding the contract method 0x24f3a51b.
//
// Solidity: function getApp(uint256 appID) view returns(uint256, string, address)
func (_AppStoreStore *AppStoreStoreCaller) GetApp(opts *bind.CallOpts, appID *big.Int) (*big.Int, string, common.Address, error) {
	var out []interface{}
	err := _AppStoreStore.contract.Call(opts, &out, "getApp", appID)

	if err != nil {
		return *new(*big.Int), *new(string), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// GetApp is a free data retrieval call binding the contract method 0x24f3a51b.
//
// Solidity: function getApp(uint256 appID) view returns(uint256, string, address)
func (_AppStoreStore *AppStoreStoreSession) GetApp(appID *big.Int) (*big.Int, string, common.Address, error) {
	return _AppStoreStore.Contract.GetApp(&_AppStoreStore.CallOpts, appID)
}

// GetApp is a free data retrieval call binding the contract method 0x24f3a51b.
//
// Solidity: function getApp(uint256 appID) view returns(uint256, string, address)
func (_AppStoreStore *AppStoreStoreCallerSession) GetApp(appID *big.Int) (*big.Int, string, common.Address, error) {
	return _AppStoreStore.Contract.GetApp(&_AppStoreStore.CallOpts, appID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AppStoreStore *AppStoreStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppStoreStore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AppStoreStore *AppStoreStoreSession) Owner() (common.Address, error) {
	return _AppStoreStore.Contract.Owner(&_AppStoreStore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AppStoreStore *AppStoreStoreCallerSession) Owner() (common.Address, error) {
	return _AppStoreStore.Contract.Owner(&_AppStoreStore.CallOpts)
}

// CreateApp is a paid mutator transaction binding the contract method 0x4ded94d1.
//
// Solidity: function createApp(string domain, address owner) returns()
func (_AppStoreStore *AppStoreStoreTransactor) CreateApp(opts *bind.TransactOpts, domain string, owner common.Address) (*types.Transaction, error) {
	return _AppStoreStore.contract.Transact(opts, "createApp", domain, owner)
}

// CreateApp is a paid mutator transaction binding the contract method 0x4ded94d1.
//
// Solidity: function createApp(string domain, address owner) returns()
func (_AppStoreStore *AppStoreStoreSession) CreateApp(domain string, owner common.Address) (*types.Transaction, error) {
	return _AppStoreStore.Contract.CreateApp(&_AppStoreStore.TransactOpts, domain, owner)
}

// CreateApp is a paid mutator transaction binding the contract method 0x4ded94d1.
//
// Solidity: function createApp(string domain, address owner) returns()
func (_AppStoreStore *AppStoreStoreTransactorSession) CreateApp(domain string, owner common.Address) (*types.Transaction, error) {
	return _AppStoreStore.Contract.CreateApp(&_AppStoreStore.TransactOpts, domain, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppStoreStore *AppStoreStoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppStoreStore.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppStoreStore *AppStoreStoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _AppStoreStore.Contract.RenounceOwnership(&_AppStoreStore.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppStoreStore *AppStoreStoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AppStoreStore.Contract.RenounceOwnership(&_AppStoreStore.TransactOpts)
}

// TransferAppOwnership is a paid mutator transaction binding the contract method 0x6de20cb5.
//
// Solidity: function transferAppOwnership(uint256 appID, address newOwner) returns()
func (_AppStoreStore *AppStoreStoreTransactor) TransferAppOwnership(opts *bind.TransactOpts, appID *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _AppStoreStore.contract.Transact(opts, "transferAppOwnership", appID, newOwner)
}

// TransferAppOwnership is a paid mutator transaction binding the contract method 0x6de20cb5.
//
// Solidity: function transferAppOwnership(uint256 appID, address newOwner) returns()
func (_AppStoreStore *AppStoreStoreSession) TransferAppOwnership(appID *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _AppStoreStore.Contract.TransferAppOwnership(&_AppStoreStore.TransactOpts, appID, newOwner)
}

// TransferAppOwnership is a paid mutator transaction binding the contract method 0x6de20cb5.
//
// Solidity: function transferAppOwnership(uint256 appID, address newOwner) returns()
func (_AppStoreStore *AppStoreStoreTransactorSession) TransferAppOwnership(appID *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _AppStoreStore.Contract.TransferAppOwnership(&_AppStoreStore.TransactOpts, appID, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AppStoreStore *AppStoreStoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AppStoreStore.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AppStoreStore *AppStoreStoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AppStoreStore.Contract.TransferOwnership(&_AppStoreStore.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AppStoreStore *AppStoreStoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AppStoreStore.Contract.TransferOwnership(&_AppStoreStore.TransactOpts, newOwner)
}

// AppStoreStoreAppCreatedIterator is returned from FilterAppCreated and is used to iterate over the raw logs and unpacked data for AppCreated events raised by the AppStoreStore contract.
type AppStoreStoreAppCreatedIterator struct {
	Event *AppStoreStoreAppCreated // Event containing the contract specifics and raw log

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
func (it *AppStoreStoreAppCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppStoreStoreAppCreated)
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
		it.Event = new(AppStoreStoreAppCreated)
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
func (it *AppStoreStoreAppCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppStoreStoreAppCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppStoreStoreAppCreated represents a AppCreated event raised by the AppStoreStore contract.
type AppStoreStoreAppCreated struct {
	AppID  *big.Int
	Domain string
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAppCreated is a free log retrieval operation binding the contract event 0x3a2de0162cddb1ff0f94dccf1fd8bb160e3acc5c1f3500094a67d151f42270bb.
//
// Solidity: event AppCreated(uint256 appID, string domain, address owner)
func (_AppStoreStore *AppStoreStoreFilterer) FilterAppCreated(opts *bind.FilterOpts) (*AppStoreStoreAppCreatedIterator, error) {

	logs, sub, err := _AppStoreStore.contract.FilterLogs(opts, "AppCreated")
	if err != nil {
		return nil, err
	}
	return &AppStoreStoreAppCreatedIterator{contract: _AppStoreStore.contract, event: "AppCreated", logs: logs, sub: sub}, nil
}

// WatchAppCreated is a free log subscription operation binding the contract event 0x3a2de0162cddb1ff0f94dccf1fd8bb160e3acc5c1f3500094a67d151f42270bb.
//
// Solidity: event AppCreated(uint256 appID, string domain, address owner)
func (_AppStoreStore *AppStoreStoreFilterer) WatchAppCreated(opts *bind.WatchOpts, sink chan<- *AppStoreStoreAppCreated) (event.Subscription, error) {

	logs, sub, err := _AppStoreStore.contract.WatchLogs(opts, "AppCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppStoreStoreAppCreated)
				if err := _AppStoreStore.contract.UnpackLog(event, "AppCreated", log); err != nil {
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

// ParseAppCreated is a log parse operation binding the contract event 0x3a2de0162cddb1ff0f94dccf1fd8bb160e3acc5c1f3500094a67d151f42270bb.
//
// Solidity: event AppCreated(uint256 appID, string domain, address owner)
func (_AppStoreStore *AppStoreStoreFilterer) ParseAppCreated(log types.Log) (*AppStoreStoreAppCreated, error) {
	event := new(AppStoreStoreAppCreated)
	if err := _AppStoreStore.contract.UnpackLog(event, "AppCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppStoreStoreAppOwnershipTransferredIterator is returned from FilterAppOwnershipTransferred and is used to iterate over the raw logs and unpacked data for AppOwnershipTransferred events raised by the AppStoreStore contract.
type AppStoreStoreAppOwnershipTransferredIterator struct {
	Event *AppStoreStoreAppOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AppStoreStoreAppOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppStoreStoreAppOwnershipTransferred)
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
		it.Event = new(AppStoreStoreAppOwnershipTransferred)
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
func (it *AppStoreStoreAppOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppStoreStoreAppOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppStoreStoreAppOwnershipTransferred represents a AppOwnershipTransferred event raised by the AppStoreStore contract.
type AppStoreStoreAppOwnershipTransferred struct {
	ApplicationID *big.Int
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAppOwnershipTransferred is a free log retrieval operation binding the contract event 0x0c2b80fc73afcaf14c31293f1619512033be1e0707118f21b755b08eea7ef01a.
//
// Solidity: event AppOwnershipTransferred(uint256 applicationID, address previousOwner, address newOwner)
func (_AppStoreStore *AppStoreStoreFilterer) FilterAppOwnershipTransferred(opts *bind.FilterOpts) (*AppStoreStoreAppOwnershipTransferredIterator, error) {

	logs, sub, err := _AppStoreStore.contract.FilterLogs(opts, "AppOwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &AppStoreStoreAppOwnershipTransferredIterator{contract: _AppStoreStore.contract, event: "AppOwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchAppOwnershipTransferred is a free log subscription operation binding the contract event 0x0c2b80fc73afcaf14c31293f1619512033be1e0707118f21b755b08eea7ef01a.
//
// Solidity: event AppOwnershipTransferred(uint256 applicationID, address previousOwner, address newOwner)
func (_AppStoreStore *AppStoreStoreFilterer) WatchAppOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AppStoreStoreAppOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _AppStoreStore.contract.WatchLogs(opts, "AppOwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppStoreStoreAppOwnershipTransferred)
				if err := _AppStoreStore.contract.UnpackLog(event, "AppOwnershipTransferred", log); err != nil {
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

// ParseAppOwnershipTransferred is a log parse operation binding the contract event 0x0c2b80fc73afcaf14c31293f1619512033be1e0707118f21b755b08eea7ef01a.
//
// Solidity: event AppOwnershipTransferred(uint256 applicationID, address previousOwner, address newOwner)
func (_AppStoreStore *AppStoreStoreFilterer) ParseAppOwnershipTransferred(log types.Log) (*AppStoreStoreAppOwnershipTransferred, error) {
	event := new(AppStoreStoreAppOwnershipTransferred)
	if err := _AppStoreStore.contract.UnpackLog(event, "AppOwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppStoreStoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AppStoreStore contract.
type AppStoreStoreOwnershipTransferredIterator struct {
	Event *AppStoreStoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AppStoreStoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppStoreStoreOwnershipTransferred)
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
		it.Event = new(AppStoreStoreOwnershipTransferred)
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
func (it *AppStoreStoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppStoreStoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppStoreStoreOwnershipTransferred represents a OwnershipTransferred event raised by the AppStoreStore contract.
type AppStoreStoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AppStoreStore *AppStoreStoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AppStoreStoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AppStoreStore.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AppStoreStoreOwnershipTransferredIterator{contract: _AppStoreStore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AppStoreStore *AppStoreStoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AppStoreStoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AppStoreStore.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppStoreStoreOwnershipTransferred)
				if err := _AppStoreStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AppStoreStore *AppStoreStoreFilterer) ParseOwnershipTransferred(log types.Log) (*AppStoreStoreOwnershipTransferred, error) {
	event := new(AppStoreStoreOwnershipTransferred)
	if err := _AppStoreStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
