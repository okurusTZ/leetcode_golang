package dynamic_programing

// 思路，枚举选哪个
func lengthOfLIS(nums []int) int {

	var dfs func(int) int
	var memo = make(map[int]int)
	// 以i结尾的最长单调子序列
	dfs = func(i int) (res int) {

		defer func() {
			memo[i] = res
		}()

		if memo[i] > 0 {
			return memo[i]
		}

		res = 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			res = max(res, dfs(j)+1)
		}
		return res
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = max(ans, dfs(i))
	}

	return ans
}
