package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/willbicks/dsdii-assembler/instruction"
	"github.com/willbicks/dsdii-assembler/output"
)

func main() {
	inFileName := flag.String("i", "", "Input file containing assembly instrucitons.")
	outFileName := flag.String("o", "stdout", "Output file to write machine code to.")
	outFileFmt := flag.String("out-fmt", "hex", "Output format (hex, vhdl-byte, vhdl-word).")
	nopBuff := flag.Uint("nop-buff", 0, "Optional number of nop instructions to include after each instruciton.")
	flag.Parse()
	inst := flag.Arg(0)

	var dest *os.File
	if *outFileName == "stdout" {
		dest = os.Stdout
	} else {
		dest, err := os.Create(*outFileName)
		if err != nil {
			log.Panicf("unable to open output file: %s", err)
		}
		defer dest.Close()
	}

	var out output.Writer
	switch *outFileFmt {
	case "hex":
		out = output.NewHex(dest)
	case "vhdl-byte":
		out = output.NewVHDLByte(dest)
	case "vhdl-word":
	default:
		log.Fatal("Invalid output format. Want one of: hex, vhdl-byte, vhdl-word.")
	}

	out.WriteStart()

	if *inFileName == "" {
		mc, err := instruction.Assemble(inst)
		if err != nil {
			log.Fatalf("parsing instruction argument: %s", err)
		}
		out.WriteInstruction(mc)
	} else {
		inFile, err := os.Open(*inFileName)
		if err != nil {
			log.Fatalf("opening input file: %s", err)
		}
		inScan := bufio.NewScanner(inFile)

		var mc uint32
		var line uint64
		for inScan.Scan() {
			line++
			// assemble and write instruction
			mc, err = instruction.Assemble(inScan.Text())
			if err != nil {
				log.Fatalf("error on line %d: %s", line, err)
			}
			out.WriteInstruction(mc)

			// add nop buffer as configured
			for i := uint(0); i <= *nopBuff; i++ {
				out.WriteInstruction(0)
			}
		}

		if err := inScan.Err(); err != nil {
			log.Fatal(err)
		}
	}

	out.WriteEnd()
}
