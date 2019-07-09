package event

import (
	"github.com/asaskevich/EventBus"
	"sync"
)

var mutex sync.Mutex = sync.Mutex{}

type IEventEmitter interface {
	On(topic string, fn interface{})
	Once(topic string, fn interface{})
	Off(topic string, fn interface{})
	Emit(topic string, args ...interface{})
}

type EventEmitter struct {
	bus EventBus.Bus
}

func (e *EventEmitter) getBus() EventBus.Bus {
	if e.bus == nil {
		mutex.Lock()
		if e.bus == nil {
			e.bus = EventBus.New()
		}
		mutex.Unlock()
	}
	return e.bus
}

func (e *EventEmitter) On(topic string, fn interface{}) {
	e.getBus().SubscribeAsync(topic, fn, false)
}

func (e *EventEmitter) Once(topic string, fn interface{}) {
	e.getBus().SubscribeOnceAsync(topic, fn)
}

func (e *EventEmitter) Off(topic string, fn interface{}) {
	e.getBus().Unsubscribe(topic, fn)
}

func (e *EventEmitter) Emit(topic string, args ...interface{}) {
	e.getBus().Publish(topic, args...)
}
