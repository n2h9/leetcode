package solution

// type value struct {
// 	vertice int
// 	stops   int
// 	cost    int
// }

// func newValue(vertice, stops, cost int) *value {
// 	return &value{
// 		vertice: vertice,
// 		stops:   stops,
// 		cost:    cost,
// 	}
// }

// type treeNode struct {
// 	val   *value
// 	left  *treeNode
// 	eq    *treeNode
// 	right *treeNode
// }

// func insert(root *treeNode, v *value) *treeNode {
// 	if root == nil {
// 		return &treeNode{
// 			val: v,
// 		}
// 	}
// 	if v.cost == root.val.cost {
// 		root.eq = insert(root.eq, v)
// 		return root
// 	}
// 	if v.cost < root.val.cost {
// 		root.left = insert(root.left, v)
// 		return root
// 	}
// 	root.right = insert(root.right, v)
// 	return root
// }

// type queue struct {
// 	storage *treeNode
// }

// func newPriorityQueue() *queue {
// 	return &queue{
// 		storage: nil,
// 	}
// }

// func (q *queue) isEmpty() bool {
// 	return q.storage == nil
// }

// func (q *queue) top() *value {
// 	t := q.storage
// 	for t.left != nil {
// 		t = t.left
// 	}
// 	return t.val
// }

// func (q *queue) pop() {
// 	if q.isEmpty() {
// 		return
// 	}

// 	tp := q.storage
// 	t := q.storage
// 	for t.left != nil {
// 		tp = t
// 		t = t.left
// 	}

// 	if tp == t {
// 		if t.eq != nil {
// 			q.storage = t.eq
// 			q.storage.right = t.right
// 			return
// 		}
// 		q.storage = q.storage.right
// 		return
// 	}

// 	if t.eq != nil {
// 		tp.left = t.eq
// 		t.eq.right = t.right
// 		return
// 	}

// 	tp.left = t.right
// }

// func (q *queue) push(v *value) {
// 	q.storage = insert(q.storage, v)
// }
