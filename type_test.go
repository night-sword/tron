package tron

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestSUN_TRX(t *testing.T) {
	tests := []struct {
		name string
		inst SUN
		want TRX
	}{
		{"case.1", SUN(1), TRX(0.000001)},
		{"case.2", SUN(10), TRX(0.00001)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.inst.TRX(); got != tt.want {
				t.Errorf("TRX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSUN_Int64E(t *testing.T) {
	tests := []struct {
		name    string
		inst    SUN
		wantU   int64
		wantErr bool
	}{
		{"case.1", SUN(1), 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotU, err := tt.inst.Int64E()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotU != tt.wantU {
				t.Errorf("Int64E() gotU = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}

func TestTRX_Float64E(t *testing.T) {
	tests := []struct {
		name    string
		inst    TRX
		wantF   float64
		wantErr bool
	}{
		{"case.1", TRX(0.000001), 0.000001, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF, err := tt.inst.Float64E()
			if (err != nil) != tt.wantErr {
				t.Errorf("Float64E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotF != tt.wantF {
				t.Errorf("Float64E() gotF = %v, want %v", gotF, tt.wantF)
			}
		})
	}
}

func TestTRX_SUN(t *testing.T) {
	tests := []struct {
		name string
		inst TRX
		want SUN
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.inst.SUN(); got != tt.want {
				t.Errorf("SUN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTRX_CeilInt64(t *testing.T) {
	tests := []struct {
		name string
		inst TRX
		want int64
	}{
		{"case.1", TRX(1.01), 2},
		{"case.2", TRX(1.61), 2},
		{"case.3", TRX(0.9), 1},
		{"case.4", TRX(0.1), 1},
		{"case.5", TRX(0), 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.inst.CeilInt64(); got != tt.want {
				t.Errorf("CeilInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTRX_CeilUint64(t *testing.T) {
	tests := []struct {
		name string
		inst TRX
		want uint64
	}{
		{"case.1", TRX(1.01), 2},
		{"case.2", TRX(1.61), 2},
		{"case.3", TRX(0.9), 1},
		{"case.4", TRX(0.1), 1},
		{"case.5", TRX(0), 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.inst.CeilUint64(); got != tt.want {
				t.Errorf("CeilUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTRX_Ceil(t *testing.T) {
	tests := []struct {
		name string
		inst TRX
		want TRX
	}{
		{"case.1", TRX(1.01), 2},
		{"case.2", TRX(1.61), 2},
		{"case.3", TRX(0.9), 1},
		{"case.4", TRX(0.1), 1},
		{"case.5", TRX(0), 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.inst.Ceil(); got != tt.want {
				t.Errorf("Ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSUN_Decimal(t *testing.T) {
	tests := []struct {
		name     string
		sun      SUN
		expected string
	}{
		{
			name:     "zero",
			sun:      SUN(0),
			expected: "0",
		},
		{
			name:     "1 SUN",
			sun:      SUN(1),
			expected: "0.000001",
		},
		{
			name:     "1000 SUN",
			sun:      SUN(1000),
			expected: "0.001",
		},
		{
			name:     "1 TRX SUN (1000000)",
			sun:      SUN(SUN_VALUE),
			expected: "1",
		},
		{
			name:     "10 TRX SUN",
			sun:      SUN(10 * SUN_VALUE),
			expected: "10",
		},
		{
			name:     "decimal number",
			sun:      SUN(1500000),
			expected: "1.5",
		},
		{
			name:     "big number",
			sun:      SUN(1000000000000),
			expected: "1000000",
		},
		{
			name:     "max uint64",
			sun:      SUN(18446744073709551615),
			expected: "18446744073709.551615",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.sun.Decimal()

			if _, ok := interface{}(result).(decimal.Decimal); !ok {
				t.Errorf("expect type decimal.Decimal, but got type: %T", result)
			}

			expected, _ := decimal.NewFromString(tt.expected)
			if !result.Equal(expected) {
				t.Errorf("expect: %s, got: %s", expected.String(), result.String())
			}
		})
	}
}

func TestNewTRXDecimalFromInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected string
	}{
		{
			name:     "zero value",
			input:    0,
			expected: "0",
		},
		{
			name:     "positive value - 1 SUN",
			input:    1,
			expected: "0.000001",
		},
		{
			name:     "positive value - 1000000 SUN equals 1 TRX",
			input:    1000000,
			expected: "1",
		},
		{
			name:     "positive value - 500000 SUN equals 0.5 TRX",
			input:    500000,
			expected: "0.5",
		},
		{
			name:     "positive value - large number",
			input:    1234567890,
			expected: "1234.56789",
		},
		{
			name:     "negative value - -1 SUN",
			input:    -1,
			expected: "-0.000001",
		},
		{
			name:     "negative value - -1000000 SUN equals -1 TRX",
			input:    -1000000,
			expected: "-1",
		},
		{
			name:     "negative value - large negative number",
			input:    -1234567890,
			expected: "-1234.56789",
		},
		{
			name:     "max int64 value",
			input:    9223372036854775807,
			expected: "9223372036854.775807",
		},
		{
			name:     "min int64 value",
			input:    -9223372036854775808,
			expected: "-9223372036854.775808",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewTRXDecimalFromInt64(tt.input)
			expected, err := decimal.NewFromString(tt.expected)
			if err != nil {
				t.Fatalf("failed to create expected decimal: %v", err)
			}

			if !result.Equal(expected) {
				t.Errorf("NewTRXDecimalFromInt64(%d) = %s, want %s", tt.input, result.String(), expected.String())
			}
		})
	}
}
