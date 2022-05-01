package coding

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]string
	}{
		{
			name: "case1",
			args: args{
				s: "aab",
			},
			wantRes: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Partition(tt.args.s); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Partition() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		s     string
		start int
		path  []string
		res   *[][]string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			partition(tt.args.s, tt.args.start, tt.args.path, tt.args.res)
		})
	}
}
