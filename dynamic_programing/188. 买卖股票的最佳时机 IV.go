package dynamic_programing

import (
	"leetcode_golang/library"
	"math"
)

func maxProfitiv(k int, prices []int) int {
	var dp = make([][][]int, len(prices)+1)

	for i := range dp {
		dp[i] = make([][]int, k+2) // 当前是第几次交易 k=0 -> -1次交易
		for j := range dp[i] {
			dp[i][j] = make([]int, 2) // 买入或卖出

			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -1 * math.MaxInt64
			}

			if j == 0 {
				dp[i][j][0] = -1 * math.MaxInt64
				dp[i][j][1] = -1 * math.MaxInt64
			}
		}
	}

	// 为什么初始化为 k+2 ？
	// 对于k=2来说，其实有三种状态 k=0 还没完成第一次交易
	// k=1 已经完成第一次交易
	// k=2 已经完成第二次交易

	// 初始化
	for i := 0; i < len(prices); i++ {
		for j := 1; j < k+2; j++ {
			// 未持有但已经完成第j次交易 = 一直未持有 OR i-1的时候持有，i的时候卖了
			dp[i+1][j][0] = library.Max(dp[i][j][0], dp[i][j-1][1]+prices[i])
			// 持有 = 一直持有 OR i-1的时候未持有
			dp[i+1][j][1] = library.Max(dp[i][j][1], dp[i][j][0]-prices[i])
		}
	}
	return dp[len(prices)][k+1][0]
}
