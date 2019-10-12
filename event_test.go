package event

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type A struct {
	Emitter
	name string
}

func TestBus(t *testing.T) {
	a := A{name: `123`}

	wg := sync.WaitGroup{}

	wg.Add(1)

	a.On("test", func(name string) {
		time.Sleep(3 * time.Second)
		wg.Done()
		fmt.Println("name:" + name)

	})
	a.Emit("test", a.name)
	wg.Wait()
}
