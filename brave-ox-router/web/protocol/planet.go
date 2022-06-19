package protocol

import "brave-ox-web/web/vo"

type ChangePlanetNameRequest struct {
	Name string `json:"name" binding:"required,lte=12,gte=2"`
}

type ChangePlanetNoticeRequest struct {
	Notice string `json:"notice" binding:"max=50"`
}

type MembersRequest struct {
	Order uint8  `json:"order,omitempty" binding:"max=1"`
	Page  uint32 `json:"page"`
}

type MembersResponse struct {
	PageResponse
	Members []*vo.PlanetMemberVO `json:"members"`
}

type ChangeMemberPosition struct {
	Member   string `json:"member" binding:"required,len=42"`
	Position uint8  `json:"position" binding:"oneof=0 1 2 3"`
}

type ChangePlanetImage struct {
	Image uint8 `json:"image" biding:"required,max=26"`
}

type PlanetTypesRequest struct {
	Order      uint8  `json:"order,omitempty" binding:"max=2"`
	Page       uint32 `json:"page"`
	PrefixName string `json:"name" binding:"lte=12"`
}

type PlanetTypesResponse struct {
	PageResponse
	List []*vo.PlanetVO `json:"list"`
}

type SubPlanetRequest struct {
	Id   uint32 `json:"id" biding:"required"`
	Page uint32 `json:"page"`
}

type SubPlanetsResponse struct {
	PageResponse
	List []*vo.SubPlanetVO `json:"list"`
}
