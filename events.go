package acogo

type CycleAccurateEvent struct {
	eventQueue *CycleAccurateEventQueue
	When       int64
	Action     func()
	Id         int64
}

type CycleAccurateEventQueue struct {
	Events         map[int64]([]*CycleAccurateEvent)
	PerCycleEvents []func()
	CurrentCycle   int64
	currentEventId int64
}

func NewCycleAccurateEventQueue() *CycleAccurateEventQueue {
	var q = CycleAccurateEventQueue{
		Events:make(map[int64]([]*CycleAccurateEvent)),
	}

	return &q
}

func (q *CycleAccurateEventQueue) Schedule(action func(), delay int) {
	q.currentEventId++

	var event = &CycleAccurateEvent{
		eventQueue:q,
		When:q.CurrentCycle + int64(delay),
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