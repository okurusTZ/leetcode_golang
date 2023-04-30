package leetcode_golang

func findAnagrams(s string, p string) []int {
	n := len(p)
	ans := []int{}
	// 定长，可以直接比较
	var sMap, pMap [26]int

	if len(p) > len(s) {
		return []int{}
	}

	// 初始化
	for _, i := range p {
		pMap[i-'a']++
	}

	for i := 0; i < n; i++ {
		sMap[s[i]-'a']++
	}

	if sMap == pMap {
		ans = append(ans, 0)
	}

	for i := 1; i < len(s)-n+1; i++ {
		sMap[s[i+n-1]-'a']++
		sMap[s[i-1]-'a']--
		if sMap == pMap {
			ans = append(ans, i)
		}
	}
	return ans
}
