package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head.Next == nil || k == 1 {
		return head
	}
	reverse := func(head *ListNode) {
		cur := head
		var pre *ListNode
		for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
	}
	dummy := &ListNode{}
	dummy.Next = head
	pre, end := dummy, dummy
	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		next := end.Next
		end.Next = nil
		begin := pre.Next
		reverse(begin)
		pre.Next = end
		begin.Next = next

		pre, end = begin, begin
	}

	return dummy.Next
}
