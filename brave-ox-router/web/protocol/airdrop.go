package protocol

type AirdropRequest struct {
	Address string `form:"address"`
	// AirdropType int    `form:"airdrop_type"`
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type AirdropResponse struct {
	Id          uint32 `gorm:"column:id" json:"id"`
	AirdropType uint8  `gorm:"column:type" json:"type"`
	TitleId     int    `gorm:"column:title_id" json:"-"`
	Title       string `gorm:"column:-" json:"title"`
	Wechat      string `gorm:"column:wechat" json:"wechat"`
	Address     string `gorm:"column:address" json:"address"`
	Category    uint8  `gorm:"column:category" json:"category"`
	FromId      uint8  `gorm:"column:from_id" json:"from_id"`
	Amount      int64  `gorm:"column:amount" json:"amount"`
	Status      int    `gorm:"column:status" json:"status"`
}

type AirdropSign struct {
	ID       uint32 `json:"id"`
	Category uint8  `json:"category"`
	FromId   uint8  `json:"from_id"`
	CardId   uint32 `json:"card_id"`
	Amount   uint32 `json:"amount"`
	R        string `json:"r"`
	S        string `json:"s"`
	V        uint8  `json:"v"`
}
