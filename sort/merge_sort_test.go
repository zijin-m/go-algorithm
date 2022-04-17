package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		arr  []int
		sort []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr:  []int{3, 1, 4, 5, 6, 3, 2, 8, 8, 9, 3, 2},
				sort: []int{1, 2, 2, 3, 3, 3, 4, 5, 6, 8, 8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.args.sort) {
				t.Errorf("excepted %v, but got %v", tt.args.sort, tt.args.arr)
			}
		})
	}
}
