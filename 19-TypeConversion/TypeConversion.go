package main

import (
	"errors"
	"fmt"
)

// 若想要將 int64 變數轉成 int8 ，注意有沒有可能發生溢位的錯誤 int64 > int8
// 同java一樣

func change() string {
	var maxInt8 int8 = 127
	greaterThanMaxInt8 := 128
	float := 9.999
	newInt := fmt.Sprintf("maxInt8: %v ,int64: %v\n", maxInt8, int64(maxInt8))                                // 將 int8 轉換成型別 int64，因為int64 型別的容量較大，所以可以正常轉換
	newInt += fmt.Sprintf("greaterThanMaxInt8: %v ,int8: %v\n", greaterThanMaxInt8, int8(greaterThanMaxInt8)) // 因為 greaterThanMaxInt8 大於 int8 的最大值，所以會發生越界繞回的問題
	newInt += fmt.Sprintf("maxInt8: %v ,float64: %v\n", maxInt8, float64(maxInt8))
	newInt += fmt.Sprintf("float: %v ,int: %v\n", float, int(float)) // 將 float 轉換為 int 小數點會直接捨去
	return newInt
}

// 接收一個 interface{} 型別的函式
func doubler1(value interface{}) (string, error) {
	if int, ok := value.(int); ok {
		return fmt.Sprint(int * 2), nil
	}
	if str, ok := value.(string); ok {
		return str + str, nil
	}
	// 型別不符前面的檢查, 傳回錯誤
	return "", errors.New("輸入的值錯誤")
}

func doubler2(value interface{}) (string, error) {
	switch t := value.(type) {
	case string:
		return t + t, nil
	case bool:
		if t {
			return "truetrue", nil
		} else {
			return "falsefalse", nil
		}
	case float32, float64:
		if f, ok := t.(float64); ok {
			return fmt.Sprint(f * 2), nil
		}
		return fmt.Sprint(t.(float32) * 2), nil
	case int:
		return fmt.Sprint(t * 2), nil
	default:
		return "", errors.New("輸入的值錯誤")
	}
}

func main() {
	fmt.Println(change())
	fmt.Println("\n--------型別斷言--------")
	res, _ := doubler1(5)
	fmt.Println("5 :", res)
	res, _ = doubler1("good")
	fmt.Println("good :", res)
	_, err := doubler1(true)
	fmt.Println("true :", err)

	fmt.Println("\n--------型別 switch--------")
	res2, _ := doubler2(-5)
	fmt.Println("-5 :", res2)
	res2, _ = doubler2(5)
	fmt.Println("5 :", res2)
	res2, _ = doubler2("good")
	fmt.Println("good :", res2)
	res2, _ = doubler2(true)
	fmt.Println("true :", res2)
	res2, _ = doubler2(float32(9.99))
	fmt.Println("9.99 :", res2)
}
