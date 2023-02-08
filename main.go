package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/willbicks/dsdii-assembler/output"
)

// generateConfig generates a config struct with the input source, output destination, and additional parameters from
// provided flag set.
//
// returns the generate config, a closer function to close the output destination, and any errors encountered while
// parsing the configuration.
func generateConfig(f flags) (c config, closer func(), err error) {
	cfg := config{
		nopBuff: f.nopBuff,
	}

	// set cfg.in (input reader)
	if f.inFile == "" {
		if len(f.instruction) != 0 {
			cfg.in = strings.NewReader(f.instruction)
		} else {
			return config{}, func() {}, fmt.Errorf("Neither an input file nor a singular instruction to assemble were provided. Need help? Try `dsdii-assembler -help`.")
		}
	} else {
		var err error
		cfg.in, err = os.Open(f.inFile)
		if err != nil {
			return config{}, func() {}, fmt.Errorf("opening input file: %s", err)
		}
	}

	// set cfg.out (output writer)
	var dest *os.File
	if f.outFile == "stdout" {
		dest = os.Stdout
	} else {
		var err error
		dest, err = os.Create(f.outFile)
		if err != nil {
			return config{}, func() {}, fmt.Errorf("unable to open output file: %s", err)
		}
	}

	close := func() {
		if err := dest.Close(); err != nil {
			fmt.Println(err)
		}
	}

	switch f.outFmt {
	case "hex":
		cfg.out = output.NewHex(dest)
	case "vhdl-byte":
		cfg.out = output.NewVHDLByte(dest)
	case "vhdl-word":
		cfg.out = output.NewVHDLWord(dest)
	case "binary":
		cfg.out = output.NewBinary(dest, 0)
	case "binary-nibble":
		cfg.out = output.NewBinary(dest, 4)
	case "binary-byte":
		cfg.out = output.NewBinary(dest, 8)
	default:
		log.Fatal("Invalid output format. Want one of: hex, vhdl-byte, vhdl-word, binary, binary-nibble, binary-byte.")
	}

	return cfg, close, nil
}

func main() {
	parseVersionInfo()
	flags := parseFlags()

	config, close, err := generateConfig(flags)
	if err != nil {
		fmt.Printf("\033[31mERROR:\033[0m %s", err)
		return
	}
	defer close()

	fmt.Printf("dsdii-assembler %v\n", version)
	start := time.Now()

	lines, err := assemble(config)
	if err != nil {
		fmt.Printf("\033[31mERROR:\033[0m %s", err)
		return
	}

	fmt.Printf("assembled %d line(s) in %d ms\n", lines, time.Since(start).Milliseconds())
}
