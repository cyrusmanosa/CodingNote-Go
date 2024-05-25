package main

import "fmt"

func main() {

	// 指標變數
	fmt.Println("-----指標變數-----")
	var name *string // 宣告 name 指標變數，初始值為 nil
	age := new(int)
	height := 160
	myHeight := &height

	if name != nil {
		fmt.Printf("name: %#v\n", *name)
	} // 因為初始值為 nil，所以 if 不成立，不會印出
	if age != nil {
		fmt.Printf("age: %#v\n", *age)
	}
	if myHeight != nil {
		fmt.Printf("myHeight: %#v\n", *myHeight)
	} // myHeight = 0(nil)

	// 列舉
	fmt.Println("------列舉------")
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
