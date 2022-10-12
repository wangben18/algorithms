package problems

type Stack struct {
	slice []int
}

func (this *Stack) Push(x int) {
	this.slice = append(this.slice, x)
}

func (this *Stack) Pop() int {
	val := this.slice[len(this.slice)-1]
	this.slice = this.slice[:len(this.slice)-1]
	return val
}

func (this *Stack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.slice[len(this.slice)-1]
}

func (this *Stack) Empty() bool {
	return len(this.slice) == 0
}

type MyQueue struct {
	in  Stack
	out Stack
}

func MyQueueConstructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	if this.out.Empty() {
		for !this.in.Empty() {
			this.out.Push(this.in.Pop())
		}
	}

	return this.out.Pop()
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	val := this.Pop()
	this.out.Push(val)
	return val
}

func (this *MyQueue) Empty() bool {
	return this.in.Empty() && this.out.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
