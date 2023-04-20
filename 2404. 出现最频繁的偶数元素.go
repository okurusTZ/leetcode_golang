package leetcode_golang

func mostFrequentEven(nums []int) int {
	hashMap := make(map[int]int)
	max := -1
	res := -1
	for _, num := range nums {
		if num%2 == 1 {
			continue
		}
		hashMap[num] += 1
		if hashMap[num] > max || ((res >= num || res == -1) && hashMap[num] == max) {
			max = hashMap[num]
			res = num
		}
	}
	return res
}
