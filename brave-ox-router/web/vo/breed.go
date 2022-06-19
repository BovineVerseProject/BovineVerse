package vo

import "brave-ox-web/model"

type BreedCattleVO struct {
	TokenId   uint32  `json:"tokenId"`
	Price     float64 `json:"price"`
	Seller    string  `json:"seller"`
	Class     uint8   `json:"class"`
	Gender    uint8   `json:"gender"`
	GenderSeq uint32  `json:"gsq"`
	Life      uint32  `json:"life"`
	Growth    uint32  `json:"growth"`
	Energy    uint32  `json:"energy"`
	IsAdult   bool    ` json:"isAdult"`
	Star      uint8   ` json:"star"`
	Attack    uint32  ` json:"attack"`
	Stamina   uint32  ` json:"stamina"`
	Defense   uint32  ` json:"defense"`
	Milk      uint32  ` json:"milk"`
	MilkRate  uint32  ` json:"milkRate"`
	Count     uint32  `json:"count"`
}

func NewBreedCenterVO(info *model.BreedCattleInfo) *BreedCattleVO {
	if info == nil {
		return nil
	}
	return &BreedCattleVO{
		TokenId:   info.TokenId,
		Price:     info.Price,
		Seller:    info.Seller,
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

}

func NewBreedCenterVOList(info []*model.BreedCattleInfo) []*BreedCattleVO {
	l := len(info)
	if l == 0 {
		return nil
	}
	vo := make([]*BreedCattleVO, l)
	for index, v := range info {
		vo[index] = NewBreedCenterVO(v)
	}
	return vo
}
