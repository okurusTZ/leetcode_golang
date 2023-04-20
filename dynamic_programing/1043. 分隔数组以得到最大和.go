package dynamic_programing

import "leetcode_golang/library"

func maxSumAfterPartitioning(arr []int, k int) int {
	var dp func(int) int
	var memo = make([]int, len(arr))

	for i := range memo {
		memo[i] = -1
	}

	dp = func(i int) (res int) {
		if i < 0 {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}

		var (
			m = 0
		)
		// i-k = j最远可以遍历到的地方
		for j := i; j > i-k && j >= 0; j-- {
			// 获取遍历到的最大值
			m = library.Max(arr[j], m)
			//
			res = library.Max(res, dp(j-1)+(i-j+1)*m)
		}

		memo[i] = res

		return res
	}

	return dp(len(arr) - 1)
}
