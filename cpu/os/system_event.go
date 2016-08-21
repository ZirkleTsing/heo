package os

import "github.com/mcai/acogo/cpu"

type SystemEventCriterion interface {
	NeedProcess(context *cpu.Context) bool
}

type TimeCriterion struct {
	When uint64
}

func NewTimeCriterion() *TimeCriterion {
	var timeCriterion = &TimeCriterion{
	}

	return timeCriterion
}

func (timeCriterion *TimeCriterion) NeedProcess(context *cpu.Context) bool {
	return timeCriterion.When <= Clock(context.Kernel.CurrentCycle)
}

type SignalCriterion struct {
}

func NewSignalCriterion() *SignalCriterion {
	var signalCriterion = &SignalCriterion{
	}

	return signalCriterion
}

func (signalCriterion *SignalCriterion) NeedProcess(context *cpu.Context) bool {
	for signal := uint32(1); signal <= MAX_SIGNAL; signal++ {
		if context.Kernel.MustProcessSignal(context, signal) {
			return true
		}
	}

	return false
}

type WaitForProcessIdCriterion struct {
	ProcessId        uint32
	HasProcessKilled bool
}

func NewWaitForProcessIdCriterion(context *cpu.Context, processId uint32) *WaitForProcessIdCriterion {
	var waitForProcessIdCriterion = &WaitForProcessIdCriterion{
		ProcessId:processId,
	}

	//TODO

	return waitForProcessIdCriterion
}

func (waitForProcessIdCriterion *WaitForProcessIdCriterion) NeedProcess(context *cpu.Context) bool {
	return false //TODO
}

type WaitForFileDescriptorCriterion struct {
	//TODO: buffer
	Address uint32
	Size    uint32
	Pufds   uint32
}

func NewWaitForFileDescriptorCriterion() *WaitForFileDescriptorCriterion {
	var waitForFileDescriptorCriterion = &WaitForFileDescriptorCriterion{
	}

	return waitForFileDescriptorCriterion
}

func (waitForFileDescriptorCriterion *WaitForFileDescriptorCriterion) NeedProcess(context *cpu.Context) bool {
	return false //TODO
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
	GetContext() *cpu.Context
	GetEventType() SystemEventType
	NeedProcess() bool
	Process()
}

type BaseSystemEvent struct {
	Context   *cpu.Context
	EventType SystemEventType
}

func NewBaseSystemEvent(context *cpu.Context, eventType SystemEventType) *BaseSystemEvent {
	var baseSystemEvent = &BaseSystemEvent{
		Context:context,
		EventType:eventType,
	}

	return baseSystemEvent
}

type PollEvent struct {
	*BaseSystemEvent
	TimeCriterion                  *TimeCriterion
	WaitForFileDescriptorCriterion *WaitForFileDescriptorCriterion
}

func NewPollEvent(context *cpu.Context) *PollEvent {
	var pollEvent = &PollEvent{
		BaseSystemEvent:NewBaseSystemEvent(context, SystemEventType_POLL),
		TimeCriterion:NewTimeCriterion(),
		WaitForFileDescriptorCriterion: NewWaitForFileDescriptorCriterion(),
	}

	return pollEvent
}

func (pollEvent *PollEvent) NeedProcess() bool {
	return pollEvent.TimeCriterion.NeedProcess(pollEvent.Context) ||
		pollEvent.WaitForFileDescriptorCriterion.NeedProcess(pollEvent.Context)
}

func (pollEvent *PollEvent) Process() {
	//TODO
}

type ReadEvent struct {
	*BaseSystemEvent
	WaitForFileDescriptorCriterion *WaitForFileDescriptorCriterion
}

func NewReadEvent(context *cpu.Context) *ReadEvent {
	var readEvent = &ReadEvent{
		BaseSystemEvent: NewBaseSystemEvent(context, SystemEventType_READ),
		WaitForFileDescriptorCriterion:NewWaitForFileDescriptorCriterion(),
	}

	return readEvent
}

func (readEvent *ReadEvent) NeedProcess() bool {
	return readEvent.WaitForFileDescriptorCriterion.NeedProcess(readEvent.Context)
}

func (readEvent *ReadEvent) Process() {
	//TODO
}

type ResumeEvent struct {
	*BaseSystemEvent
	TimeCriterion *TimeCriterion
}

func NewResumeEvent(context *cpu.Context) *ResumeEvent {
	var resumeEvent = &ResumeEvent{
		BaseSystemEvent: NewBaseSystemEvent(context, SystemEventType_RESUME),
		TimeCriterion:NewTimeCriterion(),
	}

	return resumeEvent
}

func (resumeEvent *ResumeEvent) NeedProcess(context *cpu.Context) bool {
	return resumeEvent.TimeCriterion.NeedProcess(context)
}

func (resumeEvent *ResumeEvent) Process() {
	//TODO
}

type SignalSuspendEvent struct {
	*BaseSystemEvent
	SignalCriterion *SignalCriterion
}

func NewSignalSuspendEvent(context *cpu.Context) *SignalSuspendEvent {
	var signalSuspendEvent = &SignalSuspendEvent{
		BaseSystemEvent: NewBaseSystemEvent(context, SystemEventType_SIGNAL_SUSPEND),
		SignalCriterion:NewSignalCriterion(),
	}

	return signalSuspendEvent
}

func (signalSuspendEvent *SignalSuspendEvent) NeedProcess(context *cpu.Context) bool {
	return signalSuspendEvent.SignalCriterion.NeedProcess(context)
}

func (signalSuspendEvent *SignalSuspendEvent) Process() {
	//TODO
}

type WaitEvent struct {
	*BaseSystemEvent
	WaitForProcessIdCriterion *WaitForProcessIdCriterion
	SignalCriterion           *SignalCriterion
}

func NewWaitEvent(context *cpu.Context, processId uint32) *WaitEvent {
	var waitEvent = &WaitEvent{
		BaseSystemEvent: NewBaseSystemEvent(context, SystemEventType_WAIT),
		WaitForProcessIdCriterion: NewWaitForProcessIdCriterion(context, processId),
		SignalCriterion: NewSignalCriterion(),
	}

	return waitEvent
}

func (waitEvent *WaitEvent) NeedProcess() bool {
	return waitEvent.WaitForProcessIdCriterion.NeedProcess(waitEvent.Context) ||
		waitEvent.SignalCriterion.NeedProcess(waitEvent.Context)
}

func (waitEvent *WaitEvent) Process() {
	//TODO
}


