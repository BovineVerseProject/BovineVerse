package routing

import (
	"github.com/gin-gonic/gin"

	"brave-ox-web/config"
	"brave-ox-web/web/controller/debug"
)

func debugRouter(e *gin.Engine) {
	if !config.IsDebug() {
		return
	}
	h := debug.DebugController()
	g := newGroup(e, "/test")
	g.POST("/reset-free/:address", h.ResetProfileFlag)
}
