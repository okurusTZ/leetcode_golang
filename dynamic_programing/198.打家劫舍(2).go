package dynamic_programing

import "leetcode_golang/library"

// 回溯+存储
func rob2(nums []int) int {

	f0, f1 := 0, 0 // 滚动
	// f0 -> 前两个
	// f1 -> 前一个

	for _, num := range nums {
		new_f := library.Max(f1, f0+num)
		f0 = f1
		f1 = new_f
	}

	return f1

}
