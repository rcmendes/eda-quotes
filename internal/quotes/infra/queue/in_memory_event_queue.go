package queue

import (
	"sync"

	"com.github.rcmendes/eda/quotes/internal/common/eda"
)

type InMemoryEventPublisher struct {
	observers map[string][]eda.EventHandler
}

func NewInMemoryEventPublisher() *InMemoryEventPublisher {
	return &InMemoryEventPublisher{}
}

func (q *InMemoryEventPublisher) Register(commandID string, handler eda.EventHandler) {
	list := q.observers[commandID]

	if list == nil {
		list = make([]eda.EventHandler, 1)
	}

	list = append(list, handler)

	q.observers[commandID] = list
}

func (q InMemoryEventPublisher) Publish(cmd eda.Event) {
	observers := q.observers[cmd.EventID()]
	if observers == nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(observers))

	for _, observer := range observers {
		go func(wg *sync.WaitGroup, handler eda.EventHandler) {
			defer wg.Done()
			handler.Handle(cmd)
		}(&wg, observer)
	}

	wg.Wait()
}
