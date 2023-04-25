package leetcode_golang

import "sort"

func sortPeople(names []string, heights []int) []string {

	ans := make([]int, len(names))

	for i := range names {
		ans[i] = i
	}

	sort.Slice(ans, func(i, j int) bool {
		return heights[ans[i]] >= heights[ans[j]]
	})

	// fmt.Println(ans)

	res := make([]string, len(names))

	for i, item := range ans {
		res[i] = names[item]
	}

	return res
}
