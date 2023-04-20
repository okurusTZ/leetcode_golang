package dynamic_programing

import "leetcode_golang/library"

func maxProfitii(prices []int) int {
	var dp = make([][]int, len(prices))

	for i := range dp {
		dp[i] = make([]int, 2)
	}

	// 持有股票
	dp[0][0] = -prices[0]
	// 不持有股票
	dp[0][1] = 0

	for i := 1; i < len(prices); i++ {

		// 持有
		dp[i][0] += library.Max(dp[i-1][0], dp[i-1][1]-prices[i])
		// 不持有
		dp[i][1] += library.Max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(prices)-1][1]
}
