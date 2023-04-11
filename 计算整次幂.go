package main

import (
	"fmt"
)

// 分治思想
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return QuickMul(x, n)
	}
	return 1.0 / QuickMul(x, n)
}

// 计算x^77
// x^77 -> x^38 * x^38 * x -> x^19 -> x^9 * x^9 * x -> x^4 * x^4 * x -> x^2 * x^2 -> x
func QuickMul(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	// 获取一半的数据
	tmp := QuickMul(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}

func main() {
	fmt.Println(myPow(0.00001, 2147483647))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
