package protocol

import "brave-ox-web/web/vo"

type GetProfileResponse struct {
	vo.ProfileVO
	IconList []*vo.ProfileIconVO `json:"avatar"`
}

type ProfileName struct {
	Name string `json:"name" binding:"required,lte=42,gte=2"`
}

type UpdateProfileIconRequest struct {
	Avatar string `json:"avatar" binding:"required"`
}

type PhotoCode struct {
	Avatar []string `json:"avatar" binding:"required,lte=50,gte=1"`
}

type PhotoUri struct {
	Avatar string `json:"avatar"`
	Id     uint32 `json:"id"`
	Uri    string `json:"uri"`
}

type SearchAddrResponse struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type MultiAddress struct {
	Address []string `json:"address" binding:"required,gte=1,lte=50"`
}
