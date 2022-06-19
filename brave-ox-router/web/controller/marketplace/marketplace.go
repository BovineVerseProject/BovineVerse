package marketplace

import (
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"brave-ox-web/utils"
	"brave-ox-web/web/protocol"
	"brave-ox-web/web/response"
	"brave-ox-web/web/services/market"
)

type marketplaceController struct{ service market.MarketService }

func MarketplaceController() *marketplaceController {
	return &marketplaceController{
		service: market.Service(),
	}
}

func (this *marketplaceController) CattleList(c *gin.Context) {
	req := &protocol.MarketCattleRequest{}
	if err := c.ShouldBindJSON(req); err != nil && err != io.EOF {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	//  注释原因：只有普通成年牛才有星级，但是普通成年牛不能交易
	//var l int
	//if l = len(req.Star); l > 0 {
	//	cmap := make(map[uint32]uint32)
	//	for _, v := range req.Star {
	//		if _, ok := cmap[v]; ok || v > 3 {
	//			respondErr(c, fmt.Sprintf("Invalid star: %d", v))
	//			return
	//		}
	//		cmap[v] = 0
	//	}
	//
	//	// 如果发全星级过来直接优化掉
	//	if l == 4 {
	//		req.Star = nil
	//	}
	//}

	if !this.arrayCheck(&req.Life) {
		response.RespondInvalidArgsErr(c, "Invalid life args")
		return
	}

	if !this.arrayCheck(&req.Growth) {
		response.RespondInvalidArgsErr(c, "Invalid growth args")
		return
	}

	if !this.arrayCheck(&req.Energy) {
		response.RespondInvalidArgsErr(c, "Invalid energy args")
		return
	}

	// 没选性别的情况下把所有特殊(公、母)属性都清空
	if req.Gender == 0 {
		req.Attack = nil
		req.Defence = nil
		req.Stamina = nil
		req.Milk = nil
		req.MilkRate = nil
	} else {
		// 没有选择公牛查询却选了公牛的属性
		if !this.checkOxAttr(req) {
			response.RespondInvalidArgsErr(c, fmt.Sprintf("Query condition err:ox"))
			return
		}

		// 检查公牛数值
		if req.Gender == 1 && (!this.arrayCheck(&req.Attack) || !this.arrayCheck(&req.Defence) || !this.arrayCheck(&req.Stamina)) {
			response.RespondInvalidArgsErr(c, fmt.Sprintf("Query condition err:ox"))
			return
		}

		// 没有选择母牛查询却选了母牛的属性
		if !this.checkCowAttr(req) {
			response.RespondInvalidArgsErr(c, fmt.Sprintf("Query condition err:cow"))
			return
		}

		// 检查母牛数值
		if req.Gender == 2 && !this.arrayCheck(&req.Milk) || !this.arrayCheck(&req.MilkRate) {
			response.RespondInvalidArgsErr(c, fmt.Sprintf("Query condition err:cow"))
			return
		}

	}

	if req.Page == 0 {
		req.Page = 1 // 默认第一页
	}

	result, count, err := this.service.SearchSellingCattle(req)
	if err != nil {
		response.RespondServerError(c)
		return
	}

	resp := &protocol.MarketCattleResponse{
		GoodsType: 1,
		List:      result,
	}
	resp.Total = count
	resp.Page = req.Page

	response.RespondWithData(c, resp)
}

// 检查公牛属性
func (this *marketplaceController) checkOxAttr(req *protocol.MarketCattleRequest) bool {
	if req.Gender == 1 && (len(req.Milk) > 0 || len(req.MilkRate) > 0) {
		return false
	}
	return true
}

// 检查母牛属性
func (this *marketplaceController) checkCowAttr(req *protocol.MarketCattleRequest) bool {
	if req.Gender == 2 {
		if len(req.Attack) > 0 || len(req.Defence) > 0 || len(req.Stamina) > 0 {
			return false
		}
	}
	return true
}

func (this *marketplaceController) arrayCheck(attr *[]uint32) bool {
	array := *attr
	if l := len(array); l == 0 {
		return true
	} else if l != 2 {
		return false
	}

	if array[0] < 5000 {
		array[0] = 5000
	}

	if array[1] > 15000 {
		array[1] = 15000
	}

	if array[0] > array[1] {
		array[0] = array[1]
	}

	// 如果都是默认区间直接优化掉
	if array[0] == 5000 && array[1] == 15000 {
		*attr = nil
	}
	return true
}

func (this *marketplaceController) Mysteryboxes(c *gin.Context) {
	req := &protocol.MarketRequest{}
	if err := c.ShouldBindJSON(req); err != nil && err != io.EOF {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	result, count, err := this.service.SearchSellMysteryBox(req)
	if err != nil {
		response.RespondServerError(c)
		return
	}

	resp := &protocol.MarketMysteryBoxResponse{
		GoodsType: 2,
		List:      result,
	}
	resp.Total = count
	resp.Page = req.Page

	response.RespondWithData(c, resp)
}

func (this *marketplaceController) Planet(c *gin.Context) {
	req := &protocol.MarketRequest{}
	if err := c.ShouldBindJSON(req); err != nil && err != io.EOF {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	result, total, page, err := this.service.SearchSellPlanet(req)
	if err != nil {
		response.RespondServerError(c)
		return
	}
	resp := &protocol.MarketPlanetResponse{
		GoodsType: 3,
		List:      result,
	}
	resp.Total = total
	resp.Page = page

	response.RespondWithData(c, resp)
}

func (this *marketplaceController) History(c *gin.Context) {
	req := &protocol.MarketHistoryRequest{}
	if err := c.ShouldBindJSON(req); err != nil && err != io.EOF {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.RespondInvalidArgsErr(c, response.RemoveTopStruct(errs.Translate(response.Trans())))
			return
		}
		response.RespondInvalidArgsErr(c)
		return
	}

	addr := c.Param("address")
	if !common.IsHexAddress(addr) {
		response.RespondInvalidArgsErr(c)
		return
	}
	addr = utils.GetCheckSum(addr)
	result, total, page, err := this.service.History(addr, req.Page)
	if err != nil {
		response.RespondServerError(c)
		return
	}

	resp := &protocol.MarketHistoryResponse{
		History: result,
	}
	resp.Total = total
	resp.Page = page

	response.RespondWithData(c, resp)
}
