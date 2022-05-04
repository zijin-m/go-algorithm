package coding

import "testing"

func TestLengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("LengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
