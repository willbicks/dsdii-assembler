package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

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
		var err error
		dest, err = os.Create(*outFileName)
		if err != nil {
			log.Panicf("unable to open output file: %s", err)
		}
		defer func() {
			if err := dest.Sync(); err != nil {
				fmt.Println(err)
			}
			if err := dest.Close(); err != nil {
				fmt.Println(err)
			}
		}()
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

	start := time.Now()
	if err := out.WriteStart(); err != nil {
		log.Fatalf("writing start: %v", err)
	}
	var line uint64
	fmt.Println("dsdii-assembler")

	if *inFileName == "" {
		line = 1
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
		for inScan.Scan() {
			line++
			// assemble and write instruction
			mc, err = instruction.Assemble(inScan.Text())
			if err != nil {
				log.Fatalf("error on line %d: %s", line, err)
			}
			if err := out.WriteInstruction(mc); err != nil {
				log.Fatalf("writing instruction: %v", err)
			}

			// add nop buffer as configured
			for i := uint(0); i < *nopBuff; i++ {
				if err := out.WriteInstruction(0); err != nil {
					log.Fatalf("writing nop: %v", err)
				}
			}
		}

		if err := inScan.Err(); err != nil {
			log.Fatalf("scan error: %v", err)
		}
	}

	if err := out.WriteEnd(); err != nil {
		log.Fatalf("writing end: %v", err)
	}

	fmt.Printf("asssembled %d line(s) in %d ms\n", line, time.Now().Sub(start).Milliseconds())
	return
}
