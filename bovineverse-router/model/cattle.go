package model

import (
	"github.com/lib/pq"
)

type CattleInfo struct {
	Id        uint32        `gorm:"column:id"`
	Class     uint8         `gorm:"column:class"`
	Gender    uint8         `gorm:"column:gender"`
	GenderSeq uint32        `gorm:"column:gender_seq"`
	Life      uint32        `gorm:"column:life"`
	Growth    uint32        `gorm:"column:growth"`
	Energy    uint32        `gorm:"column:energy"`
	IsAdult   bool          `gorm:"column:is_adult"`
	Star      uint8         `gorm:"column:star"`
	Attack    uint32        `gorm:"column:attack"`
	Stamina   uint32        `gorm:"column:stamina" `
	Defense   uint32        `gorm:"column:defense" `
	Milk      uint32        `gorm:"column:milk" `
	MilkRate  uint32        `gorm:"column:milk_rate" `
	Owner     string        `gorm:"column:owner"`
	DeadAt    int64         `gorm:"column:dead_at"`
	Parents   pq.Int32Array `gorm:"column:parents; type:int4[]"`
}

func (*CattleInfo) TableName() string {
	return "u_cattle"
}

type CompoundInfo struct {
	Id        uint32 `gorm:"column:id"`
	Address   string `gorm:"column:address"`
	TokenId   uint32 `gorm:"column:token_id"`
	BurnId    uint32 `gorm:"column:burn_id"`
	BlockTime int64  `gorm:"column:block_time"`
}

func (*CompoundInfo) TableName() string {
	return "r_cattle_compound"
}

type BurnCattleInfo struct {
	Id       uint32        `gorm:"column:id"`
	Gender   uint8         `gorm:"column:gender"`
	Life     uint32        `gorm:"column:life"`
	Growth   uint32        `gorm:"column:growth"`
	Energy   uint32        `gorm:"column:energy"`
	Star     uint8         `gorm:"column:star"`
	Attack   uint32        `gorm:"column:attack"`
	Stamina  uint32        `gorm:"column:stamina" `
	Defense  uint32        `gorm:"column:defense" `
	Milk     uint32        `gorm:"column:milk" `
	MilkRate uint32        `gorm:"column:milk_rate" `
	Parents  pq.Int32Array `gorm:"column:parents; type:int4[]"`
	Owner    string        `gorm:"column:owner"`
	BurnWay  uint8         `gorm:"column:burn_way"`
	BurnTime int64         `gorm:"column:burn_time"`
}

func (*BurnCattleInfo) TableName() string {
	return "r_burn_cattle"
}

func NewBurnCattle(data *CattleInfo) *BurnCattleInfo {
	return &BurnCattleInfo{
		Id:       data.Id,
		Gender:   data.Gender,
		Star:     data.Star,
		Life:     data.Life,
		Growth:   data.Growth,
		Energy:   data.Energy,
		Attack:   data.Attack,
		Defense:  data.Defense,
		Stamina:  data.Stamina,
		Milk:     data.Milk,
		MilkRate: data.MilkRate,
		Parents:  data.Parents,
		Owner:    data.Owner,
	}
}
