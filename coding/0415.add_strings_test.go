package coding

import "testing"

func TestAddStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				num1: "11",
				num2: "123",
			},
			want: "134",
		},
		{
			name: "case2",
			args: args{
				num1: "9",
				num2: "99",
			},
			want: "108",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("AddStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
