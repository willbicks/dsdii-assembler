package instruction

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/willbicks/dsdii-assembler/register"
)

// parenRegex is used to parse immediate opperands with parenthesis. See splitParenOpperand() for more information.
var parenRegex *regexp.Regexp

// parseI parses an I type instruciton with the provided op code and opperands.
func parseI(opcode uint8, opperands []string) (instructionI, error) {
	iI := instructionI{
		op: opcode,
	}

	if len(opperands) < 1 {
		return instructionI{}, fmt.Errorf("missing opperands from instruction")
	}

	// parse the destination register rT (always first)
	var err error
	iI.rT, err = register.Parse(opperands[0])
	if err != nil {
		return instructionI{}, err
	}

	var rSStr, immStr string
	if opcode == iOpCodes["lw"] || opcode == iOpCodes["sw"] {
		// split the opperands of load and store instructions, which have special parenthesis syntax
		if len(opperands) != 2 {
			return instructionI{}, fmt.Errorf("unexpected opperands (want 3 w/ paren syntax)")
		}

		immStr, rSStr, err = splitParenOpperand(opperands[1])
		if err != nil {
			return instructionI{}, err
		}
	} else {
		// otherwise, find rS and imm strings in opperands slice
		if len(opperands) != 3 {
			return instructionI{}, fmt.Errorf("unexpected number of opperands (want 3)")
		}

		rSStr = opperands[1]
		immStr = opperands[2]
	}

	iI.rS, err = register.Parse(rSStr)
	if err != nil {
		return instructionI{}, err
	}

	i64, err := strconv.ParseInt(immStr, 10, 16)
	if err != nil {
		return instructionI{}, fmt.Errorf("unable to parse immediate: %w", err)
	}
	iI.imm = int16(i64)

	return iI, nil
}

// splitParenOpperand splits an opperand string with form a(b) into its component a and b pieces using a regular expression.
func splitParenOpperand(str string) (a string, b string, err error) {
	// compile regex if not allready compiled
	if parenRegex == nil {
		// this regular expression is intentially left broad so that invalid substrings are handled by
		// descriptive error checking elsewhere instead of not providing a match at all
		parenRegex = regexp.MustCompile(`(.*)\((.*)\)`)
	}

	res := parenRegex.FindStringSubmatch(str)

	if len(res) != 3 {
		return "", "", fmt.Errorf("unable to parse parenthesis syntax in opperand: %s", str)
	}

	return res[1], res[2], nil
}
