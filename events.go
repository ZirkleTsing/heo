package acogo

type CycleAccurateEvent struct {
	eventQueue *CycleAccurateEventQueue
	When       int
	Action     func()
	Id         int
}

type CycleAccurateEventQueue struct {
	Events         map[int]([]*CycleAccurateEvent)
	PerCycleEvents []func()
	CurrentCycle   int
	currentEventId int
}

func NewCycleAccurateEventQueue() *CycleAccurateEventQueue {
	var q = CycleAccurateEventQueue{
		Events:make(map[int]([]*CycleAccurateEvent)),
	}

	return &q
}

func (q *CycleAccurateEventQueue) Schedule(action func(), delay int) {
	q.currentEventId++

	var event = &CycleAccurateEvent{
		eventQueue:q,
		When:q.CurrentCycle + delay,
		Action:action,
		Id:q.currentEventId,
	}

	q.Events[event.When] = append(q.Events[event.When], event)
}

func (q *CycleAccurateEventQueue) AddPerCycleEvent(action func()) {
	q.PerCycleEvents = append(q.PerCycleEvents, action)
}

func (q *CycleAccurateEventQueue) AdvanceOneCycle() {
	if events, exists := q.Events[q.CurrentCycle]; exists {
		for _, event := range events {
			event.Action()
		}

		delete(q.Events, q.CurrentCycle)
	}

	for _, e := range q.PerCycleEvents {
		e()
	}

	q.CurrentCycle++
}