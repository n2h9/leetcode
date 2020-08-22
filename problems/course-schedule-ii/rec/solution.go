package solution

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses <= 0 {
		return []int{}
	}

	adjList := make([][]int, numCourses)
	inEdgeCount := make([]int, numCourses)
	for _, dep := range prerequisites {
		adjList[dep[1]] = append(adjList[dep[1]], dep[0])
		inEdgeCount[dep[0]]++
	}

	resOrder := make([]int, 0, numCourses)
	visited := make([]bool, numCourses)
	for v, c := range inEdgeCount {
		if !visited[v] && c == 0 {
			resOrder = findOrderRec(v, visited, adjList, inEdgeCount, resOrder)
		}
	}

	// imposible, graph has cycles
	if len(resOrder) != numCourses {
		return []int{}
	}

	return resOrder
}

func findOrderRec(v int, visited []bool, adjList [][]int, inEdgeCount []int, resOrder []int) []int {
	visited[v] = true
	resOrder = append(resOrder, v)
	for _, u := range adjList[v] {
		if !visited[u] {
			inEdgeCount[u]--
			if inEdgeCount[u] == 0 {
				resOrder = findOrderRec(u, visited, adjList, inEdgeCount, resOrder)
			}
		}
	}
	return resOrder
}
