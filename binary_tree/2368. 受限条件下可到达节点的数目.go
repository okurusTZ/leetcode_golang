package binary_tree

// 这里的queue是int的，但我懒得改成泛型了，所以就这样吧
func reachableNodes(n int, edges [][]int, restricted []int) int {
	if n == 0 {
		return 0
	}

	g := make([][]int, 100000)

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	q := &Queue{
		0,
	}

	restrictedMap := make(map[int]bool)

	for _, restrict := range restricted {
		restrictedMap[restrict] = true
	}

	ans := 1

	for q.Length() != 0 {
		node := q.Pop().(int)

		for _, y := range g[node] {

			if restrictedMap[y] {
				continue
			}

			q.Push(y)
			ans++

		}

		restrictedMap[node] = true
	}

	return ans
}

func reachableNodes2(n int, edges [][]int, restricted []int) int {

	if n == 0 {
		return 0
	}

	g := make([][]int, 100000)

	restrictedMap := make(map[int]bool)

	for _, restrict := range restricted {
		restrictedMap[restrict] = true
	}

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int)

	ans := 0

	dfs = func(i int) {
		// 到这个节点上了，本身应该加一
		ans += 1
		for j := range g[i] {
			if !restrictedMap[j] {
				dfs(j)
			}
		}
		restrictedMap[i] = true
	}

	return ans
}

func reachableNodes(n int, edges [][]int, restricted []int) int {

	g := make([][]int, n)

	r := make(map[int]bool)

	for _, restrict := range restricted {
		r[restrict] = true
	}

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if !r[y] && !r[x] {
			g[x] = append(g[x], y)
			g[y] = append(g[y], x)
		}
	}

	var dfs func(int, int)

	ans := 0
	dfs = func(i, f int) {
		// 到这个节点上了，本身应该加一
		ans += 1
		for _, j := range g[i] {
			if j != f {
				dfs(j, i)
			}
		}
	}

	dfs(0, -1)

	return ans
}
