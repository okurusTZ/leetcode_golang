package dynamic_programing

func wordBreak(s string, wordDict []string) bool {
	var (
		inDict = make(map[string]bool)
		dp     = make([]bool, len(s)+1) // 以n个字符结尾的子字符串是否能够用字典拼起来
	)

	for _, word := range wordDict {
		inDict[word] = true
	}

	dp[0] = true // init, 代表没有字符，所以一定是true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && inDict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]

}
