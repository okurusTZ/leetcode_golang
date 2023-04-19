package binary_tree

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

type Queue []any

func (q *Queue) Push(v any) {
	*q = append(*q, v)
}

func (q *Queue) Pop() any {
	head := (*q)[0]
	if len(*q) >= 1 {
		*q = (*q)[1:]
	} else {
		*q = Queue{}
	}
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Length() int {
	return len(*q)
}
