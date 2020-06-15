package main

import (
	ctrl "lec.youtube.tucker/quiz01/controllers"
)

func main() {
	//ctrl.DispMinHeapTree()

	/*
		Quiz01
		Input: [-1, 3, -1, 5, 4]
		Output: 4
	*/
	nums := []int{-1, 3, -1, 5, 4}
	needle := 2
	ctrl.DispQuizResult(nums, needle)

	/*
		Quiz02
		Input: [2, 4, -2, -3, 8]
		Output: 4
	*/
	nums = []int{2, 4, -2, -3, 8}
	needle = 1
	ctrl.DispQuizResult(nums, needle)

	/*
		Quiz03
		Input: [-5, -3, 1]
		Output: 3
	*/
	nums = []int{-5, -3, 1}
	needle = 3
	ctrl.DispQuizResult(nums, needle)

}
