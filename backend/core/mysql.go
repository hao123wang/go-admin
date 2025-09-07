// 初始化 Mysql

package core

import (
	"fmt"
	"go-admin-server/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	config := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.DB)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(getLevel(config.LogLevel)),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	return db
}

func getLevel(level string) logger.LogLevel {
	switch level {
	case "Silent", "silent":
		return logger.Silent
	case "Error", "error":
		return logger.Error
	case "Info", "info":
		return logger.Info
	case "Warn", "warn":
		return logger.Warn
	default:
		return logger.Info
	}
}
