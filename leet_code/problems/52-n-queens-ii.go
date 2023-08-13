package problems

func totalNQueens(n int) int {
	return len(solveNQueens(n))
}
