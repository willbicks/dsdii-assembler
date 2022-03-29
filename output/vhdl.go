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

func (o VHDLByte) WriteStart() error {
	_, err := fmt.Fprintln(o.dest, "signal instructions : mem_array := (")
	return err
}

func (o VHDLByte) WriteInstruction(inst uint32) error {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, inst)

	_, err := fmt.Fprintf(o.dest, "\tx\"%02x\", x\"%02x\", x\"%02x\", x\"%02x\",\n", bs[0], bs[1], bs[2], bs[3])
	return err
}

func (o VHDLByte) WriteEnd() error {
	_, err := fmt.Fprint(o.dest, "\tothers => x\"00\"\n);")
	return err
}
