package leet_code

import (
	"fmt"
	"testing"
)

func (node *ListNode) Print() {
	n := node
	for n.Next != nil {
		fmt.Print(n.Val, "->")
		n = n.Next
	}
	fmt.Println(n.Val)
}

func TestReverseKNodeList(t *testing.T) {
	n := &ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: &ListNode{
					Next: &ListNode{
						Next: &ListNode{
							Next: &ListNode{
								Next: &ListNode{
									Next: &ListNode{
										Next: &ListNode{
											Next: &ListNode{
												Next: nil,
												Val:  11,
											},
											Val: 10,
										},
										Val: 9,
									},
									Val: 8,
								},
								Val: 7,
							},
							Val: 6,
						},
						Val: 5,
					},
					Val: 4,
				},
				Val: 3,
			},
			Val: 2,
		},
		Val: 1,
	}

	n.Print()
	n = ReverseKNodeList(n, 4)
	n.Print()
}
