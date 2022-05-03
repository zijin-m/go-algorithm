package coding

import "testing"

func TestMyAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "42",
			},
			want: 42,
		},
		{
			name: "case2",
			args: args{
				s: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "case3",
			args: args{
				s: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "case4",
			args: args{
				s: "2147483646",
			},
			want: 2147483646,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyAtoi(tt.args.s); got != tt.want {
				t.Errorf("MyAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
