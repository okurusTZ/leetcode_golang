package leetcode_golang

import "math"

func maxSubArray(nums []int) int {
	var (
		dp  = make([]int, len(nums))
		ans = math.MinInt16
	)

	dp[0] = nums[0]
	ans = max(ans, dp[0])

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}

	return ans
}
