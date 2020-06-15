package main

import "fmt"

func main() {
/*	var i int
	for {
		i++
		fmt.Println(i)
	}

	fmt.Printf("최종 i값: %v", i)
*/
/*
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
			//continue
		}
		fmt.Println(i)
	}

	fmt.Println("catch by break statement!")
*/
	var i int
	for {
		if i == 5 {
			i++
			continue
		} 
		if i == 6 {
			break
		}

		fmt.Println(i)
		i++
	}
}