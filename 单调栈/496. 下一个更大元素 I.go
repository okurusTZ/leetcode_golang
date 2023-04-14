package main

import "fmt"

// 对于nums2来说，先获取单调栈，再存map，到nums1里找
func nextGreaterElement(nums1 []int, nums2 []int) []int {

	stack := make([]int, 0)
	large := make(map[int]int)

	ans := make([]int, len(nums1))

	for i := len(nums2) - 1; i >= 0; i-- {
		// 注意这里要比较的是nums2里的大小，stack里存的是下标
		for len(stack) > 0 && nums2[stack[len(stack)-1]] < nums2[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			// 下一个大于的下标
			large[nums2[i]] = stack[len(stack)-1]
		} else {
			large[nums2[i]] = -1
		}

		stack = append(stack, i)
	}

	// fmt.Println(large)

	for i, num := range nums1 {
		// ans[i] = large[num]
		if large[num] == -1 {
			ans[i] = -1
		} else {
			ans[i] = nums2[large[num]]
		}
	}

	return ans
}

func main() {
	fmt.Println(nextGreaterElement([]int{2, 1, 3}, []int{2, 3, 1}))
}
