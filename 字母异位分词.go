package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	strMap := make(map[string][]int)
	for idx, str := range strs {
		strList := strings.Split(str, "")
		sort.Strings(strList)
		sorted := strings.Join(strList, "")
		if strMap[sorted] == nil {
			strMap[sorted] = make([]int, 0)
		}
		// 进一步优化，可以把每个已经排过序的哈希表记录下来
		strMap[sorted] = append(strMap[sorted], idx)
	}

	resp := make([][]string, len(strMap))

	var idx = 0
	for _, item := range strMap {
		resp[idx] = make([]string, 0)
		for _, stridx := range item {
			resp[idx] = append(resp[idx], strs[stridx])
		}
		idx++
	}

	return resp
}

// ai给的解法，相比原来，把string排序简化成了
func groupAnagrams2(strs []string) [][]string {

	strMap := make(map[[26]int][]string)

	for _, str := range strs {
		count := [26]int{}
		for _, s := range str {
			count[s-'a']++
		}
		if strMap[count] == nil {
			strMap[count] = make([]string, 0)
		}
		strMap[count] = append(strMap[count], str)
	}

	fmt.Println(strMap)

	resp := [][]string{}
	for _, item := range strMap {
		resp = append(resp, item)
	}

	return resp
}

func main() {
	groupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}
