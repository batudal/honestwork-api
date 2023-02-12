// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// HWRegistryWhitelist is an auto generated low-level Go binding around an user-defined struct.
type HWRegistryWhitelist struct {
	Token      common.Address
	MaxAllowed *big.Int
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_maxAllowed\",\"type\":\"uint256\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_maxAllowed\",\"type\":\"uint256\"}],\"name\":\"WhitelistedUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxAllowed\",\"type\":\"uint256\"}],\"name\":\"addWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allWhitelisted\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAllowed\",\"type\":\"uint256\"}],\"internalType\":\"structHWRegistry.Whitelist[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getNFTGrossRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"honestPayLock\",\"outputs\":[{\"internalType\":\"contractIHonestPayLock\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"isAllowedAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftGrossRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setHonestPayLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setNFTGrossRevenue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxAllowed\",\"type\":\"uint256\"}],\"name\":\"updateWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelisted\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAllowed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}],",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// AllWhitelisted is a free data retrieval call binding the contract method 0x765b9d6b.
//
// Solidity: function allWhitelisted() view returns((address,uint256)[])
func (_Registry *RegistryCaller) AllWhitelisted(opts *bind.CallOpts) ([]HWRegistryWhitelist, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "allWhitelisted")

	if err != nil {
		return *new([]HWRegistryWhitelist), err
	}

	out0 := *abi.ConvertType(out[0], new([]HWRegistryWhitelist)).(*[]HWRegistryWhitelist)

	return out0, err

}

// AllWhitelisted is a free data retrieval call binding the contract method 0x765b9d6b.
//
// Solidity: function allWhitelisted() view returns((address,uint256)[])
func (_Registry *RegistrySession) AllWhitelisted() ([]HWRegistryWhitelist, error) {
	return _Registry.Contract.AllWhitelisted(&_Registry.CallOpts)
}

// AllWhitelisted is a free data retrieval call binding the contract method 0x765b9d6b.
//
// Solidity: function allWhitelisted() view returns((address,uint256)[])
func (_Registry *RegistryCallerSession) AllWhitelisted() ([]HWRegistryWhitelist, error) {
	return _Registry.Contract.AllWhitelisted(&_Registry.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256 _value)
func (_Registry *RegistryCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256 _value)
func (_Registry *RegistrySession) Counter() (*big.Int, error) {
	return _Registry.Contract.Counter(&_Registry.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256 _value)
func (_Registry *RegistryCallerSession) Counter() (*big.Int, error) {
	return _Registry.Contract.Counter(&_Registry.CallOpts)
}

// GetNFTGrossRevenue is a free data retrieval call binding the contract method 0x13faf9df.
//
// Solidity: function getNFTGrossRevenue(uint256 _id) view returns(uint256)
func (_Registry *RegistryCaller) GetNFTGrossRevenue(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getNFTGrossRevenue", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNFTGrossRevenue is a free data retrieval call binding the contract method 0x13faf9df.
//
// Solidity: function getNFTGrossRevenue(uint256 _id) view returns(uint256)
func (_Registry *RegistrySession) GetNFTGrossRevenue(_id *big.Int) (*big.Int, error) {
	return _Registry.Contract.GetNFTGrossRevenue(&_Registry.CallOpts, _id)
}

// GetNFTGrossRevenue is a free data retrieval call binding the contract method 0x13faf9df.
//
// Solidity: function getNFTGrossRevenue(uint256 _id) view returns(uint256)
func (_Registry *RegistryCallerSession) GetNFTGrossRevenue(_id *big.Int) (*big.Int, error) {
	return _Registry.Contract.GetNFTGrossRevenue(&_Registry.CallOpts, _id)
}

// HonestPayLock is a free data retrieval call binding the contract method 0xf3d5a8a2.
//
// Solidity: function honestPayLock() view returns(address)
func (_Registry *RegistryCaller) HonestPayLock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "honestPayLock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HonestPayLock is a free data retrieval call binding the contract method 0xf3d5a8a2.
//
// Solidity: function honestPayLock() view returns(address)
func (_Registry *RegistrySession) HonestPayLock() (common.Address, error) {
	return _Registry.Contract.HonestPayLock(&_Registry.CallOpts)
}

// HonestPayLock is a free data retrieval call binding the contract method 0xf3d5a8a2.
//
// Solidity: function honestPayLock() view returns(address)
func (_Registry *RegistryCallerSession) HonestPayLock() (common.Address, error) {
	return _Registry.Contract.HonestPayLock(&_Registry.CallOpts)
}

// IsAllowedAmount is a free data retrieval call binding the contract method 0x83ed69c1.
//
// Solidity: function isAllowedAmount(address _address, uint256 _amount) view returns(bool)
func (_Registry *RegistryCaller) IsAllowedAmount(opts *bind.CallOpts, _address common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isAllowedAmount", _address, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedAmount is a free data retrieval call binding the contract method 0x83ed69c1.
//
// Solidity: function isAllowedAmount(address _address, uint256 _amount) view returns(bool)
func (_Registry *RegistrySession) IsAllowedAmount(_address common.Address, _amount *big.Int) (bool, error) {
	return _Registry.Contract.IsAllowedAmount(&_Registry.CallOpts, _address, _amount)
}

// IsAllowedAmount is a free data retrieval call binding the contract method 0x83ed69c1.
//
// Solidity: function isAllowedAmount(address _address, uint256 _amount) view returns(bool)
func (_Registry *RegistryCallerSession) IsAllowedAmount(_address common.Address, _amount *big.Int) (bool, error) {
	return _Registry.Contract.IsAllowedAmount(&_Registry.CallOpts, _address, _amount)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_Registry *RegistryCaller) IsWhitelisted(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isWhitelisted", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_Registry *RegistrySession) IsWhitelisted(_address common.Address) (bool, error) {
	return _Registry.Contract.IsWhitelisted(&_Registry.CallOpts, _address)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_Registry *RegistryCallerSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _Registry.Contract.IsWhitelisted(&_Registry.CallOpts, _address)
}

// NftGrossRevenue is a free data retrieval call binding the contract method 0x43120455.
//
// Solidity: function nftGrossRevenue(uint256 ) view returns(uint256)
func (_Registry *RegistryCaller) NftGrossRevenue(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "nftGrossRevenue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftGrossRevenue is a free data retrieval call binding the contract method 0x43120455.
//
// Solidity: function nftGrossRevenue(uint256 ) view returns(uint256)
func (_Registry *RegistrySession) NftGrossRevenue(arg0 *big.Int) (*big.Int, error) {
	return _Registry.Contract.NftGrossRevenue(&_Registry.CallOpts, arg0)
}

// NftGrossRevenue is a free data retrieval call binding the contract method 0x43120455.
//
// Solidity: function nftGrossRevenue(uint256 ) view returns(uint256)
func (_Registry *RegistryCallerSession) NftGrossRevenue(arg0 *big.Int) (*big.Int, error) {
	return _Registry.Contract.NftGrossRevenue(&_Registry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0x3d4efe09.
//
// Solidity: function whitelisted(uint256 ) view returns(address token, uint256 maxAllowed)
func (_Registry *RegistryCaller) Whitelisted(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token      common.Address
	MaxAllowed *big.Int
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "whitelisted", arg0)

	outstruct := new(struct {
		Token      common.Address
		MaxAllowed *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MaxAllowed = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Whitelisted is a free data retrieval call binding the contract method 0x3d4efe09.
//
// Solidity: function whitelisted(uint256 ) view returns(address token, uint256 maxAllowed)
func (_Registry *RegistrySession) Whitelisted(arg0 *big.Int) (struct {
	Token      common.Address
	MaxAllowed *big.Int
}, error) {
	return _Registry.Contract.Whitelisted(&_Registry.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0x3d4efe09.
//
// Solidity: function whitelisted(uint256 ) view returns(address token, uint256 maxAllowed)
func (_Registry *RegistryCallerSession) Whitelisted(arg0 *big.Int) (struct {
	Token      common.Address
	MaxAllowed *big.Int
}, error) {
	return _Registry.Contract.Whitelisted(&_Registry.CallOpts, arg0)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x8902ff86.
//
// Solidity: function addWhitelisted(address _address, uint256 _maxAllowed) returns(bool)
func (_Registry *RegistryTransactor) AddWhitelisted(opts *bind.TransactOpts, _address common.Address, _maxAllowed *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addWhitelisted", _address, _maxAllowed)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x8902ff86.
//
// Solidity: function addWhitelisted(address _address, uint256 _maxAllowed) returns(bool)
func (_Registry *RegistrySession) AddWhitelisted(_address common.Address, _maxAllowed *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.AddWhitelisted(&_Registry.TransactOpts, _address, _maxAllowed)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x8902ff86.
//
// Solidity: function addWhitelisted(address _address, uint256 _maxAllowed) returns(bool)
func (_Registry *RegistryTransactorSession) AddWhitelisted(_address common.Address, _maxAllowed *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.AddWhitelisted(&_Registry.TransactOpts, _address, _maxAllowed)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns(bool)
func (_Registry *RegistryTransactor) RemoveWhitelisted(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "removeWhitelisted", _address)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns(bool)
func (_Registry *RegistrySession) RemoveWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemoveWhitelisted(&_Registry.TransactOpts, _address)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns(bool)
func (_Registry *RegistryTransactorSession) RemoveWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemoveWhitelisted(&_Registry.TransactOpts, _address)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// SetHonestPayLock is a paid mutator transaction binding the contract method 0x20820803.
//
// Solidity: function setHonestPayLock(address _address) returns(bool)
func (_Registry *RegistryTransactor) SetHonestPayLock(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setHonestPayLock", _address)
}

// SetHonestPayLock is a paid mutator transaction binding the contract method 0x20820803.
//
// Solidity: function setHonestPayLock(address _address) returns(bool)
func (_Registry *RegistrySession) SetHonestPayLock(_address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetHonestPayLock(&_Registry.TransactOpts, _address)
}

// SetHonestPayLock is a paid mutator transaction binding the contract method 0x20820803.
//
// Solidity: function setHonestPayLock(address _address) returns(bool)
func (_Registry *RegistryTransactorSession) SetHonestPayLock(_address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetHonestPayLock(&_Registry.TransactOpts, _address)
}

// SetNFTGrossRevenue is a paid mutator transaction binding the contract method 0x155e85c4.
//
// Solidity: function setNFTGrossRevenue(uint256 _id, uint256 _amount) returns()
func (_Registry *RegistryTransactor) SetNFTGrossRevenue(opts *bind.TransactOpts, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setNFTGrossRevenue", _id, _amount)
}

// SetNFTGrossRevenue is a paid mutator transaction binding the contract method 0x155e85c4.
//
// Solidity: function setNFTGrossRevenue(uint256 _id, uint256 _amount) returns()
func (_Registry *RegistrySession) SetNFTGrossRevenue(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetNFTGrossRevenue(&_Registry.TransactOpts, _id, _amount)
}

// SetNFTGrossRevenue is a paid mutator transaction binding the contract method 0x155e85c4.
//
// Solidity: function setNFTGrossRevenue(uint256 _id, uint256 _amount) returns()
func (_Registry *RegistryTransactorSession) SetNFTGrossRevenue(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetNFTGrossRevenue(&_Registry.TransactOpts, _id, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// UpdateWhitelisted is a paid mutator transaction binding the contract method 0xcd1188a5.
//
// Solidity: function updateWhitelisted(address _address, uint256 _maxAllowed) returns(bool)
func (_Registry *RegistryTransactor) UpdateWhitelisted(opts *bind.TransactOpts, _address common.Address, _maxAllowed *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateWhitelisted", _address, _maxAllowed)
}

// UpdateWhitelisted is a paid mutator transaction binding the contract method 0xcd1188a5.
//
// Solidity: function updateWhitelisted(address _address, uint256 _maxAllowed) returns(bool)
func (_Registry *RegistrySession) UpdateWhitelisted(_address common.Address, _maxAllowed *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UpdateWhitelisted(&_Registry.TransactOpts, _address, _maxAllowed)
}

// UpdateWhitelisted is a paid mutator transaction binding the contract method 0xcd1188a5.
//
// Solidity: function updateWhitelisted(address _address, uint256 _maxAllowed) returns(bool)
func (_Registry *RegistryTransactorSession) UpdateWhitelisted(_address common.Address, _maxAllowed *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UpdateWhitelisted(&_Registry.TransactOpts, _address, _maxAllowed)
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryOwnershipTransferred, error) {
	event := new(RegistryOwnershipTransferred)
	if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryWhitelistedAddedIterator is returned from FilterWhitelistedAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAdded events raised by the Registry contract.
type RegistryWhitelistedAddedIterator struct {
	Event *RegistryWhitelistedAdded // Event containing the contract specifics and raw log

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
func (it *RegistryWhitelistedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryWhitelistedAdded)
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
		it.Event = new(RegistryWhitelistedAdded)
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
func (it *RegistryWhitelistedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryWhitelistedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryWhitelistedAdded represents a WhitelistedAdded event raised by the Registry contract.
type RegistryWhitelistedAdded struct {
	Address    common.Address
	MaxAllowed *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAdded is a free log retrieval operation binding the contract event 0xcb9ce803cb004a01f0682efa32d66c42ec72e8daa764006c95e1b551b9465a21.
//
// Solidity: event WhitelistedAdded(address indexed _address, uint256 _maxAllowed)
func (_Registry *RegistryFilterer) FilterWhitelistedAdded(opts *bind.FilterOpts, _address []common.Address) (*RegistryWhitelistedAddedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "WhitelistedAdded", _addressRule)
	if err != nil {
		return nil, err
	}
	return &RegistryWhitelistedAddedIterator{contract: _Registry.contract, event: "WhitelistedAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAdded is a free log subscription operation binding the contract event 0xcb9ce803cb004a01f0682efa32d66c42ec72e8daa764006c95e1b551b9465a21.
//
// Solidity: event WhitelistedAdded(address indexed _address, uint256 _maxAllowed)
func (_Registry *RegistryFilterer) WatchWhitelistedAdded(opts *bind.WatchOpts, sink chan<- *RegistryWhitelistedAdded, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "WhitelistedAdded", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryWhitelistedAdded)
				if err := _Registry.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
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

// ParseWhitelistedAdded is a log parse operation binding the contract event 0xcb9ce803cb004a01f0682efa32d66c42ec72e8daa764006c95e1b551b9465a21.
//
// Solidity: event WhitelistedAdded(address indexed _address, uint256 _maxAllowed)
func (_Registry *RegistryFilterer) ParseWhitelistedAdded(log types.Log) (*RegistryWhitelistedAdded, error) {
	event := new(RegistryWhitelistedAdded)
	if err := _Registry.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryWhitelistedRemovedIterator is returned from FilterWhitelistedRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedRemoved events raised by the Registry contract.
type RegistryWhitelistedRemovedIterator struct {
	Event *RegistryWhitelistedRemoved // Event containing the contract specifics and raw log

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
func (it *RegistryWhitelistedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryWhitelistedRemoved)
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
		it.Event = new(RegistryWhitelistedRemoved)
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
func (it *RegistryWhitelistedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryWhitelistedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryWhitelistedRemoved represents a WhitelistedRemoved event raised by the Registry contract.
type RegistryWhitelistedRemoved struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedRemoved is a free log retrieval operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address indexed _address)
func (_Registry *RegistryFilterer) FilterWhitelistedRemoved(opts *bind.FilterOpts, _address []common.Address) (*RegistryWhitelistedRemovedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "WhitelistedRemoved", _addressRule)
	if err != nil {
		return nil, err
	}
	return &RegistryWhitelistedRemovedIterator{contract: _Registry.contract, event: "WhitelistedRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedRemoved is a free log subscription operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address indexed _address)
func (_Registry *RegistryFilterer) WatchWhitelistedRemoved(opts *bind.WatchOpts, sink chan<- *RegistryWhitelistedRemoved, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "WhitelistedRemoved", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryWhitelistedRemoved)
				if err := _Registry.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
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

// ParseWhitelistedRemoved is a log parse operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address indexed _address)
func (_Registry *RegistryFilterer) ParseWhitelistedRemoved(log types.Log) (*RegistryWhitelistedRemoved, error) {
	event := new(RegistryWhitelistedRemoved)
	if err := _Registry.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryWhitelistedUpdatedIterator is returned from FilterWhitelistedUpdated and is used to iterate over the raw logs and unpacked data for WhitelistedUpdated events raised by the Registry contract.
type RegistryWhitelistedUpdatedIterator struct {
	Event *RegistryWhitelistedUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryWhitelistedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryWhitelistedUpdated)
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
		it.Event = new(RegistryWhitelistedUpdated)
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
func (it *RegistryWhitelistedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryWhitelistedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryWhitelistedUpdated represents a WhitelistedUpdated event raised by the Registry contract.
type RegistryWhitelistedUpdated struct {
	Address    common.Address
	MaxAllowed *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedUpdated is a free log retrieval operation binding the contract event 0x8bc3b900d96200ab851acb4ccd80670a931934f8fbae20691e95a683c0d00e60.
//
// Solidity: event WhitelistedUpdated(address indexed _address, uint256 _maxAllowed)
func (_Registry *RegistryFilterer) FilterWhitelistedUpdated(opts *bind.FilterOpts, _address []common.Address) (*RegistryWhitelistedUpdatedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "WhitelistedUpdated", _addressRule)
	if err != nil {
		return nil, err
	}
	return &RegistryWhitelistedUpdatedIterator{contract: _Registry.contract, event: "WhitelistedUpdated", logs: logs, sub: sub}, nil
}

// WatchWhitelistedUpdated is a free log subscription operation binding the contract event 0x8bc3b900d96200ab851acb4ccd80670a931934f8fbae20691e95a683c0d00e60.
//
// Solidity: event WhitelistedUpdated(address indexed _address, uint256 _maxAllowed)
func (_Registry *RegistryFilterer) WatchWhitelistedUpdated(opts *bind.WatchOpts, sink chan<- *RegistryWhitelistedUpdated, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "WhitelistedUpdated", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryWhitelistedUpdated)
				if err := _Registry.contract.UnpackLog(event, "WhitelistedUpdated", log); err != nil {
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

// ParseWhitelistedUpdated is a log parse operation binding the contract event 0x8bc3b900d96200ab851acb4ccd80670a931934f8fbae20691e95a683c0d00e60.
//
// Solidity: event WhitelistedUpdated(address indexed _address, uint256 _maxAllowed)
func (_Registry *RegistryFilterer) ParseWhitelistedUpdated(log types.Log) (*RegistryWhitelistedUpdated, error) {
	event := new(RegistryWhitelistedUpdated)
	if err := _Registry.contract.UnpackLog(event, "WhitelistedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
