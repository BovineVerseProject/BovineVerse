package cron

import (
	"sync"

	"github.com/robfig/cron/v3"

	"brave-ox-web/common"
	"brave-ox-web/log"
)

var jobChain *cron.Cron

var initializer sync.Once

func Init() {
	initializer.Do(func() {
		jobChain = cron.New(cron.WithSeconds(), cron.WithChain(cron.Recover(log.CronPanicLogger())))
	})
}

func AddTimerFunc(spec string, jobFunc func()) {
	_, err := jobChain.AddFunc(spec, jobFunc)
	common.MustNoError(err)
}

func Start() {
	jobChain.Start()
}

func StopCron() {
	jobChain.Stop()
}
