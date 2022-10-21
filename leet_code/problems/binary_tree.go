package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeQueue struct {
	s []*TreeNode
}

func (q *TreeNodeQueue) Push(x *TreeNode) {
	q.s = append(q.s, x)
}

func (q *TreeNodeQueue) Peek() *TreeNode {
	if q.Empty() {
		panic("empty queue")
	}
	return (q.s)[0]
}

func (q *TreeNodeQueue) Pop() *TreeNode {
	val := (q.s)[0]
	q.s = (q.s)[1:len(q.s)]
	return val
}

func (q *TreeNodeQueue) Empty() bool {
	return len(q.s) == 0
}

func (q *TreeNodeQueue) Size() int {
	return len(q.s)
}
