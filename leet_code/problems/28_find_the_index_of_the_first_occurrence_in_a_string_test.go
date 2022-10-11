package problems

import (
	"reflect"
	"testing"
)

func Test_getNext(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				s: "aabaaf",
			},
			want: []int{0, 1, 0, 1, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNext(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				haystack: "aabaabaaf",
				needle:   "aabaaf",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
