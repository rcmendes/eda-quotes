package service

import (
	"context"

	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/repository"
	"github.com/google/uuid"
)

type CreateQuoteInput interface {
	ID() *uuid.UUID
	Title() string
	Description() *string
	CustomerID() uuid.UUID
	ServiceProviderID() uuid.UUID
}

type QuoteService struct {
	usersRepo  repository.UsersRepository
	quotesRepo repository.QuotesRepository
}

func (svc QuoteService) CreateQuote(ctx context.Context, input CreateQuoteInput) (*uuid.UUID, error) {
	user, err := svc.usersRepo.FindByID(ctx, input.CustomerID())
	if err != nil {
		//TODO Add log
		return nil, err
	}
	customer := entity.NewCustomerFromUser(*user)

	user, err = svc.usersRepo.FindByID(ctx, input.ServiceProviderID())
	if err != nil {
		//TODO Add log
		return nil, err
	}
	provider := entity.NewServiceProviderFromUser(*user)

	builder := entity.NewQuoteBuilder()
	builder.Title(input.Title()).Customer(*customer).ServiceProvider(*provider)

	id := input.ID()
	if id != nil {
		builder.ID(*id)
	}

	description := input.Description()
	if description != nil {
		builder.Description(*description)
	}

	//TODO Builder should return a set of errors
	quote, err := builder.Build()
	if err != nil {
		//TODO Add log
		return nil, err
	}

	err = svc.quotesRepo.Save(ctx, *quote)
	if err != nil {
		//TODO Add log
		return nil, err
	}

	//TODO publish a QuoteCreatedEvent

	*id = quote.ID()

	return id, nil
}

func (svc QuoteService) SubmitQuote(quoteID uuid.UUID) {

}
