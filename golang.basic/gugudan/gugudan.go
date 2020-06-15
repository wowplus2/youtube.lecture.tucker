package main

import (
	"fmt"
)

func gugudan() {
	for dan := 2; dan < 10; dan++ {
		fmt.Printf(" [%d 단] \n", dan)

		for j := 1; j < 10; j++ {
			fmt.Printf("%d * %d = %d\n", dan, j, (dan * j))
		}

		fmt.Println()
	}
}

func starMark01() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i + 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func starMark02() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5 - i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func starMark03() {
	for i := 0; i < 3; i++ {
		for j := 0; j < i + 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2 - i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func starMark04() {
	// line수 * 2 + 1 <=== 홀수
	for i := 0; i < 4; i++ {
		for j := 0; j < 3 - i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i * 2 + 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func starMark05() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i * 2 + 1; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < i + 1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 5 - i*2; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
	
}

func main() {
	//gugudan()	
	//starMark01()
	//starMark02()
	//starMark03()
	//starMark04()
	starMark05()
}