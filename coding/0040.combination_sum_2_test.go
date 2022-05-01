package coding

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		{
			name: "case1",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			wantRes: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := CombinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("CombinationSum2() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
