package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		head = list2
		list2 = list2.Next
	} else {
		head = list1
		list1 = list1.Next
	}
	cur := head
	for {
		if list1 == nil {
			cur.Next = list2
			break
		} else if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}
	return head
}
