package tron

import "testing"

func TestIsAddressValid(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case.1", args{"T9yW5v9biiSkpsdFNjfjt2WWaz6jVhG888"}, true},
		{"case.2", args{"TSNDArxC7ejNax8aHY8hRzqn9GcSQSK777"}, true},
		{"case.3", args{"TSNDArxC7ejNax8aHY8hRzqn9GcSQSK771"}, false},
		{"case.4", args{"TSNDArxC7ejNax8aHY8hRzqn9GcSQSK772"}, false},
		{"case.5", args{"1"}, false},
		{"case.6", args{""}, false},
		{"case.7", args{"1111111111111111111111111111111111"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAddressValid(tt.args.address); got != tt.want {
				t.Errorf("IsAddressValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
