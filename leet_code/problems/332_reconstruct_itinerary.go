package problems

import "sort"

func findItinerary(tickets [][]string) []string {
	sort.Sort(Tickets(tickets)) // 先排序，保证先获得的结果排序更靠前
	result := make([]string, 0)
	var backTracking func(tickets [][]string, start string, used []bool, path []string) // start: 当前要求的出发地，path: 当前路径
	backTracking = func(tickets [][]string, start string, used []bool, path []string) {
		if len(result) > 0 { // 因为只需要一个结果，剪枝
			return
		}
		if len(path)-1 == len(tickets) { // 路径已达期望值，保存结果
			result = append([]string{}, path...)
			return
		}
		for i := 0; i < len(tickets); i++ {
			if used[i] || tickets[i][0] != start { // 当前票已经用过 或者 当前票出发地不符合要求，下一位
				continue
			}
			used[i] = true // 标记当前票已用过
			backTracking(tickets, tickets[i][1], used, append(path, tickets[i][1]))
			used[i] = false // 回溯，月光宝盒时光倒流～
		}
	}
	used := make([]bool, len(tickets))
	backTracking(tickets, "JFK", used, []string{"JFK"}) // 从 JFK 出发
	return result
}

type Tickets [][]string

func (a Tickets) Len() int      { return len(a) }
func (a Tickets) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Tickets) Less(i, j int) bool {
	if a[i][0] == a[j][0] {
		return a[i][1] < a[j][1]
	}
	return a[i][0] < a[j][0]
}
