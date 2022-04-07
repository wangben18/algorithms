package leet_code

type ListNode struct {
	Next *ListNode
	Val  int
}

func ReverseKNodeList(head *ListNode, k int) *ListNode {
	if head.Next == nil || k == 1 {
		return head
	}

	var nodes []*ListNode
	for n := head; n != nil; n = n.Next {
		nodes = append(nodes, n)
	}

	for offset := 0; offset < len(nodes)/k; offset++ {
		offsetIndex := offset * k
		nodes[offsetIndex].Next = func() *ListNode {
			nextGroupEnd := offsetIndex + 2*k - 1
			if nextGroupEnd >= len(nodes) {
				if len(nodes)%k == 0 {
					return nil
				}
				return nodes[len(nodes)-len(nodes)%k]
			}
			return nodes[nextGroupEnd]
		}()
		for i := 1; i < k; i++ {
			nodes[i+offsetIndex].Next = nodes[i+offsetIndex-1]
		}
	}

	return nodes[k-1]
}
