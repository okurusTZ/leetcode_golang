package main

import "fmt"

// 单调栈，存下自己第一个大于自己的数
func dailyTemperatures(temperatures []int) []int {

	var (
		stack  = make([]int, 0)
		higher = make([]int, len(temperatures))
	)

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			higher[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
	}

	return higher
}

func main() {
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}
