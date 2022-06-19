package model

import (
	"time"

	"github.com/lib/pq"
)

type ProfileInfo struct {
	Address   string `gorm:"column:address"`
	Name      string `gorm:"column:name"`
	Avatar    string `gorm:"column:avatar"`
	UpdateNum uint32 `gorm:"column:update_num"`
	Gender    uint8  `gorm:"column:gender"`
	UsedFree  bool   `gorm:"column:used_free"`
	UpdateAt  int64  `gorm:"update_at"`

	PlanetName  string `gorm:"-"`
	PlanetImage int8   `gorm:"-"`
}

func (*ProfileInfo) TableName() string {
	return "u_profile"
}

type RelationshipInfo struct {
	Id       uint32         `gorm:"column:id"`
	Pair     pq.StringArray `gorm:"column:pair; type:varchar[]"`
	Types    uint8          `gorm:"column:types"`
	CreateAt time.Time      `gorm:"column:create_at"`
}

func (*RelationshipInfo) TableName() string {
	return "u_relationship"
}
