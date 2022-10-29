package problems

import "sort"

func findItinerary(tickets [][]string) []string {
	sort.Sort(Tickets(tickets))
	result := make([]string, 0)
	var backTracking func(tickets [][]string, start string, used [300]bool, path []string)
	backTracking = func(tickets [][]string, start string, used [300]bool, path []string) {
		if len(result) > 0 {
			return
		}
		if len(path)-1 == len(tickets) {
			result = append([]string{}, path...)
			return
		}
		for i := 0; i < len(tickets); i++ {
			if used[i] || tickets[i][0] != start {
				continue
			}
			used[i] = true
			backTracking(tickets, tickets[i][1], used, append(path, tickets[i][1]))
			used[i] = false
		}
	}
	backTracking(tickets, "JFK", [300]bool{}, []string{"JFK"})
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
