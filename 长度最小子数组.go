package main

func minSubArrayLen(target int, nums []int) int {
	left := 0
	n := len(nums)
	res := n + 1
	sum := 0
	// 枚举右端点
	for right, x := range nums {

		sum += x // 先加上右端点

		// 大于了，left左移
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left += 1
		}
	}

	// 特殊情况，遍历了也没有命中
	if res == n+1 {
		return 0
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {}
