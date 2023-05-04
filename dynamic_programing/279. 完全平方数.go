package dynamic_programing

import "math"

func numSquares(n int) int {

	var dp = make([]int, n)

	for i := 1; i <= n; i++ {
		min := math.MaxInt
		for j := 1; j*j <= i; j++ {
			if min > dp[i-j*j] {
				min = dp[i-j*j]
			}
		}
		dp[i] = min + 1
	}
	return dp[n-1]
}
