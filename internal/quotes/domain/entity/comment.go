package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comments struct {
	id        uuid.UUID
	owner     Reporter
	createdAt time.Time
	content   string
}

func (c Comments) ID() uuid.UUID {
	return c.id
}

func (c Comments) Owner() Reporter {
	return c.owner
}

func (c Comments) CreatedAt() time.Time {
	return c.createdAt
}

func (c Comments) Content() string {
	return c.content
}
