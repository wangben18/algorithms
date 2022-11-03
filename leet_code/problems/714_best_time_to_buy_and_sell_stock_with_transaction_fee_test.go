package problems

import "testing"

func Test_maxProfitWithFee(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				prices: []int{},
				fee:    0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitWithFee(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfitWithFee() = %v, want %v", got, tt.want)
			}
		})
	}
}
