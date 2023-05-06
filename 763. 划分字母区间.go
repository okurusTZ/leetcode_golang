package leetcode_golang

func partitionLabels(s string) []int {
	// 获取每个字母的最后的下标，表示如果要包含这个字母，那么最远需要包含到的距离
	var (
		farDistance = make(map[byte]int)
		maxDistance = 0
		start       = 0
		end         = 0
		ans         = make([]int, 0)
	)

	for i, item := range s {
		farDistance[byte(item)] = i
	}

	for end < len(s) {
		current := s[end]

		if farDistance[current] > maxDistance {
			maxDistance = farDistance[current]
		}

		if maxDistance == end {
			ans = append(ans, end-start+1)
			start = end + 1
			end = end + 1
		} else {
			end++
		}
	}
	return ans
}
