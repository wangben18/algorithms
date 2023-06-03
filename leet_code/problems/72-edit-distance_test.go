package problems

import "testing"

func Test_editDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := editDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("editDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
