package pole_two_pointer

// 输入：s = "abab"
// 输出："bab"
// 解释：我们可以找出 7 个子串 ["a", "ab", "aba", "abab", "b", "ba", "bab"]。按字典序排在最后的子串是 "bab"。
func lastSubstring(s string) string {
	var (
		j = 0
		n = len(s)
		k = 0
	)

	// 从i开始，到末尾一定是最大的前缀
	// 只需要考虑开始节点？
	for i := 1; i+k < n; {
		if s[i+k] == s[j+k] {
			k++
		} else {
			// 当前的大于历史的
			if s[i+k] > s[j+k] {
				// dacdae - acdae - cdae -> dae -> j = j + 3 = 0 + 3 = 3 -> i = j + 1 = 4
				// 这里说明s[i, i+k]的字典序大于s[j, j+k] -> s[j+1,j+k]也不会是最大后缀
				// j直接跳到k+1的位置，因为前面开始不会再比[i,i+k]更大了
				// 能走到这里，说明[j,j+k]之间的都和[i,i+k]相等，所以直接跳到j+k+1就行
				j += k + 1 // 这里会导致j>=i

				// 为什么要跳过去，因为前面的都比过了，已经没意义了
				if j >= i {
					i = j + 1
				}
			} else {
				i += k + 1
			}
			k = 0
		}
	}

	return s[j:n]
}

func main() {
	lastSubstring("vmjtxddvzmwrjvfamgpoowncslddrkjhchqswkamnsitrcmnhn")
}
