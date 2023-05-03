package graph

// bfs
func orangesRotting(grid [][]int) int {
	ans := 0

	directions := [][]int{
		{0, -1}, {0, 1}, {1, 0}, {-1, 0},
	}

	type orange struct {
		axis  []int
		depth int
	}

	queue := make([]orange, 0) // 标记腐烂的橘子

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, orange{
					axis:  []int{i, j},
					depth: 0,
				})
			}
		}
	}

	for len(queue) > 0 {
		x, y, depth := queue[0].axis[0], queue[0].axis[1], queue[0].depth

		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if newX < 0 || newX > len(grid)-1 || newY < 0 || newY > len(grid[0])-1 || grid[newX][newY] != 1 {
				continue
			}
			queue = append(queue, orange{
				axis:  []int{newX, newY},
				depth: depth + 1,
			})
			if depth+1 > ans {
				ans = depth + 1
			}
			grid[newX][newY] = 2 // 烂掉了
		}
		queue = queue[1:]
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return ans
}
