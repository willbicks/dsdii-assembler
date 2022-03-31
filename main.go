package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/willbicks/dsdii-assembler/output"
)

// generateConfig parses command flags and arguments, and generates a config struct with the input source, output destination, and additoinal parameters.
//
// returns the generate config, a closer function to close the output destination, and any errors encountered while parsing the configuration.
func generateConfig() (c config, closer func(), err error) {
	var cfg config

	inFileName := flag.String("i", "", "Input file containing assembly instrucitons.")
	outFileName := flag.String("o", "stdout", "Output file to write machine code to.")
	outFileFmt := flag.String("out-fmt", "hex", "Output format (hex, vhdl-byte, vhdl-word).")
	flag.UintVar(&cfg.nopBuff, "nop-buff", 0, "Optional number of nop instructions to include after each instruciton.")
	flag.Parse()
	inst := flag.Arg(0)

	if *inFileName == "" {
		if len(inst) != 0 {
			cfg.in = strings.NewReader(inst)
		} else {
			return config{}, func() {}, fmt.Errorf("neither an input file nor a singular instruction to assemble were provided")
		}
	} else {
		var err error
		cfg.in, err = os.Open(*inFileName)
		if err != nil {
			return config{}, func() {}, fmt.Errorf("opening input file: %s", err)
		}
	}

	var dest *os.File
	if *outFileName == "stdout" {
		dest = os.Stdout
	} else {
		var err error
		dest, err = os.Create(*outFileName)
		if err != nil {
			return config{}, func() {}, fmt.Errorf("unable to open output file: %s", err)
		}
	}

	close := func() {
		if err := dest.Close(); err != nil {
			fmt.Println(err)
		}
	}

	switch *outFileFmt {
	case "hex":
		cfg.out = output.NewHex(dest)
	case "vhdl-byte":
		cfg.out = output.NewVHDLByte(dest)
	case "vhdl-word":
	default:
		log.Fatal("Invalid output format. Want one of: hex, vhdl-byte, vhdl-word.")
	}

	return cfg, close, nil
}

func main() {
	start := time.Now()
	fmt.Println("dsdii-assembler")

	config, close, err := generateConfig()
	if err != nil {
		fmt.Printf("\033[31mERROR:\033[0m %s", err)
		return
	}
	defer close()

	lines, err := assemble(config)
	if err != nil {
		fmt.Printf("\033[31mERROR:\033[0m %s", err)
		return
	}

	fmt.Printf("asssembled %d line(s) in %d ms\n", lines, time.Now().Sub(start).Milliseconds())
}
