package leetcode_golang

import "sort"

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	// 排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := 1
	tmp := 1

	// fmt.Println(nums)

	for i := 1; i < len(nums); i++ {

		if nums[i]-nums[i-1] == 0 {
			continue
		}

		if nums[i]-nums[i-1] == 1 {
			tmp++
		} else {
			tmp = 1
		}
		ans = max(ans, tmp)
	}

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
