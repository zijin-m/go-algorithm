package huawei

import "testing"

func TestEncryptStr(t *testing.T) {
	type args struct {
		str string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	name: "case1",
		// 	args: args{
		// 		str: "password_a12345678__timeout_100",
		// 		k:   1,
		// 	},
		// 	want: "password_******__timeout_100",
		// },
		{
			name: "case1",
			args: args{
				str: `password_a1#______#2345678__timeout_100`,
				k:   1,
			},
			want: "password_******__timeout_100",
		},
		{
			name: "case1",
			args: args{
				str: `password_a123#__#456#________#78__timeout__##_100`,
				k:   1,
			},
			want: "password_******__timeout__##_100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptStr(tt.args.str, tt.args.k); got != tt.want {
				t.Errorf("EncryptStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
