package output

import (
	"fmt"
	"io"
)

// Hex output formatter generates an ASCII encoded file with hex representations
// of each instruction, seperated by new lines
type Hex struct {
	dest io.Writer
}

// NewHex returns a Hex output formatter with specified destination
func NewHex(dest io.Writer) Hex {
	return Hex{
		dest: dest,
	}
}

var _ Writer = Hex{}

func (o Hex) WriteStart(comment string) error {
	return nil
}

func (o Hex) WriteInstruction(inst uint32, comment string) error {
	_, err := fmt.Fprintf(o.dest, "%08x\n", inst)
	return err
}

func (o Hex) WriteEnd() error {
	return nil
}
