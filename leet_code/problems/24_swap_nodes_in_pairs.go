package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		tmpNext := cur.Next
		tmpNextNextNext := cur.Next.Next.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = tmpNext
		cur.Next.Next.Next = tmpNextNextNext
		cur = tmpNext
	}

	return dummyHead.Next
}
