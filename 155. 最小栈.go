package leetcode_golang

import "math"

type MinStack struct {
	minStack []int // 辅助栈，存着最小元素的信息
	Stack    []int
}

func Constructor() MinStack {
	return MinStack{
		minStack: []int{
			math.MaxInt64,
		},
	}
}

func (s *MinStack) Push(val int) {
	s.Stack = append(s.Stack, val)
	min := s.minStack[len(s.minStack)-1]
	// 推入当前最小
	if val < min {
		s.minStack = append(s.minStack, val)
	} else {
		s.minStack = append(s.minStack, min)
	}
}

func (s *MinStack) Pop() {
	s.minStack = s.minStack[:len(s.minStack)-1]
	s.Stack = s.Stack[:len(s.Stack)-1]
}

func (s *MinStack) Top() int {
	return s.Stack[len(s.Stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
