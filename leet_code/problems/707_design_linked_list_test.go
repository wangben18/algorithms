package problems

import (
	"reflect"
	"testing"
)

func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
	type fields struct {
		head *ListNode
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name: "",
			fields: fields{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			args: args{
				index: 0,
			},
			want: fields{
				head: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyLinkedList{
				head: tt.fields.head,
			}
			this.DeleteAtIndex(tt.args.index)
			if !reflect.DeepEqual(this.head, tt.want.head) {
				t.Errorf("DeleteAtIndex(), got %v, want %v", this.head, tt.want.head)
			}
		})
	}
}

func TestMyLinkedList_AddAtIndex(t *testing.T) {
	type fields struct {
		head *ListNode
	}
	type args struct {
		index int
		val   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ListNode
	}{
		{
			name: "",
			fields: fields{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			args: args{
				index: 1,
				val:   2,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyLinkedList{
				head: tt.fields.head,
			}
			this.AddAtIndex(tt.args.index, tt.args.val)
			if !reflect.DeepEqual(tt.want, this.head) {
				t.Errorf("AddAtIndex(), got %v, want %v", this.head, tt.want)
			}
		})
	}
}

func TestMyLinkedList_AddAtTail(t *testing.T) {
	type fields struct {
		head *ListNode
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ListNode
	}{
		{
			name: "",
			fields: fields{
				head: nil,
			},
			args: args{
				val: 2,
			},
			want: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyLinkedList{
				head: tt.fields.head,
			}
			this.AddAtTail(tt.args.val)
			if !reflect.DeepEqual(tt.want, this.head) {
				t.Errorf("AddAtTail(), got %v, want %v", this.head, tt.want)
			}
		})
	}
}
