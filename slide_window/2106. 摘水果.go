package slide_window

import "sort"

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)

	// 找到最左的位置，开始滑动
	left := sort.Search(n, func(i int) bool { return fruits[i][0] >= startPos-k })
	var (
		right = left
		ans   = 0
		s     = 0
	)

	for ; right < n && fruits[right][0] < startPos; right++ {
		s += fruits[right][1] // 从 fruits[left][0] 到 startPos 的水果数 -> 默认全摘了
	}
	ans = s
	for ; right < n && fruits[right][0] <= startPos+k; right++ {
		s += fruits[right][1]
		// 这时候需要移动了
		// left 1 2 3 startPos 5 6 right
		// 先去左边后去右边 && 先去右边后去左边
		for fruits[right][0]*2-fruits[left][0]-startPos > k && fruits[right][0]-fruits[left][0]*2+startPos > k {
			s -= fruits[left][1] // fruits[left][0] 无法到达
			left++
		}
		if s > ans {
			ans = s
		}
	}
	return ans
}
