package cattle

import (
	"fmt"
	"strconv"

	"brave-ox-web/conf"
	"brave-ox-web/web/services/cattle"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"brave-ox-web/web/cache"
	"brave-ox-web/web/protocol"
	"brave-ox-web/web/response"
	"brave-ox-web/web/vo"
)

type cattleController struct{ service cattle.CattleService }

func CattleController() *cattleController {
	return &cattleController{
		service: cattle.Service(),
	}
}

func (this *cattleController) GetCattle(c *gin.Context) {
	id := c.Param("id")
	nftID, err := strconv.Atoi(id)
	if nftID == 0 && err != nil {
		response.RespondInvalidArgsErr(c, "invalid cattle id")
		return
	}
	info, err := this.service.GetCattleById(uint32(nftID))
	if err != nil {
		response.RespondServerError(c)
		return
	}

	var cattleVO *vo.CattleVO
	if info != nil {
		cattleVO = vo.NewCattleVO(info)
		cattleVO.Name = cache.CattleName(info.Class, info.Gender, info.GenderSeq, info.IsAdult)
		if info.Class == conf.TypeBovineHero {
			cattleVO.Name = fmt.Sprintf("Bovine Hero #%d", info.Id)
		}
		cattleVO.Image = cache.GetCattleImage(info.Class, info.Gender, info.GenderSeq, info.IsAdult)
	}
	response.RespondWithData(c, cattleVO)
}

func (this *cattleController) GetCattleList(c *gin.Context) {
	req := &protocol.CattleListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	repeatedCheck := make(map[uint32]bool)
	for _, v := range req.Ids {
		if _, ok := repeatedCheck[v]; ok {
			response.RespondInvalidArgsErr(c, fmt.Sprintf("repeated id:%v", req.Ids))
			return
		}
		repeatedCheck[v] = true
	}

	data, err := this.service.GetCattleList(req.Ids)
	if err != nil {
		response.RespondServerError(c)
		return
	}

	var cattleVOList []*vo.CattleVO
	if len(data) > 0 {
		cattleVOList = make([]*vo.CattleVO, len(data))
		for idx, v := range data {
			cattleVO := vo.NewCattleVO(v)
			cattleVO.Name = cache.CattleName(v.Class, v.Gender, v.GenderSeq, v.IsAdult)
			cattleVO.Image = cache.GetCattleImage(v.Class, v.Gender, v.GenderSeq, v.IsAdult)
			cattleVOList[idx] = cattleVO
		}
	}

	response.RespondWithData(c, cattleVOList)
}
