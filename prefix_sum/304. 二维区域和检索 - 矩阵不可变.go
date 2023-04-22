package prefix_sum

type NumMatrix struct {
	PreSum [][]int // 每一行的preSum
}

func Constructor(matrix [][]int) NumMatrix {
	preSum := make([][]int, len(matrix))
	for i := 1; i < len(matrix)+1; i++ {
		preSum[i] = make([]int, len(matrix[0])+1)
		for j := 1; j < len(matrix[0])+1; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{
		PreSum: preSum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.PreSum[row2+1][col2+1] - this.PreSum[row1][col1+1] - this.PreSum[row2+1][col1] + this.PreSum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

type NumMatrix struct {
	PreSum [][]int // 每一行的preSum
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	m, n := len(matrix), len(matrix[0])
	preSum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		preSum[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{
		PreSum: preSum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.PreSum[row2+1][col2+1] - this.PreSum[row1][col2+1] - this.PreSum[row2+1][col1] + this.PreSum[row1][col1]
}
