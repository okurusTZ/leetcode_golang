package backtracking

func generateParenthesis(n int) []string {
	var f func(int)

	var (
		current  = ""
		res      = make([]string, 0)
		leftNum  = 0
		rightNum = 0
	)

	f = func(i int) {

		if leftNum > n || rightNum > n {
			return
		}

		if i == n*2 {
			res = append(res, current)
			return
		}

		// 放哪个
		if len(current) > 0 && leftNum > rightNum {
			current = current + ")"
			rightNum++
			f(i + 1)
			current = current[:len(current)-1]
			rightNum--
		}

		current = current + "("
		leftNum++
		f(i + 1)
		current = current[:len(current)-1]
		leftNum--
	}

	f(0)
	return res
}
