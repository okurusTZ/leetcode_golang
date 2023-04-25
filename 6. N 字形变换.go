package leetcode_golang

import "strings"

// P   A   H   N
// A P L S I I G
// Y   I   R

// P     I    N
// A   L S  I G
// Y A   H R
// P     I

func convert(s string, numRows int) string {
	// 以一个V字为一个周期
	// 假如有n行，那么周期为 2 * (n-1)

	// 所以第i个字符的相对位置 i % (2*(n-1))
	// 如果 i % (2*(n-1)) > n, 已经走回去了，实际上是 n-(i % (2*(n-1)))

	ans := make([]string, len(s))

	n := 2 * (numRows - 1)
	for i, char := range s {
		k := min(i%n, n-(i%n))
		ans[k] = string(char)
	}

	return strings.Join(ans, "")
}
