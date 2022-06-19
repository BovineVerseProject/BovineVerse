package vo

import (
	"brave-ox-web/model"
)

type CattleVO struct {
	Id        uint32  `json:"id" `
	Name      string  `json:"name"`
	Parents   []int32 `json:"parents"`
	Class     uint8   `json:"class"`
	Gender    uint8   `json:"gender"`
	GenderSeq uint32  `json:"gsq"`
	Life      uint32  ` json:"life"`
	Growth    uint32  `json:"growth"`
	Energy    uint32  `json:"energy"`
	IsAdult   bool    ` json:"isAdult"`
	Star      uint8   ` json:"star"`
	Attack    uint32  ` json:"attack"`
	Stamina   uint32  ` json:"stamina"`
	Defense   uint32  ` json:"defense"`
	Milk      uint32  ` json:"milk"`
	MilkRate  uint32  ` json:"milkRate"`
	Image     string  `json:"image"`
}

func NewCattleVO(info *model.CattleInfo) *CattleVO {
	if info == nil {
		return nil
	}
	vo := &CattleVO{
		Id:        info.Id,
		Parents:   []int32(info.Parents),
		Class:     info.Class,
		Gender:    info.Gender,
		GenderSeq: info.GenderSeq,
		Life:      info.Life,
		Growth:    info.Growth,
		Energy:    info.Energy,
		IsAdult:   info.IsAdult,
		Star:      info.Star,
		Attack:    info.Attack,
		Stamina:   info.Stamina,
		Defense:   info.Defense,
		Milk:      info.Milk,
		MilkRate:  info.MilkRate,
	}

	return vo
}

func NewCattleVOList(list []*model.CattleInfo) []*CattleVO {
	l := len(list)
	if l == 0 {
		return nil
	}

	vo := make([]*CattleVO, l)
	for index, v := range list {
		vo[index] = NewCattleVO(v)
	}

	return vo
}
