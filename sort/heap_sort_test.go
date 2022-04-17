package sort

import (
	"reflect"
	"testing"
)

func TestHeap_Sort(t *testing.T) {
	type fields struct {
		arr  []int
		sort []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "case1",
			fields: fields{
				arr:  []int{3, 1, 4, 5, 6, 3, 2, 8, 8, 9, 3, 2},
				sort: []int{1, 2, 2, 3, 3, 3, 4, 5, 6, 8, 8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				arr: tt.fields.arr,
			}
			h.Sort()
			if !reflect.DeepEqual(tt.fields.arr, tt.fields.sort) {
				t.Errorf("excepted %v, but got %v", tt.fields.sort, tt.fields.arr)
			}
		})
	}
}
