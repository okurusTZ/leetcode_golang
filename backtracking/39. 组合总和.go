package backtracking

func combinationSum(candidates []int, target int) [][]int {

	var (
		res     = make([][]int, 0)
		current = make([]int, 0)
	)
	var f func(int, int)
	f = func(i int, target int) {
		if target == 0 {
			// deepcopy一份
			res = append(res, append([]int(nil), current...))
			return
		}
		if i == len(candidates) {
			return
		}

		// 两种选择
		// 不选
		f(i+1, target)
		if target-candidates[i] > 0 {
			current = append(current, candidates[i])
			f(i, target-candidates[i])
			current = current[:len(current)-1]
		}
	}
	f(0, target)
	return res
}
