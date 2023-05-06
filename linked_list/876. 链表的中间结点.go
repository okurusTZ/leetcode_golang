package linked_list

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

// 快慢指针
func middleNode(head *ListNode) *ListNode {
	var (
		slow, fast = head, head
	)

	for fast != nil && fast.Next != nil {
		// 奇偶导致，如果一共奇数个，那循环到最后是fast.Next==nil
		// 一共是偶数个，那循环到最后是fast==nil

		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow

}
