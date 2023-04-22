package dynamic_programing

func longestSubsequence(arr []int, difference int) int {
	var dp = make(map[int]int)

	// dp[0] = 0

	ans := 0
	for i := 0; i < len(arr); i++ {
		dp[arr[i]] = dp[arr[i]-difference] + 1
		ans = Max(ans, dp[arr[i]])
	}
	return ans
}
