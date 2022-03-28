// package output provides methods for taking machine code and formatting and outputting it to a destination.
package output

type Writer interface {
	WriteStart() error
	WriteInstruction(inst uint32) error
	WriteEnd() error
}
