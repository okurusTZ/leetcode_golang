package leetcode_golang

func subsets(nums []int) (resp [][]int) {
	// 选或者不选的问题
	var dfs func(int)
	n := len(nums)
	path := make([]int, 0, n)
	dfs = func(i int) {
		// 边界
		if n == i {
			tmp := make([]int, len(path))
			copy(tmp, path)
			resp = append(resp, tmp)
			return
		}

		// 不选
		dfs(i + 1)

		// 选
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
	}
	dfs(0)
	return resp
}
