package pole_two_pointer

import "leetcode_golang/library"

// 思路，计算每一个格子能接的水
// 每一个格子能接的水 =  (min(left_height, right_height) - self_height) * 1
// 所以我们需要一个数组pre_height记录到i位置的前缀最大高度
// suffix_height 记录到i位置的后缀最大高度（i - len(height))
func trap(height []int) int {

	res := 0

	var (
		suffix_height = make([]int, len(height))
		prefix_height = make([]int, len(height))
	)

	for idx, h := range height {
		if idx == 0 {
			prefix_height[idx] = h
		} else {
			prefix_height[idx] = library.Max(height[idx], prefix_height[idx-1])
		}
	}

	for i := len(height) - 1; i >= 0; i-- {
		if i == len(height)-1 {
			suffix_height[i] = height[i]
		} else {
			suffix_height[i] = library.Max(height[i], suffix_height[i+1])
		}
	}

	// fmt.Print(prefix_height, suffix_height)

	for i := 0; i < len(height); i++ {
		if library.Min(prefix_height[i], suffix_height[i])-height[i] < 0 {
			continue
		}
		res += (library.Min(prefix_height[i], suffix_height[i]) - height[i]) * 1
	}
	return res
}
