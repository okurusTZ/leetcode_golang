package pole_two_pointer

import (
	"math"
	"sort"
)

// 输入：nums = [-1,2,1,-4], target = 1
// 输出：2
// 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)

func threeSumClosest(nums []int, target int) int {

	// 排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := math.MaxInt64

	// fmt.Println(nums)

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			// 如果命中了，直接返回
			if target == nums[left]+nums[right]+nums[i] {
				return target
			}

			// 记录一下最小
			if abs(target-nums[i]-nums[left]-nums[right]) < abs(target-ans) {
				ans = nums[i] + nums[left] + nums[right]
			}

			if target < nums[left]+nums[right]+nums[i] {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
