package leetcode_golang

func queryString(s string, n int) bool {
	// 枚举子字符串，利用map记录下对应数值，最后计算map的长度
	var hit = make(map[int]bool)

	for i, b := range s {
		// 1 或 0
		x := int(b - '0')
		if x == 0 {
			// 如果以0开头，就和下一位开头是一致的
			continue
		}

		for j := i + 1; j < len(s); j++ {
			if x > n {
				break
			}
			hit[x] = true
			x = x<<1 | int(s[j]-'0')
		}

		if x <= n {
			hit[x] = true
		}

		// for j := i + 1; x <= n; j++ {

		// 	hit[x] = true

		// 	if j >= len(s) {
		// 		break
		// 	}

		// 	// O(1)计算出下一个数
		// 	x = x<<1 | int(s[j]-'0')
		// }
	}

	return len(hit) == n

}
