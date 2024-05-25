package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	// V1()
	V2()
}

func V3() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines) // Read line by line

	// Read length
	fmt.Print("Enter the length of the array: ")
	scanner.Scan()
	length, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid length:", err)
		return // Exit if length is invalid
	}

	// Read numbers
	fmt.Print("Enter the numbers separated by spaces: ")
	scanner.Scan()
	numberStrs := strings.Fields(scanner.Text()) // Split on whitespace

	if len(numberStrs) != length {
		fmt.Fprintln(os.Stderr, "incorrect number of numbers entered")
		return // Exit if the number count doesn't match
	}

	var numbers []int
	for _, str := range numberStrs {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintln(os.Stderr, "invalid number:", err)
			return // Exit if any number is invalid
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return // Exit on scanning error
	}

	fmt.Println("The numbers are:", numbers)
}
