package binary_tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 利用中序遍历的性质，记录下前一个遍历到的结果
func isValidBST(root *TreeNode) bool {

	var pre = math.MinInt

	var f func(*TreeNode) bool
	f = func(tn *TreeNode) bool {
		if tn == nil {
			return true
		}

		// 因为这是中序遍历，所以有效的二叉搜索树，上一个遍历到的结果不会大于当前值
		if !f(tn.Left) || pre >= tn.Val {
			return false
		}

		pre = tn.Val

		return f(tn.Right)

	}
	return f(root)
}
