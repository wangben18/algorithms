package problems

type DoubleLinkedNode struct {
	Pre   *DoubleLinkedNode
	Next  *DoubleLinkedNode
	Value interface{}
}
