package linked_list

func reorderList(head *ListNode) {
	// before l1 -> l2 -> l3 .... -> ln-1 -> ln
	// after l1 -> ln -> l2 -> ln-1 ...

	// tail.Next = current.Next.Next
	// current.Next -> tail

	mid := middleNode(head)
	head2 := reverseList(mid) // -> 右半段的头节点，也就是需要插入左半段的内容

	// 退出条件 head2 -> mid ok
	// head2.Next == nil
	for head2 != mid {
		// 提前记录
		next2 := head2.Next
		next := head.Next

		head2.Next = head.Next
		head.Next = head2

		head2 = next2
		head = next
	}
}

// 复用876 -> 获取中间节点
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

// 翻转链表
func reverseList(head *ListNode) *ListNode {
	var (
		prev    *ListNode = nil
		current           = head
	)

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
