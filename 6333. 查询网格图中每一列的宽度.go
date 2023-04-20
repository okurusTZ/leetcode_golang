package leetcode_golang

import (
	"fmt"
	"leetcode_golang/library"
)

func findColumnWidth(grid [][]int) []int {
	res := make([]int, len(grid[0]))
	for i := range grid[0] { // 遍历长度
		for _, jj := range grid {
			res[i] = library.Max(res[i], len(fmt.Sprintf("%d", jj[i])))
		}
	}
	return res
}
