package mathmatic

import (
	"sort"
)

func numMovesStones(a int, b int, c int) []int {
	// 先排序
	tmp := []int{a, b, c}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	a, b, c = tmp[0], tmp[1], tmp[2]
	// 最大，固定一个，其他两个向这个移动，每次移动一格
	// 1 5 9 -> 4 5 6 -> 234 876 -> 移动6次
	max := c - b + b - a - 2
	// 剩余情况都是2
	min := 2
	// 最小，分类讨论
	if c-a == 2 {
		min = 0
	} else {
		// 如果有两个已经连续
		if b-a == 1 || c-b == 1 {
			min = 1
		}
		if b-a == 2 || c-b == 2 {
			min = 1
		}
	}

	return []int{min, max}

}
