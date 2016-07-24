package acogo

import (
	"container/heap"
)

type CycleAccurateEvent struct {
	eventQueue *CycleAccurateEventQueue
	When       int
	Action     func()
	Id         int
	index      int
}

type CycleAccurateEventQueue struct {
	Events         []*CycleAccurateEvent
	PerCycleEvents []func()
	CurrentCycle   int
	currentEventId int
}

func (q CycleAccurateEventQueue) Len() int {
	return len(q.Events)
}

func (q CycleAccurateEventQueue) Less(i, j int) bool {
	var x = q.Events[i]
	var y = q.Events[j]
	return x.When < y.When || (x.When == y.When && x.Id < y.Id)
}

func (q CycleAccurateEventQueue) Swap(i, j int) {
	q.Events[i], q.Events[j] = q.Events[j], q.Events[i]
	q.Events[i].index = i
	q.Events[j].index = j
}

func (q *CycleAccurateEventQueue) Push(x interface{}) {
	n := len((*q).Events)
	item := x.(*CycleAccurateEvent)
	item.index = n
	(*q).Events = append((*q).Events, item)
}

func (q *CycleAccurateEventQueue) Pop() interface{} {
	old := (*q).Events
	n := len(old)
	item := old[n - 1]
	item.index = -1
	(*q).Events = old[0 : n - 1]
	return item
}

func NewCycleAccurateEventQueue() *CycleAccurateEventQueue {
	return &CycleAccurateEventQueue{}
}

func (q *CycleAccurateEventQueue) Schedule(action func(), delay int) {
	q.currentEventId++

	var event = &CycleAccurateEvent{
		eventQueue:q,
		When:q.CurrentCycle + delay,
		Action:action,
		Id:q.currentEventId,
	}

	heap.Push(q, event)
}

func (q *CycleAccurateEventQueue) AddPerCycleEvent(action func()) {
	q.PerCycleEvents = append(q.PerCycleEvents, action)
}

func (q *CycleAccurateEventQueue) AdvanceOneCycle() {
	for q.Len() > 0 {
		var value = q.Pop()
		var event *CycleAccurateEvent = value.(*CycleAccurateEvent)

		if event.When > q.CurrentCycle {
			q.Push(value)
			break
		}

		event.Action()
	}

	for i := 0; i < len(q.PerCycleEvents); i++ {
		q.PerCycleEvents[i]()
	}

	q.CurrentCycle++
}