package leetcode_golang

func removeElement2(nums []int, val int) int {
	var (
		left  = 0 // 下一个赋值的元素
		right = 0 // 当前处理的元素
		ans   = 0
	)

	for right < len(nums) {
		if nums[right] != val {
			nums[left], nums[right] = nums[right], nums[left]
			nums[left]++
			ans++
		}
		nums[right]++
	}
	return ans
}

func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
