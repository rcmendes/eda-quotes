package application

import (
	"com.github.rcmendes/eda/quotes/internal/common/eda"
	"github.com/google/uuid"
)

type CreateQuoteCommand struct {
	ID                uuid.UUID
	Title             string
	CustomerID        uuid.UUID
	ServiceProviderID uuid.UUID
}

func (cmd CreateQuoteCommand) CommandID() string {
	return "create-quote"
}

type QuoteApplicationService struct {
	publisher eda.Publisher
}

func NewQuoteApplicationService(publisher eda.Publisher) *QuoteApplicationService {
	return &QuoteApplicationService{
		publisher,
	}
}

func (svc QuoteApplicationService) CreateQuote(
	title string,
	customerID uuid.UUID,
	serviceProviderID uuid.UUID) (*uuid.UUID, error) {

	id := uuid.New()

	cmd := CreateQuoteCommand{
		ID:                id,
		Title:             title,
		CustomerID:        customerID,
		ServiceProviderID: serviceProviderID,
	}

	go svc.publisher.Notify(cmd)

	return &id, nil
}

/*
func (svc QuoteApplicationService) CreateQuote(cmd CreateQuoteCommand) (*uuid.UUID, error) {
	user, err := svc.usersRepo.FindByID(cmd.CustomerID)
	if err != nil {
		//TODO Add log
		return nil, err
	}
	customer := entity.NewCustomerFromUser(*user)

	user, err = svc.usersRepo.FindByID(cmd.CustomerID)
	if err != nil {
		//TODO Add log
		return nil, err
	}
	provider := entity.NewServiceProviderFromUser(*user)

	builder := entity.NewQuoteBuilder()
	builder.Title(cmd.Title).Customer(*customer).ServiceProvider(*provider)

	if cmd.Description != nil {
		builder.Description(*cmd.Description)
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
*/
func (svc QuoteApplicationService) CreateUser() {

}

func (svc QuoteApplicationService) UpdateUser(userID string) {

}

func (svc QuoteApplicationService) SubmitQuote(quoteID uuid.UUID) {

}
