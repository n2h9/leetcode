package solution

func findRedundantDirectedConnection(edges [][]int) []int {
	return newProcessor(edges).result()
}

type processor struct {
	edges                 [][]int
	verticesParentsNum    []int
	parentsNumVerticesNum map[int]int
}

func newProcessor(edges [][]int) *processor {
	// because vertices starts from 1 to N
	verticesParentsNum := make([]int, len(edges)+1)
	for _, edge := range edges {
		verticesParentsNum[edge[1]]++
	}
	parentsNumVerticesNum := make(map[int]int)
	for _, parentsNum := range verticesParentsNum {
		parentsNumVerticesNum[parentsNum]++
	}
	// because vertices starts from 1 to N
	// so verticesParentsNum[0] = 0 and 0 vertice is not exists
	parentsNumVerticesNum[0]--

	return &processor{
		edges:                 edges,
		verticesParentsNum:    verticesParentsNum,
		parentsNumVerticesNum: parentsNumVerticesNum,
	}
}

func (p *processor) result() []int {
	resultEdge := []int{}
	for i := len(p.edges) - 1; i >= 0; i-- {
		p.removeEdge(p.edges[i])
		if p.isTree() {
			resultEdge = p.edges[i]
			break
		}
		p.returnEdge(p.edges[i])
	}
	return resultEdge
}

func (p *processor) removeEdge(edge []int) {
	p.parentsNumVerticesNum[p.verticesParentsNum[edge[1]]]--
	p.verticesParentsNum[edge[1]]--
	// if p.parentsNumVerticesNum[pnvn] == 0 {
	// 	delete(p.parentsNumVerticesNum, pnvn)
	// }
	p.parentsNumVerticesNum[p.verticesParentsNum[edge[1]]]++
}

func (p *processor) returnEdge(edge []int) {
	p.parentsNumVerticesNum[p.verticesParentsNum[edge[1]]]--
	p.verticesParentsNum[edge[1]]++
	// if p.parentsNumVerticesNum[pnvn] == 0 {
	// 	delete(p.parentsNumVerticesNum, pnvn)
	// }
	p.parentsNumVerticesNum[p.verticesParentsNum[edge[1]]]++
}

func (p *processor) isTree() bool {
	return p.parentsNumVerticesNum[0] == 1 &&
		p.parentsNumVerticesNum[1] == len(p.edges)-1
}
