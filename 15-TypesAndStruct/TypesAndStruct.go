package main

import "fmt"

func main() {
	fmt.Println("\n-----自訂型別(Type)-----")
	myFirstLove, mySecondLove := getMyLoves()
	fmt.Println("我的初戀:", myFirstLove, " 我的第二個戀人:", mySecondLove) 
	fmt.Println(myFirstLove == mySecondLove) 
	fmt.Println("\n-----結構 (struct)-----")
	favorites := getFavorites()
	for i := 0; i < len(favorites); i++ {
		fmt.Println(i, favorites[i])
	}
}

// 定義一個自訂型別
type myLove string 

func getMyLoves() (myLove, myLove) {
	var myFirstLove myLove = "奇犽" // 將自訂型別建立變數
	var mySecondLove myLove       // 自訂型別也擁有零值，所以這邊會是空字串
	return myFirstLove, mySecondLove
}

// struct
// type <結構型別的名稱> struct { <欄位 1> <型別>,<欄位 2> <型別>,<欄位 3> <型別> }
type myFavorite struct { // 定義一個名為 myFavorite 的結構型別
	name     string // 定義結構裡的欄位名稱與型別
	color    string
	isCommon bool
	month    int
}

func getFavorites() []myFavorite {
	userA := myFavorite{
		color:    "yellow",
		isCommon: true,
		name:     "Krystal",
		month:    12,
	}
	userB := myFavorite{
		name:  "Andy",
		month: 8,
	}
	userC := myFavorite{
		"Coco", // 沒有欄位名稱的賦值（一定要照順序）
		"blue",
		false,
		1,
	}
	var userD myFavorite
	userD.name = "Tom" 
	userD.color = "gray"
	userD.isCommon = true
	userD.month = 7

	return []myFavorite{userA, userB, userC, userD}
}
