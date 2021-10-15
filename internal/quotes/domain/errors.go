package domain

import (
	"fmt"
)

type DomainError struct {
	domain  string
	id      string
	message string
	details *string
}

func NewDomainErrorWithDetails(domain string, id string, message string, details string) *DomainError {
	return &DomainError{
		domain: domain, id: id, message: message, details: &details,
	}
}

func NewDomainError(domain string, id string, message string) *DomainError {
	return &DomainError{
		domain: domain, id: id, message: message, details: nil,
	}
}

func (err DomainError) Error() string {
	return fmt.Sprintf("[%s] %s", err.id, err.message)
}

func (err DomainError) Domain() string {
	return err.domain
}

func (err DomainError) Message() string {
	return err.message
}

func (err DomainError) Details() *string {
	return err.details
}

func (err DomainError) ID() string {
	return err.id
}
