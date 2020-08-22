package solution

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 1 {
		return true
	}
	g := newGraph(numCourses)
	for _, deps := range prerequisites {
		g.addEdge(deps[0], deps[1])
	}
	return !containCycle(g)
}

type graph struct {
	adjList [][]int
}

func newGraph(V int) *graph {
	return &graph{
		adjList: make([][]int, V),
	}
}

func (g *graph) addEdge(v, u int) {
	if g.adjList[v] == nil {
		g.adjList[v] = make([]int, 0)
	}
	g.adjList[v] = append(g.adjList[v], u)
}

func containCycle(g *graph) bool {
	visited := make([]bool, len(g.adjList))
	recStack := make([]bool, len(g.adjList))
	for v := 0; v < len(g.adjList); v++ {
		if !visited[v] {
			if recCheckCycle(g, v, visited, recStack) {
				return true
			}
		}
	}

	return false
}

func recCheckCycle(g *graph, v int, visited, recStack []bool) bool {
	if recStack[v] {
		return true
	}
	recStack[v] = true
	for _, u := range g.adjList[v] {
		if recCheckCycle(g, u, visited, recStack) {
			return true
		}
	}
	recStack[v] = false
	visited[v] = true
	return false
}
