package mathmatic

// 抑或运算，判断两个值是否不同，如果相同返回0
// 抑或运算有交换律
func singleNumber(nums []int) int {
	var (
		ans = 0
	)

	for i := 0; i < len(nums); i++ {
		ans = ans ^ nums[i]
	}
	return ans
}
