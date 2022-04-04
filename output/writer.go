// package output provides methods for taking machine code and formatting and outputting it to a destination.
package output

type Writer interface {
	// WriteStart writes an introduciton / header to the destination before any instructions are written
	WriteStart(comment string) error
	// WriteInstruction writes an instruction to the destination with the format implemented by the Writer
	WriteInstruction(inst uint32, comment string) error
	// WriteEnd writes an end / footer to the destination after all instructions are written
	WriteEnd() error
}
