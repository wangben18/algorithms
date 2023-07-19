package problems

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int) // 统计数字出现的频率
	for _, v := range arr {
		count[v] += 1
	}
	fre := make(map[int]struct{}) // 看相同频率是否重复出现
	for _, v := range count {
		if _, ok := fre[v]; ok {
			return false
		}
		fre[v] = struct{}{}
	}
	return true
}
