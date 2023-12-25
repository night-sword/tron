package tron

import "testing"

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
