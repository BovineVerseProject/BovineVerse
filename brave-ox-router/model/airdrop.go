package model

type AirdropTitle struct {
	Id       uint32 `gorm:"column:id"`
	TitleCh  string `gorm:"column:title_ch; type:varchar(255); NOT NULL"`
	TitleEn  string `gorm:"column:title_en; type:varchar(255)"`
	TitleKr  string `gorm:"column:title_kr; type:varchar(255)"`
	DescCh   string `gorm:"column:desc_ch; type:varchar(255)"`
	DescEn   string `gorm:"column:desc_en; type:varchar(255)"`
	DescKr   string `gorm:"column:desc_kr; type:varchar(255)"`
	Begin    int64  `gorm:"column:begin; NOT NULL"`
	End      int64  `gorm:"column:end; NOT NULL"`
	Deadline bool   `gorm:"column:deadline; NOT NULL; default:0;"`
	Frozen   bool   `gorm:"column:frozen; NOT NULL; default:0;"`

	Chain string `gorm:"-"`
}

func (*AirdropTitle) TableName() string {
	return "s_airdrop_title"
}

type AirdropReward struct {
	Id       uint32 `gorm:"column:id"`
	Type     uint8  `gorm:"column:type; NOT NULL; comment:1-usdt; 2-bvg; 3-bvt;"`
	Title    string `gorm:"-"`
	TitleId  uint32 `gorm:"column:title_id; NOT NULL;"`
	Wechat   string `gorm:"column:wechat; type:varchar(255);"`
	Address  string `gorm:"column:address; type:varchar(255); NOT NULL;"`
	Category uint8  `gorm:"column:category; NOT NULL; comment:1-erc20; 2-erc1155; 3-erc721;"`
	FromId   uint8  `gorm:"column:from_id; NOT NULL; comment:1-usdt; 2-bvg; 3-bvt;"`
	Amount   uint32 `gorm:"column:amount; NOT NULL"`
	CardId   uint32 `gorm:"column:card_id; NOT NULL; default:0;"`
	Status   int    `gorm:"column:status; NOT NULL; default:0; comment: -1: 领取中; 0: 待领取; 1:已领取;"`
}

func (*AirdropReward) TableName() string {
	return "u_airdrop_reward"
}

type AirdropUser struct {
	Addr   string
	Amount uint
}

type Airdrop struct {
	Users []AirdropUser
	Title AirdropTitle
}

type AirdropInfo struct {
	Item []Airdrop
	Usdt []Airdrop
	Bvg  []Airdrop
	Bvt  []Airdrop
}
