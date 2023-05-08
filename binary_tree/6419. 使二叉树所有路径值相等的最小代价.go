package binary_tree

func minIncrements(n int, cost []int) int {
	// 对于叶子节点，比较两个叶子节点的大小，把小的补上就可以，因为上层路径都是相同的
	// 对于非叶子节点，默认下面已经补得相同

	var ans = 0

	// 这里下标从1开始
	for i := n/2 - 1; i >= 0; i-- {
		left, right := cost[i*2+1], cost[i*2+2]
		if left >= right {
			left, right = right, left
		}
		ans += right - left // 把子节点变成一样的大小
		cost[i] += right    // 这个点的路径和，就是加上right子节点的数值，因为默认right是大的一边
		// fmt.Println("访问第", i, "个节点", left, right, cost[i])
	}

	return ans
}
