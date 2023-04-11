package main

func main() {
	println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, target int) int {
	// 为正所有的和 = p
	// 所有元素的和 s
	// 负数 = s - p
	// 修改之后的和 p-(s-p) = t(target)
	// 2p = s + t
	// p = (s+t)/2 -> 选一些数字，他的和是(sum(nums)+targte)/2

	target += sum(nums)

	if target < 0 || target%2 != 0 {
		return 0
	}

	target = target / 2

	// 存一下cache
	cache := make([][]int, len(nums))

	for i := range cache {
		cache[i] = make([]int, target+1)
		for j := range cache[i] {
			cache[i][j] = -1 // -1 表示没用访问过
		}
	}

	var dfs func(int, int) int

	dfs = func(i, c int) (res int) {
		if i < 0 {
			// 找到了一个合适的方案
			if c == 0 {
				return 1
			}
			return 0
		}
		tmp := cache[i][c]
		if tmp != -1 {
			return tmp
		}

		defer func() {
			cache[i][c] = res
		}()

		// 已经不能再装了
		if nums[i] > c {
			return dfs(i-1, c)
		}
		// 从后装起
		return dfs(i-1, c) + dfs(i-1, c-nums[i])
	}

	return dfs(len(nums)-1, target)

}

func sum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
