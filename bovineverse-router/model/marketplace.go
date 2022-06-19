package model

type MarketSaleInfo struct {
	Id        uint32  `gorm:"column:id"`
	GoodsType uint8   `gorm:"column:goods_type"`
	TokenId   uint32  `gorm:"column:token_id"`
	TradeType uint8   `gorm:"column:trade_type"`
	Price     float64 `gorm:"column:price"`
	UPrice    float64 `gorm:"column:u_price"`
	//Image     string  `gorm:"column:image"`
	BlockTime int64  `gorm:"column:block_time" `
	Seller    string `gorm:"column:seller"`
}

func (*MarketSaleInfo) TableName() string {
	return "u_market"
}

type SaleGoodsRecord struct {
	Id        uint32  `gorm:"column:id"`
	GoodsType uint8   `gorm:"goods_type"`
	TokenId   uint32  `gorm:"column:token_id"`
	TradeType uint8   `gorm:"column:trade_type"`
	Price     float64 `gorm:"column:price"`
	Seller    string  `gorm:"column:seller"`
	Buyer     string  `gorm:"column:buyer"`
	BlockTime int64   `gorm:"column:block_time"`
}

func (*SaleGoodsRecord) TableName() string {
	return "r_market"
}

type MarketGoods struct {
	Id        uint32 `gorm:"column:id"`
	GoodsType uint8  `gorm:"column:goods_type"`
	GoodsName string `gorm:"column:goods_name"`
	Place     string `gorm:"column:place"`
}

func (*MarketGoods) TableName() string {
	return "s_market_goods"
}

type MarketCattleInfo struct {
	TokenId   uint32  `gorm:"column:token_id"`
	TradeType uint8   `gorm:"column:trade_type"`
	Price     float64 `gorm:"column:price"`
	Seller    string  `gorm:"seller"`
	UPrice    float64 `gorm:"column:u_price"`

	Class     uint8  `gorm:"column:class" `
	Gender    uint8  `gorm:"column:gender" `
	GenderSeq uint32 `gorm:"column:gender_seq" `
	IsAdult   bool   `gorm:"column:is_adult"`
	Star      uint8  `gorm:"column:star"`
	Image     string `gorm:"column:-"`
}

type MarketPlanetInfo struct {
	Types     uint8
	TokenId   uint32
	TradeType uint8
	Price     float64
	UPrice    float64
	Seller    string
}
