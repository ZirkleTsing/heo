package os

import (
	"fmt"
	//"math"
	"github.com/mcai/acogo/cpu/regs"
	"github.com/mcai/acogo/cpu"
)

type ErrNo uint32

const (
	EPERM ErrNo = 1

	ENOENT ErrNo = 2

	ESRCH ErrNo = 3

	EINTR ErrNo = 4

	EIO ErrNo = 5

	ENXIO ErrNo = 6

	E2BIG ErrNo = 7

	ENOEXEC ErrNo = 8

	EBADF ErrNo = 9

	ECHILD ErrNo = 10

	EAGAIN ErrNo = 11

	ENOMEM ErrNo = 12

	EACCES ErrNo = 13

	EFAULT ErrNo = 14

	ENOTBLK ErrNo = 15

	EBUSY ErrNo = 16

	EEXIST ErrNo = 17

	EXDEV ErrNo = 18

	ENODEV ErrNo = 19

	ENOTDIR ErrNo = 20

	EISDIR ErrNo = 21

	EINVAL ErrNo = 22

	ENFILE ErrNo = 23

	EMFILE ErrNo = 24

	ENOTTY ErrNo = 25

	ETXTBSY ErrNo = 26

	EFBIG ErrNo = 27

	ENOSPC ErrNo = 28

	ESPIPE ErrNo = 29

	EROFS ErrNo = 30

	EMLINK ErrNo = 31

	EPIPE ErrNo = 32

	EDOM ErrNo = 33

	ERANGE ErrNo = 34
)

const (
	MAX_BUFFER_SIZE = 1024
)

type TargetOpenFlag uint32

const (
	TargetOpenFlag_O_RDONLY TargetOpenFlag = 0

	TargetOpenFlag_O_WRONLY TargetOpenFlag = 1

	TargetOpenFlag_O_RDWR TargetOpenFlag = 2

	TargetOpenFlag_O_CREAT TargetOpenFlag = 0x100

	TargetOpenFlag_O_EXCL TargetOpenFlag = 0x400

	TargetOpenFlag_O_NOCTTY TargetOpenFlag = 0x800

	TargetOpenFlag_O_TRUNC TargetOpenFlag = 0x200

	TargetOpenFlag_O_APPEND TargetOpenFlag = 8

	TargetOpenFlag_O_NONBLOCK TargetOpenFlag = 0x80

	TargetOpenFlag_O_SYNC TargetOpenFlag = 0x10
)

type OpenFlag uint32

const (
	OpenFlag_O_RDONLY OpenFlag = 0x00000000

	OpenFlag_O_WRONLY OpenFlag = 0x00000001

	OpenFlag_O_RDWR OpenFlag = 0x00000002

	OpenFlag_O_CREAT OpenFlag = 0x00000040

	OpenFlag_O_EXCL OpenFlag = 0x00000080

	OpenFlag_O_NOCTTY OpenFlag = 0x00000100

	OpenFlag_O_TRUNC OpenFlag = 0x00000200

	OpenFlag_O_APPEND OpenFlag = 0x00000400

	OpenFlag_O_NONBLOCK OpenFlag = 0x00000800

	OpenFlag_O_SYNC OpenFlag = 0x00001000
)

type OpenFlagMapping struct {
	TargetFlag TargetOpenFlag
	HostFlag   OpenFlag
}

func NewOpenFlagMapping(targetFlag TargetOpenFlag, hostFlag OpenFlag) *OpenFlagMapping {
	var openFlagMapping = &OpenFlagMapping{
		TargetFlag:targetFlag,
		HostFlag:hostFlag,
	}

	return openFlagMapping
}

type SysCtrlArgs struct {
	Name    uint32
	Nlen    uint32
	Oldval  uint32
	Oldlenp uint32
	Newval  uint32
	Newlen  uint32
}

type SyscallHandler struct {
	Index uint32
	Name  string
	Run   func(context *cpu.Context)
}

func NewSyscallHandler(index uint32, name string, run func(context *cpu.Context)) *SyscallHandler {
	var syscallHandler = &SyscallHandler{
		Index:index,
		Name:name,
		Run:run,
	}

	return syscallHandler
}

type SyscallEmulation struct {
	Handlers         map[uint32]*SyscallHandler
	StackLimit       uint32
	Error            bool
	OpenFlagMappings []*OpenFlagMapping
}

func NewSyscallEmulation() *SyscallEmulation {
	var syscallEmulation = &SyscallEmulation{
		StackLimit: 0x800000,
	}

	syscallEmulation.OpenFlagMappings = append(
		syscallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_RDONLY, OpenFlag_O_RDONLY))

	syscallEmulation.OpenFlagMappings = append(
		syscallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_WRONLY, OpenFlag_O_WRONLY))

	syscallEmulation.OpenFlagMappings = append(
		syscallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_RDWR, OpenFlag_O_RDWR))

	syscallEmulation.OpenFlagMappings = append(
		syscallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_APPEND, OpenFlag_O_APPEND))

	syscallEmulation.OpenFlagMappings = append(
		syscallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_SYNC, OpenFlag_O_SYNC))

	syscallEmulation.OpenFlagMappings = append(
		syscallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_CREAT, OpenFlag_O_CREAT))

	syscallEmulation.OpenFlagMappings = append(
		syscallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_TRUNC, OpenFlag_O_TRUNC))

	syscallEmulation.OpenFlagMappings = append(
		syscallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_EXCL, OpenFlag_O_EXCL))

	syscallEmulation.OpenFlagMappings = append(
		syscallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_NOCTTY, OpenFlag_O_NOCTTY))

	syscallEmulation.OpenFlagMappings = append(
		syscallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag(0x2000), OpenFlag(0)))

	syscallEmulation.registerHandler(NewSyscallHandler(1, "exit", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(3, "read", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(4, "write", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(5, "open", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(6, "close", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(7, "waitpid", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(20, "getpid", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(24, "getuid", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(37, "kill", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(42, "pipe", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(45, "brk", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(47, "getgid", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(49, "geteuid", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(50, "getegid", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(54, "ioctl", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(64, "getppid", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(75, "setrlimit", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(76, "getrlimit", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(90, "mmap", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(91, "munmap", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(120, "clone", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(122, "uname", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(125, "mprotect", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(140, "_llseek", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(153, "_sysctl", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(166, "nanosleep", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(167, "mremap", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(188, "poll", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(194, "rt_sigaction", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(195, "rt_sigprocmask", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(199, "rt_sigsuspend", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(215, "fstat64", func(context *cpu.Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(246, "exit_group", func(context *cpu.Context) {}))

	return syscallEmulation
}

func (syscallEmulation *SyscallEmulation) registerHandler(handler *SyscallHandler) {
	syscallEmulation.Handlers[handler.Index] = handler
}

func (syscallEmulation *SyscallEmulation) findAndRunSystemCallHandler(syscallIndex uint32, context *cpu.Context) bool {
	if handler, ok := syscallEmulation.Handlers[syscallIndex]; ok {
		handler.Run(context)

		context.Kernel.ProcessSystemEvents()
		context.Kernel.ProcessSignals()

		return true
	}

	return false
}

func (syscallEmulation *SyscallEmulation) checkSystemCallError(context *cpu.Context) bool {
	if int32(context.Regs.Gpr[regs.REGISTER_V0]) != -1 {
		context.Regs.Gpr[regs.REGISTER_A3] = 0
		return false
	} else {
		context.Regs.Gpr[regs.REGISTER_V0] = 0
		context.Regs.Gpr[regs.REGISTER_A3] = 1
		return true
	}
}

func (syscallEmulation *SyscallEmulation) DoSystemCall(callNum uint32, context *cpu.Context) {
	var syscallIndex = callNum - 4000

	if !syscallEmulation.findAndRunSystemCallHandler(syscallIndex, context) {
		panic(fmt.Sprintf("ctx-%d: system call %d (%d) not implemented", context.Id, callNum, syscallIndex))
	}
}

func exit_impl(context *cpu.Context) {
	context.Finish()
}

func read_impl(context *cpu.Context) {
	//var readMaxSize = 1 << 25
	//
	//var fd = context.Process.TranslateFileDescriptor(context.Regs.Gpr[REGISTER_A0])
	//var bufAddr = context.Regs.Gpr[REGISTER_A1]
	//var count = math.Min(float64(readMaxSize), float64(context.Regs.Gpr[REGISTER_A2]))
	//
	//var ret uint32
	//var buf []byte
	//
	//var buffer = context.Kernel
}