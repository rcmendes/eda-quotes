package application

import (
	"context"
	"time"

	"com.github.rcmendes/eda/quotes/internal/common/eda"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/repository"
	"github.com/google/uuid"
)

type CreateQuoteCommand struct {
	title             string
	description       *string
	customerID        uuid.UUID
	serviceProviderID uuid.UUID
}

func NewCreateQuoteCommand(title string,
	customerID uuid.UUID,
	serviceProviderID uuid.UUID,
	description *string) *CreateQuoteCommand {
	return &CreateQuoteCommand{
		title:             title,
		customerID:        customerID,
		serviceProviderID: serviceProviderID,
		description:       description,
	}
}

func (cmd CreateQuoteCommand) CommandID() string {
	return "create-quote"
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

type CreateQuoteHandler interface {
	Handle(ctx context.Context, cmd CreateQuoteCommand) (*uuid.UUID, error)
}

type createQuoteHandler struct {
	quotesRepo repository.QuotesRepository
	eventQueue eda.EventQueue
}

func NewCreateQuoteHandler(quotesRepo repository.QuotesRepository, eventQueue eda.EventQueue) CreateQuoteHandler {
	return &createQuoteHandler{
		quotesRepo: quotesRepo,
		eventQueue: eventQueue,
	}
}

func (handler createQuoteHandler) Handle(ctx context.Context, cmd CreateQuoteCommand) (*uuid.UUID, error) {
	builder := entity.NewQuoteBuilder()
	builder.
		Title(cmd.Title()).
		CustomerID(cmd.CustomerID()).
		ServiceProviderID(cmd.ServiceProviderID())

	description := cmd.Description()
	if description != nil {
		builder.Description(*description)
	}

	//TODO Builder should return a set of errors
	quote, err := builder.Build()
	if err != nil {
		//TODO Add log
		return nil, err
	}

	err = handler.quotesRepo.Save(ctx, *quote)
	if err != nil {
		//TODO Add log
		return nil, err
	}

	event := NewQuoteCreatedEvent(*quote, time.Now())
	handler.eventQueue.Publish(event)

	id := quote.ID()

	return &id, nil
}
