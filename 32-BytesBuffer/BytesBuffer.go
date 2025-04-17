package main

import (
	"bytes"
	"fmt"
	"unsafe"
)

// bytes.Buffer
// 存取 s1 第二次的結果，會發現後者 s2 資料蓋掉部分 s1 資料。
// 原因是這樣，當執行第二次 parseMultipleValue 後，bytes.Rest() 只將 offset 位置移動到 0 位置，
// 並將新的內容給寫入到同樣記憶體位置前面區段（s1頭 3 個字元）被改成新的 s2 字串。

var buf bytes.Buffer

func parseMultipleValue1(n int, str string) []byte {
	buf.Reset()
	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}
	return buf.Bytes() //解決方案： return buf.String()
}

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func parseMultipleValue2(n int, str string) string {
	buf.Reset()
	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}
	s := make([]byte, len(buf.Bytes())) // make ( Type : []byte , 長度 ： buf.Bytes )
	copy(s, buf.Bytes())                // 將 buf.Bytes 既 Data，Copy到 S
	return b2s(s)
}

func main() {
	fmt.Println("---------bytes.Buffer---------")
	s1 := parseMultipleValue1(5, "1")
	fmt.Println("s1:", string(s1))

	s2 := parseMultipleValue1(3, "2")
	fmt.Println("s1:", string(s1))
	fmt.Println("s2:", string(s2))

	fmt.Println("\n---------解決方案B-------")

	s3 := parseMultipleValue2(5, "1")
	fmt.Println("s1:", string(s3))
	s4 := parseMultipleValue2(3, "2")
	fmt.Println("s1:", string(s3))
	fmt.Println("s2:", string(s4))
}
