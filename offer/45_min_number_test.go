package offer

import "testing"

func TestMinNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				nums: []int{10, 2},
			},
			want: "102",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinNumber(tt.args.nums); got != tt.want {
				t.Errorf("MinNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
