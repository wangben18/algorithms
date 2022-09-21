package problems

import "testing"

func TestPivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 7, 3, 6, 5, 6},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("PivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
