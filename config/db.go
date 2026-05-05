package config

import (
	"fmt"
	"time"
	"webproject/global"
	"webproject/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化db

func InitDB() {
	Host := AppConfig.Database.Host
	Port := AppConfig.Database.Port
	User := AppConfig.Database.User
	Password := AppConfig.Database.Password
	DatabaseName := AppConfig.Database.DatabaseName

	if Host == "" || Host == "localhost" || Host == "127.0.0.1" {
		fmt.Println("Database not configured, running in demo mode")
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		User, Password, Host, Port, DatabaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Warning: Failed to connect to database: %v\n", err)
		fmt.Println("Running in demo mode without database")
		return
	}

	fmt.Println("Successfully connected to database")

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Warning: Failed to get database instance: %v\n", err)
		return
	}

	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(AppConfig.Database.ConnMaxLifetime) * time.Hour)
	global.Db = db

	err = global.Db.AutoMigrate(&utils.ContactPostData{})
	if err != nil {
		fmt.Printf("Warning: Failed to migrate ContactPostData: %v\n", err)
	}
	err = global.Db.AutoMigrate(&utils.SubscribeData{})
	if err != nil {
		fmt.Printf("Warning: Failed to migrate SubscribeData: %v\n", err)
	}
}
