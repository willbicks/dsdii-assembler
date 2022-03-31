# ⚙ DSDII Assembler

DSDII Assembler parses MIPS instructions and generates machine code in various formats for a processor developed in RIT's Digital Systems Design II class. The class focuses on the VHDL and Verilog implementation of the processor itself, and this exercise is purely a personal side project.

## Features

#### Instruction Parsing:
- [x] R type instructions
- [x] I type instructions
- [ ] J type instructions (not implemented in our processor)
- [x] Pseudo instructions (`nop`, `clear`, `move`)
- [ ] Load immediate (pseudo `li`)
- [ ] Support for comments and blank lines
#### Output Formatting:
- [x] Raw hex machine code output
- [x] VHDL code generation (byte addressable memory)
- [ ] VHDL code generation (word addressable memory)

## Usage

```shell
$ dsdii-assembler --help

Usage of dsdii-assembler:
  -i string
        Input file containing assembly instrucitons.
  -nop-buff uint
        Optional number of nop instructions to include after each instruciton.
  -o string
        Output file to write machine code to. (default "stdout")
  -out-fmt string
        Output format (hex, vhdl-byte, vhdl-word). (default "hex")
```

## Examples

```shell
$ dsdii-assembler 'add $s0, $s1, $t3' 
dsdii-assembler
022b8020
asssembled 1 line(s) in 0 ms
```

```shell
$ dsdii-assembler .\test.asm -out-fmt vhdl-byte -nop-buff 4
dsdii-assembler
signal instructions : mem_array := (
      x"20", x"10", x"00", x"1f",
      x"00", x"00", x"00", x"00",
      x"00", x"00", x"00", x"00",
      x"00", x"00", x"00", x"00",
      x"00", x"00", x"00", x"00",
      x"20", x"11", x"00", x"03",
      ...
      x"00", x"00", x"00", x"00",
      x"8e", x"55", x"00", x"00",
      x"00", x"00", x"00", x"00",
      x"00", x"00", x"00", x"00",
      x"00", x"00", x"00", x"00",
      x"00", x"00", x"00", x"00",
      others => x"00"
);
asssembled 35 line(s) in 21 ms
```

## Contributions

Contributions are welcome via pull request, as well as discussions via issues. Please ensure that all contributions are appropriately documented and tested.