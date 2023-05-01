package dfs

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	var (
		g = make([][]int, n) // 获得每个人的下属关系
	)

	for i, m := range manager {
		if m >= 0 {
			g[m] = append(g[m], i) // 建树
		}
	}

	var dfs func(int) int
	dfs = func(i int) (ans int) {
		for _, child := range g[i] {
			ans = max(ans, dfs(child))
		}
		// 遍历完所有的child，加上当前这一次需要的通知时间
		return ans + informTime[i]
	}

	return dfs(headID)
}
