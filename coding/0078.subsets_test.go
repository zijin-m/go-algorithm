package coding

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 3},
			},
			wantRes: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 3},
				{1, 3},
				{2},
				{2, 3},
				{3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Subsets(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Subsets() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
