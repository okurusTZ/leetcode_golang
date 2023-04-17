package main

import "fmt"

func findColumnWidth(grid [][]int) []int {
	res := make([]int, len(grid[0]))
	for i := range grid[0] { // 遍历长度
		for _, jj := range grid {
			res[i] = max(res[i], len(fmt.Sprintf("%d", jj[i])))
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	findColumnWidth([][]int{
		{1, 22, 333},
	})
}
