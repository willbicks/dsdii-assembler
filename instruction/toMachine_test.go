package instruction

import "testing"

func Test_instructionR_toMachine(t *testing.T) {
	tests := []struct {
		name string
		i    instructionR
		want uint32
	}{
		{
			name: "empty",
			i:    instructionR{},
			want: 0x0,
		},
		{
			name: "full",
			i: instructionR{
				rS:       0b11111,
				rT:       0b11111,
				rD:       0b11111,
				shiftAmt: 0b11111,
				funct:    0b111111,
			},
			want: 0x03FFFFFF,
		},
		{
			name: "t1",
			i: instructionR{
				rS:       20,
				rT:       8,
				rD:       13,
				shiftAmt: 0,
				funct:    0b010101,
			},
			want: 0x02886815,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.toMachine(); got != tt.want {
				t.Errorf("instructionR.toMachine() = %#x, want %#x", got, tt.want)
			}
		})
	}
}

func Test_instructionI_toMachine(t *testing.T) {
	tests := []struct {
		name string
		i    instructionI
		want uint32
	}{
		{
			name: "empty",
			i:    instructionI{},
			want: 0x0,
		},
		{
			name: "full",
			i: instructionI{
				op:  0b111111,
				rS:  0b11111,
				rT:  0b11111,
				imm: -1,
			},
			want: 0xFFFFFFFF,
		},
		{
			name: "t1",
			i: instructionI{
				op:  8,
				rS:  17,
				rT:  16,
				imm: 5,
			},
			want: 0x22300005,
		},
		{
			name: "t2",
			i: instructionI{
				op:  8,
				rS:  19,
				rT:  8,
				imm: -12,
			},
			want: 0x2268FFF4,
		},
		{
			name: "t3",
			i: instructionI{
				op:  35,
				rS:  0,
				rT:  10,
				imm: 32,
			},
			want: 0x8C0A0020,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.toMachine(); got != tt.want {
				t.Errorf("instructionI.toMachine() = %#x, want %#x", got, tt.want)
			}
		})
	}
}
