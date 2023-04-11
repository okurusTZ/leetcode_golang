package main

import "sort"

// O(n)
// 空间复杂度O(n)
func containsDuplicate(nums []int) bool {
	var numberMap = make(map[int]int)
	for _, num := range nums {
		if numberMap[num] != 0 {
			return true
		}
		numberMap[num] = num
	}
	return false
}

// 排序O(nlogn) -> 排序算法最坏情况下 nlogn, 遍历一边 n
// 空间复杂度O(logn)
func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	for idx := range nums {
		if idx != len(nums)-1 {
			if nums[idx] == nums[idx+1] {
				return true
			}
		}
	}
	return false
}

func main() {}
