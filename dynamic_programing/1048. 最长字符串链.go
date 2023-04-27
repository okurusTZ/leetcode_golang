package dynamic_programing

func longestStrChain(words []string) int {
	// 利用一个map记载当前字符串s为结尾时，返回的词链的长度

	ws := make(map[string]int)

	for _, word := range words {
		ws[word] = 0 // init
	}

	ans := 0

	var dfs2 func(string) int
	dfs2 = func(s string) int {
		if len(s) == 0 {
			return 0
		}

		res := ws[s]

		if res != 0 {
			return res
		}

		for i := range s {
			// 拿走第i个字符
			j := s[:i] + s[i+1:]
			if _, ok := ws[j]; ok {
				res = max(dfs2(j), res)
			}
		}

		ws[s] = res + 1
		return res + 1
	}

	for s := range ws {
		ans = max(ans, dfs2(s))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
