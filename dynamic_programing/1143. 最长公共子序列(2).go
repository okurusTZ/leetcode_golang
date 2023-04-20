package dynamic_programing

import "leetcode_golang/library"

func longestCommonSubsequence2(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = library.Max(f[i][j+1], f[i+1][j])
			}
		}
	}

	return f[n][m]
}
