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

func reverseList(head *ListNode) *ListNode {

	var (
		current           = head
		prev    *ListNode = nil // 这样可以assign nil
	)

	for current != nil {
		// 存储下一个节点
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
