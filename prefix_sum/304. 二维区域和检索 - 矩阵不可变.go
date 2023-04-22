package prefix_sum

type NumMatrix struct {
	Matrix [][]int
	PreSum [][]int // 每一行的preSum
}

func Constructor(matrix [][]int) NumMatrix {
	preSum := make([][]int, len(matrix))
	for i, row := range matrix {
		preSum[i] = make([]int, len(row))
		for j := range row {
			if j == 0 {
				preSum[i][j] = row[j]
			} else {
				preSum[i][j] += row[j] + preSum[i][j-1]
			}
		}
	}

	return NumMatrix{
		Matrix: matrix,
		PreSum: preSum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := 0
	for i := row1; i <= row2; i++ {
		ans += this.PreSum[i][col2] - this.PreSum[i][col1] + this.Matrix[i][col1]
	}
	return ans
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
