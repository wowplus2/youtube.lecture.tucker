package main

import "fmt"

func main() {
	x := 33

/*	
	switch x {
		case 31:
			fmt.Println("x = 31")
		case 32:
			fmt.Println("x = 32")
		case 33:
			fmt.Println("x = 33")
	}
*/
	switch {
		case x > 40:
			fmt.Println("x는 40보다 크다.")
		case x < 20:
			fmt.Println("x는 20보다 작다.")
		case x > 30:
			fmt.Println("x는 30보다 크다.")
		case x == 33:
			fmt.Println("x는 33")
		
	}
}