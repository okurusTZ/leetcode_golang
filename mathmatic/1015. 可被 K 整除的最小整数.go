package mathmatic

func smallestRepunitDivByK(k int) int {

	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	var x = 1

	remainder := x % (k)
	if remainder == 0 {
		return 1
	}

	for i := 2; ; i++ {
		x = (10*x%k + 1) % k // 为什么这里可以直接用余数替代

		// Xnew = (10 * Xold + 1)
		// Xold = Xold mod k
		// 一定存在X = p * k + q
		// Xold = pold * k + qold
		// Xnew = 10 * pold * k + 10 * qold + 1
		// 分别取余
		// Xnew % k = (10 * pold * k + 10 * qold + 1) % k
		// Xnew % k = (10 * qold + 1) % k
		// 因为Xold = pold * k + qold -> Xold % k = qold

		// Xnew % k = (10 * Xold % k + 1) % k

		if x == 0 {
			return i
		}
	}
}
