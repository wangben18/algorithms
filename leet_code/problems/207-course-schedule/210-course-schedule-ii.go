package courseschedule

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, info := range prerequisites {
		indeg[info[0]]++
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	var q []int
	for i := range indeg {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	var result []int
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		result = append(result, cur)
		for _, next := range edges[cur] {
			indeg[next]--
			if indeg[next] == 0 {
				q = append(q, next)
			}
		}
	}
	if len(result) != numCourses {
		return []int{}
	}
	return result
}
