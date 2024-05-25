package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Print1(s string) {
	for i := 0; i < 50; i++ {
		time.Sleep(time.Microsecond * 1) // 控制時間搶先Print
		fmt.Print(s)
	}
}

func Print2(s string) {
	for i := 0; i < 50; i++ {
		fmt.Print(s)
	}
}

func Print3(s string) {
	defer wg.Done() // 操作了一次WaitGroup，要讓wg-1
	for i := 0; i < 50; i++ {
		time.Sleep(time.Microsecond * 1)
		fmt.Print(s)
	}
}

// Goroutine 的短寫 go
// 正常情況下goroutine分配給不同的邏輯處理器執行，實現程式的並行 ( 同時執行 )

var wg sync.WaitGroup // 建立全域變數wg

func main() {
	go Print1("X")
	go Print1("O")
	time.Sleep(time.Second * 2)

	fmt.Println("\n-----單處理器併發(分先後次序)------")
	runtime.GOMAXPROCS(1) // 限制能使用處理器=1
	go Print2("X")
	go Print2("O")
	time.Sleep(time.Second * 2)

	fmt.Println("\n-----(sync.WaitGroup)------")

	// Add (i int) : 計數器增加i
	// Done() : 計數器-1，相當於Add(-1)
	// Wait() : 阻塞直到所有的WaitGroup數量變為零，即計數器變為0
	runtime.GOMAXPROCS(1)
	wg.Add(2) // 添加2筆goroutine訂單
	go Print3("X")
	go Print3("O")
	fmt.Println("waiting to finish")
	wg.Wait() // 等待WaitGroup執行完成
	fmt.Println("\nfinish Program")
}
