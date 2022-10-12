package problems

type Queue struct {
	slice []int
}

func (this *Queue) Push(x int) {
	this.slice = append(this.slice, x)
}

func (this *Queue) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.slice[0]
}

func (this *Queue) Pop() int {
	val := this.slice[0]
	this.slice = this.slice[1:len(this.slice)]
	return val
}

func (this *Queue) Empty() bool {
	return len(this.slice) == 0
}

func (this *Queue) Size() int {
	return len(this.slice)
}

type MyStack struct {
	queue Queue
}

func MyStackConstructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue.Push(x)
}

func (this *MyStack) Pop() int {
	for i := 1; i <= this.queue.Size()-1; i++ {
		this.queue.Push(this.queue.Pop())
	}
	return this.queue.Pop()
}

func (this *MyStack) Top() int {
	val := this.Pop()
	this.Push(val)
	return val
}

func (this *MyStack) Empty() bool {
	return this.queue.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
