package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal",
			args: args{
				a: []int{3, 2, 4, 1},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "single item",
			args: args{
				a: []int{2},
			},
			want: []int{2},
		},
		{
			name: "empty list",
			args: args{
				a: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
