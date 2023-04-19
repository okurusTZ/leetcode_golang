package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
		}
		if sum < target {
			left++
		}
		if sum == target {
			return []int{left + 1, right + 1}
		}
	}
	return []int{left + 1, right + 1}
}

func main() {
	fmt.Println(twoSum([]int{-1, 0}, -1))
}
