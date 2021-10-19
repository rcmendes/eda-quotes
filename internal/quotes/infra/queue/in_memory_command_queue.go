package queue

import (
	"sync"

	"com.github.rcmendes/eda/quotes/internal/common/eda"
)

type inMemoryCommandPublisher struct {
	observers map[string]eda.CommandHandler
}

func NewInMemoryCommandQueue() eda.CommandQueue {
	return &inMemoryCommandPublisher{observers: make(map[string]eda.CommandHandler)}
}

func (q *inMemoryCommandPublisher) Register(commandID string, handler eda.CommandHandler) {
	q.observers[commandID] = handler
}

func (q inMemoryCommandPublisher) PublishAndForget(cmd eda.Command) {
	handler := q.observers[cmd.CommandID()]
	if handler == nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		handler.Handle(cmd)
	}()

	wg.Wait()
}

func (q inMemoryCommandPublisher) Publish(cmd eda.Command, result chan interface{}, err chan error) {
	handler := q.observers[cmd.CommandID()]
	if handler == nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer func() {
			close(result)
			close(err)
			wg.Done()
		}()

		res, e := handler.Handle(cmd)
		if res != nil {
			result <- res
		}

		if e != nil {
			err <- e
		}

	}()

	wg.Wait()
}
