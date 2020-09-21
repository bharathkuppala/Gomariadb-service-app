package gormdb

import (
	"fmt"
	"time"

	utility "github.com/mariaDB/module/utilities"
	
	"github.com/mariaDB/module/config"

	gsql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewGormConfig ...
func NewGormConfig(appConfig *config.AppConfig) (*gorm.DB, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		utility.Logger.Println("failed to get location")
		return nil, err
	}
	gconfig := &gsql.Config{
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", appConfig.Db.Host, appConfig.Db.Port),
		DBName:               appConfig.Db.Name,
		User:                 appConfig.Db.Username,
		Passwd:               appConfig.Db.Password,
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  loc,
	}

	var logLevel logger.LogLevel
	if appConfig.Debug {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}
	return gorm.Open(mysql.Open(gconfig.FormatDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	
}
