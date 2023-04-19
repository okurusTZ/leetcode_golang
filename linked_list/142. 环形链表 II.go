package main

// 找到链表成环的部分
func detectCycle(head *ListNode) *ListNode {

	// 对于快慢指针来说
	// 第一次相遇的时候，慢指针一定还没有走完第一次环
	// 假设相遇时，慢指针在环内走了b长度，剩余环长为c
	// b + c = 总环长，再假设head到成环的节点的距离为a
	// 举一个极端的例子，慢指针进环时，刚好在快指针后面，这样相对速度下，快指针下一次相遇需要走(b+c-1)步
	// 这个情况下，慢指针还没有走完一整圈
	// 慢指针移动的距离 a+b
	// 快指针移动的距离 a+b+k(b+c) -> 对于快指针来说，可能已经走完n圈了
	// 因为速度是2倍
	// -> 2(a+b) = a+b+k(b+c)
	// -> a+b-b-c = (k-1)(b+c)
	// -> a-c = (k-1)(b+c)
	// 意味着，这时候从head出发，和慢指针一样的速度，一定能遇到走完(k-1)圈的慢指针

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			// 头节点开始出发
			for slow != head {
				// 这里忽略掉了初始化的时候就相遇的情况
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}

	return nil

}

// e.g.
// 1 -> 2
// ^    ｜
// ---——-
// 2 1
// 1 1 -> 第一次相遇
//
