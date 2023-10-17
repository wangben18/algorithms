package problems

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		tmp := []byte(str)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		m[string(tmp)] = append(m[string(tmp)], str)
	}
	ret := make([][]string, 0, len(m))
	for _, val := range m {
		ret = append(ret, val)
	}
	return ret
}
