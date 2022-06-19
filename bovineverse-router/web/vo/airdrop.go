package vo

import (
	"brave-ox-web/model"
)

type AirdropVO struct {
	Id       uint32 `json:"id"`
	Type     uint8  `json:"type"`
	Title    string `json:"title"`
	Wechat   string `json:"wechat"`
	Address  string `json:"address"`
	Category uint8  `json:"category"`
	FromId   uint8  `json:"from_id"`
	CardId   uint32 `json:"card_id"`
	Amount   uint32 `json:"amount"`
	Status   int    `json:"status"`
}

func NewAirdropVO(info *model.AirdropReward) *AirdropVO {
	if info == nil {
		return nil
	}
	vo := &AirdropVO{
		info.Id,
		info.Type,
		info.Title,
		info.Wechat,
		info.Address,
		info.Category,
		info.FromId,
		info.CardId,
		info.Amount,
		info.Status,
	}

	return vo
}

func NewAirdropVOList(list []*model.AirdropReward) []*AirdropVO {
	l := len(list)
	if l == 0 {
		return nil
	}

	voList := make([]*AirdropVO, l)
	for index, v := range list {
		voList[index] = NewAirdropVO(v)
	}

	return voList
}
