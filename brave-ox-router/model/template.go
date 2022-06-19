package model

type CattleTemplate struct {
	Id          uint32 `gorm:"column:id"`
	Class       uint8  `gorm:"column:class"`
	Gender      uint8  `gorm:"column:gender"`
	Gsq         uint32 `gorm:"column:gsq"`
	IsAdult     bool   `gorm:"column:is_adult"`
	Image       string `gorm:"column:image"`
	Profile     uint32 `gorm:"column:profile"`
	NameCN      string `gorm:"column:name_cn"`
	NameEn      string `gorm:"column:name_en"`
	Description string `gorm:"column:description"`
}

func (*CattleTemplate) TableName() string {
	return "s_cattle"
}

type ProfilePhotoTemplate struct {
	Id    uint32 `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Image string `gorm:"column:image"`
}

func (*ProfilePhotoTemplate) TableName() string {
	return "s_profile_photo"
}

type CattleSkinTemplate struct {
	Id          uint32 `gorm:"column:id"`
	NameEn      string `gorm:"column:name_en"`
	Level       string `gorm:"column:level"`
	Attack      uint32 `gorm:"column:attack"`
	Defense     uint32 `gorm:"column:defense"`
	Stamina     uint32 `gorm:"column:stamina"`
	Life        uint32 `gorm:"column:life"`
	Image       string `gorm:"column:image"`
	Description string `gorm:"column:description"`
}

func (*CattleSkinTemplate) TableName() string {
	return "s_cattle_skin"
}

type PlanetTemplate struct {
	Types       uint8  `gorm:"column:types"`
	NameEn      string `gorm:"column:name_en"`
	Image       string `gorm:"column:image"`
	Description string `gorm:"column:description"`
}

func (*PlanetTemplate) TableName() string {
	return "s_planet"
}

type BoxTemplate struct {
	Id          uint32 `gorm:"column:id"`
	Types       uint32 `gorm:"column:types"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Image       string `gorm:"column:image"`
	Contract    string `gorm:"column:contract"`
}

func (*BoxTemplate) TableName() string {
	return "s_mystery_box"
}
