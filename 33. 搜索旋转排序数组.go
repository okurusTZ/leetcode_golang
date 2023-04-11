package main

import "fmt"

func search(nums []int, target int) int {
	minIdx := findMin(nums)
	println(minIdx)

	// 找到最小数的idx
	if target > nums[len(nums)-1] {
		// 在0 - i之间
		return lowerBound(nums[:minIdx+1], target)
	} else {
		// 在minIdx和len(nums)-1之间
		res := lowerBound(nums[minIdx:len(nums)], target)
		if res == -1 {
			return -1
		} else {
			return res + minIdx
		}
	}
}

// 寻找第一个
func lowerBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	fmt.Println(nums)

	for left <= right {
		mid := left + (right-left)/2
		// println(mid, left, right, nums[mid], nums[left], nums[right])
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// 缩小范围，先找出是在左半段还是右半段

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		// println(left, right, mid, nums[left], nums[right], nums[mid])
		// 注意这里应该是小于最后一个数
		if nums[mid] > nums[len(nums)-1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	println(search([]int{3, 1}, 0))
}
