package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"brave-ox-web/common"
	"brave-ox-web/conf"
	"brave-ox-web/config"
	"brave-ox-web/db"
	"brave-ox-web/log"
	"brave-ox-web/pkg"
	"brave-ox-web/pkg/cron"
	"brave-ox-web/web/endpoint"
	"brave-ox-web/web/sensitive"
	"brave-ox-web/web/services"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	sysType := runtime.GOOS
	switch sysType {
	case "windows", "darwin":
		config.Parse("./config/sys.toml.list")
	case "linux":
		if len(os.Args) < 2 {
			panic("miss config file")
		}
		config.Parse(os.Args[1])
	default:
		panic(fmt.Sprintf("error system:%s", sysType))
	}

	// if len(os.Args) < 2 {
	// 	panic("miss config file")
	// }
	// config.Parse(os.Args[1]) // 系统配置

	conf.Parse(config.ServiceConfig())
	log.InitLogger("web")

	db.Connect(config.IsShowSQL())

	endpoint.Init()

	common.MustNoError(sensitive.InitSensitiveFilter(conf.SensitivePath()))

	cron.Init()
	services.InitService()
	cron.Start()

	s := newServer()
	go s.start()

	if config.Pprof() != "" {
		go func() {
			log.Info(http.ListenAndServe(config.Pprof(), nil))
		}()
	}

	defer func() {
		s.stop()
		services.Close()
		log.Info("server stopped")
		log.Sync()
	}()
	pkg.ShutdownHook()
}
