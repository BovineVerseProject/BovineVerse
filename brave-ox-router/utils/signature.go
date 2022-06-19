package utils

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func ecRecover_raw(msg string, sigStr string, prefix string) (common.Address, error) {
	sig, err := hexutil.Decode(sigStr)
	if err != nil {
		return common.Address{}, err
	}

	var msgBytes []byte
	if prefix == "eth" {
		msg = fmt.Sprint("\x19Ethereum Signed Message:\n", len(msg), msg) // for web3
		msgBytes = crypto.Keccak256([]byte(msg))
	} else if prefix == "cfx" {
		msgBytes = []byte("0x" + common.Bytes2Hex([]byte(msg)))
		pre := fmt.Sprint("\x19Conflux Signed Message:\n", len("0x"+common.Bytes2Hex([]byte(msg)))) // for conflux portal
		preBytes := []byte(pre)
		msgBytes = append(preBytes, msgBytes...)

		msgBytes = crypto.Keccak256(msgBytes)
	}

	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("wrong sig length")
	}
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}, fmt.Errorf("wrong sig v")
	}

	sig[64] -= 27

	pubKey, err := crypto.SigToPub(msgBytes, sig)

	if err != nil {
		return common.Address{}, err
	}
	addr := crypto.PubkeyToAddress(*pubKey)
	if prefix == "cfx" {
		addr[0] = addr[0]&0x1f | 0x10
	}

	return addr, nil
}

func ecRecover_typed(msg string, sigStr string) (common.Address, error) {
	// hash data
	res1 := crypto.Keccak256Hash(
		[]byte("string Message"),
	).Bytes()
	res2 := crypto.Keccak256Hash(
		[]byte(msg),
	).Bytes()
	msgBytes := crypto.Keccak256Hash(
		common.RightPadBytes(res1, 32),
		common.RightPadBytes(res2, 32),
	).Bytes()

	// recover
	sig, err := hexutil.Decode(sigStr)
	if err != nil {
		return common.Address{}, err
	}
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("EcRecover_typed wrong sig length")
	}
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}, fmt.Errorf("EcRecover_typed wrong sig length")
	}
	sig[64] -= 27 // Transform yellow paper V from 27/28 to 0/1

	pubKey, err := crypto.SigToPub(msgBytes, sig)
	//ok := crypto.VerifySignature(sig)
	if err != nil {
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(*pubKey), nil
}

func VerifySign(msg string, sign string, address string) (err error) {
	sig := strings.Split(sign, " ")
	if len(sig) != 2 {
		return fmt.Errorf("miss wallettype for %s .", address)
	}
	var resultAddr common.Address
	var resultAddr2 common.Address

	var prefix string
	if sig[0] == "conflux" {
		prefix = "cfx"
	} else {
		prefix = "eth"
	}

	resultAddr, err = ecRecover_raw(msg, sig[1], prefix)
	if err != nil {
		return fmt.Errorf("cRecover err for %s .", address)
	}
	resultAddr2, err = ecRecover_typed(msg, sig[1])
	if err != nil {
		return fmt.Errorf("cRecover err for %s .", address)
	}

	if common.HexToAddress(address).Hex() != resultAddr.Hex() && common.HexToAddress(address).Hex() != resultAddr2.Hex() {
		return fmt.Errorf("invalid signature for %s.", address)
	}
	return nil
}
