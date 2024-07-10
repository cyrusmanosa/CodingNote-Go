package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// V1()
	// V2()
	V3()
	// V4()
}
func V1() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
		} else {
			break
		}
	}
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}

func V2() {
	var val int
	// 数値の入力
	fmt.Println("数値を入力してください")
	fmt.Scan(&val)

	fmt.Println("入力結果")
	fmt.Printf("%d\n", val)

	// 文字列の入力
	fmt.Println("文字列を入力してください")
	var str string
	fmt.Scan(&str)

	fmt.Println("入力結果")
	fmt.Println(str)
}

// 首先輸入Array的長度，再輸入內容
func V3() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	fmt.Print("Enter the length of the array: ")
	scanner.Scan()

	length, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid length:", err)
		return
	}

	fmt.Print("Enter the numbers separated by spaces: ")
	scanner.Scan()
	numberStrs := strings.Fields(scanner.Text())

	if len(numberStrs) != length {
		fmt.Fprintln(os.Stderr, "incorrect number of numbers entered")
		return
	}

	var numbers []int
	for _, str := range numberStrs {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintln(os.Stderr, "invalid number:", err)
			return
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return
	}

	fmt.Println("The numbers are:", numbers)
}

// 首先輸入struct的數目，再輸入內容
type Data struct {
	Name     string
	Age      int
	Birthday string
	City     string
}

func V4() {
	scanner := bufio.NewScanner(os.Stdin)
	// 數目
	if !scanner.Scan() {
		fmt.Println("Failed to read number of data groups")
		return
	}
	numGroups, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
		return
	}

	dataGroups := make([]Data, 0, numGroups)
	// 讀取接下來的每組數據
	for i := 0; i < numGroups; i++ {
		if !scanner.Scan() {
			fmt.Println("Failed to read data group")
			return
		}
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 4 {
			fmt.Println("Invalid data group format:", line)
			return
		}

		age, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Invalid age:", parts[1])
			return
		}

		dataGroup := Data{
			Name:     parts[0],
			Age:      age,
			Birthday: parts[2],
			City:     parts[3],
		}
		dataGroups = append(dataGroups, dataGroup)
	}

	// 打印所有數據組
	for _, dataGroup := range dataGroups {
		fmt.Printf("Name: %s, Age: %d, Birthday: %s, City: %s\n",
			dataGroup.Name, dataGroup.Age, dataGroup.Birthday, dataGroup.City)
	}
}
