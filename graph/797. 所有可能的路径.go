package graph

// 这题这么做的大前提是没有环

func allPathsSourceTarget(graph [][]int) [][]int {
	// 是否访问过该节点
	var visit = []int{}
	var ans = [][]int{}
	traverse(graph, &visit, &ans, 0)
	return ans
}

// 这里要引用传递
func traverse(graph [][]int, current *[]int, res *[][]int, node int) {

	*current = append(*current, node)

	if node == len(graph)-1 {
		// 遍历到最后一个节点
		// 因为这里是引用，需要copy一份当前的值，不然会全部是最后一个遍历的值
		*res = append(*res, append([]int(nil), *current...))
		return
	}

	// 如果没有，继续遍历相邻节点
	for _, v := range graph[node] {
		traverse(graph, current, res, v)
	}

	// 回溯
	*current = (*current)[:len(*current)-1]

}
