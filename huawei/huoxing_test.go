package huawei

import "testing"

func TestHuoxing(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "case1",
			args: args{
				str: "7#6$5#12",
			},
			wantAns: 226,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := Huoxing(tt.args.str); gotAns != tt.wantAns {
				t.Errorf("Huoxing() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
