package dfs

func numIslands(grid [][]byte) int {

	var ans = 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				dye(i, j, &grid)
			}
		}
	}
	return ans
}

func dye(x, y int, grid *[][]byte) {
	// (*grid)[x][y] != '1' 这个判断条件！！要遍历的必须是还没遍历过的联通节点
	if x < 0 || x >= len((*grid)) || y < 0 || y >= len((*grid)[0]) || (*grid)[x][y] != '1' {
		return
	}
	(*grid)[x][y] = '2'
	dye(x-1, y, grid)
	dye(x+1, y, grid)
	dye(x, y-1, grid)
	dye(x, y+1, grid)
}
