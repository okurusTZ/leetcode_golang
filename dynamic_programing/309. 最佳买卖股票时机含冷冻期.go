package dynamic_programing

func maxProfitiii(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	var dp = make([][]int, len(prices))

	for i := range dp {
		dp[i] = make([]int, 2)
	}

	// 持有股票
	dp[0][0] = -prices[0]
	// 不持有股票
	dp[0][1] = 0

	// 持有
	dp[1][0] = max(dp[0][0], dp[0][1]-prices[1])
	// 不持有
	dp[1][1] = max(dp[0][1], dp[0][0]+prices[1])

	for i := 2; i < len(prices); i++ {

		// 持有
		dp[i][0] += max(dp[i-1][0], dp[i-2][1]-prices[i])
		// 不持有 -> 为什么是i-2，因为i-1不能有卖出操作
		dp[i][1] += max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(prices)-1][1]
}
