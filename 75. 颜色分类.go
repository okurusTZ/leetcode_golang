package leetcode_golang

func sortColors(nums []int) {
	var (
		left  = 0
		right = len(nums) - 1
		i     = 0
	)

	for i <= right {
		if nums[i] == 0 {
			// 与左端点交换
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
		if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			continue
		}
		i++
	}
}
