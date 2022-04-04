package output

import (
	"encoding/binary"
	"fmt"
	"io"
)

// VHDLByte output formatter generates VHDL code for a byte atressable memory array
type VHDLByte struct {
	dest io.Writer
}

// NewVHDLByte createas a new VHDLByte output formatter
func NewVHDLByte(dest io.Writer) VHDLByte {
	return VHDLByte{
		dest: dest,
	}
}

var _ Writer = VHDLByte{}

// WriteStart writes the VHDL array opener
func (o VHDLByte) WriteStart(comment string) error {
	_, err := fmt.Fprintf(o.dest, "-- %s\nsignal instructions : mem_array := (\n", comment)
	return err
}

// WriteInstruction writes instruction to destination as 4 two character hex strings with VHDL quotes and commas,
// and a VHDL comment containing the comment parameter
func (o VHDLByte) WriteInstruction(inst uint32, comment string) error {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, inst)

	_, err := fmt.Fprintf(o.dest, "\tx\"%02x\", x\"%02x\", x\"%02x\", x\"%02x\", -- %s\n", bs[0], bs[1], bs[2], bs[3], comment)
	return err
}

// WriteEnd writes the VHDL array closer for byte adressable memory
func (o VHDLByte) WriteEnd() error {
	_, err := fmt.Fprint(o.dest, "\tothers => x\"00\"\n);\n")
	return err
}

// VHDLWord output formatter generates VHDL code for a word adressable atressable memory array
type VHDLWord struct {
	dest io.Writer
}

// NewVHDLWord createas a new VHDLWord output formatter
func NewVHDLWord(dest io.Writer) VHDLWord {
	return VHDLWord{
		dest: dest,
	}
}

var _ Writer = VHDLWord{}

// WriteStart writes the VHDL array opener
func (o VHDLWord) WriteStart(comment string) error {
	_, err := fmt.Fprintf(o.dest, "-- %s\nsignal instructions : mem_array := (\n", comment)
	return err
}

// WriteInstruction writes instruction to destination as an 8 character hex strings with VHDL quotes and commas,
// and a VHDL comment containing the comment parameter
func (o VHDLWord) WriteInstruction(inst uint32, comment string) error {
	_, err := fmt.Fprintf(o.dest, "\tx\"%08x\", -- %s\n", inst, comment)
	return err
}

// WriteEnd writes the VHDL array closer for word adressable memory
func (o VHDLWord) WriteEnd() error {
	_, err := fmt.Fprint(o.dest, "\tothers => x\"00000000\"\n);\n")
	return err
}
