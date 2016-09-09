package cpu

type RoundRobinScheduler struct {
	Resources []interface{}
	Predicate func(resource interface{}) bool
	Consume func(resource interface{}) bool
	Quant uint32

	ResourceId int32

	Stalled map[int32]bool
}

func NewRoundRobinScheduler(resources []interface{}, predicate func(resource interface{}) bool, consume func(resource interface{}) bool, quant uint32) *RoundRobinScheduler {
	var scheduler = &RoundRobinScheduler{
		Resources:resources,
		Predicate:predicate,
		Consume:consume,
		Quant:quant,

		ResourceId:0,
		Stalled:make(map[int32]bool),
	}

	return scheduler
}

func (scheduler *RoundRobinScheduler) ConsumeNext() {
	scheduler.ResourceId = scheduler.consumeNext(scheduler.ResourceId)
}

func (scheduler *RoundRobinScheduler) findNext() int32 {
	for i, resource := range scheduler.Resources {
		if scheduler.Predicate(resource) && !scheduler.Stalled[int32(i)] {
			return int32(i)
		}
	}

	return -1
}

func (scheduler *RoundRobinScheduler) consumeNext(resourceId int32) int32 {
	scheduler.Stalled = make(map[int32]bool)

	resourceId = (resourceId + 1) % int32(len(scheduler.Resources))

	for numConsumed := uint32(0); numConsumed < scheduler.Quant; numConsumed++ {
		if stall, ok := scheduler.Stalled[resourceId]; ok && stall && scheduler.Predicate(scheduler.Resources[resourceId]) {
			resourceId = scheduler.findNext()
		}

		if resourceId == -1 {
			break
		}

		if !scheduler.Consume(scheduler.Resources[resourceId]) {
			scheduler.Stalled[resourceId] = true
		}
	}

	return resourceId
}
