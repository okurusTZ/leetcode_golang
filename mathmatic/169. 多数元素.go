package mathmatic

import "sort"

func majorityElement2(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[len(nums)/2]
}

func majorityElement(nums []int) int {
	var (
		candidate = 0
		count     = 0
	)
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if candidate != num {
			count--
		} else {
			count++
		}
	}
	return candidate
}
