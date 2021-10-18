package queue

import (
	"sync"

	"com.github.rcmendes/eda/quotes/internal/common/eda"
)

type InMemoryCommandPublisher struct {
	observers map[string][]eda.CommandHandler
}

func NewInMemoryCommandPublisher() *InMemoryCommandPublisher {
	return &InMemoryCommandPublisher{observers: make(map[string][]eda.CommandHandler)}
}

func (q *InMemoryCommandPublisher) Register(commandID string, handler eda.CommandHandler) {
	list := q.observers[commandID]

	if len(list) == 0 {
		list = make([]eda.CommandHandler, 0)
	}

	list = append(list, handler)

	q.observers[commandID] = list
}

func (q InMemoryCommandPublisher) Publish(cmd eda.Command) {
	observers := q.observers[cmd.CommandID()]
	if len(observers) == 0 {
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(observers))

	for _, observer := range observers {
		go func(wg *sync.WaitGroup, handler eda.CommandHandler) {
			defer wg.Done()
			handler.Handle(cmd)
		}(&wg, observer)
	}

	wg.Wait()
}
