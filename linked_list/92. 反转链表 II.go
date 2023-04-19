package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 翻转中间一段，可以看作和I是一样的
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Next: head,
	} // 哨兵节点

	p0 := dummy

	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}

	// p0 开始翻转的前一个节点

	var (
		prev    *ListNode = nil
		current           = p0.Next
	)

	for i := 0; i < (right - left + 1); i++ {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	// p0.Next 开始翻转的头节点
	p0.Next.Next = current
	p0.Next = prev

	return dummy.Next
}
