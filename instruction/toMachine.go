package instruction

const (
	opBits    = 6
	regBits   = 5
	shiftBits = 5
	functBits = 6
	immBits   = 16
)

const (
	opMask    = 1<<opBits - 1
	regMask   = 1<<regBits - 1
	shiftMask = 1<<shiftBits - 1
	functMask = 1<<functBits - 1
	immMask   = 1<<immBits - 1
)

// toMachine convers an R type instruciton to 32 bit machine code, made up of each of it's members
// packed with bit widths specified by constants.
func (i instructionR) toMachine() uint32 {
	var u uint32

	u |= uint32(i.rS) & regMask
	u <<= regBits
	u |= uint32(i.rT) & regMask
	u <<= regBits
	u |= uint32(i.rD) & regMask
	u <<= shiftBits
	u |= uint32(i.shiftAmt) & shiftMask
	u <<= functBits
	u |= uint32(i.funct) & functMask

	return u
}

// toMachine convers an I type instruciton to 32 bit machine code, made up of each of it's members
// packed with bit widths specified by constants.
func (i instructionI) toMachine() uint32 {
	var u uint32

	u |= uint32(i.op) & opMask
	u <<= regBits
	u |= uint32(i.rS) & regMask
	u <<= regBits
	u |= uint32(i.rT) & regMask
	u <<= immBits
	u |= uint32(i.imm) & immMask

	return u
}
