package instruction

import "fmt"

// iOpCodes are 6 bit opcodes for I type instructions
var iOpCodes = map[string]uint8{
	"addi": 0b001000,
	"andi": 0b001100,
	"ori":  0b001101,
	"xori": 0b001110,
	"sw":   0b101011,
	"lw":   0b100011,
}

// rFunctCodes are 6 bit function codes for R type instructions
var rFunctCodes = map[string]uint8{
	"add":   0b100000,
	"and":   0b100100,
	"multu": 0b011001,
	"or":    0b100101,
	"sll":   0b000000,
	"sra":   0b000011,
	"srl":   0b000010,
	"sub":   0b100010,
	"xor":   0b100110,
}

// parseOperator returns the integer opcode and funct code from the assembly operator string.
//
// If the the operator is found in rFuncCodes, then the operator is R-type, and the opcode is
// set to zero. Otherwise, the operator is I-type, the correct opcode is returned, and the
// funct code returned is zero.
//
// Returns an error if the provided operator is not a valid R or I type operator.
func parseOperator(op string) (opcode uint8, funct uint8, err error) {
	if funct, ok := rFunctCodes[op]; ok {
		return 0, funct, nil
	}
	if opc, ok := iOpCodes[op]; ok {
		return opc, 0, nil
	}
	return 0, 0, fmt.Errorf("unrecognized operator '%s'", op)
}
