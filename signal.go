package acogo

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

	signalMask.Signals[signal / 32] = SetBit32(signalMask.Signals[signal / 32], signal % 32)
}

func (signalMask *SignalMask) Clear(signal uint32) {
	if signal < 1 || signal > MAX_SIGNAL {
		return
	}

	signal--

	signalMask.Signals[signal / 32] = ClearBit32(signalMask.Signals[signal / 32], signal % 32)
}

func (signalMask *SignalMask) Contains(signal uint32) bool {
	if signal < 1 || signal > MAX_SIGNAL {
		return false
	}

	signal--

	return GetBit32(signalMask.Signals[signal / 32], signal % 32) != 0
}

func (signalMask *SignalMask) LoadFrom(memory *PagedMemory, virtualAddress uint64) {
	for i := uint64(0); i < MAX_SIGNAL / 32; i++ {
		signalMask.Signals[i] = memory.ReadWord(virtualAddress + i * 4)
	}
}

func (signalMask *SignalMask) SaveTo(memory *PagedMemory, virtualAddress uint64) {
	for i := uint64(0); i < MAX_SIGNAL / 32; i++ {
		memory.WriteWord(virtualAddress + i * 4, signalMask.Signals[i])
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

func (signalAction *SignalAction) LoadFrom(memory *PagedMemory, virtualAddress uint64) {
	signalAction.Flags = memory.ReadWord(virtualAddress)
	signalAction.Handler = memory.ReadWord(virtualAddress + SignalAction_HANDLER_OFFSET)
	signalAction.Restorer = memory.ReadWord(virtualAddress + SignalAction_RESTORER_OFFSET)
	signalAction.Mask.LoadFrom(memory, virtualAddress + SignalAction_MASK_OFFSET)
}

func (signalAction *SignalAction) SaveTo(memory *PagedMemory, virtualAddress uint64) {
	memory.WriteWord(virtualAddress, signalAction.Flags)
	memory.WriteWord(virtualAddress + SignalAction_HANDLER_OFFSET, signalAction.Handler)
	memory.WriteWord(virtualAddress + SignalAction_RESTORER_OFFSET, signalAction.Restorer)
	signalAction.Mask.SaveTo(memory, virtualAddress + SignalAction_MASK_OFFSET)
}
