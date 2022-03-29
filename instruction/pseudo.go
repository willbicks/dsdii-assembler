package instruction

import (
	"errors"

	"github.com/willbicks/dsdii-assembler/register"
)

// attempts to parse a pseudo instruction, returns a nil instruction and nil error if not a pseudo instruction
//
// TODO: Implement 'li' load immediate instruction. Requires refactor to allow a single assembly instruction
// to be translated to multiple instructions in machine code
func parsePseudo(operator string, opperands []string) (instruction, error) {
	switch operator {
	case "nop":
		if len(opperands) == 0 {
			return instructionR{}, nil
		} else {
			return nil, errors.New("unexpected number of opperands (want 0)")
		}

	case "clear":
		if len(opperands) == 1 {
			reg, err := register.Parse(opperands[0])
			if err != nil {
				return nil, err
			}

			return instructionR{
				funct: rFunctCodes["add"],
				rD:    reg,
				rS:    0,
				rT:    0,
			}, nil
		} else {
			return nil, errors.New("unexpected number of opperands (want 1)")
		}

	case "move":
		if len(opperands) == 2 {
			rD, err := register.Parse(opperands[0])
			if err != nil {
				return nil, err
			}

			rS, err := register.Parse(opperands[1])
			if err != nil {
				return nil, err
			}

			return instructionR{
				funct: rFunctCodes["add"],
				rD:    rD,
				rS:    rS,
				rT:    0,
			}, nil
		} else {
			return nil, errors.New("unexpected number of opperands (want 2)")
		}

	default:
		return nil, nil
	}
}
