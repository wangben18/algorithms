package problems

type Node struct {
	Val      int
	Children []*Node
}

// for leetcode
type NodeQueue struct {
	s []*Node
}

func (q *NodeQueue) Push(x *Node) {
	q.s = append(q.s, x)
}

func (q *NodeQueue) Peek() *Node {
	if q.Empty() {
		panic("empty queue")
	}
	return (q.s)[0]
}

func (q *NodeQueue) Pop() *Node {
	val := (q.s)[0]
	q.s = (q.s)[1:len(q.s)]
	return val
}

func (q *NodeQueue) Empty() bool {
	return len(q.s) == 0
}

func (q *NodeQueue) Size() int {
	return len(q.s)
}
