package dynamic_programing

// 回溯+存储
func rob(nums []int) int {

	f0, f1 := 0, 0 // 滚动
	// f0 -> 前两个
	// f1 -> 前一个

	for _, num := range nums {
		new_f := max(f1, f0+num)
		f0 = f1
		f1 = new_f
	}

	return f1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
// }
