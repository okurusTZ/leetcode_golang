package mathmatic

func sumOfMultiples(n int) int {
	// 返回m倍数的数字之和
	var f = func(m int) int {
		// 1, 2, ... n/m
		return (1 + n/m) * (n / m) / 2 * m
	}
	// 如果不互斥的话，返回最小公倍数
	return f(3) + f(5) + f(7) - f(15) - f(21) - f(35) + f(105)
}
