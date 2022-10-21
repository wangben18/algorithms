package problems

import (
	"reflect"
	"testing"
)

func Test_nAryTreelevelOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{{
						Val: 3,
						Children: []*Node{{
							Val:      5,
							Children: []*Node{},
						}, {
							Val:      6,
							Children: []*Node{},
						}},
					}, {
						Val:      2,
						Children: []*Node{},
					}, {
						Val:      4,
						Children: []*Node{},
					}},
				},
			},
			want: [][]int{{1}, {3, 2, 4}, {5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nAryTreelevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nAryTreelevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
