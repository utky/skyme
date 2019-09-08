package ebpf

import "golang.org/x/sys/unix"

// Value including reference to register or immediate value
type Value interface {
	source() byte
}

// Reg mean register
type Reg struct {
	nr byte // actually 4 bit
}

func (r Reg) source() byte {
	return unix.BPF_X
}

// R0 is Register 0
func R0() Reg { return Reg{0} }

// R1 is Register 1
func R1() Reg { return Reg{1} }

// R2 is Register 2
func R2() Reg { return Reg{2} }

// R3 is Register 3
func R3() Reg { return Reg{3} }

// R4 is Register 4
func R4() Reg { return Reg{4} }

// R5 is Register 5
func R5() Reg { return Reg{5} }

// R6 is Register 6
func R6() Reg { return Reg{6} }

// R7 is Register 7
func R7() Reg { return Reg{7} }

// R8 is Register 8
func R8() Reg { return Reg{8} }

// R9 is Register 9
func R9() Reg { return Reg{9} }

// Imm means immediate value
type Imm struct {
	val uint32
}

// Im as immediate value
func Im(v uint32) Imm { return Imm{v} }

func (i Imm) source() byte {
	return unix.BPF_K
}

func alujmp(class byte, op byte, dest Reg, offset uint16, value Value) Inst {
	var src Reg
	var imm Imm
	if value.source() == unix.BPF_X {
		src = value.(Reg)
		imm = Imm{0}
	} else { // assuming source type : immediate unix.BPF_X
		src = R0()
		imm = value.(Imm)
	}
	regbit := (dest.nr << 4) | src.nr
	opcode := op | value.source() | class
	return Inst{opcode, regbit, offset, imm.val}
}

func alu(op byte, dest Reg, value Value) Inst {
	return alujmp(unix.BPF_ALU, op, dest, 0, value)
}

func jmp(op byte, offset uint16, value Value) Inst {
	return alujmp(unix.BPF_JMP, op, R0(), offset, value)
}

// Add is instruction for BPF_ADD
func Add(dest Reg, value Value) Inst {
	return alu(unix.BPF_ADD, dest, value)
}

// Mov is instruction for BPF_MOV
func Mov(dest Reg, value Value) Inst {
	return alu(unix.BPF_MOV, dest, value)
}

// Exit is instruction for BPF_EXIT
func Exit() Inst {
	return jmp(unix.BPF_EXIT, 0, R0())
}
