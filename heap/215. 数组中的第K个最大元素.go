package heap

import (
	"container/heap"
	"sort"
)

func findKthLargest2(nums []int, k int) int {
	p := &hp{
		make([]int, 0),
	}
	heap.Init(p)

	for _, num := range nums {
		heap.Push(p, num)
	}

	// fmt.Println(p)

	for i := 0; i < len(nums)-k; i++ {
		heap.Pop(p)
	}

	// fmt.Println(p)

	return heap.Pop(p).(int)
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(value any) {
	h.IntSlice = append(h.IntSlice, value.(int))
}

func (h *hp) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// 快排序

func Partition(nums []int, start, end int) int {
	pivot := nums[start]
	l, r := start, end

	for l < r {
		for l < r && nums[r] >= pivot {
			r--
		}

		nums[l] = nums[r]
		for l < r && nums[l] <= pivot {
			l++
		}
		nums[r] = nums[l]
	}

	nums[l] = pivot
	return l
}

func TopKSplit(nums []int, k, start, stop int) {
	if start < stop {
		index := Partition(nums, start, stop)
		if index == k {
			return
		} else if index < k {
			TopKSplit(nums, k, index+1, stop)
		} else {
			TopKSplit(nums, k, start, index-1)
		}
	}
}
