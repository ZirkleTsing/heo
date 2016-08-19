package acogo

import "fmt"

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

type SystemCallHandler struct {
	Index uint32
	Name  string
	Run   func(context *Context)
}

func NewSystemCallHandler(index uint32, name string, run func(context *Context)) *SystemCallHandler {
	var systemCallHandler = &SystemCallHandler{
		Index:index,
		Name:name,
		Run:run,
	}

	return systemCallHandler
}

type SystemCallEmulation struct {
	Handlers         map[uint32]*SystemCallHandler
	StackLimit       uint32
	Error            bool
	OpenFlagMappings []*OpenFlagMapping
}

func NewSystemCallEmulation() *SystemCallEmulation {
	var systemCallEmulation = &SystemCallEmulation{
		StackLimit: 0x800000,
	}

	systemCallEmulation.OpenFlagMappings = append(
		systemCallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_RDONLY, OpenFlag_O_RDONLY))

	systemCallEmulation.OpenFlagMappings = append(
		systemCallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_WRONLY, OpenFlag_O_WRONLY))

	systemCallEmulation.OpenFlagMappings = append(
		systemCallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_RDWR, OpenFlag_O_RDWR))

	systemCallEmulation.OpenFlagMappings = append(
		systemCallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_APPEND, OpenFlag_O_APPEND))

	systemCallEmulation.OpenFlagMappings = append(
		systemCallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_SYNC, OpenFlag_O_SYNC))

	systemCallEmulation.OpenFlagMappings = append(
		systemCallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_CREAT, OpenFlag_O_CREAT))

	systemCallEmulation.OpenFlagMappings = append(
		systemCallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_TRUNC, OpenFlag_O_TRUNC))

	systemCallEmulation.OpenFlagMappings = append(
		systemCallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_EXCL, OpenFlag_O_EXCL))

	systemCallEmulation.OpenFlagMappings = append(
		systemCallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag_O_NOCTTY, OpenFlag_O_NOCTTY))

	systemCallEmulation.OpenFlagMappings = append(
		systemCallEmulation.OpenFlagMappings,
		NewOpenFlagMapping(TargetOpenFlag(0x2000), OpenFlag(0)))

	systemCallEmulation.registerHandler(NewSystemCallHandler(1, "exit", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(3, "read", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(4, "write", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(5, "open", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(6, "close", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(7, "waitpid", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(20, "getpid", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(24, "getuid", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(37, "kill", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(42, "pipe", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(45, "brk", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(47, "getgid", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(49, "geteuid", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(50, "getegid", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(54, "ioctl", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(64, "getppid", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(75, "setrlimit", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(76, "getrlimit", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(90, "mmap", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(91, "munmap", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(120, "clone", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(122, "uname", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(125, "mprotect", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(140, "_llseek", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(153, "_sysctl", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(166, "nanosleep", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(167, "mremap", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(188, "poll", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(194, "rt_sigaction", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(195, "rt_sigprocmask", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(199, "rt_sigsuspend", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(215, "fstat64", func(context *Context) {}))
	systemCallEmulation.registerHandler(NewSystemCallHandler(246, "exit_group", func(context *Context) {}))

	return systemCallEmulation
}

func (systemCallEmulation *SystemCallEmulation) registerHandler(handler *SystemCallHandler) {
	systemCallEmulation.Handlers[handler.Index] = handler
}

func (systemCallEmulation *SystemCallEmulation) findAndRunSystemCallHandler(systemCallIndex uint32, context *Context) bool {
	if handler, ok := systemCallEmulation.Handlers[systemCallIndex]; ok {
		handler.Run(context)

		context.Kernel.ProcessSystemEvents()
		context.Kernel.ProcessSignals()

		return true
	}

	return false
}

func (systemCallEmulation *SystemCallEmulation) checkSystemCallError(context *Context) bool {
	if int32(context.Regs.Gpr[REGISTER_V0]) != -1 {
		context.Regs.Gpr[REGISTER_A3] = 0
		return false
	} else {
		context.Regs.Gpr[REGISTER_V0] = 0
		context.Regs.Gpr[REGISTER_A3] = 1
		return true
	}
}

func (systemCallEmulation *SystemCallEmulation) DoSystemCall(callNum uint32, context *Context) {
	var systemCallIndex = callNum - 4000

	if !systemCallEmulation.findAndRunSystemCallHandler(systemCallIndex, context) {
		panic(fmt.Sprintf("ctx-%d: system call %d (%d) not implemented", context.Id, callNum, systemCallIndex))
	}
}
