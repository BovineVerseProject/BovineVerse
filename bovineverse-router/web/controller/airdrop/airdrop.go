package airdrop

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"brave-ox-web/web/protocol"
	"brave-ox-web/web/response"
	"brave-ox-web/web/services/airdrop"
	"brave-ox-web/web/vo"
)

type airdropController struct {
	service airdrop.AirDropService
}

func AirdropController() *airdropController {
	return &airdropController{
		service: airdrop.Service(),
	}
}

func (a *airdropController) QueryList(c *gin.Context) {
	req := &protocol.AirdropRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}
	if len(req.Address) > 0 {
		if !common.IsHexAddress(req.Address) {
			response.RespondInvalidArgsErr(c, "wrong address")
			return
		}
		req.Address = common.HexToAddress(req.Address).Hex()
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	list, total, err := a.service.QueryRewards(req)
	if err != nil {
		response.RespondErr(c, err.Error())
		return
	}

	v := vo.NewAirdropVOList(list)
	response.RespondWithData(c, gin.H{
		"list":  v,
		"page":  req.Page,
		"limit": uint32(req.Limit),
		"total": uint32(total),
	})
}

func (a *airdropController) QuerySign(c *gin.Context) {
	recordId, err := strconv.ParseInt(c.Param("id"), 0, 0)
	if err != nil {
		response.RespondInvalidArgsErr(c)
		return
	}
	info, err := a.service.QueryRewardSign(uint32(recordId))
	if err != nil {
		response.RespondErr(c, err.Error())
		return
	}
	response.RespondWithData(c, info)
}

func (a *airdropController) Claim(c *gin.Context) {
	recordId, err := strconv.ParseInt(c.Param("id"), 0, 0)
	if err != nil {
		response.RespondInvalidArgsErr(c)
		return
	}
	if err := a.service.ClaimReward(uint32(recordId)); err != nil {
		response.RespondErr(c, err.Error())
		return
	}

	response.RespondOK(c)
}

func (a *airdropController) AddRecord(c *gin.Context) {
	files, err := c.FormFile("file")
	if err != nil {
		response.RespondInvalidArgsErr(c)
		return
	}
	if err := a.service.LoadExcel(files); err != nil {
		response.RespondErr(c, err.Error())
		return
	}
	response.RespondOK(c)
}
