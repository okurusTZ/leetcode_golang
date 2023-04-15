package main

import "fmt"

func gardenNoAdj(n int, paths [][]int) []int {

	// 一个有向图
	graph := make([][]int, n)

	for _, path := range paths {
		// 变成0下标的
		x, y := path[0]-1, path[1]-1

		// 表示双向通道
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	// 记录的是第n个花园种植的类型
	color := make([]int, n)

	// 第i个花园，与哪些相通
	for i, nodes := range graph {
		// 一共四种颜色
		used := [5]bool{}

		for _, node := range nodes {
			// 记录下已经用过的颜色
			// 这里需要注意，要不初始化一下color为-1
			// 要不就设置used[0]无意义，不然的话，未种植的花园会产生误导信息
			used[color[node]] = true
		}

		// 花的种类用1234表示
		for j := 1; j < 5; j++ {
			if !used[j] {
				color[i] = j
				break
			}
		}
	}

	return color

}

func main() {
	fmt.Println(gardenNoAdj(4, [][]int{
		{1, 2},
		{3, 4},
	}))
}
