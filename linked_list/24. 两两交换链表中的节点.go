package linked_list

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next // 交换后新的head
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}
