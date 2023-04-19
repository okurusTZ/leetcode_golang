package dynamic_programing

import (
	"fmt"
	"sort"
)

func bestTeamScore(scores []int, ages []int) int {

	n := len(scores)

	var dynamic = make([][]int, n)

	for idx, score := range scores {
		dynamic[idx] = []int{ages[idx], score}
	}

	sort.Slice(dynamic, func(i int, j int) bool {
		if dynamic[i][1] != dynamic[j][1] {
			return dynamic[i][1] < dynamic[j][1]
		}
		return dynamic[i][0] < dynamic[j][0]
	})

	fmt.Println(dynamic)

	res := make([]int, n)
	sum := 0
	// 以i编号的人为结尾，所获得的最大分数
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			// i的分数一定大于等于j的
			// 如果i的年龄也大，就可以
			if dynamic[i][0] >= dynamic[j][0] {
				res[i] = Max(res[i], res[j])
			}
		}
		res[i] += dynamic[i][1]
		sum = Max(sum, res[i])
	}

	return sum
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(bestTeamScore([]int{1, 1, 1, 1, 1}, []int{88, 89, 12, 39, 2}))
// }
