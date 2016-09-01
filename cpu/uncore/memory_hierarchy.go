package uncore

import (
	"github.com/mcai/acogo/simutil"
	"fmt"
)

type MemoryHierarchyDriver interface {
	CycleAccurateEventQueue() *simutil.CycleAccurateEventQueue
	BlockingEventDispatcher() *simutil.BlockingEventDispatcher
}

type MemoryHierarchy struct {
	Driver                         MemoryHierarchyDriver
	Config                         *MemoryHierarchyConfig

	CurrentMemoryHierarchyAccessId int32
	CurrentCacheCoherenceFlowId    int32

	PendingFlows                   []CacheCoherenceFlow

	MemoryController               *MemoryController
	L2Controller                   *DirectoryController
	L1IControllers                 []*L1IController
	L1DControllers                 []*L1DController

	ITlbs                          []*TranslationLookasideBuffer
	DTlbs                          []*TranslationLookasideBuffer

	P2PReorderBuffers              map[Controller](map[Controller]*P2PReorderBuffer)
}

func NewMemoryHierarchy(driver MemoryHierarchyDriver, config *MemoryHierarchyConfig, numCores int32, numThreadsPerCore int32) *MemoryHierarchy {
	var memoryHierarchy = &MemoryHierarchy{
		Driver:driver,
		Config:config,
	}

	memoryHierarchy.MemoryController = NewMemoryController(memoryHierarchy)

	memoryHierarchy.L2Controller = NewDirectoryController(memoryHierarchy, "l2")
	memoryHierarchy.L2Controller.SetNext(memoryHierarchy.MemoryController)

	for i := int32(0); i < numCores; i++ {
		var l1IController = NewL1IController(memoryHierarchy, fmt.Sprintf("c%d/icache", i))
		l1IController.SetNext(memoryHierarchy.L2Controller)
		memoryHierarchy.L1IControllers = append(memoryHierarchy.L1IControllers, l1IController)

		var l1DController = NewL1DController(memoryHierarchy, fmt.Sprintf("c%d/dcache", i))
		l1DController.SetNext(memoryHierarchy.L2Controller)
		memoryHierarchy.L1DControllers = append(memoryHierarchy.L1DControllers, l1DController)

		for j := int32(0); j < numThreadsPerCore; j++ {
			memoryHierarchy.ITlbs = append(
				memoryHierarchy.ITlbs,
				NewTranslationLookasideBuffer(
					memoryHierarchy,
					fmt.Sprintf("c%dt%d/itlb", i, j),
				),
			)

			memoryHierarchy.DTlbs = append(
				memoryHierarchy.DTlbs,
				NewTranslationLookasideBuffer(
					memoryHierarchy,
					fmt.Sprintf("c%dt%d/dtlb", i, j),
				),
			)
		}
	}

	memoryHierarchy.P2PReorderBuffers = make(map[Controller](map[Controller]*P2PReorderBuffer))

	return memoryHierarchy
}

func (memoryHierarchy *MemoryHierarchy) Transfer(from MemoryDevice, to MemoryDevice, size uint32, onCompletedCallback func()) {
	panic("Unimplemented") //TODO
}

func (memoryHierarchy *MemoryHierarchy) TransferMessage(from Controller, to Controller, size uint32, message CoherenceMessage) {
	if _, ok := memoryHierarchy.P2PReorderBuffers[from]; !ok {
		memoryHierarchy.P2PReorderBuffers[from] = make(map[Controller]*P2PReorderBuffer)
	}

	if _, ok := memoryHierarchy.P2PReorderBuffers[from][to]; !ok {
		memoryHierarchy.P2PReorderBuffers[from][to] = NewP2PReorderBuffer(from, to)
	}

	var p2pReorderBuffer = memoryHierarchy.P2PReorderBuffers[from][to]

	p2pReorderBuffer.Messages = append(p2pReorderBuffer.Messages, message)

	memoryHierarchy.Transfer(from, to, size, func(){
		p2pReorderBuffer.OnDestArrived(message)
	})
}

type P2PReorderBuffer struct {
	Messages               []CoherenceMessage
	From                   Controller
	To                     Controller
	LastCompletedMessageId int32
}

func NewP2PReorderBuffer(from Controller, to Controller) *P2PReorderBuffer {
	var p2pReorderBUffer = &P2PReorderBuffer{
		From:from,
		To:to,
		LastCompletedMessageId:-1,
	}

	return p2pReorderBUffer
}

func (p2pReorderBuffer *P2PReorderBuffer) OnDestArrived(message CoherenceMessage) {
	message.SetDestArrived(true)

	for len(p2pReorderBuffer.Messages) > 0 {
		var message = p2pReorderBuffer.Messages[0]

		if !message.DestArrived() {
			break
		}

		p2pReorderBuffer.Messages = p2pReorderBuffer.Messages[1:]

		p2pReorderBuffer.To.MemoryHierarchy().Driver.CycleAccurateEventQueue().Schedule(
			func() {
				message.Complete()

				p2pReorderBuffer.LastCompletedMessageId = message.Id()

				p2pReorderBuffer.To.ReceiveMessage(message)
			},
			int(p2pReorderBuffer.To.HitLatency()),
		)
	}
}