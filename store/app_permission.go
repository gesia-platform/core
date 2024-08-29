// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
	"errors"
	"fmt"
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

// AppPermissionStoreMetaData contains all meta data concerning the AppPermissionStore contract.
var AppPermissionStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractAppStore\",\"name\":\"_appStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"networkAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"NetworkAccessPermissionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"networkAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isGranted\",\"type\":\"bool\"}],\"name\":\"NetworkAccessPermissionResponsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"appStore\",\"outputs\":[{\"internalType\":\"contractAppStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"getAppID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"networkAccount\",\"type\":\"address\"}],\"name\":\"getNetworkAccessRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"networkAccount\",\"type\":\"address\"}],\"name\":\"getNetworkAccessResponse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"responsed\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"networkAccount\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"requestNetworkAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_networkAccount\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isGranted\",\"type\":\"bool\"}],\"name\":\"responseNetworkAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"networkAccount\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"setNetworkAccessRequestIP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AppPermissionStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use AppPermissionStoreMetaData.ABI instead.
var AppPermissionStoreABI = AppPermissionStoreMetaData.ABI

// AppPermissionStore is an auto generated Go binding around an Ethereum contract.
type AppPermissionStore struct {
	AppPermissionStoreCaller     // Read-only binding to the contract
	AppPermissionStoreTransactor // Write-only binding to the contract
	AppPermissionStoreFilterer   // Log filterer for contract events
}

// AppPermissionStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppPermissionStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppPermissionStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppPermissionStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppPermissionStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppPermissionStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppPermissionStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppPermissionStoreSession struct {
	Contract     *AppPermissionStore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AppPermissionStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppPermissionStoreCallerSession struct {
	Contract *AppPermissionStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AppPermissionStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppPermissionStoreTransactorSession struct {
	Contract     *AppPermissionStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AppPermissionStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppPermissionStoreRaw struct {
	Contract *AppPermissionStore // Generic contract binding to access the raw methods on
}

// AppPermissionStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppPermissionStoreCallerRaw struct {
	Contract *AppPermissionStoreCaller // Generic read-only contract binding to access the raw methods on
}

// AppPermissionStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppPermissionStoreTransactorRaw struct {
	Contract *AppPermissionStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppPermissionStore creates a new instance of AppPermissionStore, bound to a specific deployed contract.
func NewAppPermissionStore(address common.Address, backend bind.ContractBackend) (*AppPermissionStore, error) {
	contract, err := bindAppPermissionStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AppPermissionStore{AppPermissionStoreCaller: AppPermissionStoreCaller{contract: contract}, AppPermissionStoreTransactor: AppPermissionStoreTransactor{contract: contract}, AppPermissionStoreFilterer: AppPermissionStoreFilterer{contract: contract}}, nil
}

// NewAppPermissionStoreCaller creates a new read-only instance of AppPermissionStore, bound to a specific deployed contract.
func NewAppPermissionStoreCaller(address common.Address, caller bind.ContractCaller) (*AppPermissionStoreCaller, error) {
	contract, err := bindAppPermissionStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppPermissionStoreCaller{contract: contract}, nil
}

// NewAppPermissionStoreTransactor creates a new write-only instance of AppPermissionStore, bound to a specific deployed contract.
func NewAppPermissionStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*AppPermissionStoreTransactor, error) {
	contract, err := bindAppPermissionStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppPermissionStoreTransactor{contract: contract}, nil
}

// NewAppPermissionStoreFilterer creates a new log filterer instance of AppPermissionStore, bound to a specific deployed contract.
func NewAppPermissionStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*AppPermissionStoreFilterer, error) {
	contract, err := bindAppPermissionStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppPermissionStoreFilterer{contract: contract}, nil
}

// bindAppPermissionStore binds a generic wrapper to an already deployed contract.
func bindAppPermissionStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AppPermissionStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppPermissionStore *AppPermissionStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppPermissionStore.Contract.AppPermissionStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppPermissionStore *AppPermissionStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppPermissionStore.Contract.AppPermissionStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppPermissionStore *AppPermissionStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppPermissionStore.Contract.AppPermissionStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppPermissionStore *AppPermissionStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppPermissionStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppPermissionStore *AppPermissionStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppPermissionStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppPermissionStore *AppPermissionStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppPermissionStore.Contract.contract.Transact(opts, method, params...)
}

// AppStore is a free data retrieval call binding the contract method 0x72f86f05.
//
// Solidity: function appStore() view returns(address)
func (_AppPermissionStore *AppPermissionStoreCaller) AppStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppPermissionStore.contract.Call(opts, &out, "appStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AppStore is a free data retrieval call binding the contract method 0x72f86f05.
//
// Solidity: function appStore() view returns(address)
func (_AppPermissionStore *AppPermissionStoreSession) AppStore() (common.Address, error) {
	return _AppPermissionStore.Contract.AppStore(&_AppPermissionStore.CallOpts)
}

// AppStore is a free data retrieval call binding the contract method 0x72f86f05.
//
// Solidity: function appStore() view returns(address)
func (_AppPermissionStore *AppPermissionStoreCallerSession) AppStore() (common.Address, error) {
	return _AppPermissionStore.Contract.AppStore(&_AppPermissionStore.CallOpts)
}

// GetAppID is a free data retrieval call binding the contract method 0x6900f1b3.
//
// Solidity: function getAppID(string ip) view returns(bool, uint256)
func (_AppPermissionStore *AppPermissionStoreCaller) GetAppID(opts *bind.CallOpts, ip string) (bool, *big.Int, error) {
	var out []interface{}
	err := _AppPermissionStore.contract.Call(opts, &out, "getAppID", ip)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAppID is a free data retrieval call binding the contract method 0x6900f1b3.
//
// Solidity: function getAppID(string ip) view returns(bool, uint256)
func (_AppPermissionStore *AppPermissionStoreSession) GetAppID(ip string) (bool, *big.Int, error) {
	return _AppPermissionStore.Contract.GetAppID(&_AppPermissionStore.CallOpts, ip)
}

// GetAppID is a free data retrieval call binding the contract method 0x6900f1b3.
//
// Solidity: function getAppID(string ip) view returns(bool, uint256)
func (_AppPermissionStore *AppPermissionStoreCallerSession) GetAppID(ip string) (bool, *big.Int, error) {
	return _AppPermissionStore.Contract.GetAppID(&_AppPermissionStore.CallOpts, ip)
}

// GetNetworkAccessRequest is a free data retrieval call binding the contract method 0xce57b454.
//
// Solidity: function getNetworkAccessRequest(uint256 appID, address networkAccount) view returns(bool, string)
func (_AppPermissionStore *AppPermissionStoreCaller) GetNetworkAccessRequest(opts *bind.CallOpts, appID *big.Int, networkAccount common.Address) (bool, string, error) {
	var out []interface{}
	err := _AppPermissionStore.contract.Call(opts, &out, "getNetworkAccessRequest", appID, networkAccount)

	if err != nil {
		return *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetNetworkAccessRequest is a free data retrieval call binding the contract method 0xce57b454.
//
// Solidity: function getNetworkAccessRequest(uint256 appID, address networkAccount) view returns(bool, string)
func (_AppPermissionStore *AppPermissionStoreSession) GetNetworkAccessRequest(appID *big.Int, networkAccount common.Address) (bool, string, error) {
	return _AppPermissionStore.Contract.GetNetworkAccessRequest(&_AppPermissionStore.CallOpts, appID, networkAccount)
}

// GetNetworkAccessRequest is a free data retrieval call binding the contract method 0xce57b454.
//
// Solidity: function getNetworkAccessRequest(uint256 appID, address networkAccount) view returns(bool, string)
func (_AppPermissionStore *AppPermissionStoreCallerSession) GetNetworkAccessRequest(appID *big.Int, networkAccount common.Address) (bool, string, error) {
	return _AppPermissionStore.Contract.GetNetworkAccessRequest(&_AppPermissionStore.CallOpts, appID, networkAccount)
}

// GetNetworkAccessResponse is a free data retrieval call binding the contract method 0xb50926e0.
//
// Solidity: function getNetworkAccessResponse(uint256 appID, address networkAccount) view returns(bool responsed, bytes, bool)
func (_AppPermissionStore *AppPermissionStoreCaller) GetNetworkAccessResponse(opts *bind.CallOpts, appID *big.Int, networkAccount common.Address) (bool, []byte, bool, error) {
	var out []interface{}
	err := _AppPermissionStore.contract.Call(opts, &out, "getNetworkAccessResponse", appID, networkAccount)

	if err != nil {
		return *new(bool), *new([]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	fmt.Print("Test : ",out0,out1,out2)

	return out0, out1, out2, err

}

// GetNetworkAccessResponse is a free data retrieval call binding the contract method 0xb50926e0.
//
// Solidity: function getNetworkAccessResponse(uint256 appID, address networkAccount) view returns(bool responsed, bytes, bool)
func (_AppPermissionStore *AppPermissionStoreSession) GetNetworkAccessResponse(appID *big.Int, networkAccount common.Address) (bool, []byte, bool, error) {
	return _AppPermissionStore.Contract.GetNetworkAccessResponse(&_AppPermissionStore.CallOpts, appID, networkAccount)
}

// GetNetworkAccessResponse is a free data retrieval call binding the contract method 0xb50926e0.
//
// Solidity: function getNetworkAccessResponse(uint256 appID, address networkAccount) view returns(bool responsed, bytes, bool)
func (_AppPermissionStore *AppPermissionStoreCallerSession) GetNetworkAccessResponse(appID *big.Int, networkAccount common.Address) (bool, []byte, bool, error) {
	return _AppPermissionStore.Contract.GetNetworkAccessResponse(&_AppPermissionStore.CallOpts, appID, networkAccount)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AppPermissionStore *AppPermissionStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppPermissionStore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AppPermissionStore *AppPermissionStoreSession) Owner() (common.Address, error) {
	return _AppPermissionStore.Contract.Owner(&_AppPermissionStore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AppPermissionStore *AppPermissionStoreCallerSession) Owner() (common.Address, error) {
	return _AppPermissionStore.Contract.Owner(&_AppPermissionStore.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppPermissionStore *AppPermissionStoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppPermissionStore.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppPermissionStore *AppPermissionStoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _AppPermissionStore.Contract.RenounceOwnership(&_AppPermissionStore.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppPermissionStore *AppPermissionStoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AppPermissionStore.Contract.RenounceOwnership(&_AppPermissionStore.TransactOpts)
}

// RequestNetworkAccess is a paid mutator transaction binding the contract method 0x7d39cad8.
//
// Solidity: function requestNetworkAccess(uint256 appID, address networkAccount, string ip) returns()
func (_AppPermissionStore *AppPermissionStoreTransactor) RequestNetworkAccess(opts *bind.TransactOpts, appID *big.Int, networkAccount common.Address, ip string) (*types.Transaction, error) {
	return _AppPermissionStore.contract.Transact(opts, "requestNetworkAccess", appID, networkAccount, ip)
}

// RequestNetworkAccess is a paid mutator transaction binding the contract method 0x7d39cad8.
//
// Solidity: function requestNetworkAccess(uint256 appID, address networkAccount, string ip) returns()
func (_AppPermissionStore *AppPermissionStoreSession) RequestNetworkAccess(appID *big.Int, networkAccount common.Address, ip string) (*types.Transaction, error) {
	return _AppPermissionStore.Contract.RequestNetworkAccess(&_AppPermissionStore.TransactOpts, appID, networkAccount, ip)
}

// RequestNetworkAccess is a paid mutator transaction binding the contract method 0x7d39cad8.
//
// Solidity: function requestNetworkAccess(uint256 appID, address networkAccount, string ip) returns()
func (_AppPermissionStore *AppPermissionStoreTransactorSession) RequestNetworkAccess(appID *big.Int, networkAccount common.Address, ip string) (*types.Transaction, error) {
	return _AppPermissionStore.Contract.RequestNetworkAccess(&_AppPermissionStore.TransactOpts, appID, networkAccount, ip)
}

// ResponseNetworkAccess is a paid mutator transaction binding the contract method 0x77a03f1e.
//
// Solidity: function responseNetworkAccess(uint256 appID, address _networkAccount, bytes signature, bool isGranted) returns()
func (_AppPermissionStore *AppPermissionStoreTransactor) ResponseNetworkAccess(opts *bind.TransactOpts, appID *big.Int, _networkAccount common.Address, signature []byte, isGranted bool) (*types.Transaction, error) {
	return _AppPermissionStore.contract.Transact(opts, "responseNetworkAccess", appID, _networkAccount, signature, isGranted)
}

// ResponseNetworkAccess is a paid mutator transaction binding the contract method 0x77a03f1e.
//
// Solidity: function responseNetworkAccess(uint256 appID, address _networkAccount, bytes signature, bool isGranted) returns()
func (_AppPermissionStore *AppPermissionStoreSession) ResponseNetworkAccess(appID *big.Int, _networkAccount common.Address, signature []byte, isGranted bool) (*types.Transaction, error) {
	return _AppPermissionStore.Contract.ResponseNetworkAccess(&_AppPermissionStore.TransactOpts, appID, _networkAccount, signature, isGranted)
}

// ResponseNetworkAccess is a paid mutator transaction binding the contract method 0x77a03f1e.
//
// Solidity: function responseNetworkAccess(uint256 appID, address _networkAccount, bytes signature, bool isGranted) returns()
func (_AppPermissionStore *AppPermissionStoreTransactorSession) ResponseNetworkAccess(appID *big.Int, _networkAccount common.Address, signature []byte, isGranted bool) (*types.Transaction, error) {
	return _AppPermissionStore.Contract.ResponseNetworkAccess(&_AppPermissionStore.TransactOpts, appID, _networkAccount, signature, isGranted)
}

// SetNetworkAccessRequestIP is a paid mutator transaction binding the contract method 0x548734eb.
//
// Solidity: function setNetworkAccessRequestIP(uint256 appID, address networkAccount, string ip) returns()
func (_AppPermissionStore *AppPermissionStoreTransactor) SetNetworkAccessRequestIP(opts *bind.TransactOpts, appID *big.Int, networkAccount common.Address, ip string) (*types.Transaction, error) {
	return _AppPermissionStore.contract.Transact(opts, "setNetworkAccessRequestIP", appID, networkAccount, ip)
}

// SetNetworkAccessRequestIP is a paid mutator transaction binding the contract method 0x548734eb.
//
// Solidity: function setNetworkAccessRequestIP(uint256 appID, address networkAccount, string ip) returns()
func (_AppPermissionStore *AppPermissionStoreSession) SetNetworkAccessRequestIP(appID *big.Int, networkAccount common.Address, ip string) (*types.Transaction, error) {
	return _AppPermissionStore.Contract.SetNetworkAccessRequestIP(&_AppPermissionStore.TransactOpts, appID, networkAccount, ip)
}

// SetNetworkAccessRequestIP is a paid mutator transaction binding the contract method 0x548734eb.
//
// Solidity: function setNetworkAccessRequestIP(uint256 appID, address networkAccount, string ip) returns()
func (_AppPermissionStore *AppPermissionStoreTransactorSession) SetNetworkAccessRequestIP(appID *big.Int, networkAccount common.Address, ip string) (*types.Transaction, error) {
	return _AppPermissionStore.Contract.SetNetworkAccessRequestIP(&_AppPermissionStore.TransactOpts, appID, networkAccount, ip)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AppPermissionStore *AppPermissionStoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AppPermissionStore.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AppPermissionStore *AppPermissionStoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AppPermissionStore.Contract.TransferOwnership(&_AppPermissionStore.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AppPermissionStore *AppPermissionStoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AppPermissionStore.Contract.TransferOwnership(&_AppPermissionStore.TransactOpts, newOwner)
}

// AppPermissionStoreNetworkAccessPermissionRequestedIterator is returned from FilterNetworkAccessPermissionRequested and is used to iterate over the raw logs and unpacked data for NetworkAccessPermissionRequested events raised by the AppPermissionStore contract.
type AppPermissionStoreNetworkAccessPermissionRequestedIterator struct {
	Event *AppPermissionStoreNetworkAccessPermissionRequested // Event containing the contract specifics and raw log

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
func (it *AppPermissionStoreNetworkAccessPermissionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppPermissionStoreNetworkAccessPermissionRequested)
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
		it.Event = new(AppPermissionStoreNetworkAccessPermissionRequested)
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
func (it *AppPermissionStoreNetworkAccessPermissionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppPermissionStoreNetworkAccessPermissionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppPermissionStoreNetworkAccessPermissionRequested represents a NetworkAccessPermissionRequested event raised by the AppPermissionStore contract.
type AppPermissionStoreNetworkAccessPermissionRequested struct {
	AppID          *big.Int
	NetworkAccount common.Address
	Ip             string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNetworkAccessPermissionRequested is a free log retrieval operation binding the contract event 0xccfc52f65ba4e391ef7b8ac94bee5f9fd3bfe05a126052d1ef0b9c864a807cf8.
//
// Solidity: event NetworkAccessPermissionRequested(uint256 appID, address networkAccount, string ip)
func (_AppPermissionStore *AppPermissionStoreFilterer) FilterNetworkAccessPermissionRequested(opts *bind.FilterOpts) (*AppPermissionStoreNetworkAccessPermissionRequestedIterator, error) {

	logs, sub, err := _AppPermissionStore.contract.FilterLogs(opts, "NetworkAccessPermissionRequested")
	if err != nil {
		return nil, err
	}
	return &AppPermissionStoreNetworkAccessPermissionRequestedIterator{contract: _AppPermissionStore.contract, event: "NetworkAccessPermissionRequested", logs: logs, sub: sub}, nil
}

// WatchNetworkAccessPermissionRequested is a free log subscription operation binding the contract event 0xccfc52f65ba4e391ef7b8ac94bee5f9fd3bfe05a126052d1ef0b9c864a807cf8.
//
// Solidity: event NetworkAccessPermissionRequested(uint256 appID, address networkAccount, string ip)
func (_AppPermissionStore *AppPermissionStoreFilterer) WatchNetworkAccessPermissionRequested(opts *bind.WatchOpts, sink chan<- *AppPermissionStoreNetworkAccessPermissionRequested) (event.Subscription, error) {

	logs, sub, err := _AppPermissionStore.contract.WatchLogs(opts, "NetworkAccessPermissionRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppPermissionStoreNetworkAccessPermissionRequested)
				if err := _AppPermissionStore.contract.UnpackLog(event, "NetworkAccessPermissionRequested", log); err != nil {
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

// ParseNetworkAccessPermissionRequested is a log parse operation binding the contract event 0xccfc52f65ba4e391ef7b8ac94bee5f9fd3bfe05a126052d1ef0b9c864a807cf8.
//
// Solidity: event NetworkAccessPermissionRequested(uint256 appID, address networkAccount, string ip)
func (_AppPermissionStore *AppPermissionStoreFilterer) ParseNetworkAccessPermissionRequested(log types.Log) (*AppPermissionStoreNetworkAccessPermissionRequested, error) {
	event := new(AppPermissionStoreNetworkAccessPermissionRequested)
	if err := _AppPermissionStore.contract.UnpackLog(event, "NetworkAccessPermissionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppPermissionStoreNetworkAccessPermissionResponsedIterator is returned from FilterNetworkAccessPermissionResponsed and is used to iterate over the raw logs and unpacked data for NetworkAccessPermissionResponsed events raised by the AppPermissionStore contract.
type AppPermissionStoreNetworkAccessPermissionResponsedIterator struct {
	Event *AppPermissionStoreNetworkAccessPermissionResponsed // Event containing the contract specifics and raw log

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
func (it *AppPermissionStoreNetworkAccessPermissionResponsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppPermissionStoreNetworkAccessPermissionResponsed)
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
		it.Event = new(AppPermissionStoreNetworkAccessPermissionResponsed)
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
func (it *AppPermissionStoreNetworkAccessPermissionResponsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppPermissionStoreNetworkAccessPermissionResponsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppPermissionStoreNetworkAccessPermissionResponsed represents a NetworkAccessPermissionResponsed event raised by the AppPermissionStore contract.
type AppPermissionStoreNetworkAccessPermissionResponsed struct {
	AppID          *big.Int
	NetworkAccount common.Address
	Signature      []byte
	IsGranted      bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNetworkAccessPermissionResponsed is a free log retrieval operation binding the contract event 0xbe88eb6b46107091ccab269e16bee03c986d5ed4d3f36f48e59fa668375b10ab.
//
// Solidity: event NetworkAccessPermissionResponsed(uint256 appID, address networkAccount, bytes signature, bool isGranted)
func (_AppPermissionStore *AppPermissionStoreFilterer) FilterNetworkAccessPermissionResponsed(opts *bind.FilterOpts) (*AppPermissionStoreNetworkAccessPermissionResponsedIterator, error) {

	logs, sub, err := _AppPermissionStore.contract.FilterLogs(opts, "NetworkAccessPermissionResponsed")
	if err != nil {
		return nil, err
	}
	return &AppPermissionStoreNetworkAccessPermissionResponsedIterator{contract: _AppPermissionStore.contract, event: "NetworkAccessPermissionResponsed", logs: logs, sub: sub}, nil
}

// WatchNetworkAccessPermissionResponsed is a free log subscription operation binding the contract event 0xbe88eb6b46107091ccab269e16bee03c986d5ed4d3f36f48e59fa668375b10ab.
//
// Solidity: event NetworkAccessPermissionResponsed(uint256 appID, address networkAccount, bytes signature, bool isGranted)
func (_AppPermissionStore *AppPermissionStoreFilterer) WatchNetworkAccessPermissionResponsed(opts *bind.WatchOpts, sink chan<- *AppPermissionStoreNetworkAccessPermissionResponsed) (event.Subscription, error) {

	logs, sub, err := _AppPermissionStore.contract.WatchLogs(opts, "NetworkAccessPermissionResponsed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppPermissionStoreNetworkAccessPermissionResponsed)
				if err := _AppPermissionStore.contract.UnpackLog(event, "NetworkAccessPermissionResponsed", log); err != nil {
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

// ParseNetworkAccessPermissionResponsed is a log parse operation binding the contract event 0xbe88eb6b46107091ccab269e16bee03c986d5ed4d3f36f48e59fa668375b10ab.
//
// Solidity: event NetworkAccessPermissionResponsed(uint256 appID, address networkAccount, bytes signature, bool isGranted)
func (_AppPermissionStore *AppPermissionStoreFilterer) ParseNetworkAccessPermissionResponsed(log types.Log) (*AppPermissionStoreNetworkAccessPermissionResponsed, error) {
	event := new(AppPermissionStoreNetworkAccessPermissionResponsed)
	if err := _AppPermissionStore.contract.UnpackLog(event, "NetworkAccessPermissionResponsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppPermissionStoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AppPermissionStore contract.
type AppPermissionStoreOwnershipTransferredIterator struct {
	Event *AppPermissionStoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AppPermissionStoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppPermissionStoreOwnershipTransferred)
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
		it.Event = new(AppPermissionStoreOwnershipTransferred)
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
func (it *AppPermissionStoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppPermissionStoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppPermissionStoreOwnershipTransferred represents a OwnershipTransferred event raised by the AppPermissionStore contract.
type AppPermissionStoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AppPermissionStore *AppPermissionStoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AppPermissionStoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AppPermissionStore.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AppPermissionStoreOwnershipTransferredIterator{contract: _AppPermissionStore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AppPermissionStore *AppPermissionStoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AppPermissionStoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AppPermissionStore.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppPermissionStoreOwnershipTransferred)
				if err := _AppPermissionStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AppPermissionStore *AppPermissionStoreFilterer) ParseOwnershipTransferred(log types.Log) (*AppPermissionStoreOwnershipTransferred, error) {
	event := new(AppPermissionStoreOwnershipTransferred)
	if err := _AppPermissionStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
