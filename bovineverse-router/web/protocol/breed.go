package protocol

import "brave-ox-web/web/vo"

type BreedCenterRequest struct {
	Order  uint8  `json:"order" binding:"oneof=0 1 2"`
	Gender uint8  `json:"gender" binding:"oneof=0 1 2"`
	Page   uint32 `json:"page"`
}

type BreedCenterResponse struct {
	Order  uint8 `json:"order,omitempty"`
	Gender uint8 `json:"gender,omitempty"`
	PageResponse

	List []*vo.BreedCattleVO `json:"list"`
}

type BreedCenterSearchRequest struct {
	Order  uint8  `json:"order" binding:"omitempty,oneof=0 1 2"`
	Gender uint8  `json:"gender" binding:"omitempty,oneof=0 1 2"`
	Page   uint32 `json:"page,omitempty"`
	Match  string `json:"match" binding:"required,lte=42"`
}
