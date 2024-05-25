package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutex　（互斥鎖）
// 因為在total進行++時，還沒處理完並存成新的total的情況，這樣就跑了兩次++，但total只加了一次，
// 確保同一時間只有一執行續在操作變數
type SafeNumber struct {
	v   int
	mux sync.Mutex // 互斥鎖
}

func main() {
	fmt.Println("------沒有上鎖------")
	total1 := 0
	for i := 0; i < 1000; i++ {
		go func() { total1++ }()
	}
	time.Sleep(time.Second)
	fmt.Println(total1)

	fmt.Println("------變數上鎖------")
	total2 := SafeNumber{v: 0}
	for i := 0; i < 1000; i++ {
		go func() {
			total2.mux.Lock()
			total2.v++
			total2.mux.Unlock()
		}()
	}
	time.Sleep(time.Second)
	total2.mux.Lock()
	fmt.Println(total2.v)
	total2.mux.Unlock()
}
