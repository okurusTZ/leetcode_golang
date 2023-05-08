package bfs

func minPushBox(grid [][]byte) int {
	var sx, sy, bx, by = 0, 0, 0, 0
	var width, height = len(grid), len(grid[0])

	// 记录下人和箱子的位置
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if grid[i][j] == 'S' {
				sx = i
				sy = j
			}
			if grid[i][j] == 'B' {
				bx = i
				by = j
			}
		}
	}

	// 			   [["#","#","#","#","#","#"],
	// 			    ["#","T","#","#","#","#"],
	//              ["#",".",".","B",".","#"],
	//              ["#",".","#","#",".","#"],
	//              ["#",".",".",".","S","#"],
	//              ["#","#","#","#","#","#"]]

	var dp = make([][]int, width*height)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, height*width)
	}

	var dir = []int{0, -1, 0, 1, 0} // 移动方向

	for i := 0; i < width*height; i++ {
		for j := 0; j < height*width; j++ {
			dp[i][j] = 0x3f3f3f3f // 表示infinity
		}
	}

	// 表示当前的位置是否有效
	canPass := func(x, y int) bool {
		return x >= 0 && x < width && y >= 0 && y < height && grid[x][y] != '#'
	}

	dp[height*sx+sy][height*bx+by] = 0

	// 初始位置
	q := [][2]int{{sx*height + sy, bx*height + by}}

	for len(q) > 0 {
		// 这里做的是把人推到箱子面前
		q1 := [][2]int{}
		for len(q) > 0 {
			// 这里开始做的是推箱子
			s1, b1 := q[0][0], q[0][1]
			q = q[1:]

			sx1, bx1, sy1, by1 := s1/height, b1/height, s1%height, b1%height

			if grid[bx1][by1] == 'T' {
				return dp[s1][b1]
			}

			for i := 0; i < 4; i++ {
				// 玩家移动到任意一个方向
				sx2, sy2 := sx1+dir[i], sy1+dir[i+1]
				if !canPass(sx2, sy2) { // 位置不合理
					continue
				}
				s2 := sx2*height + sy2

				// 玩家与箱子重合，给箱子一个力，箱子移动了
				if bx1 == sx2 && by1 == sy2 {
					bx2, by2 := bx1+dir[i], by1+dir[i+1]
					b2 := bx2*height + by2
					if !canPass(bx2, by2) {
						continue
					}
					// fmt.Println("推到箱子了", sx2, sy2)
					// 记录下当前的状态，是上一次多走一步
					dp[s2][b2] = dp[s1][b1] + 1
					//
					q1 = append(q1, [2]int{s2, b2})
				} else {
					// 玩家和箱子没有重合
					if dp[s2][b1] <= dp[s1][b1] { // 状态已访问
						continue
					}
					// 这里没有推动过，所以无所谓
					dp[s2][b1] = dp[s1][b1]
					q = append(q, [2]int{s2, b1})
				}
			}
		}
		q = q1
	}
	return -1
}
