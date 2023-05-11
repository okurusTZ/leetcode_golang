package matrix

func setZeroes(matrix [][]int) {
	var (
		row0, col0 = false, false // 第一行第一列是否存在0
		rl, cl     = len(matrix), len(matrix[0])
	)

	for i := 0; i < rl; i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
	}

	for i := 0; i < cl; i++ {
		if matrix[0][i] == 0 {
			row0 = true
		}
	}

	// 使用第一行/列记录
	for i := 1; i < rl; i++ {
		for j := 1; j < cl; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 开始置为0
	for i := 1; i < rl; i++ {
		for j := 1; j < cl; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if row0 {
		for j := 0; j < cl; j++ {
			matrix[0][j] = 0
		}
	}

	if col0 {
		for j := 0; j < rl; j++ {
			matrix[j][0] = 0
		}
	}
}
