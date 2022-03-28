package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/willbicks/dsdii-assembler/instruction"
	"github.com/willbicks/dsdii-assembler/output"
)

func main() {
	inFileName := flag.String("i", "", "Input file containing assembly instrucitons.")
	outFileName := flag.String("o", "stdout", "Output file to write machine code to.")
	outFileFmt := flag.String("ofmt", "hex", "Output format (hex, vhdl-byte, vhdl-word).")
	inst := flag.Arg(0)
	flag.Parse()

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

		for inScan.Scan() {
			fmt.Println(inScan.Text())
		}

		if err := inScan.Err(); err != nil {
			log.Fatal(err)
		}
	}

	out.WriteEnd()

	fmt.Println(*inFileName, *outFileName, *outFileFmt, inst)
}
