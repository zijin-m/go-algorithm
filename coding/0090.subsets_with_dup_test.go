package coding

import (
	"reflect"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
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
				nums: []int{1, 2, 2},
			},
			wantRes: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 2},
				{2},
				{2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := SubsetsWithDup(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("SubsetsWithDup() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
