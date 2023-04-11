package main

import "fmt"

func longestPalindrome(s string) string {

	// 特殊判断
	if len(s) == 1 {
		return s
	}

	// 取一个字符为回文中心
	max := 0
	res := ""
	for idx := 1; idx < len(s); idx++ {

		tmp := 0
		current := string(s[idx])
		left := idx

		for left >= 0 && string(s[left]) == current {
			left--
		}

		if left < 0 {
			if idx-left > max {
				max = idx - left
				res = s[:idx]
			}
		}

		right := idx

		for right < len(s) && string(s[right]) == current {
			right++
		}

		if right >= len(s) {
			if right-idx > max {
				max = right - idx + 1
				res = s[idx:]
				continue
			}
		}

		for left >= 0 && right < len(s) && string(s[left]) == string(s[right]) {
			left--
			right++
		}
		tmp = right - left + 1
		if tmp > max {
			max = tmp
			res = s[left+1 : right]
		}
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("aaacdcaa"))
}
