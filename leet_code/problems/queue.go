package problems

type Queue[T any] []T

func (q *Queue[T]) Push(x T) {
	*q = append(*q, x)
}

func (q *Queue[T]) Peek() T {
	if q.Empty() {
		panic("empty queue")
	}
	return (*q)[0]
}

func (q *Queue[T]) Pop() T {
	val := (*q)[0]
	*q = (*q)[1:len(*q)]
	return val
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Size() int {
	return len(*q)
}
