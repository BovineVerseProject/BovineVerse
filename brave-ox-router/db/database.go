package db

import (
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"brave-ox-web/common"
	"brave-ox-web/config"
	customLog "brave-ox-web/log"
)

var mainDB *gorm.DB
var _once sync.Once

func Connect(showSQL bool) {
	_once.Do(func() {

		SQLLogger := logger.New(log.New(customLog.LoggerWriter(), "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: false,
		})

		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  config.DatabaseStr(),
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{
			Logger: SQLLogger,
		})

		common.MustNoError(err)
		tmp, _ := db.DB()
		tmp.SetMaxIdleConns(10)
		tmp.SetConnMaxIdleTime(5 * time.Minute)
		tmp.SetMaxOpenConns(40)

		if showSQL {
			db = db.Debug()
		}
		mainDB = db
	})

}

func MainDB() *gorm.DB {
	return mainDB
}
