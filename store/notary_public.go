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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"applicationID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ApplicationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"applicationID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"ApplicationNotarizationRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"applicationID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"ApplicationNotarized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"applicationID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"ApplicationOwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"createApplication\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"applicationID\",\"type\":\"uint256\"}],\"name\":\"getApplicationDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"applicationID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"getApplicationNotarizationDetails\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"applicationID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"notarizeApplication\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"applicationID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"revokeApplicationNotarization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"applicationID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferApplicationOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetApplicationDetails is a free data retrieval call binding the contract method 0xd4aa7380.
//
// Solidity: function getApplicationDetails(uint256 applicationID) view returns(string, address)
func (_NotaryPublicStore *NotaryPublicStoreCaller) GetApplicationDetails(opts *bind.CallOpts, applicationID *big.Int) (string, common.Address, error) {
	var out []interface{}
	err := _NotaryPublicStore.contract.Call(opts, &out, "getApplicationDetails", applicationID)

	if err != nil {
		return *new(string), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetApplicationDetails is a free data retrieval call binding the contract method 0xd4aa7380.
//
// Solidity: function getApplicationDetails(uint256 applicationID) view returns(string, address)
func (_NotaryPublicStore *NotaryPublicStoreSession) GetApplicationDetails(applicationID *big.Int) (string, common.Address, error) {
	return _NotaryPublicStore.Contract.GetApplicationDetails(&_NotaryPublicStore.CallOpts, applicationID)
}

// GetApplicationDetails is a free data retrieval call binding the contract method 0xd4aa7380.
//
// Solidity: function getApplicationDetails(uint256 applicationID) view returns(string, address)
func (_NotaryPublicStore *NotaryPublicStoreCallerSession) GetApplicationDetails(applicationID *big.Int) (string, common.Address, error) {
	return _NotaryPublicStore.Contract.GetApplicationDetails(&_NotaryPublicStore.CallOpts, applicationID)
}

// GetApplicationNotarizationDetails is a free data retrieval call binding the contract method 0xcf48b377.
//
// Solidity: function getApplicationNotarizationDetails(uint256 applicationID, string network, address notary) view returns(bool, bytes, string)
func (_NotaryPublicStore *NotaryPublicStoreCaller) GetApplicationNotarizationDetails(opts *bind.CallOpts, applicationID *big.Int, network string, notary common.Address) (bool, []byte, string, error) {
	var out []interface{}
	err := _NotaryPublicStore.contract.Call(opts, &out, "getApplicationNotarizationDetails", applicationID, network, notary)

	if err != nil {
		return *new(bool), *new([]byte), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetApplicationNotarizationDetails is a free data retrieval call binding the contract method 0xcf48b377.
//
// Solidity: function getApplicationNotarizationDetails(uint256 applicationID, string network, address notary) view returns(bool, bytes, string)
func (_NotaryPublicStore *NotaryPublicStoreSession) GetApplicationNotarizationDetails(applicationID *big.Int, network string, notary common.Address) (bool, []byte, string, error) {
	return _NotaryPublicStore.Contract.GetApplicationNotarizationDetails(&_NotaryPublicStore.CallOpts, applicationID, network, notary)
}

// GetApplicationNotarizationDetails is a free data retrieval call binding the contract method 0xcf48b377.
//
// Solidity: function getApplicationNotarizationDetails(uint256 applicationID, string network, address notary) view returns(bool, bytes, string)
func (_NotaryPublicStore *NotaryPublicStoreCallerSession) GetApplicationNotarizationDetails(applicationID *big.Int, network string, notary common.Address) (bool, []byte, string, error) {
	return _NotaryPublicStore.Contract.GetApplicationNotarizationDetails(&_NotaryPublicStore.CallOpts, applicationID, network, notary)
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

// CreateApplication is a paid mutator transaction binding the contract method 0xb89d1d16.
//
// Solidity: function createApplication(string domain, address account) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactor) CreateApplication(opts *bind.TransactOpts, domain string, account common.Address) (*types.Transaction, error) {
	return _NotaryPublicStore.contract.Transact(opts, "createApplication", domain, account)
}

// CreateApplication is a paid mutator transaction binding the contract method 0xb89d1d16.
//
// Solidity: function createApplication(string domain, address account) returns()
func (_NotaryPublicStore *NotaryPublicStoreSession) CreateApplication(domain string, account common.Address) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.CreateApplication(&_NotaryPublicStore.TransactOpts, domain, account)
}

// CreateApplication is a paid mutator transaction binding the contract method 0xb89d1d16.
//
// Solidity: function createApplication(string domain, address account) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactorSession) CreateApplication(domain string, account common.Address) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.CreateApplication(&_NotaryPublicStore.TransactOpts, domain, account)
}

// NotarizeApplication is a paid mutator transaction binding the contract method 0x3d8a8e19.
//
// Solidity: function notarizeApplication(uint256 applicationID, string network, string url, bytes signature) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactor) NotarizeApplication(opts *bind.TransactOpts, applicationID *big.Int, network string, url string, signature []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.contract.Transact(opts, "notarizeApplication", applicationID, network, url, signature)
}

// NotarizeApplication is a paid mutator transaction binding the contract method 0x3d8a8e19.
//
// Solidity: function notarizeApplication(uint256 applicationID, string network, string url, bytes signature) returns()
func (_NotaryPublicStore *NotaryPublicStoreSession) NotarizeApplication(applicationID *big.Int, network string, url string, signature []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.NotarizeApplication(&_NotaryPublicStore.TransactOpts, applicationID, network, url, signature)
}

// NotarizeApplication is a paid mutator transaction binding the contract method 0x3d8a8e19.
//
// Solidity: function notarizeApplication(uint256 applicationID, string network, string url, bytes signature) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactorSession) NotarizeApplication(applicationID *big.Int, network string, url string, signature []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.NotarizeApplication(&_NotaryPublicStore.TransactOpts, applicationID, network, url, signature)
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

// RevokeApplicationNotarization is a paid mutator transaction binding the contract method 0x12fbd138.
//
// Solidity: function revokeApplicationNotarization(uint256 applicationID, string network, bytes signature) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactor) RevokeApplicationNotarization(opts *bind.TransactOpts, applicationID *big.Int, network string, signature []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.contract.Transact(opts, "revokeApplicationNotarization", applicationID, network, signature)
}

// RevokeApplicationNotarization is a paid mutator transaction binding the contract method 0x12fbd138.
//
// Solidity: function revokeApplicationNotarization(uint256 applicationID, string network, bytes signature) returns()
func (_NotaryPublicStore *NotaryPublicStoreSession) RevokeApplicationNotarization(applicationID *big.Int, network string, signature []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.RevokeApplicationNotarization(&_NotaryPublicStore.TransactOpts, applicationID, network, signature)
}

// RevokeApplicationNotarization is a paid mutator transaction binding the contract method 0x12fbd138.
//
// Solidity: function revokeApplicationNotarization(uint256 applicationID, string network, bytes signature) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactorSession) RevokeApplicationNotarization(applicationID *big.Int, network string, signature []byte) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.RevokeApplicationNotarization(&_NotaryPublicStore.TransactOpts, applicationID, network, signature)
}

// TransferApplicationOwnership is a paid mutator transaction binding the contract method 0x4077e0b5.
//
// Solidity: function transferApplicationOwnership(uint256 applicationID, address newOwner) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactor) TransferApplicationOwnership(opts *bind.TransactOpts, applicationID *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _NotaryPublicStore.contract.Transact(opts, "transferApplicationOwnership", applicationID, newOwner)
}

// TransferApplicationOwnership is a paid mutator transaction binding the contract method 0x4077e0b5.
//
// Solidity: function transferApplicationOwnership(uint256 applicationID, address newOwner) returns()
func (_NotaryPublicStore *NotaryPublicStoreSession) TransferApplicationOwnership(applicationID *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.TransferApplicationOwnership(&_NotaryPublicStore.TransactOpts, applicationID, newOwner)
}

// TransferApplicationOwnership is a paid mutator transaction binding the contract method 0x4077e0b5.
//
// Solidity: function transferApplicationOwnership(uint256 applicationID, address newOwner) returns()
func (_NotaryPublicStore *NotaryPublicStoreTransactorSession) TransferApplicationOwnership(applicationID *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _NotaryPublicStore.Contract.TransferApplicationOwnership(&_NotaryPublicStore.TransactOpts, applicationID, newOwner)
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

// NotaryPublicStoreApplicationCreatedIterator is returned from FilterApplicationCreated and is used to iterate over the raw logs and unpacked data for ApplicationCreated events raised by the NotaryPublicStore contract.
type NotaryPublicStoreApplicationCreatedIterator struct {
	Event *NotaryPublicStoreApplicationCreated // Event containing the contract specifics and raw log

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
func (it *NotaryPublicStoreApplicationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryPublicStoreApplicationCreated)
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
		it.Event = new(NotaryPublicStoreApplicationCreated)
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
func (it *NotaryPublicStoreApplicationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryPublicStoreApplicationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryPublicStoreApplicationCreated represents a ApplicationCreated event raised by the NotaryPublicStore contract.
type NotaryPublicStoreApplicationCreated struct {
	ApplicationID *big.Int
	Domain        string
	Account       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterApplicationCreated is a free log retrieval operation binding the contract event 0x3d56de26da74f7b3c43307bef441e34b1d3cfab3df6b0f02d60fdd30e01d2f2e.
//
// Solidity: event ApplicationCreated(uint256 applicationID, string domain, address account)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) FilterApplicationCreated(opts *bind.FilterOpts) (*NotaryPublicStoreApplicationCreatedIterator, error) {

	logs, sub, err := _NotaryPublicStore.contract.FilterLogs(opts, "ApplicationCreated")
	if err != nil {
		return nil, err
	}
	return &NotaryPublicStoreApplicationCreatedIterator{contract: _NotaryPublicStore.contract, event: "ApplicationCreated", logs: logs, sub: sub}, nil
}

// WatchApplicationCreated is a free log subscription operation binding the contract event 0x3d56de26da74f7b3c43307bef441e34b1d3cfab3df6b0f02d60fdd30e01d2f2e.
//
// Solidity: event ApplicationCreated(uint256 applicationID, string domain, address account)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) WatchApplicationCreated(opts *bind.WatchOpts, sink chan<- *NotaryPublicStoreApplicationCreated) (event.Subscription, error) {

	logs, sub, err := _NotaryPublicStore.contract.WatchLogs(opts, "ApplicationCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryPublicStoreApplicationCreated)
				if err := _NotaryPublicStore.contract.UnpackLog(event, "ApplicationCreated", log); err != nil {
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

// ParseApplicationCreated is a log parse operation binding the contract event 0x3d56de26da74f7b3c43307bef441e34b1d3cfab3df6b0f02d60fdd30e01d2f2e.
//
// Solidity: event ApplicationCreated(uint256 applicationID, string domain, address account)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) ParseApplicationCreated(log types.Log) (*NotaryPublicStoreApplicationCreated, error) {
	event := new(NotaryPublicStoreApplicationCreated)
	if err := _NotaryPublicStore.contract.UnpackLog(event, "ApplicationCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryPublicStoreApplicationNotarizationRevokedIterator is returned from FilterApplicationNotarizationRevoked and is used to iterate over the raw logs and unpacked data for ApplicationNotarizationRevoked events raised by the NotaryPublicStore contract.
type NotaryPublicStoreApplicationNotarizationRevokedIterator struct {
	Event *NotaryPublicStoreApplicationNotarizationRevoked // Event containing the contract specifics and raw log

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
func (it *NotaryPublicStoreApplicationNotarizationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryPublicStoreApplicationNotarizationRevoked)
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
		it.Event = new(NotaryPublicStoreApplicationNotarizationRevoked)
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
func (it *NotaryPublicStoreApplicationNotarizationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryPublicStoreApplicationNotarizationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryPublicStoreApplicationNotarizationRevoked represents a ApplicationNotarizationRevoked event raised by the NotaryPublicStore contract.
type NotaryPublicStoreApplicationNotarizationRevoked struct {
	ApplicationID *big.Int
	Network       string
	Notary        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterApplicationNotarizationRevoked is a free log retrieval operation binding the contract event 0x918ffb70a8e2fd44a588ca28844edeb9ac84aa4a945ef483abb76ddb5223b2bc.
//
// Solidity: event ApplicationNotarizationRevoked(uint256 applicationID, string network, address notary)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) FilterApplicationNotarizationRevoked(opts *bind.FilterOpts) (*NotaryPublicStoreApplicationNotarizationRevokedIterator, error) {

	logs, sub, err := _NotaryPublicStore.contract.FilterLogs(opts, "ApplicationNotarizationRevoked")
	if err != nil {
		return nil, err
	}
	return &NotaryPublicStoreApplicationNotarizationRevokedIterator{contract: _NotaryPublicStore.contract, event: "ApplicationNotarizationRevoked", logs: logs, sub: sub}, nil
}

// WatchApplicationNotarizationRevoked is a free log subscription operation binding the contract event 0x918ffb70a8e2fd44a588ca28844edeb9ac84aa4a945ef483abb76ddb5223b2bc.
//
// Solidity: event ApplicationNotarizationRevoked(uint256 applicationID, string network, address notary)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) WatchApplicationNotarizationRevoked(opts *bind.WatchOpts, sink chan<- *NotaryPublicStoreApplicationNotarizationRevoked) (event.Subscription, error) {

	logs, sub, err := _NotaryPublicStore.contract.WatchLogs(opts, "ApplicationNotarizationRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryPublicStoreApplicationNotarizationRevoked)
				if err := _NotaryPublicStore.contract.UnpackLog(event, "ApplicationNotarizationRevoked", log); err != nil {
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

// ParseApplicationNotarizationRevoked is a log parse operation binding the contract event 0x918ffb70a8e2fd44a588ca28844edeb9ac84aa4a945ef483abb76ddb5223b2bc.
//
// Solidity: event ApplicationNotarizationRevoked(uint256 applicationID, string network, address notary)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) ParseApplicationNotarizationRevoked(log types.Log) (*NotaryPublicStoreApplicationNotarizationRevoked, error) {
	event := new(NotaryPublicStoreApplicationNotarizationRevoked)
	if err := _NotaryPublicStore.contract.UnpackLog(event, "ApplicationNotarizationRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryPublicStoreApplicationNotarizedIterator is returned from FilterApplicationNotarized and is used to iterate over the raw logs and unpacked data for ApplicationNotarized events raised by the NotaryPublicStore contract.
type NotaryPublicStoreApplicationNotarizedIterator struct {
	Event *NotaryPublicStoreApplicationNotarized // Event containing the contract specifics and raw log

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
func (it *NotaryPublicStoreApplicationNotarizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryPublicStoreApplicationNotarized)
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
		it.Event = new(NotaryPublicStoreApplicationNotarized)
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
func (it *NotaryPublicStoreApplicationNotarizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryPublicStoreApplicationNotarizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryPublicStoreApplicationNotarized represents a ApplicationNotarized event raised by the NotaryPublicStore contract.
type NotaryPublicStoreApplicationNotarized struct {
	ApplicationID *big.Int
	Network       string
	Notary        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterApplicationNotarized is a free log retrieval operation binding the contract event 0x3ab79fdc57369bd148d9f4751a22e9ced1eea3efcfc194323b0f2b87279de3d4.
//
// Solidity: event ApplicationNotarized(uint256 applicationID, string network, address notary)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) FilterApplicationNotarized(opts *bind.FilterOpts) (*NotaryPublicStoreApplicationNotarizedIterator, error) {

	logs, sub, err := _NotaryPublicStore.contract.FilterLogs(opts, "ApplicationNotarized")
	if err != nil {
		return nil, err
	}
	return &NotaryPublicStoreApplicationNotarizedIterator{contract: _NotaryPublicStore.contract, event: "ApplicationNotarized", logs: logs, sub: sub}, nil
}

// WatchApplicationNotarized is a free log subscription operation binding the contract event 0x3ab79fdc57369bd148d9f4751a22e9ced1eea3efcfc194323b0f2b87279de3d4.
//
// Solidity: event ApplicationNotarized(uint256 applicationID, string network, address notary)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) WatchApplicationNotarized(opts *bind.WatchOpts, sink chan<- *NotaryPublicStoreApplicationNotarized) (event.Subscription, error) {

	logs, sub, err := _NotaryPublicStore.contract.WatchLogs(opts, "ApplicationNotarized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryPublicStoreApplicationNotarized)
				if err := _NotaryPublicStore.contract.UnpackLog(event, "ApplicationNotarized", log); err != nil {
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

// ParseApplicationNotarized is a log parse operation binding the contract event 0x3ab79fdc57369bd148d9f4751a22e9ced1eea3efcfc194323b0f2b87279de3d4.
//
// Solidity: event ApplicationNotarized(uint256 applicationID, string network, address notary)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) ParseApplicationNotarized(log types.Log) (*NotaryPublicStoreApplicationNotarized, error) {
	event := new(NotaryPublicStoreApplicationNotarized)
	if err := _NotaryPublicStore.contract.UnpackLog(event, "ApplicationNotarized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryPublicStoreApplicationOwnershipTransferredIterator is returned from FilterApplicationOwnershipTransferred and is used to iterate over the raw logs and unpacked data for ApplicationOwnershipTransferred events raised by the NotaryPublicStore contract.
type NotaryPublicStoreApplicationOwnershipTransferredIterator struct {
	Event *NotaryPublicStoreApplicationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NotaryPublicStoreApplicationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryPublicStoreApplicationOwnershipTransferred)
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
		it.Event = new(NotaryPublicStoreApplicationOwnershipTransferred)
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
func (it *NotaryPublicStoreApplicationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryPublicStoreApplicationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryPublicStoreApplicationOwnershipTransferred represents a ApplicationOwnershipTransferred event raised by the NotaryPublicStore contract.
type NotaryPublicStoreApplicationOwnershipTransferred struct {
	ApplicationID *big.Int
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterApplicationOwnershipTransferred is a free log retrieval operation binding the contract event 0x6b46638de5ff197b187fba2bdcc3f5749301ea7a1db7469cbdea45155adc96d2.
//
// Solidity: event ApplicationOwnershipTransferred(uint256 applicationID, address previousOwner, address newOwner)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) FilterApplicationOwnershipTransferred(opts *bind.FilterOpts) (*NotaryPublicStoreApplicationOwnershipTransferredIterator, error) {

	logs, sub, err := _NotaryPublicStore.contract.FilterLogs(opts, "ApplicationOwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &NotaryPublicStoreApplicationOwnershipTransferredIterator{contract: _NotaryPublicStore.contract, event: "ApplicationOwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchApplicationOwnershipTransferred is a free log subscription operation binding the contract event 0x6b46638de5ff197b187fba2bdcc3f5749301ea7a1db7469cbdea45155adc96d2.
//
// Solidity: event ApplicationOwnershipTransferred(uint256 applicationID, address previousOwner, address newOwner)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) WatchApplicationOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NotaryPublicStoreApplicationOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _NotaryPublicStore.contract.WatchLogs(opts, "ApplicationOwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryPublicStoreApplicationOwnershipTransferred)
				if err := _NotaryPublicStore.contract.UnpackLog(event, "ApplicationOwnershipTransferred", log); err != nil {
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

// ParseApplicationOwnershipTransferred is a log parse operation binding the contract event 0x6b46638de5ff197b187fba2bdcc3f5749301ea7a1db7469cbdea45155adc96d2.
//
// Solidity: event ApplicationOwnershipTransferred(uint256 applicationID, address previousOwner, address newOwner)
func (_NotaryPublicStore *NotaryPublicStoreFilterer) ParseApplicationOwnershipTransferred(log types.Log) (*NotaryPublicStoreApplicationOwnershipTransferred, error) {
	event := new(NotaryPublicStoreApplicationOwnershipTransferred)
	if err := _NotaryPublicStore.contract.UnpackLog(event, "ApplicationOwnershipTransferred", log); err != nil {
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
