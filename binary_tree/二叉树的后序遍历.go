package binary_tree

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func postorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	val := make([]int, 0)

	prev := &TreeNode{}
	node := root

	for len(stack) != 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			// 【1,2,3】
			node = node.Left // 找到最左
		}

		// 此时没有最左，找到最后一个节点
		node = stack[len(stack)-1]
		// node = 3
		stack = stack[:len(stack)-1]
		// [1,2]

		// 也没有右
		// 或者这个节点的右边已经被遍历过了
		if node.Right == nil || node.Right == prev {
			// prev = node3
			prev = node
			val = append(val, node.Val)
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
			// 【1】
		}
	}
	return val
}

func postOrderRecur(head *TreeNode) (res []int) {
	if head == nil {
		return res
	}
	res = make([]int, 0)
	postOrderRecur(head.Right)
	postOrderRecur(head.Left)
	res = append(res, head.Val)
	return res
}
