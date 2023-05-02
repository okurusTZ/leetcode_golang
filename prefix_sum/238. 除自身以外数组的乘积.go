package prefix_sum

func productExceptSelf(nums []int) []int {
	// 索引左边的数字 * 索引右边的数字
	var (
		left  = make([]int, len(nums))
		right = make([]int, len(nums))
		ans   = make([]int, len(nums))
	)

	for i := range nums {
		if i == 0 {
			left[i] = 1
		} else {
			left[i] = left[i-1] * nums[i-1]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			right[i] = 1
		} else {
			right[i] = right[i+1] * nums[i+1]
		}
	}

	for i := 0; i < len(nums); i++ {
		ans[i] = left[i] * right[i]
	}
	return ans
}
