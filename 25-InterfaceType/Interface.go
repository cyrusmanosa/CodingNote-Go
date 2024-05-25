package main

import (
	"fmt"
	"math"
)

type I interface{ M() }

type T struct{ S string }

func (t *T) M() { fmt.Println(t.S) }

type F float64

func (f F) M() { fmt.Println(f) }

//--------------------------

type Salaried interface{ getSalary() int }

// Salary 實作了 getSalary() 的方法，因此可以算是 Salaried type（polymorphism）
type Salary struct{ basic, insurance, allowance int }

func (s Salary) getSalary() int { return s.basic + s.insurance + s.allowance }

type Employee struct {
	firstName, lastName string
	salary              Salaried
}

func main() {
	var i I

	i = &T{"Hello"}                // 把 type T 的值賦予給變數 i
	fmt.Printf("(%v, %T)\n", i, i) // i 的 dynamic value 是 &{Hello}、 dynamic type 是 *main.T
	i.M()                          // 意思是將 type T 對應的 value （&{Hello}） 來執行 type T 對應的 Ｍ 方法

	i = F(math.Pi)                 // 把 type F 的值賦予給變數 i
	fmt.Printf("(%v, %T)\n", i, i) // i 的 dynamic value 是 3.141、dynamic type 是 main.F
	i.M()                          // 意思是將 type F 對應的 value （3.1415） 去執行 type F 對應的 Ｍ 方法

	fmt.Println("\n-------Interface 的 polymorphism(多型)--------")
	ross := Employee{
		firstName: "Ross",
		lastName:  "Geller",
		salary:    Salary{1100, 50, 50},
	}
	fmt.Println("Ross's salary is", ross.salary.getSalary())
}
