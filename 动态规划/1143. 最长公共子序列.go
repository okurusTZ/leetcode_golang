package main

func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	var dfs func(int, int) int

	caches := make([][]int, n)
	for i := 0; i < n; i++ {
		caches[i] = make([]int, m)
		for j := 0; j < m; j++ {
			caches[i][j] = -1
		}
	}

	// 这里代表的是，text1[:i]和text2[:j]之间的最长公共子序列
	dfs = func(i, j int) (res int) {
		if i < 0 || j < 0 {
			return 0
		}

		if caches[i][j] != -1 {
			return caches[i][j]
		}

		defer func() {
			caches[i][j] = res
		}()

		if text1[i] == text2[j] {
			return dfs(i-1, j-1) + 1
		}
		return max(dfs(i-1, j), dfs(i, j-1))
	}

	return dfs(n-1, m-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var t1 = "pmjghexybyrgzczy"
	var t2 = "hafcdqbgncrcbihkd"

	println(longestCommonSubsequence(t1, t2))
}
