package main

func partition(s string) (res [][]string) {
	// 可以看作是一个回溯问题，是否在第j的位置增加逗号

	var n = len(s)
	res = make([][]string, 0)
	var dfs func(int)
	path := make([]string, 0)

	dfs = func(i int) {
		if i == n {
			res = append(res, append([]string(nil), path...))
			return
		}

		for j := i; j < n; j++ {
			tmp := s[i : j+1] // 注意边界问题
			if palindrome(tmp) {
				path = append(path, tmp)
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)
	return res
}

func palindrome(s string) bool {
	right := len(s) - 1
	left := 0

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	println(palindrome("aba"))
}
