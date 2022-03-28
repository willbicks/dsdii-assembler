# ⚙ DSDII Assembler

DSDII Assembler parses MIPS instructions and generates machine code in various formats for a processor developed in RIT's Digital Systems Design II class. The class focuses on the VHDL and Verilog implementation of the processor itself, and this exercise is purely a personal side project.

## Usage

```shellp
$ dsdii-assembler --help

Usage of dsdii-assembler:
  -i string
        Input file containing assembly instrucitons.
  -o string
        Output file to write machine code to. (default "stdout")
  -ofmt string
        Output format (hex, vhdl-byte, vhdl-word). (default "hex")
```