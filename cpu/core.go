package cpu

import "github.com/mcai/acogo/cpu/uncore"

type Core interface {
	Processor() *Processor
	Threads() []Thread
	AddThread(thread Thread)
	Num() int32
	FastForwardOneCycle()

	L1IController() *uncore.L1IController
	L1DController() *uncore.L1DController

	CanIfetch(thread Thread, virtualAddress uint32) bool
	CanLoad(thread Thread, virtualAddress uint32) bool
	CanStore(thread Thread, virtualAddress uint32) bool

	Ifetch(thread Thread, virtualAddress uint32, virtualPc uint32, onCompletedCallback func())
	Load(thread Thread, virtualAddress uint32, virtualPc uint32, onCompletedCallback func())
	Store(thread Thread, virtualAddress uint32, virtualPc uint32, onCompletedCallback func())

	WarmupOneCycle()
}

type BaseCore struct {
	processor *Processor
	threads   []Thread
	num       int32
}

func NewBaseCore(processor *Processor, num int32) *BaseCore {
	var baseCore = &BaseCore{
		processor:processor,
		num:num,
	}

	return baseCore
}

func (baseCore *BaseCore) Processor() *Processor {
	return baseCore.processor
}

func (baseCore *BaseCore) Threads() []Thread {
	return baseCore.threads
}

func (baseCore *BaseCore) AddThread(thread Thread) {
	baseCore.threads = append(baseCore.threads, thread)
}

func (baseCore *BaseCore) Num() int32 {
	return baseCore.num
}

func (baseCore *BaseCore) FastForwardOneCycle() {
	for _, thread := range baseCore.Threads() {
		thread.FastForwardOneCycle()
	}
}
