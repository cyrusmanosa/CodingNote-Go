package main

import (
	"fmt"
	"reflect"
)

// 不能是核心型別 (string, bool) 、介面型別 (interface) 或指標型別
// 方法(method)與該型別必須定義在同一個套件中。

// 加上 * 宣告一個指標變數

type id int // 自訂一個名為 name 的字串型別

type hunter struct { // 定義名為 hunter 的結構型別
	role    string
	ability string
}

type hunter2 struct { // 定義名為 hunter 的結構型別
	role    string
	ability string
}

func (i id) printID() { // 值接收器寫法， i 為接收器變數， id 是結構型別方法
	fmt.Println("id:", i) // 印出 id
}

func (h *hunter) setHunter(role, ability string) { // 指標接收器寫法， h 為接收器變數， hunter 是結構方法
	h.role = role // 存取結構欄位
	h.ability = ability
}

func (h hunter) getHunter() string { // 值接收器寫法， h 為接收器變數， hunter 是結構型別方法
	return fmt.Sprintf("(%v, %v)", h.role, h.ability)
}

func (h2 *hunter2) setHunter2(role, ability string) { // 指標接收器寫法， h 為接收器變數， hunter 是結構方法
	h2.role = role // 存取結構欄位
	h2.ability = ability
}

func (h2 hunter2) getHunter2() string {
	return fmt.Sprintf("(%v, %v)", h2.role, h2.ability)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// set multiple return type and input type
func Function1[V int | float64 | int64 | float32 | int32](t V) V {
	return t + 1
}

// input multiple type. 但return 都是要Set
func Function2[V1 int | float64 | int64 | float32 | int32, V2 int | float64 | int64 | float32 | int32](t1 V1, t2 V2) V1 {
	return t1 + 1
}

func main() {
	var i id = 1 // 定義 i 變數，且賦值
	i.printID()  // 呼叫 i 的方法 printID()

	a, b := hunter{}, hunter{}
	a.setHunter("奇犽", "變化系")               // 呼叫 a 的 setHunter 方法，且帶入參數
	b.setHunter("小傑", "強化系")               // 呼叫 b 的 setHunter 方法，且帶入參數
	fmt.Println("hunter1:", a.getHunter()) // 呼叫 a 的 getHunter 方法
	fmt.Println("hunter2:", b.getHunter()) // 呼叫 b 的 getHunter 方法

	fmt.Println("\n-----接收器值與指標的自動轉換------")
	c := hunter2{}
	d := &hunter2{} //指標 &h 《＝》 接收器 h
	c.setHunter2("奇犽", "變化系")
	d.setHunter2("小傑", "強化系")
	fmt.Println("hunter1:", c.getHunter2())
	fmt.Println("hunter2:", d.getHunter2())

	fmt.Println("\n-----適用閉包（closure）------")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
	fmt.Println("\n---------------Function1-----------------------")
	var t1 int = 123
	fmt.Printf("Function1: %v (type: %s)\n", Function1(t1), reflect.TypeOf(Function1(t1)))
	var t2 float32 = 123.123
	fmt.Printf("Function2: %v (type: %s)\n", Function2(t1, t2), reflect.TypeOf(Function2(t1, t2)))

}
