package leetcode_golang

import "math"

func canJump(nums []int) bool {
	var maxLength = 0

	// 注意边界问题
	if len(nums) <= 1 {
		return true
	}

	for i, num := range nums {
		maxLength = max(maxLength, i+num)
		// 已经可以跳到末尾的情况
		if maxLength >= len(nums)-1 {
			return true
		}
		// fmt.Println(maxLength)
		if maxLength <= i {
			return false
		}
	}
	return true
}

// 45. 跳跃游戏 II
func jump2(nums []int) int {
	var (
		dp = make([]int, len(nums))
	)

	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[0] = 0

	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i]; j++ {
			// 从i跳过来
			if i+j < len(nums) && dp[i+j] > dp[i]+1 {
				dp[i+j] = dp[i] + 1
			}
		}
	}
	// fmt.Println(dp)

	return dp[len(dp)-1]
}

func jump(nums []int) int {

	var (
		maxLength = 0
		end       = 0
		ans       = 0
	)

	for i := 0; i < len(nums)-1; i++ {
		// 当前可以到达的最远距离，每次选择跳得更远的
		maxLength = max(maxLength, i+nums[i])
		if i == end {
			end = maxLength
			ans++
		}
	}
	return ans
}
