package coding

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case1",
			args: args{
				inorder:   []int{},
				postorder: []int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildTree(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
