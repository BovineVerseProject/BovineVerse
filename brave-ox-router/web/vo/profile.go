package vo

import "brave-ox-web/model"

type ProfileVO struct {
	//Address  string `gorm:"column:address"  json:"-"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar" `
	Gender      uint8  `json:"gender"`
	PlanetImage int8   `json:"planetImage"`
	PlanetName  string `json:"planetName"`

	Relationship []*RelationshipVO `json:"relationship"`
}

func NewProfileVO(info *model.ProfileInfo) *ProfileVO {
	if info == nil {
		return nil
	}

	vo := &ProfileVO{
		Name:        info.Name,
		Avatar:      info.Avatar,
		Gender:      info.Gender,
		PlanetName:  info.PlanetName,
		PlanetImage: info.PlanetImage,
	}

	return vo
}

type ProfileIconVO struct {
	Avatar string `json:"avatar"`
	URI    string `json:"uri"`
}

type GameProfileVO struct {
	//Address  string `gorm:"column:address"  json:"-"`
	Name       string `json:"name"`
	Gender     uint8  `json:"gender"`
	Avatar     string `json:"avatar" `
	Num        uint32 `json:"num"`
	DeadCount  uint32 `json:"deadCount"`
	PlanetName string `json:"planetName"`
}

type GameProfileSummary struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Gender  uint8  `json:"gender"`
	Avatar  string `json:"avatar"`
}

type RelationshipVO struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Gender  uint8  `json:"gender"`
	Avatar  string `json:"avatar"`
	Types   uint8  `json:"types"`
}
