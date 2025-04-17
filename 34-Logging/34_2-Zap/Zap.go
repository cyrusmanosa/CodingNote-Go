package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	// 為記錄器配置（即構建器）
	config := zap.NewProductionConfig()

	// 本地中是一個文件
	config.OutputPaths = []string{"zap-demo.log"}

	// logger 建立 變數
	zapLogger, err := config.Build()

	// Check Error
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	// 刷新緩衝區
	defer zapLogger.Sync()

	// message at the INFO level
	zapLogger.Info("This is our first log message using zap!")

}
