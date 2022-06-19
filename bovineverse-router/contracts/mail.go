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

// MailABI is the input ABI used to generate the binding from.
const MailABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player2\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"types\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relationshipID\",\"type\":\"uint256\"}],\"name\":\"BondRelationship\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ClaimMail\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"relationshipID\",\"type\":\"uint256\"}],\"name\":\"UnBondRelationship\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BVG\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BVT\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avatar\",\"outputs\":[{\"internalType\":\"contractIProfilePhoto\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"banker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"types\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"bond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"box\",\"outputs\":[{\"internalType\":\"contractIBOX\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bvgClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bvtClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cattle\",\"outputs\":[{\"internalType\":\"contractICOW\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"rewardType\",\"type\":\"uint8[]\"},{\"internalType\":\"uint8[]\",\"name\":\"rewardId\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewardAmount\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"isTax\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"claimMail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compound\",\"outputs\":[{\"internalType\":\"contractICompound\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"idClaim\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"item\",\"outputs\":[{\"internalType\":\"contractICattle1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mating\",\"outputs\":[{\"internalType\":\"contractIMating\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"milk\",\"outputs\":[{\"internalType\":\"contractIMilk\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"planet\",\"outputs\":[{\"internalType\":\"contractIPlanet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"relationIdentify\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"relationTypes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setBanker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"setCattle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCompound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setMating\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setMilk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"setPlanet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setProfilePhoto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"name\":\"setStable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"BVT_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BVG_\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stable\",\"outputs\":[{\"internalType\":\"contractIStable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"relationID\",\"type\":\"uint256\"}],\"name\":\"unBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Mail is an auto generated Go binding around an Ethereum contract.
type Mail struct {
	MailCaller     // Read-only binding to the contract
	MailTransactor // Write-only binding to the contract
	MailFilterer   // Log filterer for contract events
}

// MailCaller is an auto generated read-only Go binding around an Ethereum contract.
type MailCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MailTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MailFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MailSession struct {
	Contract     *Mail             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MailCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MailCallerSession struct {
	Contract *MailCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MailTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MailTransactorSession struct {
	Contract     *MailTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MailRaw is an auto generated low-level Go binding around an Ethereum contract.
type MailRaw struct {
	Contract *Mail // Generic contract binding to access the raw methods on
}

// MailCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MailCallerRaw struct {
	Contract *MailCaller // Generic read-only contract binding to access the raw methods on
}

// MailTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MailTransactorRaw struct {
	Contract *MailTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMail creates a new instance of Mail, bound to a specific deployed contract.
func NewMail(address common.Address, backend bind.ContractBackend) (*Mail, error) {
	contract, err := bindMail(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mail{MailCaller: MailCaller{contract: contract}, MailTransactor: MailTransactor{contract: contract}, MailFilterer: MailFilterer{contract: contract}}, nil
}

// NewMailCaller creates a new read-only instance of Mail, bound to a specific deployed contract.
func NewMailCaller(address common.Address, caller bind.ContractCaller) (*MailCaller, error) {
	contract, err := bindMail(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MailCaller{contract: contract}, nil
}

// NewMailTransactor creates a new write-only instance of Mail, bound to a specific deployed contract.
func NewMailTransactor(address common.Address, transactor bind.ContractTransactor) (*MailTransactor, error) {
	contract, err := bindMail(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MailTransactor{contract: contract}, nil
}

// NewMailFilterer creates a new log filterer instance of Mail, bound to a specific deployed contract.
func NewMailFilterer(address common.Address, filterer bind.ContractFilterer) (*MailFilterer, error) {
	contract, err := bindMail(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MailFilterer{contract: contract}, nil
}

// bindMail binds a generic wrapper to an already deployed contract.
func bindMail(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MailABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mail *MailRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mail.Contract.MailCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mail *MailRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mail.Contract.MailTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mail *MailRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mail.Contract.MailTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mail *MailCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mail.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mail *MailTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mail.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mail *MailTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mail.Contract.contract.Transact(opts, method, params...)
}

// BVG is a free data retrieval call binding the contract method 0x87464389.
//
// Solidity: function BVG() view returns(address)
func (_Mail *MailCaller) BVG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "BVG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BVG is a free data retrieval call binding the contract method 0x87464389.
//
// Solidity: function BVG() view returns(address)
func (_Mail *MailSession) BVG() (common.Address, error) {
	return _Mail.Contract.BVG(&_Mail.CallOpts)
}

// BVG is a free data retrieval call binding the contract method 0x87464389.
//
// Solidity: function BVG() view returns(address)
func (_Mail *MailCallerSession) BVG() (common.Address, error) {
	return _Mail.Contract.BVG(&_Mail.CallOpts)
}

// BVT is a free data retrieval call binding the contract method 0x5f24b4a4.
//
// Solidity: function BVT() view returns(address)
func (_Mail *MailCaller) BVT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "BVT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BVT is a free data retrieval call binding the contract method 0x5f24b4a4.
//
// Solidity: function BVT() view returns(address)
func (_Mail *MailSession) BVT() (common.Address, error) {
	return _Mail.Contract.BVT(&_Mail.CallOpts)
}

// BVT is a free data retrieval call binding the contract method 0x5f24b4a4.
//
// Solidity: function BVT() view returns(address)
func (_Mail *MailCallerSession) BVT() (common.Address, error) {
	return _Mail.Contract.BVT(&_Mail.CallOpts)
}

// Avatar is a free data retrieval call binding the contract method 0x5aef7de6.
//
// Solidity: function avatar() view returns(address)
func (_Mail *MailCaller) Avatar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "avatar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avatar is a free data retrieval call binding the contract method 0x5aef7de6.
//
// Solidity: function avatar() view returns(address)
func (_Mail *MailSession) Avatar() (common.Address, error) {
	return _Mail.Contract.Avatar(&_Mail.CallOpts)
}

// Avatar is a free data retrieval call binding the contract method 0x5aef7de6.
//
// Solidity: function avatar() view returns(address)
func (_Mail *MailCallerSession) Avatar() (common.Address, error) {
	return _Mail.Contract.Avatar(&_Mail.CallOpts)
}

// Banker is a free data retrieval call binding the contract method 0x0ab9db5b.
//
// Solidity: function banker() view returns(address)
func (_Mail *MailCaller) Banker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "banker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Banker is a free data retrieval call binding the contract method 0x0ab9db5b.
//
// Solidity: function banker() view returns(address)
func (_Mail *MailSession) Banker() (common.Address, error) {
	return _Mail.Contract.Banker(&_Mail.CallOpts)
}

// Banker is a free data retrieval call binding the contract method 0x0ab9db5b.
//
// Solidity: function banker() view returns(address)
func (_Mail *MailCallerSession) Banker() (common.Address, error) {
	return _Mail.Contract.Banker(&_Mail.CallOpts)
}

// Box is a free data retrieval call binding the contract method 0x754215a1.
//
// Solidity: function box() view returns(address)
func (_Mail *MailCaller) Box(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "box")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Box is a free data retrieval call binding the contract method 0x754215a1.
//
// Solidity: function box() view returns(address)
func (_Mail *MailSession) Box() (common.Address, error) {
	return _Mail.Contract.Box(&_Mail.CallOpts)
}

// Box is a free data retrieval call binding the contract method 0x754215a1.
//
// Solidity: function box() view returns(address)
func (_Mail *MailCallerSession) Box() (common.Address, error) {
	return _Mail.Contract.Box(&_Mail.CallOpts)
}

// BvgClaimed is a free data retrieval call binding the contract method 0x1e9bcc0e.
//
// Solidity: function bvgClaimed(address ) view returns(uint256)
func (_Mail *MailCaller) BvgClaimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "bvgClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BvgClaimed is a free data retrieval call binding the contract method 0x1e9bcc0e.
//
// Solidity: function bvgClaimed(address ) view returns(uint256)
func (_Mail *MailSession) BvgClaimed(arg0 common.Address) (*big.Int, error) {
	return _Mail.Contract.BvgClaimed(&_Mail.CallOpts, arg0)
}

// BvgClaimed is a free data retrieval call binding the contract method 0x1e9bcc0e.
//
// Solidity: function bvgClaimed(address ) view returns(uint256)
func (_Mail *MailCallerSession) BvgClaimed(arg0 common.Address) (*big.Int, error) {
	return _Mail.Contract.BvgClaimed(&_Mail.CallOpts, arg0)
}

// BvtClaimed is a free data retrieval call binding the contract method 0xc542a90b.
//
// Solidity: function bvtClaimed(address ) view returns(uint256)
func (_Mail *MailCaller) BvtClaimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "bvtClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BvtClaimed is a free data retrieval call binding the contract method 0xc542a90b.
//
// Solidity: function bvtClaimed(address ) view returns(uint256)
func (_Mail *MailSession) BvtClaimed(arg0 common.Address) (*big.Int, error) {
	return _Mail.Contract.BvtClaimed(&_Mail.CallOpts, arg0)
}

// BvtClaimed is a free data retrieval call binding the contract method 0xc542a90b.
//
// Solidity: function bvtClaimed(address ) view returns(uint256)
func (_Mail *MailCallerSession) BvtClaimed(arg0 common.Address) (*big.Int, error) {
	return _Mail.Contract.BvtClaimed(&_Mail.CallOpts, arg0)
}

// Cattle is a free data retrieval call binding the contract method 0x01c211c7.
//
// Solidity: function cattle() view returns(address)
func (_Mail *MailCaller) Cattle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "cattle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cattle is a free data retrieval call binding the contract method 0x01c211c7.
//
// Solidity: function cattle() view returns(address)
func (_Mail *MailSession) Cattle() (common.Address, error) {
	return _Mail.Contract.Cattle(&_Mail.CallOpts)
}

// Cattle is a free data retrieval call binding the contract method 0x01c211c7.
//
// Solidity: function cattle() view returns(address)
func (_Mail *MailCallerSession) Cattle() (common.Address, error) {
	return _Mail.Contract.Cattle(&_Mail.CallOpts)
}

// Compound is a free data retrieval call binding the contract method 0xf69e2046.
//
// Solidity: function compound() view returns(address)
func (_Mail *MailCaller) Compound(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "compound")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Compound is a free data retrieval call binding the contract method 0xf69e2046.
//
// Solidity: function compound() view returns(address)
func (_Mail *MailSession) Compound() (common.Address, error) {
	return _Mail.Contract.Compound(&_Mail.CallOpts)
}

// Compound is a free data retrieval call binding the contract method 0xf69e2046.
//
// Solidity: function compound() view returns(address)
func (_Mail *MailCallerSession) Compound() (common.Address, error) {
	return _Mail.Contract.Compound(&_Mail.CallOpts)
}

// IdClaim is a free data retrieval call binding the contract method 0x53b9008f.
//
// Solidity: function idClaim(uint256 ) view returns(address)
func (_Mail *MailCaller) IdClaim(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "idClaim", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdClaim is a free data retrieval call binding the contract method 0x53b9008f.
//
// Solidity: function idClaim(uint256 ) view returns(address)
func (_Mail *MailSession) IdClaim(arg0 *big.Int) (common.Address, error) {
	return _Mail.Contract.IdClaim(&_Mail.CallOpts, arg0)
}

// IdClaim is a free data retrieval call binding the contract method 0x53b9008f.
//
// Solidity: function idClaim(uint256 ) view returns(address)
func (_Mail *MailCallerSession) IdClaim(arg0 *big.Int) (common.Address, error) {
	return _Mail.Contract.IdClaim(&_Mail.CallOpts, arg0)
}

// Item is a free data retrieval call binding the contract method 0xf2a4a82e.
//
// Solidity: function item() view returns(address)
func (_Mail *MailCaller) Item(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "item")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Item is a free data retrieval call binding the contract method 0xf2a4a82e.
//
// Solidity: function item() view returns(address)
func (_Mail *MailSession) Item() (common.Address, error) {
	return _Mail.Contract.Item(&_Mail.CallOpts)
}

// Item is a free data retrieval call binding the contract method 0xf2a4a82e.
//
// Solidity: function item() view returns(address)
func (_Mail *MailCallerSession) Item() (common.Address, error) {
	return _Mail.Contract.Item(&_Mail.CallOpts)
}

// Mating is a free data retrieval call binding the contract method 0xb4dac333.
//
// Solidity: function mating() view returns(address)
func (_Mail *MailCaller) Mating(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "mating")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mating is a free data retrieval call binding the contract method 0xb4dac333.
//
// Solidity: function mating() view returns(address)
func (_Mail *MailSession) Mating() (common.Address, error) {
	return _Mail.Contract.Mating(&_Mail.CallOpts)
}

// Mating is a free data retrieval call binding the contract method 0xb4dac333.
//
// Solidity: function mating() view returns(address)
func (_Mail *MailCallerSession) Mating() (common.Address, error) {
	return _Mail.Contract.Mating(&_Mail.CallOpts)
}

// Milk is a free data retrieval call binding the contract method 0xdbf8b168.
//
// Solidity: function milk() view returns(address)
func (_Mail *MailCaller) Milk(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "milk")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Milk is a free data retrieval call binding the contract method 0xdbf8b168.
//
// Solidity: function milk() view returns(address)
func (_Mail *MailSession) Milk() (common.Address, error) {
	return _Mail.Contract.Milk(&_Mail.CallOpts)
}

// Milk is a free data retrieval call binding the contract method 0xdbf8b168.
//
// Solidity: function milk() view returns(address)
func (_Mail *MailCallerSession) Milk() (common.Address, error) {
	return _Mail.Contract.Milk(&_Mail.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mail *MailCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mail *MailSession) Owner() (common.Address, error) {
	return _Mail.Contract.Owner(&_Mail.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mail *MailCallerSession) Owner() (common.Address, error) {
	return _Mail.Contract.Owner(&_Mail.CallOpts)
}

// Planet is a free data retrieval call binding the contract method 0x8443bea6.
//
// Solidity: function planet() view returns(address)
func (_Mail *MailCaller) Planet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "planet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Planet is a free data retrieval call binding the contract method 0x8443bea6.
//
// Solidity: function planet() view returns(address)
func (_Mail *MailSession) Planet() (common.Address, error) {
	return _Mail.Contract.Planet(&_Mail.CallOpts)
}

// Planet is a free data retrieval call binding the contract method 0x8443bea6.
//
// Solidity: function planet() view returns(address)
func (_Mail *MailCallerSession) Planet() (common.Address, error) {
	return _Mail.Contract.Planet(&_Mail.CallOpts)
}

// RelationIdentify is a free data retrieval call binding the contract method 0x5d3bea2c.
//
// Solidity: function relationIdentify(uint256 , uint256 ) view returns(address)
func (_Mail *MailCaller) RelationIdentify(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "relationIdentify", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelationIdentify is a free data retrieval call binding the contract method 0x5d3bea2c.
//
// Solidity: function relationIdentify(uint256 , uint256 ) view returns(address)
func (_Mail *MailSession) RelationIdentify(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Mail.Contract.RelationIdentify(&_Mail.CallOpts, arg0, arg1)
}

// RelationIdentify is a free data retrieval call binding the contract method 0x5d3bea2c.
//
// Solidity: function relationIdentify(uint256 , uint256 ) view returns(address)
func (_Mail *MailCallerSession) RelationIdentify(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Mail.Contract.RelationIdentify(&_Mail.CallOpts, arg0, arg1)
}

// RelationTypes is a free data retrieval call binding the contract method 0x8be031f4.
//
// Solidity: function relationTypes(address , address ) view returns(uint256)
func (_Mail *MailCaller) RelationTypes(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "relationTypes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelationTypes is a free data retrieval call binding the contract method 0x8be031f4.
//
// Solidity: function relationTypes(address , address ) view returns(uint256)
func (_Mail *MailSession) RelationTypes(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Mail.Contract.RelationTypes(&_Mail.CallOpts, arg0, arg1)
}

// RelationTypes is a free data retrieval call binding the contract method 0x8be031f4.
//
// Solidity: function relationTypes(address , address ) view returns(uint256)
func (_Mail *MailCallerSession) RelationTypes(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Mail.Contract.RelationTypes(&_Mail.CallOpts, arg0, arg1)
}

// Rid is a free data retrieval call binding the contract method 0xa2f1b44e.
//
// Solidity: function rid() view returns(uint256)
func (_Mail *MailCaller) Rid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "rid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rid is a free data retrieval call binding the contract method 0xa2f1b44e.
//
// Solidity: function rid() view returns(uint256)
func (_Mail *MailSession) Rid() (*big.Int, error) {
	return _Mail.Contract.Rid(&_Mail.CallOpts)
}

// Rid is a free data retrieval call binding the contract method 0xa2f1b44e.
//
// Solidity: function rid() view returns(uint256)
func (_Mail *MailCallerSession) Rid() (*big.Int, error) {
	return _Mail.Contract.Rid(&_Mail.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(address)
func (_Mail *MailCaller) Stable(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mail.contract.Call(opts, &out, "stable")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(address)
func (_Mail *MailSession) Stable() (common.Address, error) {
	return _Mail.Contract.Stable(&_Mail.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(address)
func (_Mail *MailCallerSession) Stable() (common.Address, error) {
	return _Mail.Contract.Stable(&_Mail.CallOpts)
}

// Bond is a paid mutator transaction binding the contract method 0x62d34ff7.
//
// Solidity: function bond(address addr1, address addr2, uint256 types, bytes32 r, bytes32 s, uint8 v) returns()
func (_Mail *MailTransactor) Bond(opts *bind.TransactOpts, addr1 common.Address, addr2 common.Address, types *big.Int, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "bond", addr1, addr2, types, r, s, v)
}

// Bond is a paid mutator transaction binding the contract method 0x62d34ff7.
//
// Solidity: function bond(address addr1, address addr2, uint256 types, bytes32 r, bytes32 s, uint8 v) returns()
func (_Mail *MailSession) Bond(addr1 common.Address, addr2 common.Address, types *big.Int, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Mail.Contract.Bond(&_Mail.TransactOpts, addr1, addr2, types, r, s, v)
}

// Bond is a paid mutator transaction binding the contract method 0x62d34ff7.
//
// Solidity: function bond(address addr1, address addr2, uint256 types, bytes32 r, bytes32 s, uint8 v) returns()
func (_Mail *MailTransactorSession) Bond(addr1 common.Address, addr2 common.Address, types *big.Int, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Mail.Contract.Bond(&_Mail.TransactOpts, addr1, addr2, types, r, s, v)
}

// ClaimMail is a paid mutator transaction binding the contract method 0x1ef84039.
//
// Solidity: function claimMail(uint8[] rewardType, uint8[] rewardId, uint256[] rewardAmount, bool isTax, uint256 id, bytes32 r, bytes32 s, uint8 v) returns()
func (_Mail *MailTransactor) ClaimMail(opts *bind.TransactOpts, rewardType []uint8, rewardId []uint8, rewardAmount []*big.Int, isTax bool, id *big.Int, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "claimMail", rewardType, rewardId, rewardAmount, isTax, id, r, s, v)
}

// ClaimMail is a paid mutator transaction binding the contract method 0x1ef84039.
//
// Solidity: function claimMail(uint8[] rewardType, uint8[] rewardId, uint256[] rewardAmount, bool isTax, uint256 id, bytes32 r, bytes32 s, uint8 v) returns()
func (_Mail *MailSession) ClaimMail(rewardType []uint8, rewardId []uint8, rewardAmount []*big.Int, isTax bool, id *big.Int, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Mail.Contract.ClaimMail(&_Mail.TransactOpts, rewardType, rewardId, rewardAmount, isTax, id, r, s, v)
}

// ClaimMail is a paid mutator transaction binding the contract method 0x1ef84039.
//
// Solidity: function claimMail(uint8[] rewardType, uint8[] rewardId, uint256[] rewardAmount, bool isTax, uint256 id, bytes32 r, bytes32 s, uint8 v) returns()
func (_Mail *MailTransactorSession) ClaimMail(rewardType []uint8, rewardId []uint8, rewardAmount []*big.Int, isTax bool, id *big.Int, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Mail.Contract.ClaimMail(&_Mail.TransactOpts, rewardType, rewardId, rewardAmount, isTax, id, r, s, v)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mail *MailTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mail *MailSession) Initialize() (*types.Transaction, error) {
	return _Mail.Contract.Initialize(&_Mail.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mail *MailTransactorSession) Initialize() (*types.Transaction, error) {
	return _Mail.Contract.Initialize(&_Mail.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mail *MailTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mail *MailSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mail.Contract.RenounceOwnership(&_Mail.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mail *MailTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mail.Contract.RenounceOwnership(&_Mail.TransactOpts)
}

// SetBanker is a paid mutator transaction binding the contract method 0xf1ff732b.
//
// Solidity: function setBanker(address addr) returns()
func (_Mail *MailTransactor) SetBanker(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "setBanker", addr)
}

// SetBanker is a paid mutator transaction binding the contract method 0xf1ff732b.
//
// Solidity: function setBanker(address addr) returns()
func (_Mail *MailSession) SetBanker(addr common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetBanker(&_Mail.TransactOpts, addr)
}

// SetBanker is a paid mutator transaction binding the contract method 0xf1ff732b.
//
// Solidity: function setBanker(address addr) returns()
func (_Mail *MailTransactorSession) SetBanker(addr common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetBanker(&_Mail.TransactOpts, addr)
}

// SetBox is a paid mutator transaction binding the contract method 0x537627ab.
//
// Solidity: function setBox(address addr) returns()
func (_Mail *MailTransactor) SetBox(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "setBox", addr)
}

// SetBox is a paid mutator transaction binding the contract method 0x537627ab.
//
// Solidity: function setBox(address addr) returns()
func (_Mail *MailSession) SetBox(addr common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetBox(&_Mail.TransactOpts, addr)
}

// SetBox is a paid mutator transaction binding the contract method 0x537627ab.
//
// Solidity: function setBox(address addr) returns()
func (_Mail *MailTransactorSession) SetBox(addr common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetBox(&_Mail.TransactOpts, addr)
}

// SetCattle is a paid mutator transaction binding the contract method 0x0de307fe.
//
// Solidity: function setCattle(address addr_) returns()
func (_Mail *MailTransactor) SetCattle(opts *bind.TransactOpts, addr_ common.Address) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "setCattle", addr_)
}

// SetCattle is a paid mutator transaction binding the contract method 0x0de307fe.
//
// Solidity: function setCattle(address addr_) returns()
func (_Mail *MailSession) SetCattle(addr_ common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetCattle(&_Mail.TransactOpts, addr_)
}

// SetCattle is a paid mutator transaction binding the contract method 0x0de307fe.
//
// Solidity: function setCattle(address addr_) returns()
func (_Mail *MailTransactorSession) SetCattle(addr_ common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetCattle(&_Mail.TransactOpts, addr_)
}

// SetCompound is a paid mutator transaction binding the contract method 0x2be682d4.
//
// Solidity: function setCompound(address addr) returns()
func (_Mail *MailTransactor) SetCompound(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "setCompound", addr)
}

// SetCompound is a paid mutator transaction binding the contract method 0x2be682d4.
//
// Solidity: function setCompound(address addr) returns()
func (_Mail *MailSession) SetCompound(addr common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetCompound(&_Mail.TransactOpts, addr)
}

// SetCompound is a paid mutator transaction binding the contract method 0x2be682d4.
//
// Solidity: function setCompound(address addr) returns()
func (_Mail *MailTransactorSession) SetCompound(addr common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetCompound(&_Mail.TransactOpts, addr)
}

// SetItem is a paid mutator transaction binding the contract method 0x165ed276.
//
// Solidity: function setItem(address addr_) returns()
func (_Mail *MailTransactor) SetItem(opts *bind.TransactOpts, addr_ common.Address) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "setItem", addr_)
}

// SetItem is a paid mutator transaction binding the contract method 0x165ed276.
//
// Solidity: function setItem(address addr_) returns()
func (_Mail *MailSession) SetItem(addr_ common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetItem(&_Mail.TransactOpts, addr_)
}

// SetItem is a paid mutator transaction binding the contract method 0x165ed276.
//
// Solidity: function setItem(address addr_) returns()
func (_Mail *MailTransactorSession) SetItem(addr_ common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetItem(&_Mail.TransactOpts, addr_)
}

// SetMating is a paid mutator transaction binding the contract method 0x9fd0e138.
//
// Solidity: function setMating(address addr) returns()
func (_Mail *MailTransactor) SetMating(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "setMating", addr)
}

// SetMating is a paid mutator transaction binding the contract method 0x9fd0e138.
//
// Solidity: function setMating(address addr) returns()
func (_Mail *MailSession) SetMating(addr common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetMating(&_Mail.TransactOpts, addr)
}

// SetMating is a paid mutator transaction binding the contract method 0x9fd0e138.
//
// Solidity: function setMating(address addr) returns()
func (_Mail *MailTransactorSession) SetMating(addr common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetMating(&_Mail.TransactOpts, addr)
}

// SetMilk is a paid mutator transaction binding the contract method 0x110a2ded.
//
// Solidity: function setMilk(address addr) returns()
func (_Mail *MailTransactor) SetMilk(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "setMilk", addr)
}

// SetMilk is a paid mutator transaction binding the contract method 0x110a2ded.
//
// Solidity: function setMilk(address addr) returns()
func (_Mail *MailSession) SetMilk(addr common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetMilk(&_Mail.TransactOpts, addr)
}

// SetMilk is a paid mutator transaction binding the contract method 0x110a2ded.
//
// Solidity: function setMilk(address addr) returns()
func (_Mail *MailTransactorSession) SetMilk(addr common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetMilk(&_Mail.TransactOpts, addr)
}

// SetPlanet is a paid mutator transaction binding the contract method 0xad911387.
//
// Solidity: function setPlanet(address addr_) returns()
func (_Mail *MailTransactor) SetPlanet(opts *bind.TransactOpts, addr_ common.Address) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "setPlanet", addr_)
}

// SetPlanet is a paid mutator transaction binding the contract method 0xad911387.
//
// Solidity: function setPlanet(address addr_) returns()
func (_Mail *MailSession) SetPlanet(addr_ common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetPlanet(&_Mail.TransactOpts, addr_)
}

// SetPlanet is a paid mutator transaction binding the contract method 0xad911387.
//
// Solidity: function setPlanet(address addr_) returns()
func (_Mail *MailTransactorSession) SetPlanet(addr_ common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetPlanet(&_Mail.TransactOpts, addr_)
}

// SetProfilePhoto is a paid mutator transaction binding the contract method 0x61a91290.
//
// Solidity: function setProfilePhoto(address addr) returns()
func (_Mail *MailTransactor) SetProfilePhoto(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "setProfilePhoto", addr)
}

// SetProfilePhoto is a paid mutator transaction binding the contract method 0x61a91290.
//
// Solidity: function setProfilePhoto(address addr) returns()
func (_Mail *MailSession) SetProfilePhoto(addr common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetProfilePhoto(&_Mail.TransactOpts, addr)
}

// SetProfilePhoto is a paid mutator transaction binding the contract method 0x61a91290.
//
// Solidity: function setProfilePhoto(address addr) returns()
func (_Mail *MailTransactorSession) SetProfilePhoto(addr common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetProfilePhoto(&_Mail.TransactOpts, addr)
}

// SetStable is a paid mutator transaction binding the contract method 0xdba71376.
//
// Solidity: function setStable(address addr_) returns()
func (_Mail *MailTransactor) SetStable(opts *bind.TransactOpts, addr_ common.Address) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "setStable", addr_)
}

// SetStable is a paid mutator transaction binding the contract method 0xdba71376.
//
// Solidity: function setStable(address addr_) returns()
func (_Mail *MailSession) SetStable(addr_ common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetStable(&_Mail.TransactOpts, addr_)
}

// SetStable is a paid mutator transaction binding the contract method 0xdba71376.
//
// Solidity: function setStable(address addr_) returns()
func (_Mail *MailTransactorSession) SetStable(addr_ common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetStable(&_Mail.TransactOpts, addr_)
}

// SetToken is a paid mutator transaction binding the contract method 0x1da26a8b.
//
// Solidity: function setToken(address BVT_, address BVG_) returns()
func (_Mail *MailTransactor) SetToken(opts *bind.TransactOpts, BVT_ common.Address, BVG_ common.Address) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "setToken", BVT_, BVG_)
}

// SetToken is a paid mutator transaction binding the contract method 0x1da26a8b.
//
// Solidity: function setToken(address BVT_, address BVG_) returns()
func (_Mail *MailSession) SetToken(BVT_ common.Address, BVG_ common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetToken(&_Mail.TransactOpts, BVT_, BVG_)
}

// SetToken is a paid mutator transaction binding the contract method 0x1da26a8b.
//
// Solidity: function setToken(address BVT_, address BVG_) returns()
func (_Mail *MailTransactorSession) SetToken(BVT_ common.Address, BVG_ common.Address) (*types.Transaction, error) {
	return _Mail.Contract.SetToken(&_Mail.TransactOpts, BVT_, BVG_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mail *MailTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mail *MailSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mail.Contract.TransferOwnership(&_Mail.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mail *MailTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mail.Contract.TransferOwnership(&_Mail.TransactOpts, newOwner)
}

// UnBond is a paid mutator transaction binding the contract method 0x7865d72f.
//
// Solidity: function unBond(uint256 relationID) returns()
func (_Mail *MailTransactor) UnBond(opts *bind.TransactOpts, relationID *big.Int) (*types.Transaction, error) {
	return _Mail.contract.Transact(opts, "unBond", relationID)
}

// UnBond is a paid mutator transaction binding the contract method 0x7865d72f.
//
// Solidity: function unBond(uint256 relationID) returns()
func (_Mail *MailSession) UnBond(relationID *big.Int) (*types.Transaction, error) {
	return _Mail.Contract.UnBond(&_Mail.TransactOpts, relationID)
}

// UnBond is a paid mutator transaction binding the contract method 0x7865d72f.
//
// Solidity: function unBond(uint256 relationID) returns()
func (_Mail *MailTransactorSession) UnBond(relationID *big.Int) (*types.Transaction, error) {
	return _Mail.Contract.UnBond(&_Mail.TransactOpts, relationID)
}

// MailBondRelationshipIterator is returned from FilterBondRelationship and is used to iterate over the raw logs and unpacked data for BondRelationship events raised by the Mail contract.
type MailBondRelationshipIterator struct {
	Event *MailBondRelationship // Event containing the contract specifics and raw log

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
func (it *MailBondRelationshipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailBondRelationship)
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
		it.Event = new(MailBondRelationship)
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
func (it *MailBondRelationshipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailBondRelationshipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailBondRelationship represents a BondRelationship event raised by the Mail contract.
type MailBondRelationship struct {
	Player1        common.Address
	Player2        common.Address
	Types          *big.Int
	RelationshipID *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondRelationship is a free log retrieval operation binding the contract event 0x956005f8c8d9fe678f3aa213fa207bf7d199e9ff6c0fffec713e0807b7142f00.
//
// Solidity: event BondRelationship(address indexed player1, address indexed player2, uint256 indexed types, uint256 relationshipID)
func (_Mail *MailFilterer) FilterBondRelationship(opts *bind.FilterOpts, player1 []common.Address, player2 []common.Address, types []*big.Int) (*MailBondRelationshipIterator, error) {

	var player1Rule []interface{}
	for _, player1Item := range player1 {
		player1Rule = append(player1Rule, player1Item)
	}
	var player2Rule []interface{}
	for _, player2Item := range player2 {
		player2Rule = append(player2Rule, player2Item)
	}
	var typesRule []interface{}
	for _, typesItem := range types {
		typesRule = append(typesRule, typesItem)
	}

	logs, sub, err := _Mail.contract.FilterLogs(opts, "BondRelationship", player1Rule, player2Rule, typesRule)
	if err != nil {
		return nil, err
	}
	return &MailBondRelationshipIterator{contract: _Mail.contract, event: "BondRelationship", logs: logs, sub: sub}, nil
}

// WatchBondRelationship is a free log subscription operation binding the contract event 0x956005f8c8d9fe678f3aa213fa207bf7d199e9ff6c0fffec713e0807b7142f00.
//
// Solidity: event BondRelationship(address indexed player1, address indexed player2, uint256 indexed types, uint256 relationshipID)
func (_Mail *MailFilterer) WatchBondRelationship(opts *bind.WatchOpts, sink chan<- *MailBondRelationship, player1 []common.Address, player2 []common.Address, types []*big.Int) (event.Subscription, error) {

	var player1Rule []interface{}
	for _, player1Item := range player1 {
		player1Rule = append(player1Rule, player1Item)
	}
	var player2Rule []interface{}
	for _, player2Item := range player2 {
		player2Rule = append(player2Rule, player2Item)
	}
	var typesRule []interface{}
	for _, typesItem := range types {
		typesRule = append(typesRule, typesItem)
	}

	logs, sub, err := _Mail.contract.WatchLogs(opts, "BondRelationship", player1Rule, player2Rule, typesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailBondRelationship)
				if err := _Mail.contract.UnpackLog(event, "BondRelationship", log); err != nil {
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

// ParseBondRelationship is a log parse operation binding the contract event 0x956005f8c8d9fe678f3aa213fa207bf7d199e9ff6c0fffec713e0807b7142f00.
//
// Solidity: event BondRelationship(address indexed player1, address indexed player2, uint256 indexed types, uint256 relationshipID)
func (_Mail *MailFilterer) ParseBondRelationship(log types.Log) (*MailBondRelationship, error) {
	event := new(MailBondRelationship)
	if err := _Mail.contract.UnpackLog(event, "BondRelationship", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailClaimMailIterator is returned from FilterClaimMail and is used to iterate over the raw logs and unpacked data for ClaimMail events raised by the Mail contract.
type MailClaimMailIterator struct {
	Event *MailClaimMail // Event containing the contract specifics and raw log

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
func (it *MailClaimMailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailClaimMail)
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
		it.Event = new(MailClaimMail)
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
func (it *MailClaimMailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailClaimMailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailClaimMail represents a ClaimMail event raised by the Mail contract.
type MailClaimMail struct {
	Player common.Address
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimMail is a free log retrieval operation binding the contract event 0xa24234afdca02a8d3145889376ed6307832a41930df54311c44ecd3d72769c10.
//
// Solidity: event ClaimMail(address indexed player, uint256 indexed id)
func (_Mail *MailFilterer) FilterClaimMail(opts *bind.FilterOpts, player []common.Address, id []*big.Int) (*MailClaimMailIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Mail.contract.FilterLogs(opts, "ClaimMail", playerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MailClaimMailIterator{contract: _Mail.contract, event: "ClaimMail", logs: logs, sub: sub}, nil
}

// WatchClaimMail is a free log subscription operation binding the contract event 0xa24234afdca02a8d3145889376ed6307832a41930df54311c44ecd3d72769c10.
//
// Solidity: event ClaimMail(address indexed player, uint256 indexed id)
func (_Mail *MailFilterer) WatchClaimMail(opts *bind.WatchOpts, sink chan<- *MailClaimMail, player []common.Address, id []*big.Int) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Mail.contract.WatchLogs(opts, "ClaimMail", playerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailClaimMail)
				if err := _Mail.contract.UnpackLog(event, "ClaimMail", log); err != nil {
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

// ParseClaimMail is a log parse operation binding the contract event 0xa24234afdca02a8d3145889376ed6307832a41930df54311c44ecd3d72769c10.
//
// Solidity: event ClaimMail(address indexed player, uint256 indexed id)
func (_Mail *MailFilterer) ParseClaimMail(log types.Log) (*MailClaimMail, error) {
	event := new(MailClaimMail)
	if err := _Mail.contract.UnpackLog(event, "ClaimMail", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mail contract.
type MailOwnershipTransferredIterator struct {
	Event *MailOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MailOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailOwnershipTransferred)
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
		it.Event = new(MailOwnershipTransferred)
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
func (it *MailOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailOwnershipTransferred represents a OwnershipTransferred event raised by the Mail contract.
type MailOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mail *MailFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MailOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mail.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MailOwnershipTransferredIterator{contract: _Mail.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mail *MailFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MailOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mail.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailOwnershipTransferred)
				if err := _Mail.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mail *MailFilterer) ParseOwnershipTransferred(log types.Log) (*MailOwnershipTransferred, error) {
	event := new(MailOwnershipTransferred)
	if err := _Mail.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailUnBondRelationshipIterator is returned from FilterUnBondRelationship and is used to iterate over the raw logs and unpacked data for UnBondRelationship events raised by the Mail contract.
type MailUnBondRelationshipIterator struct {
	Event *MailUnBondRelationship // Event containing the contract specifics and raw log

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
func (it *MailUnBondRelationshipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailUnBondRelationship)
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
		it.Event = new(MailUnBondRelationship)
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
func (it *MailUnBondRelationshipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailUnBondRelationshipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailUnBondRelationship represents a UnBondRelationship event raised by the Mail contract.
type MailUnBondRelationship struct {
	RelationshipID *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUnBondRelationship is a free log retrieval operation binding the contract event 0x0507e23770b0afab34ca2476b251d1aa8d4e41e7b010f947ce559c8a2ef39c65.
//
// Solidity: event UnBondRelationship(uint256 indexed relationshipID)
func (_Mail *MailFilterer) FilterUnBondRelationship(opts *bind.FilterOpts, relationshipID []*big.Int) (*MailUnBondRelationshipIterator, error) {

	var relationshipIDRule []interface{}
	for _, relationshipIDItem := range relationshipID {
		relationshipIDRule = append(relationshipIDRule, relationshipIDItem)
	}

	logs, sub, err := _Mail.contract.FilterLogs(opts, "UnBondRelationship", relationshipIDRule)
	if err != nil {
		return nil, err
	}
	return &MailUnBondRelationshipIterator{contract: _Mail.contract, event: "UnBondRelationship", logs: logs, sub: sub}, nil
}

// WatchUnBondRelationship is a free log subscription operation binding the contract event 0x0507e23770b0afab34ca2476b251d1aa8d4e41e7b010f947ce559c8a2ef39c65.
//
// Solidity: event UnBondRelationship(uint256 indexed relationshipID)
func (_Mail *MailFilterer) WatchUnBondRelationship(opts *bind.WatchOpts, sink chan<- *MailUnBondRelationship, relationshipID []*big.Int) (event.Subscription, error) {

	var relationshipIDRule []interface{}
	for _, relationshipIDItem := range relationshipID {
		relationshipIDRule = append(relationshipIDRule, relationshipIDItem)
	}

	logs, sub, err := _Mail.contract.WatchLogs(opts, "UnBondRelationship", relationshipIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailUnBondRelationship)
				if err := _Mail.contract.UnpackLog(event, "UnBondRelationship", log); err != nil {
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

// ParseUnBondRelationship is a log parse operation binding the contract event 0x0507e23770b0afab34ca2476b251d1aa8d4e41e7b010f947ce559c8a2ef39c65.
//
// Solidity: event UnBondRelationship(uint256 indexed relationshipID)
func (_Mail *MailFilterer) ParseUnBondRelationship(log types.Log) (*MailUnBondRelationship, error) {
	event := new(MailUnBondRelationship)
	if err := _Mail.contract.UnpackLog(event, "UnBondRelationship", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
