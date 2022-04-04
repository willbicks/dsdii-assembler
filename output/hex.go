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

// not required for Hex output format
func (o Hex) WriteStart(comment string) error {
	return nil
}

// WriteInstruction writes instruction to destination as an 8 character hex string
// ignores the comment parameter if set
func (o Hex) WriteInstruction(inst uint32, comment string) error {
	_, err := fmt.Fprintf(o.dest, "%08x\n", inst)
	return err
}

// not required for Hex output format
func (o Hex) WriteEnd() error {
	return nil
}
