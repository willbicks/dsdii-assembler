package instruction

import "testing"

func Test_assembleInstruction(t *testing.T) {
	tests := []struct {
		inst    string
		wantMc  uint32
		wantErr bool
	}{
		{
			inst:    "add $s0, $s1, $s2",
			wantMc:  0x02328020,
			wantErr: false,
		},
		{
			inst:    "sub $t0, $t3, $t5",
			wantMc:  0x016D4022,
			wantErr: false,
		},
		{
			inst:    "sll $t0, $s1, 4",
			wantMc:  0x00114100,
			wantErr: false,
		},
		{
			inst:    "srl $s2, $s1, 4",
			wantMc:  0x00119102,
			wantErr: false,
		},
		{
			inst:    "addi $s0, $s1, 5",
			wantMc:  0x22300005,
			wantErr: false,
		},
		{
			inst:    "addi $t0, $s3, -12",
			wantMc:  0x2268FFF4,
			wantErr: false,
		},
		{
			inst:    "lw $t2, 32($0)",
			wantMc:  0x8C0A0020,
			wantErr: false,
		},
		{
			inst:    "sw $s1, 4($t1)",
			wantMc:  0xAD310004,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.inst, func(t *testing.T) {
			got, err := assembleInstruction(tt.inst)
			if (err != nil) != tt.wantErr {
				t.Errorf("assembleInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantMc {
				t.Errorf("assembleInstruction() = %#x, want %#x", got, tt.wantMc)
			}
		})
	}
}
