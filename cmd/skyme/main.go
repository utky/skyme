package main

import (
	"fmt"
	"os"

	"github.com/utky/skyme/pkg/ebpf"
)

func main() {
	p := ebpf.NewArrayProg()
	p.Push(ebpf.Mov(ebpf.R0(), ebpf.Im(1)))
	p.Push(ebpf.Exit())
	if err := p.Write(os.Stdout); err != nil {
		fmt.Printf("Write failed")
		os.Exit(1)
	}
	os.Exit(0)
}
