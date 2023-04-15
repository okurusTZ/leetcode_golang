package main

import (
	"fmt"
	"math"
)

// o(n^2) -> 超时
func find132pattern2(nums []int) bool {
	// 枚举i
	// 找到 (j,k) -> nums[j] > nums[k] && nums[k] > nums[i]
	// 也就是说(i,k)之间存在比nums[k]更大的数字

	stack := make([]int, 0)
	left := make([]int, len(nums))
	for i := range left {
		left[i] = -1
	}

	for i, num := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= num {
			// 弹出
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			// 获取左边最大的下标
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	fmt.Println(left)

	for i, num := range nums {
		for k := i + 1; k < len(nums); k++ {
			if nums[k] <= num {
				continue
			}
			if left[k] > i {
				fmt.Println(i, left[k], k, nums[i], nums[left[k]], nums[k])
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(find132pattern([]int{1, 2, 8, -1, 2}))
}

func find132pattern(nums []int) bool {

	stack := make([]int, 0)
	mid := math.MinInt32 // 记录下k这个数值

	for i := len(nums) - 1; i >= 0; i-- {

		fmt.Println("stack: ", stack, "mid:", mid, "idx:", i, "num: ", nums[i])

		if nums[i] < mid {
			return true
		}

		// 从栈里找到除了nums[i]以外最大的数值 -> 为什么是这个数，因为这个数最方便找到i
		// 不断弹出，因此结束循环之后，栈顶就是 > nums[1]的
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			// 因为是单调递减，所以这里记录的是小于nums[1]的最大的数字(idx在i之后的)
			mid = nums[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
		}

		// 入栈 -> 单调递减的栈
		stack = append(stack, i)
	}
	return false
}
