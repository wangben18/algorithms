package problems

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: &ListNode{
					Val:  2,
					Next: nil,
				},
				n: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
