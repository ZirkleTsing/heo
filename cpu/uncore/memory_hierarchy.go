package uncore

import (
	"github.com/mcai/acogo/simutil"
	"fmt"
	"github.com/mcai/acogo/noc"
	"math"
)

type UncoreDriver interface {
	CycleAccurateEventQueue() *simutil.CycleAccurateEventQueue
	BlockingEventDispatcher() *simutil.BlockingEventDispatcher
}

type MemoryHierarchy interface {
	Driver() UncoreDriver
	Config() *UncoreConfig

	CurrentMemoryHierarchyAccessId() int32
	SetCurrentMemoryHierarchyAccessId(currentMemoryHierarchyAccessId int32)

	CurrentCacheCoherenceFlowId() int32
	SetCurrentCacheCoherenceFlowId(currentCacheCoherenceFlowId int32)

	PendingFlows() []CacheCoherenceFlow
	SetPendingFlows(pendingFlows []CacheCoherenceFlow)

	MemoryController() *MemoryController
	L2Controller() *DirectoryController
	L1IControllers() []*L1IController
	L1DControllers() []*L1DController

	ITlbs() []*TranslationLookasideBuffer
	DTlbs() []*TranslationLookasideBuffer

	Transfer(from MemoryDevice, to MemoryDevice, size uint32, onCompletedCallback func())
	TransferMessage(from Controller, to Controller, size uint32, message CoherenceMessage)
}

type BaseMemoryHierarchy struct {
	driver                         UncoreDriver
	config                         *UncoreConfig

	currentMemoryHierarchyAccessId int32
	currentCacheCoherenceFlowId    int32

	pendingFlows                   []CacheCoherenceFlow

	memoryController               *MemoryController
	l2Controller                   *DirectoryController
	l1IControllers                 []*L1IController
	l1DControllers                 []*L1DController

	iTlbs                          []*TranslationLookasideBuffer
	dTlbs                          []*TranslationLookasideBuffer

	p2PReorderBuffers              map[Controller](map[Controller]*P2PReorderBuffer)

	Network                        *noc.Network
	DevicesToNodeIds               map[interface{}]uint32
}

func NewMemoryHierarchy(driver UncoreDriver, config *UncoreConfig, nocConfig *noc.NoCConfig) *BaseMemoryHierarchy {
	var memoryHierarchy = &BaseMemoryHierarchy{
		driver:driver,
		config:config,
		DevicesToNodeIds:make(map[interface{}]uint32),
	}

	memoryHierarchy.memoryController = NewMemoryController(memoryHierarchy)

	memoryHierarchy.l2Controller = NewDirectoryController(memoryHierarchy, "l2")
	memoryHierarchy.l2Controller.SetNext(memoryHierarchy.memoryController)

	for i := int32(0); i < config.NumCores; i++ {
		var l1IController = NewL1IController(memoryHierarchy, fmt.Sprintf("c%d/icache", i))
		l1IController.SetNext(memoryHierarchy.l2Controller)
		memoryHierarchy.l1IControllers = append(memoryHierarchy.l1IControllers, l1IController)

		var l1DController = NewL1DController(memoryHierarchy, fmt.Sprintf("c%d/dcache", i))
		l1DController.SetNext(memoryHierarchy.l2Controller)
		memoryHierarchy.l1DControllers = append(memoryHierarchy.l1DControllers, l1DController)

		for j := int32(0); j < config.NumThreadsPerCore; j++ {
			memoryHierarchy.iTlbs = append(
				memoryHierarchy.iTlbs,
				NewTranslationLookasideBuffer(
					memoryHierarchy,
					fmt.Sprintf("c%dt%d/itlb", i, j),
				),
			)

			memoryHierarchy.dTlbs = append(
				memoryHierarchy.dTlbs,
				NewTranslationLookasideBuffer(
					memoryHierarchy,
					fmt.Sprintf("c%dt%d/dtlb", i, j),
				),
			)
		}
	}

	memoryHierarchy.p2PReorderBuffers = make(map[Controller](map[Controller]*P2PReorderBuffer))

	var numNodes = uint32(0)

	for i, l1IController := range memoryHierarchy.L1IControllers() {
		memoryHierarchy.DevicesToNodeIds[l1IController] = numNodes

		var l1DController = memoryHierarchy.L1DControllers()[i]

		memoryHierarchy.DevicesToNodeIds[l1DController] = numNodes

		numNodes++
	}

	memoryHierarchy.DevicesToNodeIds[memoryHierarchy.L2Controller()] = numNodes

	numNodes++

	memoryHierarchy.DevicesToNodeIds[memoryHierarchy.MemoryController()] = numNodes

	numNodes++

	var width = uint32(math.Sqrt(float64(numNodes)))

	if width * width != numNodes {
		numNodes = (width + 1) * (width + 1)
	}

	nocConfig.NumNodes = int(numNodes)
	nocConfig.MaxInputBufferSize = int(memoryHierarchy.l2Controller.Cache.LineSize() + 8)

	memoryHierarchy.Network = noc.NewNetwork(driver.(noc.NetworkDriver), nocConfig)

	return memoryHierarchy
}

func (memoryHierarchy *BaseMemoryHierarchy) Driver() UncoreDriver {
	return memoryHierarchy.driver
}

func (memoryHierarchy *BaseMemoryHierarchy) Config() *UncoreConfig {
	return memoryHierarchy.config
}

func (memoryHierarchy *BaseMemoryHierarchy) CurrentMemoryHierarchyAccessId() int32 {
	return memoryHierarchy.currentMemoryHierarchyAccessId
}

func (memoryHierarchy *BaseMemoryHierarchy) SetCurrentMemoryHierarchyAccessId(currentMemoryHierarchyAccessId int32) {
	memoryHierarchy.currentMemoryHierarchyAccessId = currentMemoryHierarchyAccessId
}

func (memoryHierarchy *BaseMemoryHierarchy) CurrentCacheCoherenceFlowId() int32 {
	return memoryHierarchy.currentCacheCoherenceFlowId
}

func (memoryHierarchy *BaseMemoryHierarchy) SetCurrentCacheCoherenceFlowId(currentCacheCoherenceFlowId int32) {
	memoryHierarchy.currentCacheCoherenceFlowId = currentCacheCoherenceFlowId
}

func (memoryHierarchy *BaseMemoryHierarchy) PendingFlows() []CacheCoherenceFlow {
	return memoryHierarchy.pendingFlows
}

func (memoryHierarchy *BaseMemoryHierarchy) SetPendingFlows(pendingFlows []CacheCoherenceFlow) {
	memoryHierarchy.pendingFlows = pendingFlows
}

func (memoryHierarchy *BaseMemoryHierarchy) MemoryController() *MemoryController {
	return memoryHierarchy.memoryController
}

func (memoryHierarchy *BaseMemoryHierarchy) L2Controller() *DirectoryController {
	return memoryHierarchy.l2Controller
}

func (memoryHierarchy *BaseMemoryHierarchy) L1IControllers() []*L1IController {
	return memoryHierarchy.l1IControllers
}

func (memoryHierarchy *BaseMemoryHierarchy) L1DControllers() []*L1DController {
	return memoryHierarchy.l1DControllers
}

func (memoryHierarchy *BaseMemoryHierarchy) ITlbs() []*TranslationLookasideBuffer {
	return memoryHierarchy.iTlbs
}

func (memoryHierarchy *BaseMemoryHierarchy) DTlbs() []*TranslationLookasideBuffer {
	return memoryHierarchy.dTlbs
}

func (memoryHierarchy *BaseMemoryHierarchy) Transfer(from MemoryDevice, to MemoryDevice, size uint32, onCompletedCallback func()) {
	var src = memoryHierarchy.DevicesToNodeIds[from]
	var dest = memoryHierarchy.DevicesToNodeIds[to]

	var packet = noc.NewDataPacket(memoryHierarchy.Network, int(src), int(dest), int(size), true, onCompletedCallback)

	memoryHierarchy.Driver().CycleAccurateEventQueue().Schedule(
		func() {
			memoryHierarchy.Network.Receive(packet)
		},
		1,
	)
}

func (memoryHierarchy *BaseMemoryHierarchy) TransferMessage(from Controller, to Controller, size uint32, message CoherenceMessage) {
	if _, ok := memoryHierarchy.p2PReorderBuffers[from]; !ok {
		memoryHierarchy.p2PReorderBuffers[from] = make(map[Controller]*P2PReorderBuffer)
	}

	if _, ok := memoryHierarchy.p2PReorderBuffers[from][to]; !ok {
		memoryHierarchy.p2PReorderBuffers[from][to] = NewP2PReorderBuffer(from, to)
	}

	var p2pReorderBuffer = memoryHierarchy.p2PReorderBuffers[from][to]

	p2pReorderBuffer.Messages = append(p2pReorderBuffer.Messages, message)

	memoryHierarchy.Transfer(from, to, size, func() {
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
	var p2pReorderBuffer = &P2PReorderBuffer{
		From:from,
		To:to,
		LastCompletedMessageId:-1,
	}

	return p2pReorderBuffer
}

func (p2pReorderBuffer *P2PReorderBuffer) OnDestArrived(message CoherenceMessage) {
	message.SetDestArrived(true)

	for len(p2pReorderBuffer.Messages) > 0 {
		var message = p2pReorderBuffer.Messages[0]

		if !message.DestArrived() {
			break
		}

		p2pReorderBuffer.Messages = p2pReorderBuffer.Messages[1:]

		p2pReorderBuffer.To.MemoryHierarchy().Driver().CycleAccurateEventQueue().Schedule(
			func() {
				message.Complete()

				p2pReorderBuffer.LastCompletedMessageId = message.Id()

				p2pReorderBuffer.To.ReceiveMessage(message)
			},
			int(p2pReorderBuffer.To.HitLatency()),
		)
	}
}