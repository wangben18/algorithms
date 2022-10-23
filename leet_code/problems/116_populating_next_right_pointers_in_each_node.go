package problems

type NextNode struct {
	Val   int
	Left  *NextNode
	Right *NextNode
	Next  *NextNode
}

func connect(root *NextNode) *NextNode {
	if root == nil {
		return nil
	}
	queue := Queue[*NextNode]{}
	queue.Push(root)
	for !queue.Empty() {
		length := queue.Size()
		var prev *NextNode
		for i := 0; i < length; i++ {
			node := queue.Pop()
			node.Next = prev
			if node.Right != nil {
				queue.Push(node.Right)
			}
			if node.Left != nil {
				queue.Push(node.Left)
			}
			prev = node
		}
	}
	return root
}
