package main

import "fmt"

// map 是一種雜湊表 (hashmap)
// key 不能是切片
// 定義：map[<key 型別>]<value 型別>{
// 		<key 1>:<value 1>, <key 2>:<value 2>, ..., <key N>:<value N>}

// make定義map：make(map[<key 型別>]<value 型別>, <容量>)

func map_part01() {
	myFavorite := map[string]string{ // 定義一個 map ，且賦予其索引鍵 (key) 跟元素值 (value)來初始化 map
		"fruit":   "mango",
		"animal":  "panda",
		"dessert": "cheesecake",
	}

	myFavorite["weather"] = "sunny" // 使用 [] 加入要新增進 map 的 key 與 value
	fmt.Println("myFavorite", myFavorite)

	for key, value := range myFavorite {
		fmt.Println("key:", key, " = value:", value) // 將 key 與 value 一一印出來
	}

	fmt.Println("\n-----從 map 讀取元素值-----")
	favoriteTitle, exists := myFavorite["fruit"] // 讀取 myFavorite 這個 map 的 key 值 fruit 的 value, exists專用
	if exists {                                  // 如果 fruit 這個 key 有相對應的 value
		fmt.Println("目前列表：")
		for key, value := range myFavorite { // 用 for range 走訪
			fmt.Println("key:", key, " = value:", value)
		}
	} else {
		fmt.Println("查無此品項", favoriteTitle)
	}
	fmt.Println("查詢品項：", favoriteTitle)
}

var myFavorites02 = map[string]string{
	"fruit":   "mango",
	"animal":  "panda",
	"dessert": "cheesecake",
	"weather": "sunny",
}

func DelMyFavorite(title string) {
	delete(myFavorites02, title)
}

func main() {
	fmt.Println("\n-----Part01-----")
	map_part01()
	fmt.Println("\n-----Part02-----")
	favoriteTitle := myFavorites02["fruit"]
	DelMyFavorite(favoriteTitle)
	fmt.Println("我的最愛：", myFavorites02)
}
