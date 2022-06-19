package controller

import (
	"github.com/gin-gonic/gin"

	"brave-ox-web/conf"
	"brave-ox-web/utils"
	"brave-ox-web/web/protocol"
	"brave-ox-web/web/response"
)

type loginController struct{}

func LoginController() *loginController {
	return &loginController{}
}

func (*loginController) OtaLogin(c *gin.Context) {
	var reqData struct {
		Code string `json:"code"`
	}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		response.RespondInvalidArgsErr(c)
		return
	}
	auth := utils.GoogleAuth{}
	ok, _ := auth.VerifyCode(conf.OTASecret(), reqData.Code)
	if !ok {
		response.RespondInvalidArgsErr(c, "invalid code")
		return
	}

	user := "admin"
	token, err := utils.GenerateToken(0, user)
	if err != nil {
		response.RespondErr(c, "invalid code")
		return
	}

	response.RespondWithData(c, protocol.LoginResult{
		User:  user,
		Token: token,
	})
}
