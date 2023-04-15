package main

// 思路： 找到中点之后反转后半部分，走完前半部value都相同就可以
// why -> 如果后半部分比前半部分长，多出来的一定是中点（奇数个），可以忽略
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	beforeMid, mid := midNode(head)
	// 断开
	beforeMid.Next = nil
	newHead := reverseList(mid)

	// fmt.Println(newHead.Val, beforeMid.Val, mid.Val)

	for head != nil {
		// fmt.Println(head.Val, newHead.Val)
		if head.Val != newHead.Val {
			return false
		}
		head = head.Next
		newHead = newHead.Next
	}
	return true
}

func midNode(head *ListNode) (*ListNode, *ListNode) {
	slow, fast, prev := head, head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return prev, slow
}

func reverseList(head *ListNode) *ListNode {
	var (
		cur            = head
		prev *ListNode = nil
	)

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
