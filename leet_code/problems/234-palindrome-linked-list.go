package problems

// 方法：快慢指针
func isPalindrome(head *ListNode) bool {
	pre, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}
	pre.Next = nil
	reversedSlow := reverseList(slow)
	for ; head != nil; head, reversedSlow = head.Next, reversedSlow.Next {
		if head.Val != reversedSlow.Val {
			return false
		}
	}
	return true
}

// 方法：导出到数组
func isPalindrome2(head *ListNode) bool {
	cur, count := head, 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	cur = head
	nums := make([]int, count)
	for i := 0; cur != nil; i, cur = i+1, cur.Next {
		nums[i] = cur.Val
	}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}
	return true
}
