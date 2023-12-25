package tron

import "testing"

func TestPadLeftStr(t *testing.T) {
	type args struct {
		str    string
		length int
		pad    rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case.1", args{"aaa", 5, '0'}, "00aaa"},
		{"case.2", args{"aaaaa", 5, '0'}, "aaaaa"},
		{"case.3", args{"aaaaaa", 5, '0'}, "aaaaaa"},
		{"case.4", args{"", 5, '0'}, "00000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadLeftStr(tt.args.str, tt.args.length, tt.args.pad); got != tt.want {
				t.Errorf("PadLeftStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
