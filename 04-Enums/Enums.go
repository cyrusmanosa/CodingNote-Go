package main

import "fmt"

func main() {
	// 列舉
	const (
		zero = iota // 宣告一個 const 時，iota 預設值是 0,
		one         // 第二個宣言會自動加一，第三個再加一
		two
		three
		four
	)
	const (
		Sunday = iota
		Monday
		Tuesday
	)

	fmt.Println(zero, one, two, three, four, Sunday, Monday, Tuesday)

	//變數作用範圍 (Scope)
	fmt.Println("-----變數作用範圍 (Scope)-----")
}
