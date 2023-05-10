package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) (ans int) {

	var dfs func(*TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}

		left := dfs(node.Left) + 1
		right := dfs(node.Right) + 1

		ans = max(ans, left+right)
		// fmt.Println(node, left, right)
		return max(left, right)
	}

	dfs(root)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
