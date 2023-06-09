package pole_two_pointer

import "leetcode_golang/library"

func minOperations(nums []int, x int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	target := sum - x // 转换后的目标

	// 越界问题就是因为left走到最后还没有凑齐，所以提前处理就行
	if target < 0 {
		return -1
	}

	left := 0
	ans := len(nums) + 1
	sum = 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum > target {
			sum -= nums[left]
			left++
		}

		if sum == target {
			ans = library.Min(ans, len(nums)-(right-left+1))
		}

	}

	if ans > len(nums) {
		return -1
	}
	return ans
}
