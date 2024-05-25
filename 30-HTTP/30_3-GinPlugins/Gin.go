package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Title   string
	Content string
}

func test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	server := gin.Default()           // 設定 gin 的啟動程式
	server.LoadHTMLGlob("template/*") // 檔案內要開一個Template目錄
	server.GET("/", test)             // 設定 http routing
	server.Run(":8880")
}
