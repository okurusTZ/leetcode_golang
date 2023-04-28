package backtracking

import (
	"fmt"
)

var phoneMap = []string{
	"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
}

// 1. 递归
func letterCombinations(digits string) (res []string) {

	n := len(digits)
	if n == 0 {
		return []string{}
	}

	path := make([]byte, n)
	var dfs func(int)
	dfs = func(i int) {
		// 满了
		if i == n {
			res = append(res, string(path))
			return
		}
		for _, c := range phoneMap[digits[i]-'0'] {

			path[i] = byte(c)
			dfs(i + 1)
		}
	}

	dfs(0)

	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
}
