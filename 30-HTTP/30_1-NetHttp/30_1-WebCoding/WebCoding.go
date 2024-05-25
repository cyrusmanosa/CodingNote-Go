package main

import (
	"log"      // log 輸出程式目前執行狀態，log 可以分成不同等級
	"net/http" // net/http將網頁運行，包含伺服器監聽與運行的方法與 request handler 都使用此 package
)

// --http.ResponseWriter, r *http.Request（資料驗證）
// request handler 中回傳 html

// 在response內回傳html格式的文字
func test(w http.ResponseWriter, r *http.Request) {
	str := `<!DOCTYPE html>
<html>
<head><title>首頁</title></head>
<body><h1>首頁</h1><p>我的第一個首頁v1</p></body>
</html>
`
	w.Write([]byte(str))
}

// --ListenAndServe
// 開始運行伺服器 => http.ListenAndServe( "address（URL or POST）" , (handle) )
// 可用作Check Server
func main() {
	http.HandleFunc("/", test)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
