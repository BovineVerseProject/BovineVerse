package debug

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	"brave-ox-web/utils"
	"brave-ox-web/web/response"
	"brave-ox-web/web/services/debug"
)

type debugController struct{ service debug.DebugService }

func DebugController() *debugController {
	return &debugController{
		service: debug.Service(),
	}
}

// ResetProfileFlag 重置改名机会
func (this *debugController) ResetProfileFlag(c *gin.Context) {
	addr := c.Param("address")
	if !common.IsHexAddress(addr) {
		response.RespondInvalidArgsErr(c)
		return
	}
	err := this.service.ResetProfileFreeFlag(utils.GetCheckSum(addr))
	if err != nil {
		response.RespondServerError(c)
		return
	}
	response.RespondOK(c)
}
