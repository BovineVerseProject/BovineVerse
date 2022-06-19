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

// Marketplace721ABI is the input ABI used to generate the binding from.
const Marketplace721ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goodsId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goodsType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tradeType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"Exchanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goodsType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"place\",\"type\":\"address\"}],\"name\":\"NewGoods\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"goodsId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goodsType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tradeType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"UpdateGoodsStatus\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"goodsType_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"cancelSell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee_\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bank\",\"type\":\"address\"}],\"name\":\"changeTradeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"payee_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tradeType_\",\"type\":\"uint256\"}],\"name\":\"divestFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"goodsType_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"erc20Purchase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGoodsSeq\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"goodsType_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"getUserSaleList\",\"outputs\":[{\"internalType\":\"uint256[4][]\",\"name\":\"data\",\"type\":\"uint256[4][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"goodsInfos\",\"outputs\":[{\"internalType\":\"contractERC721Upgradeable\",\"name\":\"place\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"goodsType_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"mainCoinPurchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"types_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"place_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"newGoodsInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bank_\",\"type\":\"address\"}],\"name\":\"newTradeType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"types_\",\"type\":\"uint256\"}],\"name\":\"queryGoodsSoldData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"goodsType_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tradeType_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sellingGoods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tradeType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tradeType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Marketplace721 is an auto generated Go binding around an Ethereum contract.
type Marketplace721 struct {
	Marketplace721Caller     // Read-only binding to the contract
	Marketplace721Transactor // Write-only binding to the contract
	Marketplace721Filterer   // Log filterer for contract events
}

// Marketplace721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Marketplace721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Marketplace721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Marketplace721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Marketplace721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Marketplace721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Marketplace721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Marketplace721Session struct {
	Contract     *Marketplace721   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Marketplace721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Marketplace721CallerSession struct {
	Contract *Marketplace721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Marketplace721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Marketplace721TransactorSession struct {
	Contract     *Marketplace721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Marketplace721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Marketplace721Raw struct {
	Contract *Marketplace721 // Generic contract binding to access the raw methods on
}

// Marketplace721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Marketplace721CallerRaw struct {
	Contract *Marketplace721Caller // Generic read-only contract binding to access the raw methods on
}

// Marketplace721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Marketplace721TransactorRaw struct {
	Contract *Marketplace721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace721 creates a new instance of Marketplace721, bound to a specific deployed contract.
func NewMarketplace721(address common.Address, backend bind.ContractBackend) (*Marketplace721, error) {
	contract, err := bindMarketplace721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace721{Marketplace721Caller: Marketplace721Caller{contract: contract}, Marketplace721Transactor: Marketplace721Transactor{contract: contract}, Marketplace721Filterer: Marketplace721Filterer{contract: contract}}, nil
}

// NewMarketplace721Caller creates a new read-only instance of Marketplace721, bound to a specific deployed contract.
func NewMarketplace721Caller(address common.Address, caller bind.ContractCaller) (*Marketplace721Caller, error) {
	contract, err := bindMarketplace721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Marketplace721Caller{contract: contract}, nil
}

// NewMarketplace721Transactor creates a new write-only instance of Marketplace721, bound to a specific deployed contract.
func NewMarketplace721Transactor(address common.Address, transactor bind.ContractTransactor) (*Marketplace721Transactor, error) {
	contract, err := bindMarketplace721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Marketplace721Transactor{contract: contract}, nil
}

// NewMarketplace721Filterer creates a new log filterer instance of Marketplace721, bound to a specific deployed contract.
func NewMarketplace721Filterer(address common.Address, filterer bind.ContractFilterer) (*Marketplace721Filterer, error) {
	contract, err := bindMarketplace721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Marketplace721Filterer{contract: contract}, nil
}

// bindMarketplace721 binds a generic wrapper to an already deployed contract.
func bindMarketplace721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Marketplace721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace721 *Marketplace721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace721.Contract.Marketplace721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace721 *Marketplace721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace721.Contract.Marketplace721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace721 *Marketplace721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace721.Contract.Marketplace721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace721 *Marketplace721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace721 *Marketplace721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace721 *Marketplace721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace721.Contract.contract.Transact(opts, method, params...)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Marketplace721 *Marketplace721Caller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace721.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Marketplace721 *Marketplace721Session) Fee() (*big.Int, error) {
	return _Marketplace721.Contract.Fee(&_Marketplace721.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Marketplace721 *Marketplace721CallerSession) Fee() (*big.Int, error) {
	return _Marketplace721.Contract.Fee(&_Marketplace721.CallOpts)
}

// GetGoodsSeq is a free data retrieval call binding the contract method 0xa2155179.
//
// Solidity: function getGoodsSeq() view returns(uint256)
func (_Marketplace721 *Marketplace721Caller) GetGoodsSeq(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace721.contract.Call(opts, &out, "getGoodsSeq")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGoodsSeq is a free data retrieval call binding the contract method 0xa2155179.
//
// Solidity: function getGoodsSeq() view returns(uint256)
func (_Marketplace721 *Marketplace721Session) GetGoodsSeq() (*big.Int, error) {
	return _Marketplace721.Contract.GetGoodsSeq(&_Marketplace721.CallOpts)
}

// GetGoodsSeq is a free data retrieval call binding the contract method 0xa2155179.
//
// Solidity: function getGoodsSeq() view returns(uint256)
func (_Marketplace721 *Marketplace721CallerSession) GetGoodsSeq() (*big.Int, error) {
	return _Marketplace721.Contract.GetGoodsSeq(&_Marketplace721.CallOpts)
}

// GetUserSaleList is a free data retrieval call binding the contract method 0x225658ee.
//
// Solidity: function getUserSaleList(uint256 goodsType_, address addr_) view returns(uint256[4][] data)
func (_Marketplace721 *Marketplace721Caller) GetUserSaleList(opts *bind.CallOpts, goodsType_ *big.Int, addr_ common.Address) ([][4]*big.Int, error) {
	var out []interface{}
	err := _Marketplace721.contract.Call(opts, &out, "getUserSaleList", goodsType_, addr_)

	if err != nil {
		return *new([][4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]*big.Int)).(*[][4]*big.Int)

	return out0, err

}

// GetUserSaleList is a free data retrieval call binding the contract method 0x225658ee.
//
// Solidity: function getUserSaleList(uint256 goodsType_, address addr_) view returns(uint256[4][] data)
func (_Marketplace721 *Marketplace721Session) GetUserSaleList(goodsType_ *big.Int, addr_ common.Address) ([][4]*big.Int, error) {
	return _Marketplace721.Contract.GetUserSaleList(&_Marketplace721.CallOpts, goodsType_, addr_)
}

// GetUserSaleList is a free data retrieval call binding the contract method 0x225658ee.
//
// Solidity: function getUserSaleList(uint256 goodsType_, address addr_) view returns(uint256[4][] data)
func (_Marketplace721 *Marketplace721CallerSession) GetUserSaleList(goodsType_ *big.Int, addr_ common.Address) ([][4]*big.Int, error) {
	return _Marketplace721.Contract.GetUserSaleList(&_Marketplace721.CallOpts, goodsType_, addr_)
}

// GoodsInfos is a free data retrieval call binding the contract method 0x3e4cfbc0.
//
// Solidity: function goodsInfos(uint256 ) view returns(address place, string name)
func (_Marketplace721 *Marketplace721Caller) GoodsInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Place common.Address
	Name  string
}, error) {
	var out []interface{}
	err := _Marketplace721.contract.Call(opts, &out, "goodsInfos", arg0)

	outstruct := new(struct {
		Place common.Address
		Name  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Place = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// GoodsInfos is a free data retrieval call binding the contract method 0x3e4cfbc0.
//
// Solidity: function goodsInfos(uint256 ) view returns(address place, string name)
func (_Marketplace721 *Marketplace721Session) GoodsInfos(arg0 *big.Int) (struct {
	Place common.Address
	Name  string
}, error) {
	return _Marketplace721.Contract.GoodsInfos(&_Marketplace721.CallOpts, arg0)
}

// GoodsInfos is a free data retrieval call binding the contract method 0x3e4cfbc0.
//
// Solidity: function goodsInfos(uint256 ) view returns(address place, string name)
func (_Marketplace721 *Marketplace721CallerSession) GoodsInfos(arg0 *big.Int) (struct {
	Place common.Address
	Name  string
}, error) {
	return _Marketplace721.Contract.GoodsInfos(&_Marketplace721.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace721 *Marketplace721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace721 *Marketplace721Session) Owner() (common.Address, error) {
	return _Marketplace721.Contract.Owner(&_Marketplace721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace721 *Marketplace721CallerSession) Owner() (common.Address, error) {
	return _Marketplace721.Contract.Owner(&_Marketplace721.CallOpts)
}

// QueryGoodsSoldData is a free data retrieval call binding the contract method 0x3fcdc391.
//
// Solidity: function queryGoodsSoldData(uint256 types_) view returns(uint256)
func (_Marketplace721 *Marketplace721Caller) QueryGoodsSoldData(opts *bind.CallOpts, types_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace721.contract.Call(opts, &out, "queryGoodsSoldData", types_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryGoodsSoldData is a free data retrieval call binding the contract method 0x3fcdc391.
//
// Solidity: function queryGoodsSoldData(uint256 types_) view returns(uint256)
func (_Marketplace721 *Marketplace721Session) QueryGoodsSoldData(types_ *big.Int) (*big.Int, error) {
	return _Marketplace721.Contract.QueryGoodsSoldData(&_Marketplace721.CallOpts, types_)
}

// QueryGoodsSoldData is a free data retrieval call binding the contract method 0x3fcdc391.
//
// Solidity: function queryGoodsSoldData(uint256 types_) view returns(uint256)
func (_Marketplace721 *Marketplace721CallerSession) QueryGoodsSoldData(types_ *big.Int) (*big.Int, error) {
	return _Marketplace721.Contract.QueryGoodsSoldData(&_Marketplace721.CallOpts, types_)
}

// SellingGoods is a free data retrieval call binding the contract method 0x9e12eae8.
//
// Solidity: function sellingGoods(uint256 , uint256 ) view returns(uint256 id, uint256 tradeType, uint256 price, address owner)
func (_Marketplace721 *Marketplace721Caller) SellingGoods(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Id        *big.Int
	TradeType *big.Int
	Price     *big.Int
	Owner     common.Address
}, error) {
	var out []interface{}
	err := _Marketplace721.contract.Call(opts, &out, "sellingGoods", arg0, arg1)

	outstruct := new(struct {
		Id        *big.Int
		TradeType *big.Int
		Price     *big.Int
		Owner     common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TradeType = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SellingGoods is a free data retrieval call binding the contract method 0x9e12eae8.
//
// Solidity: function sellingGoods(uint256 , uint256 ) view returns(uint256 id, uint256 tradeType, uint256 price, address owner)
func (_Marketplace721 *Marketplace721Session) SellingGoods(arg0 *big.Int, arg1 *big.Int) (struct {
	Id        *big.Int
	TradeType *big.Int
	Price     *big.Int
	Owner     common.Address
}, error) {
	return _Marketplace721.Contract.SellingGoods(&_Marketplace721.CallOpts, arg0, arg1)
}

// SellingGoods is a free data retrieval call binding the contract method 0x9e12eae8.
//
// Solidity: function sellingGoods(uint256 , uint256 ) view returns(uint256 id, uint256 tradeType, uint256 price, address owner)
func (_Marketplace721 *Marketplace721CallerSession) SellingGoods(arg0 *big.Int, arg1 *big.Int) (struct {
	Id        *big.Int
	TradeType *big.Int
	Price     *big.Int
	Owner     common.Address
}, error) {
	return _Marketplace721.Contract.SellingGoods(&_Marketplace721.CallOpts, arg0, arg1)
}

// TradeType is a free data retrieval call binding the contract method 0x506fac59.
//
// Solidity: function tradeType(uint256 ) view returns(address)
func (_Marketplace721 *Marketplace721Caller) TradeType(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Marketplace721.contract.Call(opts, &out, "tradeType", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TradeType is a free data retrieval call binding the contract method 0x506fac59.
//
// Solidity: function tradeType(uint256 ) view returns(address)
func (_Marketplace721 *Marketplace721Session) TradeType(arg0 *big.Int) (common.Address, error) {
	return _Marketplace721.Contract.TradeType(&_Marketplace721.CallOpts, arg0)
}

// TradeType is a free data retrieval call binding the contract method 0x506fac59.
//
// Solidity: function tradeType(uint256 ) view returns(address)
func (_Marketplace721 *Marketplace721CallerSession) TradeType(arg0 *big.Int) (common.Address, error) {
	return _Marketplace721.Contract.TradeType(&_Marketplace721.CallOpts, arg0)
}

// CancelSell is a paid mutator transaction binding the contract method 0xfc72226b.
//
// Solidity: function cancelSell(uint256 goodsType_, uint256 index_) returns()
func (_Marketplace721 *Marketplace721Transactor) CancelSell(opts *bind.TransactOpts, goodsType_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "cancelSell", goodsType_, index_)
}

// CancelSell is a paid mutator transaction binding the contract method 0xfc72226b.
//
// Solidity: function cancelSell(uint256 goodsType_, uint256 index_) returns()
func (_Marketplace721 *Marketplace721Session) CancelSell(goodsType_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.Contract.CancelSell(&_Marketplace721.TransactOpts, goodsType_, index_)
}

// CancelSell is a paid mutator transaction binding the contract method 0xfc72226b.
//
// Solidity: function cancelSell(uint256 goodsType_, uint256 index_) returns()
func (_Marketplace721 *Marketplace721TransactorSession) CancelSell(goodsType_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.Contract.CancelSell(&_Marketplace721.TransactOpts, goodsType_, index_)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 fee_) returns(bool)
func (_Marketplace721 *Marketplace721Transactor) ChangeFee(opts *bind.TransactOpts, fee_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "changeFee", fee_)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 fee_) returns(bool)
func (_Marketplace721 *Marketplace721Session) ChangeFee(fee_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.Contract.ChangeFee(&_Marketplace721.TransactOpts, fee_)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 fee_) returns(bool)
func (_Marketplace721 *Marketplace721TransactorSession) ChangeFee(fee_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.Contract.ChangeFee(&_Marketplace721.TransactOpts, fee_)
}

// ChangeTradeType is a paid mutator transaction binding the contract method 0xead7a57e.
//
// Solidity: function changeTradeType(uint256 id_, address bank) returns()
func (_Marketplace721 *Marketplace721Transactor) ChangeTradeType(opts *bind.TransactOpts, id_ *big.Int, bank common.Address) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "changeTradeType", id_, bank)
}

// ChangeTradeType is a paid mutator transaction binding the contract method 0xead7a57e.
//
// Solidity: function changeTradeType(uint256 id_, address bank) returns()
func (_Marketplace721 *Marketplace721Session) ChangeTradeType(id_ *big.Int, bank common.Address) (*types.Transaction, error) {
	return _Marketplace721.Contract.ChangeTradeType(&_Marketplace721.TransactOpts, id_, bank)
}

// ChangeTradeType is a paid mutator transaction binding the contract method 0xead7a57e.
//
// Solidity: function changeTradeType(uint256 id_, address bank) returns()
func (_Marketplace721 *Marketplace721TransactorSession) ChangeTradeType(id_ *big.Int, bank common.Address) (*types.Transaction, error) {
	return _Marketplace721.Contract.ChangeTradeType(&_Marketplace721.TransactOpts, id_, bank)
}

// DivestFee is a paid mutator transaction binding the contract method 0x97b4075e.
//
// Solidity: function divestFee(address payee_, uint256 value_, uint256 tradeType_) returns(bool)
func (_Marketplace721 *Marketplace721Transactor) DivestFee(opts *bind.TransactOpts, payee_ common.Address, value_ *big.Int, tradeType_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "divestFee", payee_, value_, tradeType_)
}

// DivestFee is a paid mutator transaction binding the contract method 0x97b4075e.
//
// Solidity: function divestFee(address payee_, uint256 value_, uint256 tradeType_) returns(bool)
func (_Marketplace721 *Marketplace721Session) DivestFee(payee_ common.Address, value_ *big.Int, tradeType_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.Contract.DivestFee(&_Marketplace721.TransactOpts, payee_, value_, tradeType_)
}

// DivestFee is a paid mutator transaction binding the contract method 0x97b4075e.
//
// Solidity: function divestFee(address payee_, uint256 value_, uint256 tradeType_) returns(bool)
func (_Marketplace721 *Marketplace721TransactorSession) DivestFee(payee_ common.Address, value_ *big.Int, tradeType_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.Contract.DivestFee(&_Marketplace721.TransactOpts, payee_, value_, tradeType_)
}

// Erc20Purchase is a paid mutator transaction binding the contract method 0x1c7c8bc8.
//
// Solidity: function erc20Purchase(uint256 goodsType_, uint256 tokenId_) returns()
func (_Marketplace721 *Marketplace721Transactor) Erc20Purchase(opts *bind.TransactOpts, goodsType_ *big.Int, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "erc20Purchase", goodsType_, tokenId_)
}

// Erc20Purchase is a paid mutator transaction binding the contract method 0x1c7c8bc8.
//
// Solidity: function erc20Purchase(uint256 goodsType_, uint256 tokenId_) returns()
func (_Marketplace721 *Marketplace721Session) Erc20Purchase(goodsType_ *big.Int, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.Contract.Erc20Purchase(&_Marketplace721.TransactOpts, goodsType_, tokenId_)
}

// Erc20Purchase is a paid mutator transaction binding the contract method 0x1c7c8bc8.
//
// Solidity: function erc20Purchase(uint256 goodsType_, uint256 tokenId_) returns()
func (_Marketplace721 *Marketplace721TransactorSession) Erc20Purchase(goodsType_ *big.Int, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.Contract.Erc20Purchase(&_Marketplace721.TransactOpts, goodsType_, tokenId_)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Marketplace721 *Marketplace721Transactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Marketplace721 *Marketplace721Session) Init() (*types.Transaction, error) {
	return _Marketplace721.Contract.Init(&_Marketplace721.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Marketplace721 *Marketplace721TransactorSession) Init() (*types.Transaction, error) {
	return _Marketplace721.Contract.Init(&_Marketplace721.TransactOpts)
}

// MainCoinPurchase is a paid mutator transaction binding the contract method 0x0b3707e6.
//
// Solidity: function mainCoinPurchase(uint256 goodsType_, uint256 tokenId_) payable returns()
func (_Marketplace721 *Marketplace721Transactor) MainCoinPurchase(opts *bind.TransactOpts, goodsType_ *big.Int, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "mainCoinPurchase", goodsType_, tokenId_)
}

// MainCoinPurchase is a paid mutator transaction binding the contract method 0x0b3707e6.
//
// Solidity: function mainCoinPurchase(uint256 goodsType_, uint256 tokenId_) payable returns()
func (_Marketplace721 *Marketplace721Session) MainCoinPurchase(goodsType_ *big.Int, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.Contract.MainCoinPurchase(&_Marketplace721.TransactOpts, goodsType_, tokenId_)
}

// MainCoinPurchase is a paid mutator transaction binding the contract method 0x0b3707e6.
//
// Solidity: function mainCoinPurchase(uint256 goodsType_, uint256 tokenId_) payable returns()
func (_Marketplace721 *Marketplace721TransactorSession) MainCoinPurchase(goodsType_ *big.Int, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.Contract.MainCoinPurchase(&_Marketplace721.TransactOpts, goodsType_, tokenId_)
}

// NewGoodsInfo is a paid mutator transaction binding the contract method 0x674e0338.
//
// Solidity: function newGoodsInfo(uint256 types_, address place_, string name_) returns(bool)
func (_Marketplace721 *Marketplace721Transactor) NewGoodsInfo(opts *bind.TransactOpts, types_ *big.Int, place_ common.Address, name_ string) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "newGoodsInfo", types_, place_, name_)
}

// NewGoodsInfo is a paid mutator transaction binding the contract method 0x674e0338.
//
// Solidity: function newGoodsInfo(uint256 types_, address place_, string name_) returns(bool)
func (_Marketplace721 *Marketplace721Session) NewGoodsInfo(types_ *big.Int, place_ common.Address, name_ string) (*types.Transaction, error) {
	return _Marketplace721.Contract.NewGoodsInfo(&_Marketplace721.TransactOpts, types_, place_, name_)
}

// NewGoodsInfo is a paid mutator transaction binding the contract method 0x674e0338.
//
// Solidity: function newGoodsInfo(uint256 types_, address place_, string name_) returns(bool)
func (_Marketplace721 *Marketplace721TransactorSession) NewGoodsInfo(types_ *big.Int, place_ common.Address, name_ string) (*types.Transaction, error) {
	return _Marketplace721.Contract.NewGoodsInfo(&_Marketplace721.TransactOpts, types_, place_, name_)
}

// NewTradeType is a paid mutator transaction binding the contract method 0x91b5a02e.
//
// Solidity: function newTradeType(uint256 id_, address bank_) returns(bool)
func (_Marketplace721 *Marketplace721Transactor) NewTradeType(opts *bind.TransactOpts, id_ *big.Int, bank_ common.Address) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "newTradeType", id_, bank_)
}

// NewTradeType is a paid mutator transaction binding the contract method 0x91b5a02e.
//
// Solidity: function newTradeType(uint256 id_, address bank_) returns(bool)
func (_Marketplace721 *Marketplace721Session) NewTradeType(id_ *big.Int, bank_ common.Address) (*types.Transaction, error) {
	return _Marketplace721.Contract.NewTradeType(&_Marketplace721.TransactOpts, id_, bank_)
}

// NewTradeType is a paid mutator transaction binding the contract method 0x91b5a02e.
//
// Solidity: function newTradeType(uint256 id_, address bank_) returns(bool)
func (_Marketplace721 *Marketplace721TransactorSession) NewTradeType(id_ *big.Int, bank_ common.Address) (*types.Transaction, error) {
	return _Marketplace721.Contract.NewTradeType(&_Marketplace721.TransactOpts, id_, bank_)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Marketplace721 *Marketplace721Transactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Marketplace721 *Marketplace721Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Marketplace721.Contract.OnERC721Received(&_Marketplace721.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Marketplace721 *Marketplace721TransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Marketplace721.Contract.OnERC721Received(&_Marketplace721.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace721 *Marketplace721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace721 *Marketplace721Session) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace721.Contract.RenounceOwnership(&_Marketplace721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace721 *Marketplace721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace721.Contract.RenounceOwnership(&_Marketplace721.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0x3620875e.
//
// Solidity: function sell(uint256 goodsType_, uint256 tokenId_, uint256 tradeType_, uint256 price_) returns()
func (_Marketplace721 *Marketplace721Transactor) Sell(opts *bind.TransactOpts, goodsType_ *big.Int, tokenId_ *big.Int, tradeType_ *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "sell", goodsType_, tokenId_, tradeType_, price_)
}

// Sell is a paid mutator transaction binding the contract method 0x3620875e.
//
// Solidity: function sell(uint256 goodsType_, uint256 tokenId_, uint256 tradeType_, uint256 price_) returns()
func (_Marketplace721 *Marketplace721Session) Sell(goodsType_ *big.Int, tokenId_ *big.Int, tradeType_ *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.Contract.Sell(&_Marketplace721.TransactOpts, goodsType_, tokenId_, tradeType_, price_)
}

// Sell is a paid mutator transaction binding the contract method 0x3620875e.
//
// Solidity: function sell(uint256 goodsType_, uint256 tokenId_, uint256 tradeType_, uint256 price_) returns()
func (_Marketplace721 *Marketplace721TransactorSession) Sell(goodsType_ *big.Int, tokenId_ *big.Int, tradeType_ *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _Marketplace721.Contract.Sell(&_Marketplace721.TransactOpts, goodsType_, tokenId_, tradeType_, price_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace721 *Marketplace721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace721 *Marketplace721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace721.Contract.TransferOwnership(&_Marketplace721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace721 *Marketplace721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace721.Contract.TransferOwnership(&_Marketplace721.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace721 *Marketplace721Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace721.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace721 *Marketplace721Session) Receive() (*types.Transaction, error) {
	return _Marketplace721.Contract.Receive(&_Marketplace721.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace721 *Marketplace721TransactorSession) Receive() (*types.Transaction, error) {
	return _Marketplace721.Contract.Receive(&_Marketplace721.TransactOpts)
}

// Marketplace721ExchangedIterator is returned from FilterExchanged and is used to iterate over the raw logs and unpacked data for Exchanged events raised by the Marketplace721 contract.
type Marketplace721ExchangedIterator struct {
	Event *Marketplace721Exchanged // Event containing the contract specifics and raw log

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
func (it *Marketplace721ExchangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplace721Exchanged)
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
		it.Event = new(Marketplace721Exchanged)
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
func (it *Marketplace721ExchangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplace721ExchangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplace721Exchanged represents a Exchanged event raised by the Marketplace721 contract.
type Marketplace721Exchanged struct {
	GoodsId   *big.Int
	GoodsType *big.Int
	TokenId   *big.Int
	TradeType *big.Int
	Price     *big.Int
	Seller    common.Address
	Buyer     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExchanged is a free log retrieval operation binding the contract event 0x0b404f6b1f311cd2e20892c58db4774102012ddae1b440a6fb03385c066fad5a.
//
// Solidity: event Exchanged(uint256 goodsId, uint256 goodsType, uint256 tokenId, uint256 tradeType, uint256 price, address seller, address buyer)
func (_Marketplace721 *Marketplace721Filterer) FilterExchanged(opts *bind.FilterOpts) (*Marketplace721ExchangedIterator, error) {

	logs, sub, err := _Marketplace721.contract.FilterLogs(opts, "Exchanged")
	if err != nil {
		return nil, err
	}
	return &Marketplace721ExchangedIterator{contract: _Marketplace721.contract, event: "Exchanged", logs: logs, sub: sub}, nil
}

// WatchExchanged is a free log subscription operation binding the contract event 0x0b404f6b1f311cd2e20892c58db4774102012ddae1b440a6fb03385c066fad5a.
//
// Solidity: event Exchanged(uint256 goodsId, uint256 goodsType, uint256 tokenId, uint256 tradeType, uint256 price, address seller, address buyer)
func (_Marketplace721 *Marketplace721Filterer) WatchExchanged(opts *bind.WatchOpts, sink chan<- *Marketplace721Exchanged) (event.Subscription, error) {

	logs, sub, err := _Marketplace721.contract.WatchLogs(opts, "Exchanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplace721Exchanged)
				if err := _Marketplace721.contract.UnpackLog(event, "Exchanged", log); err != nil {
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

// ParseExchanged is a log parse operation binding the contract event 0x0b404f6b1f311cd2e20892c58db4774102012ddae1b440a6fb03385c066fad5a.
//
// Solidity: event Exchanged(uint256 goodsId, uint256 goodsType, uint256 tokenId, uint256 tradeType, uint256 price, address seller, address buyer)
func (_Marketplace721 *Marketplace721Filterer) ParseExchanged(log types.Log) (*Marketplace721Exchanged, error) {
	event := new(Marketplace721Exchanged)
	if err := _Marketplace721.contract.UnpackLog(event, "Exchanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplace721NewGoodsIterator is returned from FilterNewGoods and is used to iterate over the raw logs and unpacked data for NewGoods events raised by the Marketplace721 contract.
type Marketplace721NewGoodsIterator struct {
	Event *Marketplace721NewGoods // Event containing the contract specifics and raw log

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
func (it *Marketplace721NewGoodsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplace721NewGoods)
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
		it.Event = new(Marketplace721NewGoods)
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
func (it *Marketplace721NewGoodsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplace721NewGoodsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplace721NewGoods represents a NewGoods event raised by the Marketplace721 contract.
type Marketplace721NewGoods struct {
	GoodsType *big.Int
	Name      string
	Place     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewGoods is a free log retrieval operation binding the contract event 0x1e03217d1b68831d5bb050c4c781ec0c230fc2242b5d5966d6079a7ea841909f.
//
// Solidity: event NewGoods(uint256 goodsType, string name, address place)
func (_Marketplace721 *Marketplace721Filterer) FilterNewGoods(opts *bind.FilterOpts) (*Marketplace721NewGoodsIterator, error) {

	logs, sub, err := _Marketplace721.contract.FilterLogs(opts, "NewGoods")
	if err != nil {
		return nil, err
	}
	return &Marketplace721NewGoodsIterator{contract: _Marketplace721.contract, event: "NewGoods", logs: logs, sub: sub}, nil
}

// WatchNewGoods is a free log subscription operation binding the contract event 0x1e03217d1b68831d5bb050c4c781ec0c230fc2242b5d5966d6079a7ea841909f.
//
// Solidity: event NewGoods(uint256 goodsType, string name, address place)
func (_Marketplace721 *Marketplace721Filterer) WatchNewGoods(opts *bind.WatchOpts, sink chan<- *Marketplace721NewGoods) (event.Subscription, error) {

	logs, sub, err := _Marketplace721.contract.WatchLogs(opts, "NewGoods")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplace721NewGoods)
				if err := _Marketplace721.contract.UnpackLog(event, "NewGoods", log); err != nil {
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

// ParseNewGoods is a log parse operation binding the contract event 0x1e03217d1b68831d5bb050c4c781ec0c230fc2242b5d5966d6079a7ea841909f.
//
// Solidity: event NewGoods(uint256 goodsType, string name, address place)
func (_Marketplace721 *Marketplace721Filterer) ParseNewGoods(log types.Log) (*Marketplace721NewGoods, error) {
	event := new(Marketplace721NewGoods)
	if err := _Marketplace721.contract.UnpackLog(event, "NewGoods", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplace721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Marketplace721 contract.
type Marketplace721OwnershipTransferredIterator struct {
	Event *Marketplace721OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Marketplace721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplace721OwnershipTransferred)
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
		it.Event = new(Marketplace721OwnershipTransferred)
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
func (it *Marketplace721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplace721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplace721OwnershipTransferred represents a OwnershipTransferred event raised by the Marketplace721 contract.
type Marketplace721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace721 *Marketplace721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Marketplace721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Marketplace721OwnershipTransferredIterator{contract: _Marketplace721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace721 *Marketplace721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Marketplace721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplace721OwnershipTransferred)
				if err := _Marketplace721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Marketplace721 *Marketplace721Filterer) ParseOwnershipTransferred(log types.Log) (*Marketplace721OwnershipTransferred, error) {
	event := new(Marketplace721OwnershipTransferred)
	if err := _Marketplace721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Marketplace721UpdateGoodsStatusIterator is returned from FilterUpdateGoodsStatus and is used to iterate over the raw logs and unpacked data for UpdateGoodsStatus events raised by the Marketplace721 contract.
type Marketplace721UpdateGoodsStatusIterator struct {
	Event *Marketplace721UpdateGoodsStatus // Event containing the contract specifics and raw log

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
func (it *Marketplace721UpdateGoodsStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Marketplace721UpdateGoodsStatus)
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
		it.Event = new(Marketplace721UpdateGoodsStatus)
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
func (it *Marketplace721UpdateGoodsStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Marketplace721UpdateGoodsStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Marketplace721UpdateGoodsStatus represents a UpdateGoodsStatus event raised by the Marketplace721 contract.
type Marketplace721UpdateGoodsStatus struct {
	GoodsId   *big.Int
	Status    bool
	GoodsType *big.Int
	TokenId   *big.Int
	TradeType *big.Int
	Price     *big.Int
	Seller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateGoodsStatus is a free log retrieval operation binding the contract event 0x7b2ffaa020bf96f1f057269cdb102e3baac0e9d1b874ed622c0a7466568e9261.
//
// Solidity: event UpdateGoodsStatus(uint256 indexed goodsId, bool indexed status, uint256 goodsType, uint256 tokenId, uint256 tradeType, uint256 price, address seller)
func (_Marketplace721 *Marketplace721Filterer) FilterUpdateGoodsStatus(opts *bind.FilterOpts, goodsId []*big.Int, status []bool) (*Marketplace721UpdateGoodsStatusIterator, error) {

	var goodsIdRule []interface{}
	for _, goodsIdItem := range goodsId {
		goodsIdRule = append(goodsIdRule, goodsIdItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Marketplace721.contract.FilterLogs(opts, "UpdateGoodsStatus", goodsIdRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &Marketplace721UpdateGoodsStatusIterator{contract: _Marketplace721.contract, event: "UpdateGoodsStatus", logs: logs, sub: sub}, nil
}

// WatchUpdateGoodsStatus is a free log subscription operation binding the contract event 0x7b2ffaa020bf96f1f057269cdb102e3baac0e9d1b874ed622c0a7466568e9261.
//
// Solidity: event UpdateGoodsStatus(uint256 indexed goodsId, bool indexed status, uint256 goodsType, uint256 tokenId, uint256 tradeType, uint256 price, address seller)
func (_Marketplace721 *Marketplace721Filterer) WatchUpdateGoodsStatus(opts *bind.WatchOpts, sink chan<- *Marketplace721UpdateGoodsStatus, goodsId []*big.Int, status []bool) (event.Subscription, error) {

	var goodsIdRule []interface{}
	for _, goodsIdItem := range goodsId {
		goodsIdRule = append(goodsIdRule, goodsIdItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Marketplace721.contract.WatchLogs(opts, "UpdateGoodsStatus", goodsIdRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Marketplace721UpdateGoodsStatus)
				if err := _Marketplace721.contract.UnpackLog(event, "UpdateGoodsStatus", log); err != nil {
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

// ParseUpdateGoodsStatus is a log parse operation binding the contract event 0x7b2ffaa020bf96f1f057269cdb102e3baac0e9d1b874ed622c0a7466568e9261.
//
// Solidity: event UpdateGoodsStatus(uint256 indexed goodsId, bool indexed status, uint256 goodsType, uint256 tokenId, uint256 tradeType, uint256 price, address seller)
func (_Marketplace721 *Marketplace721Filterer) ParseUpdateGoodsStatus(log types.Log) (*Marketplace721UpdateGoodsStatus, error) {
	event := new(Marketplace721UpdateGoodsStatus)
	if err := _Marketplace721.contract.UnpackLog(event, "UpdateGoodsStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
