package main

import "fmt"

func main() {
	/*
		var m map[string]string     // 선언
		m = make(map[string]string) // 초기화
		m["abc"] = "aaa"
		fmt.Println("m['abc']: ", m["abc"])
		fmt.Println()

		m1 := make(map[int]string) // 선언&초기화
		m1[255] = "ddd"
		fmt.Println("m1[255]: ", m1[255])
		fmt.Println("m1[55]: ", m1[55]) // 해당map에 값이 없으면, string의 기본값 ''으로 대체된다.
		fmt.Println()

		m2 := make(map[int]int)
		m2[5] = 0
		m2[4] = 4
		v1, ok1 := m2[5]
		v2, ok2 := m2[10]

		fmt.Println("m2[5]: ", v1, " - is defined? ", ok1) // 해당 map의 설정여부
		fmt.Println("m2[4]: ", m2[4])
		fmt.Println("m2[10]: ", v2, " - is defined? ", ok2) // 해당map에 값이 없으면, int의 기본값 0으로 대체된다.
		fmt.Println()
	*/
	fmt.Println("1. m3 선언 및 key의 값 setting")
	m3 := make(map[int]bool)
	m3[4] = true
	v1, ok1 := m3[4]
	v2, ok2 := m3[10]
	fmt.Println("\tm3[4]: ", v1, " - is defined? ", ok1)
	fmt.Println("\tm3[10]: ", v2, " - is defined? ", ok2) // 해당map에 값이 없으면, bool의 기본값 false로 대체된다.

	fmt.Println("2. m3 순회")
	for k, v := range m3 {
		fmt.Println("\t", k, " = ", v)
	}

	delete(m3, 4)
	v1, ok1 = m3[4]
	v2, ok2 = m3[10]
	fmt.Println("3. delete m3[4]")

	fmt.Println("\tm3[4]: ", v1, " - is defined? ", ok1)
	fmt.Println("\tm3[10]: ", v2, " - is defined? ", ok2)
}
