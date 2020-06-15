package main

import "fmt"

func usage01() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[4:8]
	c := a[4:]
	d := a[:4]

	fmt.Printf("a:%p, b:%p, c:%p, d:%p\n", a, b, c, d)

	fmt.Printf("a[4:8] = %v\n", b)
	fmt.Printf("a[4:] = %v\n", c)
	fmt.Printf("a[:4] = %v\n", d)
}

func usage02() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[4:8] // 같은 메모리 영역을 사용한다.
	b[0] = 1
	b[1] = 2

	fmt.Printf("a[] = %v\n", a)
}

func usage03() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 5; i++ {
		var back int
		a, back = removeStepBack(a)
		fmt.Printf("%d\n", back)
	}

	fmt.Printf("a[] = %v\n", a)
}

func usage04() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 5; i++ {
		var back int
		a, back = removeStepFront(a)
		fmt.Printf("%d\n", back)
	}

	fmt.Printf("a[] = %v\n", a)
}

func removeStepBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func removeStepFront(a []int) ([]int, int) {
	return a[1:], a[0]
}

func main() {
	//usage01()
	//usage02()
	//usage03()
	usage04()
}
