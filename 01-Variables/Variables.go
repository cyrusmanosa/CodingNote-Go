package main

import (
	"fmt"
	"math"
	"math/big"
)

// var MAN int      => 0
// var KA bool      => false
// var CHUN string  => " "

/*
	bool 布林 不是"真"就是"假" 真的用True 假的用False
	byte 位元數 不帶正負號的整數類型
	rune 由 Go 語言定義的特殊類型 專門用來存 Unicode 編碼 或 32位元的2進位帶有號的整數
	int / unit 有號整數類 / 不帶正負號的整數類
	int8 / uint8 8位元2進位整數類 / 不帶正負號的整數類
	int16 / uint16 16位元2進位整數類 / 不帶正負號的整數類
	int32 / uint32 32位元2進位整數類 / 不帶正負號的整數類
	int64 / uint64 64位元2進位整數類 / 不帶正負號的整數類
	float32 32位元2進位的浮點數類
	float64 64位元2進位的浮點數類
	complex64 64位元2進位的複數類
	complex128 128位元2進位的複數類
	string 一個字串類代表一個字串值的集合
*/

const numA int = 100

func main() {
	// GO 自動判斷值的變數型態
	屌你老尾 := "fuckyou"
	name := "Cyrus"
	j := 100

	fmt.Println(j + numA)
	fmt.Println("Hello world \n And ") // \n 分行
	fmt.Println(屌你老尾)
	fmt.Println(len(name))
	fmt.Println(name[1])
	fmt.Println(name[0:2])
	fmt.Println(name[:3])
	fmt.Println(name[2:])

	fmt.Println("\n-----浮點數-----")

	var x int = 1000
	var y float32 = 1000 //準確度較 float64 低（小數點）
	var z float64 = 1000 //準確度較高

	fmt.Println("int:", x/7)    // 整數除整數得整數
	fmt.Println("float32", y/7) // 整數除浮點數得浮點數
	fmt.Println("float64", z/7)

	fmt.Println("\n-----特別例01-----")
	fmt.Println("int:", x/7*7/7*7)
	fmt.Println("float32", y/7*7/7*7)
	fmt.Println("float64", z/7*7/7*7)

	fmt.Println("\n-----特別例02-----")

	var w uint8 = 253

	for i := 0; i < 5; i++ {
		w++
		fmt.Println("uint8:", w) //超出範圍
	}

	fmt.Println("\n-----大數值-----")

	maxInt := math.MaxInt64 // 最大 int 整數
	maxInt = maxInt + 1
	bigInt := big.NewInt(math.MaxInt64)
	bigInt.Add(bigInt, big.NewInt(1))

	fmt.Println("原本的maxInt:", math.MaxInt64)
	fmt.Println("加1後的maxInt:", maxInt)
	fmt.Println("加1後的bigInt:", bigInt)
}
