package linked_list

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	reversedL1, reversedL2 := l1, l2

	var (
		ans = &ListNode{
			Next: nil,
		}

		carry = 0
		sum   = 0

		dummy = ans
	)

	for reversedL1 != nil || reversedL2 != nil {

		if reversedL2 == nil {
			reversedL2 = &ListNode{
				Val: 0,
			}
		}

		if reversedL1 == nil {
			reversedL1 = &ListNode{
				Val: 0,
			}
		}

		sum = (reversedL1.Val + reversedL2.Val + carry) % 10
		carry = (reversedL1.Val + reversedL2.Val + carry) / 10

		if reversedL1 != nil {
			reversedL1 = reversedL1.Next
		}

		if reversedL2 != nil {
			reversedL2 = reversedL2.Next
		}

		dummy.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}
		dummy = dummy.Next
	}

	if carry != 0 {
		dummy.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return ans.Next
}
