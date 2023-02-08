# ⚙ DSDII Assembler

DSDII Assembler parses MIPS instructions and generates machine code in various formats for a processor developed in RIT's Digital Systems Design II class. The class focuses on the VHDL and Verilog implementation of the processor itself, and this exercise is purely a personal side project.

## Features

### Instruction Parsing:

- [x] R type instructions
- [x] I type instructions
- [ ] J type instructions (not implemented in our processor)
- [x] Pseudo instructions (`nop`, `clear`, `move`)
- [ ] Load immediate (pseudo `li`)
- [x] Support for comments and blank lines

### Code Generation:

- [x] Raw hex machine code output
- [x] VHDL code generation (byte addressable memory)
- [x] VHDL code generation (word addressable memory)
- [x] Generate VHDL comments with assembly instructions

## Usage

```shell
$ dsdii-assembler -help
Usage of dsdii-assembler:
        dsdii-assembler version
        dsdii-assembler [options...] <instruction>

Options:
  -i string
        Input file containing assembly instructions. If not set, the instruction parameter should contain the singular instruction to be assembled.
  -nop-buff uint
        Optional number of nop instructions to include after each instruction.
  -o string
        Output file to write machine code to. (default "stdout")
  -out-fmt string
        Output format (hex, vhdl-byte, vhdl-word). (default "hex")
```

## Installation

### Without Go Installed:

If you don't have Go installed on your computer, pre-compiled binaries for Windows, macOS, and Linux can be downloaded from the [releases section](https://github.com/willbicks/dsdii-assembler/releases).

Download the appropriate binary for your system, extract it, and add it your path or use it in-situ.

### With Go Installed:

If you have the Go language tools installed on your computer, downloading and installing dsdii-assembler is as simple as:

```shell
$ go install github.com/willbicks/dsdii-assembler@latest
```

This will download, compile, and install the binary in your Go path.

## Examples

```shell
$ dsdii-assembler 'add $s0, $s1, $t3'
dsdii-assembler v0.3.2
022b8020
assembled 1 line(s) in 0 ms
```

```shell
$ dsdii-assembler -i .\test.asm -out-fmt vhdl-byte
dsdii-assembler v0.3.2
-- generated by dsdii-assembler
signal instructions : mem_array := (
      x"00", x"00", x"00", x"00", -- nop
      x"00", x"00", x"00", x"00", -- nop
      x"00", x"00", x"00", x"00", -- nop
      x"00", x"00", x"00", x"00", -- nop
      x"20", x"10", x"00", x"1f", -- addi $s0, $0, 31 # assembly comments included in VHDL
      x"20", x"11", x"00", x"03", -- addi $s1, $0, 3
      x"20", x"12", x"00", x"04", -- addi $s2, $0, 4
      ...
      x"00", x"00", x"00", x"00", -- nop
      x"02", x"11", x"40", x"20", -- add $t0, $s0, $s1
      x"02", x"11", x"48", x"24", -- and $t1, $s0, $s1
      x"02", x"11", x"50", x"19", -- multu $t2, $s0, $s1
      others => x"00"
);
assembled 37 line(s) in 9 ms
```

## Contributions

Contributions are welcome via pull request, as well as discussions via issues. Please ensure that all contributions are appropriately documented and tested.
