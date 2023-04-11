package main

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		// println(left, right, mid, nums[left], nums[right], nums[mid])
		// 注意这里应该是小于最后一个数
		if nums[mid] > nums[len(nums)-1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	println(findMin([]int{2, 3, 4, 5, 1}))
}
