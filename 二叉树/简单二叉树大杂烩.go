package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 是否对称的树 101
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left, right := root.Left, root.Right
	return isSymmetricTree(left, right)
}

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	return p.Val == q.Val && isSymmetricTree(p.Left, q.Right) && isSymmetricTree(p.Right, q.Left)
}

// 是否是平衡的树
func isBalanced(root *TreeNode) bool {
	var getHeight func(*TreeNode) int

	if root == nil {
		return true
	}

	getHeight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left_h := getHeight(root.Left)
		if left_h == -1 {
			return -1
		}
		right_h := getHeight(root.Right)
		if right_h == -1 || abs(right_h, left_h) > 1 {
			return -1
		}
		return max(left_h, right_h) + 1
	}

	return getHeight(root) != -1
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {

	var f func(*TreeNode) *TreeNode

	f = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		// 存一下，不然被覆盖掉了
		tmp := node.Left
		node.Left = f(node.Right)
		node.Right = f(tmp)

		return node
	}

	return f(root)
}
