package binary_tree

// 如何判断是否要记录当前节点，因为一定是先看见右子树，再看到更深的左子树
// 记录一个深度长度的数组，如果遍历过程中，深度上已经有数值，不记录
// 优先遍历右子树
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)

	var f func(*TreeNode, int)

	f = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(ans) {
			ans = append(ans, node.Val)
		}
		f(node.Right, depth+1)
		f(node.Left, depth+1)
	}

	f(root, 0)
	return ans
}
