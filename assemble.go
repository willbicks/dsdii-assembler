package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/willbicks/dsdii-assembler/instruction"
	"github.com/willbicks/dsdii-assembler/output"
)

type config struct {
	in      io.Reader     // input reader
	out     output.Writer // output writer
	nopBuff uint          // number of nop instructions to include after each instruction
}

// assemble a series of instructions using the provided configuration, and return the number of lines assembled, and an error if one occured.
func assemble(c config) (lines uint64, err error) {
	if err := c.out.WriteStart(); err != nil {
		return 0, fmt.Errorf("writing start: %v", err)
	}

	var line uint64
	var mc uint32
	inScan := bufio.NewScanner(c.in)
	for inScan.Scan() {
		line++

		// assemble and write instruction
		mc, err = instruction.Assemble(inScan.Text())
		if err != nil {
			return 0, fmt.Errorf("on line %d: %s", line, err)
		}
		if err := c.out.WriteInstruction(mc); err != nil {
			return 0, fmt.Errorf("writing instruction: %v", err)
		}

		// add nop buffer as configured
		for i := uint(0); i < c.nopBuff; i++ {
			if err := c.out.WriteInstruction(0); err != nil {
				return 0, fmt.Errorf("writing nop: %v", err)
			}
		}
	}

	if err := inScan.Err(); err != nil {
		return 0, fmt.Errorf("scan: %v", err)
	}

	if err := c.out.WriteEnd(); err != nil {
		return 0, fmt.Errorf("writing end: %v", err)
	}

	return line, nil
}
