package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Println("----------------------------")
	Sprintf()
	fmt.Println("----------------------------")
	PrintF()
}

// Sprint => Print 但return String
func Sprintf() {
	quantity := 10
	fruits := "apples"
	text := fmt.Sprint("I have ", quantity, " ", fruits, ".")
	fmt.Printf("Value stored in text is: '%v'\n", text)
}

func PrintF() {
	// GO 自動判斷值的變數型態
	const numA int = 100 // 常數
	name := "Cyrus"
	P, R, j := 23.123, 20, 100
	win := true
	ne := complex(2.5, 3.1) // 複數=>實部（第一）和虛部（第二）

	// Printf格式化輸出、格式化樣板語言
	// PrintF (數字)
	fmt.Println("-----------PrintF(數字)-----------------")

	fmt.Println(real(ne))      // 實部（第一）
	fmt.Println(imag(ne))      // 虛部（第二）
	fmt.Println(cmplx.Abs(ne)) // ne的絕對值
	fmt.Printf("i have value: %v and type: %T \n", j, j)
	fmt.Printf("%v\n", R)   // %v   =>  j的值，%T J 的型態
	fmt.Printf("%b\n", R)   // %b   =>	 2進
	fmt.Printf("%d\n", R)   // %d   =>  10進
	fmt.Printf("%+d\n", R)  // %+d  =>  10進 ＋-
	fmt.Printf("%o\n", R)   // %o   =>  8進 0o
	fmt.Printf("%x\n", R)   // %x   =>  16進 英-細草
	fmt.Printf("%X\n", R)   // %X   =>  16進 英-大草
	fmt.Printf("%#x\n", R)  // %#x  =>  16進 0x
	fmt.Printf("%2.x\n", R) // %#x  =>  16進 兩位數字
	fmt.Printf("%4b\n", R)  // %4d  =>  對齊右的4位表示
	fmt.Printf("%-4b\n", R) // %-4d =>  對齊左
	fmt.Printf("%04b\n", R) // %04d =>  補04位

	// (中文)
	fmt.Println("-----------PrintF(中文)-----------------")

	fmt.Printf("%s\n", name)   // %s    =>  淨字
	fmt.Printf("%q\n", name)   // %q    =>  字, 雙引號
	fmt.Printf("%8s\n", name)  // %8s   =>  對齊右的8位表示
	fmt.Printf("%-8s\n", name) // %-8s  =>  對齊左
	fmt.Printf("%x\n", name)   // %x    =>  將字轉16進
	fmt.Printf("%X\n", name)   // %X    =>  將字轉16進 空香位分開表示

	// (真偽)
	fmt.Println("-----------PrintF(真偽)-----------------")
	fmt.Printf("%t\n", win)

	// (浮動數)
	fmt.Println("-----------PrintF(浮動數)-----------------")
	fmt.Printf("%e\n", P)    // %e    => e指數的科學記數法表示
	fmt.Printf("%f\n", P)    // %f    => 小數點，無指數
	fmt.Printf("%.2f\n", P)  // %.2f  => 小數點後兩個位
	fmt.Printf("%6.2f\n", P) // %6.2  => 6個字位表示，小數點後兩個位
	fmt.Printf("%g\n", P)    // %g    => 指數根據需要，只表示必要的數字
}
