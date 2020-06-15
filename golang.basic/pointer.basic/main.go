package main

import "fmt"

func main() {
	var a int
	var b int
	var p *int

	p = &a
	a = 3
	b = 2

	fmt.Println(a)
	fmt.Println(p)  // p의 메모리 주소
	fmt.Println(*p) // p의 메모리 주소가 가리키는 주소안에 있는 값
	fmt.Println()

	p = &b
	fmt.Println(b)
	fmt.Println(p)
	fmt.Println(*p)
}
