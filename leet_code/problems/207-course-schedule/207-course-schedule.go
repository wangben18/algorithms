package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, info := range prerequisites {
		indeg[info[0]]++
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	var q []int
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	count := 0
	for len(q) > 0 {
		currentCourse := q[0]
		q = q[1:]
		count++
		for _, nextCourse := range edges[currentCourse] {
			indeg[nextCourse]--
			if indeg[nextCourse] == 0 {
				q = append(q, nextCourse)
			}
		}
	}
	return count == numCourses
}
