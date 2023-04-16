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

func nextLargerNodes(head *ListNode) []int {

	reversedList, length := reverseList(head)

	ans := make([]int, length)
	stack := make([]int, 0)

	for reversedList != nil {

		length--

		for len(stack) > 0 && reversedList.Val >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			ans[length] = stack[len(stack)-1]
		}

		stack = append(stack, reversedList.Val)
		reversedList = reversedList.Next
	}

	return ans

}

func reverseList(head *ListNode) (*ListNode, int) {
	var (
		current           = head
		prev    *ListNode = nil
		length            = 0
	)

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
		length++
	}

	return prev, length
}
