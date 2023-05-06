package stack

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var (
		ans      = ""
		stack    = make([]string, 0)
		digitNum = 0
	)

	for _, item := range s {
		// 如果是数字
		if item >= '0' && item <= '9' {
			stack = append(stack, string(item))
			digitNum++
		}

		if item == '[' {
			// 合并digit
			d := stack[len(stack)-digitNum : len(stack)]
			d2 := strings.Join(d, "")
			// fmt.Println(d2)
			stack = stack[:len(stack)-digitNum]
			stack = append(stack, d2)
			stack = append(stack, string(item))
			digitNum = 0
		}

		// 如果是/a-z/
		if item >= 'a' && item <= 'z' {
			stack = append(stack, string(item))

		}
		if item == ']' {
			var tmp = ""
			for stack[len(stack)-1] != "[" {
				top := stack[len(stack)-1]
				tmp = fmt.Sprintf("%v%v", top, tmp)
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // pop [
			c := stack[len(stack)-1]
			count, _ := strconv.Atoi(c)
			stack = stack[:len(stack)-1] // pop Number
			t := strings.Repeat(tmp, count)
			stack = append(stack, t) // 应对嵌套，这里要重新放回去
		}
	}

	for _, s := range stack {
		ans += s
	}
	return ans
}
