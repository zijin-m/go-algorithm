package huawei

import (
	"reflect"
	"testing"
)

func TestHasNumCombine(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{7, 3, 5, 2},
			},
			want: []int{7, 3, 2},
		},
		{
			name: "case2",
			args: args{
				nums: []int{3, 50, 14, 1, 6, 2},
			},
			want: []int{14, 2, 6},
		},
		{
			name: "case3",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "case4",
			args: args{
				nums: []int{8, 2},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasNumCombine(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasNumCombine() = %v, want %v", got, tt.want)
			}
		})
	}
}
