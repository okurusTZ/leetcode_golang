package leetcode_golang

import "unicode"

// 双指针
func camelMatch(queries []string, pattern string) []bool {

	res := make([]bool, len(queries))

	for i, query := range queries {
		res[i] = true
		p := 0
		for _, c := range query {
			if p < len(pattern) && pattern[p] == byte(c) {
				p++
			} else {
				if unicode.IsUpper(c) {
					res[i] = false
				}
			}
		}

		if p < len(pattern) {
			res[i] = false
		}
	}

	return res
}
