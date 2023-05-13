package leetcode_golang

import "sort"

func findMaxK(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var (
		i, j = 0, len(nums) - 1
	)

	for i <= j {
		for nums[i] < 0 && nums[i]*-1 > nums[j] {
			i++
		}
		if nums[i] >= 0 {
			return -1
		}

		if -nums[i] == nums[j] {
			return -nums[i]
		}

		for nums[j] > 0 && nums[j]*-1 > nums[i] {
			j--
		}

		if nums[j] <= 0 {
			return -1
		}

		if -nums[i] == nums[j] {
			return -nums[i]
		}
	}
	return -1
}
