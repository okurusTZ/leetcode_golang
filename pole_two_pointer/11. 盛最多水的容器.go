package main

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	ans := 0

	for left < right {
		// 宽
		ans = max(ans, (right-left)*min(height[left], height[right]))

		// 迁移短的那根，因为长的再怎么迁移，都不会更大
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
