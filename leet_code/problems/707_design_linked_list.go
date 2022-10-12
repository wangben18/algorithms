package problems

type MyLinkedList struct {
	head *ListNode
}

func MyLinkedListConstructor() MyLinkedList {
	return MyLinkedList{}
}

func (list *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}
	node := list.head
	for ; index > 0; index-- {
		if node == nil {
			break
		}
		node = node.Next
	}
	if node == nil {
		return -1
	}
	return node.Val
}

func (list *MyLinkedList) AddAtHead(val int) {
	list.head = &ListNode{
		Val:  val,
		Next: list.head,
	}
}

func (list *MyLinkedList) AddAtTail(val int) {
	tail := &ListNode{Val: val}
	if list.head == nil {
		list.head = tail
		return
	}
	node := list.head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = tail
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		list.head = &ListNode{
			Val:  val,
			Next: list.head,
		}
		return
	}
	node := list.head
	for ; index > 1 && node != nil; index-- {
		node = node.Next
	}
	if node == nil {
		return
	}
	newNode := &ListNode{
		Val:  val,
		Next: node.Next,
	}
	node.Next = newNode
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	if index == 0 && list.head != nil {
		list.head = list.head.Next
		return
	}
	node := list.head
	for ; index > 1 && node != nil; index-- {
		node = node.Next
	}
	if node == nil || node.Next == nil {
		return
	}
	node.Next = node.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
