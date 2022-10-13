package problems

type MyQueue struct {
	in  Stack[int]
	out Stack[int]
}

func MyQueueConstructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.in.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.Empty() {
		return 0
	}
	if q.out.Empty() {
		for !q.in.Empty() {
			q.out.Push(q.in.Pop())
		}
	}

	return q.out.Pop()
}

func (q *MyQueue) Peek() int {
	if q.Empty() {
		return 0
	}
	val := q.Pop()
	q.out.Push(val)
	return val
}

func (q *MyQueue) Empty() bool {
	return q.in.Empty() && q.out.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
