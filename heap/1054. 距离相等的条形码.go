package heap

import "container/heap"

// 利用最大堆
func rearrangeBarcodes(barcodes []int) []int {
	h := &MaxHeap{}
	count := make(map[int]int)
	heap.Init(h)

	for _, item := range barcodes {
		count[item]++
	}

	for k, v := range count {
		heap.Push(h, []int{k, v})
	}

	ans := make([]int, len(barcodes))

	for i := 0; i < len(barcodes); i++ {
		// 获取顶部元素
		top := (heap.Pop(h)).([]int)
		num, count := top[0], top[1]

		if i == 0 || ans[i-1] != num {
			ans[i] = num
		} else {
			// 获取第二多的数字
			second := (heap.Pop(h)).([]int)
			num, count = second[0], second[1]
			ans[i] = num

			heap.Push(h, []int{top[0], top[1]})
		}

		if count > 1 {
			heap.Push(h, []int{num, count - 1})
		}
	}

	return ans
}

type MaxHeap [][]int // 最大堆

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h MaxHeap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h *MaxHeap) Less(a, b int) bool {
	return (*h)[a][1] > (*h)[b][1]
}

func (h *MaxHeap) Push(x interface{}) {
	x1 := x.([]int)
	*h = append(*h, x1)
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}
