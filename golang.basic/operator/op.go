package main

import "fmt"

func main() {
	a := 4
	b := 2
	c := 21
	d := c / 10
	e := c % 10
	f := 4

	fmt.Printf("%v\n", a&b);
	fmt.Printf("%v\n", a|b);
	fmt.Printf("%v의 첫번째 값: %v\n", c, d);
	fmt.Printf("%v의 두번째 값: %v\n", c, e);
	fmt.Printf("%v\n", f << 1);
	fmt.Printf("%v\n", f >> 1);
}