package problems

func climbStairs(n int) int {
	return climbStairsWithMemo(n, make(map[int]int))
}

func climbStairsWithMemo(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	if memoNum, ok := memo[n]; ok {
		return memoNum
	}
	num := climbStairsWithMemo(n-1, memo) + climbStairsWithMemo(n-2, memo)
	memo[n] = num
	return num
}
