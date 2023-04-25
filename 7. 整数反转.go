package leetcode_golang

import "math"

func reverse(x int) (ans int) {

	defer func() {
		if ans > math.MaxInt32 || ans < -1*math.MaxInt32 {
			ans = 0
		}
	}()

	ans = 0
	negative := false
	if x < 0 {
		negative = true
		x = -x
	}

	for x > 0 {
		current := x % 10
		ans += current
		x = x / 10
		ans = ans * 10
	}

	if negative {
		return -1 * ans / 10
	}

	return ans / 10
}
