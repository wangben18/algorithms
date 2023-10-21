package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var head *ListNode
	for i := range lists {
		head = mergeTwoLists(head, lists[i])
	}
	return head
}
