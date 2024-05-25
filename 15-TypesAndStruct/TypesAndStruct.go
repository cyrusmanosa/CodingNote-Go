package main

import "fmt"

func main() {
	fmt.Println("\n-----自訂型別(Type)-----")
	myFirstLove, mySecondLove := getMyLoves()
	fmt.Println("我的初戀:", myFirstLove, " 我的第二個戀人:", mySecondLove) // 印出自訂型別的值
	fmt.Println(myFirstLove == mySecondLove)                     // 比對自訂型別的值
	fmt.Println("\n-----結構 (struct)-----")
	favorites := getFavorites()
	for i := 0; i < len(favorites); i++ {
		fmt.Println(i, favorites[i])
	}
}

// type <自訂型別名稱> <核心型別>
type myLove string // 定義一個自訂型別
// type myFavorite func()  // 自訂一個名為 myFavorite 的函式型別，但不需要輸入參數也沒有傳回值
// type myFavorite func(string) string  // 自訂一個名為 myFavorite 的函式型別，接收一個字串參數，以及傳回一個字串型別的值

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
		color:    "yellow", // 有欄位名稱的賦值（可不用照順序）
		isCommon: true,
		name:     "Krystal",
		month:    12,
	}
	userB := myFavorite{
		name:  "Andy", // 有欄位名稱的賦值，沒有賦值的欄位會是零值
		month: 8,
	}
	userC := myFavorite{
		"Coco", // 沒有欄位名稱的賦值（一定要照順序）
		"blue",
		false,
		1,
	}
	var userD myFavorite // 先定義一個 userD 結構變數
	userD.name = "Tom"   // 使用 結構變數.欄位名稱 來賦值
	userD.color = "gray"
	userD.isCommon = true
	userD.month = 7

	return []myFavorite{userA, userB, userC, userD}
}
