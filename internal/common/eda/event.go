package eda

import (
	"time"
)

type Event interface {
	EventID() string
	CreatedAt() time.Time
	Payload() []byte
}

type EventHandler interface {
	Handle(evt Event)
}
