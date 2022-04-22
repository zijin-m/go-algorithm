package coding

import "testing"

func TestCanPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				nums: []int{3, 3, 3, 4, 5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanPartition(tt.args.nums); got != tt.want {
				t.Errorf("CanPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
