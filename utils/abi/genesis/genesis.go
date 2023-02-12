// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package genesis

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

// GenesisMetaData contains all meta data concerning the Genesis contract.
var GenesisMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseUri\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_whitelistedTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tier\",\"type\":\"uint256\"}],\"name\":\"Upgrade\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TOKEN_CAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_tokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tier\",\"type\":\"uint256\"}],\"name\":\"adminMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenTier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserTier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"publicMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"setWhitelistRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierOneFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierThreeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierTwoFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_levels\",\"type\":\"uint256\"}],\"name\":\"upgradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistCap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"whitelistMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"whitelistToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GenesisABI is the input ABI used to generate the binding from.
// Deprecated: Use GenesisMetaData.ABI instead.
var GenesisABI = GenesisMetaData.ABI

// Genesis is an auto generated Go binding around an Ethereum contract.
type Genesis struct {
	GenesisCaller     // Read-only binding to the contract
	GenesisTransactor // Write-only binding to the contract
	GenesisFilterer   // Log filterer for contract events
}

// GenesisCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenesisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenesisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenesisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenesisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenesisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenesisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenesisSession struct {
	Contract     *Genesis          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenesisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenesisCallerSession struct {
	Contract *GenesisCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GenesisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenesisTransactorSession struct {
	Contract     *GenesisTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GenesisRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenesisRaw struct {
	Contract *Genesis // Generic contract binding to access the raw methods on
}

// GenesisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenesisCallerRaw struct {
	Contract *GenesisCaller // Generic read-only contract binding to access the raw methods on
}

// GenesisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenesisTransactorRaw struct {
	Contract *GenesisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenesis creates a new instance of Genesis, bound to a specific deployed contract.
func NewGenesis(address common.Address, backend bind.ContractBackend) (*Genesis, error) {
	contract, err := bindGenesis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Genesis{GenesisCaller: GenesisCaller{contract: contract}, GenesisTransactor: GenesisTransactor{contract: contract}, GenesisFilterer: GenesisFilterer{contract: contract}}, nil
}

// NewGenesisCaller creates a new read-only instance of Genesis, bound to a specific deployed contract.
func NewGenesisCaller(address common.Address, caller bind.ContractCaller) (*GenesisCaller, error) {
	contract, err := bindGenesis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenesisCaller{contract: contract}, nil
}

// NewGenesisTransactor creates a new write-only instance of Genesis, bound to a specific deployed contract.
func NewGenesisTransactor(address common.Address, transactor bind.ContractTransactor) (*GenesisTransactor, error) {
	contract, err := bindGenesis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenesisTransactor{contract: contract}, nil
}

// NewGenesisFilterer creates a new log filterer instance of Genesis, bound to a specific deployed contract.
func NewGenesisFilterer(address common.Address, filterer bind.ContractFilterer) (*GenesisFilterer, error) {
	contract, err := bindGenesis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenesisFilterer{contract: contract}, nil
}

// bindGenesis binds a generic wrapper to an already deployed contract.
func bindGenesis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenesisABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Genesis *GenesisRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Genesis.Contract.GenesisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Genesis *GenesisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Genesis.Contract.GenesisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Genesis *GenesisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Genesis.Contract.GenesisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Genesis *GenesisCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Genesis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Genesis *GenesisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Genesis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Genesis *GenesisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Genesis.Contract.contract.Transact(opts, method, params...)
}

// TOKENCAP is a free data retrieval call binding the contract method 0x9a6524f1.
//
// Solidity: function TOKEN_CAP() view returns(uint256)
func (_Genesis *GenesisCaller) TOKENCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "TOKEN_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOKENCAP is a free data retrieval call binding the contract method 0x9a6524f1.
//
// Solidity: function TOKEN_CAP() view returns(uint256)
func (_Genesis *GenesisSession) TOKENCAP() (*big.Int, error) {
	return _Genesis.Contract.TOKENCAP(&_Genesis.CallOpts)
}

// TOKENCAP is a free data retrieval call binding the contract method 0x9a6524f1.
//
// Solidity: function TOKEN_CAP() view returns(uint256)
func (_Genesis *GenesisCallerSession) TOKENCAP() (*big.Int, error) {
	return _Genesis.Contract.TOKENCAP(&_Genesis.CallOpts)
}

// TokenIds is a free data retrieval call binding the contract method 0xaa46a400.
//
// Solidity: function _tokenIds() view returns(uint256 _value)
func (_Genesis *GenesisCaller) TokenIds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "_tokenIds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIds is a free data retrieval call binding the contract method 0xaa46a400.
//
// Solidity: function _tokenIds() view returns(uint256 _value)
func (_Genesis *GenesisSession) TokenIds() (*big.Int, error) {
	return _Genesis.Contract.TokenIds(&_Genesis.CallOpts)
}

// TokenIds is a free data retrieval call binding the contract method 0xaa46a400.
//
// Solidity: function _tokenIds() view returns(uint256 _value)
func (_Genesis *GenesisCallerSession) TokenIds() (*big.Int, error) {
	return _Genesis.Contract.TokenIds(&_Genesis.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Genesis *GenesisCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Genesis *GenesisSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Genesis.Contract.BalanceOf(&_Genesis.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Genesis *GenesisCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Genesis.Contract.BalanceOf(&_Genesis.CallOpts, owner)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_Genesis *GenesisCaller) BaseUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "baseUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_Genesis *GenesisSession) BaseUri() (string, error) {
	return _Genesis.Contract.BaseUri(&_Genesis.CallOpts)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_Genesis *GenesisCallerSession) BaseUri() (string, error) {
	return _Genesis.Contract.BaseUri(&_Genesis.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Genesis *GenesisCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Genesis *GenesisSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Genesis.Contract.GetApproved(&_Genesis.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Genesis *GenesisCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Genesis.Contract.GetApproved(&_Genesis.CallOpts, tokenId)
}

// GetTokenTier is a free data retrieval call binding the contract method 0xc26b265f.
//
// Solidity: function getTokenTier(uint256 _tokenId) view returns(uint256)
func (_Genesis *GenesisCaller) GetTokenTier(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "getTokenTier", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenTier is a free data retrieval call binding the contract method 0xc26b265f.
//
// Solidity: function getTokenTier(uint256 _tokenId) view returns(uint256)
func (_Genesis *GenesisSession) GetTokenTier(_tokenId *big.Int) (*big.Int, error) {
	return _Genesis.Contract.GetTokenTier(&_Genesis.CallOpts, _tokenId)
}

// GetTokenTier is a free data retrieval call binding the contract method 0xc26b265f.
//
// Solidity: function getTokenTier(uint256 _tokenId) view returns(uint256)
func (_Genesis *GenesisCallerSession) GetTokenTier(_tokenId *big.Int) (*big.Int, error) {
	return _Genesis.Contract.GetTokenTier(&_Genesis.CallOpts, _tokenId)
}

// GetUserTier is a free data retrieval call binding the contract method 0xe4d2620e.
//
// Solidity: function getUserTier(address _user) view returns(uint256)
func (_Genesis *GenesisCaller) GetUserTier(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "getUserTier", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserTier is a free data retrieval call binding the contract method 0xe4d2620e.
//
// Solidity: function getUserTier(address _user) view returns(uint256)
func (_Genesis *GenesisSession) GetUserTier(_user common.Address) (*big.Int, error) {
	return _Genesis.Contract.GetUserTier(&_Genesis.CallOpts, _user)
}

// GetUserTier is a free data retrieval call binding the contract method 0xe4d2620e.
//
// Solidity: function getUserTier(address _user) view returns(uint256)
func (_Genesis *GenesisCallerSession) GetUserTier(_user common.Address) (*big.Int, error) {
	return _Genesis.Contract.GetUserTier(&_Genesis.CallOpts, _user)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Genesis *GenesisCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Genesis *GenesisSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Genesis.Contract.IsApprovedForAll(&_Genesis.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Genesis *GenesisCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Genesis.Contract.IsApprovedForAll(&_Genesis.CallOpts, owner, operator)
}

// MetadataImplementation is a free data retrieval call binding the contract method 0x90eadc32.
//
// Solidity: function metadataImplementation() view returns(address)
func (_Genesis *GenesisCaller) MetadataImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "metadataImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetadataImplementation is a free data retrieval call binding the contract method 0x90eadc32.
//
// Solidity: function metadataImplementation() view returns(address)
func (_Genesis *GenesisSession) MetadataImplementation() (common.Address, error) {
	return _Genesis.Contract.MetadataImplementation(&_Genesis.CallOpts)
}

// MetadataImplementation is a free data retrieval call binding the contract method 0x90eadc32.
//
// Solidity: function metadataImplementation() view returns(address)
func (_Genesis *GenesisCallerSession) MetadataImplementation() (common.Address, error) {
	return _Genesis.Contract.MetadataImplementation(&_Genesis.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Genesis *GenesisCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Genesis *GenesisSession) Name() (string, error) {
	return _Genesis.Contract.Name(&_Genesis.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Genesis *GenesisCallerSession) Name() (string, error) {
	return _Genesis.Contract.Name(&_Genesis.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Genesis *GenesisCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Genesis *GenesisSession) Owner() (common.Address, error) {
	return _Genesis.Contract.Owner(&_Genesis.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Genesis *GenesisCallerSession) Owner() (common.Address, error) {
	return _Genesis.Contract.Owner(&_Genesis.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Genesis *GenesisCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Genesis *GenesisSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Genesis.Contract.OwnerOf(&_Genesis.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Genesis *GenesisCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Genesis.Contract.OwnerOf(&_Genesis.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Genesis *GenesisCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Genesis *GenesisSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Genesis.Contract.SupportsInterface(&_Genesis.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Genesis *GenesisCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Genesis.Contract.SupportsInterface(&_Genesis.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Genesis *GenesisCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Genesis *GenesisSession) Symbol() (string, error) {
	return _Genesis.Contract.Symbol(&_Genesis.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Genesis *GenesisCallerSession) Symbol() (string, error) {
	return _Genesis.Contract.Symbol(&_Genesis.CallOpts)
}

// Tier is a free data retrieval call binding the contract method 0x6dda34db.
//
// Solidity: function tier(uint256 ) view returns(uint256)
func (_Genesis *GenesisCaller) Tier(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "tier", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tier is a free data retrieval call binding the contract method 0x6dda34db.
//
// Solidity: function tier(uint256 ) view returns(uint256)
func (_Genesis *GenesisSession) Tier(arg0 *big.Int) (*big.Int, error) {
	return _Genesis.Contract.Tier(&_Genesis.CallOpts, arg0)
}

// Tier is a free data retrieval call binding the contract method 0x6dda34db.
//
// Solidity: function tier(uint256 ) view returns(uint256)
func (_Genesis *GenesisCallerSession) Tier(arg0 *big.Int) (*big.Int, error) {
	return _Genesis.Contract.Tier(&_Genesis.CallOpts, arg0)
}

// TierOneFee is a free data retrieval call binding the contract method 0x9a50c628.
//
// Solidity: function tierOneFee() view returns(uint256)
func (_Genesis *GenesisCaller) TierOneFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "tierOneFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierOneFee is a free data retrieval call binding the contract method 0x9a50c628.
//
// Solidity: function tierOneFee() view returns(uint256)
func (_Genesis *GenesisSession) TierOneFee() (*big.Int, error) {
	return _Genesis.Contract.TierOneFee(&_Genesis.CallOpts)
}

// TierOneFee is a free data retrieval call binding the contract method 0x9a50c628.
//
// Solidity: function tierOneFee() view returns(uint256)
func (_Genesis *GenesisCallerSession) TierOneFee() (*big.Int, error) {
	return _Genesis.Contract.TierOneFee(&_Genesis.CallOpts)
}

// TierThreeFee is a free data retrieval call binding the contract method 0x9d951b3b.
//
// Solidity: function tierThreeFee() view returns(uint256)
func (_Genesis *GenesisCaller) TierThreeFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "tierThreeFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierThreeFee is a free data retrieval call binding the contract method 0x9d951b3b.
//
// Solidity: function tierThreeFee() view returns(uint256)
func (_Genesis *GenesisSession) TierThreeFee() (*big.Int, error) {
	return _Genesis.Contract.TierThreeFee(&_Genesis.CallOpts)
}

// TierThreeFee is a free data retrieval call binding the contract method 0x9d951b3b.
//
// Solidity: function tierThreeFee() view returns(uint256)
func (_Genesis *GenesisCallerSession) TierThreeFee() (*big.Int, error) {
	return _Genesis.Contract.TierThreeFee(&_Genesis.CallOpts)
}

// TierTwoFee is a free data retrieval call binding the contract method 0x08ac6289.
//
// Solidity: function tierTwoFee() view returns(uint256)
func (_Genesis *GenesisCaller) TierTwoFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "tierTwoFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierTwoFee is a free data retrieval call binding the contract method 0x08ac6289.
//
// Solidity: function tierTwoFee() view returns(uint256)
func (_Genesis *GenesisSession) TierTwoFee() (*big.Int, error) {
	return _Genesis.Contract.TierTwoFee(&_Genesis.CallOpts)
}

// TierTwoFee is a free data retrieval call binding the contract method 0x08ac6289.
//
// Solidity: function tierTwoFee() view returns(uint256)
func (_Genesis *GenesisCallerSession) TierTwoFee() (*big.Int, error) {
	return _Genesis.Contract.TierTwoFee(&_Genesis.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Genesis *GenesisCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Genesis *GenesisSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Genesis.Contract.TokenByIndex(&_Genesis.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Genesis *GenesisCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Genesis.Contract.TokenByIndex(&_Genesis.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Genesis *GenesisCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Genesis *GenesisSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Genesis.Contract.TokenOfOwnerByIndex(&_Genesis.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Genesis *GenesisCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Genesis.Contract.TokenOfOwnerByIndex(&_Genesis.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Genesis *GenesisCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Genesis *GenesisSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Genesis.Contract.TokenURI(&_Genesis.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Genesis *GenesisCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Genesis.Contract.TokenURI(&_Genesis.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Genesis *GenesisCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Genesis *GenesisSession) TotalSupply() (*big.Int, error) {
	return _Genesis.Contract.TotalSupply(&_Genesis.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Genesis *GenesisCallerSession) TotalSupply() (*big.Int, error) {
	return _Genesis.Contract.TotalSupply(&_Genesis.CallOpts)
}

// WhitelistCap is a free data retrieval call binding the contract method 0xff5de458.
//
// Solidity: function whitelistCap(address ) view returns(bool)
func (_Genesis *GenesisCaller) WhitelistCap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "whitelistCap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistCap is a free data retrieval call binding the contract method 0xff5de458.
//
// Solidity: function whitelistCap(address ) view returns(bool)
func (_Genesis *GenesisSession) WhitelistCap(arg0 common.Address) (bool, error) {
	return _Genesis.Contract.WhitelistCap(&_Genesis.CallOpts, arg0)
}

// WhitelistCap is a free data retrieval call binding the contract method 0xff5de458.
//
// Solidity: function whitelistCap(address ) view returns(bool)
func (_Genesis *GenesisCallerSession) WhitelistCap(arg0 common.Address) (bool, error) {
	return _Genesis.Contract.WhitelistCap(&_Genesis.CallOpts, arg0)
}

// WhitelistRoot is a free data retrieval call binding the contract method 0x386bfc98.
//
// Solidity: function whitelistRoot() view returns(bytes32)
func (_Genesis *GenesisCaller) WhitelistRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "whitelistRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WhitelistRoot is a free data retrieval call binding the contract method 0x386bfc98.
//
// Solidity: function whitelistRoot() view returns(bytes32)
func (_Genesis *GenesisSession) WhitelistRoot() ([32]byte, error) {
	return _Genesis.Contract.WhitelistRoot(&_Genesis.CallOpts)
}

// WhitelistRoot is a free data retrieval call binding the contract method 0x386bfc98.
//
// Solidity: function whitelistRoot() view returns(bytes32)
func (_Genesis *GenesisCallerSession) WhitelistRoot() ([32]byte, error) {
	return _Genesis.Contract.WhitelistRoot(&_Genesis.CallOpts)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_Genesis *GenesisCaller) WhitelistedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Genesis.contract.Call(opts, &out, "whitelistedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_Genesis *GenesisSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _Genesis.Contract.WhitelistedTokens(&_Genesis.CallOpts, arg0)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_Genesis *GenesisCallerSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _Genesis.Contract.WhitelistedTokens(&_Genesis.CallOpts, arg0)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address _to, uint256 _tier) returns()
func (_Genesis *GenesisTransactor) AdminMint(opts *bind.TransactOpts, _to common.Address, _tier *big.Int) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "adminMint", _to, _tier)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address _to, uint256 _tier) returns()
func (_Genesis *GenesisSession) AdminMint(_to common.Address, _tier *big.Int) (*types.Transaction, error) {
	return _Genesis.Contract.AdminMint(&_Genesis.TransactOpts, _to, _tier)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address _to, uint256 _tier) returns()
func (_Genesis *GenesisTransactorSession) AdminMint(_to common.Address, _tier *big.Int) (*types.Transaction, error) {
	return _Genesis.Contract.AdminMint(&_Genesis.TransactOpts, _to, _tier)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Genesis *GenesisTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Genesis *GenesisSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Genesis.Contract.Approve(&_Genesis.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Genesis *GenesisTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Genesis.Contract.Approve(&_Genesis.TransactOpts, to, tokenId)
}

// PublicMint is a paid mutator transaction binding the contract method 0x32a93a3a.
//
// Solidity: function publicMint(address _token) returns(uint256)
func (_Genesis *GenesisTransactor) PublicMint(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "publicMint", _token)
}

// PublicMint is a paid mutator transaction binding the contract method 0x32a93a3a.
//
// Solidity: function publicMint(address _token) returns(uint256)
func (_Genesis *GenesisSession) PublicMint(_token common.Address) (*types.Transaction, error) {
	return _Genesis.Contract.PublicMint(&_Genesis.TransactOpts, _token)
}

// PublicMint is a paid mutator transaction binding the contract method 0x32a93a3a.
//
// Solidity: function publicMint(address _token) returns(uint256)
func (_Genesis *GenesisTransactorSession) PublicMint(_token common.Address) (*types.Transaction, error) {
	return _Genesis.Contract.PublicMint(&_Genesis.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Genesis *GenesisTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Genesis *GenesisSession) RenounceOwnership() (*types.Transaction, error) {
	return _Genesis.Contract.RenounceOwnership(&_Genesis.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Genesis *GenesisTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Genesis.Contract.RenounceOwnership(&_Genesis.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Genesis *GenesisTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Genesis *GenesisSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Genesis.Contract.SafeTransferFrom(&_Genesis.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Genesis *GenesisTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Genesis.Contract.SafeTransferFrom(&_Genesis.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Genesis *GenesisTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Genesis *GenesisSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Genesis.Contract.SafeTransferFrom0(&_Genesis.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Genesis *GenesisTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Genesis.Contract.SafeTransferFrom0(&_Genesis.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Genesis *GenesisTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Genesis *GenesisSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Genesis.Contract.SetApprovalForAll(&_Genesis.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Genesis *GenesisTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Genesis.Contract.SetApprovalForAll(&_Genesis.TransactOpts, operator, approved)
}

// SetWhitelistRoot is a paid mutator transaction binding the contract method 0xf5aa406d.
//
// Solidity: function setWhitelistRoot(bytes32 _root) returns()
func (_Genesis *GenesisTransactor) SetWhitelistRoot(opts *bind.TransactOpts, _root [32]byte) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "setWhitelistRoot", _root)
}

// SetWhitelistRoot is a paid mutator transaction binding the contract method 0xf5aa406d.
//
// Solidity: function setWhitelistRoot(bytes32 _root) returns()
func (_Genesis *GenesisSession) SetWhitelistRoot(_root [32]byte) (*types.Transaction, error) {
	return _Genesis.Contract.SetWhitelistRoot(&_Genesis.TransactOpts, _root)
}

// SetWhitelistRoot is a paid mutator transaction binding the contract method 0xf5aa406d.
//
// Solidity: function setWhitelistRoot(bytes32 _root) returns()
func (_Genesis *GenesisTransactorSession) SetWhitelistRoot(_root [32]byte) (*types.Transaction, error) {
	return _Genesis.Contract.SetWhitelistRoot(&_Genesis.TransactOpts, _root)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Genesis *GenesisTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Genesis *GenesisSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Genesis.Contract.TransferFrom(&_Genesis.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Genesis *GenesisTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Genesis.Contract.TransferFrom(&_Genesis.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Genesis *GenesisTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Genesis *GenesisSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Genesis.Contract.TransferOwnership(&_Genesis.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Genesis *GenesisTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Genesis.Contract.TransferOwnership(&_Genesis.TransactOpts, newOwner)
}

// UpgradeToken is a paid mutator transaction binding the contract method 0xc9abb2e9.
//
// Solidity: function upgradeToken(address _token, uint256 _levels) returns()
func (_Genesis *GenesisTransactor) UpgradeToken(opts *bind.TransactOpts, _token common.Address, _levels *big.Int) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "upgradeToken", _token, _levels)
}

// UpgradeToken is a paid mutator transaction binding the contract method 0xc9abb2e9.
//
// Solidity: function upgradeToken(address _token, uint256 _levels) returns()
func (_Genesis *GenesisSession) UpgradeToken(_token common.Address, _levels *big.Int) (*types.Transaction, error) {
	return _Genesis.Contract.UpgradeToken(&_Genesis.TransactOpts, _token, _levels)
}

// UpgradeToken is a paid mutator transaction binding the contract method 0xc9abb2e9.
//
// Solidity: function upgradeToken(address _token, uint256 _levels) returns()
func (_Genesis *GenesisTransactorSession) UpgradeToken(_token common.Address, _levels *big.Int) (*types.Transaction, error) {
	return _Genesis.Contract.UpgradeToken(&_Genesis.TransactOpts, _token, _levels)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0x372f657c.
//
// Solidity: function whitelistMint(bytes32[] _proof) returns(uint256)
func (_Genesis *GenesisTransactor) WhitelistMint(opts *bind.TransactOpts, _proof [][32]byte) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "whitelistMint", _proof)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0x372f657c.
//
// Solidity: function whitelistMint(bytes32[] _proof) returns(uint256)
func (_Genesis *GenesisSession) WhitelistMint(_proof [][32]byte) (*types.Transaction, error) {
	return _Genesis.Contract.WhitelistMint(&_Genesis.TransactOpts, _proof)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0x372f657c.
//
// Solidity: function whitelistMint(bytes32[] _proof) returns(uint256)
func (_Genesis *GenesisTransactorSession) WhitelistMint(_proof [][32]byte) (*types.Transaction, error) {
	return _Genesis.Contract.WhitelistMint(&_Genesis.TransactOpts, _proof)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x6247f6f2.
//
// Solidity: function whitelistToken(address _token) returns()
func (_Genesis *GenesisTransactor) WhitelistToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "whitelistToken", _token)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x6247f6f2.
//
// Solidity: function whitelistToken(address _token) returns()
func (_Genesis *GenesisSession) WhitelistToken(_token common.Address) (*types.Transaction, error) {
	return _Genesis.Contract.WhitelistToken(&_Genesis.TransactOpts, _token)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x6247f6f2.
//
// Solidity: function whitelistToken(address _token) returns()
func (_Genesis *GenesisTransactorSession) WhitelistToken(_token common.Address) (*types.Transaction, error) {
	return _Genesis.Contract.WhitelistToken(&_Genesis.TransactOpts, _token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_Genesis *GenesisTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Genesis.contract.Transact(opts, "withdraw", _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_Genesis *GenesisSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Genesis.Contract.Withdraw(&_Genesis.TransactOpts, _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_Genesis *GenesisTransactorSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Genesis.Contract.Withdraw(&_Genesis.TransactOpts, _token, _amount)
}

// GenesisApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Genesis contract.
type GenesisApprovalIterator struct {
	Event *GenesisApproval // Event containing the contract specifics and raw log

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
func (it *GenesisApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenesisApproval)
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
		it.Event = new(GenesisApproval)
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
func (it *GenesisApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenesisApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenesisApproval represents a Approval event raised by the Genesis contract.
type GenesisApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Genesis *GenesisFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*GenesisApprovalIterator, error) {

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

	logs, sub, err := _Genesis.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GenesisApprovalIterator{contract: _Genesis.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Genesis *GenesisFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GenesisApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Genesis.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenesisApproval)
				if err := _Genesis.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Genesis *GenesisFilterer) ParseApproval(log types.Log) (*GenesisApproval, error) {
	event := new(GenesisApproval)
	if err := _Genesis.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenesisApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Genesis contract.
type GenesisApprovalForAllIterator struct {
	Event *GenesisApprovalForAll // Event containing the contract specifics and raw log

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
func (it *GenesisApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenesisApprovalForAll)
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
		it.Event = new(GenesisApprovalForAll)
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
func (it *GenesisApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenesisApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenesisApprovalForAll represents a ApprovalForAll event raised by the Genesis contract.
type GenesisApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Genesis *GenesisFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*GenesisApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Genesis.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &GenesisApprovalForAllIterator{contract: _Genesis.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Genesis *GenesisFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *GenesisApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Genesis.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenesisApprovalForAll)
				if err := _Genesis.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Genesis *GenesisFilterer) ParseApprovalForAll(log types.Log) (*GenesisApprovalForAll, error) {
	event := new(GenesisApprovalForAll)
	if err := _Genesis.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenesisMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Genesis contract.
type GenesisMintIterator struct {
	Event *GenesisMint // Event containing the contract specifics and raw log

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
func (it *GenesisMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenesisMint)
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
		it.Event = new(GenesisMint)
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
func (it *GenesisMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenesisMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenesisMint represents a Mint event raised by the Genesis contract.
type GenesisMint struct {
	Id   *big.Int
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xf3cea5493d790af0133817606f7350a91d7f154ea52eaa79d179d4d231e50102.
//
// Solidity: event Mint(uint256 id, address user)
func (_Genesis *GenesisFilterer) FilterMint(opts *bind.FilterOpts) (*GenesisMintIterator, error) {

	logs, sub, err := _Genesis.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &GenesisMintIterator{contract: _Genesis.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xf3cea5493d790af0133817606f7350a91d7f154ea52eaa79d179d4d231e50102.
//
// Solidity: event Mint(uint256 id, address user)
func (_Genesis *GenesisFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *GenesisMint) (event.Subscription, error) {

	logs, sub, err := _Genesis.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenesisMint)
				if err := _Genesis.contract.UnpackLog(event, "Mint", log); err != nil {
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
func (_Genesis *GenesisFilterer) ParseMint(log types.Log) (*GenesisMint, error) {
	event := new(GenesisMint)
	if err := _Genesis.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenesisOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Genesis contract.
type GenesisOwnershipTransferredIterator struct {
	Event *GenesisOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GenesisOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenesisOwnershipTransferred)
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
		it.Event = new(GenesisOwnershipTransferred)
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
func (it *GenesisOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenesisOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenesisOwnershipTransferred represents a OwnershipTransferred event raised by the Genesis contract.
type GenesisOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Genesis *GenesisFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GenesisOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Genesis.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GenesisOwnershipTransferredIterator{contract: _Genesis.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Genesis *GenesisFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GenesisOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Genesis.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenesisOwnershipTransferred)
				if err := _Genesis.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Genesis *GenesisFilterer) ParseOwnershipTransferred(log types.Log) (*GenesisOwnershipTransferred, error) {
	event := new(GenesisOwnershipTransferred)
	if err := _Genesis.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenesisTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Genesis contract.
type GenesisTransferIterator struct {
	Event *GenesisTransfer // Event containing the contract specifics and raw log

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
func (it *GenesisTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenesisTransfer)
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
		it.Event = new(GenesisTransfer)
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
func (it *GenesisTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenesisTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenesisTransfer represents a Transfer event raised by the Genesis contract.
type GenesisTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Genesis *GenesisFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*GenesisTransferIterator, error) {

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

	logs, sub, err := _Genesis.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GenesisTransferIterator{contract: _Genesis.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Genesis *GenesisFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GenesisTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Genesis.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenesisTransfer)
				if err := _Genesis.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Genesis *GenesisFilterer) ParseTransfer(log types.Log) (*GenesisTransfer, error) {
	event := new(GenesisTransfer)
	if err := _Genesis.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenesisUpgradeIterator is returned from FilterUpgrade and is used to iterate over the raw logs and unpacked data for Upgrade events raised by the Genesis contract.
type GenesisUpgradeIterator struct {
	Event *GenesisUpgrade // Event containing the contract specifics and raw log

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
func (it *GenesisUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenesisUpgrade)
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
		it.Event = new(GenesisUpgrade)
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
func (it *GenesisUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenesisUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenesisUpgrade represents a Upgrade event raised by the Genesis contract.
type GenesisUpgrade struct {
	Id   *big.Int
	Tier *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpgrade is a free log retrieval operation binding the contract event 0xeb0ebb51128928d7b1a6419c52128a9319bfcb55f0adafea75afbf75f9f2f3e9.
//
// Solidity: event Upgrade(uint256 id, uint256 tier)
func (_Genesis *GenesisFilterer) FilterUpgrade(opts *bind.FilterOpts) (*GenesisUpgradeIterator, error) {

	logs, sub, err := _Genesis.contract.FilterLogs(opts, "Upgrade")
	if err != nil {
		return nil, err
	}
	return &GenesisUpgradeIterator{contract: _Genesis.contract, event: "Upgrade", logs: logs, sub: sub}, nil
}

// WatchUpgrade is a free log subscription operation binding the contract event 0xeb0ebb51128928d7b1a6419c52128a9319bfcb55f0adafea75afbf75f9f2f3e9.
//
// Solidity: event Upgrade(uint256 id, uint256 tier)
func (_Genesis *GenesisFilterer) WatchUpgrade(opts *bind.WatchOpts, sink chan<- *GenesisUpgrade) (event.Subscription, error) {

	logs, sub, err := _Genesis.contract.WatchLogs(opts, "Upgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenesisUpgrade)
				if err := _Genesis.contract.UnpackLog(event, "Upgrade", log); err != nil {
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
func (_Genesis *GenesisFilterer) ParseUpgrade(log types.Log) (*GenesisUpgrade, error) {
	event := new(GenesisUpgrade)
	if err := _Genesis.contract.UnpackLog(event, "Upgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
