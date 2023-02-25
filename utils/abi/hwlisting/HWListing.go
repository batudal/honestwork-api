// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hwlisting

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

// HWListingPayment is an auto generated low-level Go binding around an user-defined struct.
type HWListingPayment struct {
	Token       common.Address
	Amount      *big.Int
	ListingDate *big.Int
}

// HwlistingMetaData contains all meta data concerning the Hwlisting contract.
var HwlistingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stableCoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"PaymentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"PaymentAddedETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getEthPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getLatestPayment\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingDate\",\"type\":\"uint256\"}],\"internalType\":\"structHWListing.Payment\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getPayments\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingDate\",\"type\":\"uint256\"}],\"internalType\":\"structHWListing.Payment[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payForListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minTokensOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allowedDelay\",\"type\":\"uint256\"}],\"name\":\"payForListingEth\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIHWRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router01\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableCoin\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"updateRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawAllEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllEarningsEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HwlistingABI is the input ABI used to generate the binding from.
// Deprecated: Use HwlistingMetaData.ABI instead.
var HwlistingABI = HwlistingMetaData.ABI

// Hwlisting is an auto generated Go binding around an Ethereum contract.
type Hwlisting struct {
	HwlistingCaller     // Read-only binding to the contract
	HwlistingTransactor // Write-only binding to the contract
	HwlistingFilterer   // Log filterer for contract events
}

// HwlistingCaller is an auto generated read-only Go binding around an Ethereum contract.
type HwlistingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HwlistingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HwlistingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HwlistingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HwlistingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HwlistingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HwlistingSession struct {
	Contract     *Hwlisting        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HwlistingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HwlistingCallerSession struct {
	Contract *HwlistingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HwlistingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HwlistingTransactorSession struct {
	Contract     *HwlistingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HwlistingRaw is an auto generated low-level Go binding around an Ethereum contract.
type HwlistingRaw struct {
	Contract *Hwlisting // Generic contract binding to access the raw methods on
}

// HwlistingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HwlistingCallerRaw struct {
	Contract *HwlistingCaller // Generic read-only contract binding to access the raw methods on
}

// HwlistingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HwlistingTransactorRaw struct {
	Contract *HwlistingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHwlisting creates a new instance of Hwlisting, bound to a specific deployed contract.
func NewHwlisting(address common.Address, backend bind.ContractBackend) (*Hwlisting, error) {
	contract, err := bindHwlisting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hwlisting{HwlistingCaller: HwlistingCaller{contract: contract}, HwlistingTransactor: HwlistingTransactor{contract: contract}, HwlistingFilterer: HwlistingFilterer{contract: contract}}, nil
}

// NewHwlistingCaller creates a new read-only instance of Hwlisting, bound to a specific deployed contract.
func NewHwlistingCaller(address common.Address, caller bind.ContractCaller) (*HwlistingCaller, error) {
	contract, err := bindHwlisting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HwlistingCaller{contract: contract}, nil
}

// NewHwlistingTransactor creates a new write-only instance of Hwlisting, bound to a specific deployed contract.
func NewHwlistingTransactor(address common.Address, transactor bind.ContractTransactor) (*HwlistingTransactor, error) {
	contract, err := bindHwlisting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HwlistingTransactor{contract: contract}, nil
}

// NewHwlistingFilterer creates a new log filterer instance of Hwlisting, bound to a specific deployed contract.
func NewHwlistingFilterer(address common.Address, filterer bind.ContractFilterer) (*HwlistingFilterer, error) {
	contract, err := bindHwlisting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HwlistingFilterer{contract: contract}, nil
}

// bindHwlisting binds a generic wrapper to an already deployed contract.
func bindHwlisting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HwlistingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hwlisting *HwlistingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hwlisting.Contract.HwlistingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hwlisting *HwlistingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hwlisting.Contract.HwlistingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hwlisting *HwlistingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hwlisting.Contract.HwlistingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hwlisting *HwlistingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hwlisting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hwlisting *HwlistingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hwlisting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hwlisting *HwlistingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hwlisting.Contract.contract.Transact(opts, method, params...)
}

// GetEthPrice is a free data retrieval call binding the contract method 0x870d365d.
//
// Solidity: function getEthPrice(uint256 _amount) view returns(uint256)
func (_Hwlisting *HwlistingCaller) GetEthPrice(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwlisting.contract.Call(opts, &out, "getEthPrice", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthPrice is a free data retrieval call binding the contract method 0x870d365d.
//
// Solidity: function getEthPrice(uint256 _amount) view returns(uint256)
func (_Hwlisting *HwlistingSession) GetEthPrice(_amount *big.Int) (*big.Int, error) {
	return _Hwlisting.Contract.GetEthPrice(&_Hwlisting.CallOpts, _amount)
}

// GetEthPrice is a free data retrieval call binding the contract method 0x870d365d.
//
// Solidity: function getEthPrice(uint256 _amount) view returns(uint256)
func (_Hwlisting *HwlistingCallerSession) GetEthPrice(_amount *big.Int) (*big.Int, error) {
	return _Hwlisting.Contract.GetEthPrice(&_Hwlisting.CallOpts, _amount)
}

// GetLatestPayment is a free data retrieval call binding the contract method 0xa12a0365.
//
// Solidity: function getLatestPayment(address _user) view returns((address,uint256,uint256))
func (_Hwlisting *HwlistingCaller) GetLatestPayment(opts *bind.CallOpts, _user common.Address) (HWListingPayment, error) {
	var out []interface{}
	err := _Hwlisting.contract.Call(opts, &out, "getLatestPayment", _user)

	if err != nil {
		return *new(HWListingPayment), err
	}

	out0 := *abi.ConvertType(out[0], new(HWListingPayment)).(*HWListingPayment)

	return out0, err

}

// GetLatestPayment is a free data retrieval call binding the contract method 0xa12a0365.
//
// Solidity: function getLatestPayment(address _user) view returns((address,uint256,uint256))
func (_Hwlisting *HwlistingSession) GetLatestPayment(_user common.Address) (HWListingPayment, error) {
	return _Hwlisting.Contract.GetLatestPayment(&_Hwlisting.CallOpts, _user)
}

// GetLatestPayment is a free data retrieval call binding the contract method 0xa12a0365.
//
// Solidity: function getLatestPayment(address _user) view returns((address,uint256,uint256))
func (_Hwlisting *HwlistingCallerSession) GetLatestPayment(_user common.Address) (HWListingPayment, error) {
	return _Hwlisting.Contract.GetLatestPayment(&_Hwlisting.CallOpts, _user)
}

// GetPayments is a free data retrieval call binding the contract method 0x1d6a1711.
//
// Solidity: function getPayments(address _user) view returns((address,uint256,uint256)[])
func (_Hwlisting *HwlistingCaller) GetPayments(opts *bind.CallOpts, _user common.Address) ([]HWListingPayment, error) {
	var out []interface{}
	err := _Hwlisting.contract.Call(opts, &out, "getPayments", _user)

	if err != nil {
		return *new([]HWListingPayment), err
	}

	out0 := *abi.ConvertType(out[0], new([]HWListingPayment)).(*[]HWListingPayment)

	return out0, err

}

// GetPayments is a free data retrieval call binding the contract method 0x1d6a1711.
//
// Solidity: function getPayments(address _user) view returns((address,uint256,uint256)[])
func (_Hwlisting *HwlistingSession) GetPayments(_user common.Address) ([]HWListingPayment, error) {
	return _Hwlisting.Contract.GetPayments(&_Hwlisting.CallOpts, _user)
}

// GetPayments is a free data retrieval call binding the contract method 0x1d6a1711.
//
// Solidity: function getPayments(address _user) view returns((address,uint256,uint256)[])
func (_Hwlisting *HwlistingCallerSession) GetPayments(_user common.Address) ([]HWListingPayment, error) {
	return _Hwlisting.Contract.GetPayments(&_Hwlisting.CallOpts, _user)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x82b2e257.
//
// Solidity: function getTokenBalance() view returns(uint256)
func (_Hwlisting *HwlistingCaller) GetTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hwlisting.contract.Call(opts, &out, "getTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenBalance is a free data retrieval call binding the contract method 0x82b2e257.
//
// Solidity: function getTokenBalance() view returns(uint256)
func (_Hwlisting *HwlistingSession) GetTokenBalance() (*big.Int, error) {
	return _Hwlisting.Contract.GetTokenBalance(&_Hwlisting.CallOpts)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x82b2e257.
//
// Solidity: function getTokenBalance() view returns(uint256)
func (_Hwlisting *HwlistingCallerSession) GetTokenBalance() (*big.Int, error) {
	return _Hwlisting.Contract.GetTokenBalance(&_Hwlisting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hwlisting *HwlistingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hwlisting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hwlisting *HwlistingSession) Owner() (common.Address, error) {
	return _Hwlisting.Contract.Owner(&_Hwlisting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hwlisting *HwlistingCallerSession) Owner() (common.Address, error) {
	return _Hwlisting.Contract.Owner(&_Hwlisting.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Hwlisting *HwlistingCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hwlisting.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Hwlisting *HwlistingSession) Pool() (common.Address, error) {
	return _Hwlisting.Contract.Pool(&_Hwlisting.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Hwlisting *HwlistingCallerSession) Pool() (common.Address, error) {
	return _Hwlisting.Contract.Pool(&_Hwlisting.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Hwlisting *HwlistingCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hwlisting.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Hwlisting *HwlistingSession) Registry() (common.Address, error) {
	return _Hwlisting.Contract.Registry(&_Hwlisting.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Hwlisting *HwlistingCallerSession) Registry() (common.Address, error) {
	return _Hwlisting.Contract.Registry(&_Hwlisting.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Hwlisting *HwlistingCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hwlisting.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Hwlisting *HwlistingSession) Router() (common.Address, error) {
	return _Hwlisting.Contract.Router(&_Hwlisting.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Hwlisting *HwlistingCallerSession) Router() (common.Address, error) {
	return _Hwlisting.Contract.Router(&_Hwlisting.CallOpts)
}

// StableCoin is a free data retrieval call binding the contract method 0x992642e5.
//
// Solidity: function stableCoin() view returns(address)
func (_Hwlisting *HwlistingCaller) StableCoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hwlisting.contract.Call(opts, &out, "stableCoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StableCoin is a free data retrieval call binding the contract method 0x992642e5.
//
// Solidity: function stableCoin() view returns(address)
func (_Hwlisting *HwlistingSession) StableCoin() (common.Address, error) {
	return _Hwlisting.Contract.StableCoin(&_Hwlisting.CallOpts)
}

// StableCoin is a free data retrieval call binding the contract method 0x992642e5.
//
// Solidity: function stableCoin() view returns(address)
func (_Hwlisting *HwlistingCallerSession) StableCoin() (common.Address, error) {
	return _Hwlisting.Contract.StableCoin(&_Hwlisting.CallOpts)
}

// PayForListing is a paid mutator transaction binding the contract method 0x304db7d1.
//
// Solidity: function payForListing(address _token, uint256 _amount) returns()
func (_Hwlisting *HwlistingTransactor) PayForListing(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Hwlisting.contract.Transact(opts, "payForListing", _token, _amount)
}

// PayForListing is a paid mutator transaction binding the contract method 0x304db7d1.
//
// Solidity: function payForListing(address _token, uint256 _amount) returns()
func (_Hwlisting *HwlistingSession) PayForListing(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Hwlisting.Contract.PayForListing(&_Hwlisting.TransactOpts, _token, _amount)
}

// PayForListing is a paid mutator transaction binding the contract method 0x304db7d1.
//
// Solidity: function payForListing(address _token, uint256 _amount) returns()
func (_Hwlisting *HwlistingTransactorSession) PayForListing(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Hwlisting.Contract.PayForListing(&_Hwlisting.TransactOpts, _token, _amount)
}

// PayForListingEth is a paid mutator transaction binding the contract method 0xc99b6e99.
//
// Solidity: function payForListingEth(uint256 _minTokensOut, uint256 _allowedDelay) payable returns(uint256[])
func (_Hwlisting *HwlistingTransactor) PayForListingEth(opts *bind.TransactOpts, _minTokensOut *big.Int, _allowedDelay *big.Int) (*types.Transaction, error) {
	return _Hwlisting.contract.Transact(opts, "payForListingEth", _minTokensOut, _allowedDelay)
}

// PayForListingEth is a paid mutator transaction binding the contract method 0xc99b6e99.
//
// Solidity: function payForListingEth(uint256 _minTokensOut, uint256 _allowedDelay) payable returns(uint256[])
func (_Hwlisting *HwlistingSession) PayForListingEth(_minTokensOut *big.Int, _allowedDelay *big.Int) (*types.Transaction, error) {
	return _Hwlisting.Contract.PayForListingEth(&_Hwlisting.TransactOpts, _minTokensOut, _allowedDelay)
}

// PayForListingEth is a paid mutator transaction binding the contract method 0xc99b6e99.
//
// Solidity: function payForListingEth(uint256 _minTokensOut, uint256 _allowedDelay) payable returns(uint256[])
func (_Hwlisting *HwlistingTransactorSession) PayForListingEth(_minTokensOut *big.Int, _allowedDelay *big.Int) (*types.Transaction, error) {
	return _Hwlisting.Contract.PayForListingEth(&_Hwlisting.TransactOpts, _minTokensOut, _allowedDelay)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hwlisting *HwlistingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hwlisting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hwlisting *HwlistingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hwlisting.Contract.RenounceOwnership(&_Hwlisting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hwlisting *HwlistingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hwlisting.Contract.RenounceOwnership(&_Hwlisting.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hwlisting *HwlistingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Hwlisting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hwlisting *HwlistingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hwlisting.Contract.TransferOwnership(&_Hwlisting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hwlisting *HwlistingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hwlisting.Contract.TransferOwnership(&_Hwlisting.TransactOpts, newOwner)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x1a5da6c8.
//
// Solidity: function updateRegistry(address _registry) returns()
func (_Hwlisting *HwlistingTransactor) UpdateRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _Hwlisting.contract.Transact(opts, "updateRegistry", _registry)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x1a5da6c8.
//
// Solidity: function updateRegistry(address _registry) returns()
func (_Hwlisting *HwlistingSession) UpdateRegistry(_registry common.Address) (*types.Transaction, error) {
	return _Hwlisting.Contract.UpdateRegistry(&_Hwlisting.TransactOpts, _registry)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x1a5da6c8.
//
// Solidity: function updateRegistry(address _registry) returns()
func (_Hwlisting *HwlistingTransactorSession) UpdateRegistry(_registry common.Address) (*types.Transaction, error) {
	return _Hwlisting.Contract.UpdateRegistry(&_Hwlisting.TransactOpts, _registry)
}

// WithdrawAllEarnings is a paid mutator transaction binding the contract method 0xc4218374.
//
// Solidity: function withdrawAllEarnings(address _token) returns()
func (_Hwlisting *HwlistingTransactor) WithdrawAllEarnings(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Hwlisting.contract.Transact(opts, "withdrawAllEarnings", _token)
}

// WithdrawAllEarnings is a paid mutator transaction binding the contract method 0xc4218374.
//
// Solidity: function withdrawAllEarnings(address _token) returns()
func (_Hwlisting *HwlistingSession) WithdrawAllEarnings(_token common.Address) (*types.Transaction, error) {
	return _Hwlisting.Contract.WithdrawAllEarnings(&_Hwlisting.TransactOpts, _token)
}

// WithdrawAllEarnings is a paid mutator transaction binding the contract method 0xc4218374.
//
// Solidity: function withdrawAllEarnings(address _token) returns()
func (_Hwlisting *HwlistingTransactorSession) WithdrawAllEarnings(_token common.Address) (*types.Transaction, error) {
	return _Hwlisting.Contract.WithdrawAllEarnings(&_Hwlisting.TransactOpts, _token)
}

// WithdrawAllEarningsEth is a paid mutator transaction binding the contract method 0xc7c067af.
//
// Solidity: function withdrawAllEarningsEth() returns()
func (_Hwlisting *HwlistingTransactor) WithdrawAllEarningsEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hwlisting.contract.Transact(opts, "withdrawAllEarningsEth")
}

// WithdrawAllEarningsEth is a paid mutator transaction binding the contract method 0xc7c067af.
//
// Solidity: function withdrawAllEarningsEth() returns()
func (_Hwlisting *HwlistingSession) WithdrawAllEarningsEth() (*types.Transaction, error) {
	return _Hwlisting.Contract.WithdrawAllEarningsEth(&_Hwlisting.TransactOpts)
}

// WithdrawAllEarningsEth is a paid mutator transaction binding the contract method 0xc7c067af.
//
// Solidity: function withdrawAllEarningsEth() returns()
func (_Hwlisting *HwlistingTransactorSession) WithdrawAllEarningsEth() (*types.Transaction, error) {
	return _Hwlisting.Contract.WithdrawAllEarningsEth(&_Hwlisting.TransactOpts)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0x280da6fa.
//
// Solidity: function withdrawAllTokens() returns()
func (_Hwlisting *HwlistingTransactor) WithdrawAllTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hwlisting.contract.Transact(opts, "withdrawAllTokens")
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0x280da6fa.
//
// Solidity: function withdrawAllTokens() returns()
func (_Hwlisting *HwlistingSession) WithdrawAllTokens() (*types.Transaction, error) {
	return _Hwlisting.Contract.WithdrawAllTokens(&_Hwlisting.TransactOpts)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0x280da6fa.
//
// Solidity: function withdrawAllTokens() returns()
func (_Hwlisting *HwlistingTransactorSession) WithdrawAllTokens() (*types.Transaction, error) {
	return _Hwlisting.Contract.WithdrawAllTokens(&_Hwlisting.TransactOpts)
}

// WithdrawEarnings is a paid mutator transaction binding the contract method 0xc9211c82.
//
// Solidity: function withdrawEarnings(address _token, uint256 _amount) returns()
func (_Hwlisting *HwlistingTransactor) WithdrawEarnings(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Hwlisting.contract.Transact(opts, "withdrawEarnings", _token, _amount)
}

// WithdrawEarnings is a paid mutator transaction binding the contract method 0xc9211c82.
//
// Solidity: function withdrawEarnings(address _token, uint256 _amount) returns()
func (_Hwlisting *HwlistingSession) WithdrawEarnings(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Hwlisting.Contract.WithdrawEarnings(&_Hwlisting.TransactOpts, _token, _amount)
}

// WithdrawEarnings is a paid mutator transaction binding the contract method 0xc9211c82.
//
// Solidity: function withdrawEarnings(address _token, uint256 _amount) returns()
func (_Hwlisting *HwlistingTransactorSession) WithdrawEarnings(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Hwlisting.Contract.WithdrawEarnings(&_Hwlisting.TransactOpts, _token, _amount)
}

// HwlistingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Hwlisting contract.
type HwlistingOwnershipTransferredIterator struct {
	Event *HwlistingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HwlistingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwlistingOwnershipTransferred)
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
		it.Event = new(HwlistingOwnershipTransferred)
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
func (it *HwlistingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwlistingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwlistingOwnershipTransferred represents a OwnershipTransferred event raised by the Hwlisting contract.
type HwlistingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hwlisting *HwlistingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HwlistingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hwlisting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HwlistingOwnershipTransferredIterator{contract: _Hwlisting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hwlisting *HwlistingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HwlistingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hwlisting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwlistingOwnershipTransferred)
				if err := _Hwlisting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Hwlisting *HwlistingFilterer) ParseOwnershipTransferred(log types.Log) (*HwlistingOwnershipTransferred, error) {
	event := new(HwlistingOwnershipTransferred)
	if err := _Hwlisting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwlistingPaymentAddedIterator is returned from FilterPaymentAdded and is used to iterate over the raw logs and unpacked data for PaymentAdded events raised by the Hwlisting contract.
type HwlistingPaymentAddedIterator struct {
	Event *HwlistingPaymentAdded // Event containing the contract specifics and raw log

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
func (it *HwlistingPaymentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwlistingPaymentAdded)
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
		it.Event = new(HwlistingPaymentAdded)
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
func (it *HwlistingPaymentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwlistingPaymentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwlistingPaymentAdded represents a PaymentAdded event raised by the Hwlisting contract.
type HwlistingPaymentAdded struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentAdded is a free log retrieval operation binding the contract event 0x7ceb32c988f6127615957ec467cbb5688c54a860a192790b125e54b9c3b681e0.
//
// Solidity: event PaymentAdded(address indexed _token, uint256 _amount)
func (_Hwlisting *HwlistingFilterer) FilterPaymentAdded(opts *bind.FilterOpts, _token []common.Address) (*HwlistingPaymentAddedIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Hwlisting.contract.FilterLogs(opts, "PaymentAdded", _tokenRule)
	if err != nil {
		return nil, err
	}
	return &HwlistingPaymentAddedIterator{contract: _Hwlisting.contract, event: "PaymentAdded", logs: logs, sub: sub}, nil
}

// WatchPaymentAdded is a free log subscription operation binding the contract event 0x7ceb32c988f6127615957ec467cbb5688c54a860a192790b125e54b9c3b681e0.
//
// Solidity: event PaymentAdded(address indexed _token, uint256 _amount)
func (_Hwlisting *HwlistingFilterer) WatchPaymentAdded(opts *bind.WatchOpts, sink chan<- *HwlistingPaymentAdded, _token []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Hwlisting.contract.WatchLogs(opts, "PaymentAdded", _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwlistingPaymentAdded)
				if err := _Hwlisting.contract.UnpackLog(event, "PaymentAdded", log); err != nil {
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
func (_Hwlisting *HwlistingFilterer) ParsePaymentAdded(log types.Log) (*HwlistingPaymentAdded, error) {
	event := new(HwlistingPaymentAdded)
	if err := _Hwlisting.contract.UnpackLog(event, "PaymentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwlistingPaymentAddedETHIterator is returned from FilterPaymentAddedETH and is used to iterate over the raw logs and unpacked data for PaymentAddedETH events raised by the Hwlisting contract.
type HwlistingPaymentAddedETHIterator struct {
	Event *HwlistingPaymentAddedETH // Event containing the contract specifics and raw log

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
func (it *HwlistingPaymentAddedETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwlistingPaymentAddedETH)
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
		it.Event = new(HwlistingPaymentAddedETH)
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
func (it *HwlistingPaymentAddedETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwlistingPaymentAddedETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwlistingPaymentAddedETH represents a PaymentAddedETH event raised by the Hwlisting contract.
type HwlistingPaymentAddedETH struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentAddedETH is a free log retrieval operation binding the contract event 0xf8b14910ac2e6a8e7967441ce0230a43990764e0c669d0681e91e42749cdb18b.
//
// Solidity: event PaymentAddedETH(uint256 _amount)
func (_Hwlisting *HwlistingFilterer) FilterPaymentAddedETH(opts *bind.FilterOpts) (*HwlistingPaymentAddedETHIterator, error) {

	logs, sub, err := _Hwlisting.contract.FilterLogs(opts, "PaymentAddedETH")
	if err != nil {
		return nil, err
	}
	return &HwlistingPaymentAddedETHIterator{contract: _Hwlisting.contract, event: "PaymentAddedETH", logs: logs, sub: sub}, nil
}

// WatchPaymentAddedETH is a free log subscription operation binding the contract event 0xf8b14910ac2e6a8e7967441ce0230a43990764e0c669d0681e91e42749cdb18b.
//
// Solidity: event PaymentAddedETH(uint256 _amount)
func (_Hwlisting *HwlistingFilterer) WatchPaymentAddedETH(opts *bind.WatchOpts, sink chan<- *HwlistingPaymentAddedETH) (event.Subscription, error) {

	logs, sub, err := _Hwlisting.contract.WatchLogs(opts, "PaymentAddedETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwlistingPaymentAddedETH)
				if err := _Hwlisting.contract.UnpackLog(event, "PaymentAddedETH", log); err != nil {
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
func (_Hwlisting *HwlistingFilterer) ParsePaymentAddedETH(log types.Log) (*HwlistingPaymentAddedETH, error) {
	event := new(HwlistingPaymentAddedETH)
	if err := _Hwlisting.contract.UnpackLog(event, "PaymentAddedETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
