package service

import (
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/repository"
	"github.com/google/uuid"
)

type QuoteService struct {
	usersRepo  repository.UsersRepository
	quotesRepo repository.QuotesRepository
}

func (svc QuoteService) CreateQuote(
	title string,
	customerID uuid.UUID,
	serviceProviderID uuid.UUID,
	description *string) (*uuid.UUID, error) {
	user, err := svc.usersRepo.FindByID(customerID)
	if err != nil {
		//TODO Add log
		return nil, err
	}
	customer := entity.NewCustomerFromUser(*user)

	user, err = svc.usersRepo.FindByID(serviceProviderID)
	if err != nil {
		//TODO Add log
		return nil, err
	}
	provider := entity.NewServiceProviderFromUser(*user)

	builder := entity.NewQuoteBuilder()
	builder.Title(title).Customer(*customer).ServiceProvider(*provider)

	if description != nil {
		builder.Description(*description)
	}

	//TODO Builder should return a set of errors
	quote, err := builder.Build()
	if err != nil {
		//TODO Add log
		return nil, err
	}

	err = svc.quotesRepo.Save(*quote)
	if err != nil {
		//TODO Add log
		return nil, err
	}

	id := quote.ID()
	return &id, nil
}

func (svc QuoteService) CreateUser() {

}

func (svc QuoteService) UpdateUser(userID string) {

}

func (svc QuoteService) SubmitQuote(quoteID uuid.UUID) {

}
