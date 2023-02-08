package output

import (
	"fmt"
	"io"
)

// Binary output formatter writes instructions as formatted binary text
// (1s and 0s), optionally with spaces used to separate every spaceAfter bits.
//
// If spaceAfter is 0, no spaces are added.
type Binary struct {
	dest       io.Writer
	spaceAfter int8
}

// NewBinary returns a Binary output formatter with specified destination
// and spaceAfter. spaceAfter is the number of bits to separate with a space, or
// 0 to disable spaces.
func NewBinary(dest io.Writer, spaceAfter int8) Binary {
	return Binary{
		dest:       dest,
		spaceAfter: spaceAfter,
	}
}

var _ Writer = Binary{}

// no action required for Binary output format
func (o Binary) WriteStart(comment string) error {
	return nil
}

// WriteInstruction writes instruction to destination as a binary string
// with spaces every spaceAfter bits.
//
// Ignores the comment parameter if set.
func (o Binary) WriteInstruction(inst uint32, comment string) error {
	bits := fmt.Sprintf("%032b\n", inst)
	res := ""

	if o.spaceAfter > 0 {
		for i := int8(0); i < 32/o.spaceAfter; i++ {
			res += bits[i*o.spaceAfter:(i+1)*o.spaceAfter] + " "
		}
		res = res[:len(res)-1] + "\n"
	} else {
		res = bits
	}

	_, err := fmt.Fprint(o.dest, res)
	return err
}

// no action required for Binary output format
func (o Binary) WriteEnd() error {
	return nil
}
