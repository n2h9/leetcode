package solution

// import (
// 	"strconv"
// 	"strings"
// )

// type state byte

// const STATE_START state = 's'
// const STATE_END state = 'e'

// func exclusiveTime(n int, logs []string) []int {
// 	result := make([]int, n)
// 	st := newStack()
// 	for _, log := range logs {
// 		v := parseLog(log)
// 		if v.state == STATE_START {
// 			st.push(newNode(v))
// 			continue
// 		}
// 		// stack is not empty by problem conditions on this step
// 		if v.state != STATE_END || st.isEmpty() {
// 			continue
// 		}

// 		// also current.id must be the same function as v.id by problem condition (no check)
// 		current := st.top().val
// 		st.pop()
// 		if v.id == current.id {
// 			// increase total function execution time
// 			result[v.id] += v.timestamp - current.timestamp + 1
// 		}
// 		if !st.isEmpty() {
// 			// if stack is not empty, than function which is on top of it will continue execution
// 			previous := st.top().val
// 			// increase previous total function execution time
// 			result[previous.id] += current.timestamp - previous.timestamp
// 			// and do not pop from stack, just set new start time right after current is finished
// 			// this is a little bit tricky, if next log value is start right after current is finished
// 			// it means that previous.timestamp will be equal to next.timestamp
// 			// but when total execution time will be counted it will be next.timestamp-previous.timestamp and it will be 0 which is ok
// 			previous.timestamp = v.timestamp + 1
// 		}
// 	}
// 	return result
// }

// func parseLog(log string) *value {
// 	parts := strings.Split(log, ":")
// 	id, _ := strconv.Atoi(parts[0])
// 	timestamp, _ := strconv.Atoi(parts[2])
// 	var s state
// 	if parts[1] == "start" {
// 		s = STATE_START
// 	} else if parts[1] == "end" {
// 		s = STATE_END
// 	}
// 	return &value{
// 		id:        id,
// 		state:     s,
// 		timestamp: timestamp,
// 	}
// }

// type value struct {
// 	id        int
// 	state     state
// 	timestamp int
// }

// type node struct {
// 	val  *value
// 	next *node
// }

// func newNode(val *value) *node {
// 	return &node{val: val}
// }

// type stack struct {
// 	head *node
// }

// func newStack() *stack {
// 	return &stack{}
// }

// func (s *stack) isEmpty() bool {
// 	return s.head == nil
// }

// func (s *stack) top() *node {
// 	return s.head
// }

// // stack must not be empty
// func (s *stack) pop() {
// 	head := s.head
// 	s.head = s.head.next
// 	head.next = nil
// }

// func (s *stack) push(n *node) {
// 	n.next = s.head
// 	s.head = n
// }
