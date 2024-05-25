package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// pos 中的 sum 和 neg 中的 sum 是不同變數
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
