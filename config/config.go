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
	err = viper.Unmarshal(AppConfig)
	// 初始化数据库
	InitDB()

	//redisClient := redis.NewClient(&redis.Options{
	//	Addr: AppConfig.Redis.Host + ":" + AppConfig.Redis.Port,
	//})
	//
	//// 初始化redis
	//global.Handler = &utils.ContactHandler{
	//	Redis:        redisClient,
	//	MySQL:        global.GlobalDB,
	//	DedupSetKey:  "contact:dedup", // Set for deduplication
	//	QueueListKey: "contact:queue", // List for ordered queue
	//}

	//ctx, _ := context.WithCancel(context.Background())
	// mysql异步刷新
	//go global.Handler.ProcessPendingContacts(ctx)
}
