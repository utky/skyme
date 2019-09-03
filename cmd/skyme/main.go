package main

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

func main() {
	var instructions []byte
	r1, r2, err := unix.Syscall(unix.SYS_BPF, unix.BPF_PROG_LOAD, uintptr(unsafe.Pointer(&instructions)), 0)
	fmt.Printf("r1 = %v, r2 = %v, lastErr = %v\n", r1, r2, err)
	fmt.Printf("Error: %s", err.Error())
	os.Exit(0)
}
