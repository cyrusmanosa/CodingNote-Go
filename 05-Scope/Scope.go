package main

import "fmt"

var favoriteFruit = "kiwi" // 在最外層宣告最愛的水果是 kiwi

func main() {
	fmt.Println("package main favorite:", favoriteFruit) // 程式碼執行到這時，func main 還沒有同名 favoriteFruit 這個變數，所以找到外層的 kiwi

	favoriteFruit := "apple" // 在 func main 內更改變數值為 apple
	if true {
		favoriteFruit := "peach"                         // 在 if block 內更改變數值為 peach
		fmt.Println("if block favorite:", favoriteFruit) // 因為在 if block 內有找到favoriteFruit，所以可以直接印出值為 peach
		fruit()                                          // 呼叫 func fruit
	}
	fmt.Println("func main favorite:", favoriteFruit) // 雖然順序在 if block 下會有點讓人誤會，但是仔細劃分好區域，會發現它的 scope 是在 func main 內的，所以找到值為 apple
}

func fruit() {
	favoriteFruit := "mango"                           // 在 func fruit 內更改變數值為 mango
	fmt.Println("func fruit favorite:", favoriteFruit) // 這邊 func fruit 又是自己的 scope ，無論在哪裡呼叫，他會先找到最近的 favoriteFruit ，也就是 mango
}
