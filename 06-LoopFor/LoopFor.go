package main

// GO 是沒有while
// for range 是隨機

import (
	"fmt"
	"math/rand" // math 檔案內的 rand
)

func main() {
	Type1()
	Type2()
	Type3()
}

func Type1() {
	favoriteFruit := []string{"apple", "kiwi", "peach", "mango"}

	for i := 0; i < len(favoriteFruit); i++ { // len() 可以取得任意集合的長度，索引值從 0 起累加，用 len() 的回傳值即可判定迴圈是否走到最後一個值，是否繼續執行或是終止。
		fmt.Println(favoriteFruit[i])
	}
}

func Type2() {
	favoriteFruit := []string{"apple", "kiwi", "peach", "mango"}
	fmt.Println("\n----Type2----")

	myFavorite := map[string]string{
		"fruit":   "mango",
		"animal":  "panda",
		"dessert": "cheesecake",
	}

	for key, value := range myFavorite {
		// or
		// for _, value := range myFavorite
		fmt.Println("\nkey:", key, " = value:", value)
		// or
		fmt.Println(key)
		fmt.Println(value)
	}

	for i, value := range favoriteFruit {
		fmt.Println("index:", i, "value:", value)
	}
}

func Type3() {
	fmt.Println("\n----Type3----")
	// 無限loop 的 FOR
	for {
		i := rand.Intn(10) // 產生 0~10 的亂數
		if i%3 == 0 {      // 若是 3 的倍數
			fmt.Println("i=", i, ",用 continue 來跳過這次迴圈")
			continue
		}

		if i%2 == 0 { // 若是 2 的倍數
			fmt.Println("i=", i, ",用 break 來終止這次迴圈")
			break
		}
		fmt.Println("i=", i)
	}
}
