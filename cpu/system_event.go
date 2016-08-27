package cpu

import (
	"github.com/mcai/acogo/cpu/mem"
	"github.com/mcai/acogo/cpu/native"
	"reflect"
	"github.com/mcai/acogo/cpu/regs"
)

type SystemEventCriterion interface {
	NeedProcess(context *Context) bool
}

type TimeCriterion struct {
	When int32
}

func NewTimeCriterion() *TimeCriterion {
	var timeCriterion = &TimeCriterion{
	}

	return timeCriterion
}

func (timeCriterion *TimeCriterion) NeedProcess(context *Context) bool {
	return timeCriterion.When <= native.Clock(context.Kernel.CurrentCycle)
}

type SignalCriterion struct {
}

func NewSignalCriterion() *SignalCriterion {
	var signalCriterion = &SignalCriterion{
	}

	return signalCriterion
}

func (signalCriterion *SignalCriterion) NeedProcess(context *Context) bool {
	for signal := uint32(1); signal <= MAX_SIGNAL; signal++ {
		if context.Kernel.MustProcessSignal(context, signal) {
			return true
		}
	}

	return false
}

type WaitForProcessIdCriterion struct {
	ProcessId        int32
	HasProcessKilled bool
}

func NewWaitForProcessIdCriterion(context *Context, processId int32) *WaitForProcessIdCriterion {
	var waitForProcessIdCriterion = &WaitForProcessIdCriterion{
		ProcessId:processId,
	}

	context.Kernel.Experiment.BlockingEventDispatcher.AddListener(reflect.TypeOf((*ContextKilledEvent)(nil)), func(event interface{}) {
		waitForProcessIdCriterion.HasProcessKilled = true
	})

	return waitForProcessIdCriterion
}

func (waitForProcessIdCriterion *WaitForProcessIdCriterion) NeedProcess(context *Context) bool {
	return waitForProcessIdCriterion.ProcessId == -1 && waitForProcessIdCriterion.HasProcessKilled ||
		waitForProcessIdCriterion.ProcessId > 0 && context.Kernel.GetContextFromProcessId(waitForProcessIdCriterion.ProcessId) == nil
}

type WaitForFileDescriptorCriterion struct {
	Buffer  *mem.CircularByteBuffer
	Address uint32
	Size    uint32
	Pufds   uint32
}

func NewWaitForFileDescriptorCriterion() *WaitForFileDescriptorCriterion {
	var waitForFileDescriptorCriterion = &WaitForFileDescriptorCriterion{
	}

	return waitForFileDescriptorCriterion
}

func (waitForFileDescriptorCriterion *WaitForFileDescriptorCriterion) NeedProcess(context *Context) bool {
	return !waitForFileDescriptorCriterion.Buffer.IsEmpty()
}

const (
	SystemEventType_READ = 0
	SystemEventType_RESUME = 1
	SystemEventType_WAIT = 2
	SystemEventType_POLL = 3
	SystemEventType_SIGNAL_SUSPEND = 4
)

type SystemEventType uint32

type SystemEvent interface {
	Context() *Context
	EventType() SystemEventType
	NeedProcess() bool
	Process()
}

type BaseSystemEvent struct {
	context   *Context
	eventType SystemEventType
}

func NewBaseSystemEvent(context *Context, eventType SystemEventType) *BaseSystemEvent {
	var baseSystemEvent = &BaseSystemEvent{
		context:context,
		eventType:eventType,
	}

	return baseSystemEvent
}

func (baseSystemEvent *BaseSystemEvent) Context() *Context {
	return baseSystemEvent.context
}

func (baseSystemEvent *BaseSystemEvent) EventType() SystemEventType {
	return baseSystemEvent.eventType
}

type PollEvent struct {
	*BaseSystemEvent
	TimeCriterion                  *TimeCriterion
	WaitForFileDescriptorCriterion *WaitForFileDescriptorCriterion
}

func NewPollEvent(context *Context) *PollEvent {
	var pollEvent = &PollEvent{
		BaseSystemEvent:NewBaseSystemEvent(context, SystemEventType_POLL),
		TimeCriterion:NewTimeCriterion(),
		WaitForFileDescriptorCriterion: NewWaitForFileDescriptorCriterion(),
	}

	return pollEvent
}

func (pollEvent *PollEvent) NeedProcess() bool {
	return pollEvent.TimeCriterion.NeedProcess(pollEvent.context) ||
		pollEvent.WaitForFileDescriptorCriterion.NeedProcess(pollEvent.context)
}

func (pollEvent *PollEvent) Process() {
	if !pollEvent.WaitForFileDescriptorCriterion.Buffer.IsEmpty() {
		pollEvent.Context().Process.Memory.WriteHalfWordAt(pollEvent.WaitForFileDescriptorCriterion.Pufds + 6, 1)
		pollEvent.Context().Regs.Gpr[regs.REGISTER_V0] = 1
	} else {
		pollEvent.Context().Regs.Gpr[regs.REGISTER_V0] = 0
	}

	pollEvent.Context().Regs.Gpr[regs.REGISTER_A3] = 0
	pollEvent.context.Resume()
}

type ReadEvent struct {
	*BaseSystemEvent
	WaitForFileDescriptorCriterion *WaitForFileDescriptorCriterion
}

func NewReadEvent(context *Context) *ReadEvent {
	var readEvent = &ReadEvent{
		BaseSystemEvent: NewBaseSystemEvent(context, SystemEventType_READ),
		WaitForFileDescriptorCriterion:NewWaitForFileDescriptorCriterion(),
	}

	return readEvent
}

func (readEvent *ReadEvent) NeedProcess() bool {
	return readEvent.WaitForFileDescriptorCriterion.NeedProcess(readEvent.context)
}

func (readEvent *ReadEvent) Process() {
	readEvent.Context().Resume()

	var buf = make([]byte, readEvent.WaitForFileDescriptorCriterion.Size)

	var numRead = readEvent.WaitForFileDescriptorCriterion.Buffer.Read(&buf, uint32(len(buf)))

	readEvent.Context().Regs.Gpr[regs.REGISTER_V0] = uint32(numRead)
	readEvent.Context().Regs.Gpr[regs.REGISTER_A3] = 0

	readEvent.Context().Process.Memory.WriteBlockAt(readEvent.WaitForFileDescriptorCriterion.Address, numRead, buf)
}

type ResumeEvent struct {
	*BaseSystemEvent
	TimeCriterion *TimeCriterion
}

func NewResumeEvent(context *Context) *ResumeEvent {
	var resumeEvent = &ResumeEvent{
		BaseSystemEvent: NewBaseSystemEvent(context, SystemEventType_RESUME),
		TimeCriterion:NewTimeCriterion(),
	}

	return resumeEvent
}

func (resumeEvent *ResumeEvent) NeedProcess() bool {
	return resumeEvent.TimeCriterion.NeedProcess(resumeEvent.context)
}

func (resumeEvent *ResumeEvent) Process() {
	resumeEvent.Context().Resume()
}

type SignalSuspendEvent struct {
	*BaseSystemEvent
	SignalCriterion *SignalCriterion
}

func NewSignalSuspendEvent(context *Context) *SignalSuspendEvent {
	var signalSuspendEvent = &SignalSuspendEvent{
		BaseSystemEvent: NewBaseSystemEvent(context, SystemEventType_SIGNAL_SUSPEND),
		SignalCriterion:NewSignalCriterion(),
	}

	return signalSuspendEvent
}

func (signalSuspendEvent *SignalSuspendEvent) NeedProcess() bool {
	return signalSuspendEvent.SignalCriterion.NeedProcess(signalSuspendEvent.context)
}

func (signalSuspendEvent *SignalSuspendEvent) Process() {
	signalSuspendEvent.Context().Resume()

	signalSuspendEvent.Context().Kernel.ProcessSignals()

	signalSuspendEvent.Context().SignalMasks.Blocked = signalSuspendEvent.Context().SignalMasks.Backup.Clone()
}

type WaitEvent struct {
	*BaseSystemEvent
	WaitForProcessIdCriterion *WaitForProcessIdCriterion
	SignalCriterion           *SignalCriterion
}

func NewWaitEvent(context *Context, processId int32) *WaitEvent {
	var waitEvent = &WaitEvent{
		BaseSystemEvent: NewBaseSystemEvent(context, SystemEventType_WAIT),
		WaitForProcessIdCriterion: NewWaitForProcessIdCriterion(context, processId),
		SignalCriterion: NewSignalCriterion(),
	}

	return waitEvent
}

func (waitEvent *WaitEvent) NeedProcess() bool {
	return waitEvent.WaitForProcessIdCriterion.NeedProcess(waitEvent.context) ||
		waitEvent.SignalCriterion.NeedProcess(waitEvent.context)
}

func (waitEvent *WaitEvent) Process() {
	waitEvent.Context().Resume()

	waitEvent.Context().Regs.Gpr[regs.REGISTER_V0] = uint32(waitEvent.WaitForProcessIdCriterion.ProcessId)
	waitEvent.Context().Regs.Gpr[regs.REGISTER_A3] = 0
}


