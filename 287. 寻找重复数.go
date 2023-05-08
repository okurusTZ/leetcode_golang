package leetcode_golang

func findDuplicate(nums []int) int {
	slow, fast := 0, 0

	// 为什么这里可以这么走；因为有下面这个条件约束
	// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	// 这里走完，快慢指针第一次相遇
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
