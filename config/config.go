package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name         string
		Port         string
		TemplatePath string
	}
	Database struct {
		Host            string
		Port            string
		User            string
		Password        string
		DatabaseName    string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifetime int
	}
	GinMode struct {
		Model string
	}

	Redis struct {
		Host     string
		Port     string
		Password string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}
	// 初始化数据库
	InitDB()
}
