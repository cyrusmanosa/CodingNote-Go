package main

import (
	"fmt"
)

func main() {
	var name *string // 用 var 宣告 name 指標變數，初始值為 nil (not 0) - 特殊型態

	if name == nil {
		fmt.Println("OMG 這個 name 是 nil 值")
	} else {
		fmt.Println("恭喜你成功宣告一個指標")
	}
}
