package leetcode_golang

func distinctDifferenceArray(nums []int) []int {
	var (
		set    = make(map[int]int)
		suffix = make([]int, len(nums)+1)
		ans    = make([]int, len(nums))
	)

	for i := len(nums) - 1; i >= 0; i-- {
		set[nums[i]] = nums[i]
		suffix[i] = len(set)
	}
	suffix[len(nums)] = 0

	set = map[int]int{}

	for i, num := range nums {
		set[num] = num
		ans[i] = len(set) - suffix[i+1]
	}
	return ans
}
