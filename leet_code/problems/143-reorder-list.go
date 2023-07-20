package problems

func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}
	pre, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}
	pre.Next = nil

	cur1, cur2 := head, reverseList(slow)
	for cur1.Next != nil {
		cur1.Next, cur2.Next, cur1, cur2 = cur2, cur1.Next, cur1.Next, cur2.Next
	}

	cur1.Next = cur2
}
