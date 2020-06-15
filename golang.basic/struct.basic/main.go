package main

import "fmt"

// Person is personal information(name, age)
type Person struct {
	name string
	age  int
}

// PrintName method output user's name
func (p Person) PrintName() {
	fmt.Println(p.name)
}

func main() {
	var p Person
	p1 := Person{"홍길동", 15}
	p2 := Person{name: "이수일", age: 21}
	p3 := Person{name: "Daniel"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)

	p.name = "Smith"
	p.age = 24

	fmt.Println(p)

	p.PrintName()
}
