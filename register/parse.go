package register

import (
	"fmt"
	"strconv"
)

// Parse returns the register number given a register name. Uses the
// map technique.
func Parse(regName string) (uint8, error) {
	return parseMap(regName)
}

// parseMap returns the register number given a register name.
// Uses a map lookup with the register map.
func parseMap(regName string) (uint8, error) {
	if len(regName) < 2 {
		return 0, fmt.Errorf("invalid register %s", regName)
	}

	if regName[0] != '$' {
		return 0, fmt.Errorf("invalid register %s, must start with $", regName)
	}
	i, ok := registers[regName[1:]]
	if !ok {
		return 0, fmt.Errorf("invalid register %s", regName)
	}
	return i, nil
}

// parseConditional returns the register number given a register name.
// Uses uses a series of conditionals to find / calculate the register number.
//
// While not used by the public Parse function, it is left in place for benchmarking,
// testing, and posterity purposes.
func parseConditional(regName string) (uint8, error) {
	if len(regName) < 2 {
		return 0, fmt.Errorf("invalid register %s", regName)
	}

	if regName[0] != '$' {
		return 0, fmt.Errorf("invalid register %s, must start with $", regName)
	}

	// return zero register
	if regName == "$0" {
		return 0, nil
	}

	// if not zero register, check that is three characters
	if len(regName) != 3 {
		return 0, fmt.Errorf("invalid register %s, must be three characters", regName)
	}

	// return if reg is special
	switch regName[1:] {
	case "at":
		return 1, nil
	case "gp":
		return 28, nil
	case "sp":
		return 29, nil
	case "fp":
		return 30, nil
	case "ra":
		return 31, nil
	}

	// return general purpose registers
	errInvalidRegNum := fmt.Errorf("invalid register number %s", regName)
	num64, err := strconv.ParseUint(string(regName[2]), 10, 8)
	if err != nil {
		return 0, errInvalidRegNum
	}
	num := uint8(num64)

	switch regName[1] {
	case 'v':
		if num <= 1 {
			return num + 2, nil
		} else {
			return 0, errInvalidRegNum
		}
	case 'a':
		if num <= 3 {
			return num + 4, nil
		} else {
			return 0, errInvalidRegNum
		}
	case 't':
		if num <= 7 {
			return num + 8, nil
		} else if num <= 9 {
			return num + (24 - 8), nil
		} else {
			return 0, errInvalidRegNum
		}
	case 's':
		if num <= 7 {
			return num + 16, nil
		} else {
			return 0, errInvalidRegNum
		}
	case 'k':
		if num <= 1 {
			return num + 26, nil
		} else {
			return 0, errInvalidRegNum
		}
	}

	return 0, fmt.Errorf("unkown register %s", regName)
}
