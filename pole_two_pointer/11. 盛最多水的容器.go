package pole_two_pointer

import "leetcode_golang/library"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	ans := 0

	for left < right {
		// 宽
		ans = library.Max(ans, (right-left)*library.Min(height[left], height[right]))

		// 迁移短的那根，因为长的再怎么迁移，都不会更大
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return ans
}
