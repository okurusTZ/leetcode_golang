package leetcode_golang

import (
	"fmt"
	"leetcode_golang/library"
)

func lengthOfLongestSubstring(s string) int {
	var (
		tmpStr = ""
		maxLen = 0
	)
	for _, sByte := range s {
		str := string(sByte)
		// 如果包含，获取最前面出现过的位置
		idx := IndexOf(tmpStr, str)
		tmpStr += str
		if idx > -1 {
			tmpStr = tmpStr[idx+1:]
		}
		if len(tmpStr) > maxLen {
			maxLen = len(tmpStr)
		}
	}
	return maxLen
}

func IndexOf(arr string, substr string) int {
	for idx, b := range arr {
		if string(b) == substr {
			return idx
		}
	}
	return -1
}

func lengthOfLongestSubstring2(s string) int {

	ans := 0
	left := 0 // 左指针
	hashMap := make(map[byte]int, 26)

	for right := 0; right < len(s); right++ {
		hashMap[s[right]-'a']++

		// 右移left
		for hashMap[s[right]-'a'] > 1 {
			hashMap[s[left]-'a']--
			left++
		}

		fmt.Println(right, left, hashMap)

		// 记录
		ans = library.Max(ans, right-left+1)
	}

	return ans
}
