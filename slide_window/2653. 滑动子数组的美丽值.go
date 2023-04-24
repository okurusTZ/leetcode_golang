package slide_window

func getSubarrayBeauty(nums []int, k int, x int) []int {

	bucket := make([]int, 101) // 因为数据是-50-50

	for i := 0; i < k; i++ {
		bucket[nums[i]+50]++
	}

	ans := make([]int, len(nums)-k+1)
	sum := 0
	for idx, j := range bucket {
		sum += j
		if sum >= x {
			if idx-50 < 0 {
				ans[0] = idx - 50
			}
			break
		}
	}

	for i := 1; i < len(nums)-k+1; i++ {
		bucket[nums[i+k-1]+50]++
		bucket[nums[i-1]+50]--
		sum := 0
		for idx, j := range bucket {
			sum += j
			if sum >= x {
				if idx-50 < 0 {
					ans[i] = idx - 50
				}
				break
			}
		}
	}

	return ans
}
