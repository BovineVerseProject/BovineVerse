package vo

import "brave-ox-web/model"

type SubPlanetVO struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Score string `json:"score"`
	Image uint16 `json:"image"`
}

type PlanetVO struct {
	SubPlanetVO
	Address string `json:"address"`
	Current uint16 `json:"current"`
	Fee     string `json:"fee"`
	Owner   string `json:"owner"`
	Notice  string `json:"notice"`
}

func NewSubPlanetVOList(list []*model.PlanetInfo) []*SubPlanetVO {
	l := len(list)
	if l <= 0 {
		return nil
	}

	vo := make([]*SubPlanetVO, l)
	for idx, v := range list {
		vo[idx] = &SubPlanetVO{
			Id:    v.Id,
			Name:  v.Name,
			Score: v.TotalScore,
			Image: v.Image,
		}
	}

	return vo
}

func NewPlanetVO(info *model.PlanetInfo) *PlanetVO {
	if info == nil {
		return nil
	}
	obj := &PlanetVO{}
	obj.Id = info.Id
	obj.Name = info.Name
	obj.Score = info.TotalScore
	obj.Image = info.Image
	obj.Address = info.Owner
	obj.Current = info.Population
	obj.Fee = info.Fee
	obj.Owner = info.OwnerName
	obj.Notice = info.Notice
	return obj
}

func NewPlanetVOList(list []*model.PlanetInfo) []*PlanetVO {
	l := len(list)
	if l == 0 {
		return nil
	}

	vo := make([]*PlanetVO, l)
	for index, v := range list {
		vo[index] = NewPlanetVO(v)
	}
	return vo
}

type BasicPlanetVO struct {
	PlanetVO
	// 这些字段特定时候才返回
	UpdateNameAt   int64 `json:"updateNameAt"`
	UpdateNoticeAt int64 `json:"updateNoticeAt"`
}

func NewBasicPlanetVO(info *model.PlanetInfo) *BasicPlanetVO {
	if info == nil {
		return nil
	}
	vo := &BasicPlanetVO{}
	vo.PlanetVO = *NewPlanetVO(info)
	vo.UpdateNameAt = info.UpdateNameAt
	vo.UpdateNoticeAt = info.UpdateNoticeAt
	return vo
}

type PlanetMemberVO struct {
	Id       uint32 `json:"id"`
	Player   string `json:"player"`
	Position uint8  `json:"position"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Count    uint32 `json:"count"`
	Score    string `json:"score"`
}

func NewPlanetMemberVo(info *model.PlanetMember) *PlanetMemberVO {
	if info == nil {
		return nil
	}

	return &PlanetMemberVO{
		Id:       info.PlanetId,
		Player:   info.Player,
		Position: info.Position,
		Name:     info.Name,
		Count:    info.Count,
		Score:    info.Score,
		Avatar:   info.Avatar,
	}
}

func NewPlanetMemberVOList(infos []*model.PlanetMember) []*PlanetMemberVO {
	l := len(infos)
	if l == 0 {
		return nil
	}

	voList := make([]*PlanetMemberVO, l)
	for idx, v := range infos {
		voList[idx] = NewPlanetMemberVo(v)
	}
	return voList
}

type MemberBasicVO struct {
	Address   string `json:"address"`
	Position  uint8  `json:"position"`
	JointTime int64  `json:"joinTime"`
}

type MemberGroupBasicVO struct {
	PlanetId  uint32 `json:"planetId"`
	Position  uint8  `json:"position"`
	JointTime int64  `json:"joinTime"`
}

func NewMemberGroupBasicVO(members []model.PlanetMember) []MemberGroupBasicVO {
	var vo []MemberGroupBasicVO
	for _, v := range members {
		vo = append(vo, MemberGroupBasicVO{
			PlanetId:  v.PlanetId,
			Position:  v.Position,
			JointTime: v.CreateAt,
		})
	}
	return vo
}

type MemberGroupVO struct {
	Position uint8  `json:"position"`
	Count    uint32 `json:"count"`
}

func NewPlanetMemberGroupVO(list []*model.MemberGroup) []*MemberGroupVO {
	var vo []*MemberGroupVO
	for _, v := range list {
		vo = append(vo, &MemberGroupVO{
			Position: v.Position,
			Count:    v.Count,
		})
	}

	return vo
}
