package problems

type MyStack struct {
	queue Queue[int]
}

func MyStackConstructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.queue.Push(x)
}

func (s *MyStack) Pop() int {
	for i := 1; i <= s.queue.Size()-1; i++ {
		s.queue.Push(s.queue.Pop())
	}
	return s.queue.Pop()
}

func (s *MyStack) Top() int {
	val := s.Pop()
	s.Push(val)
	return val
}

func (s *MyStack) Empty() bool {
	return s.queue.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
