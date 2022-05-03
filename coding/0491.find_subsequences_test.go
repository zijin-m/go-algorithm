package coding

import (
	"reflect"
	"testing"
)

func TestFindSubsequences(t *testing.T) {
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
				nums: []int{4, 4, 3, 2, 1},
			},
			wantRes: [][]int{{4, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := FindSubsequences(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("FindSubsequences() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_findSubsequences(t *testing.T) {
	type args struct {
		nums  []int
		start int
		path  []int
		res   *[][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findSubsequences(tt.args.nums, tt.args.start, tt.args.path, tt.args.res)
		})
	}
}
