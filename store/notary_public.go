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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"prefix\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signatrue\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"Notarized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notaryAccount\",\"type\":\"address\"}],\"name\":\"NotaryAccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"}],\"name\":\"createNotaryAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"prefix\",\"type\":\"bytes1\"},{\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatrue\",\"type\":\"bytes\"}],\"name\":\"notarize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"membership\",\"type\":\"bool\"}],\"name\":\"setMembership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"getPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"prefix\",\"type\":\"bytes1\"},{\"internalType\":\"uint256\",\"name\":\"appID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"getNotarization\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAppIDByNotaryAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetAppIDByNotaryAccount is a free data retrieval call binding the contract method 0xbf73a847.
//
// Solidity: function getAppIDByNotaryAccount(address account) view returns(bool, uint256)
func (_NotaryPublicStore *NotaryPublicStoreCaller) GetAppIDByNotaryAccount(opts *bind.CallOpts, account common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _NotaryPublicStore.contract.Call(opts, &out, "getAppIDByNotaryAccount", account)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAppIDByNotaryAccount is a free data retrieval call binding the contract method 0xbf73a847.
//
// Solidity: function getAppIDByNotaryAccount(address account) view returns(bool, uint256)
func (_NotaryPublicStore *NotaryPublicStoreSession) GetAppIDByNotaryAccount(account common.Address) (bool, *big.Int, error) {
	return _NotaryPublicStore.Contract.GetAppIDByNotaryAccount(&_NotaryPublicStore.CallOpts, account)
}

// GetAppIDByNotaryAccount is a free data retrieval call binding the contract method 0xbf73a847.
//
// Solidity: function getAppIDByNotaryAccount(address account) view returns(bool, uint256)
func (_NotaryPublicStore *NotaryPublicStoreCallerSession) GetAppIDByNotaryAccount(account common.Address) (bool, *big.Int, error) {
	return _NotaryPublicStore.Contract.GetAppIDByNotaryAccount(&_NotaryPublicStore.CallOpts, account)
}

// GetNotarization is a free data retrieval call binding the contract method 0xedac6563.
//
// Solidity: function getNotarization(bytes1 prefix, uint256 appID, address notary) view returns(bytes, bytes)
func (_NotaryPublicStore *NotaryPublicStoreCaller) GetNotarization(opts *bind.CallOpts, prefix [1]byte, appID *big.Int, notary common.Address) ([]byte, []byte, error) {
	var out []interface{}
	err := _NotaryPublicStore.contract.Call(opts, &out, "getNotarization", prefix, appID, notary)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetNotarization is a free data retrieval call binding the contract method 0xedac6563.
//
// Solidity: function getNotarization(bytes1 prefix, uint256 appID, address notary) view returns(bytes, bytes)
func (_NotaryPublicStore *NotaryPublicStoreSession) GetNotarization(prefix [1]byte, appID *big.Int, notary common.Address) ([]byte, []byte, error) {
	return _NotaryPublicStore.Contract.GetNotarization(&_NotaryPublicStore.CallOpts, prefix, appID, notary)
}

// GetNotarization is a free data retrieval call binding the contract method 0xedac6563.
//
// Solidity: function getNotarization(bytes1 prefix, uint256 appID, address notary) view returns(bytes, bytes)
func (_NotaryPublicStore *NotaryPublicStoreCallerSession) GetNotarization(prefix [1]byte, appID *big.Int, notary common.Address) ([]byte, []byte, error) {
	return _NotaryPublicStore.Contract.GetNotarization(&_NotaryPublicStore.CallOpts, prefix, appID, notary)
}

// GetPubkey is a free data retrieval call binding the contract method 0x0f260ca0.
//
// Solidity: function getPubkey(address notary) view returns(bytes)
func (_NotaryPublicStore *NotaryPublicStoreCaller) GetPubkey(opts *bind.CallOpts, notary common.Address) ([]byte, error) {
	var out []interface{}
	err := _NotaryPublicStore.contract.Call(opts, &out, "getPubkey", notary)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPubkey is a free data retrieval call binding the contract method 0x0f260ca0.
//
// Solidity: function getPubkey(address notary) view returns(bytes)
func (_NotaryPublicStore *NotaryPublicStoreSession) GetPubkey(notary common.Address) ([]byte, error) {
	return _NotaryPublicStore.Contract.GetPubkey(&_NotaryPublicStore.CallOpts, notary)
}

// GetPubkey is a free data retrieval call binding the contract method 0x0f260ca0.
//
// Solidity: function getPubkey(address notary) view returns(bytes)
func (_NotaryPublicStore *NotaryPublicStoreCallerSession) GetPubkey(notary common.Address) ([]byte, error) {
	return _NotaryPublicStore.Contract.GetPubkey(&_NotaryPublicStore.CallOpts, notary)
}

// HasMembership is a free data retrieval call binding the contract method 0xa88db6ad.
//
// Solidity: function hasMembership(address account) view returns(bool)
func (_NotaryPublicStore *NotaryPublicStoreCaller) HasMembership(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _NotaryPublicStore.contract.Call(opts, &out, "hasMembership", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMembership is a free data retrieval call binding the contract method 0xa88db6ad.
//
// Solidity: function hasMembership(address account) view returns(bool)
func (_NotaryPublicStore *NotaryPublicStoreSession) HasMembership(account common.Address) (bool, error) {
	return _NotaryPublicStore.Contract.HasMembership(&_NotaryPublicStore.CallOpts, account)
}

// HasMembership is a free data retrieval call binding the contract method 0xa88db6ad.
//
// Solidity: function hasMembership(address account) view returns(bool)
func (_NotaryPublicStore *NotaryPublicStoreCallerSession) HasMembership(account common.Address) (bool, error) {
	return _NotaryPublicStore.Contract.HasMembership(&_NotaryPublicStore.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NotaryPublicStore *NotaryPublicStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NotaryPublicStore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NotaryPublicStore *NotaryPublicStoreSession) Owner() (common.Address, error) {
	return _NotaryPublicStore.Contract.Owner(&_NotaryPublicStore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NotaryPublicStore *NotaryPublicStoreCallerSession) Owner() (common.Address, error) {
	return _NotaryPublicStore.Contract.Owner(&_NotaryPublicStore.CallOpts)
}

// CreateNotaryAccount is a paid mutator transaction binding the contract method 0xb741bdef.
//
// Solidity: function createNotaryAccount(uint256 appID) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactor) CreateNotaryAccount(opts *bind.TransactOpts, appID *big.Int) (*types.Transaction, error) {
	return _NotaryPublicStore.contract.Transact(opts, "createNotaryAccount", appID)
}

// CreateNotaryAccount is a paid mutator transaction binding the contract method 0xb741bdef.
//
// Solidity: function createNotaryAccount(uint256 appID) returns()
func (_NotaryPublicStore *NotaryPublicStoreSession) CreateNotaryAccount(appID *big.Int) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.CreateNotaryAccount(&_NotaryPublicStore.TransactOpts, appID)
}

// CreateNotaryAccount is a paid mutator transaction binding the contract method 0xb741bdef.
//
// Solidity: function createNotaryAccount(uint256 appID) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactorSession) CreateNotaryAccount(appID *big.Int) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.CreateNotaryAccount(&_NotaryPublicStore.TransactOpts, appID)
}

// Notarize is a paid mutator transaction binding the contract method 0x7e004409.
//
// Solidity: function notarize(bytes1 prefix, uint256 appID, bytes signatrue) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactor) Notarize(opts *bind.TransactOpts, prefix [1]byte, appID *big.Int, signatrue []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.contract.Transact(opts, "notarize", prefix, appID, signatrue)
}

// Notarize is a paid mutator transaction binding the contract method 0x7e004409.
//
// Solidity: function notarize(bytes1 prefix, uint256 appID, bytes signatrue) returns()
func (_NotaryPublicStore *NotaryPublicStoreSession) Notarize(prefix [1]byte, appID *big.Int, signatrue []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.Notarize(&_NotaryPublicStore.TransactOpts, prefix, appID, signatrue)
}

// Notarize is a paid mutator transaction binding the contract method 0x7e004409.
//
// Solidity: function notarize(bytes1 prefix, uint256 appID, bytes signatrue) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactorSession) Notarize(prefix [1]byte, appID *big.Int, signatrue []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.Notarize(&_NotaryPublicStore.TransactOpts, prefix, appID, signatrue)
}

// Register is a paid mutator transaction binding the contract method 0x82fbdc9c.
//
// Solidity: function register(bytes pubkey) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactor) Register(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.contract.Transact(opts, "register", pubkey)
}

// Register is a paid mutator transaction binding the contract method 0x82fbdc9c.
//
// Solidity: function register(bytes pubkey) returns()
func (_NotaryPublicStore *NotaryPublicStoreSession) Register(pubkey []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.Register(&_NotaryPublicStore.TransactOpts, pubkey)
}

// Register is a paid mutator transaction binding the contract method 0x82fbdc9c.
//
// Solidity: function register(bytes pubkey) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactorSession) Register(pubkey []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.Register(&_NotaryPublicStore.TransactOpts, pubkey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NotaryPublicStore.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NotaryPublicStore *NotaryPublicStoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.RenounceOwnership(&_NotaryPublicStore.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.RenounceOwnership(&_NotaryPublicStore.TransactOpts)
}

// SetMembership is a paid mutator transaction binding the contract method 0x01746484.
//
// Solidity: function setMembership(address account, bool membership) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactor) SetMembership(opts *bind.TransactOpts, account common.Address, membership bool) (*types.Transaction, error) {
	return _NotaryPublicStore.contract.Transact(opts, "setMembership", account, membership)
}

// SetMembership is a paid mutator transaction binding the contract method 0x01746484.
//
// Solidity: function setMembership(address account, bool membership) returns()
func (_NotaryPublicStore *NotaryPublicStoreSession) SetMembership(account common.Address, membership bool) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.SetMembership(&_NotaryPublicStore.TransactOpts, account, membership)
}

// SetMembership is a paid mutator transaction binding the contract method 0x01746484.
//
// Solidity: function setMembership(address account, bool membership) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactorSession) SetMembership(account common.Address, membership bool) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.SetMembership(&_NotaryPublicStore.TransactOpts, account, membership)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NotaryPublicStore.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NotaryPublicStore *NotaryPublicStoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.TransferOwnership(&_NotaryPublicStore.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.TransferOwnership(&_NotaryPublicStore.TransactOpts, newOwner)
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
	AppID     *big.Int
	Signatrue []byte
	Notary    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNotarized is a free log retrieval operation binding the contract event 0x05e59c60c32fb6565e9a30277c830760329a8da45b5eb3b0b995f70fb3c9682e.
//
// Solidity: event Notarized(bytes1 prefix, uint256 appID, bytes signatrue, address notary)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) FilterNotarized(opts *bind.FilterOpts) (*NotaryPublicStoreNotarizedIterator, error) {

	logs, sub, err := _NotaryPublicStore.contract.FilterLogs(opts, "Notarized")
	if err != nil {
		return nil, err
	}
	return &NotaryPublicStoreNotarizedIterator{contract: _NotaryPublicStore.contract, event: "Notarized", logs: logs, sub: sub}, nil
}

// WatchNotarized is a free log subscription operation binding the contract event 0x05e59c60c32fb6565e9a30277c830760329a8da45b5eb3b0b995f70fb3c9682e.
//
// Solidity: event Notarized(bytes1 prefix, uint256 appID, bytes signatrue, address notary)
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

// ParseNotarized is a log parse operation binding the contract event 0x05e59c60c32fb6565e9a30277c830760329a8da45b5eb3b0b995f70fb3c9682e.
//
// Solidity: event Notarized(bytes1 prefix, uint256 appID, bytes signatrue, address notary)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) ParseNotarized(log types.Log) (*NotaryPublicStoreNotarized, error) {
	event := new(NotaryPublicStoreNotarized)
	if err := _NotaryPublicStore.contract.UnpackLog(event, "Notarized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryPublicStoreNotaryAccountCreatedIterator is returned from FilterNotaryAccountCreated and is used to iterate over the raw logs and unpacked data for NotaryAccountCreated events raised by the NotaryPublicStore contract.
type NotaryPublicStoreNotaryAccountCreatedIterator struct {
	Event *NotaryPublicStoreNotaryAccountCreated // Event containing the contract specifics and raw log

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
func (it *NotaryPublicStoreNotaryAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryPublicStoreNotaryAccountCreated)
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
		it.Event = new(NotaryPublicStoreNotaryAccountCreated)
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
func (it *NotaryPublicStoreNotaryAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryPublicStoreNotaryAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryPublicStoreNotaryAccountCreated represents a NotaryAccountCreated event raised by the NotaryPublicStore contract.
type NotaryPublicStoreNotaryAccountCreated struct {
	AppID         *big.Int
	NotaryAccount common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNotaryAccountCreated is a free log retrieval operation binding the contract event 0xa37d3f614521f572029dcaaecfd63b783ded611520b64a15e0eab5e9ec73cbcc.
//
// Solidity: event NotaryAccountCreated(uint256 appID, address notaryAccount)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) FilterNotaryAccountCreated(opts *bind.FilterOpts) (*NotaryPublicStoreNotaryAccountCreatedIterator, error) {

	logs, sub, err := _NotaryPublicStore.contract.FilterLogs(opts, "NotaryAccountCreated")
	if err != nil {
		return nil, err
	}
	return &NotaryPublicStoreNotaryAccountCreatedIterator{contract: _NotaryPublicStore.contract, event: "NotaryAccountCreated", logs: logs, sub: sub}, nil
}

// WatchNotaryAccountCreated is a free log subscription operation binding the contract event 0xa37d3f614521f572029dcaaecfd63b783ded611520b64a15e0eab5e9ec73cbcc.
//
// Solidity: event NotaryAccountCreated(uint256 appID, address notaryAccount)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) WatchNotaryAccountCreated(opts *bind.WatchOpts, sink chan<- *NotaryPublicStoreNotaryAccountCreated) (event.Subscription, error) {

	logs, sub, err := _NotaryPublicStore.contract.WatchLogs(opts, "NotaryAccountCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryPublicStoreNotaryAccountCreated)
				if err := _NotaryPublicStore.contract.UnpackLog(event, "NotaryAccountCreated", log); err != nil {
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

// ParseNotaryAccountCreated is a log parse operation binding the contract event 0xa37d3f614521f572029dcaaecfd63b783ded611520b64a15e0eab5e9ec73cbcc.
//
// Solidity: event NotaryAccountCreated(uint256 appID, address notaryAccount)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) ParseNotaryAccountCreated(log types.Log) (*NotaryPublicStoreNotaryAccountCreated, error) {
	event := new(NotaryPublicStoreNotaryAccountCreated)
	if err := _NotaryPublicStore.contract.UnpackLog(event, "NotaryAccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryPublicStoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NotaryPublicStore contract.
type NotaryPublicStoreOwnershipTransferredIterator struct {
	Event *NotaryPublicStoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NotaryPublicStoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryPublicStoreOwnershipTransferred)
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
		it.Event = new(NotaryPublicStoreOwnershipTransferred)
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
func (it *NotaryPublicStoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryPublicStoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryPublicStoreOwnershipTransferred represents a OwnershipTransferred event raised by the NotaryPublicStore contract.
type NotaryPublicStoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NotaryPublicStoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NotaryPublicStore.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NotaryPublicStoreOwnershipTransferredIterator{contract: _NotaryPublicStore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NotaryPublicStoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NotaryPublicStore.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryPublicStoreOwnershipTransferred)
				if err := _NotaryPublicStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NotaryPublicStore *NotaryPublicStoreFilterer) ParseOwnershipTransferred(log types.Log) (*NotaryPublicStoreOwnershipTransferred, error) {
	event := new(NotaryPublicStoreOwnershipTransferred)
	if err := _NotaryPublicStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
