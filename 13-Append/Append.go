package main

import "fmt"

var slice1 = []int{1, 2, 3, 4, 5} 
var slice4 = make([]int, 5, 10) 

func appendToNewArray() (int, int) {
	slice2 := slice1                 // 複製 slice1 切片為 slice2 (會使兩者指向同一個隱藏陣列)
	slice2 = append(slice2, 6, 7, 8) // 使用內建函式 append() 將 slice2 新增元素值 6.7.8 ，本來跟 slice1 指向相同隱藏陣列，但因為超過本來的容量會陣列置換，所以建立一個跟本來不同的隱藏陣列 slice2
	slice2[0] = 77 // 更改 slice2 切片索引值 0 的值為 77

	fmt.Println("\n ------new slice2------")
	for i := 0; i < len(slice2); i++ { // len() 可以取得任意集合的長度，索引值從 0 起累加，用 len() 的回傳值即可判定迴圈是否走到最後一個值，是否繼續執行或是終止。
		fmt.Print(slice2[i], " ")
	}
	fmt.Println("")
	fmt.Println("\n-----例1-----")
	return slice1[0], slice2[0]
}


func appendToOriginArray() (int, int) {
	slice5 := slice4                 // 複製 slice1 切片為 slice2 (會使兩者指向同一個隱藏陣列)
	slice5 = append(slice5, 6, 7, 8) // 使用內建函式 append() 將 slice2 新增元素值 6.7.8 因為沒有超過容量，所以會共用同一個隱藏陣列

	slice5[0] = 77 // 更改 slice2 切片索引值 0 的值為 77
	return slice4[0], slice5[0]
}

func main() {
	s1, s2 := appendToNewArray()
	fmt.Println("指向新的隱藏陣列，且更改 slice2 中的值：", "slice1[0]=", s1, "; slice2[0]=", s2) // 因為兩者已經指向不同底層隱藏陣列，故更換其中一個的值，原本切片不會受影響
	fmt.Println("\n-----例2-----")
	s4, s5 := appendToOriginArray()
	fmt.Println("指向同個隱藏陣列，且更改 slice5 中的值：", "slice4[0]=", s4, "; slice5[0]=", s5) // 因為兩者皆指向同一個底層隱藏陣列，故更換其中一個的值，其餘切片也會受影響
}
