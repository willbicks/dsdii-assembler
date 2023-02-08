package main

import (
	"flag"
	"fmt"
	"os"
)

type flags struct {
	inFile      string // input file name
	outFile     string // output file name
	outFmt      string // output format
	nopBuff     uint   // number of nop instructions to include after each instruction
	instruction string // instruction to assemble
}

// printUsage prints help information to the flag output (defaults to stderr).
func printUsage() {
	fmt.Fprint(flag.CommandLine.Output(),
		"Usage of dsdii-assembler:\n",
		"\tdsdii-assembler version\n",
		"\tdsdii-assembler [options...] <instruction>\n\n",
		"Options:\n",
	)
	flag.PrintDefaults()
}

// parseFlags parses the command line flags and returns a flags struct.
//
// If the first argument is "version", the version is printed and the program
// exits with status 0.
func parseFlags() flags {
	var f flags

	flag.Usage = printUsage

	flag.StringVar(&f.inFile, "i", "", "Input file containing assembly instructions. If not set, the instruction parameter should contain the singular instruction to be assembled.")
	flag.StringVar(&f.outFile, "o", "stdout", "Output file to write machine code to.")
	flag.UintVar(&f.nopBuff, "nop-buff", 0, "Optional number of nop instructions to include after each instruction.")

	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		os.Exit(0)
	}
	f.instruction = flag.Arg(0)

	return f
}
