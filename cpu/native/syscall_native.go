package native

import (
	"syscall"
	"unsafe"
)

const (
	CLOCKS_PER_SEC = 1000000
	CPU_FREQUENCY = 300000
)

func Getuid() int {
	return syscall.Getuid()
}

func Geteuid() int {
	return syscall.Geteuid()
}

func Getgid() int {
	return syscall.Getgid()
}

func Getegid() int {
	return syscall.Getegid()
}

func Read(fd int, buf []byte) int {
	var count, _ = syscall.Read(fd, buf)
	return count
}

func Write(fd int, buf []byte) int {
	var count, _ = syscall.Write(fd, buf)
	return count
}

func Open(path string, mode int, perm uint32) int {
	var fd, _ = syscall.Open(path, mode, perm)
	return fd
}

func Close(fd int) int {
	syscall.Close(fd)
	panic("Unimplemented")
	return 0 //TODO
}

func Clock(numCycles int) int {
	return CLOCKS_PER_SEC * numCycles / CPU_FREQUENCY
}

func Seek(fd int, offset int64, whence int) int64 {
	var off, _ = syscall.Seek(fd, offset, whence)
	return off
}

func Ioctl(fd int, request int, buf []byte) int64 {
	result, _, _ := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(request), uintptr(unsafe.Pointer(&buf[0])))
	return int64(result)
}
