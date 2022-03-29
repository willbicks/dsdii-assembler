package instruction

import (
	"fmt"
	"testing"
)

func Test_parsePseudoVsAssemble(t *testing.T) {
	type args struct {
		operator  string
		opperands []string
	}
	tests := []struct {
		name    string
		args    args
		asm     string
		wantNil bool
		wantErr bool
	}{
		{
			name: "nop",
			args: args{
				operator: "nop",
			},
			asm: "sll $0, $0, 0",
		},
		{
			name: "nop_wrong_arg",
			args: args{
				operator:  "nop",
				opperands: []string{"$s0"},
			},
			wantErr: true,
			wantNil: true,
		},
		{
			name: "move",
			args: args{
				operator:  "move",
				opperands: []string{"$s0", "$t1"},
			},
			asm: "add $s0, $t1, $0",
		},
		{
			name: "move_wrong_arg",
			args: args{
				operator:  "move",
				opperands: []string{"$s0"},
			},
			wantErr: true,
			wantNil: true,
		},
		{
			name: "clear",
			args: args{
				operator:  "clear",
				opperands: []string{"$t0"},
			},
			asm: "add $t0, $0, $0",
		},
		{
			name: "clear_wrong_arg",
			args: args{
				operator:  "clear",
				opperands: []string{"$s0", "$t1"},
			},
			wantErr: true,
			wantNil: true,
		},
		{
			name: "non-pseudo",
			args: args{
				operator:  "add",
				opperands: []string{"$t0", "$t1", "$t2"},
			},
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parsePseudo(tt.args.operator, tt.args.opperands)

			if (err != nil) != tt.wantErr {
				t.Errorf("parsePseudo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantNil {
				if got != nil {
					t.Errorf("parsePseudo() = %v, wantNil %v", err, tt.wantNil)
					return
				}
			} else {
				want, err := Assemble(tt.asm)
				if err != nil {
					panic(fmt.Errorf("assembling reference instruciton failed: %w", err))
				}
				if got.toMachine() != want {
					t.Errorf("parsePseudo() = %#08x, want %#08x", got, want)
				}
			}

		})
	}
}
