package main

import "fmt"

var favoriteFruit = "kiwi" // 外層宣告 kiwi

func main() {
	fmt.Println("package main favorite:", favoriteFruit)

	favoriteFruit := "apple" // 內favoriteFruit != 外favoriteFruit

	if true {
		favoriteFruit := "peach"                         // 在 if 內更改變數值為 peach
		fmt.Println("if block favorite:", favoriteFruit) // 因為在 if 內有找到favoriteFruit，所以可以直接印出值為 peach
		fruit()
	}
	fmt.Println("func main favorite:", favoriteFruit) // 雖然順序在 if block 下會有點讓人誤會，但是仔細劃分好區域，會發現它的 scope 是在 func main 內的，所以找到值為 apple
}

func fruit() {
	favoriteFruit := "mango" // 改變數值為 mango
	fmt.Println("func fruit favorite:", favoriteFruit)
}
