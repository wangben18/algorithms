package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	n1 := headA
	n2 := headB
	for n1 != n2 {
		if n1 != nil {
			n1 = n1.Next
		} else {
			n1 = headB
		}
		if n2 != nil {
			n2 = n2.Next
		} else {
			n2 = headA
		}
	}
	return n1
}
