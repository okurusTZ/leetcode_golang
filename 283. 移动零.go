package leetcode_golang

func moveZeroes(nums []int) {
	idx := 0
	for i := 0; i < len(nums) && idx < len(nums); i++ {
		if nums[idx] == 0 {
			tmp := append(nums[:idx], nums[idx+1:]...)
			tmp = append(tmp, 0)
			nums = tmp
		} else {
			idx++
		}
	}
}

func moveZeroes2(nums []int) {
	var (
		left  = 0 // 处理完的非0数组的末尾 -> 如果right不是0，那么放到这里就可以
		right = 0
	)

	// 0,1,0,3,12

	for right < len(nums) {

		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
