package leetcode_golang

func minOperations(nums []int) int {
	if gcds(nums) > 1 {
		return -1
	}

	count := 0

	for i := range nums {
		if nums[i] == 1 {
			count++
		}
	}

	if count > 0 {
		return len(nums) - count
	}

	min_size := len(nums)

	// 为什么是找最短的
	// 假设子数组长度为n，要操作n-1次才能获得一个1
	// 有一个1后，操作len(nums)-1次，全部变成1
	// 2,3 -> 操作1次
	// 5,10,2 -> 操作2次

	for i := 0; i < len(nums); i++ {
		g := 0

		for j := i; j < len(nums); j++ {
			g = gcd(nums[j], g)
			if g == 1 {
				min_size = min(min_size, j-i+1)
				break
			}
		}
	}

	return min_size + len(nums) - 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcds(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = gcd(res, nums[i])
	}
	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	if a == 0 {
		return b
	}
	return gcd(b, a%b)
}
