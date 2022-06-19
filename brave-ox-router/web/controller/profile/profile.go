package profile

import (
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	cusCommon "brave-ox-web/common"
	"brave-ox-web/log"
	"brave-ox-web/utils"
	"brave-ox-web/web/errorcode"
	"brave-ox-web/web/protocol"
	"brave-ox-web/web/response"
	"brave-ox-web/web/sensitive"
	"brave-ox-web/web/services/profile"
)

type profileController struct{ service profile.ProfileService }

func ProfileController() *profileController {
	return &profileController{
		service: profile.Service(),
	}
}

// ViewProfile 获取个人信息
func (this *profileController) ViewProfile(c *gin.Context) {
	addr := c.Param("address")
	if !common.IsHexAddress(addr) {
		response.RespondInvalidArgsErr(c)
		return
	}
	addr = utils.GetCheckSum(addr)
	data := this.service.GetProfile(addr)
	response.RespondWithData(c, data)
}

// ViewGameProfile 获取个人信息(游戏)
func (this *profileController) ViewGameProfile(c *gin.Context) {
	addr := c.Param("address")
	if !common.IsHexAddress(addr) {
		response.RespondInvalidArgsErr(c)
		return
	}
	addr = utils.GetCheckSum(addr)
	log.Info(addr)

	data := this.service.GetGameProfile(addr)
	response.RespondWithData(c, data)
}

// UpdateName 更新名称
func (this *profileController) UpdateName(c *gin.Context) {
	req := &protocol.ProfileName{}
	if err := c.ShouldBindJSON(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	req.Name = strings.Trim(req.Name, " ")
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
		m := utils.ChineseNameCheck(string(cr[i]))
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
	} else if character < 3 {
		response.RespondWithErrCode(c, errorcode.ErrCodeTextFormat, "length limit")
		return
	}

	claims, _ := c.Get(cusCommon.JWTKey)
	info := claims.(*utils.CustomClaims)
	addr := info.Address

	var newName = req.Name

	// 不得改成别人的地址，只能改自己的
	if n := utils.GetCheckSum(req.Name); n != cusCommon.ZeroAddress {
		if n != addr {
			response.RespondWithErrCode(c, errorcode.ErrCodeTextFormat, "only self address")
			return
		}
		newName = n
	}

	err := this.service.UpdateName(addr, newName)
	if err != nil {
		response.RespondWithErrCode2(c, err)
		return
	}

	response.RespondWithData(c, newName)
}

// UpdateProfile 更换头像
func (this *profileController) UpdateProfile(c *gin.Context) {
	req := &protocol.UpdateProfileIconRequest{}
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

	addr = utils.GetCheckSum(addr)
	err := this.service.UpdateProfilePhoto(addr, req.Avatar)
	if err != nil {
		response.RespondErr(c, err.Error())
		return
	}
	response.RespondWithData(c, req.Avatar)
}

func (this *profileController) PhotoImageUri(c *gin.Context) {
	req := &protocol.PhotoCode{}
	if err := c.ShouldBindJSON(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	response.RespondWithData(c, this.service.ParseProfilePhotoLink(req.Avatar))
}

func (this *profileController) SetGender(c *gin.Context) {
	g := c.Param("gender")
	gender, err := strconv.ParseUint(g, 0, 0)
	if err != nil || gender == 0 || (gender != 1 && gender != 2) {
		response.RespondInvalidArgsErr(c)
		return
	}

	claims, _ := c.Get(cusCommon.JWTKey)
	info := claims.(*utils.CustomClaims)

	err = this.service.SetGender(info.Address, uint8(gender))
	if err != nil {
		response.RespondWithErrCode2(c, err)
		return
	}

	response.RespondWithData(c, gender)
}

func (this *profileController) SearchAddr(c *gin.Context) {
	req := &protocol.ProfileName{}
	if err := c.ShouldBindJSON(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	if common.IsHexAddress(req.Name) {
		req.Name = utils.GetCheckSum(req.Name)
	}

	addr := this.service.SearchAddressByName(req.Name)

	response.RespondWithData(c, &protocol.SearchAddrResponse{
		Name:    req.Name,
		Address: addr,
	})

}

func (this *profileController) MultiProfile(c *gin.Context) {
	req := &protocol.MultiAddress{}
	if err := c.ShouldBindJSON(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	for idx, v := range req.Address {
		if !common.IsHexAddress(v) {
			response.RespondInvalidArgsErr(c)
		}
		req.Address[idx] = utils.GetCheckSum(v)
	}

	data, err := this.service.GetMultiProfile(req.Address)
	if err != nil {
		response.RespondWithErrCode2(c, err)
		return
	}

	response.RespondWithData(c, data)
}
