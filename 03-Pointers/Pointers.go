package main

import "fmt"

// & 對變數取得記憶體位置
// * 對記憶體位置獲取存放的數值
// 指針：儲存記憶體地址，禁止運算，安全高效。
// 大型數據：指針顯著提升性能（避免複製），支援快速修改和共享。
// 小型數據：指針和結構體（值）效果接近，但指針在需要修改或共享時更有優勢。
// Go 的定位：指針是值語義的補充，簡單安全，專注於引用語義和性能優化。

func plus_10(a *int) {
	*a += 10
}

func main() {
	// 指標變數
	fmt.Println("-----指標變數-----")

	// name 指標變數，初始值為 nil
	var name *string
	age := new(int)
	height := 160
	myHeight := &height // <變數 1> := &<變數 2>

	/// 因為初始值為 nil，所以 if 不成立，不會印出
	if name != nil {
		fmt.Printf("name: %#v\n", &name)
	}
	if age != nil {
		fmt.Printf("age: %#v\n", &age)
	}

	// myHeight = 0(nil)
	if myHeight != nil {
		fmt.Printf("myHeight: %#v\n", &myHeight)
	}

	fmt.Println("---------------")
	score := 75
	plus_10(&score)
	fmt.Println(score) //85
}
