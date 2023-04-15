package main

// 快慢指针
// 慢指针每次走一步，如果有环，一定会进去
// 快指针相对于慢指针每次向前一步，如果有环，一定会相遇
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
