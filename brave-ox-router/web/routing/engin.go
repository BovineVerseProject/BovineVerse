package routing

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"brave-ox-web/common"
	"brave-ox-web/conf"
	"brave-ox-web/config"
	"brave-ox-web/log"
	"brave-ox-web/utils"
	"brave-ox-web/web/errorcode"
	"brave-ox-web/web/response"
)

var (
	_router []func(e *gin.Engine)
)

func InitEngine() *gin.Engine {
	if !config.IsDebug() {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()
	e.Use(gin.RecoveryWithWriter(log.LoggerWriter(), handleRecovery), gin.LoggerWithWriter(log.LoggerWriter()))
	if config.IsUseCors() {
		e.Use(cors())
	}

	for _, v := range _router {
		v(e)
	}
	return e
}

func newGroup(engine *gin.Engine, relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return engine.Group(relativePath, handlers...)
}

func addRouterFunc(f func(e *gin.Engine)) {
	_router = append(_router, f)
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Origin") != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization,Token")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
			}
		}

		c.Next()
	}
}

func jwtAuth() gin.HandlerFunc {
	return parseToken(false)
}

func onlyAdmin() gin.HandlerFunc {
	return parseToken(true)
}

func parseToken(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Token")
		if token == "" {
			response.RespondWithErrCode(c, errorcode.ErrCodeNotAuthenticated, "request without token")
			c.Abort()
			return
		}

		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.RespondWithErrCode(c, errorcode.ErrCodeAuthenticationExpired, "authentication expired")
				c.Abort()
				return
			}
			response.RespondErr(c, err.Error())
			c.Abort()
			return
		}

		if isAdmin {
			if claims.Address != "admin" {
				c.Abort()
				return
			}
		}
		c.Set(common.JWTKey, claims)
	}
}

func accessCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !config.IsDebug() {
			ip := c.ClientIP()
			if !allowed(ip) {
				response.RespondWithErrCode(c, errorcode.ErrCodeNotAuthenticated, "access denied")
				c.Abort()
			}
		}
	}
}

func handleRecovery(c *gin.Context, err interface{}) {
	response.RespondWithErrCode(c, errorcode.ErrCodeServerBusy, "internal server error")
}

func newLimiter(r rate.Limit, b int, t time.Duration) gin.HandlerFunc {
	limiters := &sync.Map{}

	return func(c *gin.Context) {
		key := c.ClientIP()
		if !allowed(key) {
			l, _ := limiters.LoadOrStore(key, rate.NewLimiter(r, b))
			ctx, cancel := context.WithTimeout(c, t)
			defer cancel()

			if err := l.(*rate.Limiter).Wait(ctx); err != nil {
				response.RespondWithErrCode(c, errorcode.ErrCodeRequestBusy, err.Error())
				c.Abort()
			}
		}
	}
}

func allowed(ip string) bool {
	for _, v := range conf.IPWhiteList() {
		if ip == v {
			return true
		}
	}
	return false
}
