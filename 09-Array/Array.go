package main

import "fmt"

func array1() [5]int { // func 回傳一個長度為 5 型別為 int 的陣列
	var arr [5]int // 定義一個長度為 5 型別為 int 的陣列(沒有賦予初始值，會自動以零值為其初始值)
	return arr     // 回傳 arr
}
func array2() [3]int { // func 回傳一個長度為 3 型別為 int 的陣列
	arr := [3]int{1, 2, 3} // 定義一個長度為 3 型別為 int 的陣列(有賦予初始值，會以賦予的值為其初始值)
	return arr             // 回傳 arr
}
func array3() [4]int { // func 回傳一個長度為 4 型別為 int 的陣列
	arr := [4]int{9} // 定義一個長度為 4 型別為 int 的陣列 (只有賦予一個初始值，其餘會以零值為初始值)
	return arr       // 回傳 arr
}

// 定義與長度不一致會不能return
func array4() [6]int { // func 回傳一個長度為 6 型別為 int 的陣列
	arr := [...]int{1, 2, 3, 9, 8, 7} // 定義一個長度為 6 型別為 int 的陣列(將長度省略為...)
	return arr                        // 回傳 arr
}

func compareArrays1() (bool, bool, bool) { // func 回傳三個布林值
	arr1 := [...]int{1, 2, 3}
	var arr2 [3]int
	arr3 := [...]int{0, 0, 0}
	arr4 := [3]int{1}
	return arr1 == arr2, arr2 == arr3, arr1 == arr4 // 回傳 arr 是否兩兩相等 （Array 長度需要一致）
}

func murmur() string {
	arr := [...]string{
		"我最愛",
		"水果",
		"芒果",
		"是",
		"的",
	}
	return fmt.Sprintln(arr[0], arr[4], arr[1], arr[3], arr[2]) // fmt.Sprintln 將字串格式化
}

func main() {
	fmt.Println("array1:", array1())
	fmt.Println("array2:", array2())
	fmt.Println("array3:", array3())
	fmt.Println("array4:", array4())
	fmt.Println("\n-----Array比較------")
	compare1, compare2, compare3 := compareArrays1() // 變數メソッド
	fmt.Println("arr1 == arr2:", compare1)
	fmt.Println("arr2 == arr3:", compare2)
	fmt.Println("arr1 == arr4:", compare3)

	fmt.Println("\n-----陣列索引賦值------")
	var (
		arr1 [5]int
		// [...]int{4:0} 0 - 4( 長度為 5 )
		arr2 = [...]int{4: 0}           // 若使用...讓 Go 推算陣列長度，即使索引 4 一樣是零值，還是要賦予初始值，讓 Go 有推算根據
		arr3 = [5]int{4: 40, 2: 20, 30} // 索引賦予初始值不按照順序也 ok ，且若是沒有給予索引，則會自動推算在前一個索引 +1
	)
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("arr3:", arr3)
	fmt.Println(arr1 == arr2) // 雖寫法不同，但實質上兩個陣列內容相同

	fmt.Println("\n-----讀取陣列元素值------")
	fmt.Print(murmur()) // 利用索引值將陣列內的字串，拼成我的 murmur
}
