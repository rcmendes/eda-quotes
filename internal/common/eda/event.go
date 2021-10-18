package eda

import (
	"time"
)

type Event interface {
	EventID() string
	CreatedAt() time.Time
	Payload() string
}

type EventHandler interface {
	Handle(evt Event)
}
