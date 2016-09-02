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

	DumpPendingFlowTree()
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

	p2pReorderBuffers              map[Controller](map[Controller]*P2PReorderBuffer)

	Network                        *noc.Network
	DevicesToNodeIds               map[interface{}]uint32
}

func NewBaseMemoryHierarchy(driver UncoreDriver, config *UncoreConfig, nocConfig *noc.NoCConfig) *BaseMemoryHierarchy {
	var baseMemoryHierarchy = &BaseMemoryHierarchy{
		driver:driver,
		config:config,
		DevicesToNodeIds:make(map[interface{}]uint32),
	}

	baseMemoryHierarchy.memoryController = NewMemoryController(baseMemoryHierarchy)

	baseMemoryHierarchy.l2Controller = NewDirectoryController(baseMemoryHierarchy, "l2")
	baseMemoryHierarchy.l2Controller.SetNext(baseMemoryHierarchy.memoryController)

	for i := int32(0); i < config.NumCores; i++ {
		var l1IController = NewL1IController(baseMemoryHierarchy, fmt.Sprintf("c%d/icache", i))
		l1IController.SetNext(baseMemoryHierarchy.l2Controller)
		baseMemoryHierarchy.l1IControllers = append(baseMemoryHierarchy.l1IControllers, l1IController)

		var l1DController = NewL1DController(baseMemoryHierarchy, fmt.Sprintf("c%d/dcache", i))
		l1DController.SetNext(baseMemoryHierarchy.l2Controller)
		baseMemoryHierarchy.l1DControllers = append(baseMemoryHierarchy.l1DControllers, l1DController)

		for j := int32(0); j < config.NumThreadsPerCore; j++ {
			baseMemoryHierarchy.iTlbs = append(
				baseMemoryHierarchy.iTlbs,
				NewTranslationLookasideBuffer(
					baseMemoryHierarchy,
					fmt.Sprintf("c%dt%d/itlb", i, j),
				),
			)

			baseMemoryHierarchy.dTlbs = append(
				baseMemoryHierarchy.dTlbs,
				NewTranslationLookasideBuffer(
					baseMemoryHierarchy,
					fmt.Sprintf("c%dt%d/dtlb", i, j),
				),
			)
		}
	}

	baseMemoryHierarchy.p2pReorderBuffers = make(map[Controller](map[Controller]*P2PReorderBuffer))

	var numNodes = uint32(0)

	for i, l1IController := range baseMemoryHierarchy.L1IControllers() {
		baseMemoryHierarchy.DevicesToNodeIds[l1IController] = numNodes

		var l1DController = baseMemoryHierarchy.L1DControllers()[i]

		baseMemoryHierarchy.DevicesToNodeIds[l1DController] = numNodes

		numNodes++
	}

	baseMemoryHierarchy.DevicesToNodeIds[baseMemoryHierarchy.L2Controller()] = numNodes

	numNodes++

	baseMemoryHierarchy.DevicesToNodeIds[baseMemoryHierarchy.MemoryController()] = numNodes

	numNodes++

	var width = uint32(math.Sqrt(float64(numNodes)))

	if width * width != numNodes {
		numNodes = (width + 1) * (width + 1)
	}

	nocConfig.NumNodes = int(numNodes)
	nocConfig.MaxInputBufferSize = int(baseMemoryHierarchy.l2Controller.Cache.LineSize() + 8)

	baseMemoryHierarchy.Network = noc.NewNetwork(driver.(noc.NetworkDriver), nocConfig)

	return baseMemoryHierarchy
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) Driver() UncoreDriver {
	return baseMemoryHierarchy.driver
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) Config() *UncoreConfig {
	return baseMemoryHierarchy.config
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) CurrentMemoryHierarchyAccessId() int32 {
	return baseMemoryHierarchy.currentMemoryHierarchyAccessId
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) SetCurrentMemoryHierarchyAccessId(currentMemoryHierarchyAccessId int32) {
	baseMemoryHierarchy.currentMemoryHierarchyAccessId = currentMemoryHierarchyAccessId
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) CurrentCacheCoherenceFlowId() int32 {
	return baseMemoryHierarchy.currentCacheCoherenceFlowId
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) SetCurrentCacheCoherenceFlowId(currentCacheCoherenceFlowId int32) {
	baseMemoryHierarchy.currentCacheCoherenceFlowId = currentCacheCoherenceFlowId
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) PendingFlows() []CacheCoherenceFlow {
	return baseMemoryHierarchy.pendingFlows
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) SetPendingFlows(pendingFlows []CacheCoherenceFlow) {
	baseMemoryHierarchy.pendingFlows = pendingFlows
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) MemoryController() *MemoryController {
	return baseMemoryHierarchy.memoryController
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) L2Controller() *DirectoryController {
	return baseMemoryHierarchy.l2Controller
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) L1IControllers() []*L1IController {
	return baseMemoryHierarchy.l1IControllers
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) L1DControllers() []*L1DController {
	return baseMemoryHierarchy.l1DControllers
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) ITlbs() []*TranslationLookasideBuffer {
	return baseMemoryHierarchy.iTlbs
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) DTlbs() []*TranslationLookasideBuffer {
	return baseMemoryHierarchy.dTlbs
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) Transfer(from MemoryDevice, to MemoryDevice, size uint32, onCompletedCallback func()) {
	//var src = baseMemoryHierarchy.DevicesToNodeIds[from]
	//var dest = baseMemoryHierarchy.DevicesToNodeIds[to]
	//
	//var packet = noc.NewDataPacket(baseMemoryHierarchy.Network, int(src), int(dest), int(size), true, onCompletedCallback)
	//
	//baseMemoryHierarchy.Driver().CycleAccurateEventQueue().Schedule(
	//	func() {
	//		baseMemoryHierarchy.Network.Receive(packet)
	//	},
	//	1,
	//)

	onCompletedCallback() //TODO
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) TransferMessage(from Controller, to Controller, size uint32, message CoherenceMessage) {
	if _, ok := baseMemoryHierarchy.p2pReorderBuffers[from]; !ok {
		baseMemoryHierarchy.p2pReorderBuffers[from] = make(map[Controller]*P2PReorderBuffer)
	}

	if _, ok := baseMemoryHierarchy.p2pReorderBuffers[from][to]; !ok {
		baseMemoryHierarchy.p2pReorderBuffers[from][to] = NewP2PReorderBuffer(from, to)
	}

	var p2pReorderBuffer = baseMemoryHierarchy.p2pReorderBuffers[from][to]

	p2pReorderBuffer.Messages = append(p2pReorderBuffer.Messages, message)

	baseMemoryHierarchy.Transfer(from, to, size, func() {
		p2pReorderBuffer.OnDestArrived(message)
	})
}

func (baseMemoryHierarchy *BaseMemoryHierarchy) DumpPendingFlowTree() {
	for _, pendingFlow := range baseMemoryHierarchy.pendingFlows {
		simutil.PrintNode(pendingFlow)
		fmt.Println()
	}

	fmt.Println()
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