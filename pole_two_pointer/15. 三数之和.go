package pole_two_pointer

import (
	"fmt"
	"sort"
)

// 双指针

// 1,2,3,4,5 n=5
func threeSum(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)

	// 升序排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		target := 0 - nums[i]
		left := i + 1
		right := n - 1

		for left < right {
			if nums[left]+nums[right] == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 跳过重复数字
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}

				for right--; left < right && nums[right] == nums[right+1]; right-- {
				}
			}
			if nums[left]+nums[right] < target {
				left++
			}
			if nums[left]+nums[right] > target {
				right--
			}

		}
	}
	return ans
}
