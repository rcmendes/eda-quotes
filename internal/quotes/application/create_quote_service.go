package application

import (
	"context"

	"github.com/google/uuid"
)

type CreateQuoteCommand struct {
	id                *uuid.UUID
	title             string
	description       *string
	customerID        uuid.UUID
	serviceProviderID uuid.UUID
}

func (cmd CreateQuoteCommand) CommandID() string {
	return "create-quote"
}
func (cmd CreateQuoteCommand) ID() *uuid.UUID {
	return cmd.id
}
func (cmd CreateQuoteCommand) Title() string {
	return cmd.title
}

func (cmd CreateQuoteCommand) Description() *string {
	return cmd.description
}

func (cmd CreateQuoteCommand) CustomerID() uuid.UUID {
	return cmd.customerID
}

func (cmd CreateQuoteCommand) ServiceProviderID() uuid.UUID {
	return cmd.serviceProviderID
}

func (svc QuotesApplicationService) CreateQuote(
	ctx context.Context,
	title string,
	description *string,
	customerID uuid.UUID,
	serviceProviderID uuid.UUID) (*uuid.UUID, error) {

	id := uuid.New()

	//TODO Use Value Objects for validation and consistency? e.g.: QuoteTitle, QuoteDescription
	cmd := CreateQuoteCommand{
		id:                &id,
		title:             title,
		description:       description,
		customerID:        customerID,
		serviceProviderID: serviceProviderID,
	}

	if _, err := svc.quotesService.CreateQuote(ctx, cmd); err != nil {
		return nil, err
	}

	go svc.publisher.Publish(cmd)

	return &id, nil
}
