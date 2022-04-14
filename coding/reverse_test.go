package coding

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		out  []byte
	}{
		{
			name: "case1",
			in:   []byte{1, 2, 3, 4, 5},
			out:  []byte{5, 4, 3, 2, 1},
		},
		{
			name: "case2",
			in:   []byte{1, 2},
			out:  []byte{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Reverse(tt.in); !reflect.DeepEqual(tt.in, tt.out) {
				t.Errorf("Reverse %v excepted %v, but got %v", tt.in, tt.out, tt.in)
			}
		})
	}
}
