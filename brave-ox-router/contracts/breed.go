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

// BreedCenterABI is the input ABI used to generate the binding from.
const BreedCenterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"targetTokenID\",\"type\":\"uint256\"}],\"name\":\"Mate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OffSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"UpLoad\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BVG\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BVT\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"box\",\"outputs\":[{\"internalType\":\"contractIBOX\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cattle\",\"outputs\":[{\"internalType\":\"contractICOW\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"list\",\"type\":\"uint256[]\"}],\"name\":\"checkMatingTimeBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"out\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"energyCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"getUserUploadList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"index\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastMatingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"myTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetTokenID\",\"type\":\"uint256\"}],\"name\":\"mating\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"matingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"offSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"onSale\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"photo\",\"outputs\":[{\"internalType\":\"contractIProfilePhoto\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"planet\",\"outputs\":[{\"internalType\":\"contractIPlanet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId2\",\"type\":\"uint256\"}],\"name\":\"selfMating\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"box_\",\"type\":\"address\"}],\"name\":\"setBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cattle_\",\"type\":\"address\"}],\"name\":\"setCow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost_\",\"type\":\"uint256\"}],\"name\":\"setEnergyCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"planet_\",\"type\":\"address\"}],\"name\":\"setPlanet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"setProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stable_\",\"type\":\"address\"}],\"name\":\"setStable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"BVT_\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stable\",\"outputs\":[{\"internalType\":\"contractIStable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"upLoad\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMatingTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userUploadList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BreedCenter is an auto generated Go binding around an Ethereum contract.
type BreedCenter struct {
	BreedCenterCaller     // Read-only binding to the contract
	BreedCenterTransactor // Write-only binding to the contract
	BreedCenterFilterer   // Log filterer for contract events
}

// BreedCenterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BreedCenterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BreedCenterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BreedCenterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BreedCenterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BreedCenterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BreedCenterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BreedCenterSession struct {
	Contract     *BreedCenter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BreedCenterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BreedCenterCallerSession struct {
	Contract *BreedCenterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BreedCenterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BreedCenterTransactorSession struct {
	Contract     *BreedCenterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BreedCenterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BreedCenterRaw struct {
	Contract *BreedCenter // Generic contract binding to access the raw methods on
}

// BreedCenterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BreedCenterCallerRaw struct {
	Contract *BreedCenterCaller // Generic read-only contract binding to access the raw methods on
}

// BreedCenterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BreedCenterTransactorRaw struct {
	Contract *BreedCenterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBreedCenter creates a new instance of BreedCenter, bound to a specific deployed contract.
func NewBreedCenter(address common.Address, backend bind.ContractBackend) (*BreedCenter, error) {
	contract, err := bindBreedCenter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BreedCenter{BreedCenterCaller: BreedCenterCaller{contract: contract}, BreedCenterTransactor: BreedCenterTransactor{contract: contract}, BreedCenterFilterer: BreedCenterFilterer{contract: contract}}, nil
}

// NewBreedCenterCaller creates a new read-only instance of BreedCenter, bound to a specific deployed contract.
func NewBreedCenterCaller(address common.Address, caller bind.ContractCaller) (*BreedCenterCaller, error) {
	contract, err := bindBreedCenter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BreedCenterCaller{contract: contract}, nil
}

// NewBreedCenterTransactor creates a new write-only instance of BreedCenter, bound to a specific deployed contract.
func NewBreedCenterTransactor(address common.Address, transactor bind.ContractTransactor) (*BreedCenterTransactor, error) {
	contract, err := bindBreedCenter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BreedCenterTransactor{contract: contract}, nil
}

// NewBreedCenterFilterer creates a new log filterer instance of BreedCenter, bound to a specific deployed contract.
func NewBreedCenterFilterer(address common.Address, filterer bind.ContractFilterer) (*BreedCenterFilterer, error) {
	contract, err := bindBreedCenter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BreedCenterFilterer{contract: contract}, nil
}

// bindBreedCenter binds a generic wrapper to an already deployed contract.
func bindBreedCenter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BreedCenterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BreedCenter *BreedCenterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BreedCenter.Contract.BreedCenterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BreedCenter *BreedCenterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BreedCenter.Contract.BreedCenterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BreedCenter *BreedCenterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BreedCenter.Contract.BreedCenterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BreedCenter *BreedCenterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BreedCenter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BreedCenter *BreedCenterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BreedCenter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BreedCenter *BreedCenterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BreedCenter.Contract.contract.Transact(opts, method, params...)
}

// BVG is a free data retrieval call binding the contract method 0x87464389.
//
// Solidity: function BVG() view returns(address)
func (_BreedCenter *BreedCenterCaller) BVG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "BVG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BVG is a free data retrieval call binding the contract method 0x87464389.
//
// Solidity: function BVG() view returns(address)
func (_BreedCenter *BreedCenterSession) BVG() (common.Address, error) {
	return _BreedCenter.Contract.BVG(&_BreedCenter.CallOpts)
}

// BVG is a free data retrieval call binding the contract method 0x87464389.
//
// Solidity: function BVG() view returns(address)
func (_BreedCenter *BreedCenterCallerSession) BVG() (common.Address, error) {
	return _BreedCenter.Contract.BVG(&_BreedCenter.CallOpts)
}

// BVT is a free data retrieval call binding the contract method 0x5f24b4a4.
//
// Solidity: function BVT() view returns(address)
func (_BreedCenter *BreedCenterCaller) BVT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "BVT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BVT is a free data retrieval call binding the contract method 0x5f24b4a4.
//
// Solidity: function BVT() view returns(address)
func (_BreedCenter *BreedCenterSession) BVT() (common.Address, error) {
	return _BreedCenter.Contract.BVT(&_BreedCenter.CallOpts)
}

// BVT is a free data retrieval call binding the contract method 0x5f24b4a4.
//
// Solidity: function BVT() view returns(address)
func (_BreedCenter *BreedCenterCallerSession) BVT() (common.Address, error) {
	return _BreedCenter.Contract.BVT(&_BreedCenter.CallOpts)
}

// Box is a free data retrieval call binding the contract method 0x754215a1.
//
// Solidity: function box() view returns(address)
func (_BreedCenter *BreedCenterCaller) Box(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "box")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Box is a free data retrieval call binding the contract method 0x754215a1.
//
// Solidity: function box() view returns(address)
func (_BreedCenter *BreedCenterSession) Box() (common.Address, error) {
	return _BreedCenter.Contract.Box(&_BreedCenter.CallOpts)
}

// Box is a free data retrieval call binding the contract method 0x754215a1.
//
// Solidity: function box() view returns(address)
func (_BreedCenter *BreedCenterCallerSession) Box() (common.Address, error) {
	return _BreedCenter.Contract.Box(&_BreedCenter.CallOpts)
}

// Cattle is a free data retrieval call binding the contract method 0x01c211c7.
//
// Solidity: function cattle() view returns(address)
func (_BreedCenter *BreedCenterCaller) Cattle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "cattle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cattle is a free data retrieval call binding the contract method 0x01c211c7.
//
// Solidity: function cattle() view returns(address)
func (_BreedCenter *BreedCenterSession) Cattle() (common.Address, error) {
	return _BreedCenter.Contract.Cattle(&_BreedCenter.CallOpts)
}

// Cattle is a free data retrieval call binding the contract method 0x01c211c7.
//
// Solidity: function cattle() view returns(address)
func (_BreedCenter *BreedCenterCallerSession) Cattle() (common.Address, error) {
	return _BreedCenter.Contract.Cattle(&_BreedCenter.CallOpts)
}

// CheckMatingTimeBatch is a free data retrieval call binding the contract method 0x761d6418.
//
// Solidity: function checkMatingTimeBatch(uint256[] list) view returns(uint256[] out)
func (_BreedCenter *BreedCenterCaller) CheckMatingTimeBatch(opts *bind.CallOpts, list []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "checkMatingTimeBatch", list)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CheckMatingTimeBatch is a free data retrieval call binding the contract method 0x761d6418.
//
// Solidity: function checkMatingTimeBatch(uint256[] list) view returns(uint256[] out)
func (_BreedCenter *BreedCenterSession) CheckMatingTimeBatch(list []*big.Int) ([]*big.Int, error) {
	return _BreedCenter.Contract.CheckMatingTimeBatch(&_BreedCenter.CallOpts, list)
}

// CheckMatingTimeBatch is a free data retrieval call binding the contract method 0x761d6418.
//
// Solidity: function checkMatingTimeBatch(uint256[] list) view returns(uint256[] out)
func (_BreedCenter *BreedCenterCallerSession) CheckMatingTimeBatch(list []*big.Int) ([]*big.Int, error) {
	return _BreedCenter.Contract.CheckMatingTimeBatch(&_BreedCenter.CallOpts, list)
}

// EnergyCost is a free data retrieval call binding the contract method 0x44af49c5.
//
// Solidity: function energyCost() view returns(uint256)
func (_BreedCenter *BreedCenterCaller) EnergyCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "energyCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnergyCost is a free data retrieval call binding the contract method 0x44af49c5.
//
// Solidity: function energyCost() view returns(uint256)
func (_BreedCenter *BreedCenterSession) EnergyCost() (*big.Int, error) {
	return _BreedCenter.Contract.EnergyCost(&_BreedCenter.CallOpts)
}

// EnergyCost is a free data retrieval call binding the contract method 0x44af49c5.
//
// Solidity: function energyCost() view returns(uint256)
func (_BreedCenter *BreedCenterCallerSession) EnergyCost() (*big.Int, error) {
	return _BreedCenter.Contract.EnergyCost(&_BreedCenter.CallOpts)
}

// GetUserUploadList is a free data retrieval call binding the contract method 0x72329e0d.
//
// Solidity: function getUserUploadList(address addr_) view returns(uint256[])
func (_BreedCenter *BreedCenterCaller) GetUserUploadList(opts *bind.CallOpts, addr_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "getUserUploadList", addr_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserUploadList is a free data retrieval call binding the contract method 0x72329e0d.
//
// Solidity: function getUserUploadList(address addr_) view returns(uint256[])
func (_BreedCenter *BreedCenterSession) GetUserUploadList(addr_ common.Address) ([]*big.Int, error) {
	return _BreedCenter.Contract.GetUserUploadList(&_BreedCenter.CallOpts, addr_)
}

// GetUserUploadList is a free data retrieval call binding the contract method 0x72329e0d.
//
// Solidity: function getUserUploadList(address addr_) view returns(uint256[])
func (_BreedCenter *BreedCenterCallerSession) GetUserUploadList(addr_ common.Address) ([]*big.Int, error) {
	return _BreedCenter.Contract.GetUserUploadList(&_BreedCenter.CallOpts, addr_)
}

// Index is a free data retrieval call binding the contract method 0x335932fc.
//
// Solidity: function index(uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterCaller) Index(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "index", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Index is a free data retrieval call binding the contract method 0x335932fc.
//
// Solidity: function index(uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterSession) Index(arg0 *big.Int) (*big.Int, error) {
	return _BreedCenter.Contract.Index(&_BreedCenter.CallOpts, arg0)
}

// Index is a free data retrieval call binding the contract method 0x335932fc.
//
// Solidity: function index(uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterCallerSession) Index(arg0 *big.Int) (*big.Int, error) {
	return _BreedCenter.Contract.Index(&_BreedCenter.CallOpts, arg0)
}

// LastMatingTime is a free data retrieval call binding the contract method 0xc4c1a187.
//
// Solidity: function lastMatingTime(uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterCaller) LastMatingTime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "lastMatingTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMatingTime is a free data retrieval call binding the contract method 0xc4c1a187.
//
// Solidity: function lastMatingTime(uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterSession) LastMatingTime(arg0 *big.Int) (*big.Int, error) {
	return _BreedCenter.Contract.LastMatingTime(&_BreedCenter.CallOpts, arg0)
}

// LastMatingTime is a free data retrieval call binding the contract method 0xc4c1a187.
//
// Solidity: function lastMatingTime(uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterCallerSession) LastMatingTime(arg0 *big.Int) (*big.Int, error) {
	return _BreedCenter.Contract.LastMatingTime(&_BreedCenter.CallOpts, arg0)
}

// MatingTime is a free data retrieval call binding the contract method 0x134f4282.
//
// Solidity: function matingTime(uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterCaller) MatingTime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "matingTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MatingTime is a free data retrieval call binding the contract method 0x134f4282.
//
// Solidity: function matingTime(uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterSession) MatingTime(arg0 *big.Int) (*big.Int, error) {
	return _BreedCenter.Contract.MatingTime(&_BreedCenter.CallOpts, arg0)
}

// MatingTime is a free data retrieval call binding the contract method 0x134f4282.
//
// Solidity: function matingTime(uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterCallerSession) MatingTime(arg0 *big.Int) (*big.Int, error) {
	return _BreedCenter.Contract.MatingTime(&_BreedCenter.CallOpts, arg0)
}

// OnSale is a free data retrieval call binding the contract method 0xb43a76f1.
//
// Solidity: function onSale(uint256 ) view returns(bool)
func (_BreedCenter *BreedCenterCaller) OnSale(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "onSale", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnSale is a free data retrieval call binding the contract method 0xb43a76f1.
//
// Solidity: function onSale(uint256 ) view returns(bool)
func (_BreedCenter *BreedCenterSession) OnSale(arg0 *big.Int) (bool, error) {
	return _BreedCenter.Contract.OnSale(&_BreedCenter.CallOpts, arg0)
}

// OnSale is a free data retrieval call binding the contract method 0xb43a76f1.
//
// Solidity: function onSale(uint256 ) view returns(bool)
func (_BreedCenter *BreedCenterCallerSession) OnSale(arg0 *big.Int) (bool, error) {
	return _BreedCenter.Contract.OnSale(&_BreedCenter.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BreedCenter *BreedCenterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BreedCenter *BreedCenterSession) Owner() (common.Address, error) {
	return _BreedCenter.Contract.Owner(&_BreedCenter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BreedCenter *BreedCenterCallerSession) Owner() (common.Address, error) {
	return _BreedCenter.Contract.Owner(&_BreedCenter.CallOpts)
}

// Photo is a free data retrieval call binding the contract method 0xd76ec1a3.
//
// Solidity: function photo() view returns(address)
func (_BreedCenter *BreedCenterCaller) Photo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "photo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Photo is a free data retrieval call binding the contract method 0xd76ec1a3.
//
// Solidity: function photo() view returns(address)
func (_BreedCenter *BreedCenterSession) Photo() (common.Address, error) {
	return _BreedCenter.Contract.Photo(&_BreedCenter.CallOpts)
}

// Photo is a free data retrieval call binding the contract method 0xd76ec1a3.
//
// Solidity: function photo() view returns(address)
func (_BreedCenter *BreedCenterCallerSession) Photo() (common.Address, error) {
	return _BreedCenter.Contract.Photo(&_BreedCenter.CallOpts)
}

// Planet is a free data retrieval call binding the contract method 0x8443bea6.
//
// Solidity: function planet() view returns(address)
func (_BreedCenter *BreedCenterCaller) Planet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "planet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Planet is a free data retrieval call binding the contract method 0x8443bea6.
//
// Solidity: function planet() view returns(address)
func (_BreedCenter *BreedCenterSession) Planet() (common.Address, error) {
	return _BreedCenter.Contract.Planet(&_BreedCenter.CallOpts)
}

// Planet is a free data retrieval call binding the contract method 0x8443bea6.
//
// Solidity: function planet() view returns(address)
func (_BreedCenter *BreedCenterCallerSession) Planet() (common.Address, error) {
	return _BreedCenter.Contract.Planet(&_BreedCenter.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0x26a49e37.
//
// Solidity: function price(uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterCaller) Price(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "price", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0x26a49e37.
//
// Solidity: function price(uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterSession) Price(arg0 *big.Int) (*big.Int, error) {
	return _BreedCenter.Contract.Price(&_BreedCenter.CallOpts, arg0)
}

// Price is a free data retrieval call binding the contract method 0x26a49e37.
//
// Solidity: function price(uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterCallerSession) Price(arg0 *big.Int) (*big.Int, error) {
	return _BreedCenter.Contract.Price(&_BreedCenter.CallOpts, arg0)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(address)
func (_BreedCenter *BreedCenterCaller) Stable(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "stable")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(address)
func (_BreedCenter *BreedCenterSession) Stable() (common.Address, error) {
	return _BreedCenter.Contract.Stable(&_BreedCenter.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(address)
func (_BreedCenter *BreedCenterCallerSession) Stable() (common.Address, error) {
	return _BreedCenter.Contract.Stable(&_BreedCenter.CallOpts)
}

// UserMatingTimes is a free data retrieval call binding the contract method 0x18d774a4.
//
// Solidity: function userMatingTimes(address ) view returns(uint256)
func (_BreedCenter *BreedCenterCaller) UserMatingTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "userMatingTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMatingTimes is a free data retrieval call binding the contract method 0x18d774a4.
//
// Solidity: function userMatingTimes(address ) view returns(uint256)
func (_BreedCenter *BreedCenterSession) UserMatingTimes(arg0 common.Address) (*big.Int, error) {
	return _BreedCenter.Contract.UserMatingTimes(&_BreedCenter.CallOpts, arg0)
}

// UserMatingTimes is a free data retrieval call binding the contract method 0x18d774a4.
//
// Solidity: function userMatingTimes(address ) view returns(uint256)
func (_BreedCenter *BreedCenterCallerSession) UserMatingTimes(arg0 common.Address) (*big.Int, error) {
	return _BreedCenter.Contract.UserMatingTimes(&_BreedCenter.CallOpts, arg0)
}

// UserUploadList is a free data retrieval call binding the contract method 0x63f42bde.
//
// Solidity: function userUploadList(address , uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterCaller) UserUploadList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BreedCenter.contract.Call(opts, &out, "userUploadList", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUploadList is a free data retrieval call binding the contract method 0x63f42bde.
//
// Solidity: function userUploadList(address , uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterSession) UserUploadList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BreedCenter.Contract.UserUploadList(&_BreedCenter.CallOpts, arg0, arg1)
}

// UserUploadList is a free data retrieval call binding the contract method 0x63f42bde.
//
// Solidity: function userUploadList(address , uint256 ) view returns(uint256)
func (_BreedCenter *BreedCenterCallerSession) UserUploadList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BreedCenter.Contract.UserUploadList(&_BreedCenter.CallOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_BreedCenter *BreedCenterTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_BreedCenter *BreedCenterSession) Initialize() (*types.Transaction, error) {
	return _BreedCenter.Contract.Initialize(&_BreedCenter.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_BreedCenter *BreedCenterTransactorSession) Initialize() (*types.Transaction, error) {
	return _BreedCenter.Contract.Initialize(&_BreedCenter.TransactOpts)
}

// Mating is a paid mutator transaction binding the contract method 0x09d8dd7d.
//
// Solidity: function mating(uint256 myTokenId, uint256 targetTokenID) returns()
func (_BreedCenter *BreedCenterTransactor) Mating(opts *bind.TransactOpts, myTokenId *big.Int, targetTokenID *big.Int) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "mating", myTokenId, targetTokenID)
}

// Mating is a paid mutator transaction binding the contract method 0x09d8dd7d.
//
// Solidity: function mating(uint256 myTokenId, uint256 targetTokenID) returns()
func (_BreedCenter *BreedCenterSession) Mating(myTokenId *big.Int, targetTokenID *big.Int) (*types.Transaction, error) {
	return _BreedCenter.Contract.Mating(&_BreedCenter.TransactOpts, myTokenId, targetTokenID)
}

// Mating is a paid mutator transaction binding the contract method 0x09d8dd7d.
//
// Solidity: function mating(uint256 myTokenId, uint256 targetTokenID) returns()
func (_BreedCenter *BreedCenterTransactorSession) Mating(myTokenId *big.Int, targetTokenID *big.Int) (*types.Transaction, error) {
	return _BreedCenter.Contract.Mating(&_BreedCenter.TransactOpts, myTokenId, targetTokenID)
}

// OffSale is a paid mutator transaction binding the contract method 0x97fd3a64.
//
// Solidity: function offSale(uint256 tokenId) returns()
func (_BreedCenter *BreedCenterTransactor) OffSale(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "offSale", tokenId)
}

// OffSale is a paid mutator transaction binding the contract method 0x97fd3a64.
//
// Solidity: function offSale(uint256 tokenId) returns()
func (_BreedCenter *BreedCenterSession) OffSale(tokenId *big.Int) (*types.Transaction, error) {
	return _BreedCenter.Contract.OffSale(&_BreedCenter.TransactOpts, tokenId)
}

// OffSale is a paid mutator transaction binding the contract method 0x97fd3a64.
//
// Solidity: function offSale(uint256 tokenId) returns()
func (_BreedCenter *BreedCenterTransactorSession) OffSale(tokenId *big.Int) (*types.Transaction, error) {
	return _BreedCenter.Contract.OffSale(&_BreedCenter.TransactOpts, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BreedCenter *BreedCenterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BreedCenter *BreedCenterSession) RenounceOwnership() (*types.Transaction, error) {
	return _BreedCenter.Contract.RenounceOwnership(&_BreedCenter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BreedCenter *BreedCenterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BreedCenter.Contract.RenounceOwnership(&_BreedCenter.TransactOpts)
}

// SelfMating is a paid mutator transaction binding the contract method 0x0c3f4a2d.
//
// Solidity: function selfMating(uint256 tokenId1, uint256 tokenId2) returns()
func (_BreedCenter *BreedCenterTransactor) SelfMating(opts *bind.TransactOpts, tokenId1 *big.Int, tokenId2 *big.Int) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "selfMating", tokenId1, tokenId2)
}

// SelfMating is a paid mutator transaction binding the contract method 0x0c3f4a2d.
//
// Solidity: function selfMating(uint256 tokenId1, uint256 tokenId2) returns()
func (_BreedCenter *BreedCenterSession) SelfMating(tokenId1 *big.Int, tokenId2 *big.Int) (*types.Transaction, error) {
	return _BreedCenter.Contract.SelfMating(&_BreedCenter.TransactOpts, tokenId1, tokenId2)
}

// SelfMating is a paid mutator transaction binding the contract method 0x0c3f4a2d.
//
// Solidity: function selfMating(uint256 tokenId1, uint256 tokenId2) returns()
func (_BreedCenter *BreedCenterTransactorSession) SelfMating(tokenId1 *big.Int, tokenId2 *big.Int) (*types.Transaction, error) {
	return _BreedCenter.Contract.SelfMating(&_BreedCenter.TransactOpts, tokenId1, tokenId2)
}

// SetBox is a paid mutator transaction binding the contract method 0x537627ab.
//
// Solidity: function setBox(address box_) returns()
func (_BreedCenter *BreedCenterTransactor) SetBox(opts *bind.TransactOpts, box_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "setBox", box_)
}

// SetBox is a paid mutator transaction binding the contract method 0x537627ab.
//
// Solidity: function setBox(address box_) returns()
func (_BreedCenter *BreedCenterSession) SetBox(box_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetBox(&_BreedCenter.TransactOpts, box_)
}

// SetBox is a paid mutator transaction binding the contract method 0x537627ab.
//
// Solidity: function setBox(address box_) returns()
func (_BreedCenter *BreedCenterTransactorSession) SetBox(box_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetBox(&_BreedCenter.TransactOpts, box_)
}

// SetCow is a paid mutator transaction binding the contract method 0x1ddbd167.
//
// Solidity: function setCow(address cattle_) returns()
func (_BreedCenter *BreedCenterTransactor) SetCow(opts *bind.TransactOpts, cattle_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "setCow", cattle_)
}

// SetCow is a paid mutator transaction binding the contract method 0x1ddbd167.
//
// Solidity: function setCow(address cattle_) returns()
func (_BreedCenter *BreedCenterSession) SetCow(cattle_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetCow(&_BreedCenter.TransactOpts, cattle_)
}

// SetCow is a paid mutator transaction binding the contract method 0x1ddbd167.
//
// Solidity: function setCow(address cattle_) returns()
func (_BreedCenter *BreedCenterTransactorSession) SetCow(cattle_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetCow(&_BreedCenter.TransactOpts, cattle_)
}

// SetEnergyCost is a paid mutator transaction binding the contract method 0xa91c39a9.
//
// Solidity: function setEnergyCost(uint256 cost_) returns()
func (_BreedCenter *BreedCenterTransactor) SetEnergyCost(opts *bind.TransactOpts, cost_ *big.Int) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "setEnergyCost", cost_)
}

// SetEnergyCost is a paid mutator transaction binding the contract method 0xa91c39a9.
//
// Solidity: function setEnergyCost(uint256 cost_) returns()
func (_BreedCenter *BreedCenterSession) SetEnergyCost(cost_ *big.Int) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetEnergyCost(&_BreedCenter.TransactOpts, cost_)
}

// SetEnergyCost is a paid mutator transaction binding the contract method 0xa91c39a9.
//
// Solidity: function setEnergyCost(uint256 cost_) returns()
func (_BreedCenter *BreedCenterTransactorSession) SetEnergyCost(cost_ *big.Int) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetEnergyCost(&_BreedCenter.TransactOpts, cost_)
}

// SetPlanet is a paid mutator transaction binding the contract method 0xad911387.
//
// Solidity: function setPlanet(address planet_) returns()
func (_BreedCenter *BreedCenterTransactor) SetPlanet(opts *bind.TransactOpts, planet_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "setPlanet", planet_)
}

// SetPlanet is a paid mutator transaction binding the contract method 0xad911387.
//
// Solidity: function setPlanet(address planet_) returns()
func (_BreedCenter *BreedCenterSession) SetPlanet(planet_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetPlanet(&_BreedCenter.TransactOpts, planet_)
}

// SetPlanet is a paid mutator transaction binding the contract method 0xad911387.
//
// Solidity: function setPlanet(address planet_) returns()
func (_BreedCenter *BreedCenterTransactorSession) SetPlanet(planet_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetPlanet(&_BreedCenter.TransactOpts, planet_)
}

// SetProfile is a paid mutator transaction binding the contract method 0xe0b6a1e3.
//
// Solidity: function setProfile(address addr_) returns()
func (_BreedCenter *BreedCenterTransactor) SetProfile(opts *bind.TransactOpts, addr_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "setProfile", addr_)
}

// SetProfile is a paid mutator transaction binding the contract method 0xe0b6a1e3.
//
// Solidity: function setProfile(address addr_) returns()
func (_BreedCenter *BreedCenterSession) SetProfile(addr_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetProfile(&_BreedCenter.TransactOpts, addr_)
}

// SetProfile is a paid mutator transaction binding the contract method 0xe0b6a1e3.
//
// Solidity: function setProfile(address addr_) returns()
func (_BreedCenter *BreedCenterTransactorSession) SetProfile(addr_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetProfile(&_BreedCenter.TransactOpts, addr_)
}

// SetStable is a paid mutator transaction binding the contract method 0xdba71376.
//
// Solidity: function setStable(address stable_) returns()
func (_BreedCenter *BreedCenterTransactor) SetStable(opts *bind.TransactOpts, stable_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "setStable", stable_)
}

// SetStable is a paid mutator transaction binding the contract method 0xdba71376.
//
// Solidity: function setStable(address stable_) returns()
func (_BreedCenter *BreedCenterSession) SetStable(stable_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetStable(&_BreedCenter.TransactOpts, stable_)
}

// SetStable is a paid mutator transaction binding the contract method 0xdba71376.
//
// Solidity: function setStable(address stable_) returns()
func (_BreedCenter *BreedCenterTransactorSession) SetStable(stable_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetStable(&_BreedCenter.TransactOpts, stable_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address BVT_) returns()
func (_BreedCenter *BreedCenterTransactor) SetToken(opts *bind.TransactOpts, BVT_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "setToken", BVT_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address BVT_) returns()
func (_BreedCenter *BreedCenterSession) SetToken(BVT_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetToken(&_BreedCenter.TransactOpts, BVT_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address BVT_) returns()
func (_BreedCenter *BreedCenterTransactorSession) SetToken(BVT_ common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.SetToken(&_BreedCenter.TransactOpts, BVT_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BreedCenter *BreedCenterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BreedCenter *BreedCenterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.TransferOwnership(&_BreedCenter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BreedCenter *BreedCenterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BreedCenter.Contract.TransferOwnership(&_BreedCenter.TransactOpts, newOwner)
}

// UpLoad is a paid mutator transaction binding the contract method 0x1e1068d5.
//
// Solidity: function upLoad(uint256 tokenId, uint256 price_) returns()
func (_BreedCenter *BreedCenterTransactor) UpLoad(opts *bind.TransactOpts, tokenId *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _BreedCenter.contract.Transact(opts, "upLoad", tokenId, price_)
}

// UpLoad is a paid mutator transaction binding the contract method 0x1e1068d5.
//
// Solidity: function upLoad(uint256 tokenId, uint256 price_) returns()
func (_BreedCenter *BreedCenterSession) UpLoad(tokenId *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _BreedCenter.Contract.UpLoad(&_BreedCenter.TransactOpts, tokenId, price_)
}

// UpLoad is a paid mutator transaction binding the contract method 0x1e1068d5.
//
// Solidity: function upLoad(uint256 tokenId, uint256 price_) returns()
func (_BreedCenter *BreedCenterTransactorSession) UpLoad(tokenId *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _BreedCenter.Contract.UpLoad(&_BreedCenter.TransactOpts, tokenId, price_)
}

// BreedCenterMateIterator is returned from FilterMate and is used to iterate over the raw logs and unpacked data for Mate events raised by the BreedCenter contract.
type BreedCenterMateIterator struct {
	Event *BreedCenterMate // Event containing the contract specifics and raw log

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
func (it *BreedCenterMateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BreedCenterMate)
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
		it.Event = new(BreedCenterMate)
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
func (it *BreedCenterMateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BreedCenterMateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BreedCenterMate represents a Mate event raised by the BreedCenter contract.
type BreedCenterMate struct {
	Player        common.Address
	TokenId       *big.Int
	TargetTokenID *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMate is a free log retrieval operation binding the contract event 0x5242693aa1335ee93aec1a37b90a06ab4ebee2b7f8ad1871b50719efd040a649.
//
// Solidity: event Mate(address indexed player_, uint256 indexed tokenId, uint256 indexed targetTokenID)
func (_BreedCenter *BreedCenterFilterer) FilterMate(opts *bind.FilterOpts, player_ []common.Address, tokenId []*big.Int, targetTokenID []*big.Int) (*BreedCenterMateIterator, error) {

	var player_Rule []interface{}
	for _, player_Item := range player_ {
		player_Rule = append(player_Rule, player_Item)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var targetTokenIDRule []interface{}
	for _, targetTokenIDItem := range targetTokenID {
		targetTokenIDRule = append(targetTokenIDRule, targetTokenIDItem)
	}

	logs, sub, err := _BreedCenter.contract.FilterLogs(opts, "Mate", player_Rule, tokenIdRule, targetTokenIDRule)
	if err != nil {
		return nil, err
	}
	return &BreedCenterMateIterator{contract: _BreedCenter.contract, event: "Mate", logs: logs, sub: sub}, nil
}

// WatchMate is a free log subscription operation binding the contract event 0x5242693aa1335ee93aec1a37b90a06ab4ebee2b7f8ad1871b50719efd040a649.
//
// Solidity: event Mate(address indexed player_, uint256 indexed tokenId, uint256 indexed targetTokenID)
func (_BreedCenter *BreedCenterFilterer) WatchMate(opts *bind.WatchOpts, sink chan<- *BreedCenterMate, player_ []common.Address, tokenId []*big.Int, targetTokenID []*big.Int) (event.Subscription, error) {

	var player_Rule []interface{}
	for _, player_Item := range player_ {
		player_Rule = append(player_Rule, player_Item)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var targetTokenIDRule []interface{}
	for _, targetTokenIDItem := range targetTokenID {
		targetTokenIDRule = append(targetTokenIDRule, targetTokenIDItem)
	}

	logs, sub, err := _BreedCenter.contract.WatchLogs(opts, "Mate", player_Rule, tokenIdRule, targetTokenIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BreedCenterMate)
				if err := _BreedCenter.contract.UnpackLog(event, "Mate", log); err != nil {
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

// ParseMate is a log parse operation binding the contract event 0x5242693aa1335ee93aec1a37b90a06ab4ebee2b7f8ad1871b50719efd040a649.
//
// Solidity: event Mate(address indexed player_, uint256 indexed tokenId, uint256 indexed targetTokenID)
func (_BreedCenter *BreedCenterFilterer) ParseMate(log types.Log) (*BreedCenterMate, error) {
	event := new(BreedCenterMate)
	if err := _BreedCenter.contract.UnpackLog(event, "Mate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BreedCenterOffSaleIterator is returned from FilterOffSale and is used to iterate over the raw logs and unpacked data for OffSale events raised by the BreedCenter contract.
type BreedCenterOffSaleIterator struct {
	Event *BreedCenterOffSale // Event containing the contract specifics and raw log

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
func (it *BreedCenterOffSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BreedCenterOffSale)
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
		it.Event = new(BreedCenterOffSale)
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
func (it *BreedCenterOffSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BreedCenterOffSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BreedCenterOffSale represents a OffSale event raised by the BreedCenter contract.
type BreedCenterOffSale struct {
	Sender  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOffSale is a free log retrieval operation binding the contract event 0x6a70ef5e3ae6ca2a87f56abe5fb27baf70a2dfdd7a3884810c61fa7d76b4585c.
//
// Solidity: event OffSale(address indexed sender_, uint256 indexed tokenId)
func (_BreedCenter *BreedCenterFilterer) FilterOffSale(opts *bind.FilterOpts, sender_ []common.Address, tokenId []*big.Int) (*BreedCenterOffSaleIterator, error) {

	var sender_Rule []interface{}
	for _, sender_Item := range sender_ {
		sender_Rule = append(sender_Rule, sender_Item)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BreedCenter.contract.FilterLogs(opts, "OffSale", sender_Rule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BreedCenterOffSaleIterator{contract: _BreedCenter.contract, event: "OffSale", logs: logs, sub: sub}, nil
}

// WatchOffSale is a free log subscription operation binding the contract event 0x6a70ef5e3ae6ca2a87f56abe5fb27baf70a2dfdd7a3884810c61fa7d76b4585c.
//
// Solidity: event OffSale(address indexed sender_, uint256 indexed tokenId)
func (_BreedCenter *BreedCenterFilterer) WatchOffSale(opts *bind.WatchOpts, sink chan<- *BreedCenterOffSale, sender_ []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sender_Rule []interface{}
	for _, sender_Item := range sender_ {
		sender_Rule = append(sender_Rule, sender_Item)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BreedCenter.contract.WatchLogs(opts, "OffSale", sender_Rule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BreedCenterOffSale)
				if err := _BreedCenter.contract.UnpackLog(event, "OffSale", log); err != nil {
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

// ParseOffSale is a log parse operation binding the contract event 0x6a70ef5e3ae6ca2a87f56abe5fb27baf70a2dfdd7a3884810c61fa7d76b4585c.
//
// Solidity: event OffSale(address indexed sender_, uint256 indexed tokenId)
func (_BreedCenter *BreedCenterFilterer) ParseOffSale(log types.Log) (*BreedCenterOffSale, error) {
	event := new(BreedCenterOffSale)
	if err := _BreedCenter.contract.UnpackLog(event, "OffSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BreedCenterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BreedCenter contract.
type BreedCenterOwnershipTransferredIterator struct {
	Event *BreedCenterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BreedCenterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BreedCenterOwnershipTransferred)
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
		it.Event = new(BreedCenterOwnershipTransferred)
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
func (it *BreedCenterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BreedCenterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BreedCenterOwnershipTransferred represents a OwnershipTransferred event raised by the BreedCenter contract.
type BreedCenterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BreedCenter *BreedCenterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BreedCenterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BreedCenter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BreedCenterOwnershipTransferredIterator{contract: _BreedCenter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BreedCenter *BreedCenterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BreedCenterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BreedCenter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BreedCenterOwnershipTransferred)
				if err := _BreedCenter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BreedCenter *BreedCenterFilterer) ParseOwnershipTransferred(log types.Log) (*BreedCenterOwnershipTransferred, error) {
	event := new(BreedCenterOwnershipTransferred)
	if err := _BreedCenter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BreedCenterUpLoadIterator is returned from FilterUpLoad and is used to iterate over the raw logs and unpacked data for UpLoad events raised by the BreedCenter contract.
type BreedCenterUpLoadIterator struct {
	Event *BreedCenterUpLoad // Event containing the contract specifics and raw log

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
func (it *BreedCenterUpLoadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BreedCenterUpLoad)
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
		it.Event = new(BreedCenterUpLoad)
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
func (it *BreedCenterUpLoadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BreedCenterUpLoadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BreedCenterUpLoad represents a UpLoad event raised by the BreedCenter contract.
type BreedCenterUpLoad struct {
	Sender  common.Address
	Price   *big.Int
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpLoad is a free log retrieval operation binding the contract event 0x7f4d84239409ccb115d1ad979c17ca067a25d979e719f6fedab335a45595f8a2.
//
// Solidity: event UpLoad(address indexed sender_, uint256 indexed price, uint256 indexed tokenId)
func (_BreedCenter *BreedCenterFilterer) FilterUpLoad(opts *bind.FilterOpts, sender_ []common.Address, price []*big.Int, tokenId []*big.Int) (*BreedCenterUpLoadIterator, error) {

	var sender_Rule []interface{}
	for _, sender_Item := range sender_ {
		sender_Rule = append(sender_Rule, sender_Item)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BreedCenter.contract.FilterLogs(opts, "UpLoad", sender_Rule, priceRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BreedCenterUpLoadIterator{contract: _BreedCenter.contract, event: "UpLoad", logs: logs, sub: sub}, nil
}

// WatchUpLoad is a free log subscription operation binding the contract event 0x7f4d84239409ccb115d1ad979c17ca067a25d979e719f6fedab335a45595f8a2.
//
// Solidity: event UpLoad(address indexed sender_, uint256 indexed price, uint256 indexed tokenId)
func (_BreedCenter *BreedCenterFilterer) WatchUpLoad(opts *bind.WatchOpts, sink chan<- *BreedCenterUpLoad, sender_ []common.Address, price []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

	var sender_Rule []interface{}
	for _, sender_Item := range sender_ {
		sender_Rule = append(sender_Rule, sender_Item)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BreedCenter.contract.WatchLogs(opts, "UpLoad", sender_Rule, priceRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BreedCenterUpLoad)
				if err := _BreedCenter.contract.UnpackLog(event, "UpLoad", log); err != nil {
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

// ParseUpLoad is a log parse operation binding the contract event 0x7f4d84239409ccb115d1ad979c17ca067a25d979e719f6fedab335a45595f8a2.
//
// Solidity: event UpLoad(address indexed sender_, uint256 indexed price, uint256 indexed tokenId)
func (_BreedCenter *BreedCenterFilterer) ParseUpLoad(log types.Log) (*BreedCenterUpLoad, error) {
	event := new(BreedCenterUpLoad)
	if err := _BreedCenter.contract.UnpackLog(event, "UpLoad", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
