package dynamic_programing

// 思路，对比两个字符串的时候
// abc
// bb
// 删除操作w1的操作可以看成dfs(i-1,j)
// 添加w1的操作，可以看作dfs(i, j-1) -> 插入一个b，相当于把w2的b去掉了
// 替换的操作，可以看作dfs(i-1,j-1)
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	var dfs func(int, int) int

	var cache = make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	dfs = func(i, j int) (res int) {
		if i < 0 {
			// 边界情况，如果w1为空，那么w2需要修改 j+1次 [0,j]
			return j + 1
		}
		if j < 0 {
			return i + 1
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}

		defer func() {
			cache[i][j] = res
		}()

		if word1[i] == word2[j] {
			return dfs(i-1, j-1)
		}

		return min(min(dfs(i-1, j-1), dfs(i, j-1)), dfs(i-1, j)) + 1
	}

	return dfs(m-1, n-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
