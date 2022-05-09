package coding

import (
	"reflect"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			wantRes: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := FindDuplicates(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("FindDuplicates() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
