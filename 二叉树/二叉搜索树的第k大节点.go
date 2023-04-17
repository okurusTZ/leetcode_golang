package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthLargest(root *TreeNode, k int) int {
	val := Order(root)
	return val[k-1]
}

// 逆序中序遍历
func Order(root *TreeNode) []int {
	val := make([]int, 0)
	node := root
	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Right
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			val = append(val, node.Val)
			node = node.Left
		}
	}
	return val
}
