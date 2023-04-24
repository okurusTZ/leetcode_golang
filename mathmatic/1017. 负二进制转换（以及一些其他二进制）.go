package mathmatic

import (
	"fmt"
	"math"
)

/**
* 输入：n = 2
* 输出："110"
* 解释：(-2)2 + (-2)1 = 2
**/

// 类似于十进制

func baseNeg2(n int) string {
	res := ""
	for n != 0 {
		res = fmt.Sprintf("%v%s", math.Abs(float64(n%-2)), res)
		n = -(n >> 1) // 每次除以-2，相当于右移之后，变为相反数
	}
	if len(res) == 0 {
		return "0"
	}
	return res
}

func main() {
	fmt.Println(baseNeg2(2))
}

// 汉明重量
// 2进制中1的个数
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res += int(num % 2)
		num = num / 2
	}
	return res
}
