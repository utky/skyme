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
type Prog struct {
	ProgType int
	Insts    []Inst
}

// NewProg creates a new instance of BPF program
func NewProg(progType int) *Prog {
	is := make([]Inst, 0)
	return &Prog{progType, is}
}

// Push appends instruction to tail of array of instruction
func (p *Prog) Push(i Inst) {
	p.Insts = append(p.Insts, i)
}

func (p *Prog) Write(w io.Writer) error {
	for _, i := range p.Insts {
		if err := i.Write(w); err != nil {
			return err
		}
	}

	return nil
}

// Map is generic ke-value store to interact with eBPF program
type Map interface {
}
