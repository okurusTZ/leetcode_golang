package main

import "fmt"

const _WOKR_TIME_THRESHOLD = 8

// 前缀和
// 假设劳累的天数hours[i] = 1; 不劳累的天数hours[i] = -1
// 那么就是求 前缀和 i < j && s[i] < s[j] -> j-i中间的和都是大于0的，求i-j的最大长度
// 单调递减栈 [0,-1,-2,-3,]
func longestWPI2(hours []int) int {

	n := len(hours)

	s := make([]int, n+1) // 前缀和
	stack := []int{0}

	for i, hour := range hours {
		if hour > _WOKR_TIME_THRESHOLD {
			s[i+1] = s[i] + 1
		} else {
			s[i+1] = s[i] - 1
		}

		// 小于的时候append上去，是个递减栈
		if s[i+1] < s[stack[len(stack)-1]] {
			stack = append(stack, i+1)
		}
	}

	ans := 0

	// 这里本身i已经是原来的idx基础上+1了的
	for i := n; i >= 1; i-- {

		for len(stack) > 0 && s[i] > s[stack[len(stack)-1]] {
			ans = max(ans, i-stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	// fmt.Println(s)

	return ans
}

// "要想算出比 s[i]−1 更小的数，必然会先算出 s[i]−1，那么这些更小数必然在 s[i]−1 首次出现的位置的右边。"
func longestWPI(hours []int) (ans int) {

	// 对于len(hours) 来说，因为每一次都是+1/-1，所以取值范围可以认为是
	// n = 1 -> [-1,1]
	// n = 2 -> [-2,2]
	// n = 3 -> [-3,3]
	pos := make([]int, len(hours)*2) // 记录前缀和首次出现的位置
	s := 0

	for i := 1; i <= len(hours); i++ {

		hour := hours[i-1]

		if hour > _WOKR_TIME_THRESHOLD {
			s--
		} else {
			s++
		}

		if s < 0 {
			// 前缀和大于0，可以认为[0,i]都是可用的
			ans = i
		} else {
			// 因为s < 0 的时候s-1一定小于0
			if pos[s+1] > 0 {
				ans = max(ans, i-pos[s+1])
			}
			if pos[s] == 0 {
				pos[s] = i
			}
		}

	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9}))
}
