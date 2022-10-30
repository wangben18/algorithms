package problems

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	isValid := func(row, col int, board []string) bool {
		// 垂直往上查
		for r := row - 1; r >= 0; r-- {
			if board[r][col] == 'Q' {
				return false
			}
		}
		// 45 度角查
		for r, c := row-1, col+1; r >= 0 && c < n; r, c = r-1, c+1 {
			if board[r][c] == 'Q' {
				return false
			}
		}
		// 135 度角查
		for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
			if board[r][c] == 'Q' {
				return false
			}
		}
		return true
	}
	var backTracking func(n, row int, board []string)
	backTracking = func(n, row int, board []string) {
		if len(board) == n {
			result = append(result, append([]string{}, board...))
			return
		}
		for i := 0; i < n; i++ {
			if !isValid(row, i, board) {
				continue
			}
			line := make([]byte, n)
			for j := 0; j < n; j++ {
				if j == i {
					line[j] = 'Q'
				} else {
					line[j] = '.'
				}
			}
			backTracking(n, row+1, append(board, string(line)))
		}
	}

	backTracking(n, 0, nil)
	return result
}
