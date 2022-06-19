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

// CattleABI is the input ABI used to generate the binding from.
const CattleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"AddDeadTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"creationId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"life_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"energy_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"grow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"attack_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defense_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stamina_\",\"type\":\"uint256\"}],\"name\":\"BornCreationBull\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"creationId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"life_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"energy_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"grow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"milk_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"milkRate_\",\"type\":\"uint256\"}],\"name\":\"BornCreationCow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"life_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"energy_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"grow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"star\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[2]\",\"name\":\"parents\",\"type\":\"uint256[2]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"attack_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defense_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stamina_\",\"type\":\"uint256\"}],\"name\":\"BornNormalBull\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"life_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"energy_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"grow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"star\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[2]\",\"name\":\"parents\",\"type\":\"uint256[2]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"milk_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"milkRate_\",\"type\":\"uint256\"}],\"name\":\"BornNormalCow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"GrowUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stars\",\"type\":\"uint256\"}],\"name\":\"UpGradeStar\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time_\",\"type\":\"uint256\"}],\"name\":\"addDeadTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"addExp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"checkUserCowList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"creation_\",\"type\":\"bool\"}],\"name\":\"checkUserCowListType\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cowInfoes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gender\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bornTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"energy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"life\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"growth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAdult\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"attack\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stamina\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defense\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"milk\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"milkRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"star\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cowLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"deadTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"femaleCreation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getAdult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getAttack\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getBronTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getCowParents\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getDefense\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getEnergy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getGender\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getGrowth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getLife\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getMilk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getMilkRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getStamina\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getStar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"growUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isApprove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isCreation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maleCreation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLife\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"parents\",\"type\":\"uint256[2]\"}],\"name\":\"mintNormall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"mintNormallWithParents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintersCreation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintersNormal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"com_\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"male_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"female_\",\"type\":\"uint256\"}],\"name\":\"setGender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"limit_\",\"type\":\"uint256[]\"}],\"name\":\"setLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"setMintersCreation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"setMintersNormal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setSuperMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"starLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"upGradeStar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Cattle is an auto generated Go binding around an Ethereum contract.
type Cattle struct {
	CattleCaller     // Read-only binding to the contract
	CattleTransactor // Write-only binding to the contract
	CattleFilterer   // Log filterer for contract events
}

// CattleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CattleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CattleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CattleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CattleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CattleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CattleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CattleSession struct {
	Contract     *Cattle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CattleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CattleCallerSession struct {
	Contract *CattleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CattleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CattleTransactorSession struct {
	Contract     *CattleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CattleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CattleRaw struct {
	Contract *Cattle // Generic contract binding to access the raw methods on
}

// CattleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CattleCallerRaw struct {
	Contract *CattleCaller // Generic read-only contract binding to access the raw methods on
}

// CattleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CattleTransactorRaw struct {
	Contract *CattleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCattle creates a new instance of Cattle, bound to a specific deployed contract.
func NewCattle(address common.Address, backend bind.ContractBackend) (*Cattle, error) {
	contract, err := bindCattle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cattle{CattleCaller: CattleCaller{contract: contract}, CattleTransactor: CattleTransactor{contract: contract}, CattleFilterer: CattleFilterer{contract: contract}}, nil
}

// NewCattleCaller creates a new read-only instance of Cattle, bound to a specific deployed contract.
func NewCattleCaller(address common.Address, caller bind.ContractCaller) (*CattleCaller, error) {
	contract, err := bindCattle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CattleCaller{contract: contract}, nil
}

// NewCattleTransactor creates a new write-only instance of Cattle, bound to a specific deployed contract.
func NewCattleTransactor(address common.Address, transactor bind.ContractTransactor) (*CattleTransactor, error) {
	contract, err := bindCattle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CattleTransactor{contract: contract}, nil
}

// NewCattleFilterer creates a new log filterer instance of Cattle, bound to a specific deployed contract.
func NewCattleFilterer(address common.Address, filterer bind.ContractFilterer) (*CattleFilterer, error) {
	contract, err := bindCattle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CattleFilterer{contract: contract}, nil
}

// bindCattle binds a generic wrapper to an already deployed contract.
func bindCattle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CattleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cattle *CattleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cattle.Contract.CattleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cattle *CattleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cattle.Contract.CattleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cattle *CattleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cattle.Contract.CattleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cattle *CattleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cattle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cattle *CattleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cattle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cattle *CattleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cattle.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Cattle *CattleCaller) Admin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "admin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Cattle *CattleSession) Admin(arg0 common.Address) (bool, error) {
	return _Cattle.Contract.Admin(&_Cattle.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Cattle *CattleCallerSession) Admin(arg0 common.Address) (bool, error) {
	return _Cattle.Contract.Admin(&_Cattle.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cattle *CattleCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cattle *CattleSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Cattle.Contract.BalanceOf(&_Cattle.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cattle *CattleCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Cattle.Contract.BalanceOf(&_Cattle.CallOpts, owner)
}

// Burned is a free data retrieval call binding the contract method 0x73f42561.
//
// Solidity: function burned() view returns(uint256)
func (_Cattle *CattleCaller) Burned(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "burned")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Burned is a free data retrieval call binding the contract method 0x73f42561.
//
// Solidity: function burned() view returns(uint256)
func (_Cattle *CattleSession) Burned() (*big.Int, error) {
	return _Cattle.Contract.Burned(&_Cattle.CallOpts)
}

// Burned is a free data retrieval call binding the contract method 0x73f42561.
//
// Solidity: function burned() view returns(uint256)
func (_Cattle *CattleCallerSession) Burned() (*big.Int, error) {
	return _Cattle.Contract.Burned(&_Cattle.CallOpts)
}

// CheckUserCowList is a free data retrieval call binding the contract method 0x611f51ac.
//
// Solidity: function checkUserCowList(address player) view returns(uint256[])
func (_Cattle *CattleCaller) CheckUserCowList(opts *bind.CallOpts, player common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "checkUserCowList", player)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CheckUserCowList is a free data retrieval call binding the contract method 0x611f51ac.
//
// Solidity: function checkUserCowList(address player) view returns(uint256[])
func (_Cattle *CattleSession) CheckUserCowList(player common.Address) ([]*big.Int, error) {
	return _Cattle.Contract.CheckUserCowList(&_Cattle.CallOpts, player)
}

// CheckUserCowList is a free data retrieval call binding the contract method 0x611f51ac.
//
// Solidity: function checkUserCowList(address player) view returns(uint256[])
func (_Cattle *CattleCallerSession) CheckUserCowList(player common.Address) ([]*big.Int, error) {
	return _Cattle.Contract.CheckUserCowList(&_Cattle.CallOpts, player)
}

// CheckUserCowListType is a free data retrieval call binding the contract method 0x9b7c6af8.
//
// Solidity: function checkUserCowListType(address player, bool creation_) view returns(uint256[])
func (_Cattle *CattleCaller) CheckUserCowListType(opts *bind.CallOpts, player common.Address, creation_ bool) ([]*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "checkUserCowListType", player, creation_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CheckUserCowListType is a free data retrieval call binding the contract method 0x9b7c6af8.
//
// Solidity: function checkUserCowListType(address player, bool creation_) view returns(uint256[])
func (_Cattle *CattleSession) CheckUserCowListType(player common.Address, creation_ bool) ([]*big.Int, error) {
	return _Cattle.Contract.CheckUserCowListType(&_Cattle.CallOpts, player, creation_)
}

// CheckUserCowListType is a free data retrieval call binding the contract method 0x9b7c6af8.
//
// Solidity: function checkUserCowListType(address player, bool creation_) view returns(uint256[])
func (_Cattle *CattleCallerSession) CheckUserCowListType(player common.Address, creation_ bool) ([]*big.Int, error) {
	return _Cattle.Contract.CheckUserCowListType(&_Cattle.CallOpts, player, creation_)
}

// CowInfoes is a free data retrieval call binding the contract method 0xf610e982.
//
// Solidity: function cowInfoes(uint256 ) view returns(uint256 gender, uint256 bornTime, uint256 energy, uint256 life, uint256 growth, uint256 exp, bool isAdult, uint256 attack, uint256 stamina, uint256 defense, uint256 milk, uint256 milkRate, uint256 star, uint256 deadTime)
func (_Cattle *CattleCaller) CowInfoes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Gender   *big.Int
	BornTime *big.Int
	Energy   *big.Int
	Life     *big.Int
	Growth   *big.Int
	Exp      *big.Int
	IsAdult  bool
	Attack   *big.Int
	Stamina  *big.Int
	Defense  *big.Int
	Milk     *big.Int
	MilkRate *big.Int
	Star     *big.Int
	DeadTime *big.Int
}, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "cowInfoes", arg0)

	outstruct := new(struct {
		Gender   *big.Int
		BornTime *big.Int
		Energy   *big.Int
		Life     *big.Int
		Growth   *big.Int
		Exp      *big.Int
		IsAdult  bool
		Attack   *big.Int
		Stamina  *big.Int
		Defense  *big.Int
		Milk     *big.Int
		MilkRate *big.Int
		Star     *big.Int
		DeadTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gender = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BornTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Energy = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Life = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Growth = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Exp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.IsAdult = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Attack = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Stamina = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Defense = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Milk = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.MilkRate = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.Star = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.DeadTime = *abi.ConvertType(out[13], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CowInfoes is a free data retrieval call binding the contract method 0xf610e982.
//
// Solidity: function cowInfoes(uint256 ) view returns(uint256 gender, uint256 bornTime, uint256 energy, uint256 life, uint256 growth, uint256 exp, bool isAdult, uint256 attack, uint256 stamina, uint256 defense, uint256 milk, uint256 milkRate, uint256 star, uint256 deadTime)
func (_Cattle *CattleSession) CowInfoes(arg0 *big.Int) (struct {
	Gender   *big.Int
	BornTime *big.Int
	Energy   *big.Int
	Life     *big.Int
	Growth   *big.Int
	Exp      *big.Int
	IsAdult  bool
	Attack   *big.Int
	Stamina  *big.Int
	Defense  *big.Int
	Milk     *big.Int
	MilkRate *big.Int
	Star     *big.Int
	DeadTime *big.Int
}, error) {
	return _Cattle.Contract.CowInfoes(&_Cattle.CallOpts, arg0)
}

// CowInfoes is a free data retrieval call binding the contract method 0xf610e982.
//
// Solidity: function cowInfoes(uint256 ) view returns(uint256 gender, uint256 bornTime, uint256 energy, uint256 life, uint256 growth, uint256 exp, bool isAdult, uint256 attack, uint256 stamina, uint256 defense, uint256 milk, uint256 milkRate, uint256 star, uint256 deadTime)
func (_Cattle *CattleCallerSession) CowInfoes(arg0 *big.Int) (struct {
	Gender   *big.Int
	BornTime *big.Int
	Energy   *big.Int
	Life     *big.Int
	Growth   *big.Int
	Exp      *big.Int
	IsAdult  bool
	Attack   *big.Int
	Stamina  *big.Int
	Defense  *big.Int
	Milk     *big.Int
	MilkRate *big.Int
	Star     *big.Int
	DeadTime *big.Int
}, error) {
	return _Cattle.Contract.CowInfoes(&_Cattle.CallOpts, arg0)
}

// CowLimit is a free data retrieval call binding the contract method 0xc8877f8e.
//
// Solidity: function cowLimit(uint256 ) view returns(uint256)
func (_Cattle *CattleCaller) CowLimit(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "cowLimit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CowLimit is a free data retrieval call binding the contract method 0xc8877f8e.
//
// Solidity: function cowLimit(uint256 ) view returns(uint256)
func (_Cattle *CattleSession) CowLimit(arg0 *big.Int) (*big.Int, error) {
	return _Cattle.Contract.CowLimit(&_Cattle.CallOpts, arg0)
}

// CowLimit is a free data retrieval call binding the contract method 0xc8877f8e.
//
// Solidity: function cowLimit(uint256 ) view returns(uint256)
func (_Cattle *CattleCallerSession) CowLimit(arg0 *big.Int) (*big.Int, error) {
	return _Cattle.Contract.CowLimit(&_Cattle.CallOpts, arg0)
}

// CurrentId is a free data retrieval call binding the contract method 0xe00dd161.
//
// Solidity: function currentId() view returns(uint256)
func (_Cattle *CattleCaller) CurrentId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "currentId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentId is a free data retrieval call binding the contract method 0xe00dd161.
//
// Solidity: function currentId() view returns(uint256)
func (_Cattle *CattleSession) CurrentId() (*big.Int, error) {
	return _Cattle.Contract.CurrentId(&_Cattle.CallOpts)
}

// CurrentId is a free data retrieval call binding the contract method 0xe00dd161.
//
// Solidity: function currentId() view returns(uint256)
func (_Cattle *CattleCallerSession) CurrentId() (*big.Int, error) {
	return _Cattle.Contract.CurrentId(&_Cattle.CallOpts)
}

// DeadTime is a free data retrieval call binding the contract method 0x96822dae.
//
// Solidity: function deadTime(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCaller) DeadTime(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "deadTime", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeadTime is a free data retrieval call binding the contract method 0x96822dae.
//
// Solidity: function deadTime(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleSession) DeadTime(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.DeadTime(&_Cattle.CallOpts, tokenId_)
}

// DeadTime is a free data retrieval call binding the contract method 0x96822dae.
//
// Solidity: function deadTime(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCallerSession) DeadTime(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.DeadTime(&_Cattle.CallOpts, tokenId_)
}

// FemaleCreation is a free data retrieval call binding the contract method 0x072c4907.
//
// Solidity: function femaleCreation() view returns(uint256)
func (_Cattle *CattleCaller) FemaleCreation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "femaleCreation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FemaleCreation is a free data retrieval call binding the contract method 0x072c4907.
//
// Solidity: function femaleCreation() view returns(uint256)
func (_Cattle *CattleSession) FemaleCreation() (*big.Int, error) {
	return _Cattle.Contract.FemaleCreation(&_Cattle.CallOpts)
}

// FemaleCreation is a free data retrieval call binding the contract method 0x072c4907.
//
// Solidity: function femaleCreation() view returns(uint256)
func (_Cattle *CattleCallerSession) FemaleCreation() (*big.Int, error) {
	return _Cattle.Contract.FemaleCreation(&_Cattle.CallOpts)
}

// GetAdult is a free data retrieval call binding the contract method 0x9ad1f0bb.
//
// Solidity: function getAdult(uint256 tokenId_) view returns(bool)
func (_Cattle *CattleCaller) GetAdult(opts *bind.CallOpts, tokenId_ *big.Int) (bool, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getAdult", tokenId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAdult is a free data retrieval call binding the contract method 0x9ad1f0bb.
//
// Solidity: function getAdult(uint256 tokenId_) view returns(bool)
func (_Cattle *CattleSession) GetAdult(tokenId_ *big.Int) (bool, error) {
	return _Cattle.Contract.GetAdult(&_Cattle.CallOpts, tokenId_)
}

// GetAdult is a free data retrieval call binding the contract method 0x9ad1f0bb.
//
// Solidity: function getAdult(uint256 tokenId_) view returns(bool)
func (_Cattle *CattleCallerSession) GetAdult(tokenId_ *big.Int) (bool, error) {
	return _Cattle.Contract.GetAdult(&_Cattle.CallOpts, tokenId_)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cattle *CattleCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cattle *CattleSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Cattle.Contract.GetApproved(&_Cattle.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cattle *CattleCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Cattle.Contract.GetApproved(&_Cattle.CallOpts, tokenId)
}

// GetAttack is a free data retrieval call binding the contract method 0xcff840f0.
//
// Solidity: function getAttack(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCaller) GetAttack(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getAttack", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAttack is a free data retrieval call binding the contract method 0xcff840f0.
//
// Solidity: function getAttack(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleSession) GetAttack(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetAttack(&_Cattle.CallOpts, tokenId_)
}

// GetAttack is a free data retrieval call binding the contract method 0xcff840f0.
//
// Solidity: function getAttack(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCallerSession) GetAttack(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetAttack(&_Cattle.CallOpts, tokenId_)
}

// GetBronTime is a free data retrieval call binding the contract method 0xbba0f646.
//
// Solidity: function getBronTime(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCaller) GetBronTime(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getBronTime", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBronTime is a free data retrieval call binding the contract method 0xbba0f646.
//
// Solidity: function getBronTime(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleSession) GetBronTime(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetBronTime(&_Cattle.CallOpts, tokenId_)
}

// GetBronTime is a free data retrieval call binding the contract method 0xbba0f646.
//
// Solidity: function getBronTime(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCallerSession) GetBronTime(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetBronTime(&_Cattle.CallOpts, tokenId_)
}

// GetCowParents is a free data retrieval call binding the contract method 0x5f64d7f3.
//
// Solidity: function getCowParents(uint256 tokenId_) view returns(uint256[2])
func (_Cattle *CattleCaller) GetCowParents(opts *bind.CallOpts, tokenId_ *big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getCowParents", tokenId_)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetCowParents is a free data retrieval call binding the contract method 0x5f64d7f3.
//
// Solidity: function getCowParents(uint256 tokenId_) view returns(uint256[2])
func (_Cattle *CattleSession) GetCowParents(tokenId_ *big.Int) ([2]*big.Int, error) {
	return _Cattle.Contract.GetCowParents(&_Cattle.CallOpts, tokenId_)
}

// GetCowParents is a free data retrieval call binding the contract method 0x5f64d7f3.
//
// Solidity: function getCowParents(uint256 tokenId_) view returns(uint256[2])
func (_Cattle *CattleCallerSession) GetCowParents(tokenId_ *big.Int) ([2]*big.Int, error) {
	return _Cattle.Contract.GetCowParents(&_Cattle.CallOpts, tokenId_)
}

// GetDefense is a free data retrieval call binding the contract method 0xeb22daa3.
//
// Solidity: function getDefense(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCaller) GetDefense(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getDefense", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDefense is a free data retrieval call binding the contract method 0xeb22daa3.
//
// Solidity: function getDefense(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleSession) GetDefense(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetDefense(&_Cattle.CallOpts, tokenId_)
}

// GetDefense is a free data retrieval call binding the contract method 0xeb22daa3.
//
// Solidity: function getDefense(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCallerSession) GetDefense(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetDefense(&_Cattle.CallOpts, tokenId_)
}

// GetEnergy is a free data retrieval call binding the contract method 0x7e0358a6.
//
// Solidity: function getEnergy(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCaller) GetEnergy(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getEnergy", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEnergy is a free data retrieval call binding the contract method 0x7e0358a6.
//
// Solidity: function getEnergy(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleSession) GetEnergy(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetEnergy(&_Cattle.CallOpts, tokenId_)
}

// GetEnergy is a free data retrieval call binding the contract method 0x7e0358a6.
//
// Solidity: function getEnergy(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCallerSession) GetEnergy(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetEnergy(&_Cattle.CallOpts, tokenId_)
}

// GetGender is a free data retrieval call binding the contract method 0xe06d2eb5.
//
// Solidity: function getGender(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCaller) GetGender(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getGender", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGender is a free data retrieval call binding the contract method 0xe06d2eb5.
//
// Solidity: function getGender(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleSession) GetGender(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetGender(&_Cattle.CallOpts, tokenId_)
}

// GetGender is a free data retrieval call binding the contract method 0xe06d2eb5.
//
// Solidity: function getGender(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCallerSession) GetGender(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetGender(&_Cattle.CallOpts, tokenId_)
}

// GetGrowth is a free data retrieval call binding the contract method 0x51c18e0c.
//
// Solidity: function getGrowth(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCaller) GetGrowth(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getGrowth", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGrowth is a free data retrieval call binding the contract method 0x51c18e0c.
//
// Solidity: function getGrowth(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleSession) GetGrowth(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetGrowth(&_Cattle.CallOpts, tokenId_)
}

// GetGrowth is a free data retrieval call binding the contract method 0x51c18e0c.
//
// Solidity: function getGrowth(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCallerSession) GetGrowth(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetGrowth(&_Cattle.CallOpts, tokenId_)
}

// GetLife is a free data retrieval call binding the contract method 0x11251f1c.
//
// Solidity: function getLife(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCaller) GetLife(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getLife", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLife is a free data retrieval call binding the contract method 0x11251f1c.
//
// Solidity: function getLife(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleSession) GetLife(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetLife(&_Cattle.CallOpts, tokenId_)
}

// GetLife is a free data retrieval call binding the contract method 0x11251f1c.
//
// Solidity: function getLife(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCallerSession) GetLife(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetLife(&_Cattle.CallOpts, tokenId_)
}

// GetMilk is a free data retrieval call binding the contract method 0x184e9c2d.
//
// Solidity: function getMilk(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCaller) GetMilk(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getMilk", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMilk is a free data retrieval call binding the contract method 0x184e9c2d.
//
// Solidity: function getMilk(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleSession) GetMilk(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetMilk(&_Cattle.CallOpts, tokenId_)
}

// GetMilk is a free data retrieval call binding the contract method 0x184e9c2d.
//
// Solidity: function getMilk(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCallerSession) GetMilk(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetMilk(&_Cattle.CallOpts, tokenId_)
}

// GetMilkRate is a free data retrieval call binding the contract method 0x92d9495d.
//
// Solidity: function getMilkRate(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCaller) GetMilkRate(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getMilkRate", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMilkRate is a free data retrieval call binding the contract method 0x92d9495d.
//
// Solidity: function getMilkRate(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleSession) GetMilkRate(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetMilkRate(&_Cattle.CallOpts, tokenId_)
}

// GetMilkRate is a free data retrieval call binding the contract method 0x92d9495d.
//
// Solidity: function getMilkRate(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCallerSession) GetMilkRate(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetMilkRate(&_Cattle.CallOpts, tokenId_)
}

// GetStamina is a free data retrieval call binding the contract method 0xe6225630.
//
// Solidity: function getStamina(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCaller) GetStamina(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getStamina", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStamina is a free data retrieval call binding the contract method 0xe6225630.
//
// Solidity: function getStamina(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleSession) GetStamina(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetStamina(&_Cattle.CallOpts, tokenId_)
}

// GetStamina is a free data retrieval call binding the contract method 0xe6225630.
//
// Solidity: function getStamina(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCallerSession) GetStamina(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetStamina(&_Cattle.CallOpts, tokenId_)
}

// GetStar is a free data retrieval call binding the contract method 0xc5426124.
//
// Solidity: function getStar(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCaller) GetStar(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "getStar", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStar is a free data retrieval call binding the contract method 0xc5426124.
//
// Solidity: function getStar(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleSession) GetStar(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetStar(&_Cattle.CallOpts, tokenId_)
}

// GetStar is a free data retrieval call binding the contract method 0xc5426124.
//
// Solidity: function getStar(uint256 tokenId_) view returns(uint256)
func (_Cattle *CattleCallerSession) GetStar(tokenId_ *big.Int) (*big.Int, error) {
	return _Cattle.Contract.GetStar(&_Cattle.CallOpts, tokenId_)
}

// IsApprove is a free data retrieval call binding the contract method 0xff45fec8.
//
// Solidity: function isApprove(address , address , uint256 ) view returns(bool)
func (_Cattle *CattleCaller) IsApprove(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "isApprove", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprove is a free data retrieval call binding the contract method 0xff45fec8.
//
// Solidity: function isApprove(address , address , uint256 ) view returns(bool)
func (_Cattle *CattleSession) IsApprove(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _Cattle.Contract.IsApprove(&_Cattle.CallOpts, arg0, arg1, arg2)
}

// IsApprove is a free data retrieval call binding the contract method 0xff45fec8.
//
// Solidity: function isApprove(address , address , uint256 ) view returns(bool)
func (_Cattle *CattleCallerSession) IsApprove(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _Cattle.Contract.IsApprove(&_Cattle.CallOpts, arg0, arg1, arg2)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cattle *CattleCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cattle *CattleSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Cattle.Contract.IsApprovedForAll(&_Cattle.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cattle *CattleCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Cattle.Contract.IsApprovedForAll(&_Cattle.CallOpts, owner, operator)
}

// IsCreation is a free data retrieval call binding the contract method 0xf0a90ff0.
//
// Solidity: function isCreation(uint256 ) view returns(bool)
func (_Cattle *CattleCaller) IsCreation(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "isCreation", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCreation is a free data retrieval call binding the contract method 0xf0a90ff0.
//
// Solidity: function isCreation(uint256 ) view returns(bool)
func (_Cattle *CattleSession) IsCreation(arg0 *big.Int) (bool, error) {
	return _Cattle.Contract.IsCreation(&_Cattle.CallOpts, arg0)
}

// IsCreation is a free data retrieval call binding the contract method 0xf0a90ff0.
//
// Solidity: function isCreation(uint256 ) view returns(bool)
func (_Cattle *CattleCallerSession) IsCreation(arg0 *big.Int) (bool, error) {
	return _Cattle.Contract.IsCreation(&_Cattle.CallOpts, arg0)
}

// MaleCreation is a free data retrieval call binding the contract method 0xb969f222.
//
// Solidity: function maleCreation() view returns(uint256)
func (_Cattle *CattleCaller) MaleCreation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "maleCreation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaleCreation is a free data retrieval call binding the contract method 0xb969f222.
//
// Solidity: function maleCreation() view returns(uint256)
func (_Cattle *CattleSession) MaleCreation() (*big.Int, error) {
	return _Cattle.Contract.MaleCreation(&_Cattle.CallOpts)
}

// MaleCreation is a free data retrieval call binding the contract method 0xb969f222.
//
// Solidity: function maleCreation() view returns(uint256)
func (_Cattle *CattleCallerSession) MaleCreation() (*big.Int, error) {
	return _Cattle.Contract.MaleCreation(&_Cattle.CallOpts)
}

// MaxLife is a free data retrieval call binding the contract method 0x232cba0d.
//
// Solidity: function maxLife() view returns(uint256)
func (_Cattle *CattleCaller) MaxLife(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "maxLife")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLife is a free data retrieval call binding the contract method 0x232cba0d.
//
// Solidity: function maxLife() view returns(uint256)
func (_Cattle *CattleSession) MaxLife() (*big.Int, error) {
	return _Cattle.Contract.MaxLife(&_Cattle.CallOpts)
}

// MaxLife is a free data retrieval call binding the contract method 0x232cba0d.
//
// Solidity: function maxLife() view returns(uint256)
func (_Cattle *CattleCallerSession) MaxLife() (*big.Int, error) {
	return _Cattle.Contract.MaxLife(&_Cattle.CallOpts)
}

// MintersCreation is a free data retrieval call binding the contract method 0xd6eac26c.
//
// Solidity: function mintersCreation(address ) view returns(uint256)
func (_Cattle *CattleCaller) MintersCreation(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "mintersCreation", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintersCreation is a free data retrieval call binding the contract method 0xd6eac26c.
//
// Solidity: function mintersCreation(address ) view returns(uint256)
func (_Cattle *CattleSession) MintersCreation(arg0 common.Address) (*big.Int, error) {
	return _Cattle.Contract.MintersCreation(&_Cattle.CallOpts, arg0)
}

// MintersCreation is a free data retrieval call binding the contract method 0xd6eac26c.
//
// Solidity: function mintersCreation(address ) view returns(uint256)
func (_Cattle *CattleCallerSession) MintersCreation(arg0 common.Address) (*big.Int, error) {
	return _Cattle.Contract.MintersCreation(&_Cattle.CallOpts, arg0)
}

// MintersNormal is a free data retrieval call binding the contract method 0xb9eea722.
//
// Solidity: function mintersNormal(address ) view returns(uint256)
func (_Cattle *CattleCaller) MintersNormal(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "mintersNormal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintersNormal is a free data retrieval call binding the contract method 0xb9eea722.
//
// Solidity: function mintersNormal(address ) view returns(uint256)
func (_Cattle *CattleSession) MintersNormal(arg0 common.Address) (*big.Int, error) {
	return _Cattle.Contract.MintersNormal(&_Cattle.CallOpts, arg0)
}

// MintersNormal is a free data retrieval call binding the contract method 0xb9eea722.
//
// Solidity: function mintersNormal(address ) view returns(uint256)
func (_Cattle *CattleCallerSession) MintersNormal(arg0 common.Address) (*big.Int, error) {
	return _Cattle.Contract.MintersNormal(&_Cattle.CallOpts, arg0)
}

// MyBaseURI is a free data retrieval call binding the contract method 0x372a3b82.
//
// Solidity: function myBaseURI() view returns(string)
func (_Cattle *CattleCaller) MyBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "myBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MyBaseURI is a free data retrieval call binding the contract method 0x372a3b82.
//
// Solidity: function myBaseURI() view returns(string)
func (_Cattle *CattleSession) MyBaseURI() (string, error) {
	return _Cattle.Contract.MyBaseURI(&_Cattle.CallOpts)
}

// MyBaseURI is a free data retrieval call binding the contract method 0x372a3b82.
//
// Solidity: function myBaseURI() view returns(string)
func (_Cattle *CattleCallerSession) MyBaseURI() (string, error) {
	return _Cattle.Contract.MyBaseURI(&_Cattle.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cattle *CattleCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cattle *CattleSession) Name() (string, error) {
	return _Cattle.Contract.Name(&_Cattle.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cattle *CattleCallerSession) Name() (string, error) {
	return _Cattle.Contract.Name(&_Cattle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cattle *CattleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cattle *CattleSession) Owner() (common.Address, error) {
	return _Cattle.Contract.Owner(&_Cattle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cattle *CattleCallerSession) Owner() (common.Address, error) {
	return _Cattle.Contract.Owner(&_Cattle.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cattle *CattleCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cattle *CattleSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Cattle.Contract.OwnerOf(&_Cattle.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cattle *CattleCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Cattle.Contract.OwnerOf(&_Cattle.CallOpts, tokenId)
}

// StarLimit is a free data retrieval call binding the contract method 0xe2608438.
//
// Solidity: function starLimit(uint256 ) view returns(uint256)
func (_Cattle *CattleCaller) StarLimit(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "starLimit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StarLimit is a free data retrieval call binding the contract method 0xe2608438.
//
// Solidity: function starLimit(uint256 ) view returns(uint256)
func (_Cattle *CattleSession) StarLimit(arg0 *big.Int) (*big.Int, error) {
	return _Cattle.Contract.StarLimit(&_Cattle.CallOpts, arg0)
}

// StarLimit is a free data retrieval call binding the contract method 0xe2608438.
//
// Solidity: function starLimit(uint256 ) view returns(uint256)
func (_Cattle *CattleCallerSession) StarLimit(arg0 *big.Int) (*big.Int, error) {
	return _Cattle.Contract.StarLimit(&_Cattle.CallOpts, arg0)
}

// SuperMinter is a free data retrieval call binding the contract method 0x62e42cb0.
//
// Solidity: function superMinter() view returns(address)
func (_Cattle *CattleCaller) SuperMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "superMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperMinter is a free data retrieval call binding the contract method 0x62e42cb0.
//
// Solidity: function superMinter() view returns(address)
func (_Cattle *CattleSession) SuperMinter() (common.Address, error) {
	return _Cattle.Contract.SuperMinter(&_Cattle.CallOpts)
}

// SuperMinter is a free data retrieval call binding the contract method 0x62e42cb0.
//
// Solidity: function superMinter() view returns(address)
func (_Cattle *CattleCallerSession) SuperMinter() (common.Address, error) {
	return _Cattle.Contract.SuperMinter(&_Cattle.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cattle *CattleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cattle *CattleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cattle.Contract.SupportsInterface(&_Cattle.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cattle *CattleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cattle.Contract.SupportsInterface(&_Cattle.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cattle *CattleCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cattle *CattleSession) Symbol() (string, error) {
	return _Cattle.Contract.Symbol(&_Cattle.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cattle *CattleCallerSession) Symbol() (string, error) {
	return _Cattle.Contract.Symbol(&_Cattle.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Cattle *CattleCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Cattle *CattleSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Cattle.Contract.TokenByIndex(&_Cattle.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Cattle *CattleCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Cattle.Contract.TokenByIndex(&_Cattle.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Cattle *CattleCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Cattle *CattleSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Cattle.Contract.TokenOfOwnerByIndex(&_Cattle.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Cattle *CattleCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Cattle.Contract.TokenOfOwnerByIndex(&_Cattle.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_Cattle *CattleCaller) TokenURI(opts *bind.CallOpts, tokenId_ *big.Int) (string, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "tokenURI", tokenId_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_Cattle *CattleSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _Cattle.Contract.TokenURI(&_Cattle.CallOpts, tokenId_)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_Cattle *CattleCallerSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _Cattle.Contract.TokenURI(&_Cattle.CallOpts, tokenId_)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cattle *CattleCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cattle.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cattle *CattleSession) TotalSupply() (*big.Int, error) {
	return _Cattle.Contract.TotalSupply(&_Cattle.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cattle *CattleCallerSession) TotalSupply() (*big.Int, error) {
	return _Cattle.Contract.TotalSupply(&_Cattle.CallOpts)
}

// AddDeadTime is a paid mutator transaction binding the contract method 0x5f70a1d5.
//
// Solidity: function addDeadTime(uint256 tokenId, uint256 time_) returns()
func (_Cattle *CattleTransactor) AddDeadTime(opts *bind.TransactOpts, tokenId *big.Int, time_ *big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "addDeadTime", tokenId, time_)
}

// AddDeadTime is a paid mutator transaction binding the contract method 0x5f70a1d5.
//
// Solidity: function addDeadTime(uint256 tokenId, uint256 time_) returns()
func (_Cattle *CattleSession) AddDeadTime(tokenId *big.Int, time_ *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.AddDeadTime(&_Cattle.TransactOpts, tokenId, time_)
}

// AddDeadTime is a paid mutator transaction binding the contract method 0x5f70a1d5.
//
// Solidity: function addDeadTime(uint256 tokenId, uint256 time_) returns()
func (_Cattle *CattleTransactorSession) AddDeadTime(tokenId *big.Int, time_ *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.AddDeadTime(&_Cattle.TransactOpts, tokenId, time_)
}

// AddExp is a paid mutator transaction binding the contract method 0x4147da13.
//
// Solidity: function addExp(uint256 tokenId_, uint256 amount_) returns()
func (_Cattle *CattleTransactor) AddExp(opts *bind.TransactOpts, tokenId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "addExp", tokenId_, amount_)
}

// AddExp is a paid mutator transaction binding the contract method 0x4147da13.
//
// Solidity: function addExp(uint256 tokenId_, uint256 amount_) returns()
func (_Cattle *CattleSession) AddExp(tokenId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.AddExp(&_Cattle.TransactOpts, tokenId_, amount_)
}

// AddExp is a paid mutator transaction binding the contract method 0x4147da13.
//
// Solidity: function addExp(uint256 tokenId_, uint256 amount_) returns()
func (_Cattle *CattleTransactorSession) AddExp(tokenId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.AddExp(&_Cattle.TransactOpts, tokenId_, amount_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Cattle *CattleTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Cattle *CattleSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.Approve(&_Cattle.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Cattle *CattleTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.Approve(&_Cattle.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId_) returns(bool)
func (_Cattle *CattleTransactor) Burn(opts *bind.TransactOpts, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "burn", tokenId_)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId_) returns(bool)
func (_Cattle *CattleSession) Burn(tokenId_ *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.Burn(&_Cattle.TransactOpts, tokenId_)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId_) returns(bool)
func (_Cattle *CattleTransactorSession) Burn(tokenId_ *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.Burn(&_Cattle.TransactOpts, tokenId_)
}

// GrowUp is a paid mutator transaction binding the contract method 0x84686c64.
//
// Solidity: function growUp(uint256 tokenId) returns()
func (_Cattle *CattleTransactor) GrowUp(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "growUp", tokenId)
}

// GrowUp is a paid mutator transaction binding the contract method 0x84686c64.
//
// Solidity: function growUp(uint256 tokenId) returns()
func (_Cattle *CattleSession) GrowUp(tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.GrowUp(&_Cattle.TransactOpts, tokenId)
}

// GrowUp is a paid mutator transaction binding the contract method 0x84686c64.
//
// Solidity: function growUp(uint256 tokenId) returns()
func (_Cattle *CattleTransactorSession) GrowUp(tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.GrowUp(&_Cattle.TransactOpts, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Cattle *CattleTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Cattle *CattleSession) Initialize() (*types.Transaction, error) {
	return _Cattle.Contract.Initialize(&_Cattle.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Cattle *CattleTransactorSession) Initialize() (*types.Transaction, error) {
	return _Cattle.Contract.Initialize(&_Cattle.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address player) returns()
func (_Cattle *CattleTransactor) Mint(opts *bind.TransactOpts, player common.Address) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "mint", player)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address player) returns()
func (_Cattle *CattleSession) Mint(player common.Address) (*types.Transaction, error) {
	return _Cattle.Contract.Mint(&_Cattle.TransactOpts, player)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address player) returns()
func (_Cattle *CattleTransactorSession) Mint(player common.Address) (*types.Transaction, error) {
	return _Cattle.Contract.Mint(&_Cattle.TransactOpts, player)
}

// MintNormall is a paid mutator transaction binding the contract method 0xae9921de.
//
// Solidity: function mintNormall(address player, uint256[2] parents) returns()
func (_Cattle *CattleTransactor) MintNormall(opts *bind.TransactOpts, player common.Address, parents [2]*big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "mintNormall", player, parents)
}

// MintNormall is a paid mutator transaction binding the contract method 0xae9921de.
//
// Solidity: function mintNormall(address player, uint256[2] parents) returns()
func (_Cattle *CattleSession) MintNormall(player common.Address, parents [2]*big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.MintNormall(&_Cattle.TransactOpts, player, parents)
}

// MintNormall is a paid mutator transaction binding the contract method 0xae9921de.
//
// Solidity: function mintNormall(address player, uint256[2] parents) returns()
func (_Cattle *CattleTransactorSession) MintNormall(player common.Address, parents [2]*big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.MintNormall(&_Cattle.TransactOpts, player, parents)
}

// MintNormallWithParents is a paid mutator transaction binding the contract method 0xbb3b9740.
//
// Solidity: function mintNormallWithParents(address player) returns()
func (_Cattle *CattleTransactor) MintNormallWithParents(opts *bind.TransactOpts, player common.Address) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "mintNormallWithParents", player)
}

// MintNormallWithParents is a paid mutator transaction binding the contract method 0xbb3b9740.
//
// Solidity: function mintNormallWithParents(address player) returns()
func (_Cattle *CattleSession) MintNormallWithParents(player common.Address) (*types.Transaction, error) {
	return _Cattle.Contract.MintNormallWithParents(&_Cattle.TransactOpts, player)
}

// MintNormallWithParents is a paid mutator transaction binding the contract method 0xbb3b9740.
//
// Solidity: function mintNormallWithParents(address player) returns()
func (_Cattle *CattleTransactorSession) MintNormallWithParents(player common.Address) (*types.Transaction, error) {
	return _Cattle.Contract.MintNormallWithParents(&_Cattle.TransactOpts, player)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cattle *CattleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cattle *CattleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cattle.Contract.RenounceOwnership(&_Cattle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cattle *CattleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cattle.Contract.RenounceOwnership(&_Cattle.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cattle *CattleTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cattle *CattleSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.SafeTransferFrom(&_Cattle.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cattle *CattleTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.SafeTransferFrom(&_Cattle.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Cattle *CattleTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Cattle *CattleSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Cattle.Contract.SafeTransferFrom0(&_Cattle.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Cattle *CattleTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Cattle.Contract.SafeTransferFrom0(&_Cattle.TransactOpts, from, to, tokenId, _data)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr_, bool com_) returns()
func (_Cattle *CattleTransactor) SetAdmin(opts *bind.TransactOpts, addr_ common.Address, com_ bool) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "setAdmin", addr_, com_)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr_, bool com_) returns()
func (_Cattle *CattleSession) SetAdmin(addr_ common.Address, com_ bool) (*types.Transaction, error) {
	return _Cattle.Contract.SetAdmin(&_Cattle.TransactOpts, addr_, com_)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address addr_, bool com_) returns()
func (_Cattle *CattleTransactorSession) SetAdmin(addr_ common.Address, com_ bool) (*types.Transaction, error) {
	return _Cattle.Contract.SetAdmin(&_Cattle.TransactOpts, addr_, com_)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cattle *CattleTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cattle *CattleSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Cattle.Contract.SetApprovalForAll(&_Cattle.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cattle *CattleTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Cattle.Contract.SetApprovalForAll(&_Cattle.TransactOpts, operator, approved)
}

// SetGender is a paid mutator transaction binding the contract method 0x780ff402.
//
// Solidity: function setGender(uint256 male_, uint256 female_) returns()
func (_Cattle *CattleTransactor) SetGender(opts *bind.TransactOpts, male_ *big.Int, female_ *big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "setGender", male_, female_)
}

// SetGender is a paid mutator transaction binding the contract method 0x780ff402.
//
// Solidity: function setGender(uint256 male_, uint256 female_) returns()
func (_Cattle *CattleSession) SetGender(male_ *big.Int, female_ *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.SetGender(&_Cattle.TransactOpts, male_, female_)
}

// SetGender is a paid mutator transaction binding the contract method 0x780ff402.
//
// Solidity: function setGender(uint256 male_, uint256 female_) returns()
func (_Cattle *CattleTransactorSession) SetGender(male_ *big.Int, female_ *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.SetGender(&_Cattle.TransactOpts, male_, female_)
}

// SetLimit is a paid mutator transaction binding the contract method 0xa435ff65.
//
// Solidity: function setLimit(uint256[] limit_) returns()
func (_Cattle *CattleTransactor) SetLimit(opts *bind.TransactOpts, limit_ []*big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "setLimit", limit_)
}

// SetLimit is a paid mutator transaction binding the contract method 0xa435ff65.
//
// Solidity: function setLimit(uint256[] limit_) returns()
func (_Cattle *CattleSession) SetLimit(limit_ []*big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.SetLimit(&_Cattle.TransactOpts, limit_)
}

// SetLimit is a paid mutator transaction binding the contract method 0xa435ff65.
//
// Solidity: function setLimit(uint256[] limit_) returns()
func (_Cattle *CattleTransactorSession) SetLimit(limit_ []*big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.SetLimit(&_Cattle.TransactOpts, limit_)
}

// SetMintersCreation is a paid mutator transaction binding the contract method 0x33523433.
//
// Solidity: function setMintersCreation(address addr_, uint256 amount_) returns()
func (_Cattle *CattleTransactor) SetMintersCreation(opts *bind.TransactOpts, addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "setMintersCreation", addr_, amount_)
}

// SetMintersCreation is a paid mutator transaction binding the contract method 0x33523433.
//
// Solidity: function setMintersCreation(address addr_, uint256 amount_) returns()
func (_Cattle *CattleSession) SetMintersCreation(addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.SetMintersCreation(&_Cattle.TransactOpts, addr_, amount_)
}

// SetMintersCreation is a paid mutator transaction binding the contract method 0x33523433.
//
// Solidity: function setMintersCreation(address addr_, uint256 amount_) returns()
func (_Cattle *CattleTransactorSession) SetMintersCreation(addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.SetMintersCreation(&_Cattle.TransactOpts, addr_, amount_)
}

// SetMintersNormal is a paid mutator transaction binding the contract method 0x148eebd7.
//
// Solidity: function setMintersNormal(address addr_, uint256 amount_) returns()
func (_Cattle *CattleTransactor) SetMintersNormal(opts *bind.TransactOpts, addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "setMintersNormal", addr_, amount_)
}

// SetMintersNormal is a paid mutator transaction binding the contract method 0x148eebd7.
//
// Solidity: function setMintersNormal(address addr_, uint256 amount_) returns()
func (_Cattle *CattleSession) SetMintersNormal(addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.SetMintersNormal(&_Cattle.TransactOpts, addr_, amount_)
}

// SetMintersNormal is a paid mutator transaction binding the contract method 0x148eebd7.
//
// Solidity: function setMintersNormal(address addr_, uint256 amount_) returns()
func (_Cattle *CattleTransactorSession) SetMintersNormal(addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.SetMintersNormal(&_Cattle.TransactOpts, addr_, amount_)
}

// SetSuperMinter is a paid mutator transaction binding the contract method 0xb41d74d8.
//
// Solidity: function setSuperMinter(address addr) returns()
func (_Cattle *CattleTransactor) SetSuperMinter(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "setSuperMinter", addr)
}

// SetSuperMinter is a paid mutator transaction binding the contract method 0xb41d74d8.
//
// Solidity: function setSuperMinter(address addr) returns()
func (_Cattle *CattleSession) SetSuperMinter(addr common.Address) (*types.Transaction, error) {
	return _Cattle.Contract.SetSuperMinter(&_Cattle.TransactOpts, addr)
}

// SetSuperMinter is a paid mutator transaction binding the contract method 0xb41d74d8.
//
// Solidity: function setSuperMinter(address addr) returns()
func (_Cattle *CattleTransactorSession) SetSuperMinter(addr common.Address) (*types.Transaction, error) {
	return _Cattle.Contract.SetSuperMinter(&_Cattle.TransactOpts, addr)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cattle *CattleTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cattle *CattleSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.TransferFrom(&_Cattle.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cattle *CattleTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.TransferFrom(&_Cattle.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cattle *CattleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cattle *CattleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cattle.Contract.TransferOwnership(&_Cattle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cattle *CattleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cattle.Contract.TransferOwnership(&_Cattle.TransactOpts, newOwner)
}

// UpGradeStar is a paid mutator transaction binding the contract method 0x25421013.
//
// Solidity: function upGradeStar(uint256 tokenId) returns()
func (_Cattle *CattleTransactor) UpGradeStar(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.contract.Transact(opts, "upGradeStar", tokenId)
}

// UpGradeStar is a paid mutator transaction binding the contract method 0x25421013.
//
// Solidity: function upGradeStar(uint256 tokenId) returns()
func (_Cattle *CattleSession) UpGradeStar(tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.UpGradeStar(&_Cattle.TransactOpts, tokenId)
}

// UpGradeStar is a paid mutator transaction binding the contract method 0x25421013.
//
// Solidity: function upGradeStar(uint256 tokenId) returns()
func (_Cattle *CattleTransactorSession) UpGradeStar(tokenId *big.Int) (*types.Transaction, error) {
	return _Cattle.Contract.UpGradeStar(&_Cattle.TransactOpts, tokenId)
}

// CattleAddDeadTimeIterator is returned from FilterAddDeadTime and is used to iterate over the raw logs and unpacked data for AddDeadTime events raised by the Cattle contract.
type CattleAddDeadTimeIterator struct {
	Event *CattleAddDeadTime // Event containing the contract specifics and raw log

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
func (it *CattleAddDeadTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CattleAddDeadTime)
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
		it.Event = new(CattleAddDeadTime)
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
func (it *CattleAddDeadTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CattleAddDeadTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CattleAddDeadTime represents a AddDeadTime event raised by the Cattle contract.
type CattleAddDeadTime struct {
	TokenId *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddDeadTime is a free log retrieval operation binding the contract event 0xc079f36d36fb34d06c2882fd7be02c1969f2dc8ce432ad6f087f977271f53f38.
//
// Solidity: event AddDeadTime(uint256 indexed tokenId, uint256 indexed time)
func (_Cattle *CattleFilterer) FilterAddDeadTime(opts *bind.FilterOpts, tokenId []*big.Int, time []*big.Int) (*CattleAddDeadTimeIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var timeRule []interface{}
	for _, timeItem := range time {
		timeRule = append(timeRule, timeItem)
	}

	logs, sub, err := _Cattle.contract.FilterLogs(opts, "AddDeadTime", tokenIdRule, timeRule)
	if err != nil {
		return nil, err
	}
	return &CattleAddDeadTimeIterator{contract: _Cattle.contract, event: "AddDeadTime", logs: logs, sub: sub}, nil
}

// WatchAddDeadTime is a free log subscription operation binding the contract event 0xc079f36d36fb34d06c2882fd7be02c1969f2dc8ce432ad6f087f977271f53f38.
//
// Solidity: event AddDeadTime(uint256 indexed tokenId, uint256 indexed time)
func (_Cattle *CattleFilterer) WatchAddDeadTime(opts *bind.WatchOpts, sink chan<- *CattleAddDeadTime, tokenId []*big.Int, time []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var timeRule []interface{}
	for _, timeItem := range time {
		timeRule = append(timeRule, timeItem)
	}

	logs, sub, err := _Cattle.contract.WatchLogs(opts, "AddDeadTime", tokenIdRule, timeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CattleAddDeadTime)
				if err := _Cattle.contract.UnpackLog(event, "AddDeadTime", log); err != nil {
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

// ParseAddDeadTime is a log parse operation binding the contract event 0xc079f36d36fb34d06c2882fd7be02c1969f2dc8ce432ad6f087f977271f53f38.
//
// Solidity: event AddDeadTime(uint256 indexed tokenId, uint256 indexed time)
func (_Cattle *CattleFilterer) ParseAddDeadTime(log types.Log) (*CattleAddDeadTime, error) {
	event := new(CattleAddDeadTime)
	if err := _Cattle.contract.UnpackLog(event, "AddDeadTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CattleApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cattle contract.
type CattleApprovalIterator struct {
	Event *CattleApproval // Event containing the contract specifics and raw log

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
func (it *CattleApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CattleApproval)
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
		it.Event = new(CattleApproval)
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
func (it *CattleApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CattleApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CattleApproval represents a Approval event raised by the Cattle contract.
type CattleApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cattle *CattleFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CattleApprovalIterator, error) {

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

	logs, sub, err := _Cattle.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CattleApprovalIterator{contract: _Cattle.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cattle *CattleFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CattleApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Cattle.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CattleApproval)
				if err := _Cattle.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Cattle *CattleFilterer) ParseApproval(log types.Log) (*CattleApproval, error) {
	event := new(CattleApproval)
	if err := _Cattle.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CattleApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Cattle contract.
type CattleApprovalForAllIterator struct {
	Event *CattleApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CattleApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CattleApprovalForAll)
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
		it.Event = new(CattleApprovalForAll)
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
func (it *CattleApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CattleApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CattleApprovalForAll represents a ApprovalForAll event raised by the Cattle contract.
type CattleApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cattle *CattleFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CattleApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Cattle.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CattleApprovalForAllIterator{contract: _Cattle.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cattle *CattleFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CattleApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Cattle.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CattleApprovalForAll)
				if err := _Cattle.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Cattle *CattleFilterer) ParseApprovalForAll(log types.Log) (*CattleApprovalForAll, error) {
	event := new(CattleApprovalForAll)
	if err := _Cattle.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CattleBornCreationBullIterator is returned from FilterBornCreationBull and is used to iterate over the raw logs and unpacked data for BornCreationBull events raised by the Cattle contract.
type CattleBornCreationBullIterator struct {
	Event *CattleBornCreationBull // Event containing the contract specifics and raw log

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
func (it *CattleBornCreationBullIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CattleBornCreationBull)
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
		it.Event = new(CattleBornCreationBull)
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
func (it *CattleBornCreationBullIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CattleBornCreationBullIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CattleBornCreationBull represents a BornCreationBull event raised by the Cattle contract.
type CattleBornCreationBull struct {
	Sender     common.Address
	TokenId    *big.Int
	CreationId *big.Int
	Life       *big.Int
	Energy     *big.Int
	Grow       *big.Int
	DeadTime   *big.Int
	Attack     *big.Int
	Defense    *big.Int
	Stamina    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBornCreationBull is a free log retrieval operation binding the contract event 0x6cadede40797b8dd19173bfb548d15376b7a853d2d84dc523423b656731deaf5.
//
// Solidity: event BornCreationBull(address indexed sender_, uint256 indexed tokenId, uint256 indexed creationId, uint256 life_, uint256 energy_, uint256 grow, uint256 deadTime, uint256 attack_, uint256 defense_, uint256 stamina_)
func (_Cattle *CattleFilterer) FilterBornCreationBull(opts *bind.FilterOpts, sender_ []common.Address, tokenId []*big.Int, creationId []*big.Int) (*CattleBornCreationBullIterator, error) {

	var sender_Rule []interface{}
	for _, sender_Item := range sender_ {
		sender_Rule = append(sender_Rule, sender_Item)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var creationIdRule []interface{}
	for _, creationIdItem := range creationId {
		creationIdRule = append(creationIdRule, creationIdItem)
	}

	logs, sub, err := _Cattle.contract.FilterLogs(opts, "BornCreationBull", sender_Rule, tokenIdRule, creationIdRule)
	if err != nil {
		return nil, err
	}
	return &CattleBornCreationBullIterator{contract: _Cattle.contract, event: "BornCreationBull", logs: logs, sub: sub}, nil
}

// WatchBornCreationBull is a free log subscription operation binding the contract event 0x6cadede40797b8dd19173bfb548d15376b7a853d2d84dc523423b656731deaf5.
//
// Solidity: event BornCreationBull(address indexed sender_, uint256 indexed tokenId, uint256 indexed creationId, uint256 life_, uint256 energy_, uint256 grow, uint256 deadTime, uint256 attack_, uint256 defense_, uint256 stamina_)
func (_Cattle *CattleFilterer) WatchBornCreationBull(opts *bind.WatchOpts, sink chan<- *CattleBornCreationBull, sender_ []common.Address, tokenId []*big.Int, creationId []*big.Int) (event.Subscription, error) {

	var sender_Rule []interface{}
	for _, sender_Item := range sender_ {
		sender_Rule = append(sender_Rule, sender_Item)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var creationIdRule []interface{}
	for _, creationIdItem := range creationId {
		creationIdRule = append(creationIdRule, creationIdItem)
	}

	logs, sub, err := _Cattle.contract.WatchLogs(opts, "BornCreationBull", sender_Rule, tokenIdRule, creationIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CattleBornCreationBull)
				if err := _Cattle.contract.UnpackLog(event, "BornCreationBull", log); err != nil {
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

// ParseBornCreationBull is a log parse operation binding the contract event 0x6cadede40797b8dd19173bfb548d15376b7a853d2d84dc523423b656731deaf5.
//
// Solidity: event BornCreationBull(address indexed sender_, uint256 indexed tokenId, uint256 indexed creationId, uint256 life_, uint256 energy_, uint256 grow, uint256 deadTime, uint256 attack_, uint256 defense_, uint256 stamina_)
func (_Cattle *CattleFilterer) ParseBornCreationBull(log types.Log) (*CattleBornCreationBull, error) {
	event := new(CattleBornCreationBull)
	if err := _Cattle.contract.UnpackLog(event, "BornCreationBull", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CattleBornCreationCowIterator is returned from FilterBornCreationCow and is used to iterate over the raw logs and unpacked data for BornCreationCow events raised by the Cattle contract.
type CattleBornCreationCowIterator struct {
	Event *CattleBornCreationCow // Event containing the contract specifics and raw log

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
func (it *CattleBornCreationCowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CattleBornCreationCow)
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
		it.Event = new(CattleBornCreationCow)
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
func (it *CattleBornCreationCowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CattleBornCreationCowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CattleBornCreationCow represents a BornCreationCow event raised by the Cattle contract.
type CattleBornCreationCow struct {
	Sender     common.Address
	TokenId    *big.Int
	CreationId *big.Int
	Life       *big.Int
	Energy     *big.Int
	Grow       *big.Int
	DeadTime   *big.Int
	Milk       *big.Int
	MilkRate   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBornCreationCow is a free log retrieval operation binding the contract event 0x3be0cc335618acb20037383e05a5f336f1e6aa524c959ad29d1b91cf64772bc0.
//
// Solidity: event BornCreationCow(address indexed sender_, uint256 indexed tokenId, uint256 indexed creationId, uint256 life_, uint256 energy_, uint256 grow, uint256 deadTime, uint256 milk_, uint256 milkRate_)
func (_Cattle *CattleFilterer) FilterBornCreationCow(opts *bind.FilterOpts, sender_ []common.Address, tokenId []*big.Int, creationId []*big.Int) (*CattleBornCreationCowIterator, error) {

	var sender_Rule []interface{}
	for _, sender_Item := range sender_ {
		sender_Rule = append(sender_Rule, sender_Item)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var creationIdRule []interface{}
	for _, creationIdItem := range creationId {
		creationIdRule = append(creationIdRule, creationIdItem)
	}

	logs, sub, err := _Cattle.contract.FilterLogs(opts, "BornCreationCow", sender_Rule, tokenIdRule, creationIdRule)
	if err != nil {
		return nil, err
	}
	return &CattleBornCreationCowIterator{contract: _Cattle.contract, event: "BornCreationCow", logs: logs, sub: sub}, nil
}

// WatchBornCreationCow is a free log subscription operation binding the contract event 0x3be0cc335618acb20037383e05a5f336f1e6aa524c959ad29d1b91cf64772bc0.
//
// Solidity: event BornCreationCow(address indexed sender_, uint256 indexed tokenId, uint256 indexed creationId, uint256 life_, uint256 energy_, uint256 grow, uint256 deadTime, uint256 milk_, uint256 milkRate_)
func (_Cattle *CattleFilterer) WatchBornCreationCow(opts *bind.WatchOpts, sink chan<- *CattleBornCreationCow, sender_ []common.Address, tokenId []*big.Int, creationId []*big.Int) (event.Subscription, error) {

	var sender_Rule []interface{}
	for _, sender_Item := range sender_ {
		sender_Rule = append(sender_Rule, sender_Item)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var creationIdRule []interface{}
	for _, creationIdItem := range creationId {
		creationIdRule = append(creationIdRule, creationIdItem)
	}

	logs, sub, err := _Cattle.contract.WatchLogs(opts, "BornCreationCow", sender_Rule, tokenIdRule, creationIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CattleBornCreationCow)
				if err := _Cattle.contract.UnpackLog(event, "BornCreationCow", log); err != nil {
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

// ParseBornCreationCow is a log parse operation binding the contract event 0x3be0cc335618acb20037383e05a5f336f1e6aa524c959ad29d1b91cf64772bc0.
//
// Solidity: event BornCreationCow(address indexed sender_, uint256 indexed tokenId, uint256 indexed creationId, uint256 life_, uint256 energy_, uint256 grow, uint256 deadTime, uint256 milk_, uint256 milkRate_)
func (_Cattle *CattleFilterer) ParseBornCreationCow(log types.Log) (*CattleBornCreationCow, error) {
	event := new(CattleBornCreationCow)
	if err := _Cattle.contract.UnpackLog(event, "BornCreationCow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CattleBornNormalBullIterator is returned from FilterBornNormalBull and is used to iterate over the raw logs and unpacked data for BornNormalBull events raised by the Cattle contract.
type CattleBornNormalBullIterator struct {
	Event *CattleBornNormalBull // Event containing the contract specifics and raw log

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
func (it *CattleBornNormalBullIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CattleBornNormalBull)
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
		it.Event = new(CattleBornNormalBull)
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
func (it *CattleBornNormalBullIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CattleBornNormalBullIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CattleBornNormalBull represents a BornNormalBull event raised by the Cattle contract.
type CattleBornNormalBull struct {
	Sender   common.Address
	TokenId  *big.Int
	Life     *big.Int
	Energy   *big.Int
	Grow     *big.Int
	Star     *big.Int
	Parents  [2]*big.Int
	DeadTime *big.Int
	Attack   *big.Int
	Defense  *big.Int
	Stamina  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBornNormalBull is a free log retrieval operation binding the contract event 0x18a18192f0826a437b79df65b06727f8310307aea2539b1e90284ee01bcff60b.
//
// Solidity: event BornNormalBull(address indexed sender_, uint256 indexed tokenId, uint256 life_, uint256 energy_, uint256 grow, uint256 star, uint256[2] parents, uint256 deadTime, uint256 attack_, uint256 defense_, uint256 stamina_)
func (_Cattle *CattleFilterer) FilterBornNormalBull(opts *bind.FilterOpts, sender_ []common.Address, tokenId []*big.Int) (*CattleBornNormalBullIterator, error) {

	var sender_Rule []interface{}
	for _, sender_Item := range sender_ {
		sender_Rule = append(sender_Rule, sender_Item)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cattle.contract.FilterLogs(opts, "BornNormalBull", sender_Rule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CattleBornNormalBullIterator{contract: _Cattle.contract, event: "BornNormalBull", logs: logs, sub: sub}, nil
}

// WatchBornNormalBull is a free log subscription operation binding the contract event 0x18a18192f0826a437b79df65b06727f8310307aea2539b1e90284ee01bcff60b.
//
// Solidity: event BornNormalBull(address indexed sender_, uint256 indexed tokenId, uint256 life_, uint256 energy_, uint256 grow, uint256 star, uint256[2] parents, uint256 deadTime, uint256 attack_, uint256 defense_, uint256 stamina_)
func (_Cattle *CattleFilterer) WatchBornNormalBull(opts *bind.WatchOpts, sink chan<- *CattleBornNormalBull, sender_ []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sender_Rule []interface{}
	for _, sender_Item := range sender_ {
		sender_Rule = append(sender_Rule, sender_Item)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cattle.contract.WatchLogs(opts, "BornNormalBull", sender_Rule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CattleBornNormalBull)
				if err := _Cattle.contract.UnpackLog(event, "BornNormalBull", log); err != nil {
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

// ParseBornNormalBull is a log parse operation binding the contract event 0x18a18192f0826a437b79df65b06727f8310307aea2539b1e90284ee01bcff60b.
//
// Solidity: event BornNormalBull(address indexed sender_, uint256 indexed tokenId, uint256 life_, uint256 energy_, uint256 grow, uint256 star, uint256[2] parents, uint256 deadTime, uint256 attack_, uint256 defense_, uint256 stamina_)
func (_Cattle *CattleFilterer) ParseBornNormalBull(log types.Log) (*CattleBornNormalBull, error) {
	event := new(CattleBornNormalBull)
	if err := _Cattle.contract.UnpackLog(event, "BornNormalBull", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CattleBornNormalCowIterator is returned from FilterBornNormalCow and is used to iterate over the raw logs and unpacked data for BornNormalCow events raised by the Cattle contract.
type CattleBornNormalCowIterator struct {
	Event *CattleBornNormalCow // Event containing the contract specifics and raw log

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
func (it *CattleBornNormalCowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CattleBornNormalCow)
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
		it.Event = new(CattleBornNormalCow)
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
func (it *CattleBornNormalCowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CattleBornNormalCowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CattleBornNormalCow represents a BornNormalCow event raised by the Cattle contract.
type CattleBornNormalCow struct {
	Sender   common.Address
	TokenId  *big.Int
	Life     *big.Int
	Energy   *big.Int
	Grow     *big.Int
	Star     *big.Int
	Parents  [2]*big.Int
	DeadTime *big.Int
	Milk     *big.Int
	MilkRate *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBornNormalCow is a free log retrieval operation binding the contract event 0xc94528f8de56d922e972ef26ec9f62079bff7d5376c2ee08e5c5316271cea71a.
//
// Solidity: event BornNormalCow(address indexed sender_, uint256 indexed tokenId, uint256 life_, uint256 energy_, uint256 grow, uint256 star, uint256[2] parents, uint256 deadTime, uint256 milk_, uint256 milkRate_)
func (_Cattle *CattleFilterer) FilterBornNormalCow(opts *bind.FilterOpts, sender_ []common.Address, tokenId []*big.Int) (*CattleBornNormalCowIterator, error) {

	var sender_Rule []interface{}
	for _, sender_Item := range sender_ {
		sender_Rule = append(sender_Rule, sender_Item)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cattle.contract.FilterLogs(opts, "BornNormalCow", sender_Rule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CattleBornNormalCowIterator{contract: _Cattle.contract, event: "BornNormalCow", logs: logs, sub: sub}, nil
}

// WatchBornNormalCow is a free log subscription operation binding the contract event 0xc94528f8de56d922e972ef26ec9f62079bff7d5376c2ee08e5c5316271cea71a.
//
// Solidity: event BornNormalCow(address indexed sender_, uint256 indexed tokenId, uint256 life_, uint256 energy_, uint256 grow, uint256 star, uint256[2] parents, uint256 deadTime, uint256 milk_, uint256 milkRate_)
func (_Cattle *CattleFilterer) WatchBornNormalCow(opts *bind.WatchOpts, sink chan<- *CattleBornNormalCow, sender_ []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sender_Rule []interface{}
	for _, sender_Item := range sender_ {
		sender_Rule = append(sender_Rule, sender_Item)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cattle.contract.WatchLogs(opts, "BornNormalCow", sender_Rule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CattleBornNormalCow)
				if err := _Cattle.contract.UnpackLog(event, "BornNormalCow", log); err != nil {
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

// ParseBornNormalCow is a log parse operation binding the contract event 0xc94528f8de56d922e972ef26ec9f62079bff7d5376c2ee08e5c5316271cea71a.
//
// Solidity: event BornNormalCow(address indexed sender_, uint256 indexed tokenId, uint256 life_, uint256 energy_, uint256 grow, uint256 star, uint256[2] parents, uint256 deadTime, uint256 milk_, uint256 milkRate_)
func (_Cattle *CattleFilterer) ParseBornNormalCow(log types.Log) (*CattleBornNormalCow, error) {
	event := new(CattleBornNormalCow)
	if err := _Cattle.contract.UnpackLog(event, "BornNormalCow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CattleGrowUpIterator is returned from FilterGrowUp and is used to iterate over the raw logs and unpacked data for GrowUp events raised by the Cattle contract.
type CattleGrowUpIterator struct {
	Event *CattleGrowUp // Event containing the contract specifics and raw log

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
func (it *CattleGrowUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CattleGrowUp)
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
		it.Event = new(CattleGrowUp)
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
func (it *CattleGrowUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CattleGrowUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CattleGrowUp represents a GrowUp event raised by the Cattle contract.
type CattleGrowUp struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGrowUp is a free log retrieval operation binding the contract event 0x27308be09b242076a36ef5c74033829078220c467a50c92f69adf8c8f93f1f65.
//
// Solidity: event GrowUp(uint256 indexed tokenId)
func (_Cattle *CattleFilterer) FilterGrowUp(opts *bind.FilterOpts, tokenId []*big.Int) (*CattleGrowUpIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cattle.contract.FilterLogs(opts, "GrowUp", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CattleGrowUpIterator{contract: _Cattle.contract, event: "GrowUp", logs: logs, sub: sub}, nil
}

// WatchGrowUp is a free log subscription operation binding the contract event 0x27308be09b242076a36ef5c74033829078220c467a50c92f69adf8c8f93f1f65.
//
// Solidity: event GrowUp(uint256 indexed tokenId)
func (_Cattle *CattleFilterer) WatchGrowUp(opts *bind.WatchOpts, sink chan<- *CattleGrowUp, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cattle.contract.WatchLogs(opts, "GrowUp", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CattleGrowUp)
				if err := _Cattle.contract.UnpackLog(event, "GrowUp", log); err != nil {
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

// ParseGrowUp is a log parse operation binding the contract event 0x27308be09b242076a36ef5c74033829078220c467a50c92f69adf8c8f93f1f65.
//
// Solidity: event GrowUp(uint256 indexed tokenId)
func (_Cattle *CattleFilterer) ParseGrowUp(log types.Log) (*CattleGrowUp, error) {
	event := new(CattleGrowUp)
	if err := _Cattle.contract.UnpackLog(event, "GrowUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CattleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cattle contract.
type CattleOwnershipTransferredIterator struct {
	Event *CattleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CattleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CattleOwnershipTransferred)
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
		it.Event = new(CattleOwnershipTransferred)
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
func (it *CattleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CattleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CattleOwnershipTransferred represents a OwnershipTransferred event raised by the Cattle contract.
type CattleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cattle *CattleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CattleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cattle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CattleOwnershipTransferredIterator{contract: _Cattle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cattle *CattleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CattleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cattle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CattleOwnershipTransferred)
				if err := _Cattle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Cattle *CattleFilterer) ParseOwnershipTransferred(log types.Log) (*CattleOwnershipTransferred, error) {
	event := new(CattleOwnershipTransferred)
	if err := _Cattle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CattleTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cattle contract.
type CattleTransferIterator struct {
	Event *CattleTransfer // Event containing the contract specifics and raw log

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
func (it *CattleTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CattleTransfer)
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
		it.Event = new(CattleTransfer)
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
func (it *CattleTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CattleTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CattleTransfer represents a Transfer event raised by the Cattle contract.
type CattleTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cattle *CattleFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CattleTransferIterator, error) {

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

	logs, sub, err := _Cattle.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CattleTransferIterator{contract: _Cattle.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cattle *CattleFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CattleTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Cattle.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CattleTransfer)
				if err := _Cattle.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Cattle *CattleFilterer) ParseTransfer(log types.Log) (*CattleTransfer, error) {
	event := new(CattleTransfer)
	if err := _Cattle.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CattleUpGradeStarIterator is returned from FilterUpGradeStar and is used to iterate over the raw logs and unpacked data for UpGradeStar events raised by the Cattle contract.
type CattleUpGradeStarIterator struct {
	Event *CattleUpGradeStar // Event containing the contract specifics and raw log

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
func (it *CattleUpGradeStarIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CattleUpGradeStar)
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
		it.Event = new(CattleUpGradeStar)
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
func (it *CattleUpGradeStarIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CattleUpGradeStarIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CattleUpGradeStar represents a UpGradeStar event raised by the Cattle contract.
type CattleUpGradeStar struct {
	TokenId *big.Int
	Stars   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpGradeStar is a free log retrieval operation binding the contract event 0xc345ee01e5afbb825f5ac65a2f0f2fa781c3c3498ac6c90c8d908009b12b31d8.
//
// Solidity: event UpGradeStar(uint256 indexed tokenId, uint256 indexed stars)
func (_Cattle *CattleFilterer) FilterUpGradeStar(opts *bind.FilterOpts, tokenId []*big.Int, stars []*big.Int) (*CattleUpGradeStarIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var starsRule []interface{}
	for _, starsItem := range stars {
		starsRule = append(starsRule, starsItem)
	}

	logs, sub, err := _Cattle.contract.FilterLogs(opts, "UpGradeStar", tokenIdRule, starsRule)
	if err != nil {
		return nil, err
	}
	return &CattleUpGradeStarIterator{contract: _Cattle.contract, event: "UpGradeStar", logs: logs, sub: sub}, nil
}

// WatchUpGradeStar is a free log subscription operation binding the contract event 0xc345ee01e5afbb825f5ac65a2f0f2fa781c3c3498ac6c90c8d908009b12b31d8.
//
// Solidity: event UpGradeStar(uint256 indexed tokenId, uint256 indexed stars)
func (_Cattle *CattleFilterer) WatchUpGradeStar(opts *bind.WatchOpts, sink chan<- *CattleUpGradeStar, tokenId []*big.Int, stars []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var starsRule []interface{}
	for _, starsItem := range stars {
		starsRule = append(starsRule, starsItem)
	}

	logs, sub, err := _Cattle.contract.WatchLogs(opts, "UpGradeStar", tokenIdRule, starsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CattleUpGradeStar)
				if err := _Cattle.contract.UnpackLog(event, "UpGradeStar", log); err != nil {
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

// ParseUpGradeStar is a log parse operation binding the contract event 0xc345ee01e5afbb825f5ac65a2f0f2fa781c3c3498ac6c90c8d908009b12b31d8.
//
// Solidity: event UpGradeStar(uint256 indexed tokenId, uint256 indexed stars)
func (_Cattle *CattleFilterer) ParseUpGradeStar(log types.Log) (*CattleUpGradeStar, error) {
	event := new(CattleUpGradeStar)
	if err := _Cattle.contract.UnpackLog(event, "UpGradeStar", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
