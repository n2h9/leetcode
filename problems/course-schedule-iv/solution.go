package solution

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	al := make([][]int, n)
	for _, prerequisite := range prerequisites {
		al[prerequisite[0]] = append(al[prerequisite[0]], prerequisite[1])
	}
	// paths also is a visited array paths[i] == nil also means vertex us not visited yet
	// this is work because by condition grapth does not contain cycles
	// so we could not reach same vertex while DFS from it
	paths := make([]map[int]bool, n)
	answer := make([]bool, len(queries))
	for v := 0; v < n; v++ {
		if paths[v] == nil {
			paths[v] = calcAccessibleVertices(v, al, paths)
		}
	}
	for i, q := range queries {
		answer[i] = paths[q[0]] != nil && paths[q[0]][q[1]]
	}
	return answer
}

func calcAccessibleVertices(u int, al [][]int, paths []map[int]bool) map[int]bool {
	p := make(map[int]bool, 100)
	for _, v := range al[u] {
		p[v] = true
		if paths[v] == nil {
			paths[v] = calcAccessibleVertices(v, al, paths)
		}
		for k, val := range paths[v] {
			p[k] = val
		}
	}
	return p
}
