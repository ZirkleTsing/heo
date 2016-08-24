package simutil

import (
	"reflect"
)

type BlockingEventDispatcher struct {
	Listeners map[reflect.Type]([](func(event interface{})))
}

func NewBlockingEventDispatcher() *BlockingEventDispatcher {
	var blockingEventDispatcher = &BlockingEventDispatcher{
		Listeners:make(map[reflect.Type]([](func(event interface{})))),
	}

	return blockingEventDispatcher
}

func (blockingEventDispatcher *BlockingEventDispatcher) Dispatch(event interface{}) {
	var t = reflect.TypeOf(event)

	if listeners, ok := blockingEventDispatcher.Listeners[t]; ok {
		for _, listener := range listeners {
			listener(event)
		}
	}
}

func (blockingEventDispatcher *BlockingEventDispatcher) AddListener(t reflect.Type, listener func(event interface{})) {
	if _, ok := blockingEventDispatcher.Listeners[t]; !ok {
		blockingEventDispatcher.Listeners[t] = make([](func(event interface{})), 0)
	}

	blockingEventDispatcher.Listeners[t] = append(blockingEventDispatcher.Listeners[t], listener)
}
