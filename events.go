package acogo

import (
	"container/heap"
)

type CycleAccurateEvent struct {
	eventQueue *CycleAccurateEventQueue
	when       int
	action     func()
	id         int
	index      int
}

type CycleAccurateEventQueue struct {
	events         []*CycleAccurateEvent
	perCycleEvents []func()
	currentCycle   int
	currentEventId int
}

func (q CycleAccurateEventQueue) Len() int {
	return len(q.events)
}

func (q CycleAccurateEventQueue) Less(i, j int) bool {
	var x = q.events[i]
	var y = q.events[j]
	return x.when < y.when || (x.when == y.when && x.id < y.id)
}

func (q CycleAccurateEventQueue) Swap(i, j int) {
	q.events[i], q.events[j] = q.events[j], q.events[i]
	q.events[i].index = i
	q.events[j].index = j
}

func (q *CycleAccurateEventQueue) Push(x interface{}) {
	n := len((*q).events)
	item := x.(*CycleAccurateEvent)
	item.index = n
	(*q).events = append((*q).events, item)
}

func (q *CycleAccurateEventQueue) Pop() interface{} {
	old := (*q).events
	n := len(old)
	item := old[n - 1]
	item.index = -1
	(*q).events = old[0 : n - 1]
	return item
}

func NewCycleAccurateEventQueue() *CycleAccurateEventQueue {
	return &CycleAccurateEventQueue{}
}

func (q *CycleAccurateEventQueue) Schedule(action func(), delay int) {
	q.currentEventId++

	var event = &CycleAccurateEvent{
		eventQueue:q,
		when:q.currentCycle + delay,
		action:action,
		id:q.currentEventId,
	}

	heap.Push(q, event)
}

func (q *CycleAccurateEventQueue) AddPerCycleEvent(action func()) {
	q.perCycleEvents = append(q.perCycleEvents, action)
}

func (q *CycleAccurateEventQueue) AdvanceOneCycle() {
	for q.Len() > 0 {
		var value = q.Pop()
		var event *CycleAccurateEvent = value.(*CycleAccurateEvent)

		if event.when > q.currentCycle {
			q.Push(value)
			break
		}

		event.action()
	}

	for i := 0; i < len(q.perCycleEvents); i++ {
		q.perCycleEvents[i]()
	}

	q.currentCycle++
}