package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func f1(x int) {
	if x == 0 {
		return
	}
	fmt.Printf("f1(%d) before call f1(%d)\n", x, x-1)
	f1(x - 1)
	fmt.Printf("f1(%d) after call f1(%d)\n", x, x-1)
}

func sum(x int, s int) int {
	if x == 0 {
		return s
	}
	s += x
	return sum(x-1, s)
}

func sum2(x int, s int) int {
	for i := 0; i <= x; i++ {
		s += i
	}
	return s
}

func main() {
	/*
		for i := 0; i < 10; i++ {
			fmt.Printf("%d + %d = %d\n", i, i + 2, add(i, i + 2))
		}
	*/
	/*
		f1(10);
	*/
	/*
		s := sum(10, 0)
		fmt.Println(s)
	*/
	s := sum2(10, 0)
	fmt.Println(s)
}
