package instruction

import (
	"strings"
)

type instruction interface {
	toMachine() uint32
}

// instructionI represents the fields of an I type instruciton
type instructionI struct {
	op  uint8
	rS  uint8
	rT  uint8
	imm int16
}

// instrucitonR represents the feilds of an R type instruciton
type instructionR struct {
	rS       uint8
	rT       uint8
	rD       uint8
	shiftAmt uint8
	funct    uint8
}

// Assemble parses a complete instruction (less semicolon and loewrcase) such as add $t0, $t0, $t1, and returns it's machine code binary value.
func Assemble(istr string) (uint32, error) {
	// split the instruction into operator and opperands
	inst := strings.SplitN(istr, " ", 2)

	// split the opperands by comma and strip whitespace
	opperands := strings.Split(inst[1], ",")
	for i := range opperands {
		opperands[i] = strings.TrimSpace(opperands[i])
	}

	opcode, funct, err := parseOperator(inst[0])
	if err != nil {
		return 0, err
	}

	var instr instruction

	if opcode == 0 {
		// process R type instruction
		instr, err = parseR(funct, opperands)
		if err != nil {
			return 0, err
		}
	} else {
		// process I type instruction
		instr, err = parseI(opcode, opperands)
		if err != nil {
			return 0, err
		}
	}

	return instr.toMachine(), nil
}
