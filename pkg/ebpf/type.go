package ebpf

import (
	"encoding/binary"
	"io"
)

// Inst means eBPF instruction
type Inst struct {
	op     byte   // opcode
	reg    byte   // 4 bit dst reg, 4 bit src reg
	offset uint16 // offset for jump
	imm    uint32 // immidiate
}

func (i Inst) Write(writer io.Writer) error {
	if err := binary.Write(writer, binary.LittleEndian, i.op); err != nil {
		return err
	}
	if err := binary.Write(writer, binary.LittleEndian, i.reg); err != nil {
		return err
	}
	if err := binary.Write(writer, binary.LittleEndian, i.offset); err != nil {
		return err
	}
	if err := binary.Write(writer, binary.LittleEndian, i.imm); err != nil {
		return err
	}
	return nil
}

// Prog contains set of instructions
type Prog interface {
	ProgType() int
	Insts() []Inst
}

// ArrayProg builds array of instructions
type ArrayProg struct {
	progType int
	insts    []Inst
}

func NewArrayProg(progType int) *ArrayProg {
	is := make([]Inst, 0)
	return &ArrayProg{progType, is}
}

func (p *ArrayProg) Push(i Inst) {
	p.insts = append(p.insts, i)
}

func (p *ArrayProg) Write(w io.Writer) error {
	for _, i := range p.insts {
		if err := i.Write(w); err != nil {
			return err
		}
	}

	return nil
}

// Map is generic ke-value store to interact with eBPF program
type Map interface {
}
