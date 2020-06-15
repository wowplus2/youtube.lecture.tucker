package main

import "fmt"

func aryCopy() {
	arr := [5]int{1, 2, 3, 4, 5}
	clone := [5]int{}

	for i := 0; i < len(arr); i++ {
		clone[i] = arr[i]
	}

	fmt.Println("arr:\t", arr)
	fmt.Println("clone:\t", clone)
}

func aryReverse() {
	arr := [5]int{1, 2, 3, 4, 5}
	temp := [5]int{}
	arrLen := len(arr) - 1

	for i := 0; i <= arrLen; i++ {
		temp[i] = arr[arrLen-i]
	}

	for i := 0; i <= arrLen; i++ {
		arr[i] = temp[i]
	}

	fmt.Println("temp:\t", temp)
	fmt.Println("arr:\t", arr)
}

func aryReverseNew() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	arrLen := len(arr) - 1

	fmt.Println("before reverse arr:\t", arr)
	for i := 0; i <= arrLen/2; i++ {
		arr[i], arr[arrLen-i] = arr[arrLen-i], arr[i]
	}
	fmt.Println("after reverse arr:\t", arr)
}

// 정렬(with Radix알고리즘)
func arySortWithRadix() {
	arr := [11]int{0, 5, 4, 9, 1, 2, 8, 3, 6, 4, 5}
	temp := [10]int{}

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}

	fmt.Println("arr:\t", arr)
	//fmt.Println("temp:\t", temp)
}

func main() {
	//aryCopy()
	//aryReverse()
	//aryReverseNew()
	arySortWithRadix()
}
