package main

import "fmt"

// 沒有函式名稱
// 只能在其他函式內被宣告
// 只能使用一次
// 可以實作閉包
// 可以搭配 defer 敘述延後執行程式碼

func addFunc() func() int {
	i := 0              // 定義 i 為匿名函式的父函式內的變數
	return func() int { // 回傳一個匿名函式
		i++
		return i
	}
}

func main() {
	myFavorite := "奇犽"
	func(str string) {
		fmt.Println(str)
	}(myFavorite) // 傳給 func(str string)

	fmt.Println("\n--------變數來呼叫匿名函式---------")
	f := func() { fmt.Println("變數裡的匿名函式") }
	fmt.Println("我是本來 main 函式裡要印出的字") // func 是逐行執行，所以會先印出這句
	f()                               // 在此行才通過變數呼叫匿名函式，這時控制權才轉向上方的匿名函式，進而印出函式內的字

	fmt.Println("\n--------語意閉包(lexical closure) ---------")
	// 閉包即使離開父函式執行範圍，也能記住父函式變數的特性。
	add := addFunc() // 接收傳回的 addFunc 函式
	fmt.Println(add())
	fmt.Println(add())
	fmt.Println(add())
}
