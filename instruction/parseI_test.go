package instruction

import (
	"reflect"
	"testing"
)

func Test_parseI(t *testing.T) {
	type args struct {
		opcode    uint8
		opperands []string
	}
	tests := []struct {
		name    string
		args    args
		want    instructionI
		wantErr bool
	}{
		{
			name: "addi $s0, $s1, 5",
			args: args{
				opcode:    8,
				opperands: []string{"$s0", "$s1", "5"},
			},
			want: instructionI{
				op:  8,
				rS:  17,
				rT:  16,
				imm: 5,
			},
		},
		{
			name: "addi $t0, $s3, -12",
			args: args{
				opcode:    8,
				opperands: []string{"$t0", "$s3", "-12"},
			},
			want: instructionI{
				op:  8,
				rS:  19,
				rT:  8,
				imm: -12,
			},
		},
		{
			name: "lw $t2, 32($0)",
			args: args{
				opcode:    35,
				opperands: []string{"$t2", "32($0)"},
			},
			want: instructionI{
				op:  35,
				rS:  0,
				rT:  10,
				imm: 32,
			},
		},
		{
			name: "sw $s1, -4($t1)",
			args: args{
				opcode:    43,
				opperands: []string{"$s1", "-4($t1)"},
			},
			want: instructionI{
				op:  43,
				rS:  9,
				rT:  17,
				imm: -4,
			},
		},
		{
			name: "wrong_opperand_syntax: sw $s1, $t1, -4",
			args: args{
				opcode:    43,
				opperands: []string{"$s1", "$t1", "-4"},
			},
			wantErr: true,
		},
		{
			name: "wrong_opperand_syntax: addi $s1, 4($t1)",
			args: args{
				opcode:    8,
				opperands: []string{"$t2", "4($t1)"},
			},
			wantErr: true,
		},
		{
			name: "wrong_opperand_syntax: addi s1, $t1, 4",
			args: args{
				opcode:    8,
				opperands: []string{"s1", "$t1", "4"},
			},
			wantErr: true,
		},
		{
			name: "too_many_opperands: addi $s1, $t1, $s0, 4",
			args: args{
				opcode:    8,
				opperands: []string{"$s1", "$t1", "$s0", "4"},
			},
			wantErr: true,
		},
		{
			name: "not_enough_opperands: addi $s1, 4",
			args: args{
				opcode:    8,
				opperands: []string{"$s1", "4"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseI(tt.args.opcode, tt.args.opperands)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitParenOpperand(t *testing.T) {
	tests := []struct {
		str     string
		wantA   string
		wantB   string
		wantErr bool
	}{
		{
			str:   "a(b)",
			wantA: "a",
			wantB: "b",
		},
		{
			str:   "-2($s1)",
			wantA: "-2",
			wantB: "$s1",
		},
		{
			str:   "128($0)",
			wantA: "128",
			wantB: "$0",
		},
		{
			str:     "-2($s1",
			wantErr: true,
		},
		{
			str:     "-2$s1)",
			wantErr: true,
		},
		{
			str:     "-2$s1",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			gotA, gotB, err := splitParenOpperand(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitParenOpperand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotA != tt.wantA {
				t.Errorf("splitParenOpperand() gotA = %v, want %v", gotA, tt.wantA)
			}
			if gotB != tt.wantB {
				t.Errorf("splitParenOpperand() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
