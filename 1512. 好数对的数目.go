package leetcode_golang

func numIdenticalPairs(nums []int) int {
	var (
		hashMap = make(map[int]int)
		ans     = 0
	)

	for _, num := range nums {
		if hashMap[num] > 0 {
			ans += hashMap[num]
		}
		hashMap[num]++
	}
	return ans
}
