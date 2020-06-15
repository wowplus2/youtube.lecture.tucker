package main

import "fmt"

// Student is student's info
type Student struct {
	name  string
	class int

	score Score
}

// Score is student's scord
type Score struct {
	name  string
	grade string
}

// ViewScore print student's score information
func (s Student) ViewScore() {
	fmt.Println(s.score)
}

// ViewScore 의 함수 표현
func ViewScore(s Student) {
	fmt.Println(s.score)
}

// InputScore 는 과목,성적을 입력 받는다.
func (s Student) InputScore(name string, grade string) {
	s.score.name = name
	s.score.grade = grade
}

// InputScore 의 함수 표현
func InputScore(s Student, name string, grade string) {
	s.score.name = name
	s.score.grade = grade
}

func main() {
	var s Student
	s.name = "Daniel"
	s.class = 1
	s.score.name = "수학"
	s.score.grade = "C"

	s.ViewScore()
	ViewScore(s)

	s.InputScore("과학", "A") // 입력값이 출력되지 않는다.
	s.ViewScore()
}
