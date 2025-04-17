package main

import (
	"fmt"
)


// copy(<目標切片>, <來源切片>)

var slice1 = []int{1, 2, 3, 4, 5} 
var slice4 = []int{1, 2, 3, 4, 5}
var slice6 = []int{1, 2, 3}

func sameHideArray() (int, int, int) {
	slice2 := slice1    // 複製 slice1 切片為 slice2 (會使兩者指向同一個隱藏陣列)
	slice3 := slice1[:] // 從 slice1 擷取原本長度的切片為 slice2 (會使兩者指向同一個隱藏陣列)

	slice2[0] = 77 // 更改 slice2 切片索引值 0 的值為 77
	return slice1[0], slice2[0], slice3[0]
}


func appendToOriginArray01() (int, int) {
	fmt.Println("Slice4 : ", slice4)
	slice5 := make([]int, 5, 10) // 定義一個長度為 5 ，容量為 10 的初始化切片 slice5
	copy(slice5, slice4)         // 使用內建函式 copy() 將 slice5 當作目標切片， slice4 的值複製進 slice5 ，因為內建函式 copy() 會使其用不同隱藏陣列
	fmt.Println("SLice4 copy to Slice5: ", slice5)

	slice5[0] = 77 // 更改 slice5 切片索引值 0 的值為 77
	return slice4[0], slice5[0]
}


func appendToOriginArray02() ([]int, []int, int) {
	slice7 := make([]int, 5)   // 定義一個長度為 5 ，容量為 10 的初始化切片 slice2
	cl := copy(slice6, slice7) // 使用內建函式 copy() 將 slice1 當作目標切片， slice2 的值複製進 slice1 ，因為內建函式 copy() 會使其用不同隱藏陣列

	slice7[0] = 77 // 更改 slice2 切片索引值 0 的值為 77
	return slice6, slice7, cl
}

// slice8 := []int{1, 2, 3, 4, 5}
// slice9 := append(slice1[:0:0], slice1...)
// <切片>[<起始索引值>:<結束索引值(不含)>:<容量>]

func main() {
	fmt.Println("\n-----<目標切片>, <來源切片>------")
	s1, s2, s3 := sameHideArray()
	fmt.Println("指向同一個隱藏陣列，且更改中一個的值：", "slice1[0]=", s1, "; slice2[0]=", s2, "; slice3[0]=", s3)

	// 因為三者皆指向同一個底層隱藏陣列，故更換其中一個的值，其餘切片也會受影響
	fmt.Println("\n-----copy()，使其指向不同隱藏陣列-----")
	s4, s5 := appendToOriginArray01()
	fmt.Println("指向不同隱藏陣列，且更改 slice2 中的值：", "slice4[0]=", s4, "; slice5[0]=", s5) // 因為兩者指向不同一個底層隱藏陣列，故更換其中一個的值，其餘切片不會受影響

	fmt.Println("\n-----copy()，使將長的切片複製到短的切片中-----")
	s6, s7, len := appendToOriginArray02()
	fmt.Println("指向不同隱藏陣列，且更改 slice2 中的值：", "slice6=", s6, "; slice7=", s7, "; 複製的元素量=", len)
	// 因為兩者指向不同一個底層隱藏陣列，故更換其中一個的值，其餘切片不會受影響，且因為 slice1 長度比 slice2 短，故只會從 slice2 複製 3 個元素到 slice1
}
