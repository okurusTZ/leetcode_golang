package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 先入右子树，最后一个出队列的数据，就是答案
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := Queue{
		root,
	}

	var (
		node *TreeNode = nil
	)

	for !q.IsEmpty() {
		// 这里不需要再写一个循环了，直接出队就可以
		node = q.Pop()
		if node.Right != nil {
			q.Push(node.Right)
		}
		if node.Left != nil {
			q.Push(node.Left)
		}
	}

	return node.Val
}
