package cpu

import (
	"fmt"
	"math"
	"github.com/mcai/acogo/cpu/regs"
	"github.com/mcai/acogo/cpu/native"
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
	Run   func(context *Context)
}

func NewSyscallHandler(index uint32, name string, run func(context *Context)) *SyscallHandler {
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
		Handlers:make(map[uint32]*SyscallHandler),
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

	syscallEmulation.registerHandler(NewSyscallHandler(1, "exit", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(3, "read", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(4, "write", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(5, "open", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(6, "close", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(7, "waitpid", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(20, "getpid", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(24, "getuid", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(37, "kill", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(42, "pipe", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(45, "brk", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(47, "getgid", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(49, "geteuid", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(50, "getegid", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(54, "ioctl", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(64, "getppid", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(75, "setrlimit", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(76, "getrlimit", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(90, "mmap", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(91, "munmap", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(120, "clone", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(122, "uname", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(125, "mprotect", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(140, "_llseek", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(153, "_sysctl", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(166, "nanosleep", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(167, "mremap", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(188, "poll", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(194, "rt_sigaction", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(195, "rt_sigprocmask", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(199, "rt_sigsuspend", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(215, "fstat64", func(context *Context) {}))
	syscallEmulation.registerHandler(NewSyscallHandler(246, "exit_group", func(context *Context) {}))

	return syscallEmulation
}

func (syscallEmulation *SyscallEmulation) registerHandler(handler *SyscallHandler) {
	syscallEmulation.Handlers[handler.Index] = handler
}

func (syscallEmulation *SyscallEmulation) findAndRunSystemCallHandler(syscallIndex uint32, context *Context) bool {
	if handler, ok := syscallEmulation.Handlers[syscallIndex]; ok {
		handler.Run(context)

		context.Kernel.ProcessSystemEvents()
		context.Kernel.ProcessSignals()

		return true
	}

	return false
}

func (syscallEmulation *SyscallEmulation) checkSystemCallError(context *Context) bool {
	if int32(context.Regs.Gpr[regs.REGISTER_V0]) != -1 {
		context.Regs.Gpr[regs.REGISTER_A3] = 0
		return false
	} else {
		context.Regs.Gpr[regs.REGISTER_V0] = 0
		context.Regs.Gpr[regs.REGISTER_A3] = 1
		return true
	}
}

func (syscallEmulation *SyscallEmulation) DoSystemCall(callNum uint32, context *Context) {
	var syscallIndex = callNum - 4000

	if !syscallEmulation.findAndRunSystemCallHandler(syscallIndex, context) {
		panic(fmt.Sprintf("ctx-%d: system call %d (%d) not implemented", context.Id, callNum, syscallIndex))
	}
}

func (syscallEmulation *SyscallEmulation) exit_impl(context *Context) {
	context.Finish()
}

func (syscallEmulation *SyscallEmulation) read_impl(context *Context) {
	var readMaxSize = uint64(1 << 25)

	var fd = int(context.Process.TranslateFileDescriptor(context.Regs.Gpr[regs.REGISTER_A0]))
	var bufAddr = uint64(context.Regs.Gpr[regs.REGISTER_A1])
	var count = uint64(math.Min(float64(readMaxSize), float64(context.Regs.Gpr[regs.REGISTER_A2])))

	var ret uint64
	var buf []byte

	var buffer = context.Kernel.GetReadBuffer(fd)
	if buffer != nil {
		if buffer.IsEmpty() {
			var e = NewReadEvent(context)
			e.WaitForFileDescriptorCriterion.Buffer = buffer
			e.WaitForFileDescriptorCriterion.Address = bufAddr
			e.WaitForFileDescriptorCriterion.Size = count
			context.Kernel.SystemEvents = append(context.Kernel.SystemEvents, e)
			context.Suspend()
			return
		} else {
			buf = make([]byte, count)
			ret = buffer.Read(&buf, count)
		}
	} else {
		buf = make([]byte, count)
		ret = uint64(native.Read(fd, buf))
	}

	if uint64(ret) >= readMaxSize {
		panic("Impossible")
	}

	context.Regs.Gpr[regs.REGISTER_V0] = uint32(ret)
	syscallEmulation.Error = syscallEmulation.checkSystemCallError(context)

	context.Process.Memory.WriteBlockAt(bufAddr, ret, buf)
}

func (syscallEmulation *SyscallEmulation) write_impl(context *Context) {
	var fd = int(context.Process.TranslateFileDescriptor(context.Regs.Gpr[regs.REGISTER_A0]))
	var bufAddr = uint64(context.Regs.Gpr[regs.REGISTER_A1])
	var count = uint64(context.Regs.Gpr[regs.REGISTER_A2])

	var buf = context.Process.Memory.ReadBlockAt(bufAddr, count)

	var ret uint64

	var buffer = context.Kernel.GetWriteBuffer(fd)
	if buffer != nil {
		buffer.Write(&buf, count)
		ret = count
	} else {
		ret = uint64(native.Write(fd, buf))
	}

	context.Regs.Gpr[regs.REGISTER_V0] = uint32(ret)
	syscallEmulation.Error = syscallEmulation.checkSystemCallError(context)
}

func (syscallEmulation *SyscallEmulation) open_impl(context *Context) {
	var addr = uint64(context.Regs.Gpr[regs.REGISTER_A0])
	var targetFlags = context.Regs.Gpr[regs.REGISTER_A1]
	var mode = int(context.Regs.Gpr[regs.REGISTER_A2])

	var hostFlags = uint32(0)
	for _, mapping := range syscallEmulation.OpenFlagMappings {
		if targetFlags & uint32(mapping.TargetFlag) != 0 {
			targetFlags &= ^uint32(mapping.TargetFlag)
			hostFlags |= uint32(mapping.HostFlag)
		}
	}

	if targetFlags != 0 {
		panic(fmt.Sprintf("syscall open: cannot decode flags 0x%08x", targetFlags))
	}

	var path = context.Process.Memory.ReadStringAt(addr, MAX_BUFFER_SIZE)

	var ret = native.Open(path, mode, hostFlags)

	context.Regs.Gpr[regs.REGISTER_V0] = uint32(ret)
	syscallEmulation.Error = syscallEmulation.checkSystemCallError(context)
}

func (syscallEmulation *SyscallEmulation) close_impl(context *Context) {
	var fd = int(context.Regs.Gpr[regs.REGISTER_A0])

	if fd == 0 || fd == 1 || fd == 2 {
		context.Regs.Gpr[regs.REGISTER_A3] = 0
		return
	}

	var ret = native.Close(fd)

	context.Regs.Gpr[regs.REGISTER_V0] = uint32(ret)
	syscallEmulation.Error = syscallEmulation.checkSystemCallError(context)
}

func (syscallEmulation *SyscallEmulation) waitpid_impl(context *Context) {
	var pid = context.Regs.Gpr[regs.REGISTER_A0]
	var pstatus = context.Regs.Gpr[regs.REGISTER_A1]

	if pid < 1 {
		panic("Impossible")
	}

	if context.Kernel.GetContextFromProcessId(pid) == nil {
		context.Regs.Gpr[regs.REGISTER_A3] = uint32(ECHILD)
		context.Regs.SetSgpr(regs.REGISTER_V0, -1)
		return
	}

	var e = NewWaitEvent(context, pid)
	context.Kernel.SystemEvents = append(context.Kernel.SystemEvents, e)
	context.Suspend()
	if pstatus != 0 {
		panic("Impossible")
	}
}

func (syscallEmulation *SyscallEmulation) getpid_impl(context *Context) {
	context.Regs.Gpr[regs.REGISTER_V0] = context.ProcessId
	syscallEmulation.Error = syscallEmulation.checkSystemCallError(context)
}

func (syscallEmulation *SyscallEmulation) getuid_impl(context *Context) {
	context.Regs.Gpr[regs.REGISTER_V0] = context.UserId
	syscallEmulation.Error = syscallEmulation.checkSystemCallError(context)
}

func (SyscallEmulation *SyscallEmulation) kill_impl(context *Context) {
	var pid = int(context.Regs.Gpr[regs.REGISTER_A0])
	var sig = context.Regs.Gpr[regs.REGISTER_A1]
	if pid < 0 {
		panic("Impossible")
	}

	var destContext = context.Kernel.GetContextFromProcessId(uint32(pid))
	if destContext == nil {
		context.Regs.Gpr[regs.REGISTER_A3] = uint32(ESRCH)
		context.Regs.SetSgpr(regs.REGISTER_V0, -1)
		return
	}

	destContext.SignalMasks.Pending.Set(sig)
	context.Regs.Gpr[regs.REGISTER_A3] = 0
	context.Regs.Gpr[regs.REGISTER_V0] = 0
}

func (syscallEmulation *SyscallEmulation) pipe_impl(context *Context) {
	var fileDescriptors =  context.Kernel.CreatePipe()

	context.Regs.Gpr[regs.REGISTER_V0] = uint32(fileDescriptors[0])
	context.Regs.Gpr[regs.REGISTER_V1] = uint32(fileDescriptors[1])

	context.Regs.Gpr[regs.REGISTER_A3] = 0
}

func (syscallEmulation *SyscallEmulation) brk_impl(context *Context) {
	context.Regs.Gpr[regs.REGISTER_A3] = 0
	context.Regs.Gpr[regs.REGISTER_V0] = 0
}

func (SyscallEmulation *SyscallEmulation) getgid_impl(context *Context) {
	context.Regs.Gpr[regs.REGISTER_V0] = context.GroupId
	SyscallEmulation.Error = SyscallEmulation.checkSystemCallError(context)
}

func (syscallEmulation *SyscallEmulation) geteuid_impl(context *Context) {
	context.Regs.Gpr[regs.REGISTER_V0] = context.EffectiveUserId
	syscallEmulation.Error = syscallEmulation.checkSystemCallError(context)
}

func (syscallEmulation *SyscallEmulation) getegid_impl(context *Context) {
	context.Regs.Gpr[regs.REGISTER_V0] = context.EffectiveGroupId
	syscallEmulation.Error = syscallEmulation.checkSystemCallError(context)
}

func (syscallEmulation *SyscallEmulation) ioctl_impl(context *Context) {
	var buf = make([]byte, 128)

	if context.Regs.Gpr[regs.REGISTER_A2] != 0 {
		buf = context.Process.Memory.ReadBlockAt(uint64(context.Regs.Gpr[regs.REGISTER_A2]), 128)
	}

	var fd = int(context.Process.TranslateFileDescriptor(context.Regs.Gpr[regs.REGISTER_A0]))
	if fd < 3 {
		context.Regs.Gpr[regs.REGISTER_V0] = uint32(native.Ioctl(fd, int(context.Regs.Gpr[regs.REGISTER_A1]), buf))

		syscallEmulation.Error = syscallEmulation.checkSystemCallError(context)

		if context.Regs.Gpr[regs.REGISTER_A2] != 0 {
			context.Process.Memory.WriteBlockAt(uint64(context.Regs.Gpr[regs.REGISTER_A2]), 128, buf)
		}
	} else {
		context.Regs.Gpr[regs.REGISTER_A3] = 0
		context.Regs.Gpr[regs.REGISTER_V0] = 0
	}
}

func (syscallEmulation *SyscallEmulation) getppid_impl(context *Context) {
	context.Regs.Gpr[regs.REGISTER_V0] = context.GetParentProcessId()
	syscallEmulation.Error = syscallEmulation.checkSystemCallError(context)
}

func (syscallEmulation *SyscallEmulation) setrlimit_impl(context *Context) {
	if context.Regs.Gpr[regs.REGISTER_A0] != 3 {
		panic("Impossbile")
	}

	syscallEmulation.StackLimit = context.Regs.Gpr[regs.REGISTER_A1]

	context.Regs.Gpr[regs.REGISTER_A3] = 0
	context.Regs.Gpr[regs.REGISTER_V0] = 0
}

func (syscallEmulation *SyscallEmulation) getrlimit_impl(context *Context) {
	var prlimit = uint64(context.Regs.Gpr[regs.REGISTER_A1])

	if context.Regs.Gpr[regs.REGISTER_A0] != 3 {
		panic("Impossible")
	}

	context.Process.Memory.WriteWordAt(prlimit, syscallEmulation.StackLimit)
	context.Process.Memory.WriteWordAt(prlimit + 4, 0xffffffff)

	context.Regs.Gpr[regs.REGISTER_A3] = 0
	context.Regs.Gpr[regs.REGISTER_V0] = 0
}

func (syscallEmulation *SyscallEmulation) mmap_impl(context *Context) {
	var start = uint64(context.Regs.Gpr[regs.REGISTER_A0])
	var length = uint64(context.Regs.Gpr[regs.REGISTER_A1])

	var fd = int(context.Process.Memory.ReadWordAt(uint64(context.Regs.Gpr[regs.REGISTER_SP] + 16)))

	if fd != -1 {
		panic("syscall mmap: syscall is only supported with fd = -1")
	}

	if start == 0 {
		start = uint64(context.Process.HeapTop)
	}

	start = context.Process.Memory.Map(start, length)

	context.Regs.Gpr[regs.REGISTER_A3] = 0
	context.Regs.Gpr[regs.REGISTER_V0] = uint32(start)
}

func (syscallEmulation *SyscallEmulation) munmap_impl(context *Context) {
	var start = uint64(context.Regs.Gpr[regs.REGISTER_A0])
	var length = uint64(context.Regs.Gpr[regs.REGISTER_A1])
	
	context.Process.Memory.Unmap(start, length)

	context.Regs.Gpr[regs.REGISTER_A3] = 0
	context.Regs.Gpr[regs.REGISTER_V0] = 0
}