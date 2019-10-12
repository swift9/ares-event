package event

import (
	"github.com/asaskevich/EventBus"
	"sync"
)

var mutex = sync.Mutex{}

type IEmitter interface {
	On(topic string, fn interface{})
	OnSync(topic string, fn interface{})
	Once(topic string, fn interface{})
	OnceSync(topic string, fn interface{})
	Off(topic string, fn interface{})
	Emit(topic string, args ...interface{})
}

type Emitter struct {
	bus EventBus.Bus
}

func (e *Emitter) getBus() EventBus.Bus {
	if e.bus == nil {
		mutex.Lock()
		if e.bus == nil {
			e.bus = EventBus.New()
		}
		mutex.Unlock()
	}
	return e.bus
}

func (e *Emitter) On(topic string, fn interface{}) {
	e.getBus().SubscribeAsync(topic, fn, false)
}

func (e *Emitter) OnSync(topic string, fn interface{}) {
	e.getBus().Subscribe(topic, fn)
}

func (e *Emitter) Once(topic string, fn interface{}) {
	e.getBus().SubscribeOnceAsync(topic, fn)
}

func (e *Emitter) OnceSync(topic string, fn interface{}) {
	e.getBus().SubscribeOnce(topic, fn)
}

func (e *Emitter) Off(topic string, fn interface{}) {
	e.getBus().Unsubscribe(topic, fn)
}

func (e *Emitter) Emit(topic string, args ...interface{}) {
	e.getBus().Publish(topic, args...)
}
