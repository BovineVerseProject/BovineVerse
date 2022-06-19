package routing

import (
	"github.com/gin-gonic/gin"

	"brave-ox-web/web/controller"
	"brave-ox-web/web/controller/airdrop"
	"brave-ox-web/web/controller/auth"
	"brave-ox-web/web/controller/breed"
	"brave-ox-web/web/controller/cattle"
	"brave-ox-web/web/controller/currency"
	"brave-ox-web/web/controller/marketplace"
	"brave-ox-web/web/controller/metadata"
	"brave-ox-web/web/controller/planet"
	"brave-ox-web/web/controller/profile"
)

func init() {
	addRouterFunc(cattleRouter)
	addRouterFunc(planetRouter)
	addRouterFunc(profileRouter)
	addRouterFunc(nftRouter)
	addRouterFunc(marketRouter)
	addRouterFunc(breedRouter)
	addRouterFunc(airdropRouter)
	addRouterFunc(userAuthRouter)
	addRouterFunc(usdtRateRouter)
	addRouterFunc(debugRouter)
	addRouterFunc(adminRouter)
}

func marketRouter(e *gin.Engine) {
	h := marketplace.MarketplaceController()
	g := newGroup(e, "/marketplace")
	g.POST("/cattles", h.CattleList)
	g.POST("/mysteryboxes", h.Mysteryboxes)
	g.POST("/planet", h.Planet)
	g.POST("/history/:address", h.History)
}

func planetRouter(e *gin.Engine) {
	h := planet.PlanetController()
	g := newGroup(e, "/planet")
	g.POST("/list/:types", h.QueryPlanets)
	g.GET("/:id", h.QueryPlanet)
	g.POST("/member/:id", h.Members)
	g.GET("/member/:address", h.MembersPosition)
	g.GET("/member/group/:id", h.QueryMemberGroup)
	g.POST("/sub-planets", h.GetSubPlanets)
	g.Use(jwtAuth())
	g.PUT("/:id", h.UpdateName)
	g.PUT("/member", h.ChangeMemberPosition)
	g.POST("/image/:imageId", h.ChangeImage)
	g.POST("/:id", h.UpdateNotice)
}

func cattleRouter(e *gin.Engine) {
	h := cattle.CattleController()
	g := newGroup(e, "/cattle")
	g.GET("/:id", h.GetCattle)
	g.POST("/list", h.GetCattleList)
}

func nftRouter(e *gin.Engine) {
	h := metadata.MetadataController()
	g := newGroup(e, "/nft")
	g.GET("/bovine-hero/:id", h.GetCattleNFT)
	g.GET("/box/:address/:types", h.Box)
	g.GET("/skin/:id", h.CattleSkin)
	g.GET("/planet/:id", h.Planet)
}

func profileRouter(e *gin.Engine) {
	h := profile.ProfileController()
	g := newGroup(e, "/profile")
	g.GET("/:address", h.ViewProfile)
	g.GET("/game/:address", h.ViewGameProfile)
	g.POST("/photo-link", h.PhotoImageUri)
	g.Use(jwtAuth())
	g.POST("/game/multi", h.MultiProfile)
	g.POST("/search", h.SearchAddr)
	g.PUT("/name", h.UpdateName)
	g.PUT("/photo", h.UpdateProfile)
	g.POST("/:gender", h.SetGender)
}

func breedRouter(e *gin.Engine) {
	h := breed.BreedController()
	g := newGroup(e, "/breed")
	g.POST("/list", h.BreedList)
	g.POST("/search", h.BreedSearch)
}

func airdropRouter(e *gin.Engine) {
	h := airdrop.AirdropController()
	g := newGroup(e, "/airdrop")
	g.GET("/list", h.QueryList)
	g.GET("/sign/:id", h.QuerySign)
	g.GET("/claim/:id", h.Claim)
	g.POST("", h.AddRecord)
}

func userAuthRouter(e *gin.Engine) {
	h := auth.AuthenticationController()
	g := newGroup(e, "/auth")
	g.POST("/verify", h.Verify)
	g.Use(jwtAuth())
	g.GET("/refresh-token", h.RefreshToken)
}

func usdtRateRouter(e *gin.Engine) {
	h := currency.CurrencyController()
	g := newGroup(e, "/currency")
	g.GET("/rate", h.USDTRate)
}

func adminRouter(e *gin.Engine) {
	l := controller.LoginController()
	e.GET("/login", l.OtaLogin)
	e.Use(onlyAdmin())

	a := airdrop.AirdropController()
	e.POST("/airdrop", a.AddRecord)
}
