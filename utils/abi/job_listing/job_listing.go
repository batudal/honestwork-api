// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package job_listing

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

// JobListingPayment is an auto generated low-level Go binding around an user-defined struct.
type JobListingPayment struct {
	Token       common.Address
	Amount      *big.Int
	ListingDate *big.Int
}

// JobListingMetaData contains all meta data concerning the JobListing contract.
var JobListingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"PaymentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"PaymentAddedETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getLatestPayment\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingDate\",\"type\":\"uint256\"}],\"internalType\":\"structJobListing.Payment\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getPaymentsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingDate\",\"type\":\"uint256\"}],\"internalType\":\"structJobListing.Payment[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payForListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payForListingEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIHWRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawAllEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllEarningsEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// JobListingABI is the input ABI used to generate the binding from.
// Deprecated: Use JobListingMetaData.ABI instead.
var JobListingABI = JobListingMetaData.ABI

// JobListing is an auto generated Go binding around an Ethereum contract.
type JobListing struct {
	JobListingCaller     // Read-only binding to the contract
	JobListingTransactor // Write-only binding to the contract
	JobListingFilterer   // Log filterer for contract events
}

// JobListingCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobListingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobListingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobListingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobListingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobListingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobListingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobListingSession struct {
	Contract     *JobListing       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobListingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobListingCallerSession struct {
	Contract *JobListingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// JobListingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobListingTransactorSession struct {
	Contract     *JobListingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// JobListingRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobListingRaw struct {
	Contract *JobListing // Generic contract binding to access the raw methods on
}

// JobListingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobListingCallerRaw struct {
	Contract *JobListingCaller // Generic read-only contract binding to access the raw methods on
}

// JobListingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobListingTransactorRaw struct {
	Contract *JobListingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobListing creates a new instance of JobListing, bound to a specific deployed contract.
func NewJobListing(address common.Address, backend bind.ContractBackend) (*JobListing, error) {
	contract, err := bindJobListing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JobListing{JobListingCaller: JobListingCaller{contract: contract}, JobListingTransactor: JobListingTransactor{contract: contract}, JobListingFilterer: JobListingFilterer{contract: contract}}, nil
}

// NewJobListingCaller creates a new read-only instance of JobListing, bound to a specific deployed contract.
func NewJobListingCaller(address common.Address, caller bind.ContractCaller) (*JobListingCaller, error) {
	contract, err := bindJobListing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobListingCaller{contract: contract}, nil
}

// NewJobListingTransactor creates a new write-only instance of JobListing, bound to a specific deployed contract.
func NewJobListingTransactor(address common.Address, transactor bind.ContractTransactor) (*JobListingTransactor, error) {
	contract, err := bindJobListing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobListingTransactor{contract: contract}, nil
}

// NewJobListingFilterer creates a new log filterer instance of JobListing, bound to a specific deployed contract.
func NewJobListingFilterer(address common.Address, filterer bind.ContractFilterer) (*JobListingFilterer, error) {
	contract, err := bindJobListing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobListingFilterer{contract: contract}, nil
}

// bindJobListing binds a generic wrapper to an already deployed contract.
func bindJobListing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JobListingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobListing *JobListingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JobListing.Contract.JobListingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobListing *JobListingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobListing.Contract.JobListingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobListing *JobListingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobListing.Contract.JobListingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobListing *JobListingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JobListing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobListing *JobListingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobListing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobListing *JobListingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobListing.Contract.contract.Transact(opts, method, params...)
}

// GetLatestPayment is a free data retrieval call binding the contract method 0xa12a0365.
//
// Solidity: function getLatestPayment(address _user) view returns((address,uint256,uint256))
func (_JobListing *JobListingCaller) GetLatestPayment(opts *bind.CallOpts, _user common.Address) (JobListingPayment, error) {
	var out []interface{}
	err := _JobListing.contract.Call(opts, &out, "getLatestPayment", _user)

	if err != nil {
		return *new(JobListingPayment), err
	}

	out0 := *abi.ConvertType(out[0], new(JobListingPayment)).(*JobListingPayment)

	return out0, err

}

// GetLatestPayment is a free data retrieval call binding the contract method 0xa12a0365.
//
// Solidity: function getLatestPayment(address _user) view returns((address,uint256,uint256))
func (_JobListing *JobListingSession) GetLatestPayment(_user common.Address) (JobListingPayment, error) {
	return _JobListing.Contract.GetLatestPayment(&_JobListing.CallOpts, _user)
}

// GetLatestPayment is a free data retrieval call binding the contract method 0xa12a0365.
//
// Solidity: function getLatestPayment(address _user) view returns((address,uint256,uint256))
func (_JobListing *JobListingCallerSession) GetLatestPayment(_user common.Address) (JobListingPayment, error) {
	return _JobListing.Contract.GetLatestPayment(&_JobListing.CallOpts, _user)
}

// GetPaymentsOf is a free data retrieval call binding the contract method 0x8e446afb.
//
// Solidity: function getPaymentsOf(address _user) view returns((address,uint256,uint256)[])
func (_JobListing *JobListingCaller) GetPaymentsOf(opts *bind.CallOpts, _user common.Address) ([]JobListingPayment, error) {
	var out []interface{}
	err := _JobListing.contract.Call(opts, &out, "getPaymentsOf", _user)

	if err != nil {
		return *new([]JobListingPayment), err
	}

	out0 := *abi.ConvertType(out[0], new([]JobListingPayment)).(*[]JobListingPayment)

	return out0, err

}

// GetPaymentsOf is a free data retrieval call binding the contract method 0x8e446afb.
//
// Solidity: function getPaymentsOf(address _user) view returns((address,uint256,uint256)[])
func (_JobListing *JobListingSession) GetPaymentsOf(_user common.Address) ([]JobListingPayment, error) {
	return _JobListing.Contract.GetPaymentsOf(&_JobListing.CallOpts, _user)
}

// GetPaymentsOf is a free data retrieval call binding the contract method 0x8e446afb.
//
// Solidity: function getPaymentsOf(address _user) view returns((address,uint256,uint256)[])
func (_JobListing *JobListingCallerSession) GetPaymentsOf(_user common.Address) ([]JobListingPayment, error) {
	return _JobListing.Contract.GetPaymentsOf(&_JobListing.CallOpts, _user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JobListing *JobListingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JobListing.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JobListing *JobListingSession) Owner() (common.Address, error) {
	return _JobListing.Contract.Owner(&_JobListing.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JobListing *JobListingCallerSession) Owner() (common.Address, error) {
	return _JobListing.Contract.Owner(&_JobListing.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_JobListing *JobListingCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JobListing.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_JobListing *JobListingSession) Registry() (common.Address, error) {
	return _JobListing.Contract.Registry(&_JobListing.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_JobListing *JobListingCallerSession) Registry() (common.Address, error) {
	return _JobListing.Contract.Registry(&_JobListing.CallOpts)
}

// PayForListing is a paid mutator transaction binding the contract method 0x304db7d1.
//
// Solidity: function payForListing(address _token, uint256 _amount) returns()
func (_JobListing *JobListingTransactor) PayForListing(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _JobListing.contract.Transact(opts, "payForListing", _token, _amount)
}

// PayForListing is a paid mutator transaction binding the contract method 0x304db7d1.
//
// Solidity: function payForListing(address _token, uint256 _amount) returns()
func (_JobListing *JobListingSession) PayForListing(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _JobListing.Contract.PayForListing(&_JobListing.TransactOpts, _token, _amount)
}

// PayForListing is a paid mutator transaction binding the contract method 0x304db7d1.
//
// Solidity: function payForListing(address _token, uint256 _amount) returns()
func (_JobListing *JobListingTransactorSession) PayForListing(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _JobListing.Contract.PayForListing(&_JobListing.TransactOpts, _token, _amount)
}

// PayForListingEth is a paid mutator transaction binding the contract method 0x2d61732b.
//
// Solidity: function payForListingEth() payable returns()
func (_JobListing *JobListingTransactor) PayForListingEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobListing.contract.Transact(opts, "payForListingEth")
}

// PayForListingEth is a paid mutator transaction binding the contract method 0x2d61732b.
//
// Solidity: function payForListingEth() payable returns()
func (_JobListing *JobListingSession) PayForListingEth() (*types.Transaction, error) {
	return _JobListing.Contract.PayForListingEth(&_JobListing.TransactOpts)
}

// PayForListingEth is a paid mutator transaction binding the contract method 0x2d61732b.
//
// Solidity: function payForListingEth() payable returns()
func (_JobListing *JobListingTransactorSession) PayForListingEth() (*types.Transaction, error) {
	return _JobListing.Contract.PayForListingEth(&_JobListing.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JobListing *JobListingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobListing.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JobListing *JobListingSession) RenounceOwnership() (*types.Transaction, error) {
	return _JobListing.Contract.RenounceOwnership(&_JobListing.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JobListing *JobListingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _JobListing.Contract.RenounceOwnership(&_JobListing.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JobListing *JobListingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _JobListing.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JobListing *JobListingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _JobListing.Contract.TransferOwnership(&_JobListing.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JobListing *JobListingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _JobListing.Contract.TransferOwnership(&_JobListing.TransactOpts, newOwner)
}

// WithdrawAllEarnings is a paid mutator transaction binding the contract method 0xc4218374.
//
// Solidity: function withdrawAllEarnings(address _token) returns()
func (_JobListing *JobListingTransactor) WithdrawAllEarnings(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _JobListing.contract.Transact(opts, "withdrawAllEarnings", _token)
}

// WithdrawAllEarnings is a paid mutator transaction binding the contract method 0xc4218374.
//
// Solidity: function withdrawAllEarnings(address _token) returns()
func (_JobListing *JobListingSession) WithdrawAllEarnings(_token common.Address) (*types.Transaction, error) {
	return _JobListing.Contract.WithdrawAllEarnings(&_JobListing.TransactOpts, _token)
}

// WithdrawAllEarnings is a paid mutator transaction binding the contract method 0xc4218374.
//
// Solidity: function withdrawAllEarnings(address _token) returns()
func (_JobListing *JobListingTransactorSession) WithdrawAllEarnings(_token common.Address) (*types.Transaction, error) {
	return _JobListing.Contract.WithdrawAllEarnings(&_JobListing.TransactOpts, _token)
}

// WithdrawAllEarningsEth is a paid mutator transaction binding the contract method 0xc7c067af.
//
// Solidity: function withdrawAllEarningsEth() returns()
func (_JobListing *JobListingTransactor) WithdrawAllEarningsEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobListing.contract.Transact(opts, "withdrawAllEarningsEth")
}

// WithdrawAllEarningsEth is a paid mutator transaction binding the contract method 0xc7c067af.
//
// Solidity: function withdrawAllEarningsEth() returns()
func (_JobListing *JobListingSession) WithdrawAllEarningsEth() (*types.Transaction, error) {
	return _JobListing.Contract.WithdrawAllEarningsEth(&_JobListing.TransactOpts)
}

// WithdrawAllEarningsEth is a paid mutator transaction binding the contract method 0xc7c067af.
//
// Solidity: function withdrawAllEarningsEth() returns()
func (_JobListing *JobListingTransactorSession) WithdrawAllEarningsEth() (*types.Transaction, error) {
	return _JobListing.Contract.WithdrawAllEarningsEth(&_JobListing.TransactOpts)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0x280da6fa.
//
// Solidity: function withdrawAllTokens() returns()
func (_JobListing *JobListingTransactor) WithdrawAllTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobListing.contract.Transact(opts, "withdrawAllTokens")
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0x280da6fa.
//
// Solidity: function withdrawAllTokens() returns()
func (_JobListing *JobListingSession) WithdrawAllTokens() (*types.Transaction, error) {
	return _JobListing.Contract.WithdrawAllTokens(&_JobListing.TransactOpts)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0x280da6fa.
//
// Solidity: function withdrawAllTokens() returns()
func (_JobListing *JobListingTransactorSession) WithdrawAllTokens() (*types.Transaction, error) {
	return _JobListing.Contract.WithdrawAllTokens(&_JobListing.TransactOpts)
}

// WithdrawEarnings is a paid mutator transaction binding the contract method 0xc9211c82.
//
// Solidity: function withdrawEarnings(address _token, uint256 _amount) returns()
func (_JobListing *JobListingTransactor) WithdrawEarnings(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _JobListing.contract.Transact(opts, "withdrawEarnings", _token, _amount)
}

// WithdrawEarnings is a paid mutator transaction binding the contract method 0xc9211c82.
//
// Solidity: function withdrawEarnings(address _token, uint256 _amount) returns()
func (_JobListing *JobListingSession) WithdrawEarnings(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _JobListing.Contract.WithdrawEarnings(&_JobListing.TransactOpts, _token, _amount)
}

// WithdrawEarnings is a paid mutator transaction binding the contract method 0xc9211c82.
//
// Solidity: function withdrawEarnings(address _token, uint256 _amount) returns()
func (_JobListing *JobListingTransactorSession) WithdrawEarnings(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _JobListing.Contract.WithdrawEarnings(&_JobListing.TransactOpts, _token, _amount)
}

// JobListingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the JobListing contract.
type JobListingOwnershipTransferredIterator struct {
	Event *JobListingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *JobListingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobListingOwnershipTransferred)
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
		it.Event = new(JobListingOwnershipTransferred)
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
func (it *JobListingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobListingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobListingOwnershipTransferred represents a OwnershipTransferred event raised by the JobListing contract.
type JobListingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_JobListing *JobListingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*JobListingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _JobListing.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &JobListingOwnershipTransferredIterator{contract: _JobListing.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_JobListing *JobListingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *JobListingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _JobListing.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobListingOwnershipTransferred)
				if err := _JobListing.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_JobListing *JobListingFilterer) ParseOwnershipTransferred(log types.Log) (*JobListingOwnershipTransferred, error) {
	event := new(JobListingOwnershipTransferred)
	if err := _JobListing.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobListingPaymentAddedIterator is returned from FilterPaymentAdded and is used to iterate over the raw logs and unpacked data for PaymentAdded events raised by the JobListing contract.
type JobListingPaymentAddedIterator struct {
	Event *JobListingPaymentAdded // Event containing the contract specifics and raw log

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
func (it *JobListingPaymentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobListingPaymentAdded)
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
		it.Event = new(JobListingPaymentAdded)
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
func (it *JobListingPaymentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobListingPaymentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobListingPaymentAdded represents a PaymentAdded event raised by the JobListing contract.
type JobListingPaymentAdded struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentAdded is a free log retrieval operation binding the contract event 0x7ceb32c988f6127615957ec467cbb5688c54a860a192790b125e54b9c3b681e0.
//
// Solidity: event PaymentAdded(address indexed _token, uint256 _amount)
func (_JobListing *JobListingFilterer) FilterPaymentAdded(opts *bind.FilterOpts, _token []common.Address) (*JobListingPaymentAddedIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _JobListing.contract.FilterLogs(opts, "PaymentAdded", _tokenRule)
	if err != nil {
		return nil, err
	}
	return &JobListingPaymentAddedIterator{contract: _JobListing.contract, event: "PaymentAdded", logs: logs, sub: sub}, nil
}

// WatchPaymentAdded is a free log subscription operation binding the contract event 0x7ceb32c988f6127615957ec467cbb5688c54a860a192790b125e54b9c3b681e0.
//
// Solidity: event PaymentAdded(address indexed _token, uint256 _amount)
func (_JobListing *JobListingFilterer) WatchPaymentAdded(opts *bind.WatchOpts, sink chan<- *JobListingPaymentAdded, _token []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _JobListing.contract.WatchLogs(opts, "PaymentAdded", _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobListingPaymentAdded)
				if err := _JobListing.contract.UnpackLog(event, "PaymentAdded", log); err != nil {
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

// ParsePaymentAdded is a log parse operation binding the contract event 0x7ceb32c988f6127615957ec467cbb5688c54a860a192790b125e54b9c3b681e0.
//
// Solidity: event PaymentAdded(address indexed _token, uint256 _amount)
func (_JobListing *JobListingFilterer) ParsePaymentAdded(log types.Log) (*JobListingPaymentAdded, error) {
	event := new(JobListingPaymentAdded)
	if err := _JobListing.contract.UnpackLog(event, "PaymentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobListingPaymentAddedETHIterator is returned from FilterPaymentAddedETH and is used to iterate over the raw logs and unpacked data for PaymentAddedETH events raised by the JobListing contract.
type JobListingPaymentAddedETHIterator struct {
	Event *JobListingPaymentAddedETH // Event containing the contract specifics and raw log

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
func (it *JobListingPaymentAddedETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobListingPaymentAddedETH)
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
		it.Event = new(JobListingPaymentAddedETH)
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
func (it *JobListingPaymentAddedETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobListingPaymentAddedETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobListingPaymentAddedETH represents a PaymentAddedETH event raised by the JobListing contract.
type JobListingPaymentAddedETH struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentAddedETH is a free log retrieval operation binding the contract event 0xf8b14910ac2e6a8e7967441ce0230a43990764e0c669d0681e91e42749cdb18b.
//
// Solidity: event PaymentAddedETH(uint256 _amount)
func (_JobListing *JobListingFilterer) FilterPaymentAddedETH(opts *bind.FilterOpts) (*JobListingPaymentAddedETHIterator, error) {

	logs, sub, err := _JobListing.contract.FilterLogs(opts, "PaymentAddedETH")
	if err != nil {
		return nil, err
	}
	return &JobListingPaymentAddedETHIterator{contract: _JobListing.contract, event: "PaymentAddedETH", logs: logs, sub: sub}, nil
}

// WatchPaymentAddedETH is a free log subscription operation binding the contract event 0xf8b14910ac2e6a8e7967441ce0230a43990764e0c669d0681e91e42749cdb18b.
//
// Solidity: event PaymentAddedETH(uint256 _amount)
func (_JobListing *JobListingFilterer) WatchPaymentAddedETH(opts *bind.WatchOpts, sink chan<- *JobListingPaymentAddedETH) (event.Subscription, error) {

	logs, sub, err := _JobListing.contract.WatchLogs(opts, "PaymentAddedETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobListingPaymentAddedETH)
				if err := _JobListing.contract.UnpackLog(event, "PaymentAddedETH", log); err != nil {
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

// ParsePaymentAddedETH is a log parse operation binding the contract event 0xf8b14910ac2e6a8e7967441ce0230a43990764e0c669d0681e91e42749cdb18b.
//
// Solidity: event PaymentAddedETH(uint256 _amount)
func (_JobListing *JobListingFilterer) ParsePaymentAddedETH(log types.Log) (*JobListingPaymentAddedETH, error) {
	event := new(JobListingPaymentAddedETH)
	if err := _JobListing.contract.UnpackLog(event, "PaymentAddedETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
