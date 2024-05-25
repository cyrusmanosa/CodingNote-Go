package main

import "fmt"

// 在函式前加入 defer 便會在 main() 函式最後才會執行
// 越先寫越遲做

func main() {
	fmt.Println("在 defer 函式前印出的字")
	defer defer1()
	fmt.Println("在 defer 函式後印出的字")
	defer defer2()
	defer defer3()
	defer defer4()
	func() {
		fmt.Println("我是 main() 裡的匿名函式")
	}()
	name := "奇犽"
	age := 12
	defer myFavorite(name, age)
	name = "西索"
	age = 26
	myFavorite(name, age)

}

func defer1() { fmt.Println("我是 defer 函式 1") }
func defer2() { fmt.Println("我是 defer 函式 2") }
func defer3() { fmt.Println("我是 defer 函式 3") }
func defer4() { fmt.Println("我是 defer 函式 4") }

func myFavorite(name string, age int) {
	fmt.Println("\n-----defer 對變數值的副作用-------")
	fmt.Print("我最喜歡的人是：", name, " 他", age, "歲\n")
}
