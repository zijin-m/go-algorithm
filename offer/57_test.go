package offer

import (
	"reflect"
	"testing"
)

func Test_findContinuousSequence(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		{
			name: "case1",
			args: args{
				target: 9,
			},
			wantRes: [][]int{
				{2, 3, 4},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findContinuousSequence(tt.args.target); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("findContinuousSequence() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
