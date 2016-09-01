package cpu

import (
	"github.com/mcai/acogo/cpu/uncore"
	"github.com/mcai/acogo/simutil"
)

type MemoryHierarchyCore struct {
	*BaseCore
}

func NewMemoryHierarchyCore(processor *Processor, num int32) *MemoryHierarchyCore {
	var memoryHierarchyCore = &MemoryHierarchyCore{
		BaseCore:NewBaseCore(processor, num),
	}

	return memoryHierarchyCore
}

func (core *BaseCore) L1IController() *uncore.L1IController {
	return core.Processor().Experiment.MemoryHierarchy.L1IControllers[core.Num()]
}

func (core *BaseCore) L1DController() *uncore.L1DController {
	return core.Processor().Experiment.MemoryHierarchy.L1DControllers[core.Num()]
}

func (baseCore *BaseCore) WarmupOneCycle() {
	for _, thread := range baseCore.Threads() {
		thread.WarmupOneCycle()
	}
}

func (core *BaseCore) CanIfetch(thread Thread, virtualAddress uint32) bool {
	var physicalTag = core.L1IController().Cache.GetTag(
		thread.Context().Process.Memory.GetPhysicalAddress(virtualAddress),
	)

	return core.L1IController().CanAccess(uncore.MemoryHierarchyAccessType_IFETCH, physicalTag)
}

func (core *BaseCore) CanLoad(thread Thread, virtualAddress uint32) bool {
	var physicalTag = core.L1DController().Cache.GetTag(
		thread.Context().Process.Memory.GetPhysicalAddress(virtualAddress),
	)

	return core.L1DController().CanAccess(uncore.MemoryHierarchyAccessType_LOAD, physicalTag)
}

func (core *BaseCore) CanStore(thread Thread, virtualAddress uint32) bool {
	var physicalTag = core.L1DController().Cache.GetTag(
		thread.Context().Process.Memory.GetPhysicalAddress(virtualAddress),
	)

	return core.L1DController().CanAccess(uncore.MemoryHierarchyAccessType_STORE, physicalTag)
}

func (core *BaseCore) Ifetch(thread Thread, virtualAddress uint32, virtualPc uint32, onCompletedCallback func()) {
	var physicalAddress = thread.Context().Process.Memory.GetPhysicalAddress(virtualAddress)
	var physicalTag = core.L1IController().Cache.GetTag(physicalAddress)

	var counterPending = simutil.NewCounter(uint32(0))

	counterPending.Increment()

	var alias = core.L1IController().FindAccess(physicalTag)
	var access = core.L1IController().BeginAccess(
		uncore.MemoryHierarchyAccessType_IFETCH,
		thread.Id(),
		int32(virtualPc),
		physicalAddress,
		physicalTag,
		func() {
			counterPending.Decrement()

			if counterPending.Value() == 0 {
				onCompletedCallback()
			}
		},
	)

	if alias == nil {
		counterPending.Increment()

		thread.Itlb().Access(
			access,
			func(){
				counterPending.Decrement()

				if counterPending.Value() == 0 {
					onCompletedCallback()
				}
			},
		)

		core.L1IController().ReceiveIfetch(
			access,
			func(){
				core.L1IController().EndAccess(physicalTag)
			},
		)
	}

	//TODO
}

func (core *BaseCore) Load(thread Thread, virtualAddress uint32, virtualPc uint32, onCompletedCallback func()) {
	var physicalAddress = thread.Context().Process.Memory.GetPhysicalAddress(virtualAddress)
	var physicalTag = core.L1DController().Cache.GetTag(physicalAddress)

	var counterPending = simutil.NewCounter(uint32(0))

	counterPending.Increment()

	var alias = core.L1DController().FindAccess(physicalTag)
	var access = core.L1DController().BeginAccess(
		uncore.MemoryHierarchyAccessType_LOAD,
		thread.Id(),
		int32(virtualPc),
		physicalAddress,
		physicalTag,
		func() {
			counterPending.Decrement()

			if counterPending.Value() == 0 {
				onCompletedCallback()
			}
		},
	)

	if alias == nil {
		counterPending.Increment()

		thread.Dtlb().Access(
			access,
			func(){
				counterPending.Decrement()

				if counterPending.Value() == 0 {
					onCompletedCallback()
				}
			},
		)

		core.L1DController().ReceiveLoad(
			access,
			func(){
				core.L1DController().EndAccess(physicalTag)
			},
		)
	}

	//TODO
}

func (core *BaseCore) Store(thread Thread, virtualAddress uint32, virtualPc uint32, onCompletedCallback func()) {
	var physicalAddress = thread.Context().Process.Memory.GetPhysicalAddress(virtualAddress)
	var physicalTag = core.L1DController().Cache.GetTag(physicalAddress)

	var counterPending = simutil.NewCounter(uint32(0))

	counterPending.Increment()

	var alias = core.L1DController().FindAccess(physicalTag)
	var access = core.L1DController().BeginAccess(
		uncore.MemoryHierarchyAccessType_STORE,
		thread.Id(),
		int32(virtualPc),
		physicalAddress,
		physicalTag,
		func(){
			counterPending.Decrement()

			if counterPending.Value() == 0 {
				onCompletedCallback()
			}
		},
	)

	if alias == nil {
		counterPending.Increment()

		thread.Dtlb().Access(
			access,
			func(){
				counterPending.Decrement()

				if counterPending.Value() == 0 {
					onCompletedCallback()
				}
			},
		)

		core.L1DController().ReceiveStore(
			access,
			func() {
				core.L1DController().EndAccess(physicalTag)
			},
		)
	}

	//TODO
}
