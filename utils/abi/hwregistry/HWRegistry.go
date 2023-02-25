// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hwregistry

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

// HwregistryMetaData contains all meta data concerning the Hwregistry contract.
var HwregistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_maxAllowed\",\"type\":\"uint256\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_maxAllowed\",\"type\":\"uint256\"}],\"name\":\"WhitelistedUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxAllowed\",\"type\":\"uint256\"}],\"name\":\"addWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allWhitelisted\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAllowed\",\"type\":\"uint256\"}],\"internalType\":\"structHWRegistry.Whitelist[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getNFTGrossRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hwEscrow\",\"outputs\":[{\"internalType\":\"contractIHWEscrow\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"isAllowedAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftGrossRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setHWEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setNFTGrossRevenue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxAllowed\",\"type\":\"uint256\"}],\"name\":\"updateWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelisted\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAllowed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HwregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use HwregistryMetaData.ABI instead.
var HwregistryABI = HwregistryMetaData.ABI

// Hwregistry is an auto generated Go binding around an Ethereum contract.
type Hwregistry struct {
	HwregistryCaller     // Read-only binding to the contract
	HwregistryTransactor // Write-only binding to the contract
	HwregistryFilterer   // Log filterer for contract events
}

// HwregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type HwregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HwregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HwregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HwregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HwregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HwregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HwregistrySession struct {
	Contract     *Hwregistry       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HwregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HwregistryCallerSession struct {
	Contract *HwregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HwregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HwregistryTransactorSession struct {
	Contract     *HwregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HwregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type HwregistryRaw struct {
	Contract *Hwregistry // Generic contract binding to access the raw methods on
}

// HwregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HwregistryCallerRaw struct {
	Contract *HwregistryCaller // Generic read-only contract binding to access the raw methods on
}

// HwregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HwregistryTransactorRaw struct {
	Contract *HwregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHwregistry creates a new instance of Hwregistry, bound to a specific deployed contract.
func NewHwregistry(address common.Address, backend bind.ContractBackend) (*Hwregistry, error) {
	contract, err := bindHwregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hwregistry{HwregistryCaller: HwregistryCaller{contract: contract}, HwregistryTransactor: HwregistryTransactor{contract: contract}, HwregistryFilterer: HwregistryFilterer{contract: contract}}, nil
}

// NewHwregistryCaller creates a new read-only instance of Hwregistry, bound to a specific deployed contract.
func NewHwregistryCaller(address common.Address, caller bind.ContractCaller) (*HwregistryCaller, error) {
	contract, err := bindHwregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HwregistryCaller{contract: contract}, nil
}

// NewHwregistryTransactor creates a new write-only instance of Hwregistry, bound to a specific deployed contract.
func NewHwregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*HwregistryTransactor, error) {
	contract, err := bindHwregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HwregistryTransactor{contract: contract}, nil
}

// NewHwregistryFilterer creates a new log filterer instance of Hwregistry, bound to a specific deployed contract.
func NewHwregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*HwregistryFilterer, error) {
	contract, err := bindHwregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HwregistryFilterer{contract: contract}, nil
}

// bindHwregistry binds a generic wrapper to an already deployed contract.
func bindHwregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HwregistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hwregistry *HwregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hwregistry.Contract.HwregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hwregistry *HwregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hwregistry.Contract.HwregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hwregistry *HwregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hwregistry.Contract.HwregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hwregistry *HwregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hwregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hwregistry *HwregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hwregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hwregistry *HwregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hwregistry.Contract.contract.Transact(opts, method, params...)
}

// AllWhitelisted is a free data retrieval call binding the contract method 0x765b9d6b.
//
// Solidity: function allWhitelisted() view returns((address,uint256)[])
func (_Hwregistry *HwregistryCaller) AllWhitelisted(opts *bind.CallOpts) ([]HWRegistryWhitelist, error) {
	var out []interface{}
	err := _Hwregistry.contract.Call(opts, &out, "allWhitelisted")

	if err != nil {
		return *new([]HWRegistryWhitelist), err
	}

	out0 := *abi.ConvertType(out[0], new([]HWRegistryWhitelist)).(*[]HWRegistryWhitelist)

	return out0, err

}

// AllWhitelisted is a free data retrieval call binding the contract method 0x765b9d6b.
//
// Solidity: function allWhitelisted() view returns((address,uint256)[])
func (_Hwregistry *HwregistrySession) AllWhitelisted() ([]HWRegistryWhitelist, error) {
	return _Hwregistry.Contract.AllWhitelisted(&_Hwregistry.CallOpts)
}

// AllWhitelisted is a free data retrieval call binding the contract method 0x765b9d6b.
//
// Solidity: function allWhitelisted() view returns((address,uint256)[])
func (_Hwregistry *HwregistryCallerSession) AllWhitelisted() ([]HWRegistryWhitelist, error) {
	return _Hwregistry.Contract.AllWhitelisted(&_Hwregistry.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256 _value)
func (_Hwregistry *HwregistryCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hwregistry.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256 _value)
func (_Hwregistry *HwregistrySession) Counter() (*big.Int, error) {
	return _Hwregistry.Contract.Counter(&_Hwregistry.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256 _value)
func (_Hwregistry *HwregistryCallerSession) Counter() (*big.Int, error) {
	return _Hwregistry.Contract.Counter(&_Hwregistry.CallOpts)
}

// GetNFTGrossRevenue is a free data retrieval call binding the contract method 0x13faf9df.
//
// Solidity: function getNFTGrossRevenue(uint256 _id) view returns(uint256)
func (_Hwregistry *HwregistryCaller) GetNFTGrossRevenue(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwregistry.contract.Call(opts, &out, "getNFTGrossRevenue", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNFTGrossRevenue is a free data retrieval call binding the contract method 0x13faf9df.
//
// Solidity: function getNFTGrossRevenue(uint256 _id) view returns(uint256)
func (_Hwregistry *HwregistrySession) GetNFTGrossRevenue(_id *big.Int) (*big.Int, error) {
	return _Hwregistry.Contract.GetNFTGrossRevenue(&_Hwregistry.CallOpts, _id)
}

// GetNFTGrossRevenue is a free data retrieval call binding the contract method 0x13faf9df.
//
// Solidity: function getNFTGrossRevenue(uint256 _id) view returns(uint256)
func (_Hwregistry *HwregistryCallerSession) GetNFTGrossRevenue(_id *big.Int) (*big.Int, error) {
	return _Hwregistry.Contract.GetNFTGrossRevenue(&_Hwregistry.CallOpts, _id)
}

// HwEscrow is a free data retrieval call binding the contract method 0x24989f6f.
//
// Solidity: function hwEscrow() view returns(address)
func (_Hwregistry *HwregistryCaller) HwEscrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hwregistry.contract.Call(opts, &out, "hwEscrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HwEscrow is a free data retrieval call binding the contract method 0x24989f6f.
//
// Solidity: function hwEscrow() view returns(address)
func (_Hwregistry *HwregistrySession) HwEscrow() (common.Address, error) {
	return _Hwregistry.Contract.HwEscrow(&_Hwregistry.CallOpts)
}

// HwEscrow is a free data retrieval call binding the contract method 0x24989f6f.
//
// Solidity: function hwEscrow() view returns(address)
func (_Hwregistry *HwregistryCallerSession) HwEscrow() (common.Address, error) {
	return _Hwregistry.Contract.HwEscrow(&_Hwregistry.CallOpts)
}

// IsAllowedAmount is a free data retrieval call binding the contract method 0x83ed69c1.
//
// Solidity: function isAllowedAmount(address _address, uint256 _amount) view returns(bool)
func (_Hwregistry *HwregistryCaller) IsAllowedAmount(opts *bind.CallOpts, _address common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _Hwregistry.contract.Call(opts, &out, "isAllowedAmount", _address, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedAmount is a free data retrieval call binding the contract method 0x83ed69c1.
//
// Solidity: function isAllowedAmount(address _address, uint256 _amount) view returns(bool)
func (_Hwregistry *HwregistrySession) IsAllowedAmount(_address common.Address, _amount *big.Int) (bool, error) {
	return _Hwregistry.Contract.IsAllowedAmount(&_Hwregistry.CallOpts, _address, _amount)
}

// IsAllowedAmount is a free data retrieval call binding the contract method 0x83ed69c1.
//
// Solidity: function isAllowedAmount(address _address, uint256 _amount) view returns(bool)
func (_Hwregistry *HwregistryCallerSession) IsAllowedAmount(_address common.Address, _amount *big.Int) (bool, error) {
	return _Hwregistry.Contract.IsAllowedAmount(&_Hwregistry.CallOpts, _address, _amount)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_Hwregistry *HwregistryCaller) IsWhitelisted(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Hwregistry.contract.Call(opts, &out, "isWhitelisted", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_Hwregistry *HwregistrySession) IsWhitelisted(_address common.Address) (bool, error) {
	return _Hwregistry.Contract.IsWhitelisted(&_Hwregistry.CallOpts, _address)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_Hwregistry *HwregistryCallerSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _Hwregistry.Contract.IsWhitelisted(&_Hwregistry.CallOpts, _address)
}

// NftGrossRevenue is a free data retrieval call binding the contract method 0x43120455.
//
// Solidity: function nftGrossRevenue(uint256 ) view returns(uint256)
func (_Hwregistry *HwregistryCaller) NftGrossRevenue(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwregistry.contract.Call(opts, &out, "nftGrossRevenue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftGrossRevenue is a free data retrieval call binding the contract method 0x43120455.
//
// Solidity: function nftGrossRevenue(uint256 ) view returns(uint256)
func (_Hwregistry *HwregistrySession) NftGrossRevenue(arg0 *big.Int) (*big.Int, error) {
	return _Hwregistry.Contract.NftGrossRevenue(&_Hwregistry.CallOpts, arg0)
}

// NftGrossRevenue is a free data retrieval call binding the contract method 0x43120455.
//
// Solidity: function nftGrossRevenue(uint256 ) view returns(uint256)
func (_Hwregistry *HwregistryCallerSession) NftGrossRevenue(arg0 *big.Int) (*big.Int, error) {
	return _Hwregistry.Contract.NftGrossRevenue(&_Hwregistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hwregistry *HwregistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hwregistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hwregistry *HwregistrySession) Owner() (common.Address, error) {
	return _Hwregistry.Contract.Owner(&_Hwregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hwregistry *HwregistryCallerSession) Owner() (common.Address, error) {
	return _Hwregistry.Contract.Owner(&_Hwregistry.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0x3d4efe09.
//
// Solidity: function whitelisted(uint256 ) view returns(address token, uint256 maxAllowed)
func (_Hwregistry *HwregistryCaller) Whitelisted(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token      common.Address
	MaxAllowed *big.Int
}, error) {
	var out []interface{}
	err := _Hwregistry.contract.Call(opts, &out, "whitelisted", arg0)

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
func (_Hwregistry *HwregistrySession) Whitelisted(arg0 *big.Int) (struct {
	Token      common.Address
	MaxAllowed *big.Int
}, error) {
	return _Hwregistry.Contract.Whitelisted(&_Hwregistry.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0x3d4efe09.
//
// Solidity: function whitelisted(uint256 ) view returns(address token, uint256 maxAllowed)
func (_Hwregistry *HwregistryCallerSession) Whitelisted(arg0 *big.Int) (struct {
	Token      common.Address
	MaxAllowed *big.Int
}, error) {
	return _Hwregistry.Contract.Whitelisted(&_Hwregistry.CallOpts, arg0)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x8902ff86.
//
// Solidity: function addWhitelisted(address _address, uint256 _maxAllowed) returns(bool)
func (_Hwregistry *HwregistryTransactor) AddWhitelisted(opts *bind.TransactOpts, _address common.Address, _maxAllowed *big.Int) (*types.Transaction, error) {
	return _Hwregistry.contract.Transact(opts, "addWhitelisted", _address, _maxAllowed)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x8902ff86.
//
// Solidity: function addWhitelisted(address _address, uint256 _maxAllowed) returns(bool)
func (_Hwregistry *HwregistrySession) AddWhitelisted(_address common.Address, _maxAllowed *big.Int) (*types.Transaction, error) {
	return _Hwregistry.Contract.AddWhitelisted(&_Hwregistry.TransactOpts, _address, _maxAllowed)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x8902ff86.
//
// Solidity: function addWhitelisted(address _address, uint256 _maxAllowed) returns(bool)
func (_Hwregistry *HwregistryTransactorSession) AddWhitelisted(_address common.Address, _maxAllowed *big.Int) (*types.Transaction, error) {
	return _Hwregistry.Contract.AddWhitelisted(&_Hwregistry.TransactOpts, _address, _maxAllowed)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns(bool)
func (_Hwregistry *HwregistryTransactor) RemoveWhitelisted(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Hwregistry.contract.Transact(opts, "removeWhitelisted", _address)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns(bool)
func (_Hwregistry *HwregistrySession) RemoveWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Hwregistry.Contract.RemoveWhitelisted(&_Hwregistry.TransactOpts, _address)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns(bool)
func (_Hwregistry *HwregistryTransactorSession) RemoveWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Hwregistry.Contract.RemoveWhitelisted(&_Hwregistry.TransactOpts, _address)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hwregistry *HwregistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hwregistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hwregistry *HwregistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Hwregistry.Contract.RenounceOwnership(&_Hwregistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hwregistry *HwregistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hwregistry.Contract.RenounceOwnership(&_Hwregistry.TransactOpts)
}

// SetHWEscrow is a paid mutator transaction binding the contract method 0xf5f55714.
//
// Solidity: function setHWEscrow(address _address) returns(bool)
func (_Hwregistry *HwregistryTransactor) SetHWEscrow(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Hwregistry.contract.Transact(opts, "setHWEscrow", _address)
}

// SetHWEscrow is a paid mutator transaction binding the contract method 0xf5f55714.
//
// Solidity: function setHWEscrow(address _address) returns(bool)
func (_Hwregistry *HwregistrySession) SetHWEscrow(_address common.Address) (*types.Transaction, error) {
	return _Hwregistry.Contract.SetHWEscrow(&_Hwregistry.TransactOpts, _address)
}

// SetHWEscrow is a paid mutator transaction binding the contract method 0xf5f55714.
//
// Solidity: function setHWEscrow(address _address) returns(bool)
func (_Hwregistry *HwregistryTransactorSession) SetHWEscrow(_address common.Address) (*types.Transaction, error) {
	return _Hwregistry.Contract.SetHWEscrow(&_Hwregistry.TransactOpts, _address)
}

// SetNFTGrossRevenue is a paid mutator transaction binding the contract method 0x155e85c4.
//
// Solidity: function setNFTGrossRevenue(uint256 _id, uint256 _amount) returns()
func (_Hwregistry *HwregistryTransactor) SetNFTGrossRevenue(opts *bind.TransactOpts, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Hwregistry.contract.Transact(opts, "setNFTGrossRevenue", _id, _amount)
}

// SetNFTGrossRevenue is a paid mutator transaction binding the contract method 0x155e85c4.
//
// Solidity: function setNFTGrossRevenue(uint256 _id, uint256 _amount) returns()
func (_Hwregistry *HwregistrySession) SetNFTGrossRevenue(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Hwregistry.Contract.SetNFTGrossRevenue(&_Hwregistry.TransactOpts, _id, _amount)
}

// SetNFTGrossRevenue is a paid mutator transaction binding the contract method 0x155e85c4.
//
// Solidity: function setNFTGrossRevenue(uint256 _id, uint256 _amount) returns()
func (_Hwregistry *HwregistryTransactorSession) SetNFTGrossRevenue(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Hwregistry.Contract.SetNFTGrossRevenue(&_Hwregistry.TransactOpts, _id, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hwregistry *HwregistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Hwregistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hwregistry *HwregistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hwregistry.Contract.TransferOwnership(&_Hwregistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hwregistry *HwregistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hwregistry.Contract.TransferOwnership(&_Hwregistry.TransactOpts, newOwner)
}

// UpdateWhitelisted is a paid mutator transaction binding the contract method 0xcd1188a5.
//
// Solidity: function updateWhitelisted(address _address, uint256 _maxAllowed) returns(bool)
func (_Hwregistry *HwregistryTransactor) UpdateWhitelisted(opts *bind.TransactOpts, _address common.Address, _maxAllowed *big.Int) (*types.Transaction, error) {
	return _Hwregistry.contract.Transact(opts, "updateWhitelisted", _address, _maxAllowed)
}

// UpdateWhitelisted is a paid mutator transaction binding the contract method 0xcd1188a5.
//
// Solidity: function updateWhitelisted(address _address, uint256 _maxAllowed) returns(bool)
func (_Hwregistry *HwregistrySession) UpdateWhitelisted(_address common.Address, _maxAllowed *big.Int) (*types.Transaction, error) {
	return _Hwregistry.Contract.UpdateWhitelisted(&_Hwregistry.TransactOpts, _address, _maxAllowed)
}

// UpdateWhitelisted is a paid mutator transaction binding the contract method 0xcd1188a5.
//
// Solidity: function updateWhitelisted(address _address, uint256 _maxAllowed) returns(bool)
func (_Hwregistry *HwregistryTransactorSession) UpdateWhitelisted(_address common.Address, _maxAllowed *big.Int) (*types.Transaction, error) {
	return _Hwregistry.Contract.UpdateWhitelisted(&_Hwregistry.TransactOpts, _address, _maxAllowed)
}

// HwregistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Hwregistry contract.
type HwregistryOwnershipTransferredIterator struct {
	Event *HwregistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HwregistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwregistryOwnershipTransferred)
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
		it.Event = new(HwregistryOwnershipTransferred)
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
func (it *HwregistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwregistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwregistryOwnershipTransferred represents a OwnershipTransferred event raised by the Hwregistry contract.
type HwregistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hwregistry *HwregistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HwregistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hwregistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HwregistryOwnershipTransferredIterator{contract: _Hwregistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hwregistry *HwregistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HwregistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hwregistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwregistryOwnershipTransferred)
				if err := _Hwregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Hwregistry *HwregistryFilterer) ParseOwnershipTransferred(log types.Log) (*HwregistryOwnershipTransferred, error) {
	event := new(HwregistryOwnershipTransferred)
	if err := _Hwregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwregistryWhitelistedAddedIterator is returned from FilterWhitelistedAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAdded events raised by the Hwregistry contract.
type HwregistryWhitelistedAddedIterator struct {
	Event *HwregistryWhitelistedAdded // Event containing the contract specifics and raw log

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
func (it *HwregistryWhitelistedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwregistryWhitelistedAdded)
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
		it.Event = new(HwregistryWhitelistedAdded)
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
func (it *HwregistryWhitelistedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwregistryWhitelistedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwregistryWhitelistedAdded represents a WhitelistedAdded event raised by the Hwregistry contract.
type HwregistryWhitelistedAdded struct {
	Address    common.Address
	MaxAllowed *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAdded is a free log retrieval operation binding the contract event 0xcb9ce803cb004a01f0682efa32d66c42ec72e8daa764006c95e1b551b9465a21.
//
// Solidity: event WhitelistedAdded(address indexed _address, uint256 _maxAllowed)
func (_Hwregistry *HwregistryFilterer) FilterWhitelistedAdded(opts *bind.FilterOpts, _address []common.Address) (*HwregistryWhitelistedAddedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Hwregistry.contract.FilterLogs(opts, "WhitelistedAdded", _addressRule)
	if err != nil {
		return nil, err
	}
	return &HwregistryWhitelistedAddedIterator{contract: _Hwregistry.contract, event: "WhitelistedAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAdded is a free log subscription operation binding the contract event 0xcb9ce803cb004a01f0682efa32d66c42ec72e8daa764006c95e1b551b9465a21.
//
// Solidity: event WhitelistedAdded(address indexed _address, uint256 _maxAllowed)
func (_Hwregistry *HwregistryFilterer) WatchWhitelistedAdded(opts *bind.WatchOpts, sink chan<- *HwregistryWhitelistedAdded, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Hwregistry.contract.WatchLogs(opts, "WhitelistedAdded", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwregistryWhitelistedAdded)
				if err := _Hwregistry.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
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
func (_Hwregistry *HwregistryFilterer) ParseWhitelistedAdded(log types.Log) (*HwregistryWhitelistedAdded, error) {
	event := new(HwregistryWhitelistedAdded)
	if err := _Hwregistry.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwregistryWhitelistedRemovedIterator is returned from FilterWhitelistedRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedRemoved events raised by the Hwregistry contract.
type HwregistryWhitelistedRemovedIterator struct {
	Event *HwregistryWhitelistedRemoved // Event containing the contract specifics and raw log

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
func (it *HwregistryWhitelistedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwregistryWhitelistedRemoved)
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
		it.Event = new(HwregistryWhitelistedRemoved)
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
func (it *HwregistryWhitelistedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwregistryWhitelistedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwregistryWhitelistedRemoved represents a WhitelistedRemoved event raised by the Hwregistry contract.
type HwregistryWhitelistedRemoved struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedRemoved is a free log retrieval operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address indexed _address)
func (_Hwregistry *HwregistryFilterer) FilterWhitelistedRemoved(opts *bind.FilterOpts, _address []common.Address) (*HwregistryWhitelistedRemovedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Hwregistry.contract.FilterLogs(opts, "WhitelistedRemoved", _addressRule)
	if err != nil {
		return nil, err
	}
	return &HwregistryWhitelistedRemovedIterator{contract: _Hwregistry.contract, event: "WhitelistedRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedRemoved is a free log subscription operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address indexed _address)
func (_Hwregistry *HwregistryFilterer) WatchWhitelistedRemoved(opts *bind.WatchOpts, sink chan<- *HwregistryWhitelistedRemoved, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Hwregistry.contract.WatchLogs(opts, "WhitelistedRemoved", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwregistryWhitelistedRemoved)
				if err := _Hwregistry.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
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
func (_Hwregistry *HwregistryFilterer) ParseWhitelistedRemoved(log types.Log) (*HwregistryWhitelistedRemoved, error) {
	event := new(HwregistryWhitelistedRemoved)
	if err := _Hwregistry.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwregistryWhitelistedUpdatedIterator is returned from FilterWhitelistedUpdated and is used to iterate over the raw logs and unpacked data for WhitelistedUpdated events raised by the Hwregistry contract.
type HwregistryWhitelistedUpdatedIterator struct {
	Event *HwregistryWhitelistedUpdated // Event containing the contract specifics and raw log

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
func (it *HwregistryWhitelistedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwregistryWhitelistedUpdated)
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
		it.Event = new(HwregistryWhitelistedUpdated)
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
func (it *HwregistryWhitelistedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwregistryWhitelistedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwregistryWhitelistedUpdated represents a WhitelistedUpdated event raised by the Hwregistry contract.
type HwregistryWhitelistedUpdated struct {
	Address    common.Address
	MaxAllowed *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedUpdated is a free log retrieval operation binding the contract event 0x8bc3b900d96200ab851acb4ccd80670a931934f8fbae20691e95a683c0d00e60.
//
// Solidity: event WhitelistedUpdated(address indexed _address, uint256 _maxAllowed)
func (_Hwregistry *HwregistryFilterer) FilterWhitelistedUpdated(opts *bind.FilterOpts, _address []common.Address) (*HwregistryWhitelistedUpdatedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Hwregistry.contract.FilterLogs(opts, "WhitelistedUpdated", _addressRule)
	if err != nil {
		return nil, err
	}
	return &HwregistryWhitelistedUpdatedIterator{contract: _Hwregistry.contract, event: "WhitelistedUpdated", logs: logs, sub: sub}, nil
}

// WatchWhitelistedUpdated is a free log subscription operation binding the contract event 0x8bc3b900d96200ab851acb4ccd80670a931934f8fbae20691e95a683c0d00e60.
//
// Solidity: event WhitelistedUpdated(address indexed _address, uint256 _maxAllowed)
func (_Hwregistry *HwregistryFilterer) WatchWhitelistedUpdated(opts *bind.WatchOpts, sink chan<- *HwregistryWhitelistedUpdated, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Hwregistry.contract.WatchLogs(opts, "WhitelistedUpdated", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwregistryWhitelistedUpdated)
				if err := _Hwregistry.contract.UnpackLog(event, "WhitelistedUpdated", log); err != nil {
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
func (_Hwregistry *HwregistryFilterer) ParseWhitelistedUpdated(log types.Log) (*HwregistryWhitelistedUpdated, error) {
	event := new(HwregistryWhitelistedUpdated)
	if err := _Hwregistry.contract.UnpackLog(event, "WhitelistedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
