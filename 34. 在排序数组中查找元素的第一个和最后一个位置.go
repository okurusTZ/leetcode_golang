package main

import "fmt"

func searchRange(nums []int, target int) []int {
	start := lowerBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := lowerBound(nums, target+1) - 1
	return []int{start, end}
}

func lowerBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1 // [left, right]

	// 只有一个的情况
	// [2], target=1
	for left <= right {
		mid := left + (right-left)/2
		// println(mid, left, right, nums[mid], nums[left], nums[right])
		// 注意这里等于号的取值
		// 因为是lowerBound，所以=的时候，应该收缩右区间
		if nums[mid] >= target {
			right = mid - 1 // [left, mid-1]
		} else {
			left = mid + 1 // [mid+1, right]
		}
	}

	// print(left)
	return left
}

func main() {

	// 0 5 -> mid = 2
	fmt.Print(searchRange([]int{2}, 1))
}
