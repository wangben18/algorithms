package problems

func nAryTreelevelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := NodeQueue{}
	queue.Push(root)
	for !queue.Empty() {
		levelOrder := make([]int, 0)
		length := queue.Size()
		for i := 0; i < length; i++ {
			node := queue.Pop()
			levelOrder = append(levelOrder, node.Val)
			for _, n := range node.Children {
				if n == nil {
					continue
				}
				queue.Push(n)
			}
		}
		result = append(result, levelOrder)
	}
	return result
}
