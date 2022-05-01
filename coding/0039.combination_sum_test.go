package coding

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
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
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			wantRes: [][]int{
				{2, 2, 3},
				{7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := CombinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("CombinationSum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
