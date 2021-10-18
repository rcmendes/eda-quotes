package entity

import (
	"errors"

	"github.com/google/uuid"
)

type QuoteBuilder struct {
	id          *uuid.UUID
	title       string
	description string
	customer    *Customer
	provider    *ServiceProvider
}

func NewQuoteBuilder() *QuoteBuilder {
	return &QuoteBuilder{}
}

func (qb *QuoteBuilder) Title(title string) *QuoteBuilder {
	qb.title = title
	return qb
}

func (qb *QuoteBuilder) Description(description string) *QuoteBuilder {
	qb.description = description
	return qb
}

func (qb *QuoteBuilder) Customer(customer Customer) *QuoteBuilder {
	qb.customer = &customer
	return qb
}

func (qb *QuoteBuilder) ServiceProvider(provider ServiceProvider) *QuoteBuilder {
	qb.provider = &provider
	return qb
}

func (qb *QuoteBuilder) ID(id uuid.UUID) *QuoteBuilder {
	qb.id = &id
	return qb
}

func (qb QuoteBuilder) Build() (*Quote, error) {
	title, err := NewQuoteTitle(qb.title)
	if err != nil {
		return nil, err
	}

	description, err := NewQuoteDescription(qb.description)
	if err != nil {
		return nil, err
	}

	if qb.customer == nil {
		//TODO Create custom error (Domain Error)
		return nil, errors.New("customer was not set")
	}

	if qb.provider == nil {
		//TODO Create custom error (Domain Error)
		return nil, errors.New("service provider was not set")
	}

	id := qb.id
	if id == nil {
		tmp := uuid.New()
		id = &tmp
	}

	quote := &Quote{
		id:          *id,
		title:       *title,
		description: *description,
		customer:    *qb.customer,
		provider:    *qb.provider,
		status:      DraftStatus,
	}

	return quote, nil
}
