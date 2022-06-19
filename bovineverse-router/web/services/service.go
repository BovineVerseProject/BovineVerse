package services

import (
	"brave-ox-web/conf"
	"brave-ox-web/pkg/cron"
	"brave-ox-web/web/cache"
	"brave-ox-web/web/services/airdrop"
	"brave-ox-web/web/services/breed"
	"brave-ox-web/web/services/cattle"
	"brave-ox-web/web/services/currency"
	"brave-ox-web/web/services/item"
	"brave-ox-web/web/services/market"
	"brave-ox-web/web/services/metadata"
	"brave-ox-web/web/services/planet"
	"brave-ox-web/web/services/profile"
	"brave-ox-web/web/services/query"
)

func InitService() {
	conf.LoadAllCurrency()
	cache.LoadTemplateData()

	currency.Init()
	market.Init()

	airdrop.Init()
	breed.Init()
	cattle.Init()

	item.Init()

	planet.Init()
	profile.Init()
	metadata.Init()
	query.Init()

}

func Close() {
	cron.StopCron()
}
