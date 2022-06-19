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

// QueryAPIABI is the input ABI used to generate the binding from.
const QueryAPIABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"avatar\",\"outputs\":[{\"internalType\":\"contractIProfilePhoto\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"battleInfo\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"info\",\"type\":\"uint256[3]\"},{\"internalType\":\"bool\",\"name\":\"isDead\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"bullPenInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cattle\",\"outputs\":[{\"internalType\":\"contractICOW\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compound\",\"outputs\":[{\"internalType\":\"contractICompound\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"targetId\",\"type\":\"uint256[]\"}],\"name\":\"compoundInfo\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"info\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cowInfoes\",\"outputs\":[{\"internalType\":\"uint256[23]\",\"name\":\"info1\",\"type\":\"uint256[23]\"},{\"internalType\":\"bool[3]\",\"name\":\"info2\",\"type\":\"bool[3]\"},{\"internalType\":\"uint256[2]\",\"name\":\"parents\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenId\",\"type\":\"uint256[]\"}],\"name\":\"getMattingTimeBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"getUserProfilePhoto\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"item\",\"outputs\":[{\"internalType\":\"contractICattle1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mail\",\"outputs\":[{\"internalType\":\"contractIMail\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"market\",\"outputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mating\",\"outputs\":[{\"internalType\":\"contractIMating\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"milk\",\"outputs\":[{\"internalType\":\"contractIMilk\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"planet\",\"outputs\":[{\"internalType\":\"contractIPlanet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"setCattle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCompound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setMail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setMating\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setMilk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"setPlanet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setProfilePhoto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"setStable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setTec\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stable\",\"outputs\":[{\"internalType\":\"contractIStable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tec\",\"outputs\":[{\"internalType\":\"contractITec\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"userCenter\",\"outputs\":[{\"internalType\":\"uint256[10]\",\"name\":\"info\",\"type\":\"uint256[10]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// QueryAPI is an auto generated Go binding around an Ethereum contract.
type QueryAPI struct {
	QueryAPICaller     // Read-only binding to the contract
	QueryAPITransactor // Write-only binding to the contract
	QueryAPIFilterer   // Log filterer for contract events
}

// QueryAPICaller is an auto generated read-only Go binding around an Ethereum contract.
type QueryAPICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueryAPITransactor is an auto generated write-only Go binding around an Ethereum contract.
type QueryAPITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueryAPIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QueryAPIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueryAPISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QueryAPISession struct {
	Contract     *QueryAPI         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QueryAPICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QueryAPICallerSession struct {
	Contract *QueryAPICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// QueryAPITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QueryAPITransactorSession struct {
	Contract     *QueryAPITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QueryAPIRaw is an auto generated low-level Go binding around an Ethereum contract.
type QueryAPIRaw struct {
	Contract *QueryAPI // Generic contract binding to access the raw methods on
}

// QueryAPICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QueryAPICallerRaw struct {
	Contract *QueryAPICaller // Generic read-only contract binding to access the raw methods on
}

// QueryAPITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QueryAPITransactorRaw struct {
	Contract *QueryAPITransactor // Generic write-only contract binding to access the raw methods on
}

// NewQueryAPI creates a new instance of QueryAPI, bound to a specific deployed contract.
func NewQueryAPI(address common.Address, backend bind.ContractBackend) (*QueryAPI, error) {
	contract, err := bindQueryAPI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QueryAPI{QueryAPICaller: QueryAPICaller{contract: contract}, QueryAPITransactor: QueryAPITransactor{contract: contract}, QueryAPIFilterer: QueryAPIFilterer{contract: contract}}, nil
}

// NewQueryAPICaller creates a new read-only instance of QueryAPI, bound to a specific deployed contract.
func NewQueryAPICaller(address common.Address, caller bind.ContractCaller) (*QueryAPICaller, error) {
	contract, err := bindQueryAPI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QueryAPICaller{contract: contract}, nil
}

// NewQueryAPITransactor creates a new write-only instance of QueryAPI, bound to a specific deployed contract.
func NewQueryAPITransactor(address common.Address, transactor bind.ContractTransactor) (*QueryAPITransactor, error) {
	contract, err := bindQueryAPI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QueryAPITransactor{contract: contract}, nil
}

// NewQueryAPIFilterer creates a new log filterer instance of QueryAPI, bound to a specific deployed contract.
func NewQueryAPIFilterer(address common.Address, filterer bind.ContractFilterer) (*QueryAPIFilterer, error) {
	contract, err := bindQueryAPI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QueryAPIFilterer{contract: contract}, nil
}

// bindQueryAPI binds a generic wrapper to an already deployed contract.
func bindQueryAPI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QueryAPIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueryAPI *QueryAPIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QueryAPI.Contract.QueryAPICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueryAPI *QueryAPIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueryAPI.Contract.QueryAPITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueryAPI *QueryAPIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueryAPI.Contract.QueryAPITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueryAPI *QueryAPICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QueryAPI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueryAPI *QueryAPITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueryAPI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueryAPI *QueryAPITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueryAPI.Contract.contract.Transact(opts, method, params...)
}

// Avatar is a free data retrieval call binding the contract method 0x5aef7de6.
//
// Solidity: function avatar() view returns(address)
func (_QueryAPI *QueryAPICaller) Avatar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "avatar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avatar is a free data retrieval call binding the contract method 0x5aef7de6.
//
// Solidity: function avatar() view returns(address)
func (_QueryAPI *QueryAPISession) Avatar() (common.Address, error) {
	return _QueryAPI.Contract.Avatar(&_QueryAPI.CallOpts)
}

// Avatar is a free data retrieval call binding the contract method 0x5aef7de6.
//
// Solidity: function avatar() view returns(address)
func (_QueryAPI *QueryAPICallerSession) Avatar() (common.Address, error) {
	return _QueryAPI.Contract.Avatar(&_QueryAPI.CallOpts)
}

// BattleInfo is a free data retrieval call binding the contract method 0xdb71dcdc.
//
// Solidity: function battleInfo(uint256 tokenId) view returns(uint256[3] info, bool isDead, address owner_)
func (_QueryAPI *QueryAPICaller) BattleInfo(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Info   [3]*big.Int
	IsDead bool
	Owner  common.Address
}, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "battleInfo", tokenId)

	outstruct := new(struct {
		Info   [3]*big.Int
		IsDead bool
		Owner  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Info = *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)
	outstruct.IsDead = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// BattleInfo is a free data retrieval call binding the contract method 0xdb71dcdc.
//
// Solidity: function battleInfo(uint256 tokenId) view returns(uint256[3] info, bool isDead, address owner_)
func (_QueryAPI *QueryAPISession) BattleInfo(tokenId *big.Int) (struct {
	Info   [3]*big.Int
	IsDead bool
	Owner  common.Address
}, error) {
	return _QueryAPI.Contract.BattleInfo(&_QueryAPI.CallOpts, tokenId)
}

// BattleInfo is a free data retrieval call binding the contract method 0xdb71dcdc.
//
// Solidity: function battleInfo(uint256 tokenId) view returns(uint256[3] info, bool isDead, address owner_)
func (_QueryAPI *QueryAPICallerSession) BattleInfo(tokenId *big.Int) (struct {
	Info   [3]*big.Int
	IsDead bool
	Owner  common.Address
}, error) {
	return _QueryAPI.Contract.BattleInfo(&_QueryAPI.CallOpts, tokenId)
}

// BullPenInfo is a free data retrieval call binding the contract method 0x3144286b.
//
// Solidity: function bullPenInfo(address addr_) view returns(uint256, uint256, uint256, uint256[])
func (_QueryAPI *QueryAPICaller) BullPenInfo(opts *bind.CallOpts, addr_ common.Address) (*big.Int, *big.Int, *big.Int, []*big.Int, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "bullPenInfo", addr_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, out3, err

}

// BullPenInfo is a free data retrieval call binding the contract method 0x3144286b.
//
// Solidity: function bullPenInfo(address addr_) view returns(uint256, uint256, uint256, uint256[])
func (_QueryAPI *QueryAPISession) BullPenInfo(addr_ common.Address) (*big.Int, *big.Int, *big.Int, []*big.Int, error) {
	return _QueryAPI.Contract.BullPenInfo(&_QueryAPI.CallOpts, addr_)
}

// BullPenInfo is a free data retrieval call binding the contract method 0x3144286b.
//
// Solidity: function bullPenInfo(address addr_) view returns(uint256, uint256, uint256, uint256[])
func (_QueryAPI *QueryAPICallerSession) BullPenInfo(addr_ common.Address) (*big.Int, *big.Int, *big.Int, []*big.Int, error) {
	return _QueryAPI.Contract.BullPenInfo(&_QueryAPI.CallOpts, addr_)
}

// Cattle is a free data retrieval call binding the contract method 0x01c211c7.
//
// Solidity: function cattle() view returns(address)
func (_QueryAPI *QueryAPICaller) Cattle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "cattle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cattle is a free data retrieval call binding the contract method 0x01c211c7.
//
// Solidity: function cattle() view returns(address)
func (_QueryAPI *QueryAPISession) Cattle() (common.Address, error) {
	return _QueryAPI.Contract.Cattle(&_QueryAPI.CallOpts)
}

// Cattle is a free data retrieval call binding the contract method 0x01c211c7.
//
// Solidity: function cattle() view returns(address)
func (_QueryAPI *QueryAPICallerSession) Cattle() (common.Address, error) {
	return _QueryAPI.Contract.Cattle(&_QueryAPI.CallOpts)
}

// Compound is a free data retrieval call binding the contract method 0xf69e2046.
//
// Solidity: function compound() view returns(address)
func (_QueryAPI *QueryAPICaller) Compound(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "compound")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Compound is a free data retrieval call binding the contract method 0xf69e2046.
//
// Solidity: function compound() view returns(address)
func (_QueryAPI *QueryAPISession) Compound() (common.Address, error) {
	return _QueryAPI.Contract.Compound(&_QueryAPI.CallOpts)
}

// Compound is a free data retrieval call binding the contract method 0xf69e2046.
//
// Solidity: function compound() view returns(address)
func (_QueryAPI *QueryAPICallerSession) Compound() (common.Address, error) {
	return _QueryAPI.Contract.Compound(&_QueryAPI.CallOpts)
}

// CompoundInfo is a free data retrieval call binding the contract method 0x2e8e7d7a.
//
// Solidity: function compoundInfo(uint256 tokenId, uint256[] targetId) view returns(uint256[5] info)
func (_QueryAPI *QueryAPICaller) CompoundInfo(opts *bind.CallOpts, tokenId *big.Int, targetId []*big.Int) ([5]*big.Int, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "compoundInfo", tokenId, targetId)

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// CompoundInfo is a free data retrieval call binding the contract method 0x2e8e7d7a.
//
// Solidity: function compoundInfo(uint256 tokenId, uint256[] targetId) view returns(uint256[5] info)
func (_QueryAPI *QueryAPISession) CompoundInfo(tokenId *big.Int, targetId []*big.Int) ([5]*big.Int, error) {
	return _QueryAPI.Contract.CompoundInfo(&_QueryAPI.CallOpts, tokenId, targetId)
}

// CompoundInfo is a free data retrieval call binding the contract method 0x2e8e7d7a.
//
// Solidity: function compoundInfo(uint256 tokenId, uint256[] targetId) view returns(uint256[5] info)
func (_QueryAPI *QueryAPICallerSession) CompoundInfo(tokenId *big.Int, targetId []*big.Int) ([5]*big.Int, error) {
	return _QueryAPI.Contract.CompoundInfo(&_QueryAPI.CallOpts, tokenId, targetId)
}

// CowInfoes is a free data retrieval call binding the contract method 0xf610e982.
//
// Solidity: function cowInfoes(uint256 tokenId) view returns(uint256[23] info1, bool[3] info2, uint256[2] parents)
func (_QueryAPI *QueryAPICaller) CowInfoes(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Info1   [23]*big.Int
	Info2   [3]bool
	Parents [2]*big.Int
}, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "cowInfoes", tokenId)

	outstruct := new(struct {
		Info1   [23]*big.Int
		Info2   [3]bool
		Parents [2]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Info1 = *abi.ConvertType(out[0], new([23]*big.Int)).(*[23]*big.Int)
	outstruct.Info2 = *abi.ConvertType(out[1], new([3]bool)).(*[3]bool)
	outstruct.Parents = *abi.ConvertType(out[2], new([2]*big.Int)).(*[2]*big.Int)

	return *outstruct, err

}

// CowInfoes is a free data retrieval call binding the contract method 0xf610e982.
//
// Solidity: function cowInfoes(uint256 tokenId) view returns(uint256[23] info1, bool[3] info2, uint256[2] parents)
func (_QueryAPI *QueryAPISession) CowInfoes(tokenId *big.Int) (struct {
	Info1   [23]*big.Int
	Info2   [3]bool
	Parents [2]*big.Int
}, error) {
	return _QueryAPI.Contract.CowInfoes(&_QueryAPI.CallOpts, tokenId)
}

// CowInfoes is a free data retrieval call binding the contract method 0xf610e982.
//
// Solidity: function cowInfoes(uint256 tokenId) view returns(uint256[23] info1, bool[3] info2, uint256[2] parents)
func (_QueryAPI *QueryAPICallerSession) CowInfoes(tokenId *big.Int) (struct {
	Info1   [23]*big.Int
	Info2   [3]bool
	Parents [2]*big.Int
}, error) {
	return _QueryAPI.Contract.CowInfoes(&_QueryAPI.CallOpts, tokenId)
}

// GetMattingTimeBatch is a free data retrieval call binding the contract method 0x3b6b12f1.
//
// Solidity: function getMattingTimeBatch(uint256[] tokenId) view returns(uint256[])
func (_QueryAPI *QueryAPICaller) GetMattingTimeBatch(opts *bind.CallOpts, tokenId []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "getMattingTimeBatch", tokenId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetMattingTimeBatch is a free data retrieval call binding the contract method 0x3b6b12f1.
//
// Solidity: function getMattingTimeBatch(uint256[] tokenId) view returns(uint256[])
func (_QueryAPI *QueryAPISession) GetMattingTimeBatch(tokenId []*big.Int) ([]*big.Int, error) {
	return _QueryAPI.Contract.GetMattingTimeBatch(&_QueryAPI.CallOpts, tokenId)
}

// GetMattingTimeBatch is a free data retrieval call binding the contract method 0x3b6b12f1.
//
// Solidity: function getMattingTimeBatch(uint256[] tokenId) view returns(uint256[])
func (_QueryAPI *QueryAPICallerSession) GetMattingTimeBatch(tokenId []*big.Int) ([]*big.Int, error) {
	return _QueryAPI.Contract.GetMattingTimeBatch(&_QueryAPI.CallOpts, tokenId)
}

// GetUserProfilePhoto is a free data retrieval call binding the contract method 0x1d97b39e.
//
// Solidity: function getUserProfilePhoto(address addr_) view returns(string[])
func (_QueryAPI *QueryAPICaller) GetUserProfilePhoto(opts *bind.CallOpts, addr_ common.Address) ([]string, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "getUserProfilePhoto", addr_)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetUserProfilePhoto is a free data retrieval call binding the contract method 0x1d97b39e.
//
// Solidity: function getUserProfilePhoto(address addr_) view returns(string[])
func (_QueryAPI *QueryAPISession) GetUserProfilePhoto(addr_ common.Address) ([]string, error) {
	return _QueryAPI.Contract.GetUserProfilePhoto(&_QueryAPI.CallOpts, addr_)
}

// GetUserProfilePhoto is a free data retrieval call binding the contract method 0x1d97b39e.
//
// Solidity: function getUserProfilePhoto(address addr_) view returns(string[])
func (_QueryAPI *QueryAPICallerSession) GetUserProfilePhoto(addr_ common.Address) ([]string, error) {
	return _QueryAPI.Contract.GetUserProfilePhoto(&_QueryAPI.CallOpts, addr_)
}

// Item is a free data retrieval call binding the contract method 0xf2a4a82e.
//
// Solidity: function item() view returns(address)
func (_QueryAPI *QueryAPICaller) Item(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "item")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Item is a free data retrieval call binding the contract method 0xf2a4a82e.
//
// Solidity: function item() view returns(address)
func (_QueryAPI *QueryAPISession) Item() (common.Address, error) {
	return _QueryAPI.Contract.Item(&_QueryAPI.CallOpts)
}

// Item is a free data retrieval call binding the contract method 0xf2a4a82e.
//
// Solidity: function item() view returns(address)
func (_QueryAPI *QueryAPICallerSession) Item() (common.Address, error) {
	return _QueryAPI.Contract.Item(&_QueryAPI.CallOpts)
}

// Mail is a free data retrieval call binding the contract method 0x9a68308f.
//
// Solidity: function mail() view returns(address)
func (_QueryAPI *QueryAPICaller) Mail(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "mail")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mail is a free data retrieval call binding the contract method 0x9a68308f.
//
// Solidity: function mail() view returns(address)
func (_QueryAPI *QueryAPISession) Mail() (common.Address, error) {
	return _QueryAPI.Contract.Mail(&_QueryAPI.CallOpts)
}

// Mail is a free data retrieval call binding the contract method 0x9a68308f.
//
// Solidity: function mail() view returns(address)
func (_QueryAPI *QueryAPICallerSession) Mail() (common.Address, error) {
	return _QueryAPI.Contract.Mail(&_QueryAPI.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_QueryAPI *QueryAPICaller) Market(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "market")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_QueryAPI *QueryAPISession) Market() (common.Address, error) {
	return _QueryAPI.Contract.Market(&_QueryAPI.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_QueryAPI *QueryAPICallerSession) Market() (common.Address, error) {
	return _QueryAPI.Contract.Market(&_QueryAPI.CallOpts)
}

// Mating is a free data retrieval call binding the contract method 0xb4dac333.
//
// Solidity: function mating() view returns(address)
func (_QueryAPI *QueryAPICaller) Mating(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "mating")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mating is a free data retrieval call binding the contract method 0xb4dac333.
//
// Solidity: function mating() view returns(address)
func (_QueryAPI *QueryAPISession) Mating() (common.Address, error) {
	return _QueryAPI.Contract.Mating(&_QueryAPI.CallOpts)
}

// Mating is a free data retrieval call binding the contract method 0xb4dac333.
//
// Solidity: function mating() view returns(address)
func (_QueryAPI *QueryAPICallerSession) Mating() (common.Address, error) {
	return _QueryAPI.Contract.Mating(&_QueryAPI.CallOpts)
}

// Milk is a free data retrieval call binding the contract method 0xdbf8b168.
//
// Solidity: function milk() view returns(address)
func (_QueryAPI *QueryAPICaller) Milk(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "milk")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Milk is a free data retrieval call binding the contract method 0xdbf8b168.
//
// Solidity: function milk() view returns(address)
func (_QueryAPI *QueryAPISession) Milk() (common.Address, error) {
	return _QueryAPI.Contract.Milk(&_QueryAPI.CallOpts)
}

// Milk is a free data retrieval call binding the contract method 0xdbf8b168.
//
// Solidity: function milk() view returns(address)
func (_QueryAPI *QueryAPICallerSession) Milk() (common.Address, error) {
	return _QueryAPI.Contract.Milk(&_QueryAPI.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QueryAPI *QueryAPICaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QueryAPI *QueryAPISession) Owner() (common.Address, error) {
	return _QueryAPI.Contract.Owner(&_QueryAPI.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QueryAPI *QueryAPICallerSession) Owner() (common.Address, error) {
	return _QueryAPI.Contract.Owner(&_QueryAPI.CallOpts)
}

// Planet is a free data retrieval call binding the contract method 0x8443bea6.
//
// Solidity: function planet() view returns(address)
func (_QueryAPI *QueryAPICaller) Planet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "planet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Planet is a free data retrieval call binding the contract method 0x8443bea6.
//
// Solidity: function planet() view returns(address)
func (_QueryAPI *QueryAPISession) Planet() (common.Address, error) {
	return _QueryAPI.Contract.Planet(&_QueryAPI.CallOpts)
}

// Planet is a free data retrieval call binding the contract method 0x8443bea6.
//
// Solidity: function planet() view returns(address)
func (_QueryAPI *QueryAPICallerSession) Planet() (common.Address, error) {
	return _QueryAPI.Contract.Planet(&_QueryAPI.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(address)
func (_QueryAPI *QueryAPICaller) Stable(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "stable")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(address)
func (_QueryAPI *QueryAPISession) Stable() (common.Address, error) {
	return _QueryAPI.Contract.Stable(&_QueryAPI.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(address)
func (_QueryAPI *QueryAPICallerSession) Stable() (common.Address, error) {
	return _QueryAPI.Contract.Stable(&_QueryAPI.CallOpts)
}

// Tec is a free data retrieval call binding the contract method 0x0660cf4c.
//
// Solidity: function tec() view returns(address)
func (_QueryAPI *QueryAPICaller) Tec(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "tec")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tec is a free data retrieval call binding the contract method 0x0660cf4c.
//
// Solidity: function tec() view returns(address)
func (_QueryAPI *QueryAPISession) Tec() (common.Address, error) {
	return _QueryAPI.Contract.Tec(&_QueryAPI.CallOpts)
}

// Tec is a free data retrieval call binding the contract method 0x0660cf4c.
//
// Solidity: function tec() view returns(address)
func (_QueryAPI *QueryAPICallerSession) Tec() (common.Address, error) {
	return _QueryAPI.Contract.Tec(&_QueryAPI.CallOpts)
}

// UserCenter is a free data retrieval call binding the contract method 0xf36eb671.
//
// Solidity: function userCenter(address player) view returns(uint256[10] info)
func (_QueryAPI *QueryAPICaller) UserCenter(opts *bind.CallOpts, player common.Address) ([10]*big.Int, error) {
	var out []interface{}
	err := _QueryAPI.contract.Call(opts, &out, "userCenter", player)

	if err != nil {
		return *new([10]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([10]*big.Int)).(*[10]*big.Int)

	return out0, err

}

// UserCenter is a free data retrieval call binding the contract method 0xf36eb671.
//
// Solidity: function userCenter(address player) view returns(uint256[10] info)
func (_QueryAPI *QueryAPISession) UserCenter(player common.Address) ([10]*big.Int, error) {
	return _QueryAPI.Contract.UserCenter(&_QueryAPI.CallOpts, player)
}

// UserCenter is a free data retrieval call binding the contract method 0xf36eb671.
//
// Solidity: function userCenter(address player) view returns(uint256[10] info)
func (_QueryAPI *QueryAPICallerSession) UserCenter(player common.Address) ([10]*big.Int, error) {
	return _QueryAPI.Contract.UserCenter(&_QueryAPI.CallOpts, player)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_QueryAPI *QueryAPITransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_QueryAPI *QueryAPISession) Initialize() (*types.Transaction, error) {
	return _QueryAPI.Contract.Initialize(&_QueryAPI.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_QueryAPI *QueryAPITransactorSession) Initialize() (*types.Transaction, error) {
	return _QueryAPI.Contract.Initialize(&_QueryAPI.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QueryAPI *QueryAPITransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QueryAPI *QueryAPISession) RenounceOwnership() (*types.Transaction, error) {
	return _QueryAPI.Contract.RenounceOwnership(&_QueryAPI.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QueryAPI *QueryAPITransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QueryAPI.Contract.RenounceOwnership(&_QueryAPI.TransactOpts)
}

// SetCattle is a paid mutator transaction binding the contract method 0x0de307fe.
//
// Solidity: function setCattle(address addr_) returns()
func (_QueryAPI *QueryAPITransactor) SetCattle(opts *bind.TransactOpts, addr_ common.Address) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "setCattle", addr_)
}

// SetCattle is a paid mutator transaction binding the contract method 0x0de307fe.
//
// Solidity: function setCattle(address addr_) returns()
func (_QueryAPI *QueryAPISession) SetCattle(addr_ common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetCattle(&_QueryAPI.TransactOpts, addr_)
}

// SetCattle is a paid mutator transaction binding the contract method 0x0de307fe.
//
// Solidity: function setCattle(address addr_) returns()
func (_QueryAPI *QueryAPITransactorSession) SetCattle(addr_ common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetCattle(&_QueryAPI.TransactOpts, addr_)
}

// SetCompound is a paid mutator transaction binding the contract method 0x2be682d4.
//
// Solidity: function setCompound(address addr) returns()
func (_QueryAPI *QueryAPITransactor) SetCompound(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "setCompound", addr)
}

// SetCompound is a paid mutator transaction binding the contract method 0x2be682d4.
//
// Solidity: function setCompound(address addr) returns()
func (_QueryAPI *QueryAPISession) SetCompound(addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetCompound(&_QueryAPI.TransactOpts, addr)
}

// SetCompound is a paid mutator transaction binding the contract method 0x2be682d4.
//
// Solidity: function setCompound(address addr) returns()
func (_QueryAPI *QueryAPITransactorSession) SetCompound(addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetCompound(&_QueryAPI.TransactOpts, addr)
}

// SetItem is a paid mutator transaction binding the contract method 0x165ed276.
//
// Solidity: function setItem(address addr_) returns()
func (_QueryAPI *QueryAPITransactor) SetItem(opts *bind.TransactOpts, addr_ common.Address) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "setItem", addr_)
}

// SetItem is a paid mutator transaction binding the contract method 0x165ed276.
//
// Solidity: function setItem(address addr_) returns()
func (_QueryAPI *QueryAPISession) SetItem(addr_ common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetItem(&_QueryAPI.TransactOpts, addr_)
}

// SetItem is a paid mutator transaction binding the contract method 0x165ed276.
//
// Solidity: function setItem(address addr_) returns()
func (_QueryAPI *QueryAPITransactorSession) SetItem(addr_ common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetItem(&_QueryAPI.TransactOpts, addr_)
}

// SetMail is a paid mutator transaction binding the contract method 0x773b1543.
//
// Solidity: function setMail(address addr) returns()
func (_QueryAPI *QueryAPITransactor) SetMail(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "setMail", addr)
}

// SetMail is a paid mutator transaction binding the contract method 0x773b1543.
//
// Solidity: function setMail(address addr) returns()
func (_QueryAPI *QueryAPISession) SetMail(addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetMail(&_QueryAPI.TransactOpts, addr)
}

// SetMail is a paid mutator transaction binding the contract method 0x773b1543.
//
// Solidity: function setMail(address addr) returns()
func (_QueryAPI *QueryAPITransactorSession) SetMail(addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetMail(&_QueryAPI.TransactOpts, addr)
}

// SetMating is a paid mutator transaction binding the contract method 0x9fd0e138.
//
// Solidity: function setMating(address addr) returns()
func (_QueryAPI *QueryAPITransactor) SetMating(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "setMating", addr)
}

// SetMating is a paid mutator transaction binding the contract method 0x9fd0e138.
//
// Solidity: function setMating(address addr) returns()
func (_QueryAPI *QueryAPISession) SetMating(addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetMating(&_QueryAPI.TransactOpts, addr)
}

// SetMating is a paid mutator transaction binding the contract method 0x9fd0e138.
//
// Solidity: function setMating(address addr) returns()
func (_QueryAPI *QueryAPITransactorSession) SetMating(addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetMating(&_QueryAPI.TransactOpts, addr)
}

// SetMilk is a paid mutator transaction binding the contract method 0x110a2ded.
//
// Solidity: function setMilk(address addr) returns()
func (_QueryAPI *QueryAPITransactor) SetMilk(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "setMilk", addr)
}

// SetMilk is a paid mutator transaction binding the contract method 0x110a2ded.
//
// Solidity: function setMilk(address addr) returns()
func (_QueryAPI *QueryAPISession) SetMilk(addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetMilk(&_QueryAPI.TransactOpts, addr)
}

// SetMilk is a paid mutator transaction binding the contract method 0x110a2ded.
//
// Solidity: function setMilk(address addr) returns()
func (_QueryAPI *QueryAPITransactorSession) SetMilk(addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetMilk(&_QueryAPI.TransactOpts, addr)
}

// SetPlanet is a paid mutator transaction binding the contract method 0xad911387.
//
// Solidity: function setPlanet(address addr_) returns()
func (_QueryAPI *QueryAPITransactor) SetPlanet(opts *bind.TransactOpts, addr_ common.Address) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "setPlanet", addr_)
}

// SetPlanet is a paid mutator transaction binding the contract method 0xad911387.
//
// Solidity: function setPlanet(address addr_) returns()
func (_QueryAPI *QueryAPISession) SetPlanet(addr_ common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetPlanet(&_QueryAPI.TransactOpts, addr_)
}

// SetPlanet is a paid mutator transaction binding the contract method 0xad911387.
//
// Solidity: function setPlanet(address addr_) returns()
func (_QueryAPI *QueryAPITransactorSession) SetPlanet(addr_ common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetPlanet(&_QueryAPI.TransactOpts, addr_)
}

// SetProfilePhoto is a paid mutator transaction binding the contract method 0x61a91290.
//
// Solidity: function setProfilePhoto(address addr) returns()
func (_QueryAPI *QueryAPITransactor) SetProfilePhoto(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "setProfilePhoto", addr)
}

// SetProfilePhoto is a paid mutator transaction binding the contract method 0x61a91290.
//
// Solidity: function setProfilePhoto(address addr) returns()
func (_QueryAPI *QueryAPISession) SetProfilePhoto(addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetProfilePhoto(&_QueryAPI.TransactOpts, addr)
}

// SetProfilePhoto is a paid mutator transaction binding the contract method 0x61a91290.
//
// Solidity: function setProfilePhoto(address addr) returns()
func (_QueryAPI *QueryAPITransactorSession) SetProfilePhoto(addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetProfilePhoto(&_QueryAPI.TransactOpts, addr)
}

// SetStable is a paid mutator transaction binding the contract method 0xdba71376.
//
// Solidity: function setStable(address addr_) returns()
func (_QueryAPI *QueryAPITransactor) SetStable(opts *bind.TransactOpts, addr_ common.Address) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "setStable", addr_)
}

// SetStable is a paid mutator transaction binding the contract method 0xdba71376.
//
// Solidity: function setStable(address addr_) returns()
func (_QueryAPI *QueryAPISession) SetStable(addr_ common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetStable(&_QueryAPI.TransactOpts, addr_)
}

// SetStable is a paid mutator transaction binding the contract method 0xdba71376.
//
// Solidity: function setStable(address addr_) returns()
func (_QueryAPI *QueryAPITransactorSession) SetStable(addr_ common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetStable(&_QueryAPI.TransactOpts, addr_)
}

// SetTec is a paid mutator transaction binding the contract method 0x72ad7cc1.
//
// Solidity: function setTec(address addr) returns()
func (_QueryAPI *QueryAPITransactor) SetTec(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "setTec", addr)
}

// SetTec is a paid mutator transaction binding the contract method 0x72ad7cc1.
//
// Solidity: function setTec(address addr) returns()
func (_QueryAPI *QueryAPISession) SetTec(addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetTec(&_QueryAPI.TransactOpts, addr)
}

// SetTec is a paid mutator transaction binding the contract method 0x72ad7cc1.
//
// Solidity: function setTec(address addr) returns()
func (_QueryAPI *QueryAPITransactorSession) SetTec(addr common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.SetTec(&_QueryAPI.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QueryAPI *QueryAPITransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QueryAPI.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QueryAPI *QueryAPISession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.TransferOwnership(&_QueryAPI.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QueryAPI *QueryAPITransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QueryAPI.Contract.TransferOwnership(&_QueryAPI.TransactOpts, newOwner)
}

// QueryAPIInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the QueryAPI contract.
type QueryAPIInitializedIterator struct {
	Event *QueryAPIInitialized // Event containing the contract specifics and raw log

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
func (it *QueryAPIInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QueryAPIInitialized)
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
		it.Event = new(QueryAPIInitialized)
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
func (it *QueryAPIInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QueryAPIInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QueryAPIInitialized represents a Initialized event raised by the QueryAPI contract.
type QueryAPIInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_QueryAPI *QueryAPIFilterer) FilterInitialized(opts *bind.FilterOpts) (*QueryAPIInitializedIterator, error) {

	logs, sub, err := _QueryAPI.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &QueryAPIInitializedIterator{contract: _QueryAPI.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_QueryAPI *QueryAPIFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *QueryAPIInitialized) (event.Subscription, error) {

	logs, sub, err := _QueryAPI.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QueryAPIInitialized)
				if err := _QueryAPI.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_QueryAPI *QueryAPIFilterer) ParseInitialized(log types.Log) (*QueryAPIInitialized, error) {
	event := new(QueryAPIInitialized)
	if err := _QueryAPI.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QueryAPIOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QueryAPI contract.
type QueryAPIOwnershipTransferredIterator struct {
	Event *QueryAPIOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QueryAPIOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QueryAPIOwnershipTransferred)
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
		it.Event = new(QueryAPIOwnershipTransferred)
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
func (it *QueryAPIOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QueryAPIOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QueryAPIOwnershipTransferred represents a OwnershipTransferred event raised by the QueryAPI contract.
type QueryAPIOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QueryAPI *QueryAPIFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QueryAPIOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QueryAPI.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QueryAPIOwnershipTransferredIterator{contract: _QueryAPI.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QueryAPI *QueryAPIFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QueryAPIOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QueryAPI.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QueryAPIOwnershipTransferred)
				if err := _QueryAPI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QueryAPI *QueryAPIFilterer) ParseOwnershipTransferred(log types.Log) (*QueryAPIOwnershipTransferred, error) {
	event := new(QueryAPIOwnershipTransferred)
	if err := _QueryAPI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
