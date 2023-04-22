package dynamic_programing

func longestArithSeqLength(nums []int) int {
	// 双指针的话是n^2

	ans := 0

	var dp = make([]map[int]int, len(nums)) // dp[i][d] -> 以i结尾，以d为等差数的时候的值

	dp[0] = map[int]int{}

	for i := 1; i < len(nums); i++ {
		dp[i] = map[int]int{}
		for j := i - 1; j >= 0; j-- {
			d := nums[i] - nums[j]
			// 没有初始化过 -> 如果有，那么需要+1
			if dp[i][d] == 0 {
				dp[i][d] = dp[j][d] + 1
				ans = Max(dp[i][d], ans)
				// fmt.Println(dp[i][d], i, d)
			}
		}
	}
	return ans + 1
}
