package os

import (
	"github.com/mcai/acogo/cpu/mem"
	"github.com/mcai/acogo/cpu/cpuutil"
)

const MAX_SIGNAL = 64

type SignalMask struct {
	Signals []uint32
}

func NewSignalMask() *SignalMask {
	var signalMask = &SignalMask{
		Signals:make([]uint32, MAX_SIGNAL),
	}

	return signalMask
}

func (signalMask *SignalMask) Set(signal uint32) {
	if signal < 1 || signal > MAX_SIGNAL {
		return
	}

	signal--

	signalMask.Signals[signal / 32] = cpuutil.SetBit32(signalMask.Signals[signal / 32], signal % 32)
}

func (signalMask *SignalMask) Clear(signal uint32) {
	if signal < 1 || signal > MAX_SIGNAL {
		return
	}

	signal--

	signalMask.Signals[signal / 32] = cpuutil.ClearBit32(signalMask.Signals[signal / 32], signal % 32)
}

func (signalMask *SignalMask) Contains(signal uint32) bool {
	if signal < 1 || signal > MAX_SIGNAL {
		return false
	}

	signal--

	return cpuutil.GetBit32(signalMask.Signals[signal / 32], signal % 32) != 0
}

func (signalMask *SignalMask) LoadFrom(memory *mem.PagedMemory, virtualAddress uint64) {
	for i := uint64(0); i < MAX_SIGNAL / 32; i++ {
		signalMask.Signals[i] = memory.ReadWordAt(virtualAddress + i * 4)
	}
}

func (signalMask *SignalMask) SaveTo(memory *mem.PagedMemory, virtualAddress uint64) {
	for i := uint64(0); i < MAX_SIGNAL / 32; i++ {
		memory.WriteWordAt(virtualAddress + i * 4, signalMask.Signals[i])
	}
}

type SignalMasks struct {
	Pending *SignalMask
	Blocked *SignalMask
	Backup  *SignalMask
}

func NewSignalMasks() *SignalMasks {
	var signalMasks = &SignalMasks{
		Pending:NewSignalMask(),
		Blocked:NewSignalMask(),
		Backup:NewSignalMask(),
	}

	return signalMasks
}

const (
	SignalAction_HANDLER_OFFSET = 4
	SignalAction_RESTORER_OFFSET = 136
	SignalAction_MASK_OFFSET = 8
)

type SignalAction struct {
	Flags    uint32
	Handler  uint32
	Restorer uint32
	Mask     *SignalMask
}

func NewSignalAction() *SignalAction {
	var signalAction = &SignalAction{
		Mask:NewSignalMask(),
	}

	return signalAction
}

func (signalAction *SignalAction) LoadFrom(memory *mem.PagedMemory, virtualAddress uint64) {
	signalAction.Flags = memory.ReadWordAt(virtualAddress)
	signalAction.Handler = memory.ReadWordAt(virtualAddress + SignalAction_HANDLER_OFFSET)
	signalAction.Restorer = memory.ReadWordAt(virtualAddress + SignalAction_RESTORER_OFFSET)
	signalAction.Mask.LoadFrom(memory, virtualAddress + SignalAction_MASK_OFFSET)
}

func (signalAction *SignalAction) SaveTo(memory *mem.PagedMemory, virtualAddress uint64) {
	memory.WriteWordAt(virtualAddress, signalAction.Flags)
	memory.WriteWordAt(virtualAddress + SignalAction_HANDLER_OFFSET, signalAction.Handler)
	memory.WriteWordAt(virtualAddress + SignalAction_RESTORER_OFFSET, signalAction.Restorer)
	signalAction.Mask.SaveTo(memory, virtualAddress + SignalAction_MASK_OFFSET)
}
