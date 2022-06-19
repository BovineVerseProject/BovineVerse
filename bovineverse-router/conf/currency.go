package conf

import (
	"fmt"

	"brave-ox-web/common"
	"brave-ox-web/db"
	"brave-ox-web/model"
)

var _currency map[uint8]*model.CurrencyInfo

func init() {
	_currency = make(map[uint8]*model.CurrencyInfo)
}

func LoadAllCurrency() {
	var currency []*model.CurrencyInfo
	err := db.QueryAll(&currency)
	common.MustNoError(err)

	if len(currency) <= 0 {
		panic(fmt.Errorf("no currency config table"))
	}

	for _, v := range currency {
		_currency[v.Types] = v
	}
}

func AllCurrency() map[uint8]*model.CurrencyInfo {
	return _currency
}

func GetCurrencyBasicInfo(types uint8) *model.CurrencyInfo {
	return _currency[types]
}
