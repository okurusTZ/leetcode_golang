package main

// 输入：s = "QQQW"
// 输出：2
// 解释：我们可以把前面的 "QQ" 替换成 "ER"

func balancedString(s string) int {
	left := 0
	ans := len(s) // 注意求最小值时候的初始化问题

	k := len(s) / 4    // 目标长度
	list := ['X']int{} // X在W的下一位，方便创建一个计数器

	println(len(list))

	// [l, r]需要替换

	for _, c := range s {
		list[c]++
	}

	if list['Q'] == k && list['W'] == k && list['E'] == k && list['R'] == k {
		return 0
	}

	for right := 0; right < len(s); right++ {
		list[s[right]]-- // 需要变

		// 有可能的情况，外面的case都小于k
		for list['Q'] <= k && list['W'] <= k && list['E'] <= k && list['R'] <= k {
			ans = min(ans, right-left+1)
			// 看看有没有更小的
			list[s[left]]++
			left++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	balancedString("QQQQ")
}
