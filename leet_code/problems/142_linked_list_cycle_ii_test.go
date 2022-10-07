package problems

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	cycleNode := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  -4,
				Next: nil,
			},
		},
	}
	cycleNode.Next.Next.Next = cycleNode
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: &ListNode{
					Val:  3,
					Next: cycleNode,
				},
			},
			want: cycleNode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
