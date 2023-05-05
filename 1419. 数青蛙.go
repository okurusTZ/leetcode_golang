package leetcode_golang

func minNumberOfFrogs(croakOfFrogs string) int {
	var (
		croak = make([]int, 5)
		ans   = 0
	)
	for _, item := range croakOfFrogs {
		switch byte(item) {
		case 'c':
			croak[0]++
		case 'r':
			croak[1]++
		case 'o':
			croak[2]++
		case 'a':
			croak[3]++
		case 'k':
			croak[4]++
		}

		for j := 1; j <= 4; j++ {
			if croak[j] > croak[j-1] {
				return -1
			}
		}

		if croak[0]-croak[4] > ans {
			ans = croak[0] - croak[4] // 同时开始叫，但还没结束的个数 = 青蛙的并发数
		}

	}

	// 没叫完
	for j := 1; j <= 4; j++ {
		if croak[j] != croak[j-1] {
			return -1
		}
	}

	return ans
}
