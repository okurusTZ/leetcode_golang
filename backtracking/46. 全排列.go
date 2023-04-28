package backtracking

func permute(nums []int) [][]int {

	var f func(int)

	var n = len(nums)
	var res = make([][]int, 0)
	var path = make([]int, n) // 当前遍历路径
	var onPath = make([]bool, n)

	f = func(i int) {
		if i == n {
			res = append(res, append([]int(nil), path...))
			return
		}

		for j, on := range onPath {
			if on {
				continue
			}
			path[i] = nums[j]
			onPath[j] = true
			f(i + 1)
			onPath[j] = false
		}
	}
	f(0)
	return res
}
