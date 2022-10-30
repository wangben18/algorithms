package problems

func solveSudoku(board [][]byte) {
	isValid := func(row, col int, num byte) bool {
		// 检查同行
		for c := 0; c < len(board); c++ {
			if board[row][c] == num {
				return false
			}
		}
		// 检查同列
		for r := 0; r < len(board); r++ {
			if board[r][col] == num {
				return false
			}
		}
		// 检查同 3*3 宫
		for r := 0; r < len(board); r++ {
			for c := 0; c < len(board); c++ {
				if r/3 == row/3 && c/3 == col/3 && board[r][c] == num {
					return false
				}
			}
		}
		return true
	}
	var backTracking func() bool
	backTracking = func() bool {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				if board[i][j] != '.' {
					continue
				}
				for num := '1'; num <= '9'; num++ {
					if !isValid(i, j, byte(num)) {
						continue
					}
					board[i][j] = byte(num)
					if backTracking() {
						return true
					}
					board[i][j] = '.'
				}
				return false
			}
		}
		return true
	}
	backTracking()
}
