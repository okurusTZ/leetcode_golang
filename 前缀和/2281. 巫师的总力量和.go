package main

func totalStrength(strength []int) int {

	var (
		stack = make([]int, 0)
		left  = make([]int, len(strength)) // left[i] -> strength[i]为最小值时， 向左找到的第一个比strength[i]小的值
		right = make([]int, len(strength)) // right[i] -> strength[i]为最小值时， 向右找到的第一个比strength[i]小的值
	)

	for i := range left {
		left[i] = -1
	}

	// 如果右边不存在，下标就是i
	for i := range right {
		right[i] = len(strength)
	}

	for i, s := range strength {
		// 这里应该是左边第一个小于等于的数，不然有重复的数字出现时，会重复计算
		for len(stack) > 0 && strength[stack[len(stack)-1]] >= s {
			stack = stack[:len(stack)-1]
		}
		// 找到第一个小于的数字
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)
	for i := len(strength) - 1; i >= 0; i-- {
		for len(stack) > 0 && strength[stack[len(stack)-1]] > strength[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// 前缀和
	var (
		s       = 0
		ss      = make([]int, len(strength)+2)
		ans     = 0
		mod int = 1e9 + 7 // 防止溢出
	)

	// ss[0] = 0
	// ss[1] = ss[0]+strength[1]
	for i, v := range strength {
		s += v
		ss[i+2] = (ss[i+1] + s) % mod
	}

	for i, v := range strength {
		l, r := left[i]+1, right[i]-1 // [l,r] 左闭右闭
		// 数学知识，懒得想了
		total := ((i-l+1)*(ss[r+2]-ss[i+1]) - (r-i+1)*(ss[i+1]-ss[l])) % mod
		ans = (ans + v*total) % mod // 累加贡献
	}

	return (ans + mod) % mod

}
