package config

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"webproject/global"
)
import "gorm.io/driver/mysql"

// 初始化db

func InitDB() {
	Host := AppConfig.Database.Host
	Port := AppConfig.Database.Port
	User := AppConfig.Database.User
	Password := AppConfig.Database.Password
	DatabaseName := AppConfig.Database.DatabaseName
	// 拼接 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		User, Password, Host, Port, DatabaseName)

	// 使用 GORM 打开 MySQL 数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Successfully connected to database")

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect to database")
	}
	// 最大闲置数
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	// 最大打开数
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	// 最大链接时间
	sqlDB.SetConnMaxLifetime(time.Duration(AppConfig.Database.ConnMaxLifetime) * time.Hour)
	global.GlobalDB = db
}
