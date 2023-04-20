package leetcode_golang

// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4]

// 数组翻转
// 1. 翻转所有
// 2. 翻转[0,k(mod)n−1]
// 3. 翻转[k(mod)n, n-1]
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums)
	// 注意取subArr的边界
	reverse(nums[:k])
	reverse(nums[k:n])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate3(nums []int, k int) {
	// 旋转之后的位置
	n := len(nums)
	newNums := make([]int, n)
	for idx := range nums {
		newIdx := (idx + k) % n
		newNums[newIdx] = nums[idx]
	}
	copy(nums, newNums)
}

// 可以做，会超时，因为本质上所有的大于len(nums)的k都是无意义的
// 修改之后还是超时
func rotate2(nums []int, k int) {
	newNums := nums
	k = k % len(nums)
	for k > 0 {
		tmp := newNums
		newNums = []int{tmp[len(tmp)-1]}
		newNums = append(newNums, tmp[:len(tmp)-1]...)
		k--
	}
	copy(nums, newNums)
}
