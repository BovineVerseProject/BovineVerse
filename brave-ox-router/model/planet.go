package model

import "time"

type BasicPlanet struct {
	Id         uint32    `gorm:"column:id"`
	PlanetType uint32    `gorm:"column:planet_type"`
	BlockTime  time.Time `gorm:"column:block_time"`
}

func (*BasicPlanet) TableName() string {
	return "u_planet721"
}

type PlanetInfo struct {
	Id             uint32 `gorm:"column:id"`
	PlanetType     uint32 `gorm:"column:planet_type"`
	Name           string `gorm:"column:name"`
	Owner          string `gorm:"column:master"`
	Notice         string `gorm:"column:notice"`
	Fee            string `gorm:"column:fee"`
	UpdateNameAt   int64  `gorm:"column:update_name_at"`
	UpdateNoticeAt int64  `gorm:"column:update_notice_at"`
	ParentId       uint32 `gorm:"column:parent_id"`
	CreateAt       int64  `gorm:"column:create_at"`
	Image          uint16 `gorm:"column:image"`
	TotalScore     string `gorm:"column:total_score"`
	Population     uint16 `gorm:"column:population"`

	OwnerName string `gorm:"-"`
}

func (*PlanetInfo) TableName() string {
	return "u_planet"
}

type PlanetMember struct {
	Id       uint32 `gorm:"column:id"`
	PlanetId uint32 `gorm:"column:planet_id"`
	Player   string `gorm:"column:player"`
	Position uint8  `gorm:"column:position"`
	CreateAt int64  `gorm:"column:create_at"`

	Name   string `gorm:"-"`
	Avatar string `gorm:"-"`
	Count  uint32 `gorm:"-"`
	Score  string `gorm:"-"`
}

func (*PlanetMember) TableName() string {
	return "u_planet_member"
}

type MemberGroup struct {
	Position uint8  `gorm:"column:position"`
	Count    uint32 `gorm:"column:people"`
}
