// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package honestworknft

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

// HonestworknftMetaData contains all meta data concerning the Honestworknft contract.
var HonestworknftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseUri\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_whitelistedTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tier\",\"type\":\"uint256\"}],\"name\":\"Upgrade\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_tokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tier\",\"type\":\"uint256\"}],\"name\":\"adminMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenTier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserTier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getWhitelistToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"publicMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"removeWhitelistToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseUri\",\"type\":\"string\"}],\"name\":\"setBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tierOneFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tierTwoFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tierThreeFee\",\"type\":\"uint256\"}],\"name\":\"setTiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenCap\",\"type\":\"uint256\"}],\"name\":\"setTokenCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"setWhitelistRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierOneFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierThreeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierTwoFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_levels\",\"type\":\"uint256\"}],\"name\":\"upgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistCap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"whitelistMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"whitelistToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HonestworknftABI is the input ABI used to generate the binding from.
// Deprecated: Use HonestworknftMetaData.ABI instead.
var HonestworknftABI = HonestworknftMetaData.ABI

// Honestworknft is an auto generated Go binding around an Ethereum contract.
type Honestworknft struct {
	HonestworknftCaller     // Read-only binding to the contract
	HonestworknftTransactor // Write-only binding to the contract
	HonestworknftFilterer   // Log filterer for contract events
}

// HonestworknftCaller is an auto generated read-only Go binding around an Ethereum contract.
type HonestworknftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HonestworknftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HonestworknftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HonestworknftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HonestworknftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HonestworknftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HonestworknftSession struct {
	Contract     *Honestworknft    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HonestworknftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HonestworknftCallerSession struct {
	Contract *HonestworknftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// HonestworknftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HonestworknftTransactorSession struct {
	Contract     *HonestworknftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// HonestworknftRaw is an auto generated low-level Go binding around an Ethereum contract.
type HonestworknftRaw struct {
	Contract *Honestworknft // Generic contract binding to access the raw methods on
}

// HonestworknftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HonestworknftCallerRaw struct {
	Contract *HonestworknftCaller // Generic read-only contract binding to access the raw methods on
}

// HonestworknftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HonestworknftTransactorRaw struct {
	Contract *HonestworknftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHonestworknft creates a new instance of Honestworknft, bound to a specific deployed contract.
func NewHonestworknft(address common.Address, backend bind.ContractBackend) (*Honestworknft, error) {
	contract, err := bindHonestworknft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Honestworknft{HonestworknftCaller: HonestworknftCaller{contract: contract}, HonestworknftTransactor: HonestworknftTransactor{contract: contract}, HonestworknftFilterer: HonestworknftFilterer{contract: contract}}, nil
}

// NewHonestworknftCaller creates a new read-only instance of Honestworknft, bound to a specific deployed contract.
func NewHonestworknftCaller(address common.Address, caller bind.ContractCaller) (*HonestworknftCaller, error) {
	contract, err := bindHonestworknft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HonestworknftCaller{contract: contract}, nil
}

// NewHonestworknftTransactor creates a new write-only instance of Honestworknft, bound to a specific deployed contract.
func NewHonestworknftTransactor(address common.Address, transactor bind.ContractTransactor) (*HonestworknftTransactor, error) {
	contract, err := bindHonestworknft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HonestworknftTransactor{contract: contract}, nil
}

// NewHonestworknftFilterer creates a new log filterer instance of Honestworknft, bound to a specific deployed contract.
func NewHonestworknftFilterer(address common.Address, filterer bind.ContractFilterer) (*HonestworknftFilterer, error) {
	contract, err := bindHonestworknft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HonestworknftFilterer{contract: contract}, nil
}

// bindHonestworknft binds a generic wrapper to an already deployed contract.
func bindHonestworknft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HonestworknftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Honestworknft *HonestworknftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Honestworknft.Contract.HonestworknftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Honestworknft *HonestworknftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Honestworknft.Contract.HonestworknftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Honestworknft *HonestworknftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Honestworknft.Contract.HonestworknftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Honestworknft *HonestworknftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Honestworknft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Honestworknft *HonestworknftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Honestworknft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Honestworknft *HonestworknftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Honestworknft.Contract.contract.Transact(opts, method, params...)
}

// TokenIds is a free data retrieval call binding the contract method 0xaa46a400.
//
// Solidity: function _tokenIds() view returns(uint256 _value)
func (_Honestworknft *HonestworknftCaller) TokenIds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "_tokenIds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIds is a free data retrieval call binding the contract method 0xaa46a400.
//
// Solidity: function _tokenIds() view returns(uint256 _value)
func (_Honestworknft *HonestworknftSession) TokenIds() (*big.Int, error) {
	return _Honestworknft.Contract.TokenIds(&_Honestworknft.CallOpts)
}

// TokenIds is a free data retrieval call binding the contract method 0xaa46a400.
//
// Solidity: function _tokenIds() view returns(uint256 _value)
func (_Honestworknft *HonestworknftCallerSession) TokenIds() (*big.Int, error) {
	return _Honestworknft.Contract.TokenIds(&_Honestworknft.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Honestworknft *HonestworknftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Honestworknft *HonestworknftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Honestworknft.Contract.BalanceOf(&_Honestworknft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Honestworknft *HonestworknftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Honestworknft.Contract.BalanceOf(&_Honestworknft.CallOpts, owner)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_Honestworknft *HonestworknftCaller) BaseUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "baseUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_Honestworknft *HonestworknftSession) BaseUri() (string, error) {
	return _Honestworknft.Contract.BaseUri(&_Honestworknft.CallOpts)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_Honestworknft *HonestworknftCallerSession) BaseUri() (string, error) {
	return _Honestworknft.Contract.BaseUri(&_Honestworknft.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Honestworknft *HonestworknftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Honestworknft *HonestworknftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Honestworknft.Contract.GetApproved(&_Honestworknft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Honestworknft *HonestworknftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Honestworknft.Contract.GetApproved(&_Honestworknft.CallOpts, tokenId)
}

// GetTokenTier is a free data retrieval call binding the contract method 0xc26b265f.
//
// Solidity: function getTokenTier(uint256 _tokenId) view returns(uint256)
func (_Honestworknft *HonestworknftCaller) GetTokenTier(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "getTokenTier", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenTier is a free data retrieval call binding the contract method 0xc26b265f.
//
// Solidity: function getTokenTier(uint256 _tokenId) view returns(uint256)
func (_Honestworknft *HonestworknftSession) GetTokenTier(_tokenId *big.Int) (*big.Int, error) {
	return _Honestworknft.Contract.GetTokenTier(&_Honestworknft.CallOpts, _tokenId)
}

// GetTokenTier is a free data retrieval call binding the contract method 0xc26b265f.
//
// Solidity: function getTokenTier(uint256 _tokenId) view returns(uint256)
func (_Honestworknft *HonestworknftCallerSession) GetTokenTier(_tokenId *big.Int) (*big.Int, error) {
	return _Honestworknft.Contract.GetTokenTier(&_Honestworknft.CallOpts, _tokenId)
}

// GetUserTier is a free data retrieval call binding the contract method 0xe4d2620e.
//
// Solidity: function getUserTier(address _user) view returns(uint256)
func (_Honestworknft *HonestworknftCaller) GetUserTier(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "getUserTier", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserTier is a free data retrieval call binding the contract method 0xe4d2620e.
//
// Solidity: function getUserTier(address _user) view returns(uint256)
func (_Honestworknft *HonestworknftSession) GetUserTier(_user common.Address) (*big.Int, error) {
	return _Honestworknft.Contract.GetUserTier(&_Honestworknft.CallOpts, _user)
}

// GetUserTier is a free data retrieval call binding the contract method 0xe4d2620e.
//
// Solidity: function getUserTier(address _user) view returns(uint256)
func (_Honestworknft *HonestworknftCallerSession) GetUserTier(_user common.Address) (*big.Int, error) {
	return _Honestworknft.Contract.GetUserTier(&_Honestworknft.CallOpts, _user)
}

// GetWhitelistToken is a free data retrieval call binding the contract method 0xac4e357d.
//
// Solidity: function getWhitelistToken(address _token) view returns(bool)
func (_Honestworknft *HonestworknftCaller) GetWhitelistToken(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "getWhitelistToken", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetWhitelistToken is a free data retrieval call binding the contract method 0xac4e357d.
//
// Solidity: function getWhitelistToken(address _token) view returns(bool)
func (_Honestworknft *HonestworknftSession) GetWhitelistToken(_token common.Address) (bool, error) {
	return _Honestworknft.Contract.GetWhitelistToken(&_Honestworknft.CallOpts, _token)
}

// GetWhitelistToken is a free data retrieval call binding the contract method 0xac4e357d.
//
// Solidity: function getWhitelistToken(address _token) view returns(bool)
func (_Honestworknft *HonestworknftCallerSession) GetWhitelistToken(_token common.Address) (bool, error) {
	return _Honestworknft.Contract.GetWhitelistToken(&_Honestworknft.CallOpts, _token)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Honestworknft *HonestworknftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Honestworknft *HonestworknftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Honestworknft.Contract.IsApprovedForAll(&_Honestworknft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Honestworknft *HonestworknftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Honestworknft.Contract.IsApprovedForAll(&_Honestworknft.CallOpts, owner, operator)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Honestworknft *HonestworknftCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Honestworknft *HonestworknftSession) IsPaused() (bool, error) {
	return _Honestworknft.Contract.IsPaused(&_Honestworknft.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Honestworknft *HonestworknftCallerSession) IsPaused() (bool, error) {
	return _Honestworknft.Contract.IsPaused(&_Honestworknft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Honestworknft *HonestworknftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Honestworknft *HonestworknftSession) Name() (string, error) {
	return _Honestworknft.Contract.Name(&_Honestworknft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Honestworknft *HonestworknftCallerSession) Name() (string, error) {
	return _Honestworknft.Contract.Name(&_Honestworknft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Honestworknft *HonestworknftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Honestworknft *HonestworknftSession) Owner() (common.Address, error) {
	return _Honestworknft.Contract.Owner(&_Honestworknft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Honestworknft *HonestworknftCallerSession) Owner() (common.Address, error) {
	return _Honestworknft.Contract.Owner(&_Honestworknft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Honestworknft *HonestworknftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Honestworknft *HonestworknftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Honestworknft.Contract.OwnerOf(&_Honestworknft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Honestworknft *HonestworknftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Honestworknft.Contract.OwnerOf(&_Honestworknft.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Honestworknft *HonestworknftCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Honestworknft *HonestworknftSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Honestworknft.Contract.SupportsInterface(&_Honestworknft.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Honestworknft *HonestworknftCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Honestworknft.Contract.SupportsInterface(&_Honestworknft.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Honestworknft *HonestworknftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Honestworknft *HonestworknftSession) Symbol() (string, error) {
	return _Honestworknft.Contract.Symbol(&_Honestworknft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Honestworknft *HonestworknftCallerSession) Symbol() (string, error) {
	return _Honestworknft.Contract.Symbol(&_Honestworknft.CallOpts)
}

// Tier is a free data retrieval call binding the contract method 0x6dda34db.
//
// Solidity: function tier(uint256 ) view returns(uint256)
func (_Honestworknft *HonestworknftCaller) Tier(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "tier", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tier is a free data retrieval call binding the contract method 0x6dda34db.
//
// Solidity: function tier(uint256 ) view returns(uint256)
func (_Honestworknft *HonestworknftSession) Tier(arg0 *big.Int) (*big.Int, error) {
	return _Honestworknft.Contract.Tier(&_Honestworknft.CallOpts, arg0)
}

// Tier is a free data retrieval call binding the contract method 0x6dda34db.
//
// Solidity: function tier(uint256 ) view returns(uint256)
func (_Honestworknft *HonestworknftCallerSession) Tier(arg0 *big.Int) (*big.Int, error) {
	return _Honestworknft.Contract.Tier(&_Honestworknft.CallOpts, arg0)
}

// TierOneFee is a free data retrieval call binding the contract method 0x9a50c628.
//
// Solidity: function tierOneFee() view returns(uint256)
func (_Honestworknft *HonestworknftCaller) TierOneFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "tierOneFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierOneFee is a free data retrieval call binding the contract method 0x9a50c628.
//
// Solidity: function tierOneFee() view returns(uint256)
func (_Honestworknft *HonestworknftSession) TierOneFee() (*big.Int, error) {
	return _Honestworknft.Contract.TierOneFee(&_Honestworknft.CallOpts)
}

// TierOneFee is a free data retrieval call binding the contract method 0x9a50c628.
//
// Solidity: function tierOneFee() view returns(uint256)
func (_Honestworknft *HonestworknftCallerSession) TierOneFee() (*big.Int, error) {
	return _Honestworknft.Contract.TierOneFee(&_Honestworknft.CallOpts)
}

// TierThreeFee is a free data retrieval call binding the contract method 0x9d951b3b.
//
// Solidity: function tierThreeFee() view returns(uint256)
func (_Honestworknft *HonestworknftCaller) TierThreeFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "tierThreeFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierThreeFee is a free data retrieval call binding the contract method 0x9d951b3b.
//
// Solidity: function tierThreeFee() view returns(uint256)
func (_Honestworknft *HonestworknftSession) TierThreeFee() (*big.Int, error) {
	return _Honestworknft.Contract.TierThreeFee(&_Honestworknft.CallOpts)
}

// TierThreeFee is a free data retrieval call binding the contract method 0x9d951b3b.
//
// Solidity: function tierThreeFee() view returns(uint256)
func (_Honestworknft *HonestworknftCallerSession) TierThreeFee() (*big.Int, error) {
	return _Honestworknft.Contract.TierThreeFee(&_Honestworknft.CallOpts)
}

// TierTwoFee is a free data retrieval call binding the contract method 0x08ac6289.
//
// Solidity: function tierTwoFee() view returns(uint256)
func (_Honestworknft *HonestworknftCaller) TierTwoFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "tierTwoFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierTwoFee is a free data retrieval call binding the contract method 0x08ac6289.
//
// Solidity: function tierTwoFee() view returns(uint256)
func (_Honestworknft *HonestworknftSession) TierTwoFee() (*big.Int, error) {
	return _Honestworknft.Contract.TierTwoFee(&_Honestworknft.CallOpts)
}

// TierTwoFee is a free data retrieval call binding the contract method 0x08ac6289.
//
// Solidity: function tierTwoFee() view returns(uint256)
func (_Honestworknft *HonestworknftCallerSession) TierTwoFee() (*big.Int, error) {
	return _Honestworknft.Contract.TierTwoFee(&_Honestworknft.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Honestworknft *HonestworknftCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Honestworknft *HonestworknftSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Honestworknft.Contract.TokenByIndex(&_Honestworknft.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Honestworknft *HonestworknftCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Honestworknft.Contract.TokenByIndex(&_Honestworknft.CallOpts, index)
}

// TokenCap is a free data retrieval call binding the contract method 0xdd54291b.
//
// Solidity: function tokenCap() view returns(uint256)
func (_Honestworknft *HonestworknftCaller) TokenCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "tokenCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCap is a free data retrieval call binding the contract method 0xdd54291b.
//
// Solidity: function tokenCap() view returns(uint256)
func (_Honestworknft *HonestworknftSession) TokenCap() (*big.Int, error) {
	return _Honestworknft.Contract.TokenCap(&_Honestworknft.CallOpts)
}

// TokenCap is a free data retrieval call binding the contract method 0xdd54291b.
//
// Solidity: function tokenCap() view returns(uint256)
func (_Honestworknft *HonestworknftCallerSession) TokenCap() (*big.Int, error) {
	return _Honestworknft.Contract.TokenCap(&_Honestworknft.CallOpts)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Honestworknft *HonestworknftCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Honestworknft *HonestworknftSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Honestworknft.Contract.TokenOfOwnerByIndex(&_Honestworknft.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Honestworknft *HonestworknftCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Honestworknft.Contract.TokenOfOwnerByIndex(&_Honestworknft.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Honestworknft *HonestworknftCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Honestworknft *HonestworknftSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Honestworknft.Contract.TokenURI(&_Honestworknft.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Honestworknft *HonestworknftCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Honestworknft.Contract.TokenURI(&_Honestworknft.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Honestworknft *HonestworknftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Honestworknft *HonestworknftSession) TotalSupply() (*big.Int, error) {
	return _Honestworknft.Contract.TotalSupply(&_Honestworknft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Honestworknft *HonestworknftCallerSession) TotalSupply() (*big.Int, error) {
	return _Honestworknft.Contract.TotalSupply(&_Honestworknft.CallOpts)
}

// WhitelistCap is a free data retrieval call binding the contract method 0xff5de458.
//
// Solidity: function whitelistCap(address ) view returns(bool)
func (_Honestworknft *HonestworknftCaller) WhitelistCap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "whitelistCap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistCap is a free data retrieval call binding the contract method 0xff5de458.
//
// Solidity: function whitelistCap(address ) view returns(bool)
func (_Honestworknft *HonestworknftSession) WhitelistCap(arg0 common.Address) (bool, error) {
	return _Honestworknft.Contract.WhitelistCap(&_Honestworknft.CallOpts, arg0)
}

// WhitelistCap is a free data retrieval call binding the contract method 0xff5de458.
//
// Solidity: function whitelistCap(address ) view returns(bool)
func (_Honestworknft *HonestworknftCallerSession) WhitelistCap(arg0 common.Address) (bool, error) {
	return _Honestworknft.Contract.WhitelistCap(&_Honestworknft.CallOpts, arg0)
}

// WhitelistRoot is a free data retrieval call binding the contract method 0x386bfc98.
//
// Solidity: function whitelistRoot() view returns(bytes32)
func (_Honestworknft *HonestworknftCaller) WhitelistRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "whitelistRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WhitelistRoot is a free data retrieval call binding the contract method 0x386bfc98.
//
// Solidity: function whitelistRoot() view returns(bytes32)
func (_Honestworknft *HonestworknftSession) WhitelistRoot() ([32]byte, error) {
	return _Honestworknft.Contract.WhitelistRoot(&_Honestworknft.CallOpts)
}

// WhitelistRoot is a free data retrieval call binding the contract method 0x386bfc98.
//
// Solidity: function whitelistRoot() view returns(bytes32)
func (_Honestworknft *HonestworknftCallerSession) WhitelistRoot() ([32]byte, error) {
	return _Honestworknft.Contract.WhitelistRoot(&_Honestworknft.CallOpts)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_Honestworknft *HonestworknftCaller) WhitelistedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Honestworknft.contract.Call(opts, &out, "whitelistedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_Honestworknft *HonestworknftSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _Honestworknft.Contract.WhitelistedTokens(&_Honestworknft.CallOpts, arg0)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_Honestworknft *HonestworknftCallerSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _Honestworknft.Contract.WhitelistedTokens(&_Honestworknft.CallOpts, arg0)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address _to, uint256 _tier) returns()
func (_Honestworknft *HonestworknftTransactor) AdminMint(opts *bind.TransactOpts, _to common.Address, _tier *big.Int) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "adminMint", _to, _tier)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address _to, uint256 _tier) returns()
func (_Honestworknft *HonestworknftSession) AdminMint(_to common.Address, _tier *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.AdminMint(&_Honestworknft.TransactOpts, _to, _tier)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address _to, uint256 _tier) returns()
func (_Honestworknft *HonestworknftTransactorSession) AdminMint(_to common.Address, _tier *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.AdminMint(&_Honestworknft.TransactOpts, _to, _tier)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Honestworknft *HonestworknftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Honestworknft *HonestworknftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.Approve(&_Honestworknft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Honestworknft *HonestworknftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.Approve(&_Honestworknft.TransactOpts, to, tokenId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Honestworknft *HonestworknftTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Honestworknft *HonestworknftSession) Pause() (*types.Transaction, error) {
	return _Honestworknft.Contract.Pause(&_Honestworknft.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Honestworknft *HonestworknftTransactorSession) Pause() (*types.Transaction, error) {
	return _Honestworknft.Contract.Pause(&_Honestworknft.TransactOpts)
}

// PublicMint is a paid mutator transaction binding the contract method 0x32a93a3a.
//
// Solidity: function publicMint(address _token) returns()
func (_Honestworknft *HonestworknftTransactor) PublicMint(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "publicMint", _token)
}

// PublicMint is a paid mutator transaction binding the contract method 0x32a93a3a.
//
// Solidity: function publicMint(address _token) returns()
func (_Honestworknft *HonestworknftSession) PublicMint(_token common.Address) (*types.Transaction, error) {
	return _Honestworknft.Contract.PublicMint(&_Honestworknft.TransactOpts, _token)
}

// PublicMint is a paid mutator transaction binding the contract method 0x32a93a3a.
//
// Solidity: function publicMint(address _token) returns()
func (_Honestworknft *HonestworknftTransactorSession) PublicMint(_token common.Address) (*types.Transaction, error) {
	return _Honestworknft.Contract.PublicMint(&_Honestworknft.TransactOpts, _token)
}

// RemoveWhitelistToken is a paid mutator transaction binding the contract method 0xe9ee9107.
//
// Solidity: function removeWhitelistToken(address _token) returns()
func (_Honestworknft *HonestworknftTransactor) RemoveWhitelistToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "removeWhitelistToken", _token)
}

// RemoveWhitelistToken is a paid mutator transaction binding the contract method 0xe9ee9107.
//
// Solidity: function removeWhitelistToken(address _token) returns()
func (_Honestworknft *HonestworknftSession) RemoveWhitelistToken(_token common.Address) (*types.Transaction, error) {
	return _Honestworknft.Contract.RemoveWhitelistToken(&_Honestworknft.TransactOpts, _token)
}

// RemoveWhitelistToken is a paid mutator transaction binding the contract method 0xe9ee9107.
//
// Solidity: function removeWhitelistToken(address _token) returns()
func (_Honestworknft *HonestworknftTransactorSession) RemoveWhitelistToken(_token common.Address) (*types.Transaction, error) {
	return _Honestworknft.Contract.RemoveWhitelistToken(&_Honestworknft.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Honestworknft *HonestworknftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Honestworknft *HonestworknftSession) RenounceOwnership() (*types.Transaction, error) {
	return _Honestworknft.Contract.RenounceOwnership(&_Honestworknft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Honestworknft *HonestworknftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Honestworknft.Contract.RenounceOwnership(&_Honestworknft.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Honestworknft *HonestworknftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Honestworknft *HonestworknftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.SafeTransferFrom(&_Honestworknft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Honestworknft *HonestworknftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.SafeTransferFrom(&_Honestworknft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Honestworknft *HonestworknftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Honestworknft *HonestworknftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Honestworknft.Contract.SafeTransferFrom0(&_Honestworknft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Honestworknft *HonestworknftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Honestworknft.Contract.SafeTransferFrom0(&_Honestworknft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Honestworknft *HonestworknftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Honestworknft *HonestworknftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Honestworknft.Contract.SetApprovalForAll(&_Honestworknft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Honestworknft *HonestworknftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Honestworknft.Contract.SetApprovalForAll(&_Honestworknft.TransactOpts, operator, approved)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string _baseUri) returns()
func (_Honestworknft *HonestworknftTransactor) SetBaseUri(opts *bind.TransactOpts, _baseUri string) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "setBaseUri", _baseUri)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string _baseUri) returns()
func (_Honestworknft *HonestworknftSession) SetBaseUri(_baseUri string) (*types.Transaction, error) {
	return _Honestworknft.Contract.SetBaseUri(&_Honestworknft.TransactOpts, _baseUri)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string _baseUri) returns()
func (_Honestworknft *HonestworknftTransactorSession) SetBaseUri(_baseUri string) (*types.Transaction, error) {
	return _Honestworknft.Contract.SetBaseUri(&_Honestworknft.TransactOpts, _baseUri)
}

// SetTiers is a paid mutator transaction binding the contract method 0x0f8343e5.
//
// Solidity: function setTiers(uint256 _tierOneFee, uint256 _tierTwoFee, uint256 _tierThreeFee) returns()
func (_Honestworknft *HonestworknftTransactor) SetTiers(opts *bind.TransactOpts, _tierOneFee *big.Int, _tierTwoFee *big.Int, _tierThreeFee *big.Int) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "setTiers", _tierOneFee, _tierTwoFee, _tierThreeFee)
}

// SetTiers is a paid mutator transaction binding the contract method 0x0f8343e5.
//
// Solidity: function setTiers(uint256 _tierOneFee, uint256 _tierTwoFee, uint256 _tierThreeFee) returns()
func (_Honestworknft *HonestworknftSession) SetTiers(_tierOneFee *big.Int, _tierTwoFee *big.Int, _tierThreeFee *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.SetTiers(&_Honestworknft.TransactOpts, _tierOneFee, _tierTwoFee, _tierThreeFee)
}

// SetTiers is a paid mutator transaction binding the contract method 0x0f8343e5.
//
// Solidity: function setTiers(uint256 _tierOneFee, uint256 _tierTwoFee, uint256 _tierThreeFee) returns()
func (_Honestworknft *HonestworknftTransactorSession) SetTiers(_tierOneFee *big.Int, _tierTwoFee *big.Int, _tierThreeFee *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.SetTiers(&_Honestworknft.TransactOpts, _tierOneFee, _tierTwoFee, _tierThreeFee)
}

// SetTokenCap is a paid mutator transaction binding the contract method 0x2854bc7e.
//
// Solidity: function setTokenCap(uint256 _tokenCap) returns()
func (_Honestworknft *HonestworknftTransactor) SetTokenCap(opts *bind.TransactOpts, _tokenCap *big.Int) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "setTokenCap", _tokenCap)
}

// SetTokenCap is a paid mutator transaction binding the contract method 0x2854bc7e.
//
// Solidity: function setTokenCap(uint256 _tokenCap) returns()
func (_Honestworknft *HonestworknftSession) SetTokenCap(_tokenCap *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.SetTokenCap(&_Honestworknft.TransactOpts, _tokenCap)
}

// SetTokenCap is a paid mutator transaction binding the contract method 0x2854bc7e.
//
// Solidity: function setTokenCap(uint256 _tokenCap) returns()
func (_Honestworknft *HonestworknftTransactorSession) SetTokenCap(_tokenCap *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.SetTokenCap(&_Honestworknft.TransactOpts, _tokenCap)
}

// SetWhitelistRoot is a paid mutator transaction binding the contract method 0xf5aa406d.
//
// Solidity: function setWhitelistRoot(bytes32 _root) returns()
func (_Honestworknft *HonestworknftTransactor) SetWhitelistRoot(opts *bind.TransactOpts, _root [32]byte) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "setWhitelistRoot", _root)
}

// SetWhitelistRoot is a paid mutator transaction binding the contract method 0xf5aa406d.
//
// Solidity: function setWhitelistRoot(bytes32 _root) returns()
func (_Honestworknft *HonestworknftSession) SetWhitelistRoot(_root [32]byte) (*types.Transaction, error) {
	return _Honestworknft.Contract.SetWhitelistRoot(&_Honestworknft.TransactOpts, _root)
}

// SetWhitelistRoot is a paid mutator transaction binding the contract method 0xf5aa406d.
//
// Solidity: function setWhitelistRoot(bytes32 _root) returns()
func (_Honestworknft *HonestworknftTransactorSession) SetWhitelistRoot(_root [32]byte) (*types.Transaction, error) {
	return _Honestworknft.Contract.SetWhitelistRoot(&_Honestworknft.TransactOpts, _root)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Honestworknft *HonestworknftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Honestworknft *HonestworknftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.TransferFrom(&_Honestworknft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Honestworknft *HonestworknftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.TransferFrom(&_Honestworknft.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Honestworknft *HonestworknftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Honestworknft *HonestworknftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Honestworknft.Contract.TransferOwnership(&_Honestworknft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Honestworknft *HonestworknftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Honestworknft.Contract.TransferOwnership(&_Honestworknft.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Honestworknft *HonestworknftTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Honestworknft *HonestworknftSession) Unpause() (*types.Transaction, error) {
	return _Honestworknft.Contract.Unpause(&_Honestworknft.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Honestworknft *HonestworknftTransactorSession) Unpause() (*types.Transaction, error) {
	return _Honestworknft.Contract.Unpause(&_Honestworknft.TransactOpts)
}

// UpgradeToken is a paid mutator transaction binding the contract method 0xc9abb2e9.
//
// Solidity: function upgradeToken(address _token, uint256 _levels) returns()
func (_Honestworknft *HonestworknftTransactor) UpgradeToken(opts *bind.TransactOpts, _token common.Address, _levels *big.Int) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "upgradeToken", _token, _levels)
}

// UpgradeToken is a paid mutator transaction binding the contract method 0xc9abb2e9.
//
// Solidity: function upgradeToken(address _token, uint256 _levels) returns()
func (_Honestworknft *HonestworknftSession) UpgradeToken(_token common.Address, _levels *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.UpgradeToken(&_Honestworknft.TransactOpts, _token, _levels)
}

// UpgradeToken is a paid mutator transaction binding the contract method 0xc9abb2e9.
//
// Solidity: function upgradeToken(address _token, uint256 _levels) returns()
func (_Honestworknft *HonestworknftTransactorSession) UpgradeToken(_token common.Address, _levels *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.UpgradeToken(&_Honestworknft.TransactOpts, _token, _levels)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0x372f657c.
//
// Solidity: function whitelistMint(bytes32[] _proof) returns()
func (_Honestworknft *HonestworknftTransactor) WhitelistMint(opts *bind.TransactOpts, _proof [][32]byte) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "whitelistMint", _proof)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0x372f657c.
//
// Solidity: function whitelistMint(bytes32[] _proof) returns()
func (_Honestworknft *HonestworknftSession) WhitelistMint(_proof [][32]byte) (*types.Transaction, error) {
	return _Honestworknft.Contract.WhitelistMint(&_Honestworknft.TransactOpts, _proof)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0x372f657c.
//
// Solidity: function whitelistMint(bytes32[] _proof) returns()
func (_Honestworknft *HonestworknftTransactorSession) WhitelistMint(_proof [][32]byte) (*types.Transaction, error) {
	return _Honestworknft.Contract.WhitelistMint(&_Honestworknft.TransactOpts, _proof)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x6247f6f2.
//
// Solidity: function whitelistToken(address _token) returns()
func (_Honestworknft *HonestworknftTransactor) WhitelistToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "whitelistToken", _token)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x6247f6f2.
//
// Solidity: function whitelistToken(address _token) returns()
func (_Honestworknft *HonestworknftSession) WhitelistToken(_token common.Address) (*types.Transaction, error) {
	return _Honestworknft.Contract.WhitelistToken(&_Honestworknft.TransactOpts, _token)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x6247f6f2.
//
// Solidity: function whitelistToken(address _token) returns()
func (_Honestworknft *HonestworknftTransactorSession) WhitelistToken(_token common.Address) (*types.Transaction, error) {
	return _Honestworknft.Contract.WhitelistToken(&_Honestworknft.TransactOpts, _token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_Honestworknft *HonestworknftTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Honestworknft.contract.Transact(opts, "withdraw", _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_Honestworknft *HonestworknftSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.Withdraw(&_Honestworknft.TransactOpts, _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_Honestworknft *HonestworknftTransactorSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Honestworknft.Contract.Withdraw(&_Honestworknft.TransactOpts, _token, _amount)
}

// HonestworknftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Honestworknft contract.
type HonestworknftApprovalIterator struct {
	Event *HonestworknftApproval // Event containing the contract specifics and raw log

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
func (it *HonestworknftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HonestworknftApproval)
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
		it.Event = new(HonestworknftApproval)
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
func (it *HonestworknftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HonestworknftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HonestworknftApproval represents a Approval event raised by the Honestworknft contract.
type HonestworknftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Honestworknft *HonestworknftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*HonestworknftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Honestworknft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HonestworknftApprovalIterator{contract: _Honestworknft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Honestworknft *HonestworknftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HonestworknftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Honestworknft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HonestworknftApproval)
				if err := _Honestworknft.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Honestworknft *HonestworknftFilterer) ParseApproval(log types.Log) (*HonestworknftApproval, error) {
	event := new(HonestworknftApproval)
	if err := _Honestworknft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HonestworknftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Honestworknft contract.
type HonestworknftApprovalForAllIterator struct {
	Event *HonestworknftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *HonestworknftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HonestworknftApprovalForAll)
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
		it.Event = new(HonestworknftApprovalForAll)
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
func (it *HonestworknftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HonestworknftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HonestworknftApprovalForAll represents a ApprovalForAll event raised by the Honestworknft contract.
type HonestworknftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Honestworknft *HonestworknftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*HonestworknftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Honestworknft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &HonestworknftApprovalForAllIterator{contract: _Honestworknft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Honestworknft *HonestworknftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *HonestworknftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Honestworknft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HonestworknftApprovalForAll)
				if err := _Honestworknft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Honestworknft *HonestworknftFilterer) ParseApprovalForAll(log types.Log) (*HonestworknftApprovalForAll, error) {
	event := new(HonestworknftApprovalForAll)
	if err := _Honestworknft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HonestworknftMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Honestworknft contract.
type HonestworknftMintIterator struct {
	Event *HonestworknftMint // Event containing the contract specifics and raw log

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
func (it *HonestworknftMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HonestworknftMint)
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
		it.Event = new(HonestworknftMint)
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
func (it *HonestworknftMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HonestworknftMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HonestworknftMint represents a Mint event raised by the Honestworknft contract.
type HonestworknftMint struct {
	Id   *big.Int
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xf3cea5493d790af0133817606f7350a91d7f154ea52eaa79d179d4d231e50102.
//
// Solidity: event Mint(uint256 id, address user)
func (_Honestworknft *HonestworknftFilterer) FilterMint(opts *bind.FilterOpts) (*HonestworknftMintIterator, error) {

	logs, sub, err := _Honestworknft.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &HonestworknftMintIterator{contract: _Honestworknft.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xf3cea5493d790af0133817606f7350a91d7f154ea52eaa79d179d4d231e50102.
//
// Solidity: event Mint(uint256 id, address user)
func (_Honestworknft *HonestworknftFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *HonestworknftMint) (event.Subscription, error) {

	logs, sub, err := _Honestworknft.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HonestworknftMint)
				if err := _Honestworknft.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xf3cea5493d790af0133817606f7350a91d7f154ea52eaa79d179d4d231e50102.
//
// Solidity: event Mint(uint256 id, address user)
func (_Honestworknft *HonestworknftFilterer) ParseMint(log types.Log) (*HonestworknftMint, error) {
	event := new(HonestworknftMint)
	if err := _Honestworknft.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HonestworknftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Honestworknft contract.
type HonestworknftOwnershipTransferredIterator struct {
	Event *HonestworknftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HonestworknftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HonestworknftOwnershipTransferred)
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
		it.Event = new(HonestworknftOwnershipTransferred)
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
func (it *HonestworknftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HonestworknftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HonestworknftOwnershipTransferred represents a OwnershipTransferred event raised by the Honestworknft contract.
type HonestworknftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Honestworknft *HonestworknftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HonestworknftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Honestworknft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HonestworknftOwnershipTransferredIterator{contract: _Honestworknft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Honestworknft *HonestworknftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HonestworknftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Honestworknft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HonestworknftOwnershipTransferred)
				if err := _Honestworknft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Honestworknft *HonestworknftFilterer) ParseOwnershipTransferred(log types.Log) (*HonestworknftOwnershipTransferred, error) {
	event := new(HonestworknftOwnershipTransferred)
	if err := _Honestworknft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HonestworknftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Honestworknft contract.
type HonestworknftTransferIterator struct {
	Event *HonestworknftTransfer // Event containing the contract specifics and raw log

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
func (it *HonestworknftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HonestworknftTransfer)
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
		it.Event = new(HonestworknftTransfer)
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
func (it *HonestworknftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HonestworknftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HonestworknftTransfer represents a Transfer event raised by the Honestworknft contract.
type HonestworknftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Honestworknft *HonestworknftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*HonestworknftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Honestworknft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HonestworknftTransferIterator{contract: _Honestworknft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Honestworknft *HonestworknftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HonestworknftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Honestworknft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HonestworknftTransfer)
				if err := _Honestworknft.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Honestworknft *HonestworknftFilterer) ParseTransfer(log types.Log) (*HonestworknftTransfer, error) {
	event := new(HonestworknftTransfer)
	if err := _Honestworknft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HonestworknftUpgradeIterator is returned from FilterUpgrade and is used to iterate over the raw logs and unpacked data for Upgrade events raised by the Honestworknft contract.
type HonestworknftUpgradeIterator struct {
	Event *HonestworknftUpgrade // Event containing the contract specifics and raw log

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
func (it *HonestworknftUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HonestworknftUpgrade)
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
		it.Event = new(HonestworknftUpgrade)
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
func (it *HonestworknftUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HonestworknftUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HonestworknftUpgrade represents a Upgrade event raised by the Honestworknft contract.
type HonestworknftUpgrade struct {
	Id   *big.Int
	Tier *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpgrade is a free log retrieval operation binding the contract event 0xeb0ebb51128928d7b1a6419c52128a9319bfcb55f0adafea75afbf75f9f2f3e9.
//
// Solidity: event Upgrade(uint256 id, uint256 tier)
func (_Honestworknft *HonestworknftFilterer) FilterUpgrade(opts *bind.FilterOpts) (*HonestworknftUpgradeIterator, error) {

	logs, sub, err := _Honestworknft.contract.FilterLogs(opts, "Upgrade")
	if err != nil {
		return nil, err
	}
	return &HonestworknftUpgradeIterator{contract: _Honestworknft.contract, event: "Upgrade", logs: logs, sub: sub}, nil
}

// WatchUpgrade is a free log subscription operation binding the contract event 0xeb0ebb51128928d7b1a6419c52128a9319bfcb55f0adafea75afbf75f9f2f3e9.
//
// Solidity: event Upgrade(uint256 id, uint256 tier)
func (_Honestworknft *HonestworknftFilterer) WatchUpgrade(opts *bind.WatchOpts, sink chan<- *HonestworknftUpgrade) (event.Subscription, error) {

	logs, sub, err := _Honestworknft.contract.WatchLogs(opts, "Upgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HonestworknftUpgrade)
				if err := _Honestworknft.contract.UnpackLog(event, "Upgrade", log); err != nil {
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

// ParseUpgrade is a log parse operation binding the contract event 0xeb0ebb51128928d7b1a6419c52128a9319bfcb55f0adafea75afbf75f9f2f3e9.
//
// Solidity: event Upgrade(uint256 id, uint256 tier)
func (_Honestworknft *HonestworknftFilterer) ParseUpgrade(log types.Log) (*HonestworknftUpgrade, error) {
	event := new(HonestworknftUpgrade)
	if err := _Honestworknft.contract.UnpackLog(event, "Upgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
