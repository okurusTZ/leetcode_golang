package leetcode_golang

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for idx, num := range nums {
		if tmp[target-num] != 0 {
			return []int{tmp[target-num] - 1, idx}
		}
		tmp[num] = idx + 1
	}
	return []int{0, 0}
}
