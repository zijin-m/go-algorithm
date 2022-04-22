package huawei

import "testing"

func TestChangeStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				str: "abcdef",
			},
			want: "abcdef",
		},
		{
			name: "case2",
			args: args{
				str: "fsfaf",
			},
			want: "asfff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChangeStr(tt.args.str); got != tt.want {
				t.Errorf("ChangeStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
