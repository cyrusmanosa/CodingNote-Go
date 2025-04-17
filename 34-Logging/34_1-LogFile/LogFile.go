package main

import (
	"log"
	"os"
)

func main() {
	// Create Log 檔
	f, err := os.OpenFile("golang-demo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)

	// log print 時間，日期，內容
	log.Println("This is our first log message in Go.")

	defer f.Close()
}

// Logging Level

// log.Trace("一些非常低的級別。")
// log.Debug("有用的調試信息。")
// log.Info("發生了一些值得注意的事情！")
// log.Warn("你應該看看這個。")
// log.Error("有些事情失敗了，但我沒有放棄。")

// log.Fatal("logout用")
// log.Panic("出左事用")
