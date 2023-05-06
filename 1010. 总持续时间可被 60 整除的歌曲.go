package leetcode_golang

func numPairsDivisibleBy60(time []int) int {
	var (
		hashMap = [60]int{}
		sum     = 0
	)

	for _, t := range time {
		sum += hashMap[(60-t%60)%60]
		hashMap[t%60]++
	}
	return sum
}
