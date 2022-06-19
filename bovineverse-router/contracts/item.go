// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ItemABI is the input ABI used to generate the binding from.
const ItemABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"checkItemEffect\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"checkTypeBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"itemId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"effect_\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"types\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"tradeable_\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"}],\"name\":\"editItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"itemAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"itemInfoes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"currentAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"tradeable\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"itemType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"itemId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"mintBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"itemId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"effect_\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"types\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"tradeable_\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"}],\"name\":\"newItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinter_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"itemId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinter_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"setMinterBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"setMyBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSuperMinter_\",\"type\":\"address\"}],\"name\":\"setSuperMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"itemId_\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Item is an auto generated Go binding around an Ethereum contract.
type Item struct {
	ItemCaller     // Read-only binding to the contract
	ItemTransactor // Write-only binding to the contract
	ItemFilterer   // Log filterer for contract events
}

// ItemCaller is an auto generated read-only Go binding around an Ethereum contract.
type ItemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ItemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ItemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ItemSession struct {
	Contract     *Item             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ItemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ItemCallerSession struct {
	Contract *ItemCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ItemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ItemTransactorSession struct {
	Contract     *ItemTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ItemRaw is an auto generated low-level Go binding around an Ethereum contract.
type ItemRaw struct {
	Contract *Item // Generic contract binding to access the raw methods on
}

// ItemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ItemCallerRaw struct {
	Contract *ItemCaller // Generic read-only contract binding to access the raw methods on
}

// ItemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ItemTransactorRaw struct {
	Contract *ItemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewItem creates a new instance of Item, bound to a specific deployed contract.
func NewItem(address common.Address, backend bind.ContractBackend) (*Item, error) {
	contract, err := bindItem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Item{ItemCaller: ItemCaller{contract: contract}, ItemTransactor: ItemTransactor{contract: contract}, ItemFilterer: ItemFilterer{contract: contract}}, nil
}

// NewItemCaller creates a new read-only instance of Item, bound to a specific deployed contract.
func NewItemCaller(address common.Address, caller bind.ContractCaller) (*ItemCaller, error) {
	contract, err := bindItem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ItemCaller{contract: contract}, nil
}

// NewItemTransactor creates a new write-only instance of Item, bound to a specific deployed contract.
func NewItemTransactor(address common.Address, transactor bind.ContractTransactor) (*ItemTransactor, error) {
	contract, err := bindItem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ItemTransactor{contract: contract}, nil
}

// NewItemFilterer creates a new log filterer instance of Item, bound to a specific deployed contract.
func NewItemFilterer(address common.Address, filterer bind.ContractFilterer) (*ItemFilterer, error) {
	contract, err := bindItem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ItemFilterer{contract: contract}, nil
}

// bindItem binds a generic wrapper to an already deployed contract.
func bindItem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ItemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Item *ItemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Item.Contract.ItemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Item *ItemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Item.Contract.ItemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Item *ItemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Item.Contract.ItemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Item *ItemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Item.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Item *ItemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Item.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Item *ItemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Item.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Item *ItemCaller) Admin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "admin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Item *ItemSession) Admin(arg0 common.Address) (bool, error) {
	return _Item.Contract.Admin(&_Item.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Item *ItemCallerSession) Admin(arg0 common.Address) (bool, error) {
	return _Item.Contract.Admin(&_Item.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Item *ItemCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Item *ItemSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Item.Contract.BalanceOf(&_Item.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Item *ItemCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Item.Contract.BalanceOf(&_Item.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Item *ItemCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Item *ItemSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Item.Contract.BalanceOfBatch(&_Item.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Item *ItemCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Item.Contract.BalanceOfBatch(&_Item.CallOpts, accounts, ids)
}

// Burned is a free data retrieval call binding the contract method 0x73f42561.
//
// Solidity: function burned() view returns(uint256)
func (_Item *ItemCaller) Burned(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "burned")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Burned is a free data retrieval call binding the contract method 0x73f42561.
//
// Solidity: function burned() view returns(uint256)
func (_Item *ItemSession) Burned() (*big.Int, error) {
	return _Item.Contract.Burned(&_Item.CallOpts)
}

// Burned is a free data retrieval call binding the contract method 0x73f42561.
//
// Solidity: function burned() view returns(uint256)
func (_Item *ItemCallerSession) Burned() (*big.Int, error) {
	return _Item.Contract.Burned(&_Item.CallOpts)
}

// CheckItemEffect is a free data retrieval call binding the contract method 0xce220cda.
//
// Solidity: function checkItemEffect(uint256 id_) view returns(uint256[3])
func (_Item *ItemCaller) CheckItemEffect(opts *bind.CallOpts, id_ *big.Int) ([3]*big.Int, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "checkItemEffect", id_)

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// CheckItemEffect is a free data retrieval call binding the contract method 0xce220cda.
//
// Solidity: function checkItemEffect(uint256 id_) view returns(uint256[3])
func (_Item *ItemSession) CheckItemEffect(id_ *big.Int) ([3]*big.Int, error) {
	return _Item.Contract.CheckItemEffect(&_Item.CallOpts, id_)
}

// CheckItemEffect is a free data retrieval call binding the contract method 0xce220cda.
//
// Solidity: function checkItemEffect(uint256 id_) view returns(uint256[3])
func (_Item *ItemCallerSession) CheckItemEffect(id_ *big.Int) ([3]*big.Int, error) {
	return _Item.Contract.CheckItemEffect(&_Item.CallOpts, id_)
}

// CheckTypeBatch is a free data retrieval call binding the contract method 0x649d0c57.
//
// Solidity: function checkTypeBatch(uint256[] ids) view returns(uint256[])
func (_Item *ItemCaller) CheckTypeBatch(opts *bind.CallOpts, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "checkTypeBatch", ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CheckTypeBatch is a free data retrieval call binding the contract method 0x649d0c57.
//
// Solidity: function checkTypeBatch(uint256[] ids) view returns(uint256[])
func (_Item *ItemSession) CheckTypeBatch(ids []*big.Int) ([]*big.Int, error) {
	return _Item.Contract.CheckTypeBatch(&_Item.CallOpts, ids)
}

// CheckTypeBatch is a free data retrieval call binding the contract method 0x649d0c57.
//
// Solidity: function checkTypeBatch(uint256[] ids) view returns(uint256[])
func (_Item *ItemCallerSession) CheckTypeBatch(ids []*big.Int) ([]*big.Int, error) {
	return _Item.Contract.CheckTypeBatch(&_Item.CallOpts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Item *ItemCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Item *ItemSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Item.Contract.IsApprovedForAll(&_Item.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Item *ItemCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Item.Contract.IsApprovedForAll(&_Item.CallOpts, account, operator)
}

// ItemAmount is a free data retrieval call binding the contract method 0xc0373561.
//
// Solidity: function itemAmount() view returns(uint256)
func (_Item *ItemCaller) ItemAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "itemAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ItemAmount is a free data retrieval call binding the contract method 0xc0373561.
//
// Solidity: function itemAmount() view returns(uint256)
func (_Item *ItemSession) ItemAmount() (*big.Int, error) {
	return _Item.Contract.ItemAmount(&_Item.CallOpts)
}

// ItemAmount is a free data retrieval call binding the contract method 0xc0373561.
//
// Solidity: function itemAmount() view returns(uint256)
func (_Item *ItemCallerSession) ItemAmount() (*big.Int, error) {
	return _Item.Contract.ItemAmount(&_Item.CallOpts)
}

// ItemInfoes is a free data retrieval call binding the contract method 0x86e32a77.
//
// Solidity: function itemInfoes(uint256 ) view returns(uint256 itemId, string name, uint256 currentAmount, uint256 burnedAmount, uint256 maxAmount, bool tradeable, string tokenURI)
func (_Item *ItemCaller) ItemInfoes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ItemId        *big.Int
	Name          string
	CurrentAmount *big.Int
	BurnedAmount  *big.Int
	MaxAmount     *big.Int
	Tradeable     bool
	TokenURI      string
}, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "itemInfoes", arg0)

	outstruct := new(struct {
		ItemId        *big.Int
		Name          string
		CurrentAmount *big.Int
		BurnedAmount  *big.Int
		MaxAmount     *big.Int
		Tradeable     bool
		TokenURI      string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ItemId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.CurrentAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BurnedAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Tradeable = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.TokenURI = *abi.ConvertType(out[6], new(string)).(*string)

	return *outstruct, err

}

// ItemInfoes is a free data retrieval call binding the contract method 0x86e32a77.
//
// Solidity: function itemInfoes(uint256 ) view returns(uint256 itemId, string name, uint256 currentAmount, uint256 burnedAmount, uint256 maxAmount, bool tradeable, string tokenURI)
func (_Item *ItemSession) ItemInfoes(arg0 *big.Int) (struct {
	ItemId        *big.Int
	Name          string
	CurrentAmount *big.Int
	BurnedAmount  *big.Int
	MaxAmount     *big.Int
	Tradeable     bool
	TokenURI      string
}, error) {
	return _Item.Contract.ItemInfoes(&_Item.CallOpts, arg0)
}

// ItemInfoes is a free data retrieval call binding the contract method 0x86e32a77.
//
// Solidity: function itemInfoes(uint256 ) view returns(uint256 itemId, string name, uint256 currentAmount, uint256 burnedAmount, uint256 maxAmount, bool tradeable, string tokenURI)
func (_Item *ItemCallerSession) ItemInfoes(arg0 *big.Int) (struct {
	ItemId        *big.Int
	Name          string
	CurrentAmount *big.Int
	BurnedAmount  *big.Int
	MaxAmount     *big.Int
	Tradeable     bool
	TokenURI      string
}, error) {
	return _Item.Contract.ItemInfoes(&_Item.CallOpts, arg0)
}

// ItemType is a free data retrieval call binding the contract method 0xac09aa0f.
//
// Solidity: function itemType(uint256 ) view returns(uint256)
func (_Item *ItemCaller) ItemType(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "itemType", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ItemType is a free data retrieval call binding the contract method 0xac09aa0f.
//
// Solidity: function itemType(uint256 ) view returns(uint256)
func (_Item *ItemSession) ItemType(arg0 *big.Int) (*big.Int, error) {
	return _Item.Contract.ItemType(&_Item.CallOpts, arg0)
}

// ItemType is a free data retrieval call binding the contract method 0xac09aa0f.
//
// Solidity: function itemType(uint256 ) view returns(uint256)
func (_Item *ItemCallerSession) ItemType(arg0 *big.Int) (*big.Int, error) {
	return _Item.Contract.ItemType(&_Item.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(uint256)
func (_Item *ItemCaller) Minters(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "minters", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(uint256)
func (_Item *ItemSession) Minters(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Item.Contract.Minters(&_Item.CallOpts, arg0, arg1)
}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(uint256)
func (_Item *ItemCallerSession) Minters(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Item.Contract.Minters(&_Item.CallOpts, arg0, arg1)
}

// MyBaseURI is a free data retrieval call binding the contract method 0x372a3b82.
//
// Solidity: function myBaseURI() view returns(string)
func (_Item *ItemCaller) MyBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "myBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MyBaseURI is a free data retrieval call binding the contract method 0x372a3b82.
//
// Solidity: function myBaseURI() view returns(string)
func (_Item *ItemSession) MyBaseURI() (string, error) {
	return _Item.Contract.MyBaseURI(&_Item.CallOpts)
}

// MyBaseURI is a free data retrieval call binding the contract method 0x372a3b82.
//
// Solidity: function myBaseURI() view returns(string)
func (_Item *ItemCallerSession) MyBaseURI() (string, error) {
	return _Item.Contract.MyBaseURI(&_Item.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Item *ItemCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Item *ItemSession) Name() (string, error) {
	return _Item.Contract.Name(&_Item.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Item *ItemCallerSession) Name() (string, error) {
	return _Item.Contract.Name(&_Item.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Item *ItemCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Item *ItemSession) Owner() (common.Address, error) {
	return _Item.Contract.Owner(&_Item.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Item *ItemCallerSession) Owner() (common.Address, error) {
	return _Item.Contract.Owner(&_Item.CallOpts)
}

// SuperMinter is a free data retrieval call binding the contract method 0x62e42cb0.
//
// Solidity: function superMinter() view returns(address)
func (_Item *ItemCaller) SuperMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "superMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperMinter is a free data retrieval call binding the contract method 0x62e42cb0.
//
// Solidity: function superMinter() view returns(address)
func (_Item *ItemSession) SuperMinter() (common.Address, error) {
	return _Item.Contract.SuperMinter(&_Item.CallOpts)
}

// SuperMinter is a free data retrieval call binding the contract method 0x62e42cb0.
//
// Solidity: function superMinter() view returns(address)
func (_Item *ItemCallerSession) SuperMinter() (common.Address, error) {
	return _Item.Contract.SuperMinter(&_Item.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Item *ItemCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Item *ItemSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Item.Contract.SupportsInterface(&_Item.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Item *ItemCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Item.Contract.SupportsInterface(&_Item.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Item *ItemCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Item *ItemSession) Symbol() (string, error) {
	return _Item.Contract.Symbol(&_Item.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Item *ItemCallerSession) Symbol() (string, error) {
	return _Item.Contract.Symbol(&_Item.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 itemId_) view returns(string)
func (_Item *ItemCaller) TokenURI(opts *bind.CallOpts, itemId_ *big.Int) (string, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "tokenURI", itemId_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 itemId_) view returns(string)
func (_Item *ItemSession) TokenURI(itemId_ *big.Int) (string, error) {
	return _Item.Contract.TokenURI(&_Item.CallOpts, itemId_)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 itemId_) view returns(string)
func (_Item *ItemCallerSession) TokenURI(itemId_ *big.Int) (string, error) {
	return _Item.Contract.TokenURI(&_Item.CallOpts, itemId_)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Item *ItemCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Item *ItemSession) Uri(arg0 *big.Int) (string, error) {
	return _Item.Contract.Uri(&_Item.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Item *ItemCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _Item.Contract.Uri(&_Item.CallOpts, arg0)
}

// UserBurn is a free data retrieval call binding the contract method 0xdfcea3b4.
//
// Solidity: function userBurn(address , uint256 ) view returns(uint256)
func (_Item *ItemCaller) UserBurn(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Item.contract.Call(opts, &out, "userBurn", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBurn is a free data retrieval call binding the contract method 0xdfcea3b4.
//
// Solidity: function userBurn(address , uint256 ) view returns(uint256)
func (_Item *ItemSession) UserBurn(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Item.Contract.UserBurn(&_Item.CallOpts, arg0, arg1)
}

// UserBurn is a free data retrieval call binding the contract method 0xdfcea3b4.
//
// Solidity: function userBurn(address , uint256 ) view returns(uint256)
func (_Item *ItemCallerSession) UserBurn(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Item.Contract.UserBurn(&_Item.CallOpts, arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_Item *ItemTransactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "burn", account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_Item *ItemSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Item.Contract.Burn(&_Item.TransactOpts, account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_Item *ItemTransactorSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Item.Contract.Burn(&_Item.TransactOpts, account, id, value)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_Item *ItemTransactor) BurnBatch(opts *bind.TransactOpts, account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "burnBatch", account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_Item *ItemSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Item.Contract.BurnBatch(&_Item.TransactOpts, account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_Item *ItemTransactorSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Item.Contract.BurnBatch(&_Item.TransactOpts, account, ids, values)
}

// EditItem is a paid mutator transaction binding the contract method 0x40d97d39.
//
// Solidity: function editItem(string name_, uint256 itemId_, uint256 maxAmount_, uint256[3] effect_, uint256 types, bool tradeable_, string tokenURI_) returns()
func (_Item *ItemTransactor) EditItem(opts *bind.TransactOpts, name_ string, itemId_ *big.Int, maxAmount_ *big.Int, effect_ [3]*big.Int, types *big.Int, tradeable_ bool, tokenURI_ string) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "editItem", name_, itemId_, maxAmount_, effect_, types, tradeable_, tokenURI_)
}

// EditItem is a paid mutator transaction binding the contract method 0x40d97d39.
//
// Solidity: function editItem(string name_, uint256 itemId_, uint256 maxAmount_, uint256[3] effect_, uint256 types, bool tradeable_, string tokenURI_) returns()
func (_Item *ItemSession) EditItem(name_ string, itemId_ *big.Int, maxAmount_ *big.Int, effect_ [3]*big.Int, types *big.Int, tradeable_ bool, tokenURI_ string) (*types.Transaction, error) {
	return _Item.Contract.EditItem(&_Item.TransactOpts, name_, itemId_, maxAmount_, effect_, types, tradeable_, tokenURI_)
}

// EditItem is a paid mutator transaction binding the contract method 0x40d97d39.
//
// Solidity: function editItem(string name_, uint256 itemId_, uint256 maxAmount_, uint256[3] effect_, uint256 types, bool tradeable_, string tokenURI_) returns()
func (_Item *ItemTransactorSession) EditItem(name_ string, itemId_ *big.Int, maxAmount_ *big.Int, effect_ [3]*big.Int, types *big.Int, tradeable_ bool, tokenURI_ string) (*types.Transaction, error) {
	return _Item.Contract.EditItem(&_Item.TransactOpts, name_, itemId_, maxAmount_, effect_, types, tradeable_, tokenURI_)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to_, uint256 itemId_, uint256 amount_) returns(bool)
func (_Item *ItemTransactor) Mint(opts *bind.TransactOpts, to_ common.Address, itemId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "mint", to_, itemId_, amount_)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to_, uint256 itemId_, uint256 amount_) returns(bool)
func (_Item *ItemSession) Mint(to_ common.Address, itemId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Item.Contract.Mint(&_Item.TransactOpts, to_, itemId_, amount_)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to_, uint256 itemId_, uint256 amount_) returns(bool)
func (_Item *ItemTransactorSession) Mint(to_ common.Address, itemId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Item.Contract.Mint(&_Item.TransactOpts, to_, itemId_, amount_)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd81d0a15.
//
// Solidity: function mintBatch(address to_, uint256[] ids_, uint256[] amounts_) returns(bool)
func (_Item *ItemTransactor) MintBatch(opts *bind.TransactOpts, to_ common.Address, ids_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "mintBatch", to_, ids_, amounts_)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd81d0a15.
//
// Solidity: function mintBatch(address to_, uint256[] ids_, uint256[] amounts_) returns(bool)
func (_Item *ItemSession) MintBatch(to_ common.Address, ids_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Item.Contract.MintBatch(&_Item.TransactOpts, to_, ids_, amounts_)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd81d0a15.
//
// Solidity: function mintBatch(address to_, uint256[] ids_, uint256[] amounts_) returns(bool)
func (_Item *ItemTransactorSession) MintBatch(to_ common.Address, ids_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Item.Contract.MintBatch(&_Item.TransactOpts, to_, ids_, amounts_)
}

// NewItem is a paid mutator transaction binding the contract method 0x01eb7014.
//
// Solidity: function newItem(string name_, uint256 itemId_, uint256 maxAmount_, uint256[3] effect_, uint256 types, bool tradeable_, string tokenURI_) returns()
func (_Item *ItemTransactor) NewItem(opts *bind.TransactOpts, name_ string, itemId_ *big.Int, maxAmount_ *big.Int, effect_ [3]*big.Int, types *big.Int, tradeable_ bool, tokenURI_ string) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "newItem", name_, itemId_, maxAmount_, effect_, types, tradeable_, tokenURI_)
}

// NewItem is a paid mutator transaction binding the contract method 0x01eb7014.
//
// Solidity: function newItem(string name_, uint256 itemId_, uint256 maxAmount_, uint256[3] effect_, uint256 types, bool tradeable_, string tokenURI_) returns()
func (_Item *ItemSession) NewItem(name_ string, itemId_ *big.Int, maxAmount_ *big.Int, effect_ [3]*big.Int, types *big.Int, tradeable_ bool, tokenURI_ string) (*types.Transaction, error) {
	return _Item.Contract.NewItem(&_Item.TransactOpts, name_, itemId_, maxAmount_, effect_, types, tradeable_, tokenURI_)
}

// NewItem is a paid mutator transaction binding the contract method 0x01eb7014.
//
// Solidity: function newItem(string name_, uint256 itemId_, uint256 maxAmount_, uint256[3] effect_, uint256 types, bool tradeable_, string tokenURI_) returns()
func (_Item *ItemTransactorSession) NewItem(name_ string, itemId_ *big.Int, maxAmount_ *big.Int, effect_ [3]*big.Int, types *big.Int, tradeable_ bool, tokenURI_ string) (*types.Transaction, error) {
	return _Item.Contract.NewItem(&_Item.TransactOpts, name_, itemId_, maxAmount_, effect_, types, tradeable_, tokenURI_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Item *ItemTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Item *ItemSession) RenounceOwnership() (*types.Transaction, error) {
	return _Item.Contract.RenounceOwnership(&_Item.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Item *ItemTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Item.Contract.RenounceOwnership(&_Item.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Item *ItemTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Item *ItemSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Item.Contract.SafeBatchTransferFrom(&_Item.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Item *ItemTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Item.Contract.SafeBatchTransferFrom(&_Item.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Item *ItemTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Item *ItemSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Item.Contract.SafeTransferFrom(&_Item.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Item *ItemTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Item.Contract.SafeTransferFrom(&_Item.TransactOpts, from, to, id, amount, data)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr, bool b) returns()
func (_Item *ItemTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address, b bool) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "setAdmin", addr, b)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr, bool b) returns()
func (_Item *ItemSession) SetAdmin(addr common.Address, b bool) (*types.Transaction, error) {
	return _Item.Contract.SetAdmin(&_Item.TransactOpts, addr, b)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr, bool b) returns()
func (_Item *ItemTransactorSession) SetAdmin(addr common.Address, b bool) (*types.Transaction, error) {
	return _Item.Contract.SetAdmin(&_Item.TransactOpts, addr, b)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Item *ItemTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Item *ItemSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Item.Contract.SetApprovalForAll(&_Item.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Item *ItemTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Item.Contract.SetApprovalForAll(&_Item.TransactOpts, operator, approved)
}

// SetMinter is a paid mutator transaction binding the contract method 0xbb2fa28c.
//
// Solidity: function setMinter(address newMinter_, uint256 itemId_, uint256 amount_) returns()
func (_Item *ItemTransactor) SetMinter(opts *bind.TransactOpts, newMinter_ common.Address, itemId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "setMinter", newMinter_, itemId_, amount_)
}

// SetMinter is a paid mutator transaction binding the contract method 0xbb2fa28c.
//
// Solidity: function setMinter(address newMinter_, uint256 itemId_, uint256 amount_) returns()
func (_Item *ItemSession) SetMinter(newMinter_ common.Address, itemId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Item.Contract.SetMinter(&_Item.TransactOpts, newMinter_, itemId_, amount_)
}

// SetMinter is a paid mutator transaction binding the contract method 0xbb2fa28c.
//
// Solidity: function setMinter(address newMinter_, uint256 itemId_, uint256 amount_) returns()
func (_Item *ItemTransactorSession) SetMinter(newMinter_ common.Address, itemId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Item.Contract.SetMinter(&_Item.TransactOpts, newMinter_, itemId_, amount_)
}

// SetMinterBatch is a paid mutator transaction binding the contract method 0xf882aa8d.
//
// Solidity: function setMinterBatch(address newMinter_, uint256[] ids_, uint256[] amounts_) returns(bool)
func (_Item *ItemTransactor) SetMinterBatch(opts *bind.TransactOpts, newMinter_ common.Address, ids_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "setMinterBatch", newMinter_, ids_, amounts_)
}

// SetMinterBatch is a paid mutator transaction binding the contract method 0xf882aa8d.
//
// Solidity: function setMinterBatch(address newMinter_, uint256[] ids_, uint256[] amounts_) returns(bool)
func (_Item *ItemSession) SetMinterBatch(newMinter_ common.Address, ids_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Item.Contract.SetMinterBatch(&_Item.TransactOpts, newMinter_, ids_, amounts_)
}

// SetMinterBatch is a paid mutator transaction binding the contract method 0xf882aa8d.
//
// Solidity: function setMinterBatch(address newMinter_, uint256[] ids_, uint256[] amounts_) returns(bool)
func (_Item *ItemTransactorSession) SetMinterBatch(newMinter_ common.Address, ids_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Item.Contract.SetMinterBatch(&_Item.TransactOpts, newMinter_, ids_, amounts_)
}

// SetMyBaseURI is a paid mutator transaction binding the contract method 0x5dcce234.
//
// Solidity: function setMyBaseURI(string uri_) returns()
func (_Item *ItemTransactor) SetMyBaseURI(opts *bind.TransactOpts, uri_ string) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "setMyBaseURI", uri_)
}

// SetMyBaseURI is a paid mutator transaction binding the contract method 0x5dcce234.
//
// Solidity: function setMyBaseURI(string uri_) returns()
func (_Item *ItemSession) SetMyBaseURI(uri_ string) (*types.Transaction, error) {
	return _Item.Contract.SetMyBaseURI(&_Item.TransactOpts, uri_)
}

// SetMyBaseURI is a paid mutator transaction binding the contract method 0x5dcce234.
//
// Solidity: function setMyBaseURI(string uri_) returns()
func (_Item *ItemTransactorSession) SetMyBaseURI(uri_ string) (*types.Transaction, error) {
	return _Item.Contract.SetMyBaseURI(&_Item.TransactOpts, uri_)
}

// SetSuperMinter is a paid mutator transaction binding the contract method 0xb41d74d8.
//
// Solidity: function setSuperMinter(address newSuperMinter_) returns()
func (_Item *ItemTransactor) SetSuperMinter(opts *bind.TransactOpts, newSuperMinter_ common.Address) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "setSuperMinter", newSuperMinter_)
}

// SetSuperMinter is a paid mutator transaction binding the contract method 0xb41d74d8.
//
// Solidity: function setSuperMinter(address newSuperMinter_) returns()
func (_Item *ItemSession) SetSuperMinter(newSuperMinter_ common.Address) (*types.Transaction, error) {
	return _Item.Contract.SetSuperMinter(&_Item.TransactOpts, newSuperMinter_)
}

// SetSuperMinter is a paid mutator transaction binding the contract method 0xb41d74d8.
//
// Solidity: function setSuperMinter(address newSuperMinter_) returns()
func (_Item *ItemTransactorSession) SetSuperMinter(newSuperMinter_ common.Address) (*types.Transaction, error) {
	return _Item.Contract.SetSuperMinter(&_Item.TransactOpts, newSuperMinter_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Item *ItemTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Item.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Item *ItemSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Item.Contract.TransferOwnership(&_Item.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Item *ItemTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Item.Contract.TransferOwnership(&_Item.TransactOpts, newOwner)
}

// ItemApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Item contract.
type ItemApprovalForAllIterator struct {
	Event *ItemApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ItemApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItemApprovalForAll)
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
		it.Event = new(ItemApprovalForAll)
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
func (it *ItemApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItemApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItemApprovalForAll represents a ApprovalForAll event raised by the Item contract.
type ItemApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Item *ItemFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ItemApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Item.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ItemApprovalForAllIterator{contract: _Item.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Item *ItemFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ItemApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Item.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItemApprovalForAll)
				if err := _Item.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Item *ItemFilterer) ParseApprovalForAll(log types.Log) (*ItemApprovalForAll, error) {
	event := new(ItemApprovalForAll)
	if err := _Item.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ItemOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Item contract.
type ItemOwnershipTransferredIterator struct {
	Event *ItemOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ItemOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItemOwnershipTransferred)
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
		it.Event = new(ItemOwnershipTransferred)
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
func (it *ItemOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItemOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItemOwnershipTransferred represents a OwnershipTransferred event raised by the Item contract.
type ItemOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Item *ItemFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ItemOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Item.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ItemOwnershipTransferredIterator{contract: _Item.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Item *ItemFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ItemOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Item.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItemOwnershipTransferred)
				if err := _Item.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Item *ItemFilterer) ParseOwnershipTransferred(log types.Log) (*ItemOwnershipTransferred, error) {
	event := new(ItemOwnershipTransferred)
	if err := _Item.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ItemTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Item contract.
type ItemTransferBatchIterator struct {
	Event *ItemTransferBatch // Event containing the contract specifics and raw log

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
func (it *ItemTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItemTransferBatch)
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
		it.Event = new(ItemTransferBatch)
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
func (it *ItemTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItemTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItemTransferBatch represents a TransferBatch event raised by the Item contract.
type ItemTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Item *ItemFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ItemTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Item.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ItemTransferBatchIterator{contract: _Item.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Item *ItemFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ItemTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Item.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItemTransferBatch)
				if err := _Item.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Item *ItemFilterer) ParseTransferBatch(log types.Log) (*ItemTransferBatch, error) {
	event := new(ItemTransferBatch)
	if err := _Item.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ItemTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Item contract.
type ItemTransferSingleIterator struct {
	Event *ItemTransferSingle // Event containing the contract specifics and raw log

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
func (it *ItemTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItemTransferSingle)
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
		it.Event = new(ItemTransferSingle)
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
func (it *ItemTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItemTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItemTransferSingle represents a TransferSingle event raised by the Item contract.
type ItemTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Item *ItemFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ItemTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Item.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ItemTransferSingleIterator{contract: _Item.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Item *ItemFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ItemTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Item.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItemTransferSingle)
				if err := _Item.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Item *ItemFilterer) ParseTransferSingle(log types.Log) (*ItemTransferSingle, error) {
	event := new(ItemTransferSingle)
	if err := _Item.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ItemURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Item contract.
type ItemURIIterator struct {
	Event *ItemURI // Event containing the contract specifics and raw log

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
func (it *ItemURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItemURI)
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
		it.Event = new(ItemURI)
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
func (it *ItemURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItemURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItemURI represents a URI event raised by the Item contract.
type ItemURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Item *ItemFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ItemURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Item.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ItemURIIterator{contract: _Item.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Item *ItemFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ItemURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Item.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItemURI)
				if err := _Item.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Item *ItemFilterer) ParseURI(log types.Log) (*ItemURI, error) {
	event := new(ItemURI)
	if err := _Item.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
