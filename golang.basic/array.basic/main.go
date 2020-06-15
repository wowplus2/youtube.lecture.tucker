package main

import (
	"fmt"
)

func step01() {
	var arr [10]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i * i
	}

	fmt.Println(arr)
}

func step02() {
	strEn := "Hello World"
	strKr := "Hello 월드"
	str2RuneKr := []rune(strKr)

	for i := 0; i < len(strEn); i++ {
		fmt.Print(strEn[i], ", ")
	}
	fmt.Println()
	for i := 0; i < len(strEn); i++ {
		fmt.Print(string(strEn[i]), ", ")
	}
	fmt.Println()

	for i := 0; i < len(strKr); i++ {
		fmt.Print(strKr[i], ", ")
	}
	fmt.Println()
	for i := 0; i < len(strKr); i++ {
		fmt.Print(string(strKr[i]), ", ")
	}
	fmt.Println()
	for i := 0; i < len(str2RuneKr); i++ {
		fmt.Print(str2RuneKr[i], ", ")
	}
	fmt.Println()
	for i := 0; i < len(str2RuneKr); i++ {
		fmt.Print(string(str2RuneKr[i]), ", ")
	}
}

func main() {
	//step01()
	step02()
}
