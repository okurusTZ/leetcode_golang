package main

import "math"

// 完全背包问题 -> 可以取多次

func coinChange(coins []int, amount int) int {
	n := len(coins)

	// 缓存
	cache := make([][]int, n)

	for i := range cache {
		cache[i] = make([]int, amount+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(i, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 0
			}
			// 说明凑不出来
			return math.MaxInt64
		}

		if cache[i][c] != -1 {
			return cache[i][c]
		}

		defer func() {
			cache[i][c] = res
		}()

		if coins[i] > c {
			return dfs(i-1, c)
		}
		// 一种是不选；一种是选了 -> 因为选了还可以选，所以仍然是dfs(i,...)
		return min(dfs(i-1, c), dfs(i, c-coins[i])+1)
	}

	res := dfs(n-1, amount)
	// println(res)
	if res < math.MaxInt64 && res >= 0 {
		return res
	}
	return -1
}

// 这里防止溢出，可以改成如果<0,就自动退出比较
func min(a, b int) int {
	if a < 0 && b < 0 {
		return 0
	}
	if a < 0 {
		return b
	}
	if b < 0 {
		return a
	}
	if a < b {
		return a
	}
	return b
}

func main() {
	println(coinChange([]int{2, 5, 10, 1}, 27))
}
