package main

// 一样是回溯
// 存在单调性
func numSubarrayProductLessThanK(nums []int, k int) int {

	// 左端点 l 右端点 r
	// 如果[l,..,r] 乘积小于target
	// [l,r], [l+1,r] ... [r,r] 都是满足要求的
	// 满足要求的有r-l+1个

	ans := 0
	product := 1
	left := 0

	// 严格小于，先排除边界条件
	if k <= 1 {
		return 0
	}

	for right := 0; right < len(nums); right++ {
		// 先计算一次
		product = product * nums[right]
		// 如果不符合要求，左移left
		for product >= k {
			product = product / nums[left]
			left++
		}

		ans += right - left + 1
	}
	return ans

}

func main() {
	println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 3))
}
