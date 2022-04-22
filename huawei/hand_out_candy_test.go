package huawei

import (
	"reflect"
	"testing"
)

func TestHandOutCandy(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAva int
		wantR1  []int
		wantR2  []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{7, 4, 5, 3, 3},
			},
			wantAva: 11,
			wantR1:  []int{4, 7},
			wantR2:  []int{5, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAva, gotR1, gotR2 := HandOutCandy(tt.args.nums)
			if gotAva != tt.wantAva {
				t.Errorf("HandOutCandy() gotAva = %v, want %v", gotAva, tt.wantAva)
			}
			if !reflect.DeepEqual(gotR1, tt.wantR1) {
				t.Errorf("HandOutCandy() gotR1 = %v, want %v", gotR1, tt.wantR1)
			}
			if !reflect.DeepEqual(gotR2, tt.wantR2) {
				t.Errorf("HandOutCandy() gotR2 = %v, want %v", gotR2, tt.wantR2)
			}
		})
	}
}
