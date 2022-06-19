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

// PlanetABI is the input ABI used to generate the binding from.
const PlanetABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"PlanetID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AddTaxAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tax\",\"type\":\"uint256\"}],\"name\":\"ApplyFederalPlanet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BondPlanet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"CancelApply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"types_\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"motherPlanet\",\"type\":\"uint256\"}],\"name\":\"NewPlanet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"UpGradePlanet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tecNum\",\"type\":\"uint256\"}],\"name\":\"UpGradeTechnology\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BVG\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BVT\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addTaxAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tax_\",\"type\":\"uint256\"}],\"name\":\"applyFederalPlanet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"applyInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"applyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"applyTax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approveFedApply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"battleTaxRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"bondPlanet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bvInfo\",\"outputs\":[{\"internalType\":\"contractIBvInfo\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelApply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkPlanetOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"createNewPlanet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentPlanet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"febLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"federalPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"findTax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBVTPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"getUserPlanet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"isBonding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownerOfPlanet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"planet\",\"outputs\":[{\"internalType\":\"contractIPlanet721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"planetInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"population\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"normalTaxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"battleTaxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"motherPlanet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"types\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"membershipFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"populationLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"federalLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"federalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTax\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"planetType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"populationLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"federalLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"planetTax\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"pullOutPlanetCard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"BvInfo\",\"type\":\"address\"}],\"name\":\"setBvInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"setMemberShipFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"planet721_\",\"type\":\"address\"}],\"name\":\"setPlanet721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"types_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"populationLimit_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"federalLimit_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"planetTax_\",\"type\":\"uint256\"}],\"name\":\"setPlanetType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"BVG_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BVT_\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"upGradePlanet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upGradePlanetPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"planet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taxAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"userTaxAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Planet is an auto generated Go binding around an Ethereum contract.
type Planet struct {
	PlanetCaller     // Read-only binding to the contract
	PlanetTransactor // Write-only binding to the contract
	PlanetFilterer   // Log filterer for contract events
}

// PlanetCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlanetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlanetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlanetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlanetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlanetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlanetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlanetSession struct {
	Contract     *Planet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlanetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlanetCallerSession struct {
	Contract *PlanetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PlanetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlanetTransactorSession struct {
	Contract     *PlanetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlanetRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlanetRaw struct {
	Contract *Planet // Generic contract binding to access the raw methods on
}

// PlanetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlanetCallerRaw struct {
	Contract *PlanetCaller // Generic read-only contract binding to access the raw methods on
}

// PlanetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlanetTransactorRaw struct {
	Contract *PlanetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlanet creates a new instance of Planet, bound to a specific deployed contract.
func NewPlanet(address common.Address, backend bind.ContractBackend) (*Planet, error) {
	contract, err := bindPlanet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Planet{PlanetCaller: PlanetCaller{contract: contract}, PlanetTransactor: PlanetTransactor{contract: contract}, PlanetFilterer: PlanetFilterer{contract: contract}}, nil
}

// NewPlanetCaller creates a new read-only instance of Planet, bound to a specific deployed contract.
func NewPlanetCaller(address common.Address, caller bind.ContractCaller) (*PlanetCaller, error) {
	contract, err := bindPlanet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlanetCaller{contract: contract}, nil
}

// NewPlanetTransactor creates a new write-only instance of Planet, bound to a specific deployed contract.
func NewPlanetTransactor(address common.Address, transactor bind.ContractTransactor) (*PlanetTransactor, error) {
	contract, err := bindPlanet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlanetTransactor{contract: contract}, nil
}

// NewPlanetFilterer creates a new log filterer instance of Planet, bound to a specific deployed contract.
func NewPlanetFilterer(address common.Address, filterer bind.ContractFilterer) (*PlanetFilterer, error) {
	contract, err := bindPlanet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlanetFilterer{contract: contract}, nil
}

// bindPlanet binds a generic wrapper to an already deployed contract.
func bindPlanet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlanetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Planet *PlanetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Planet.Contract.PlanetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Planet *PlanetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Planet.Contract.PlanetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Planet *PlanetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Planet.Contract.PlanetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Planet *PlanetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Planet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Planet *PlanetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Planet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Planet *PlanetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Planet.Contract.contract.Transact(opts, method, params...)
}

// BVG is a free data retrieval call binding the contract method 0x87464389.
//
// Solidity: function BVG() view returns(address)
func (_Planet *PlanetCaller) BVG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "BVG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BVG is a free data retrieval call binding the contract method 0x87464389.
//
// Solidity: function BVG() view returns(address)
func (_Planet *PlanetSession) BVG() (common.Address, error) {
	return _Planet.Contract.BVG(&_Planet.CallOpts)
}

// BVG is a free data retrieval call binding the contract method 0x87464389.
//
// Solidity: function BVG() view returns(address)
func (_Planet *PlanetCallerSession) BVG() (common.Address, error) {
	return _Planet.Contract.BVG(&_Planet.CallOpts)
}

// BVT is a free data retrieval call binding the contract method 0x5f24b4a4.
//
// Solidity: function BVT() view returns(address)
func (_Planet *PlanetCaller) BVT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "BVT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BVT is a free data retrieval call binding the contract method 0x5f24b4a4.
//
// Solidity: function BVT() view returns(address)
func (_Planet *PlanetSession) BVT() (common.Address, error) {
	return _Planet.Contract.BVT(&_Planet.CallOpts)
}

// BVT is a free data retrieval call binding the contract method 0x5f24b4a4.
//
// Solidity: function BVT() view returns(address)
func (_Planet *PlanetCallerSession) BVT() (common.Address, error) {
	return _Planet.Contract.BVT(&_Planet.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Planet *PlanetCaller) Admin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "admin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Planet *PlanetSession) Admin(arg0 common.Address) (bool, error) {
	return _Planet.Contract.Admin(&_Planet.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Planet *PlanetCallerSession) Admin(arg0 common.Address) (bool, error) {
	return _Planet.Contract.Admin(&_Planet.CallOpts, arg0)
}

// ApplyInfo is a free data retrieval call binding the contract method 0xf9c0eace.
//
// Solidity: function applyInfo(address ) view returns(uint256 applyAmount, uint256 applyTax, uint256 lockAmount)
func (_Planet *PlanetCaller) ApplyInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	ApplyAmount *big.Int
	ApplyTax    *big.Int
	LockAmount  *big.Int
}, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "applyInfo", arg0)

	outstruct := new(struct {
		ApplyAmount *big.Int
		ApplyTax    *big.Int
		LockAmount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ApplyAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ApplyTax = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LockAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ApplyInfo is a free data retrieval call binding the contract method 0xf9c0eace.
//
// Solidity: function applyInfo(address ) view returns(uint256 applyAmount, uint256 applyTax, uint256 lockAmount)
func (_Planet *PlanetSession) ApplyInfo(arg0 common.Address) (struct {
	ApplyAmount *big.Int
	ApplyTax    *big.Int
	LockAmount  *big.Int
}, error) {
	return _Planet.Contract.ApplyInfo(&_Planet.CallOpts, arg0)
}

// ApplyInfo is a free data retrieval call binding the contract method 0xf9c0eace.
//
// Solidity: function applyInfo(address ) view returns(uint256 applyAmount, uint256 applyTax, uint256 lockAmount)
func (_Planet *PlanetCallerSession) ApplyInfo(arg0 common.Address) (struct {
	ApplyAmount *big.Int
	ApplyTax    *big.Int
	LockAmount  *big.Int
}, error) {
	return _Planet.Contract.ApplyInfo(&_Planet.CallOpts, arg0)
}

// BattleTaxRate is a free data retrieval call binding the contract method 0x29cf0227.
//
// Solidity: function battleTaxRate() view returns(uint256)
func (_Planet *PlanetCaller) BattleTaxRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "battleTaxRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BattleTaxRate is a free data retrieval call binding the contract method 0x29cf0227.
//
// Solidity: function battleTaxRate() view returns(uint256)
func (_Planet *PlanetSession) BattleTaxRate() (*big.Int, error) {
	return _Planet.Contract.BattleTaxRate(&_Planet.CallOpts)
}

// BattleTaxRate is a free data retrieval call binding the contract method 0x29cf0227.
//
// Solidity: function battleTaxRate() view returns(uint256)
func (_Planet *PlanetCallerSession) BattleTaxRate() (*big.Int, error) {
	return _Planet.Contract.BattleTaxRate(&_Planet.CallOpts)
}

// BvInfo is a free data retrieval call binding the contract method 0x85ef848a.
//
// Solidity: function bvInfo() view returns(address)
func (_Planet *PlanetCaller) BvInfo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "bvInfo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BvInfo is a free data retrieval call binding the contract method 0x85ef848a.
//
// Solidity: function bvInfo() view returns(address)
func (_Planet *PlanetSession) BvInfo() (common.Address, error) {
	return _Planet.Contract.BvInfo(&_Planet.CallOpts)
}

// BvInfo is a free data retrieval call binding the contract method 0x85ef848a.
//
// Solidity: function bvInfo() view returns(address)
func (_Planet *PlanetCallerSession) BvInfo() (common.Address, error) {
	return _Planet.Contract.BvInfo(&_Planet.CallOpts)
}

// CheckPlanetOwner is a free data retrieval call binding the contract method 0x7547ad6d.
//
// Solidity: function checkPlanetOwner() view returns(uint256[], address[])
func (_Planet *PlanetCaller) CheckPlanetOwner(opts *bind.CallOpts) ([]*big.Int, []common.Address, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "checkPlanetOwner")

	if err != nil {
		return *new([]*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// CheckPlanetOwner is a free data retrieval call binding the contract method 0x7547ad6d.
//
// Solidity: function checkPlanetOwner() view returns(uint256[], address[])
func (_Planet *PlanetSession) CheckPlanetOwner() ([]*big.Int, []common.Address, error) {
	return _Planet.Contract.CheckPlanetOwner(&_Planet.CallOpts)
}

// CheckPlanetOwner is a free data retrieval call binding the contract method 0x7547ad6d.
//
// Solidity: function checkPlanetOwner() view returns(uint256[], address[])
func (_Planet *PlanetCallerSession) CheckPlanetOwner() ([]*big.Int, []common.Address, error) {
	return _Planet.Contract.CheckPlanetOwner(&_Planet.CallOpts)
}

// CurrentPlanet is a free data retrieval call binding the contract method 0xc1c9085b.
//
// Solidity: function currentPlanet(uint256 ) view returns(uint256)
func (_Planet *PlanetCaller) CurrentPlanet(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "currentPlanet", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPlanet is a free data retrieval call binding the contract method 0xc1c9085b.
//
// Solidity: function currentPlanet(uint256 ) view returns(uint256)
func (_Planet *PlanetSession) CurrentPlanet(arg0 *big.Int) (*big.Int, error) {
	return _Planet.Contract.CurrentPlanet(&_Planet.CallOpts, arg0)
}

// CurrentPlanet is a free data retrieval call binding the contract method 0xc1c9085b.
//
// Solidity: function currentPlanet(uint256 ) view returns(uint256)
func (_Planet *PlanetCallerSession) CurrentPlanet(arg0 *big.Int) (*big.Int, error) {
	return _Planet.Contract.CurrentPlanet(&_Planet.CallOpts, arg0)
}

// FebLimit is a free data retrieval call binding the contract method 0x863b3ec5.
//
// Solidity: function febLimit() view returns(uint256)
func (_Planet *PlanetCaller) FebLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "febLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FebLimit is a free data retrieval call binding the contract method 0x863b3ec5.
//
// Solidity: function febLimit() view returns(uint256)
func (_Planet *PlanetSession) FebLimit() (*big.Int, error) {
	return _Planet.Contract.FebLimit(&_Planet.CallOpts)
}

// FebLimit is a free data retrieval call binding the contract method 0x863b3ec5.
//
// Solidity: function febLimit() view returns(uint256)
func (_Planet *PlanetCallerSession) FebLimit() (*big.Int, error) {
	return _Planet.Contract.FebLimit(&_Planet.CallOpts)
}

// FederalPrice is a free data retrieval call binding the contract method 0xa445f0ee.
//
// Solidity: function federalPrice() view returns(uint256)
func (_Planet *PlanetCaller) FederalPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "federalPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FederalPrice is a free data retrieval call binding the contract method 0xa445f0ee.
//
// Solidity: function federalPrice() view returns(uint256)
func (_Planet *PlanetSession) FederalPrice() (*big.Int, error) {
	return _Planet.Contract.FederalPrice(&_Planet.CallOpts)
}

// FederalPrice is a free data retrieval call binding the contract method 0xa445f0ee.
//
// Solidity: function federalPrice() view returns(uint256)
func (_Planet *PlanetCallerSession) FederalPrice() (*big.Int, error) {
	return _Planet.Contract.FederalPrice(&_Planet.CallOpts)
}

// FindTax is a free data retrieval call binding the contract method 0xe86fc831.
//
// Solidity: function findTax(address addr_) view returns(uint256)
func (_Planet *PlanetCaller) FindTax(opts *bind.CallOpts, addr_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "findTax", addr_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FindTax is a free data retrieval call binding the contract method 0xe86fc831.
//
// Solidity: function findTax(address addr_) view returns(uint256)
func (_Planet *PlanetSession) FindTax(addr_ common.Address) (*big.Int, error) {
	return _Planet.Contract.FindTax(&_Planet.CallOpts, addr_)
}

// FindTax is a free data retrieval call binding the contract method 0xe86fc831.
//
// Solidity: function findTax(address addr_) view returns(uint256)
func (_Planet *PlanetCallerSession) FindTax(addr_ common.Address) (*big.Int, error) {
	return _Planet.Contract.FindTax(&_Planet.CallOpts, addr_)
}

// GetBVTPrice is a free data retrieval call binding the contract method 0x80dd5882.
//
// Solidity: function getBVTPrice() view returns(uint256)
func (_Planet *PlanetCaller) GetBVTPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "getBVTPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBVTPrice is a free data retrieval call binding the contract method 0x80dd5882.
//
// Solidity: function getBVTPrice() view returns(uint256)
func (_Planet *PlanetSession) GetBVTPrice() (*big.Int, error) {
	return _Planet.Contract.GetBVTPrice(&_Planet.CallOpts)
}

// GetBVTPrice is a free data retrieval call binding the contract method 0x80dd5882.
//
// Solidity: function getBVTPrice() view returns(uint256)
func (_Planet *PlanetCallerSession) GetBVTPrice() (*big.Int, error) {
	return _Planet.Contract.GetBVTPrice(&_Planet.CallOpts)
}

// GetUserPlanet is a free data retrieval call binding the contract method 0x49b10297.
//
// Solidity: function getUserPlanet(address addr_) view returns(uint256)
func (_Planet *PlanetCaller) GetUserPlanet(opts *bind.CallOpts, addr_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "getUserPlanet", addr_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserPlanet is a free data retrieval call binding the contract method 0x49b10297.
//
// Solidity: function getUserPlanet(address addr_) view returns(uint256)
func (_Planet *PlanetSession) GetUserPlanet(addr_ common.Address) (*big.Int, error) {
	return _Planet.Contract.GetUserPlanet(&_Planet.CallOpts, addr_)
}

// GetUserPlanet is a free data retrieval call binding the contract method 0x49b10297.
//
// Solidity: function getUserPlanet(address addr_) view returns(uint256)
func (_Planet *PlanetCallerSession) GetUserPlanet(addr_ common.Address) (*big.Int, error) {
	return _Planet.Contract.GetUserPlanet(&_Planet.CallOpts, addr_)
}

// IsBonding is a free data retrieval call binding the contract method 0x13ce3d69.
//
// Solidity: function isBonding(address addr_) view returns(bool)
func (_Planet *PlanetCaller) IsBonding(opts *bind.CallOpts, addr_ common.Address) (bool, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "isBonding", addr_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBonding is a free data retrieval call binding the contract method 0x13ce3d69.
//
// Solidity: function isBonding(address addr_) view returns(bool)
func (_Planet *PlanetSession) IsBonding(addr_ common.Address) (bool, error) {
	return _Planet.Contract.IsBonding(&_Planet.CallOpts, addr_)
}

// IsBonding is a free data retrieval call binding the contract method 0x13ce3d69.
//
// Solidity: function isBonding(address addr_) view returns(bool)
func (_Planet *PlanetCallerSession) IsBonding(addr_ common.Address) (bool, error) {
	return _Planet.Contract.IsBonding(&_Planet.CallOpts, addr_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Planet *PlanetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Planet *PlanetSession) Owner() (common.Address, error) {
	return _Planet.Contract.Owner(&_Planet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Planet *PlanetCallerSession) Owner() (common.Address, error) {
	return _Planet.Contract.Owner(&_Planet.CallOpts)
}

// OwnerOfPlanet is a free data retrieval call binding the contract method 0xaa85a43f.
//
// Solidity: function ownerOfPlanet(address ) view returns(uint256)
func (_Planet *PlanetCaller) OwnerOfPlanet(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "ownerOfPlanet", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerOfPlanet is a free data retrieval call binding the contract method 0xaa85a43f.
//
// Solidity: function ownerOfPlanet(address ) view returns(uint256)
func (_Planet *PlanetSession) OwnerOfPlanet(arg0 common.Address) (*big.Int, error) {
	return _Planet.Contract.OwnerOfPlanet(&_Planet.CallOpts, arg0)
}

// OwnerOfPlanet is a free data retrieval call binding the contract method 0xaa85a43f.
//
// Solidity: function ownerOfPlanet(address ) view returns(uint256)
func (_Planet *PlanetCallerSession) OwnerOfPlanet(arg0 common.Address) (*big.Int, error) {
	return _Planet.Contract.OwnerOfPlanet(&_Planet.CallOpts, arg0)
}

// Planet is a free data retrieval call binding the contract method 0x8443bea6.
//
// Solidity: function planet() view returns(address)
func (_Planet *PlanetCaller) Planet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "planet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Planet is a free data retrieval call binding the contract method 0x8443bea6.
//
// Solidity: function planet() view returns(address)
func (_Planet *PlanetSession) Planet() (common.Address, error) {
	return _Planet.Contract.Planet(&_Planet.CallOpts)
}

// Planet is a free data retrieval call binding the contract method 0x8443bea6.
//
// Solidity: function planet() view returns(address)
func (_Planet *PlanetCallerSession) Planet() (common.Address, error) {
	return _Planet.Contract.Planet(&_Planet.CallOpts)
}

// PlanetInfo is a free data retrieval call binding the contract method 0x7de50469.
//
// Solidity: function planetInfo(uint256 ) view returns(address owner, uint256 tax, uint256 population, uint256 normalTaxAmount, uint256 battleTaxAmount, uint256 motherPlanet, uint256 types, uint256 membershipFee, uint256 populationLimit, uint256 federalLimit, uint256 federalAmount, uint256 totalTax)
func (_Planet *PlanetCaller) PlanetInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner           common.Address
	Tax             *big.Int
	Population      *big.Int
	NormalTaxAmount *big.Int
	BattleTaxAmount *big.Int
	MotherPlanet    *big.Int
	Types           *big.Int
	MembershipFee   *big.Int
	PopulationLimit *big.Int
	FederalLimit    *big.Int
	FederalAmount   *big.Int
	TotalTax        *big.Int
}, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "planetInfo", arg0)

	outstruct := new(struct {
		Owner           common.Address
		Tax             *big.Int
		Population      *big.Int
		NormalTaxAmount *big.Int
		BattleTaxAmount *big.Int
		MotherPlanet    *big.Int
		Types           *big.Int
		MembershipFee   *big.Int
		PopulationLimit *big.Int
		FederalLimit    *big.Int
		FederalAmount   *big.Int
		TotalTax        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Tax = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Population = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NormalTaxAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BattleTaxAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MotherPlanet = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Types = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.MembershipFee = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.PopulationLimit = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FederalLimit = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.FederalAmount = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.TotalTax = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PlanetInfo is a free data retrieval call binding the contract method 0x7de50469.
//
// Solidity: function planetInfo(uint256 ) view returns(address owner, uint256 tax, uint256 population, uint256 normalTaxAmount, uint256 battleTaxAmount, uint256 motherPlanet, uint256 types, uint256 membershipFee, uint256 populationLimit, uint256 federalLimit, uint256 federalAmount, uint256 totalTax)
func (_Planet *PlanetSession) PlanetInfo(arg0 *big.Int) (struct {
	Owner           common.Address
	Tax             *big.Int
	Population      *big.Int
	NormalTaxAmount *big.Int
	BattleTaxAmount *big.Int
	MotherPlanet    *big.Int
	Types           *big.Int
	MembershipFee   *big.Int
	PopulationLimit *big.Int
	FederalLimit    *big.Int
	FederalAmount   *big.Int
	TotalTax        *big.Int
}, error) {
	return _Planet.Contract.PlanetInfo(&_Planet.CallOpts, arg0)
}

// PlanetInfo is a free data retrieval call binding the contract method 0x7de50469.
//
// Solidity: function planetInfo(uint256 ) view returns(address owner, uint256 tax, uint256 population, uint256 normalTaxAmount, uint256 battleTaxAmount, uint256 motherPlanet, uint256 types, uint256 membershipFee, uint256 populationLimit, uint256 federalLimit, uint256 federalAmount, uint256 totalTax)
func (_Planet *PlanetCallerSession) PlanetInfo(arg0 *big.Int) (struct {
	Owner           common.Address
	Tax             *big.Int
	Population      *big.Int
	NormalTaxAmount *big.Int
	BattleTaxAmount *big.Int
	MotherPlanet    *big.Int
	Types           *big.Int
	MembershipFee   *big.Int
	PopulationLimit *big.Int
	FederalLimit    *big.Int
	FederalAmount   *big.Int
	TotalTax        *big.Int
}, error) {
	return _Planet.Contract.PlanetInfo(&_Planet.CallOpts, arg0)
}

// PlanetType is a free data retrieval call binding the contract method 0x64e316f1.
//
// Solidity: function planetType(uint256 ) view returns(uint256 populationLimit, uint256 federalLimit, uint256 planetTax)
func (_Planet *PlanetCaller) PlanetType(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PopulationLimit *big.Int
	FederalLimit    *big.Int
	PlanetTax       *big.Int
}, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "planetType", arg0)

	outstruct := new(struct {
		PopulationLimit *big.Int
		FederalLimit    *big.Int
		PlanetTax       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PopulationLimit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FederalLimit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PlanetTax = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PlanetType is a free data retrieval call binding the contract method 0x64e316f1.
//
// Solidity: function planetType(uint256 ) view returns(uint256 populationLimit, uint256 federalLimit, uint256 planetTax)
func (_Planet *PlanetSession) PlanetType(arg0 *big.Int) (struct {
	PopulationLimit *big.Int
	FederalLimit    *big.Int
	PlanetTax       *big.Int
}, error) {
	return _Planet.Contract.PlanetType(&_Planet.CallOpts, arg0)
}

// PlanetType is a free data retrieval call binding the contract method 0x64e316f1.
//
// Solidity: function planetType(uint256 ) view returns(uint256 populationLimit, uint256 federalLimit, uint256 planetTax)
func (_Planet *PlanetCallerSession) PlanetType(arg0 *big.Int) (struct {
	PopulationLimit *big.Int
	FederalLimit    *big.Int
	PlanetTax       *big.Int
}, error) {
	return _Planet.Contract.PlanetType(&_Planet.CallOpts, arg0)
}

// UpGradePlanetPrice is a free data retrieval call binding the contract method 0x258214d3.
//
// Solidity: function upGradePlanetPrice() view returns(uint256)
func (_Planet *PlanetCaller) UpGradePlanetPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "upGradePlanetPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpGradePlanetPrice is a free data retrieval call binding the contract method 0x258214d3.
//
// Solidity: function upGradePlanetPrice() view returns(uint256)
func (_Planet *PlanetSession) UpGradePlanetPrice() (*big.Int, error) {
	return _Planet.Contract.UpGradePlanetPrice(&_Planet.CallOpts)
}

// UpGradePlanetPrice is a free data retrieval call binding the contract method 0x258214d3.
//
// Solidity: function upGradePlanetPrice() view returns(uint256)
func (_Planet *PlanetCallerSession) UpGradePlanetPrice() (*big.Int, error) {
	return _Planet.Contract.UpGradePlanetPrice(&_Planet.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 level, uint256 planet, uint256 taxAmount)
func (_Planet *PlanetCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Level     *big.Int
	Planet    *big.Int
	TaxAmount *big.Int
}, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Level     *big.Int
		Planet    *big.Int
		TaxAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Level = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Planet = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TaxAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 level, uint256 planet, uint256 taxAmount)
func (_Planet *PlanetSession) UserInfo(arg0 common.Address) (struct {
	Level     *big.Int
	Planet    *big.Int
	TaxAmount *big.Int
}, error) {
	return _Planet.Contract.UserInfo(&_Planet.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 level, uint256 planet, uint256 taxAmount)
func (_Planet *PlanetCallerSession) UserInfo(arg0 common.Address) (struct {
	Level     *big.Int
	Planet    *big.Int
	TaxAmount *big.Int
}, error) {
	return _Planet.Contract.UserInfo(&_Planet.CallOpts, arg0)
}

// UserTaxAmount is a free data retrieval call binding the contract method 0x22c76608.
//
// Solidity: function userTaxAmount(address addr) view returns(uint256)
func (_Planet *PlanetCaller) UserTaxAmount(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Planet.contract.Call(opts, &out, "userTaxAmount", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTaxAmount is a free data retrieval call binding the contract method 0x22c76608.
//
// Solidity: function userTaxAmount(address addr) view returns(uint256)
func (_Planet *PlanetSession) UserTaxAmount(addr common.Address) (*big.Int, error) {
	return _Planet.Contract.UserTaxAmount(&_Planet.CallOpts, addr)
}

// UserTaxAmount is a free data retrieval call binding the contract method 0x22c76608.
//
// Solidity: function userTaxAmount(address addr) view returns(uint256)
func (_Planet *PlanetCallerSession) UserTaxAmount(addr common.Address) (*big.Int, error) {
	return _Planet.Contract.UserTaxAmount(&_Planet.CallOpts, addr)
}

// AddTaxAmount is a paid mutator transaction binding the contract method 0x832403f3.
//
// Solidity: function addTaxAmount(address addr, uint256 amount) returns()
func (_Planet *PlanetTransactor) AddTaxAmount(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "addTaxAmount", addr, amount)
}

// AddTaxAmount is a paid mutator transaction binding the contract method 0x832403f3.
//
// Solidity: function addTaxAmount(address addr, uint256 amount) returns()
func (_Planet *PlanetSession) AddTaxAmount(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.AddTaxAmount(&_Planet.TransactOpts, addr, amount)
}

// AddTaxAmount is a paid mutator transaction binding the contract method 0x832403f3.
//
// Solidity: function addTaxAmount(address addr, uint256 amount) returns()
func (_Planet *PlanetTransactorSession) AddTaxAmount(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.AddTaxAmount(&_Planet.TransactOpts, addr, amount)
}

// ApplyFederalPlanet is a paid mutator transaction binding the contract method 0x3cad17cf.
//
// Solidity: function applyFederalPlanet(uint256 amount, uint256 tax_) returns()
func (_Planet *PlanetTransactor) ApplyFederalPlanet(opts *bind.TransactOpts, amount *big.Int, tax_ *big.Int) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "applyFederalPlanet", amount, tax_)
}

// ApplyFederalPlanet is a paid mutator transaction binding the contract method 0x3cad17cf.
//
// Solidity: function applyFederalPlanet(uint256 amount, uint256 tax_) returns()
func (_Planet *PlanetSession) ApplyFederalPlanet(amount *big.Int, tax_ *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.ApplyFederalPlanet(&_Planet.TransactOpts, amount, tax_)
}

// ApplyFederalPlanet is a paid mutator transaction binding the contract method 0x3cad17cf.
//
// Solidity: function applyFederalPlanet(uint256 amount, uint256 tax_) returns()
func (_Planet *PlanetTransactorSession) ApplyFederalPlanet(amount *big.Int, tax_ *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.ApplyFederalPlanet(&_Planet.TransactOpts, amount, tax_)
}

// ApproveFedApply is a paid mutator transaction binding the contract method 0x730f6d70.
//
// Solidity: function approveFedApply(address addr_, uint256 tokenId) returns()
func (_Planet *PlanetTransactor) ApproveFedApply(opts *bind.TransactOpts, addr_ common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "approveFedApply", addr_, tokenId)
}

// ApproveFedApply is a paid mutator transaction binding the contract method 0x730f6d70.
//
// Solidity: function approveFedApply(address addr_, uint256 tokenId) returns()
func (_Planet *PlanetSession) ApproveFedApply(addr_ common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.ApproveFedApply(&_Planet.TransactOpts, addr_, tokenId)
}

// ApproveFedApply is a paid mutator transaction binding the contract method 0x730f6d70.
//
// Solidity: function approveFedApply(address addr_, uint256 tokenId) returns()
func (_Planet *PlanetTransactorSession) ApproveFedApply(addr_ common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.ApproveFedApply(&_Planet.TransactOpts, addr_, tokenId)
}

// BondPlanet is a paid mutator transaction binding the contract method 0xbab82717.
//
// Solidity: function bondPlanet(uint256 tokenId) returns()
func (_Planet *PlanetTransactor) BondPlanet(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "bondPlanet", tokenId)
}

// BondPlanet is a paid mutator transaction binding the contract method 0xbab82717.
//
// Solidity: function bondPlanet(uint256 tokenId) returns()
func (_Planet *PlanetSession) BondPlanet(tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.BondPlanet(&_Planet.TransactOpts, tokenId)
}

// BondPlanet is a paid mutator transaction binding the contract method 0xbab82717.
//
// Solidity: function bondPlanet(uint256 tokenId) returns()
func (_Planet *PlanetTransactorSession) BondPlanet(tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.BondPlanet(&_Planet.TransactOpts, tokenId)
}

// CancelApply is a paid mutator transaction binding the contract method 0x17cd1849.
//
// Solidity: function cancelApply() returns()
func (_Planet *PlanetTransactor) CancelApply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "cancelApply")
}

// CancelApply is a paid mutator transaction binding the contract method 0x17cd1849.
//
// Solidity: function cancelApply() returns()
func (_Planet *PlanetSession) CancelApply() (*types.Transaction, error) {
	return _Planet.Contract.CancelApply(&_Planet.TransactOpts)
}

// CancelApply is a paid mutator transaction binding the contract method 0x17cd1849.
//
// Solidity: function cancelApply() returns()
func (_Planet *PlanetTransactorSession) CancelApply() (*types.Transaction, error) {
	return _Planet.Contract.CancelApply(&_Planet.TransactOpts)
}

// CreateNewPlanet is a paid mutator transaction binding the contract method 0x5bb5fa92.
//
// Solidity: function createNewPlanet(uint256 tokenId) returns()
func (_Planet *PlanetTransactor) CreateNewPlanet(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "createNewPlanet", tokenId)
}

// CreateNewPlanet is a paid mutator transaction binding the contract method 0x5bb5fa92.
//
// Solidity: function createNewPlanet(uint256 tokenId) returns()
func (_Planet *PlanetSession) CreateNewPlanet(tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.CreateNewPlanet(&_Planet.TransactOpts, tokenId)
}

// CreateNewPlanet is a paid mutator transaction binding the contract method 0x5bb5fa92.
//
// Solidity: function createNewPlanet(uint256 tokenId) returns()
func (_Planet *PlanetTransactorSession) CreateNewPlanet(tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.CreateNewPlanet(&_Planet.TransactOpts, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Planet *PlanetTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Planet *PlanetSession) Initialize() (*types.Transaction, error) {
	return _Planet.Contract.Initialize(&_Planet.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Planet *PlanetTransactorSession) Initialize() (*types.Transaction, error) {
	return _Planet.Contract.Initialize(&_Planet.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Planet *PlanetTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Planet *PlanetSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Planet.Contract.OnERC721Received(&_Planet.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Planet *PlanetTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Planet.Contract.OnERC721Received(&_Planet.TransactOpts, arg0, arg1, arg2, arg3)
}

// PullOutPlanetCard is a paid mutator transaction binding the contract method 0x16929e20.
//
// Solidity: function pullOutPlanetCard(uint256 tokenId) returns()
func (_Planet *PlanetTransactor) PullOutPlanetCard(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "pullOutPlanetCard", tokenId)
}

// PullOutPlanetCard is a paid mutator transaction binding the contract method 0x16929e20.
//
// Solidity: function pullOutPlanetCard(uint256 tokenId) returns()
func (_Planet *PlanetSession) PullOutPlanetCard(tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.PullOutPlanetCard(&_Planet.TransactOpts, tokenId)
}

// PullOutPlanetCard is a paid mutator transaction binding the contract method 0x16929e20.
//
// Solidity: function pullOutPlanetCard(uint256 tokenId) returns()
func (_Planet *PlanetTransactorSession) PullOutPlanetCard(tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.PullOutPlanetCard(&_Planet.TransactOpts, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Planet *PlanetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Planet *PlanetSession) RenounceOwnership() (*types.Transaction, error) {
	return _Planet.Contract.RenounceOwnership(&_Planet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Planet *PlanetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Planet.Contract.RenounceOwnership(&_Planet.TransactOpts)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xa5e107bd.
//
// Solidity: function replaceOwner(uint256 tokenId) returns()
func (_Planet *PlanetTransactor) ReplaceOwner(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "replaceOwner", tokenId)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xa5e107bd.
//
// Solidity: function replaceOwner(uint256 tokenId) returns()
func (_Planet *PlanetSession) ReplaceOwner(tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.ReplaceOwner(&_Planet.TransactOpts, tokenId)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xa5e107bd.
//
// Solidity: function replaceOwner(uint256 tokenId) returns()
func (_Planet *PlanetTransactorSession) ReplaceOwner(tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.ReplaceOwner(&_Planet.TransactOpts, tokenId)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr, bool b) returns()
func (_Planet *PlanetTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address, b bool) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "setAdmin", addr, b)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr, bool b) returns()
func (_Planet *PlanetSession) SetAdmin(addr common.Address, b bool) (*types.Transaction, error) {
	return _Planet.Contract.SetAdmin(&_Planet.TransactOpts, addr, b)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr, bool b) returns()
func (_Planet *PlanetTransactorSession) SetAdmin(addr common.Address, b bool) (*types.Transaction, error) {
	return _Planet.Contract.SetAdmin(&_Planet.TransactOpts, addr, b)
}

// SetBvInfo is a paid mutator transaction binding the contract method 0xd589e7ff.
//
// Solidity: function setBvInfo(address BvInfo) returns()
func (_Planet *PlanetTransactor) SetBvInfo(opts *bind.TransactOpts, BvInfo common.Address) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "setBvInfo", BvInfo)
}

// SetBvInfo is a paid mutator transaction binding the contract method 0xd589e7ff.
//
// Solidity: function setBvInfo(address BvInfo) returns()
func (_Planet *PlanetSession) SetBvInfo(BvInfo common.Address) (*types.Transaction, error) {
	return _Planet.Contract.SetBvInfo(&_Planet.TransactOpts, BvInfo)
}

// SetBvInfo is a paid mutator transaction binding the contract method 0xd589e7ff.
//
// Solidity: function setBvInfo(address BvInfo) returns()
func (_Planet *PlanetTransactorSession) SetBvInfo(BvInfo common.Address) (*types.Transaction, error) {
	return _Planet.Contract.SetBvInfo(&_Planet.TransactOpts, BvInfo)
}

// SetMemberShipFee is a paid mutator transaction binding the contract method 0xfbce174b.
//
// Solidity: function setMemberShipFee(uint256 tokenId, uint256 price_) returns()
func (_Planet *PlanetTransactor) SetMemberShipFee(opts *bind.TransactOpts, tokenId *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "setMemberShipFee", tokenId, price_)
}

// SetMemberShipFee is a paid mutator transaction binding the contract method 0xfbce174b.
//
// Solidity: function setMemberShipFee(uint256 tokenId, uint256 price_) returns()
func (_Planet *PlanetSession) SetMemberShipFee(tokenId *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.SetMemberShipFee(&_Planet.TransactOpts, tokenId, price_)
}

// SetMemberShipFee is a paid mutator transaction binding the contract method 0xfbce174b.
//
// Solidity: function setMemberShipFee(uint256 tokenId, uint256 price_) returns()
func (_Planet *PlanetTransactorSession) SetMemberShipFee(tokenId *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.SetMemberShipFee(&_Planet.TransactOpts, tokenId, price_)
}

// SetPlanet721 is a paid mutator transaction binding the contract method 0x477fd37a.
//
// Solidity: function setPlanet721(address planet721_) returns()
func (_Planet *PlanetTransactor) SetPlanet721(opts *bind.TransactOpts, planet721_ common.Address) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "setPlanet721", planet721_)
}

// SetPlanet721 is a paid mutator transaction binding the contract method 0x477fd37a.
//
// Solidity: function setPlanet721(address planet721_) returns()
func (_Planet *PlanetSession) SetPlanet721(planet721_ common.Address) (*types.Transaction, error) {
	return _Planet.Contract.SetPlanet721(&_Planet.TransactOpts, planet721_)
}

// SetPlanet721 is a paid mutator transaction binding the contract method 0x477fd37a.
//
// Solidity: function setPlanet721(address planet721_) returns()
func (_Planet *PlanetTransactorSession) SetPlanet721(planet721_ common.Address) (*types.Transaction, error) {
	return _Planet.Contract.SetPlanet721(&_Planet.TransactOpts, planet721_)
}

// SetPlanetType is a paid mutator transaction binding the contract method 0x02a12647.
//
// Solidity: function setPlanetType(uint256 types_, uint256 populationLimit_, uint256 federalLimit_, uint256 planetTax_) returns()
func (_Planet *PlanetTransactor) SetPlanetType(opts *bind.TransactOpts, types_ *big.Int, populationLimit_ *big.Int, federalLimit_ *big.Int, planetTax_ *big.Int) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "setPlanetType", types_, populationLimit_, federalLimit_, planetTax_)
}

// SetPlanetType is a paid mutator transaction binding the contract method 0x02a12647.
//
// Solidity: function setPlanetType(uint256 types_, uint256 populationLimit_, uint256 federalLimit_, uint256 planetTax_) returns()
func (_Planet *PlanetSession) SetPlanetType(types_ *big.Int, populationLimit_ *big.Int, federalLimit_ *big.Int, planetTax_ *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.SetPlanetType(&_Planet.TransactOpts, types_, populationLimit_, federalLimit_, planetTax_)
}

// SetPlanetType is a paid mutator transaction binding the contract method 0x02a12647.
//
// Solidity: function setPlanetType(uint256 types_, uint256 populationLimit_, uint256 federalLimit_, uint256 planetTax_) returns()
func (_Planet *PlanetTransactorSession) SetPlanetType(types_ *big.Int, populationLimit_ *big.Int, federalLimit_ *big.Int, planetTax_ *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.SetPlanetType(&_Planet.TransactOpts, types_, populationLimit_, federalLimit_, planetTax_)
}

// SetToken is a paid mutator transaction binding the contract method 0x1da26a8b.
//
// Solidity: function setToken(address BVG_, address BVT_) returns()
func (_Planet *PlanetTransactor) SetToken(opts *bind.TransactOpts, BVG_ common.Address, BVT_ common.Address) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "setToken", BVG_, BVT_)
}

// SetToken is a paid mutator transaction binding the contract method 0x1da26a8b.
//
// Solidity: function setToken(address BVG_, address BVT_) returns()
func (_Planet *PlanetSession) SetToken(BVG_ common.Address, BVT_ common.Address) (*types.Transaction, error) {
	return _Planet.Contract.SetToken(&_Planet.TransactOpts, BVG_, BVT_)
}

// SetToken is a paid mutator transaction binding the contract method 0x1da26a8b.
//
// Solidity: function setToken(address BVG_, address BVT_) returns()
func (_Planet *PlanetTransactorSession) SetToken(BVG_ common.Address, BVT_ common.Address) (*types.Transaction, error) {
	return _Planet.Contract.SetToken(&_Planet.TransactOpts, BVG_, BVT_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Planet *PlanetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Planet *PlanetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Planet.Contract.TransferOwnership(&_Planet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Planet *PlanetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Planet.Contract.TransferOwnership(&_Planet.TransactOpts, newOwner)
}

// UpGradePlanet is a paid mutator transaction binding the contract method 0x5bec01af.
//
// Solidity: function upGradePlanet(uint256 tokenId) returns()
func (_Planet *PlanetTransactor) UpGradePlanet(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.contract.Transact(opts, "upGradePlanet", tokenId)
}

// UpGradePlanet is a paid mutator transaction binding the contract method 0x5bec01af.
//
// Solidity: function upGradePlanet(uint256 tokenId) returns()
func (_Planet *PlanetSession) UpGradePlanet(tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.UpGradePlanet(&_Planet.TransactOpts, tokenId)
}

// UpGradePlanet is a paid mutator transaction binding the contract method 0x5bec01af.
//
// Solidity: function upGradePlanet(uint256 tokenId) returns()
func (_Planet *PlanetTransactorSession) UpGradePlanet(tokenId *big.Int) (*types.Transaction, error) {
	return _Planet.Contract.UpGradePlanet(&_Planet.TransactOpts, tokenId)
}

// PlanetAddTaxAmountIterator is returned from FilterAddTaxAmount and is used to iterate over the raw logs and unpacked data for AddTaxAmount events raised by the Planet contract.
type PlanetAddTaxAmountIterator struct {
	Event *PlanetAddTaxAmount // Event containing the contract specifics and raw log

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
func (it *PlanetAddTaxAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlanetAddTaxAmount)
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
		it.Event = new(PlanetAddTaxAmount)
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
func (it *PlanetAddTaxAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlanetAddTaxAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlanetAddTaxAmount represents a AddTaxAmount event raised by the Planet contract.
type PlanetAddTaxAmount struct {
	PlanetID *big.Int
	Player   common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddTaxAmount is a free log retrieval operation binding the contract event 0x9d3f41b946c5cf37743d55b073e1304471fdaad27065eca98021870c4d318d59.
//
// Solidity: event AddTaxAmount(uint256 indexed PlanetID, address indexed player, uint256 indexed amount)
func (_Planet *PlanetFilterer) FilterAddTaxAmount(opts *bind.FilterOpts, PlanetID []*big.Int, player []common.Address, amount []*big.Int) (*PlanetAddTaxAmountIterator, error) {

	var PlanetIDRule []interface{}
	for _, PlanetIDItem := range PlanetID {
		PlanetIDRule = append(PlanetIDRule, PlanetIDItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Planet.contract.FilterLogs(opts, "AddTaxAmount", PlanetIDRule, playerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &PlanetAddTaxAmountIterator{contract: _Planet.contract, event: "AddTaxAmount", logs: logs, sub: sub}, nil
}

// WatchAddTaxAmount is a free log subscription operation binding the contract event 0x9d3f41b946c5cf37743d55b073e1304471fdaad27065eca98021870c4d318d59.
//
// Solidity: event AddTaxAmount(uint256 indexed PlanetID, address indexed player, uint256 indexed amount)
func (_Planet *PlanetFilterer) WatchAddTaxAmount(opts *bind.WatchOpts, sink chan<- *PlanetAddTaxAmount, PlanetID []*big.Int, player []common.Address, amount []*big.Int) (event.Subscription, error) {

	var PlanetIDRule []interface{}
	for _, PlanetIDItem := range PlanetID {
		PlanetIDRule = append(PlanetIDRule, PlanetIDItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Planet.contract.WatchLogs(opts, "AddTaxAmount", PlanetIDRule, playerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlanetAddTaxAmount)
				if err := _Planet.contract.UnpackLog(event, "AddTaxAmount", log); err != nil {
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

// ParseAddTaxAmount is a log parse operation binding the contract event 0x9d3f41b946c5cf37743d55b073e1304471fdaad27065eca98021870c4d318d59.
//
// Solidity: event AddTaxAmount(uint256 indexed PlanetID, address indexed player, uint256 indexed amount)
func (_Planet *PlanetFilterer) ParseAddTaxAmount(log types.Log) (*PlanetAddTaxAmount, error) {
	event := new(PlanetAddTaxAmount)
	if err := _Planet.contract.UnpackLog(event, "AddTaxAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlanetApplyFederalPlanetIterator is returned from FilterApplyFederalPlanet and is used to iterate over the raw logs and unpacked data for ApplyFederalPlanet events raised by the Planet contract.
type PlanetApplyFederalPlanetIterator struct {
	Event *PlanetApplyFederalPlanet // Event containing the contract specifics and raw log

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
func (it *PlanetApplyFederalPlanetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlanetApplyFederalPlanet)
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
		it.Event = new(PlanetApplyFederalPlanet)
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
func (it *PlanetApplyFederalPlanetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlanetApplyFederalPlanetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlanetApplyFederalPlanet represents a ApplyFederalPlanet event raised by the Planet contract.
type PlanetApplyFederalPlanet struct {
	Player common.Address
	Amount *big.Int
	Tax    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApplyFederalPlanet is a free log retrieval operation binding the contract event 0xc6969413e2a222432feb274f4dcc7e8a6d77066246483fc02dea6ad21422b994.
//
// Solidity: event ApplyFederalPlanet(address indexed player, uint256 indexed amount, uint256 tax)
func (_Planet *PlanetFilterer) FilterApplyFederalPlanet(opts *bind.FilterOpts, player []common.Address, amount []*big.Int) (*PlanetApplyFederalPlanetIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Planet.contract.FilterLogs(opts, "ApplyFederalPlanet", playerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &PlanetApplyFederalPlanetIterator{contract: _Planet.contract, event: "ApplyFederalPlanet", logs: logs, sub: sub}, nil
}

// WatchApplyFederalPlanet is a free log subscription operation binding the contract event 0xc6969413e2a222432feb274f4dcc7e8a6d77066246483fc02dea6ad21422b994.
//
// Solidity: event ApplyFederalPlanet(address indexed player, uint256 indexed amount, uint256 tax)
func (_Planet *PlanetFilterer) WatchApplyFederalPlanet(opts *bind.WatchOpts, sink chan<- *PlanetApplyFederalPlanet, player []common.Address, amount []*big.Int) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Planet.contract.WatchLogs(opts, "ApplyFederalPlanet", playerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlanetApplyFederalPlanet)
				if err := _Planet.contract.UnpackLog(event, "ApplyFederalPlanet", log); err != nil {
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

// ParseApplyFederalPlanet is a log parse operation binding the contract event 0xc6969413e2a222432feb274f4dcc7e8a6d77066246483fc02dea6ad21422b994.
//
// Solidity: event ApplyFederalPlanet(address indexed player, uint256 indexed amount, uint256 tax)
func (_Planet *PlanetFilterer) ParseApplyFederalPlanet(log types.Log) (*PlanetApplyFederalPlanet, error) {
	event := new(PlanetApplyFederalPlanet)
	if err := _Planet.contract.UnpackLog(event, "ApplyFederalPlanet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlanetBondPlanetIterator is returned from FilterBondPlanet and is used to iterate over the raw logs and unpacked data for BondPlanet events raised by the Planet contract.
type PlanetBondPlanetIterator struct {
	Event *PlanetBondPlanet // Event containing the contract specifics and raw log

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
func (it *PlanetBondPlanetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlanetBondPlanet)
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
		it.Event = new(PlanetBondPlanet)
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
func (it *PlanetBondPlanetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlanetBondPlanetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlanetBondPlanet represents a BondPlanet event raised by the Planet contract.
type PlanetBondPlanet struct {
	Player  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBondPlanet is a free log retrieval operation binding the contract event 0x54ccf7ed3f9c5ae68462150a3fd9a0818be77ad3c1697988aa3b2b57a2ac99bd.
//
// Solidity: event BondPlanet(address indexed player, uint256 indexed tokenId)
func (_Planet *PlanetFilterer) FilterBondPlanet(opts *bind.FilterOpts, player []common.Address, tokenId []*big.Int) (*PlanetBondPlanetIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Planet.contract.FilterLogs(opts, "BondPlanet", playerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PlanetBondPlanetIterator{contract: _Planet.contract, event: "BondPlanet", logs: logs, sub: sub}, nil
}

// WatchBondPlanet is a free log subscription operation binding the contract event 0x54ccf7ed3f9c5ae68462150a3fd9a0818be77ad3c1697988aa3b2b57a2ac99bd.
//
// Solidity: event BondPlanet(address indexed player, uint256 indexed tokenId)
func (_Planet *PlanetFilterer) WatchBondPlanet(opts *bind.WatchOpts, sink chan<- *PlanetBondPlanet, player []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Planet.contract.WatchLogs(opts, "BondPlanet", playerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlanetBondPlanet)
				if err := _Planet.contract.UnpackLog(event, "BondPlanet", log); err != nil {
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

// ParseBondPlanet is a log parse operation binding the contract event 0x54ccf7ed3f9c5ae68462150a3fd9a0818be77ad3c1697988aa3b2b57a2ac99bd.
//
// Solidity: event BondPlanet(address indexed player, uint256 indexed tokenId)
func (_Planet *PlanetFilterer) ParseBondPlanet(log types.Log) (*PlanetBondPlanet, error) {
	event := new(PlanetBondPlanet)
	if err := _Planet.contract.UnpackLog(event, "BondPlanet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlanetCancelApplyIterator is returned from FilterCancelApply and is used to iterate over the raw logs and unpacked data for CancelApply events raised by the Planet contract.
type PlanetCancelApplyIterator struct {
	Event *PlanetCancelApply // Event containing the contract specifics and raw log

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
func (it *PlanetCancelApplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlanetCancelApply)
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
		it.Event = new(PlanetCancelApply)
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
func (it *PlanetCancelApplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlanetCancelApplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlanetCancelApply represents a CancelApply event raised by the Planet contract.
type PlanetCancelApply struct {
	Player common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelApply is a free log retrieval operation binding the contract event 0x4f73428f48569ba9bb5dd10e57b5d434590f19118fec434e2dacaf7c2e423151.
//
// Solidity: event CancelApply(address indexed player)
func (_Planet *PlanetFilterer) FilterCancelApply(opts *bind.FilterOpts, player []common.Address) (*PlanetCancelApplyIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _Planet.contract.FilterLogs(opts, "CancelApply", playerRule)
	if err != nil {
		return nil, err
	}
	return &PlanetCancelApplyIterator{contract: _Planet.contract, event: "CancelApply", logs: logs, sub: sub}, nil
}

// WatchCancelApply is a free log subscription operation binding the contract event 0x4f73428f48569ba9bb5dd10e57b5d434590f19118fec434e2dacaf7c2e423151.
//
// Solidity: event CancelApply(address indexed player)
func (_Planet *PlanetFilterer) WatchCancelApply(opts *bind.WatchOpts, sink chan<- *PlanetCancelApply, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _Planet.contract.WatchLogs(opts, "CancelApply", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlanetCancelApply)
				if err := _Planet.contract.UnpackLog(event, "CancelApply", log); err != nil {
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

// ParseCancelApply is a log parse operation binding the contract event 0x4f73428f48569ba9bb5dd10e57b5d434590f19118fec434e2dacaf7c2e423151.
//
// Solidity: event CancelApply(address indexed player)
func (_Planet *PlanetFilterer) ParseCancelApply(log types.Log) (*PlanetCancelApply, error) {
	event := new(PlanetCancelApply)
	if err := _Planet.contract.UnpackLog(event, "CancelApply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlanetNewPlanetIterator is returned from FilterNewPlanet and is used to iterate over the raw logs and unpacked data for NewPlanet events raised by the Planet contract.
type PlanetNewPlanetIterator struct {
	Event *PlanetNewPlanet // Event containing the contract specifics and raw log

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
func (it *PlanetNewPlanetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlanetNewPlanet)
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
		it.Event = new(PlanetNewPlanet)
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
func (it *PlanetNewPlanetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlanetNewPlanetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlanetNewPlanet represents a NewPlanet event raised by the Planet contract.
type PlanetNewPlanet struct {
	TokenId      *big.Int
	Types        *big.Int
	MotherPlanet *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewPlanet is a free log retrieval operation binding the contract event 0x5bd0809b02b94ebe0237c5fc7961b0732a88b9bb8af76cfdbaef67a888714ef0.
//
// Solidity: event NewPlanet(uint256 indexed tokenId, uint256 indexed types_, uint256 indexed motherPlanet)
func (_Planet *PlanetFilterer) FilterNewPlanet(opts *bind.FilterOpts, tokenId []*big.Int, types_ []*big.Int, motherPlanet []*big.Int) (*PlanetNewPlanetIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var types_Rule []interface{}
	for _, types_Item := range types_ {
		types_Rule = append(types_Rule, types_Item)
	}
	var motherPlanetRule []interface{}
	for _, motherPlanetItem := range motherPlanet {
		motherPlanetRule = append(motherPlanetRule, motherPlanetItem)
	}

	logs, sub, err := _Planet.contract.FilterLogs(opts, "NewPlanet", tokenIdRule, types_Rule, motherPlanetRule)
	if err != nil {
		return nil, err
	}
	return &PlanetNewPlanetIterator{contract: _Planet.contract, event: "NewPlanet", logs: logs, sub: sub}, nil
}

// WatchNewPlanet is a free log subscription operation binding the contract event 0x5bd0809b02b94ebe0237c5fc7961b0732a88b9bb8af76cfdbaef67a888714ef0.
//
// Solidity: event NewPlanet(uint256 indexed tokenId, uint256 indexed types_, uint256 indexed motherPlanet)
func (_Planet *PlanetFilterer) WatchNewPlanet(opts *bind.WatchOpts, sink chan<- *PlanetNewPlanet, tokenId []*big.Int, types_ []*big.Int, motherPlanet []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var types_Rule []interface{}
	for _, types_Item := range types_ {
		types_Rule = append(types_Rule, types_Item)
	}
	var motherPlanetRule []interface{}
	for _, motherPlanetItem := range motherPlanet {
		motherPlanetRule = append(motherPlanetRule, motherPlanetItem)
	}

	logs, sub, err := _Planet.contract.WatchLogs(opts, "NewPlanet", tokenIdRule, types_Rule, motherPlanetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlanetNewPlanet)
				if err := _Planet.contract.UnpackLog(event, "NewPlanet", log); err != nil {
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

// ParseNewPlanet is a log parse operation binding the contract event 0x5bd0809b02b94ebe0237c5fc7961b0732a88b9bb8af76cfdbaef67a888714ef0.
//
// Solidity: event NewPlanet(uint256 indexed tokenId, uint256 indexed types_, uint256 indexed motherPlanet)
func (_Planet *PlanetFilterer) ParseNewPlanet(log types.Log) (*PlanetNewPlanet, error) {
	event := new(PlanetNewPlanet)
	if err := _Planet.contract.UnpackLog(event, "NewPlanet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlanetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Planet contract.
type PlanetOwnershipTransferredIterator struct {
	Event *PlanetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PlanetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlanetOwnershipTransferred)
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
		it.Event = new(PlanetOwnershipTransferred)
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
func (it *PlanetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlanetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlanetOwnershipTransferred represents a OwnershipTransferred event raised by the Planet contract.
type PlanetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Planet *PlanetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PlanetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Planet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlanetOwnershipTransferredIterator{contract: _Planet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Planet *PlanetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PlanetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Planet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlanetOwnershipTransferred)
				if err := _Planet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Planet *PlanetFilterer) ParseOwnershipTransferred(log types.Log) (*PlanetOwnershipTransferred, error) {
	event := new(PlanetOwnershipTransferred)
	if err := _Planet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlanetUpGradePlanetIterator is returned from FilterUpGradePlanet and is used to iterate over the raw logs and unpacked data for UpGradePlanet events raised by the Planet contract.
type PlanetUpGradePlanetIterator struct {
	Event *PlanetUpGradePlanet // Event containing the contract specifics and raw log

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
func (it *PlanetUpGradePlanetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlanetUpGradePlanet)
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
		it.Event = new(PlanetUpGradePlanet)
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
func (it *PlanetUpGradePlanetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlanetUpGradePlanetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlanetUpGradePlanet represents a UpGradePlanet event raised by the Planet contract.
type PlanetUpGradePlanet struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpGradePlanet is a free log retrieval operation binding the contract event 0xe0d98b5429c829ebf62be16ab2d63c076c373f8921c854a5a2c94922e042693c.
//
// Solidity: event UpGradePlanet(uint256 indexed tokenId)
func (_Planet *PlanetFilterer) FilterUpGradePlanet(opts *bind.FilterOpts, tokenId []*big.Int) (*PlanetUpGradePlanetIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Planet.contract.FilterLogs(opts, "UpGradePlanet", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PlanetUpGradePlanetIterator{contract: _Planet.contract, event: "UpGradePlanet", logs: logs, sub: sub}, nil
}

// WatchUpGradePlanet is a free log subscription operation binding the contract event 0xe0d98b5429c829ebf62be16ab2d63c076c373f8921c854a5a2c94922e042693c.
//
// Solidity: event UpGradePlanet(uint256 indexed tokenId)
func (_Planet *PlanetFilterer) WatchUpGradePlanet(opts *bind.WatchOpts, sink chan<- *PlanetUpGradePlanet, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Planet.contract.WatchLogs(opts, "UpGradePlanet", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlanetUpGradePlanet)
				if err := _Planet.contract.UnpackLog(event, "UpGradePlanet", log); err != nil {
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

// ParseUpGradePlanet is a log parse operation binding the contract event 0xe0d98b5429c829ebf62be16ab2d63c076c373f8921c854a5a2c94922e042693c.
//
// Solidity: event UpGradePlanet(uint256 indexed tokenId)
func (_Planet *PlanetFilterer) ParseUpGradePlanet(log types.Log) (*PlanetUpGradePlanet, error) {
	event := new(PlanetUpGradePlanet)
	if err := _Planet.contract.UnpackLog(event, "UpGradePlanet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlanetUpGradeTechnologyIterator is returned from FilterUpGradeTechnology and is used to iterate over the raw logs and unpacked data for UpGradeTechnology events raised by the Planet contract.
type PlanetUpGradeTechnologyIterator struct {
	Event *PlanetUpGradeTechnology // Event containing the contract specifics and raw log

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
func (it *PlanetUpGradeTechnologyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlanetUpGradeTechnology)
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
		it.Event = new(PlanetUpGradeTechnology)
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
func (it *PlanetUpGradeTechnologyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlanetUpGradeTechnologyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlanetUpGradeTechnology represents a UpGradeTechnology event raised by the Planet contract.
type PlanetUpGradeTechnology struct {
	TokenId *big.Int
	TecNum  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpGradeTechnology is a free log retrieval operation binding the contract event 0x277c6abac3e1dcaa91126c388145e7f6d6c64dd2218b4c0a8d2a98cf8bad186a.
//
// Solidity: event UpGradeTechnology(uint256 indexed tokenId, uint256 indexed tecNum)
func (_Planet *PlanetFilterer) FilterUpGradeTechnology(opts *bind.FilterOpts, tokenId []*big.Int, tecNum []*big.Int) (*PlanetUpGradeTechnologyIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var tecNumRule []interface{}
	for _, tecNumItem := range tecNum {
		tecNumRule = append(tecNumRule, tecNumItem)
	}

	logs, sub, err := _Planet.contract.FilterLogs(opts, "UpGradeTechnology", tokenIdRule, tecNumRule)
	if err != nil {
		return nil, err
	}
	return &PlanetUpGradeTechnologyIterator{contract: _Planet.contract, event: "UpGradeTechnology", logs: logs, sub: sub}, nil
}

// WatchUpGradeTechnology is a free log subscription operation binding the contract event 0x277c6abac3e1dcaa91126c388145e7f6d6c64dd2218b4c0a8d2a98cf8bad186a.
//
// Solidity: event UpGradeTechnology(uint256 indexed tokenId, uint256 indexed tecNum)
func (_Planet *PlanetFilterer) WatchUpGradeTechnology(opts *bind.WatchOpts, sink chan<- *PlanetUpGradeTechnology, tokenId []*big.Int, tecNum []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var tecNumRule []interface{}
	for _, tecNumItem := range tecNum {
		tecNumRule = append(tecNumRule, tecNumItem)
	}

	logs, sub, err := _Planet.contract.WatchLogs(opts, "UpGradeTechnology", tokenIdRule, tecNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlanetUpGradeTechnology)
				if err := _Planet.contract.UnpackLog(event, "UpGradeTechnology", log); err != nil {
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

// ParseUpGradeTechnology is a log parse operation binding the contract event 0x277c6abac3e1dcaa91126c388145e7f6d6c64dd2218b4c0a8d2a98cf8bad186a.
//
// Solidity: event UpGradeTechnology(uint256 indexed tokenId, uint256 indexed tecNum)
func (_Planet *PlanetFilterer) ParseUpGradeTechnology(log types.Log) (*PlanetUpGradeTechnology, error) {
	event := new(PlanetUpGradeTechnology)
	if err := _Planet.contract.UnpackLog(event, "UpGradeTechnology", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
