package protocol

import (
	"brave-ox-web/web/vo"
)

type MarketCattleRequest struct {
	Page   uint32   `json:"page" binding:"number,omitempty"`
	Sort   uint32   `json:"sort" binding:"oneof=0 1 2"`
	Class  uint32   `json:"class" binding:"max=2,omitempty"`
	Gender uint32   `json:"gender" binding:"max=2,omitempty"`
	Life   []uint32 `json:"life" binding:"lte=2,omitempty"`
	Growth []uint32 `json:"growth" binding:"lte=2,omitempty"`
	Energy []uint32 `json:"energy" binding:"lte=2,omitempty"`

	Attack  []uint32 `json:"attack" binding:"lte=2,omitempty"`
	Stamina []uint32 `json:"stamina" binding:"lte=2,omitempty"`
	Defence []uint32 `json:"defence" binding:"lte=2,omitempty"`

	Milk     []uint32 `json:"milk" binding:"lte=2,omitempty"`
	MilkRate []uint32 `json:"milkRate" binding:"lte=2,omitempty"`
}

type MarketCattleResponse struct {
	GoodsType uint8 `json:"goodsType"`
	PageResponse
	List []*vo.MarketCattleInfoVO `json:"list"`
}

type PageResponse struct {
	Total uint32 `json:"total"`
	Page  uint32 `json:"page"`
}

type MarketRequest struct {
	Page uint32 `json:"page" binding:"number,omitempty"`
	Sort uint32 `json:"sort" binding:"oneof=0 1 2"`
}

type MarketMysteryBoxResponse struct {
	GoodsType uint8 `json:"goodsType"`
	PageResponse
	List []*vo.MarketSaleVO `json:"list"`
}

type MarketPlanetResponse struct {
	GoodsType uint8 `json:"goodsType"`
	PageResponse
	List []*vo.MarketPlanetVO `json:"list"`
}

type MarketHistoryRequest struct {
	Page uint32 `json:"page"`
}

type MarketHistoryResponse struct {
	PageResponse
	History []*vo.MarketHistoryVO `json:"history"`
}
