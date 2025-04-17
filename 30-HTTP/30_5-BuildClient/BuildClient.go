package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/socket", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()
	for {
		time.Sleep(time.Second)
		err = conn.WriteMessage(websocket.TextMessage, []byte("WebSocket Success!"))
		if err != nil {
			log.Println(err)
			return
		}

		_, msg, err := conn.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("receive: %s\n", msg)
	}
}
