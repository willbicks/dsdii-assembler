package instruction

import (
	"fmt"
	"strconv"

	"github.com/willbicks/dsdii-assembler/register"
)

// parseR parses an R type instruciton with the provided function code and opperands.
func parseR(funct uint8, op []string) (instructionR, error) {
	var ir = instructionR{
		funct: funct,
	}

	if len(op) != 3 {
		return instructionR{}, fmt.Errorf("unexpected number of opperands (want 3)")
	}

	// parse the destination register (always first)
	var err error
	ir.rD, err = register.Parse(op[0])
	if err != nil {
		return instructionR{}, err
	}

	if funct == rFunctCodes["sra"] || funct == rFunctCodes["srl"] || funct == rFunctCodes["sll"] {
		// if shift instruciton, parse rT and shift ammount
		ir.rT, err = register.Parse(op[1])
		if err != nil {
			return instructionR{}, err
		}

		s64, err := strconv.ParseUint(op[2], 10, 8)
		if err != nil {
			return instructionR{}, fmt.Errorf("unable to parse shift ammount: %w", err)
		}
		ir.shiftAmt = uint8(s64)
	} else {
		// if not shift instruciton, parse rS and rT
		ir.rS, err = register.Parse(op[1])
		if err != nil {
			return instructionR{}, err
		}

		ir.rT, err = register.Parse(op[2])
		if err != nil {
			return instructionR{}, err
		}

	}

	return ir, nil
}
