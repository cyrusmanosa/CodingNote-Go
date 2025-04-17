package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 0
	for i := 0; i < 1000; i++ {
		go func() { ch1 <- (<-ch1 + 1) }()
	}
	time.Sleep(time.Second)
	fmt.Println(<-ch1)

	// select讓程式在進入等待期時也能有一點輸出。
	fmt.Println("\n--------Unbuffered、Buffered(緩衝區) Channel---------")
	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Second) //模擬費時運算
		ch2 <- "FINISH"
	}()
	for {
		select {
		case (<-ch2): // Channel 中有資料執行此區域
			fmt.Println("main完成")
			return
		default: // Channel 阻塞的話執行此區域
			fmt.Println("waiting...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
