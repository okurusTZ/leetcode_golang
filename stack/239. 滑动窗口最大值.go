package stack

import (
	"container/heap"
	"sort"
)

var a []int

func maxSlidingWindow2(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

type hp struct{ sort.IntSlice }

func (h *hp) Pop() any {
	a := h.IntSlice // 获取这个堆数据
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) Push(value any) {
	h.IntSlice = append(h.IntSlice, value.(int))
}
func (h *hp) Less(i, j int) bool {
	return h.IntSlice[i] < h.IntSlice[j]
}

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		// 保证顶上这个永远是最大值
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		// 如果最大值的idx
		// fmt.Println(q[0], i-k, nums[q[0]])
		// 左边界大于最大值的idx了，说明要被丢掉了
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
