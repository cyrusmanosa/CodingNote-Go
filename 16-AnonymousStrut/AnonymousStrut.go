package main

import "fmt"

// 沒有名稱
// 在宣告時可一併賦值
// 只能有一個變數，無法再建立其他結構

// <結構變數名稱> := struct {<欄位 1> <型別>,<欄位 2><型別>}{<值 1>,<值 2>}

type myFavorites struct { // 定義名為 myFavorites 的具名結構
	name string
	age  int
}

func compare() (bool, bool) {
	MyFavorite := struct { // 匿名結構（有初始值）
		name string
		age  int
	}{
		"奇犽", // 省略欄位名稱
		12,
	}

	sisFavorite := struct { // 匿名結構（沒有初始值）
		name string
		age  int
	}{}

	sisFavorite.name = "庫拉皮卡" // 透過結構.欄位賦值
	sisFavorite.age = 17
	broFavorite := myFavorites{"奇犽", 12}                        // 具名結構直接賦值
	return MyFavorite == sisFavorite, MyFavorite == broFavorite // 將欄位相同的匿名結構型別與具名結構型別，兩兩相比較
}

func main() {
	sis, bro := compare()
	fmt.Println("MyFavorite == sisFavorite", sis)
	fmt.Println("MyFavorite == broFavorite", bro)
}
