package main

import (
	"fmt"
)

/*
   ***** 포인터를 함수 인자로 받으면 메모리 주소만 복사되고, 값을 함수 인자로 받으면 전체 속성이 복사된다.
   Pointer 호출: 메모리 주소만 복사 된다.
   값 호출: 전체 값이 복사 된다.
*/
type Student struct {
	name string
	age  int

	class string
	grade string
}

func (s *Student) outputScore() {
	fmt.Println(s.class, ":", s.grade)
}

func (s *Student) inputScore(class string, grade string) {
	s.class = class
	s.grade = grade
}

func incNumber(x *int) {
	// x++
	*x = *x + 1
}

func main() {
	var a int
	a = 1

	incNumber(&a)
	//fmt.Println(a)

	var s Student = Student{name: "Daniel", age: 23, class: "math", grade: "A+"}

	s.inputScore("Science", "C")
	s.outputScore()
}
