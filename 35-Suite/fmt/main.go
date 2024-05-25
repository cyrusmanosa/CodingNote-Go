package main

import (
	"fmt"
	"os"
)

func init() {
	// 打開文件
	file, err := os.OpenFile("fmt.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("打開文件失敗:", err)
		return
	}
	defer file.Close()

	// 讀取文件內容
	data := make([]byte, 8192)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("讀取文件失敗:", err)
		return
	}
	fmt.Printf("文件內容：\n%s\n", string(data[:count]))
}

func main() {
	fmt.Println(" ")
	type Person struct {
		Name string
	}
	var people = Person{Name: "mark"}

	//1.普通占位符
	//%v(相應值的默認格式)
	fmt.Printf("%v", people) //{mark}
	fmt.Println(" ")
	//%+v(打印結構體時，會添加字段名)
	fmt.Printf("%+v", people) //{Name:mark}
	fmt.Println(" ")

	//%#v(相應值的Go語法表示)
	fmt.Printf("%#v", people) //main.Person{Name:"mark"}
	fmt.Println(" ")

	//%T(相應值的類型的Go語法表示)
	fmt.Printf("%T", people) //main.Person
	fmt.Println(" ")

	//%%(字面上的百分號，並非值的占位符)
	fmt.Printf("%%") //%
	fmt.Println(" ")

	//2.布爾占位符
	//%t(true 或 false)
	fmt.Printf("%t", true) //true
	fmt.Println(" ")

	//3.整數占位符
	//%b(二進制表示)
	fmt.Printf("%b", 5) //101
	fmt.Println(" ")

	//%c(相應Unicode碼點所表示的字符)
	fmt.Printf("%c", 0x4E2D) //中
	fmt.Println(" ")

	//%d(十進制表示)
	fmt.Printf("%d", 0x12) //18
	fmt.Println(" ")

	//%o(八進制表示)
	fmt.Printf("%o", 10) //12
	fmt.Println(" ")

	//%q(單引號圍繞的字符字面值，由Go語法安全地轉義)
	fmt.Printf("%q", 0x4E2D) //'中'
	fmt.Println(" ")

	//%x(十六進制表示，字母形式為小寫a-f)
	fmt.Printf("%x", 13) //d
	fmt.Println(" ")

	//%X(十六進制表示，字母形式為小寫A-F)
	fmt.Printf("%X", 13) //D
	fmt.Println(" ")

	//%U(Unicode格式：U+1234，等同於 "U+%04X")
	fmt.Printf("%U", 0x4E2D) //U+4E2D
	fmt.Println(" ")

	//4.浮點數和覆數的組成部分
	//%b(無小數部分的，指數為二的冪的科學計數法)
	fmt.Printf("%b", 10.2) //5742089524897382p-49
	fmt.Println(" ")

	//%e(科學計數法,例如 -1234.456e+78)
	fmt.Printf("%e", 10.2) //1.020000e+01
	fmt.Println(" ")

	//%E(科學計數法,例如 -1234.456E+78)
	fmt.Printf("%E", 10.2) //1.020000E+01
	fmt.Println(" ")

	//%f(有小數點而無指數，例如123.456)
	fmt.Printf("%f", 10.2) //10.200000
	fmt.Println(" ")

	//%g(根據情況選擇%e或%f以產生更緊湊的(無末尾的0))
	fmt.Printf("%g", 10.20) //10.2
	fmt.Println(" ")

	//%G(根據情況選擇%E或%f以產生更緊湊的(無末尾的0))
	fmt.Printf("%G", 10.20+2i) //(10.2+2i)
	fmt.Println(" ")

	//5.字符串與字節切片
	//%s(輸出字符串表示(string類型或[]byte))
	fmt.Printf("%s", []byte("Go語言")) //Go語言
	fmt.Println(" ")

	//%q(雙引號圍繞的字符串，由Go語法安全地轉義)
	fmt.Printf("%q", "Go語言") //"Go語言"
	fmt.Println(" ")

	//%x(十六進制，小寫字母，每字節兩個字符)
	fmt.Printf("%x", "golang") //676f6c616e67
	fmt.Println(" ")

	//%X(十六進制，大寫字母，每字節兩個字符)
	fmt.Printf("%X", "golang") //676F6C616E67
	fmt.Println(" ")

	//6.指針
	//%p(十六進制表示，前綴0x)
	fmt.Printf("%p", &people) //0xc0420421d0
	fmt.Println(" ")
}
