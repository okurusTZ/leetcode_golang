package hashmap

func firstMissingPositive(nums []int) int {
	// 对于长度为n的nums，最小未出现的正数一定在[1, n]之间
	// 所以遍历，对于遍历到的num，把nums[num-1]打上标记，表示该数已经出现过
	// 如何在保持原有数据的情况下还打上标记， 变成对应的负数就可以
	// 所以先要把本身是负数的情况消除影响

	var (
		n = len(nums)
	)

	for i, num := range nums {
		if num <= 0 {
			nums[i] = n + 1 // 移除负数的影响
		}
	}

	for _, num := range nums {
		if abs(num) <= n {
			nums[abs(num)-1] = -abs(nums[abs(num)-1])
		}
	}

	// fmt.Println(nums)

	for idx, num := range nums {
		if num > 0 {
			return idx + 1
		}
	}
	// 如果每一个数都是负数，那么说明恰好是[1,n]
	return n + 1
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
