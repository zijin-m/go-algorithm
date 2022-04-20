package huawei

import (
	"reflect"
	"testing"
)

func TestStreetLight(t *testing.T) {
	type args struct {
		light []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				light: []int{50, 50},
				n:     2,
			},
			want: 0,
		},
		{
			name: "case2",
			args: args{
				light: []int{50, 70, 20, 70},
				n:     4,
			},
			want: 20,
		},
		{
			name: "case3",
			args: args{
				light: []int{0, 400, 0, 0},
				n:     4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StreetLight(tt.args.light, tt.args.n); got != tt.want {
				t.Errorf("StreetLight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		{
			name: "case1",
			args: args{
				intervals: [][]int{{1, 4}, {0, 4}},
			},
			wantRes: [][]int{{0, 4}},
		},
		{
			name: "case2",
			args: args{
				intervals: [][]int{{1, 4}, {0, 2}, {3, 5}},
			},
			wantRes: [][]int{{0, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Merge(tt.args.intervals); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Merge() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
