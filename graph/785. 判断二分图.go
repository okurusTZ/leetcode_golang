package graph

func isBipartite(graph [][]int) bool {
	// 是否是二分图，一个部分的内的图一定不邻接
	var (
		color   = make([]bool, len(graph)) // 表示每个节点是什么颜色的
		visited = make([]bool, len(graph))
		ok      = true
	)

	var traverse func(int, *[]bool, *[]bool, *bool)
	traverse = func(node int, color, visited *[]bool, ok *bool) {
		if !*ok {
			return
		}

		(*visited)[node] = true
		for _, next := range graph[node] {
			if !(*visited)[next] {
				(*color)[next] = !(*color)[node]
				traverse(next, color, visited, ok)
			} else {
				if (*color)[next] == (*color)[node] {
					*ok = false
				}
			}
		}

	}

	for i := range graph {
		if !visited[i] {
			traverse(i, &color, &visited, &ok)
		}
	}
	return ok
}
