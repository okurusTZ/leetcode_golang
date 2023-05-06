package dynamic_programing

func maxProduct(nums []int) int {
	var (
		maxP, minP = nums[0], nums[0]
		ans        = nums[0]
	)

	for idx, num := range nums {
		if idx == 0 {
			continue
		}
		tmpMax, tmpMin := maxP, minP

		maxP = max(num, max(tmpMax*num, tmpMin*num)) // 因为是子数组，所以不连续的情况下，只能保存当前数值num
		minP = min(num, min(tmpMax*num, tmpMin*num)) // 需要注意为负数的情况
		ans = max(maxP, ans)
		// fmt.Println(maxP, minP, num)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
