package solution

// const BRACKET_OPEN = '('
// const BRACKET_CLOSE = ')'

// func removeInvalidParentheses(s string) []string {
// 	if isValid(s) || len(s) <= 0 {
// 		return []string{s}
// 	}

// 	resMap := make(map[string]bool)
// 	q := newQueue()

// 	s = preProcess0(s)
// 	s = preProcess1(s)
// 	process(s, 0, q)

// 	// 0 means no result has already found
// 	resultDepth := 0

// 	for !q.isEmpty() {
// 		variant := q.top().val
// 		q.pop()
// 		if resultDepth > 0 && variant.depth > resultDepth {
// 			break
// 		}
// 		if isValid(variant.s) {
// 			resultDepth = variant.depth
// 			resMap[variant.s] = true
// 		}
// 		if resultDepth == 0 {
// 			process(variant.s, variant.depth, q)
// 		}
// 	}

// 	res := make([]string, 0, len(resMap))
// 	for k := range resMap {
// 		res = append(res, k)
// 	}

// 	if len(res) <= 0 {
// 		res = append(res, deleteAllBrackets(s))
// 	}

// 	return res
// }

// // remove all clothing brackets from the start and all open bracckets from the end
// func preProcess0(s string) string {
// 	i := 0
// 	for ; i < len(s) && s[i] == BRACKET_CLOSE; i++ {
// 	}
// 	s = s[i:]

// 	i = len(s) - 1
// 	for ; i >= 0 && s[i] == BRACKET_OPEN; i-- {
// 		s = s[:i]
// 	}

// 	return s
// }

// // detect base on number of open bracket and close bracket maximal possible nesting level
// // and replace sequences which are longer than max nesting level with max nesting level
// // f.e. solving "((()((s((((()" is equal to solving ""(()((s(()"
// // because there is only two close brackets so we could close maximum two open brackets
// func preProcess1(s string) string {
// 	countOpen := 0
// 	countClose := 0
// 	for i := range s {
// 		switch {
// 		case s[i] == BRACKET_OPEN:
// 			countOpen++
// 		case s[i] == BRACKET_CLOSE:
// 			countClose++
// 		}
// 	}
// 	if countClose == countOpen {
// 		return s
// 	}
// 	bracketSequence := make([]byte, 0)
// 	bracket := byte(BRACKET_OPEN)
// 	minCount := countClose
// 	if countClose > countOpen {
// 		bracket = BRACKET_CLOSE
// 		minCount = countOpen
// 	}
// 	for i := minCount; i > 0; i-- {
// 		bracketSequence = append(bracketSequence, bracket)
// 	}
// 	newS, buff := make([]byte, 0), make([]byte, 0)
// 	for i := range s {
// 		if s[i] == bracket {
// 			buff = append(buff, s[i])
// 			continue
// 		}
// 		if len(buff) > minCount {
// 			newS = append(newS, bracketSequence...)
// 		} else {
// 			newS = append(newS, buff...)
// 		}
// 		buff = make([]byte, 0)
// 		newS = append(newS, s[i])
// 	}
// 	newS = append(newS, buff...)
// 	return string(newS)
// }

// func process(s string, depth int, q *queue) {
// 	newDepth := depth + 1
// 	variants := make([]string, 0, len(s))
// 	b := []byte(s)
// 	for i := range s {
// 		if !isBracket(s[i]) {
// 			continue
// 		}
// 		variants = append(
// 			variants,
// 			string(
// 				append(
// 					append(
// 						[]byte{},
// 						b[:i]...,
// 					),
// 					b[i+1:]...,
// 				),
// 			),
// 		)
// 	}
// 	if len(variants) < 2 {
// 		return
// 	}
// 	for _, variant := range variants {
// 		q.push(newNode(&value{s: variant, depth: newDepth}))
// 	}
// }

// func isBracket(ch byte) bool {
// 	return ch == BRACKET_OPEN || ch == BRACKET_CLOSE
// }

// func isValid(s string) bool {
// 	val := 0
// 	for i := range s {
// 		switch {
// 		case s[i] == BRACKET_OPEN:
// 			val++
// 		case s[i] == BRACKET_CLOSE:
// 			val--
// 			if val < 0 {
// 				return false
// 			}
// 		}
// 	}
// 	return val == 0
// }

// func deleteAllBrackets(s string) string {
// 	b := []byte(s)
// 	numOfBrackets := 0
// 	for i := range s {
// 		if isBracket(s[i]) {
// 			b2 := b[i+1-numOfBrackets:]
// 			b = b[:i-numOfBrackets]
// 			b = append(b, b2...)
// 			numOfBrackets++
// 		}
// 	}
// 	return string(b)
// }

// type value struct {
// 	s     string
// 	depth int
// }

// type node struct {
// 	val  *value
// 	next *node
// }

// func newNode(val *value) *node {
// 	return &node{val: val}
// }

// type queue struct {
// 	head *node
// 	tail *node
// }

// func newQueue() *queue {
// 	return &queue{}
// }

// func (q *queue) isEmpty() bool {
// 	return q.head == nil
// }

// func (q *queue) top() *node {
// 	return q.head
// }

// func (q *queue) pop() {
// 	n := q.head
// 	q.head = q.head.next
// 	n.next = nil

// 	if q.head == nil {
// 		q.tail = nil
// 	}
// }

// func (q *queue) push(n *node) {
// 	if q.head == nil {
// 		q.head = n
// 		q.tail = n
// 		return
// 	}

// 	q.tail.next = n
// 	q.tail = n
// }
