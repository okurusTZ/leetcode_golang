package dynamic_programing

func climbStairs(n int) int {
	var (
		// 上一次
		last  = 1
		last2 = 1
		cur   = 1
	)

	for i := 2; i < n+1; i++ {
		last2 = last
		last = cur
		cur = last + last2
	}

	return cur
}
