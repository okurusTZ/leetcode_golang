package leetcode_golang

func isValid(s string) bool {
	var stack = []rune{}

	for _, item := range s {
		if item == 'a' {
			stack = append(stack, item)
		}
		if item == 'b' {
			if len(stack) == 0 || stack[len(stack)-1] != 'a' {
				return false
			}
			stack = append(stack, item)
		}
		if item == 'c' {
			if len(stack) == 0 || stack[len(stack)-1] != 'b' {
				return false
			}
			stack = stack[:len(stack)-2]
		}
	}
	return len(stack) == 0
}
