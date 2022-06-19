package utils

import (
	"fmt"
	"math/big"
	"math/rand"
	"runtime"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"

	"brave-ox-web/log"
)

func toHexInt(n *big.Int) string {
	return fmt.Sprintf("%#x", n) // or %X or upper case
}

func ParseAddress(data string) (string, bool) {
	data = strings.ToUpper(data)
	var addr string
	if strings.HasPrefix(data, "X0X") {
		addr = strings.TrimPrefix(data, "X")
	} else {
		info, _ := new(big.Int).SetString(data, 10)
		addr = toHexInt(info)
	}

	if !common.IsHexAddress(addr) {
		return addr, false
	}
	return common.HexToAddress(addr).Hex(), true

}

func PrintPanic() {
	if err := recover(); err != nil {
		var buf = make([]byte, 1024)
		runtime.Stack(buf[:], false)
		log.Errorf("panic: %v, %v \n", err, string(buf))
	}
}

func GetOneEndpoint(endpoints []string) string {
	n := rand.Int() % len(endpoints)
	return endpoints[n]
}

func GetCheckSum(address string) string {
	return common.HexToAddress(ethEnsure0x(address)).Hex()
}

func ethEnsure0x(str string) string {
	if strings.HasPrefix(str, "0x") || strings.HasPrefix(str, "0X") {
		return str
	}
	return "0x" + str
}

func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}
	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)
	return R, S, V
}

func DivRound(d decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d.DivRound(d2, int32(18))
}

func SigWithData(datas [][]byte, srcKey string) (error, string) {
	key, err := crypto.HexToECDSA(srcKey)
	if err != nil {
		return err, ""
	}

	digest := crypto.Keccak256Hash(datas...)
	sig, err := crypto.Sign(digest.Bytes(), key)
	if err != nil {
		return err, ""
	}
	return nil, string(sig)
}
