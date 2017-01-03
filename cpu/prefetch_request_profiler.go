package cpu

import (
	"github.com/mcai/acogo/cpu/uncore"
	"reflect"
)

type L2PrefetchRequestProfiler struct {
	L2Controller                                    *uncore.DirectoryController

	L2PrefetchRequestStates                         map[int32](map[int32]*L2PrefetchRequestState)

	NumL2DemandHits                                 int32
	NumL2DemandMisses                               int32

	NumL2PrefetchHits                               int32
	NumL2PrefetchMisses                             int32

	NumRedundantHitToTransientTagL2PrefetchRequests int32
	NumRedundantHitToCacheL2PrefetchRequests        int32

	NumGoodL2PrefetchRequests                       int32

	NumTimelyL2PrefetchRequests                     int32
	NumLateL2PrefetchRequests                       int32

	NumBadL2PrefetchRequests                        int32

	NumEarlyL2PrefetchRequests                      int32
}

func NewL2PrefetchRequestProfiler(experiment *CPUExperiment) *L2PrefetchRequestProfiler {
	var l2PrefetchRequestProfiler = &L2PrefetchRequestProfiler{
		L2Controller:experiment.MemoryHierarchy.L2Controller(),
	}

	l2PrefetchRequestProfiler.L2PrefetchRequestStates = make(map[int32](map[int32]*L2PrefetchRequestState))

	for set := uint32(0); set < l2PrefetchRequestProfiler.L2Controller.Cache.NumSets(); set++ {
		l2PrefetchRequestProfiler.L2PrefetchRequestStates[int32(set)] = make(map[int32]*L2PrefetchRequestState)

		for way := uint32(0); way < l2PrefetchRequestProfiler.L2Controller.Cache.Assoc(); way++ {
			l2PrefetchRequestProfiler.L2PrefetchRequestStates[int32(set)][int32(way)] = NewL2PrefetchRequestState()
		}
	}

	experiment.BlockingEventDispatcher().AddListener(reflect.TypeOf((*uncore.GeneralCacheControllerServiceNonblockingRequestEvent)(nil)), func(event interface{}) {
		var e = event.(*uncore.GeneralCacheControllerServiceNonblockingRequestEvent)

		if e.CacheController == experiment.MemoryHierarchy.L2Controller() {
			//var requesterIsPrefetchThread =
		}

	})

	return l2PrefetchRequestProfiler
}
