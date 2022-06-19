package breed

import (
	"fmt"
	"io"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"brave-ox-web/utils"
	"brave-ox-web/web/protocol"
	"brave-ox-web/web/response"
	"brave-ox-web/web/services/breed"
)

type breedController struct{ service breed.BreedService }

func BreedController() *breedController {
	return &breedController{
		service: breed.Service(),
	}
}

func (this *breedController) BreedList(c *gin.Context) {
	req := &protocol.BreedCenterRequest{}
	if err := c.ShouldBindJSON(req); err != nil && err != io.EOF {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	data, total, page := this.service.BreedList(req.Order, req.Gender, req.Page)
	resp := &protocol.BreedCenterResponse{
		Order:  req.Order,
		Gender: req.Gender,
		List:   data,
	}
	resp.Page = page
	resp.Total = total

	response.RespondWithData(c, resp)
}

func (this *breedController) BreedSearch(c *gin.Context) {
	req := &protocol.BreedCenterSearchRequest{}
	if err := c.ShouldBindJSON(req); err != nil && err != io.EOF {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	var addr string
	tid, err := strconv.ParseUint(req.Match, 0, 0)
	if err != nil || tid == 0 {
		if !common.IsHexAddress(req.Match) {
			response.RespondInvalidArgsErr(c, fmt.Sprintf("tid:%d", tid))
			return
		}
		addr = utils.GetCheckSum(req.Match)
	}

	data, total, page := this.service.BreedCenterSearch(addr, uint32(tid), req.Order, req.Gender, req.Page)
	resp := &protocol.BreedCenterResponse{
		Order:  req.Order,
		Gender: req.Gender,
		List:   data,
	}
	resp.Page = page
	resp.Total = total

	response.RespondWithData(c, resp)
}
