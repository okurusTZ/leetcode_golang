package dynamic_programing

import "leetcode_golang/library"

func maxProfit(prices []int) int {
	var dp = make([][]int, len(prices))
	// dp[i][0] -> 第i天持有股票 dp[i][1] -> 第i天不持有股票

	dp[0] = make([]int, 2)
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < len(prices); i++ {
		dp[i] = make([]int, 2)
		// 持有 = 继续持有 / 买入
		dp[i][0] = library.Max(dp[i-1][0], -prices[i])
		// 不持有 = 卖出
		dp[i][1] = library.Max(dp[i-1][0]+prices[i], dp[i-1][1])
	}

	return dp[len(prices)-1][1]
}
