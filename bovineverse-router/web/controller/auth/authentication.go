package auth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	cus "brave-ox-web/common"
	"brave-ox-web/utils"
	"brave-ox-web/web/errorcode"
	"brave-ox-web/web/protocol"
	"brave-ox-web/web/response"
)

type authenticationController struct{}

func AuthenticationController() *authenticationController {
	return &authenticationController{}
}

const identifyMsgTemplate = "Confirm to connect your wallet to Bovine-Verse successfully"

func (*authenticationController) Verify(c *gin.Context) {
	req := &protocol.AuthRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	if !common.IsHexAddress(req.Address) {
		response.RespondInvalidArgsErr(c)
		return
	}

	adr := utils.GetCheckSum(req.Address)
	err := utils.VerifySign(identifyMsgTemplate, req.Sig, adr)
	if err != nil {
		response.RespondWithErrCode(c, errorcode.ErrCodeWrongSig, req.Sig)
		return
	}

	t, err := utils.GenerateToken(0, adr)
	if err != nil {
		response.RespondWithErrCode(c, errorcode.ErrCodeCommon, err.Error())
		return
	}

	resp := &protocol.AuthResponse{
		Address: adr,
		Token:   t,
	}

	response.RespondWithData(c, resp)

}

func (*authenticationController) RefreshToken(c *gin.Context) {
	claims, _ := c.Get(cus.JWTKey)
	info := claims.(*utils.CustomClaims)

	t, err := utils.GenerateToken(0, info.Address)
	if err != nil {
		response.RespondWithErrCode(c, errorcode.ErrCodeCommon, err.Error())
		return
	}

	resp := &protocol.AuthResponse{
		Address: info.Address,
		Token:   t,
	}

	response.RespondWithData(c, resp)

}
