package problems

type Stack[T any] []T

func (s *Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *Stack[T]) Pop() T {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *Stack[T]) Top() T {
	if s.Empty() {
		panic("empty stack")
	}
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) Empty() bool {
	return len(*s) == 0
}
