package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 假设链表A长度为a 链表B长度为b
	// 共同长度为c，那么假设存在共同节点，A到共同节点的个数为a-c,B为b-c
	// 双指针，指针A先走完链表A，再走链表B，走到共同节点时，走了a+(b-c)
	// 指针B先走链表B再走链表A，走到共同节点时，走了b+(a-c)
	// 所以，此时指针重合

	var (
		pointerA = headA
		pointerB = headB
	)

	for pointerA != pointerB {
		// 走完了没
		if pointerA != nil {
			pointerA = pointerA.Next
		} else {
			// A走完了
			pointerA = headB
		}
		if pointerB != nil {
			pointerB = pointerB.Next
		} else {
			// B走完了
			pointerB = headA
		}
	}

	return pointerA
}
