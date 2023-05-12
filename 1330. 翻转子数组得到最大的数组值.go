package leetcode_golang

import "math"

func maxValueAfterReverse(nums []int) int {

	// 假设原来的数组值为base
	var (
		base = 0
		mn   = math.MaxInt
		d    = 0
		mx   = math.MinInt
	)

	// 如果翻转的是 i - j, i < j
	// 存在 1 <= i < j < len(nums)-1, 假设a=nums[i-1], b = nums[i], c = nums[j], d = nums[j+1]
	// 原来为 abs(a-b) + ... + abs(c-d)
	// 现在变为 abs(a-c) + ... + abs(b-d)
	// 变化的值为 abs(a-c) + abs(b-d) - abs(a-b) - abs(c-d)， 所以要求这个值最大
	// 获取这4个值的大小关系，一共有 4! 种可能
	// 只需要讨论哪两个数是最小的就行
	// 其实这里只需要讨论 a b < c d 的情况就可以
	// d = 2 * min(a,b) - 2 * max(c,d) -> 为什么其他不讨论，因为都小于0
	// d = 2 * min(c,d) - 2 * max(a,b) -> 对称的

	for i := 1; i < len(nums); i++ {
		// 基础值
		a := nums[i-1]
		b := nums[i]
		base += abs(a - b)
		mn = min(mn, max(a, b))
		mx = max(mx, min(a, b))

		// 特殊情况，翻转边界
		d = max(d, max(abs(nums[0]-b)-abs(a-b), // i=0
			abs(nums[len(nums)-1]-a)-abs(a-b))) // j=n-1
	}

	return base + max(d, 2*(mx-mn))
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
