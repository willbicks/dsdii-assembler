package instruction

import (
	"reflect"
	"testing"
)

func Test_parseR(t *testing.T) {
	type args struct {
		funct uint8
		op    []string
	}
	tests := []struct {
		name    string
		args    args
		want    instructionR
		wantErr bool
	}{
		{
			name: "good: add $s0, $s1, $s2",
			args: args{
				funct: 32,
				op:    []string{"$s0", "$s1", "$s2"},
			},
			want: instructionR{
				rS:       17,
				rT:       18,
				rD:       16,
				shiftAmt: 0,
				funct:    32,
			},
			wantErr: false,
		},
		{
			name: "good: sub $t0, $t3, $t5",
			args: args{
				funct: 34,
				op:    []string{"$t0", "$t3", "$t5"},
			},
			want: instructionR{
				rS:       11,
				rT:       13,
				rD:       8,
				shiftAmt: 0,
				funct:    34,
			},
			wantErr: false,
		},
		{
			name: "good: sll $t0, $s1, 4",
			args: args{
				funct: 0,
				op:    []string{"$t0", "$s1", "4"},
			},
			want: instructionR{
				rS:       0,
				rT:       17,
				rD:       8,
				shiftAmt: 4,
				funct:    0,
			},
			wantErr: false,
		},
		{
			name: "good: srl $s2, $s1, 8",
			args: args{
				funct: 2,
				op:    []string{"$s2", "$s1", "8"},
			},
			want: instructionR{
				rS:       0,
				rT:       17,
				rD:       18,
				shiftAmt: 8,
				funct:    2,
			},
			wantErr: false,
		},
		{
			name: "good: sra $s3, $s1, 8",
			args: args{
				funct: 3,
				op:    []string{"$s3", "$s1", "1"},
			},
			want: instructionR{
				rS:       0,
				rT:       17,
				rD:       19,
				shiftAmt: 1,
				funct:    3,
			},
			wantErr: false,
		},
		{
			name: "too many operrands: sub $t0, $t3, $t5, 3",
			args: args{
				funct: 34,
				op:    []string{"$t0", "$t3", "$t5", "3"},
			},
			want:    instructionR{},
			wantErr: true,
		},
		{
			name: "not enough operrands: sub $t0, $t3",
			args: args{
				funct: 34,
				op:    []string{"$t0", "$t3"},
			},
			want:    instructionR{},
			wantErr: true,
		},
		{
			name: "wrong type of opperand: sub $t0, $t3, 3",
			args: args{
				funct: 34,
				op:    []string{"$t0", "$t3", "3"},
			},
			want:    instructionR{},
			wantErr: true,
		},
		{
			name: "wrong type of opperand: sra $s2, $s1, $s0",
			args: args{
				funct: 2,
				op:    []string{"$t0", "$t3", "$s0"},
			},
			want:    instructionR{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseR(tt.args.funct, tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseR() = %v, want %v", got, tt.want)
			}
		})
	}
}
