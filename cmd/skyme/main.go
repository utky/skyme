package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"

	"github.com/utky/skyme/pkg/ebpf"
)

func main() {
	p := ebpf.NewProg(unix.BPF_PROG_TYPE_SOCKET_FILTER)
	p.Push(ebpf.Mov(ebpf.R0, ebpf.I(1)))
	p.Push(ebpf.Exit())
	if err := p.Write(os.Stdout); err != nil {
		fmt.Printf("Write failed")
		os.Exit(1)
	}
	os.Exit(0)
}
