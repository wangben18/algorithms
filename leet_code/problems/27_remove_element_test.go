package problems

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name: "",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want:     2,
			wantNums: []int{2, 2},
		},
		{
			name: "",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want:     5,
			wantNums: []int{0, 1, 4, 0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want && func() bool {
				for i, num := range tt.wantNums {
					if tt.args.nums[i] != num {
						return false
					}
				}
				return true
			}() {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
