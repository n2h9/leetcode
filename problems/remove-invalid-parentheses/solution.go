package solution

const BRACKET_OPEN = '('
const BRACKET_CLOSE = ')'

func removeInvalidParentheses(s string) []string {
	return newProcessor(s).run().result()
}

type processor struct {
	s             string
	results       map[string]bool
	minRemovedPar int
}

func newProcessor(s string) *processor {
	return &processor{
		s:             s,
		results:       make(map[string]bool),
		minRemovedPar: 1<<31 - 1,
	}
}

func (p *processor) run() *processor {
	p._rec(0, 0, 0, 0, "")
	return p
}

func (p *processor) result() []string {
	res := []string{}
	for v := range p.results {
		res = append(res, v)
	}
	return res
}

func (p *processor) _rec(index, countLeft, countRight, countRemoval int, exp string) {
	if index == len(p.s) {
		if countLeft != countRight || countRemoval > p.minRemovedPar {
			return
		}
		if countRemoval < p.minRemovedPar {
			p.minRemovedPar = countRemoval
			p.results = make(map[string]bool)
		}
		p.results[exp] = true
		return
	}
	if p.s[index] != BRACKET_OPEN && p.s[index] != BRACKET_CLOSE {
		p._rec(index+1, countLeft, countRight, countRemoval, exp+string(p.s[index]))
		return
	}

	// exclude char case
	p._rec(index+1, countLeft, countRight, countRemoval+1, exp)

	// include char case for open par
	if p.s[index] == BRACKET_OPEN {
		p._rec(index+1, countLeft+1, countRight, countRemoval, exp+string(p.s[index]))
		return
	}

	// include char case for close bracket (only if number of left par > number of right par)
	if countLeft > countRight {
		p._rec(index+1, countLeft, countRight+1, countRemoval, exp+string(p.s[index]))
	}
}
