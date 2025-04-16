package Logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LogInfo(Info string) {
	logFile, err := os.OpenFile("logs/info.logs", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Error opening logs file: %v", err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			fmt.Print("close logs file error...")
		}
	}(logFile)

	// 使用 MultiWriter 同时输出到控制台和日志文件
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// 设置 logs 包的输出目标为 multiWriter
	log.SetOutput(multiWriter)
	log.Printf("[info]:%v", Info)
}

func LogError(Info error) {
	logFile, err := os.OpenFile("logs/error.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Error opening logs file: %v", err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			fmt.Print("close logs file error...")
		}
	}(logFile)

	// 使用 MultiWriter 同时输出到控制台和日志文件
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// 设置 logs 包的输出目标为 multiWriter
	log.SetOutput(multiWriter)
	log.Fatalf("[Error]:%v", err)
}
