package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	fisrt, second := head, dummy

	for i := 0; i < n; i++ {
		fisrt = fisrt.Next // 前进n步
	}

	for fisrt != nil {
		fisrt = fisrt.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next // 为什么不直接返回head? 需要考虑到只有一个节点的情况

}
