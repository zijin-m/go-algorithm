package huawei

import "testing"

func TestTLV(t *testing.T) {
	type args struct {
		data []string
		tag  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				data: []string{"32", "01", "00", "AE", "90", "02", "00", "01", "02", "30", "03", "00", "AB", "32", "31", "31", "02", "00", "32", "33", "33", "01", "00", "CC"},
				tag:  "31",
			},
			want: "32 33",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TLV(tt.args.data, tt.args.tag); got != tt.want {
				t.Errorf("TLV() = %v, want %v", got, tt.want)
			}
		})
	}
}
