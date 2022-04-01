package main

import (
	"flag"
	"fmt"
)

type flags struct {
	inFile      string // input file name
	outFile     string // output file name
	outFmt      string // output format
	nopBuff     uint   // number of nop instructions to include after each instruction
	instruction string // instruction to assemble
}

// parseFlags parses the command line flags and returns a flags struct.
func parseFlags() flags {
	var f flags

	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(),
			"Usage of dsdii-assembler:\n",
			"\tdsdii-assembler [options...] [<instruction>]\n\n",
			"Options:\n",
		)

		flag.PrintDefaults()
	}

	flag.StringVar(&f.inFile, "i", "", "Input file containing assembly instrucitons. If not set, the instruction parameter should contain the singular instruciton to be assembled.")
	flag.StringVar(&f.outFile, "o", "stdout", "Output file to write machine code to.")
	flag.StringVar(&f.outFmt, "out-fmt", "hex", "Output format (hex, vhdl-byte, vhdl-word).")
	flag.UintVar(&f.nopBuff, "nop-buff", 0, "Optional number of nop instructions to include after each instruciton.")

	flag.Parse()

	f.instruction = flag.Arg(0)

	return f
}
