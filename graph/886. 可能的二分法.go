package graph

func possibleBipartition(n int, dislikes [][]int) bool {

	var (
		visited = make([]bool, n)
		color   = make([]bool, n)
		ok      = true
	)

	graph := buildGraph(n, dislikes)

	var dfs func(int, *[]bool, *[]bool, *bool)

	dfs = func(node int, color, visited *[]bool, ok *bool) {

		if !*ok {
			return
		}

		(*visited)[node] = true

		for _, next := range graph[node] {

			if !(*visited)[next] {
				(*color)[next] = !(*color)[node]
				dfs(next, color, visited, ok)
			} else {
				if (*color)[next] == (*color)[node] {
					// 无法组成二分
					*ok = false
					return
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i, &color, &visited, &ok)
		}
	}

	return ok
}
