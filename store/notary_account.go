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

// NotaryAccountStoreMetaData contains all meta data concerning the NotaryAccountStore contract.
var NotaryAccountStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractNotaryPublic\",\"name\":\"_notaryPublic\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"notaryPublic\",\"outputs\":[{\"internalType\":\"contractNotaryPublic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"notaryCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NotaryAccountStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use NotaryAccountStoreMetaData.ABI instead.
var NotaryAccountStoreABI = NotaryAccountStoreMetaData.ABI

// NotaryAccountStore is an auto generated Go binding around an Ethereum contract.
type NotaryAccountStore struct {
	NotaryAccountStoreCaller     // Read-only binding to the contract
	NotaryAccountStoreTransactor // Write-only binding to the contract
	NotaryAccountStoreFilterer   // Log filterer for contract events
}

// NotaryAccountStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type NotaryAccountStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryAccountStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NotaryAccountStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryAccountStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NotaryAccountStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryAccountStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NotaryAccountStoreSession struct {
	Contract     *NotaryAccountStore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NotaryAccountStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NotaryAccountStoreCallerSession struct {
	Contract *NotaryAccountStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// NotaryAccountStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NotaryAccountStoreTransactorSession struct {
	Contract     *NotaryAccountStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// NotaryAccountStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type NotaryAccountStoreRaw struct {
	Contract *NotaryAccountStore // Generic contract binding to access the raw methods on
}

// NotaryAccountStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NotaryAccountStoreCallerRaw struct {
	Contract *NotaryAccountStoreCaller // Generic read-only contract binding to access the raw methods on
}

// NotaryAccountStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NotaryAccountStoreTransactorRaw struct {
	Contract *NotaryAccountStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNotaryAccountStore creates a new instance of NotaryAccountStore, bound to a specific deployed contract.
func NewNotaryAccountStore(address common.Address, backend bind.ContractBackend) (*NotaryAccountStore, error) {
	contract, err := bindNotaryAccountStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NotaryAccountStore{NotaryAccountStoreCaller: NotaryAccountStoreCaller{contract: contract}, NotaryAccountStoreTransactor: NotaryAccountStoreTransactor{contract: contract}, NotaryAccountStoreFilterer: NotaryAccountStoreFilterer{contract: contract}}, nil
}

// NewNotaryAccountStoreCaller creates a new read-only instance of NotaryAccountStore, bound to a specific deployed contract.
func NewNotaryAccountStoreCaller(address common.Address, caller bind.ContractCaller) (*NotaryAccountStoreCaller, error) {
	contract, err := bindNotaryAccountStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NotaryAccountStoreCaller{contract: contract}, nil
}

// NewNotaryAccountStoreTransactor creates a new write-only instance of NotaryAccountStore, bound to a specific deployed contract.
func NewNotaryAccountStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*NotaryAccountStoreTransactor, error) {
	contract, err := bindNotaryAccountStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NotaryAccountStoreTransactor{contract: contract}, nil
}

// NewNotaryAccountStoreFilterer creates a new log filterer instance of NotaryAccountStore, bound to a specific deployed contract.
func NewNotaryAccountStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*NotaryAccountStoreFilterer, error) {
	contract, err := bindNotaryAccountStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NotaryAccountStoreFilterer{contract: contract}, nil
}

// bindNotaryAccountStore binds a generic wrapper to an already deployed contract.
func bindNotaryAccountStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NotaryAccountStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NotaryAccountStore *NotaryAccountStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NotaryAccountStore.Contract.NotaryAccountStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NotaryAccountStore *NotaryAccountStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NotaryAccountStore.Contract.NotaryAccountStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NotaryAccountStore *NotaryAccountStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NotaryAccountStore.Contract.NotaryAccountStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NotaryAccountStore *NotaryAccountStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NotaryAccountStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NotaryAccountStore *NotaryAccountStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NotaryAccountStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NotaryAccountStore *NotaryAccountStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NotaryAccountStore.Contract.contract.Transact(opts, method, params...)
}

// NotaryPublic is a free data retrieval call binding the contract method 0xa70ff2ba.
//
// Solidity: function notaryPublic() view returns(address)
func (_NotaryAccountStore *NotaryAccountStoreCaller) NotaryPublic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NotaryAccountStore.contract.Call(opts, &out, "notaryPublic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NotaryPublic is a free data retrieval call binding the contract method 0xa70ff2ba.
//
// Solidity: function notaryPublic() view returns(address)
func (_NotaryAccountStore *NotaryAccountStoreSession) NotaryPublic() (common.Address, error) {
	return _NotaryAccountStore.Contract.NotaryPublic(&_NotaryAccountStore.CallOpts)
}

// NotaryPublic is a free data retrieval call binding the contract method 0xa70ff2ba.
//
// Solidity: function notaryPublic() view returns(address)
func (_NotaryAccountStore *NotaryAccountStoreCallerSession) NotaryPublic() (common.Address, error) {
	return _NotaryAccountStore.Contract.NotaryPublic(&_NotaryAccountStore.CallOpts)
}

// NotaryCall is a paid mutator transaction binding the contract method 0x668a7146.
//
// Solidity: function notaryCall(address target, bytes data) returns()
func (_NotaryAccountStore *NotaryAccountStoreTransactor) NotaryCall(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _NotaryAccountStore.contract.Transact(opts, "notaryCall", target, data)
}

// NotaryCall is a paid mutator transaction binding the contract method 0x668a7146.
//
// Solidity: function notaryCall(address target, bytes data) returns()
func (_NotaryAccountStore *NotaryAccountStoreSession) NotaryCall(target common.Address, data []byte) (*types.Transaction, error) {
	return _NotaryAccountStore.Contract.NotaryCall(&_NotaryAccountStore.TransactOpts, target, data)
}

// NotaryCall is a paid mutator transaction binding the contract method 0x668a7146.
//
// Solidity: function notaryCall(address target, bytes data) returns()
func (_NotaryAccountStore *NotaryAccountStoreTransactorSession) NotaryCall(target common.Address, data []byte) (*types.Transaction, error) {
	return _NotaryAccountStore.Contract.NotaryCall(&_NotaryAccountStore.TransactOpts, target, data)
}
