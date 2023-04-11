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

	// 为什么是2 -> 每次只需要用到前两个f，所以递推的时候，存两个数就行
	f := [2][]int{make([]int, target+1), make([]int, target+1)} // 避免出现负数下标
	// 可以参考递归的边界情况 -> 递归结束时，c=0，那么肯定是一种case
	f[0][0] = 1 // 初始化边界情况

	for i, num := range nums {
		for c := 0; c <= target; c++ {
			if num > c {
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = f[i%2][c] + f[i%2][c-num]
			}
		}
	}

	return f[len(nums)%2][target]
}

func sum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
