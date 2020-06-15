package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	strikes int
	balls   int
}

// MakeNumbers : 0 ~ 9사이의 겹치치 않는 무작위 숫자 3개를 반환한다.
func MakeNumbers() [3]int {
	var res [3]int

	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			isDuplicated := false
			for j := 0; j < i; j++ {
				if res[j] == n {
					// duplicated
					isDuplicated = true
					break
				}
			}

			if !isDuplicated {
				res[i] = n
				break
			}
		}
	}

	return res
}

// InpNumbers : 키보드로부터 0 ~9 사이의 겹치치않는 숫자를 입력받아 반환한다.
func InpNumbers() [3]int {
	var res [3]int
	len := len(res) - 1

	for {
		fmt.Print("중복되지 않는 0 ~ 9사이의 숫자를 입력하세요: ")
		var no int
		_, err := fmt.Scanf("%d\n", &no)
		if err != nil {
			fmt.Println("\t[e] Wrong input. Try again.")
			continue
		}

		isSuccess := true
		idx := 0
		for no > 0 {
			n := no % 10
			no = no / 10

			isDuplicated := false
			for j := 0; j < idx; j++ {
				if res[j] == n {
					// duplicated
					isDuplicated = true
					break
				}
			}

			if isDuplicated {
				fmt.Println("\t[i] 숫자가 중복되지 않아야합니다.")
				isSuccess = false
				break
			}

			if idx >= 3 {
				fmt.Println("\t[i] 3개의 숫자를 입력해야 하세요.")
				isSuccess = false
				break
			}

			res[idx] = n
			idx++
		}

		if isSuccess && idx < 3 {
			fmt.Println("\t[i] 3개의 숫자를 입력해야 하세요.")
			isSuccess = false
		}

		if !isSuccess {
			continue
		}
		break
	}

	// 배열 elements reverse
	var temp [3]int
	for i := 0; i <= len; i++ {
		temp[i] = res[len-i]
	}

	for i := 0; i <= len; i++ {
		res[i] = temp[i]
	}

	return res
}

// CompareNumbers : 2개의 숫자그룹을 입력받아 비교결과를 출력한다.
func CompareNumbers(nums, inums [3]int) Result {
	strikes := 0
	balls := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == inums[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
				break
			}
		}
	}

	return Result{strikes, balls}
}

// PrintResult : 결과출력
func PrintResult(result Result) {
	fmt.Printf("결과: %dS %dB\n", result.strikes, result.balls)
}

// IsGameEnd : result를 전달받아 완료조건을 출력한다.
func IsGameEnd(result Result) bool {
	res := false
	if result.strikes == 3 {
		res = true
	}

	return res
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// 무작위(0 ~ 9) 숫자 3개를 만든다.
	setNums := MakeNumbers()
	fmt.Println(setNums)

	cnt := 0
	for {
		cnt++
		// 사용자의 입력을 받는다.
		inpNums := InpNumbers()

		// 결과를 비교한다.
		result := CompareNumbers(setNums, inpNums)

		PrintResult(result)

		if IsGameEnd(result) {
			// game end
			break
		}
	}

	// 게임끝. 결과 print
	fmt.Printf("%d번만에 클리어!\n", cnt)
}
