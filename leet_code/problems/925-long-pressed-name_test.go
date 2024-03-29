package problems

import "testing"

func Test_isLongPressedName(t *testing.T) {
	type args struct {
		name  string
		typed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				name:  "alex",
				typed: "aaleex",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				name:  "saeed",
				typed: "ssaaedd",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				name:  "abc",
				typed: "aabcc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLongPressedName(tt.args.name, tt.args.typed); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
