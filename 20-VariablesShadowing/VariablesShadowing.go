package main

import "fmt"

func message() (message string, err error) {
	message = "錯誤"
	if message == "錯誤" {
		err := fmt.Errorf("這是錯誤的訊息呦！")
		return message, err
	}
	return // 沒有在 return 敘述後面指定要傳回的變數， 會自己將傳回值name, age
}

// 所有符合型別的引數，都可放進此參數並打包成切片
func myFavorite1(age int, name ...string) { // name ...string 必須放後面
	fmt.Println(name, age)
}
func myFavorite2(name ...string) {
	fmt.Println(name)
	fmt.Println("len:", len(name), ", cap:", cap(name))
	fmt.Printf("%T\n", name) // 使用 Printf 加上 %T 印出型別
}

func main() {
	fmt.Println(message())
	fmt.Println("\n------pack operator-----")
	myFavorite1(12, "奇犽", "小傑")
	myFavorite1(12, "XXXX")
	fmt.Println("\n------pack operator slice test-------")
	myFavorite2("奇犽", "小傑")
	name2 := []string{"奇犽", "小傑"}
	myFavorite2(name2...) // 用Slice輸入 Pack Operator
}
