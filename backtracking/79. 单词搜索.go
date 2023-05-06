package backtracking

func exist(board [][]byte, word string) bool {

	var f func(int, int, int) bool
	var visited = make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	f = func(i int, x, y int) bool {
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || visited[x][y] || board[x][y] != word[i] {
			return false
		}

		visited[x][y] = true

		if i == len(word)-1 {
			return true
		}

		if f(i+1, x-1, y) {
			return true
		}
		if f(i+1, x+1, y) {
			return true
		}
		if f(i+1, x, y-1) {
			return true
		}
		if f(i+1, x, y+1) {
			return true
		}

		visited[x][y] = false
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if f(0, i, j) {
				return true
			}
		}
	}
	return false
}
