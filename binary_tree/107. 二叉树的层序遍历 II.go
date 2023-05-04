package binary_tree

func levelOrderBottom(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	var (
		queue = make([]*TreeNode, 0)
		depth = 0
		ans   = make([][]int, 0)
	)

	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue) // 这里必须要存下来，不然queue是变长的，遍历有误
		for i := 0; i < n; i++ {
			// 取出头
			head := queue[0]
			if len(ans) < depth+1 {
				ans = append(ans, []int{})
			}
			ans[depth] = append(ans[depth], head.Val)
			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
			// 移除头部
			queue = queue[1:]
		}
		depth++
	}

	// 倒叙
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}
