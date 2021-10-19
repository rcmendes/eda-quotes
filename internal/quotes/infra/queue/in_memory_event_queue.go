package queue

import (
	"sync"

	"com.github.rcmendes/eda/quotes/internal/common/eda"
)

type inMemoryEventQueue struct {
	handlers map[string]eda.EventHandler
}

func NewInMemoryEventQueue() eda.EventQueue {
	return &inMemoryEventQueue{
		handlers: make(map[string]eda.EventHandler),
	}
}

func (q *inMemoryEventQueue) Register(eventID string, handler eda.EventHandler) {
	q.handlers[eventID] = handler
}

func (queue inMemoryEventQueue) Publish(evt eda.Event) {
	handler := queue.handlers[evt.EventID()]
	if handler == nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		handler.Handle(evt)
	}()

	wg.Wait()
}
