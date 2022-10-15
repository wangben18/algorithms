package problems

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: []int{1, 2, 3, 6, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
