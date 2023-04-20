package leetcode_golang

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[i][1]
	})

	res := make([][]int, 0)
	res = append(res, intervals[0])

	for idx := range intervals {
		if idx != len(intervals)-1 {
			if res[len(res)-1][1] >= intervals[idx+1][0] && res[len(res)-1][1] < intervals[idx+1][1] {
				res[len(res)-1][1] = intervals[idx+1][1]
			}

			if intervals[idx+1][0] > res[len(res)-1][1] {
				res = append(res, intervals[idx+1])
			}
			fmt.Println(res)
		}
	}
	return res
}

// 双指针？
func merge2(intervals [][]int) [][]int {
	i, j := 0, 1

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)

	for j < len(intervals) {
		// 可以合并
		if intervals[j][0] <= intervals[i][1] {
			if intervals[i][1] <= intervals[j][1] {
				intervals[i][1] = intervals[j][1]
			}
		} else {
			res = append(res, intervals[i])
			i = j
		}
		j++
	}
	// append最后一个元素
	res = append(res, intervals[i])
	return res

}
