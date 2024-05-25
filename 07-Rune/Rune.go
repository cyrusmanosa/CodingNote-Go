package main

import "fmt"

// 特定文字的切片

func main() {
	myName := "Krystal_奇犽老婆" // 我的名字是 12 個字元組成（含底線）

	for i := 0; i < len(myName); i++ {
		fmt.Print(myName[i]) // String 配列型態Print ＝》 byte
		fmt.Printf("%T", myName)
	}

	fmt.Print("\n")
	fmt.Println("\n-----再轉文字型態-----")

	for i := 0; i < len(myName); i++ {
		fmt.Print(string(myName[i]))
	}

	fmt.Print("\n")
	fmt.Println("\n-----使用rune-----")
	runes := []rune(myName) // 先把 myName 字串轉成 rune 切片()

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i])) // 再把剛剛轉為 rune 切片轉字串印出
	}
	fmt.Println(" ")
	fmt.Println("bytes:", len(myName))         // 字串長度 bytes 數
	fmt.Println("runes:", len([]rune(myName))) // rune 集合字數長度
}
