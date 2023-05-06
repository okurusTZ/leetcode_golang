package leetcode_golang

func containsNearbyDuplicate(nums []int, k int) bool {
	var (
		hashMap = make(map[int]int)
	)

	for idx, num := range nums {
		if pos, ok := hashMap[num]; ok && idx-pos <= k {
			return true
		}
		hashMap[num] = idx
	}
	return false
}
