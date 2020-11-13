package solution

import (
	"strconv"
	"strings"
)

type state byte

const STATE_START state = 's'
const STATE_END state = 'e'

func exclusiveTime(n int, logs []string) []int {
	result := make([]int, n)
	st := newStack()
	startTime := 0
	for _, log := range logs {
		logItem := parseLog(log)
		if logItem.state == STATE_START {
			if !st.isEmpty() {
				// stop current executing function and calculate its excution time
				current := st.top().val
				result[current.id] += logItem.timestamp - startTime
			}
			st.push(newNode(&value{id: logItem.id}))
			startTime = logItem.timestamp
			continue
		}
		// stack is not empty by problem conditions on this step
		if logItem.state != STATE_END || st.isEmpty() {
			continue
		}

		// also current.id must be the same function as logItem.id by problem condition
		current := st.top().val
		st.pop()
		if logItem.id == current.id {
			// increase total function execution time
			result[logItem.id] += logItem.timestamp - startTime + 1
		}
		startTime = logItem.timestamp + 1
	}
	return result
}

type logItem struct {
	id        int
	state     state
	timestamp int
}

func parseLog(log string) logItem {
	parts := strings.Split(log, ":")
	id, _ := strconv.Atoi(parts[0])
	timestamp, _ := strconv.Atoi(parts[2])
	var s state
	if parts[1] == "start" {
		s = STATE_START
	} else if parts[1] == "end" {
		s = STATE_END
	}
	return logItem{
		id:        id,
		state:     s,
		timestamp: timestamp,
	}
}

type value struct {
	id int
}

type node struct {
	val  *value
	next *node
}

func newNode(val *value) *node {
	return &node{val: val}
}

type stack struct {
	head *node
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) isEmpty() bool {
	return s.head == nil
}

func (s *stack) top() *node {
	return s.head
}

// stack must not be empty
func (s *stack) pop() {
	head := s.head
	s.head = s.head.next
	head.next = nil
}

func (s *stack) push(n *node) {
	n.next = s.head
	s.head = n
}
