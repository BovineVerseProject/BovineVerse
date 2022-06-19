package currency

import (
	"github.com/gin-gonic/gin"

	"brave-ox-web/web/response"
	"brave-ox-web/web/services/currency"
)

type currencyController struct{ service currency.CurrencyService }

func CurrencyController() *currencyController {
	return &currencyController{
		service: currency.Service(),
	}
}

func (this *currencyController) USDTRate(c *gin.Context) {
	response.RespondWithData(c, this.service.GetAllTokenUSDTRate())
}
