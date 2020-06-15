package main

import "fmt"

func usageSample01() {
	//var a []int
	//a := []int{1, 3, 5, 7, 9}
	a := make([]int, 0, 8)

	fmt.Printf("len(a): %d\n", len(a))
	fmt.Printf("cap(a): %d\n", cap(a))

	fmt.Println("------------------------")
	a = append(a, 1)
	fmt.Printf("len(a): %d\n", len(a))
	fmt.Printf("cap(a): %d\n", cap(a))
}

func usageSample02() {
	a := []int{1, 2}
	b := append(a, 3)

	b[0] = 4
	b[1] = 5

	fmt.Printf("a: %p, b: %p\n", a, b)

	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
	for i := 0; i < len(b); i++ {
		fmt.Printf("%d ", b[i])
	}
	fmt.Println()

	fmt.Printf("len(a): %d, cap(a): %d\n", len(a), cap(a))
	fmt.Printf("len(b): %d, cap(b): %d\n", len(b), cap(b))
}

func usageSample03() {
	a := make([]int, 2, 4)
	b := append(a, 3)

	a[0] = 1
	a[1] = 2

	b[0] = 4
	b[1] = 5

	fmt.Printf("a: %p, b: %p\n", a, b)

	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
	for i := 0; i < len(b); i++ {
		fmt.Printf("%d ", b[i])
	}
	fmt.Println()

	fmt.Printf("len(a): %d, cap(a): %d\n", len(a), cap(a))
	fmt.Printf("len(b): %d, cap(b): %d\n", len(b), cap(b))
}

func usageSample04() {
	a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2

	b := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}

	b = append(b, 3)
	b[0] = 4
	b[1] = 5

	fmt.Printf("a: %p, b: %p\n", a, b)

	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
	for i := 0; i < len(b); i++ {
		fmt.Printf("%d ", b[i])
	}
	fmt.Println()

	fmt.Printf("len(a): %d, cap(a): %d\n", len(a), cap(a))
	fmt.Printf("len(b): %d, cap(b): %d\n", len(b), cap(b))
}

func main() {
	//usageSample01()
	//usageSample02() // 생선된 slice는 서로 다른 메모리공간에 배치된다.
	//usageSample03()	// make에 의해 생성된 slice는 같은 메모리공간을 공유한다.
	usageSample04() // make에 의해 생성된 slice는 서로 다른 메모리공간에 배치된다.
}
