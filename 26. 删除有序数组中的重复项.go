package leetcode_golang

func removeDuplicates(nums []int) int {
	var left = 0
	for idx, num := range nums {
		if idx == 0 {
			left++
		} else {
			if num != nums[idx-1] {
				nums[left] = num
				left++
			}
		}
	}
	return left
}

// 80. 删除有序数组中的重复项 II
func removeDuplicatesii(nums []int) int {
	var (
		fast = 0
		slow = 0
	)

	for fast < len(nums) {
		// 如果等于，中间的都可以跳过
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++

	}

	return slow
}
