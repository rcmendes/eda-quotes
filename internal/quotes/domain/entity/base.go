package entity

import (
	"github.com/google/uuid"
)

type DomainEntity interface {
	ID() uuid.UUID
}
