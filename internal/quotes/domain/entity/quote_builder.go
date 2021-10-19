package entity

import (
	"errors"

	"github.com/google/uuid"
)

type QuoteBuilder struct {
	id          *uuid.UUID
	title       string
	description string
	customerID  *uuid.UUID
	providerID  *uuid.UUID
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

func (qb *QuoteBuilder) CustomerID(id uuid.UUID) *QuoteBuilder {
	qb.customerID = &id
	return qb
}

func (qb *QuoteBuilder) ServiceProviderID(id uuid.UUID) *QuoteBuilder {
	qb.providerID = &id
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

	if qb.customerID == nil {
		//TODO Create custom error (Domain Error)
		return nil, errors.New("customer ID was not set")
	}

	if qb.providerID == nil {
		//TODO Create custom error (Domain Error)
		return nil, errors.New("service provider ID was not set")
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
		customerID:  *qb.customerID,
		providerID:  *qb.providerID,
		status:      DraftStatus,
	}

	return quote, nil
}
