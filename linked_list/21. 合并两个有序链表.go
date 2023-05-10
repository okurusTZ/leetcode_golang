package linked_list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var dummy = &ListNode{
		Next: nil,
	}

	var current = dummy

	for list1 != nil && list2 != nil {
		h1 := list1.Val
		h2 := list2.Val

		if h1 < h2 {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}
	return dummy.Next
}
