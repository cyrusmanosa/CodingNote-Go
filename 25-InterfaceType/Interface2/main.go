package main

import (
	"fmt"
)

// step 1
type Car struct {
	name     string
	age      int
	odometer int
}
type Robot struct {
	name  string
	age   int
	power int
}

func main() {
	// 設定 R 為 Robot 型的值
	r := &Robot{
		name:  "GUMDAM",
		age:   1,
		power: 100,
	}

	// 設定 C 為 Car 型的值
	c := &Car{
		name:     "BENZ",
		age:      2,
		odometer: 100,
	}

	// 送到中轉的func
	Cleaning(c)
	Cleaning(r)
}

// step 2 設定 return 的func
func (r *Car) GetName() string {
	return r.name
}
func (r *Car) GetAge() int {
	return r.age
}

func (r *Robot) GetName() string {
	return r.name
}
func (r *Robot) GetAge() int {
	return r.age
}

// step 3 設定 interface 只要有同名的func 就會包括
type Member interface {
	GetName() string
	GetAge() int
}

func Cleaning(m Member) {
	fmt.Printf("Consumer Name:%s, Age:%d\n", m.GetName(), m.GetAge())
}
