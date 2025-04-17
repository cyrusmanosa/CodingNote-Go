package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n-----NEW------")
	make0()
	fmt.Println("\n-----MAKE1-----")
	make1()
	fmt.Println("\n-----MAKE2-----")
	make2()
}

// new(#) 宣告會直接拿到儲存位置，並且配置 Zero Value (初始化)，也就是數字型態為 0
// new 可以快速的達到初始化
// new 是不能加入特定的初始化值

type SyncedBuffer struct {
	foo int
	bar string
}

func make0() {
	p := new(SyncedBuffer)
	k := SyncedBuffer{}
	fmt.Println("foo:", p.foo)
	fmt.Println("bar:", p.bar)
	fmt.Println("------------ P --------------")
	fmt.Printf("p 值: %#v\n", p)
	fmt.Printf("p type: %T\n", p)
	fmt.Println("------------ K --------------")
	fmt.Printf("k 值: %#v\n", k)
	fmt.Printf("k type: %T\n", k)
}

// make() 長度參數是必填設定
// 1. 一個隱藏陣列可能同時被多個切片指向 ＝》可能會 連同一個底層陣列的其他切片元素值一同被更改
// 2. 當切片擴充到超過本來隱藏陣列容量時，隱藏陣列指向的地方會不同
// 3. make 用於三個地方，分別是 slice, map 及 channel
// 4. 在沒有return type 的情況下，都可以改內容或增加內容

func make1() {
	// make
	foobar := make(map[string]string)
	foobar["foo"] = "bar"
	foobar["bar"] = "foo"
	fmt.Println(foobar)
}

func make2() {
	var slice1 []string
	slice2 := make([]bool, 5)    // 用 make() 函式建立 bool 型別，長度為 5 的切片( 沒設容量,容量預設跟長度一樣為 5 )
	slice3 := make([]int, 5, 10) // 用 make() 函式建立 int 型別，長度為 5 ，容量為 10 的切片

	fmt.Println("slice1:", slice1, " len:", len(slice1), " cap:", cap(slice1))
	fmt.Println("slice2:", slice2, " len:", len(slice2), " cap:", cap(slice2))
	fmt.Println("slice3:", slice3, " len:", len(slice3), " cap:", cap(slice3))
}
