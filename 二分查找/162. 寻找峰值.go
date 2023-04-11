package main

// 等于寻找最大值
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1 // [left,right)
	//[3,4)

	count := 0

	for left < right && count < 10 {
		mid := left + (right-left)/2
		println(left, right, mid, nums[left], nums[right], nums[mid])

		// 1,2,1,5,3,6,4 mid=5,mid+1=3
		// 1,2,1,3,5,6,4 mid=3;mid+1=5
		// [5,6,4] mid=5, mid+1=6
		// [5,6,4] 4 是蓝色，封顶或者封顶右侧
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
		count++
	}

	return left
}

func main() {
	println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
