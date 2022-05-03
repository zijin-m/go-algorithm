package coding

import (
	"reflect"
	"testing"
)

func TestReorderLogFiles(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{
				logs: []string{"j mo", "5 m w", "g 07", "o 2 0", "t q h"},
			},
			want: []string{"5 m w", "j mo", "t q h", "g 07", "o 2 0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReorderLogFiles(tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReorderLogFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
