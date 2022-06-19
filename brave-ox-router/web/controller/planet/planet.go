package planet

import (
	"io"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	cusCommon "brave-ox-web/common"
	"brave-ox-web/utils"
	"brave-ox-web/web/errorcode"
	"brave-ox-web/web/protocol"
	"brave-ox-web/web/response"
	"brave-ox-web/web/sensitive"
	"brave-ox-web/web/services/planet"
	"brave-ox-web/web/vo"
)

type planetController struct{ service planet.PlanetService }

func PlanetController() *planetController {
	return &planetController{
		service: planet.Service(),
	}
}

// QueryPlanets 查询某类型下的所有星球
func (this *planetController) QueryPlanets(c *gin.Context) {
	types := c.Param("types")
	t, err := strconv.ParseUint(types, 0, 0)

	// 0全部、1家园星球、2联邦星球、3拓荒星球
	if err != nil || t != 0 && t != 1 && t != 2 && t != 3 {
		response.RespondInvalidArgsErr(c)
		return
	}

	req := &protocol.PlanetTypesRequest{}
	if err = c.ShouldBindJSON(req); err != nil && err != io.EOF {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	if len(req.PrefixName) > 0 {
		m := utils.SpecialRegexCheck(req.PrefixName)
		if m {
			response.RespondWithErrCode(c, errorcode.ErrCodeTextFormat, "invalid name format")
			return
		}

		ok, dirty := sensitive.Validate(req.PrefixName)
		if !ok {
			response.RespondWithErrCode(c, errorcode.ErrCodeTextFormat, "sensitive character:"+dirty)
			return
		}
	}

	data, total, page := this.service.ListPlanetByTypes(uint32(t), req.Order, req.Page, req.PrefixName)
	resp := &protocol.PlanetTypesResponse{}
	resp.Total = total
	resp.Page = page
	resp.List = vo.NewPlanetVOList(data)

	response.RespondWithData(c, resp)
}

// QueryPlanet 获取星球基本信息
func (this *planetController) QueryPlanet(c *gin.Context) {
	types := c.Param("id")
	t, err := strconv.ParseUint(types, 0, 0)
	if err != nil || t == 0 {
		response.RespondInvalidArgsErr(c)
		return
	}

	data := this.service.QueryPlanetById(uint32(t))
	response.RespondWithData(c, vo.NewBasicPlanetVO(data))
}

// UpdateName 修改星球名称
func (this *planetController) UpdateName(c *gin.Context) {
	id := c.Param("id")
	tid, err := strconv.ParseUint(id, 0, 0)
	if err != nil || tid == 0 {
		response.RespondInvalidArgsErr(c)
		return
	}

	req := &protocol.ChangePlanetNameRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	m := utils.SpecialRegexCheck(req.Name)
	if m {
		response.RespondWithErrCode(c, errorcode.ErrCodeTextFormat, "invalid name format")
		return
	}

	ok, dirty := sensitive.Validate(req.Name)
	if !ok {
		response.RespondWithErrCode(c, errorcode.ErrCodeTextFormat, "sensitive character:"+dirty)
		return
	}

	var chineseCount uint8
	var character uint8

	cr := []rune(req.Name)
	for i := 0; i < len(cr); i++ {
		m = utils.ChineseNameCheck(string(cr[i]))
		if m {
			chineseCount++
		} else {
			character++
		}
	}

	if chineseCount > 0 {
		if chineseCount < 2 || chineseCount > 6 {
			response.RespondWithErrCode(c, errorcode.ErrCodeTextFormat, "length limit")
			return
		}
	} else if character < 3 { // 非中文必须3个字符
		response.RespondWithErrCode(c, errorcode.ErrCodeTextFormat, "length limit")
		return
	}

	claims, _ := c.Get(cusCommon.JWTKey)
	info := claims.(*utils.CustomClaims)
	addr := info.Address

	err = this.service.ChangeName(uint32(tid), req.Name, addr)
	if err != nil {
		response.RespondWithErrCode2(c, err)
		return
	}
	response.RespondWithData(c, req.Name)

}

// UpdateNotice 修改公告
func (this *planetController) UpdateNotice(c *gin.Context) {
	id := c.Param("id")
	tid, err := strconv.ParseUint(id, 0, 0)
	if err != nil || tid == 0 {
		response.RespondInvalidArgsErr(c)
		return
	}

	req := &protocol.ChangePlanetNoticeRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	ok, dirty := sensitive.Validate(req.Notice)
	if !ok {
		response.RespondWithErrCode(c, errorcode.ErrCodeTextFormat, "sensitive character:"+dirty)
		return
	}

	claims, _ := c.Get(cusCommon.JWTKey)
	info := claims.(*utils.CustomClaims)
	addr := info.Address

	err = this.service.ChangeNotice(uint32(tid), req.Notice, addr)
	if err != nil {
		response.RespondWithErrCode2(c, err)
		return
	}
	response.RespondWithData(c, req)

}

func (this *planetController) Members(c *gin.Context) {
	id := c.Param("id")
	tid, err := strconv.ParseUint(id, 0, 0)
	if err != nil || tid == 0 {
		response.RespondInvalidArgsErr(c)
		return
	}

	req := &protocol.MembersRequest{}
	if err = c.ShouldBindJSON(req); err != nil && err != io.EOF {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	member, total, page, err := this.service.QueryMembers(uint32(tid), req.Order, req.Page)
	if err != nil {
		response.RespondWithErrCode2(c, err)
		return
	}

	resp := &protocol.MembersResponse{}
	resp.Total = total
	resp.Page = page
	resp.Members = vo.NewPlanetMemberVOList(member)

	response.RespondWithData(c, resp)
}

func (this *planetController) ChangeMemberPosition(c *gin.Context) {
	req := &protocol.ChangeMemberPosition{}
	if err := c.ShouldBindJSON(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	claims, _ := c.Get(cusCommon.JWTKey)
	info := claims.(*utils.CustomClaims)
	addr := info.Address
	err := this.service.ChangeMemberPosition(addr, req.Member, req.Position)
	if err != nil {
		response.RespondWithErrCode2(c, err)
		return
	}
	response.RespondWithData(c, req)
}

func (this *planetController) ChangeImage(c *gin.Context) {
	id := c.Param("imageId")
	tid, err := strconv.ParseUint(id, 0, 0)
	if err != nil || tid > 26 {
		response.RespondInvalidArgsErr(c)
		return
	}

	claims, _ := c.Get(cusCommon.JWTKey)
	info := claims.(*utils.CustomClaims)
	addr := info.Address
	err = this.service.ChangeImage(addr, uint8(tid))
	if err != nil {
		response.RespondWithErrCode2(c, err)
		return
	}
	response.RespondWithData(c, tid)
}

func (this *planetController) GetSubPlanets(c *gin.Context) {
	req := &protocol.SubPlanetRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}

	data, total, page := this.service.QuerySubPlanets(req.Id, req.Page)
	resp := &protocol.SubPlanetsResponse{}
	resp.Total = total
	resp.Page = page
	resp.List = vo.NewSubPlanetVOList(data)
	response.RespondWithData(c, resp)
}

func (this *planetController) MembersPosition(c *gin.Context) {
	addr := c.Param("address")
	if !common.IsHexAddress(addr) {
		response.RespondInvalidArgsErr(c)
		return
	}
	checkSum := utils.GetCheckSum(addr)

	members, err := this.service.GetMember(checkSum)
	if err != nil {
		response.RespondErr(c, err.Error())
		return
	}

	response.RespondWithData(c, vo.NewMemberGroupBasicVO(members))
}

func (this *planetController) QueryMemberGroup(c *gin.Context) {
	id := c.Param("id")
	tid, err := strconv.ParseUint(id, 0, 0)
	if err != nil || tid == 0 {
		response.RespondInvalidArgsErr(c)
		return
	}

	d, err := this.service.MemberGroup(uint32(tid))
	if err != nil {
		response.RespondErr(c, err.Error())
		return
	}

	response.RespondWithData(c, vo.NewPlanetMemberGroupVO(d))
}
