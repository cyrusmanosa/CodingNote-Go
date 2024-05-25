package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket是一種網路傳輸協定，使得客戶端和伺服器交換資料變得更簡單，允許伺服器端主動向客戶端推播資料。
// WebSocket中，只需要完成一次交握，兩者之間就可以建立永續性的連接，並進行雙向資料傳輸。

// 建立server端

var upgrader = websocket.Upgrader{}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// 將原始 HTTP 連接升級為基於 websocket 的連接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer func() {
		log.Println("disconnect !!")
		conn.Close()
	}()

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Index Page") }

func main() {
	http.HandleFunc("/socket", socketHandler)
	log.Println("server start at :8080")
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
