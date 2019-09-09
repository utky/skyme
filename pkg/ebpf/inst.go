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

// R0 is Register 0, rax in x86_64
var R0 = Reg{0}

// R1 is Register 1, rdi in x86_64
var R1 = Reg{1}

// R2 is Register 2, rsi in x86_64
var R2 = Reg{2}

// R3 is Register 3, rdx in x86_64
var R3 = Reg{3}

// R4 is Register 4, rcx in x86_64
var R4 = Reg{4}

// R5 is Register 5, r8  in x86_64
var R5 = Reg{5}

// R6 is Register 6, rdx in x86_64
var R6 = Reg{6}

// R7 is Register 7, r13 in x86_64
var R7 = Reg{7}

// R8 is Register 8, r14 in x86_64
var R8 = Reg{8}

// R9 is Register 9, r15 in x86_64
var R9 = Reg{9}

// R10 is Register 10, rbp in x86_64
var R10 = Reg{10}

// Imm means immediate value
type Imm struct {
	val uint32
}

// I as immediate value
func I(v uint32) Imm { return Imm{v} }

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
		src = R0
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
	return alujmp(unix.BPF_JMP, op, R0, offset, value)
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
	return jmp(unix.BPF_EXIT, 0, R0)
}
