package model

import (
	"github.com/lib/pq"
)

type GameRewardInfo struct {
	Id         uint64         `gorm:"column:id"`
	Player     string         `gorm:"column:player"`
	Types      uint8          `gorm:"column:types"`
	TypeSeq    uint32         `gorm:"column:param_id"`
	RewardType pq.Int32Array  `gorm:"column:reward_type; type:int4[]"`
	ItemId     pq.Int32Array  `gorm:"column:item_id; type:int4[]"`
	Count      pq.StringArray `gorm:"column:count; type:text[]"`
	CreateAt   int64          `gorm:"column:create_at"`
	FinishedAt int64          `gorm:"column:finished_at"`
}

func (*GameRewardInfo) TableName() string {
	return "r_game_reward"
}
