package dynamic_programing

import "leetcode_golang/library"

// 回溯+存储
func rob(nums []int) int {
	n := len(nums)
	// 记录一下，每个位置已经计算过的钱
	list := make([]int, n)

	for i := 0; i < n; i++ {
		list[i] = -1 // init
	}

	return dfs(n-1, nums, list)

}

func dfs(i int, nums, list []int) int {
	if i < 0 {
		return 0
	}
	if list[i] != -1 {
		return list[i]
	}
	res := library.Max(dfs(i-1, nums, list), dfs(i-2, nums, list)+nums[i])
	list[i] = res
	return res
}
