package main

import (
	"fmt"
	"os"
)

// Slice 不能被複製
// 使用 os.Args 將我們在 go run 後面輸入的參數，加到新切片中
// os.Args 是數量不定的字串Slice參數
// os.Args 的第一個參數是執行的程式名稱(No Value)，第二個參數起才是真的參數。
// 從長度7的 Array 做出的 Slice [1:3]，長度為 2，容量為 6. 因為是從1到3，所以0番地會被捨棄

func getArgs0() (string, string) {
	args := os.Args
	A := fmt.Sprintln("A:", args)
	fmt.Println("0番地是執行Programme的Location")
	B := fmt.Sprintln("B(番地:1):", args[1])
	return A, B
}

func getArgs1() []string {
	if len(os.Args) < 1 { // 確保os.Args不是空
		fmt.Println("您沒有輸入參數,至少需傳入1個參數")
		os.Exit(1) // 若沒有傳入參數就強制結束程式
	}

	var argsSlice []string // 建立一個名為 argsSlice 的空切片

	for i := 1; i < len(os.Args); i++ { // 因為 os.Args 的第一個參數是執行的程式名稱，所以迴圈從索引 1 開始走訪
		argsSlice = append(argsSlice, os.Args[i]) // 將我們在 go run 後面輸入的參數，使用 os.Args 一一使用 append() 加到新切片中
	}

	return argsSlice // 回傳新切片
}

func getArgs2() []string {
	if len(os.Args) < 1 {
		fmt.Println("您沒有輸入參數，至少需傳入1個參數")
		os.Exit(1) // 強制結束程式 ， 1 是錯誤代碼
	}

	var argsSlice []string              // 建立一個名為 argsSlice 的空切片
	for i := 1; i < len(os.Args); i++ { // 因為 os.Args 的第一個參數是執行的程式名稱，所以迴圈從索引 1 開始走訪
		argsSlice = append(argsSlice, os.Args[i]) // 將我們在 go run 後面輸入的參數，使用 os.Args 一一使用 append() 加到新切片中
	}
	return argsSlice // 回傳新切片
}

// 將傳入的切片內的元素一一遍歷，找出長度最長的字串
func addMultipleArgs(argsSlice []string) []string {
	var multipleArgsString []string
	multipleArgsString = append(multipleArgsString, argsSlice...) // 在 argsSlice 切片後方加上 ... 來解開切片將值一一傳入 (argsSlice 內沒有東西，所以沒有分別)
	multipleArgsString = append(multipleArgsString, "水蜜桃", "哈密瓜") // 再一次傳入多個值到切片中
	fmt.Println("容量：", cap(multipleArgsString))                   // Cap() 可以check Slice 容量長度
	return multipleArgsString
}

// 將傳入的切片內的元素一一遍歷，找出長度最長的字串
func findLongestString1(argsSlice []string) string {
	var longestString string // 建立一個字串型別的空字串
	for i := 0; i < len(argsSlice); i++ {
		if len(argsSlice[i]) > len(longestString) { // 如果argsSlice 索引值內的字串長度 > longestString 這個字串
			longestString = argsSlice[i] // 將 longestString 的值改為 argsSlice[i] 的值
		}
	}
	return longestString // 回傳長度最長的字串
}

func Slice0() string {
	m := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("原本的陣列或切片:", m)
	message := fmt.Sprintln("第一個元素:", m[0], m[0:1], m[:1])
	message += fmt.Sprintln("最後一個元素:", m[len(m)-1], m[len(m)-1:])
	message += fmt.Sprintln("前兩個元素:", m[0:2], m[:2])
	message += fmt.Sprintln("倒數兩個元素:", m[len(m)-2:])
	message += fmt.Sprintln("第 2 到第 6 個元素:", m[2:7])
	message += fmt.Sprintln("全部的元素:", m[:])
	return message
}

// 如果Slice arr2 內 增加值 [3 6 0 4] ＝》[4 7 1 5]， 原Array arr1 都會被影響 [7 3 6 0 4 9 10] ＝》[7 4 7 1 5 9 10]
func Slice1() {
	var arr1 [7]int = [7]int{7, 3, 6, 0, 4, 9, 10}
	fmt.Print("Array Data")
	fmt.Println("arr1: ", arr1, "\nLen: ", len(arr1), "\nCap: ", cap(arr1))
	fmt.Println(" ")
	///
	var arr2 []int = arr1[1:3]
	fmt.Print("Slice Data")
	fmt.Println("arr2: ", arr2, "\nLen: ", len(arr2), "\nCap: ", cap(arr2))
	fmt.Println(" ")
	///
	arr2 = arr2[0 : len(arr2)+2]
	fmt.Println("arr2 增加自身", arr2, "\nLen: ", len(arr2), "\nCap: ", cap(arr2))
	fmt.Println(" ")
	///
	for k := range arr2 {
		arr2[k] += 1
	}
	fmt.Println("updated arr2: ", arr2)
	fmt.Println("updated arr1(被影響): ", arr1)
	fmt.Println(" ")
}

func main() {
	fmt.Println("\n-----os.Args Basic1-----")
	fmt.Println(getArgs0()) // 印出不同
	fmt.Println("\n-----os.Args Basic2-----")
	fmt.Println(getArgs1())                     // 印出新切片
	fmt.Println(findLongestString1(getArgs1())) // 找出切片內長度最長的字串元素
	fmt.Println("\n-----append-----")
	fmt.Println(getArgs2())                  // 印出新切片
	fmt.Println(addMultipleArgs(getArgs2())) // 找出切片內長度最長的字串元素
	fmt.Println("\n----------------Slice0---------------------")
	fmt.Println(Slice0())
	fmt.Println("\n----------------Slice1---------------------")
	Slice1()
}
