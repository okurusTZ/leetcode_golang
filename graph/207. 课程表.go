package graph

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 什么时候无法修完？ 存在依赖循环的时候
	graph := buildGraph(numCourses, prerequisites)

	var (
		onPath   = make([]bool, numCourses)
		visit    = make([]bool, numCourses)
		hasCycle = false
	)

	var traverse func([][]int, int, *[]bool, *[]bool, *bool)

	traverse = func(graph [][]int, node int, onPath *[]bool, visit *[]bool, hasCycle *bool) {
		if (*onPath)[node] {
			*hasCycle = true
		}
		if (*visit)[node] || *hasCycle {
			return
		}

		(*onPath)[node] = true
		(*visit)[node] = true
		for _, next := range graph[node] {
			traverse(graph, next, onPath, visit, hasCycle)
		}
		(*onPath)[node] = false
	}

	for i := 0; i < numCourses; i++ {
		traverse(graph, i, &onPath, &visit, &hasCycle)
	} // 应该遍历所有节点，而不是只有0开始

	return !hasCycle

}

// 先建立一个有向图
func buildGraph(nodeNum int, graph [][]int) [][]int {
	ans := make([][]int, nodeNum)

	for _, g := range graph {
		if ans[g[0]] == nil {
			ans[g[0]] = make([]int, 0)
		}
		ans[g[0]] = append(ans[g[0]], g[1])
	}

	return ans
}

// 如何找到拓扑关系
// 后序遍历之后进行翻转
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		postorder = []int{} // 后序遍历的结果
		onPath    = make([]bool, numCourses)
		visited   = make([]bool, numCourses)
		hasCycle  = false
	)

	graph := buildGraph(numCourses, prerequisites) // 建图 -> 这里需要建图的顺序翻译下，更符合直觉

	// node -> 当前遍历的节点编号
	// postorder -> 后序遍历的结果
	// onPath -> 当前遍历到的路径
	// visited -> 访问过的节点
	// 类比贪吃蛇游戏，visited 记录蛇经过过的格子，而 onPath 仅仅记录蛇身。onPath 用于判断是否成环，类比当贪吃蛇自己咬到自己（成环）的场景
	var traverse func(int, *[]int, *[]bool, *[]bool, *bool)

	traverse = func(node int, postorder *[]int, onPath, visited *[]bool, hasCycle *bool) {

		if (*onPath)[node] {
			*hasCycle = true
		}

		if *hasCycle || (*visited)[node] {
			return
		}

		// 正式开始遍历
		(*onPath)[node] = true
		(*visited)[node] = true

		for _, next := range graph[node] {
			traverse(next, postorder, onPath, visited, hasCycle)
		}

		(*postorder) = append((*postorder), node)

		(*onPath)[node] = false
	}

	for i := 0; i < numCourses; i++ {
		traverse(i, &postorder, &onPath, &visited, &hasCycle)
	}

	if hasCycle {
		return []int{}
	}

	reverse(&postorder)
	return postorder
}

func reverse(arr *[]int) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}
