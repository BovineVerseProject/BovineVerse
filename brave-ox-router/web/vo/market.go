package vo

import "brave-ox-web/model"

type UsdtRate struct {
	Price       float64 `json:"price"`
	CurrentUSDT float64 `json:"usdt"`
}

type MarketSaleVO struct {
	TokenId   uint32 `json:"id"`
	Holder    string `json:"holder"`
	TradeType uint8  `json:"tradeType"`
	UsdtRate
}

type MarketCattleInfoVO struct {
	TokenId   uint32 `json:"id"`
	Class     uint8  `json:"class"`
	Name      string `json:"name"`
	Gender    uint8  `json:"gender"`
	TradeType uint8  `json:"tradeType"`
	UsdtRate
	Image  string `json:"image"`
	Holder string `json:"holder"`
}

func NewMarketCattleVO(info *model.MarketCattleInfo) *MarketCattleInfoVO {
	if info == nil {
		return nil
	}
	obj := &MarketCattleInfoVO{}
	obj.TokenId = info.TokenId
	obj.Holder = info.Seller
	obj.Gender = info.Gender
	obj.TradeType = info.TradeType
	obj.Class = info.Class
	obj.Image = info.Image
	obj.Price = info.Price
	return obj
}

func NewMarketSaleVO(tokenID uint32, tradeType uint8, seller string, price float64, uPrice float64) *MarketSaleVO {
	obj := &MarketSaleVO{}
	obj.TokenId = tokenID
	obj.TradeType = tradeType
	obj.Holder = seller
	obj.Price = price
	obj.CurrentUSDT = uPrice
	return obj
}

type MarketHistoryVO struct {
	GoodsType uint8   `json:"goodsType"`
	TokenId   uint32  `json:"tokenId"`
	Name      string  `json:"name"`
	TradeType uint8   `json:"tradeType"`
	Price     float64 `json:"price"`
	USDT      float64 `json:"usdt"`
	Status    uint8   `json:"status"`
	DealTime  string  `json:"dealTime"`
}

func NewMarketHistoryVO(info *model.SaleGoodsRecord) *MarketHistoryVO {
	if info == nil {
		return nil
	}

	return &MarketHistoryVO{
		TokenId:   info.TokenId,
		GoodsType: info.GoodsType,
		TradeType: info.TradeType,
		Price:     info.Price,
	}
}

type MarketPlanetVO struct {
	TokenId   uint32 `json:"id"`
	Types     uint8  `json:"types"`
	TradeType uint8  `json:"tradeType"`
	Holder    string `json:"holder"`
	UsdtRate
}
