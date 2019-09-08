package ebpf

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

type ProgId = uintptr

// ProgLoadAttr is control parameters when you load eBPF program
type ProgLoadAttr struct { /* Used by BPF_PROG_LOAD */
	ProgType    uint32
	IsnsCnt     uint32
	Insns       []Inst
	License     []byte
	LogLevel    uint32
	LogSize     uint32
	LogBuf      []byte
	KernVersion uint32
}

// Load accepts BPF program
func Load(prog Prog) (ProgId, error) {
	logBuf := make([]byte, 1024)
	//	var progidptr uintptr
	attr := ProgLoadAttr{
		unix.BPF_PROG_TYPE_SOCKET_FILTER,
		uint32(len(prog.Insts())),
		prog.Insts(),
		[]byte("GPL"),
		1,
		1024,
		logBuf,
	}
	progidptr, _, errno := unix.Syscall(
		unix.SYS_BPF, unix.BPF_PROG_LOAD,
		uintptr(unsafe.Pointer(&attr)),
		uintptr(unsafe.Sizeof(attr)))
	if errno < 0 {
		return 0, fmt.Errorf("Failed on calling bpf: %s", errno.Error())
	}
	return progidptr, nil
}
