package main

// 思路，对比两个字符串的时候
// abc
// bb
// 删除操作w1的操作可以看成dfs(i-1,j)
// 添加w1的操作，可以看作dfs(i, j-1) -> 插入一个b，相当于把w2的b去掉了
// 替换的操作，可以看作dfs(i-1,j-1)
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	var f = make([][]int, m+1)

	// 需要初始化
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			// 这里因为都+1了，所以i=0的实际意义就是w1的长度是0，这时候w2到j的长度，所以应该修改j个字符
			if i == 0 {
				f[i][j] = j
			}
			if j == 0 {
				f[i][j] = i
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i+1][j], min(f[i][j], f[i][j+1])) + 1
			}
		}
	}
	return f[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
