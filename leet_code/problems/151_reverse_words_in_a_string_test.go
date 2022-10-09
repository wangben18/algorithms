package problems

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		s     []byte
		begin int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "",
			args: args{
				s:     []byte("abc rab"),
				begin: 0,
				end:   2,
			},
			want: []byte("cba rab"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if reverse(tt.args.s, tt.args.begin, tt.args.end); !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverse() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

func Test_removeExtraSpace(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "",
			args: args{
				s: []byte("  a bb  ccc "),
			},
			want: []byte("a bb ccc"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeExtraSpace(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeExtraSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "  world   hello ",
			},
			want: "hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
