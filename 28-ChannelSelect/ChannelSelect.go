package main

import (
	"fmt"
	"time"
)

// Run 後發現有時候拿到 random 01 有時候拿到 random 02
// select 的特性之一，case 是隨機選取，所以當 select 有兩個 channel 以上時，如果同時對全部 channel 送資料，則會隨機選取到不同的 Channel。

// BasicSelect結果： 隨機
// SelectMistake結果：Error
// CheckChannel結果：channel value is 1, channel value is 2
func main() {
	// BasicSelect()
	// SelectMistake()
	// SelectTimeout()
	// CheckChannel()
	LoopSelect()
}

func BasicSelect() {
	ch := make(chan int, 1)
	ch <- 1

	ch2 := make(chan int, 1)
	ch2 <- 1

	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch2:
		fmt.Println("random 02")
	}
}

// 沒有送 value 進去 Channel 就會造成 panic
func SelectMistake() {
	ch := make(chan int, 1)
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
		// default:							// 改正Error
		//     fmt.Println("exit")
	}
}

func SelectTimeout() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	ch := make(chan int)
	select {
	case <-ch: // 不會因為處理沒有結果的東西而等留，一定時間後會處理其他
	case <-timeout:
		fmt.Println("timeout 01")
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 02")
	}
}

func CheckChannel() {
	// Channel 同 Select 的容量需一致。就算Channel的容量較大，Run的時候都會被提醒
	ch := make(chan int, 2)
	ch <- 1
	select {
	case ch <- 2:
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)
	default:
		fmt.Println("channel blocking")
	}
}

func LoopSelect() {
	i := 0
	ch := make(chan string, 0)
	defer func() { close(ch) }()
	go func() {
	LOOP: // for 的 變數
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Unix()) // Unix() 取得當前時間戳記（秒）
			i++
			select {
			case m := <-ch:
				println(m)
				break LOOP
			}
		}
	}()
	time.Sleep(time.Second * 4)
	ch <- "stop"
}
