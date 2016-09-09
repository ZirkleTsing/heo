package cpu

import "github.com/mcai/acogo/simutil"

type BranchPredictorType string

const (
	BranchPredictorType_PERFECT = BranchPredictorType("PERFECT")

	BranchPredictorType_TAKEN = BranchPredictorType("TAKEN")

	BranchPredictorType_NOT_TAKEN = BranchPredictorType("NOT_TAKEN")

	BranchPredictorType_TWO_BIT = BranchPredictorType("TWO_BIT")
)

const (
	BRANCH_SHIFT = 2
)

type BranchTargetBufferEntry struct {
	Source uint32
	Target uint32
}

func NewBranchTargetBufferEntry() *BranchTargetBufferEntry {
	var branchTargetBufferEntry = &BranchTargetBufferEntry{
	}

	return branchTargetBufferEntry
}

type BranchTargetBuffer struct {
	NumSets uint32
	Assoc   uint32
	Entries []([]*BranchTargetBufferEntry)
}

func NewBranchTargetBuffer(numSets uint32, assoc uint32) *BranchTargetBuffer {
	var branchTargetBuffer = &BranchTargetBuffer{
		NumSets:numSets,
		Assoc:assoc,
	}

	for i := uint32(0); i < numSets; i++ {
		var entriesPerSet []*BranchTargetBufferEntry

		for j := uint32(0); j < assoc; j++ {
			entriesPerSet = append(entriesPerSet, NewBranchTargetBufferEntry())
		}

		branchTargetBuffer.Entries = append(branchTargetBuffer.Entries, entriesPerSet)
	}

	return branchTargetBuffer
}

func (branchTargetBuffer *BranchTargetBuffer) GetSet(branchAddress uint32) uint32 {
	return branchAddress >> BRANCH_SHIFT & (branchTargetBuffer.NumSets - 1)
}

func (branchTargetBuffer *BranchTargetBuffer) Lookup(branchAddress uint32) *BranchTargetBufferEntry {
	var set = branchTargetBuffer.GetSet(branchAddress)

	for _, entry := range branchTargetBuffer.Entries[set] {
		if entry.Source == branchAddress {
			return entry
		}

	}

	return nil
}

func (branchTargetBuffer *BranchTargetBuffer) Update(branchAddress uint32, branchTarget uint32, taken bool) {
	if !taken {
		return
	}

	var set = branchTargetBuffer.GetSet(branchAddress)

	var entryFound *BranchTargetBufferEntry

	for _, entry := range branchTargetBuffer.Entries[set] {
		if entry.Source == branchAddress {
			entryFound = entry
			break
		}
	}

	if entryFound == nil {
		entryFound = branchTargetBuffer.Entries[set][branchTargetBuffer.Assoc - 1]
		entryFound.Source = branchAddress
	}

	entryFound.Target = branchTarget

	branchTargetBuffer.removeFromEntries(set, entryFound)

	branchTargetBuffer.Entries[set] = append(
		[]*BranchTargetBufferEntry{entryFound},
		branchTargetBuffer.Entries[set]...,
	)
}

func (branchTargetBuffer *BranchTargetBuffer) removeFromEntries(set uint32, entryToRemove *BranchTargetBufferEntry) {
	var entriesToReserve []*BranchTargetBufferEntry

	for _, entry := range branchTargetBuffer.Entries[set] {
		if entry != entryToRemove {
			entriesToReserve = append(entriesToReserve, entry)
		}
	}

	branchTargetBuffer.Entries[set] = entriesToReserve
}

type ReturnAddressStack struct {
	size    uint32
	top     uint32
	entries []*BranchTargetBufferEntry
}

func NewReturnAddressStack(size uint32) *ReturnAddressStack {
	var returnAddressStack = &ReturnAddressStack{
		size:size,
		top:size - 1,
	}

	for i := uint32(0); i < size; i++ {
		returnAddressStack.entries = append(
			returnAddressStack.entries,
			NewBranchTargetBufferEntry(),
		)
	}

	return returnAddressStack
}

func (returnAddressStack *ReturnAddressStack) Size() uint32 {
	return returnAddressStack.size
}

func (returnAddressStack *ReturnAddressStack) Top() uint32 {
	if returnAddressStack.size > 0 {
		return returnAddressStack.top
	}

	return 0
}

func (returnAddressStack *ReturnAddressStack) Recover(top uint32) {
	returnAddressStack.top = top
}

func (returnAddressStack *ReturnAddressStack) Push(branchAddress uint32) {
	returnAddressStack.top = (returnAddressStack.top + 1) % returnAddressStack.size
	returnAddressStack.entries[returnAddressStack.top].Target = branchAddress + 8
}

func (returnAddressStack *ReturnAddressStack) Pop() uint32 {
	var target = returnAddressStack.entries[returnAddressStack.top].Target
	returnAddressStack.top = (returnAddressStack.top + returnAddressStack.size - 1) % returnAddressStack.size
	return target
}

type BranchPredictorUpdate struct {
	SaturatingCounter *simutil.SaturatingCounter
	Ras               bool
}

func NewBranchPredictorUpdate() *BranchPredictorUpdate {
	var branchPredictorUpdate = &BranchPredictorUpdate{
	}

	return branchPredictorUpdate
}

type BranchPredictor struct {
	BranchTargetBuffer *BranchTargetBuffer
	ReturnAddressStack *ReturnAddressStack

	Size               uint32
	SaturatingCounters []*simutil.SaturatingCounter

	NumHits            int32
	NumMisses          int32
}

func NewBranchPredictor(branchTargetBufferNumSets uint32, branchTargetBufferAssoc uint32, returnAddressStackSize uint32, size uint32) *BranchPredictor {
	var branchPredictor = &BranchPredictor{
		BranchTargetBuffer:NewBranchTargetBuffer(branchTargetBufferNumSets, branchTargetBufferAssoc),
		ReturnAddressStack:NewReturnAddressStack(returnAddressStackSize),
		Size:size,
	}

	var flipFlop = uint32(1)

	for i := uint32(0); i < size; i++ {
		branchPredictor.SaturatingCounters = append(
			branchPredictor.SaturatingCounters,
			simutil.NewSaturatingCounter(0, 2, 3, flipFlop),
		)

		flipFlop = 3 - flipFlop
	}

	return branchPredictor
}

func (branchPredictor *BranchPredictor) NumAccesses() int32 {
	return branchPredictor.NumHits + branchPredictor.NumMisses
}

func (branchPredictor *BranchPredictor) HitRatio() float32 {
	if branchPredictor.NumAccesses() > 0 {
		return float32(branchPredictor.NumHits) / float32(branchPredictor.NumAccesses())
	} else {
		return 0
	}
}

func (branchPredictor *BranchPredictor) GetSaturatingCounter(branchAddress uint32) *simutil.SaturatingCounter {
	var index =(branchAddress >> BRANCH_SHIFT) & (branchPredictor.Size - 1)

	return branchPredictor.SaturatingCounters[index]
}

func (branchPredictor *BranchPredictor) Predict(branchAddress uint32, mnemonic *Mnemonic, branchPredictorUpdate *BranchPredictorUpdate) (uint32, uint32) {
	if mnemonic.StaticInstType == StaticInstType_COND {
		branchPredictorUpdate.SaturatingCounter = branchPredictor.GetSaturatingCounter(branchAddress)
	}

	var returnAddressStackRecoverTop = branchPredictor.ReturnAddressStack.Top()

	if mnemonic.StaticInstType == StaticInstType_FUNC_RET && branchPredictor.ReturnAddressStack.Size() > 0 {
		branchPredictorUpdate.Ras = true
		return branchPredictor.ReturnAddressStack.Pop(), returnAddressStackRecoverTop
	}

	if mnemonic.StaticInstType == StaticInstType_FUNC_CALL && branchPredictor.ReturnAddressStack.Size() > 0 {
		branchPredictor.ReturnAddressStack.Push(branchAddress)
	}

	if mnemonic.StaticInstType != StaticInstType_COND || branchPredictorUpdate.SaturatingCounter.Taken() {
		var branchTargetBufferEntry = branchPredictor.BranchTargetBuffer.Lookup(branchAddress)

		if branchTargetBufferEntry != nil {
			return branchTargetBufferEntry.Target, returnAddressStackRecoverTop
		} else {
			return 0, returnAddressStackRecoverTop
		}
	} else {
		return 0, returnAddressStackRecoverTop
	}
}

func (branchPredictor *BranchPredictor) Update(branchAddress uint32, branchTarget uint32, taken bool, correct bool, mnemonic *Mnemonic, branchPredictorUpdate *BranchPredictorUpdate) {
	if correct {
		branchPredictor.NumHits++
	} else {
		branchPredictor.NumMisses++
	}

	if mnemonic.StaticInstType == StaticInstType_FUNC_RET {
		if !branchPredictorUpdate.Ras {
			return
		}
	}

	if mnemonic.StaticInstType == StaticInstType_COND {
		branchPredictorUpdate.SaturatingCounter.Update(taken)
	}

	branchPredictor.BranchTargetBuffer.Update(branchAddress, branchTarget, taken)
}