package handler

import (
	"com.github.rcmendes/eda/quotes/internal/common/eda"
	"com.github.rcmendes/eda/quotes/internal/quotes/application"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/repository"
)

type CreateQuoteHandler struct {
	usersRepo  repository.UsersRepository
	quotesRepo repository.QuotesRepository
}

func NewCreateQuoteHandler(usersRepo repository.UsersRepository,
	quotesRepo repository.QuotesRepository) *CreateQuoteHandler {
	return &CreateQuoteHandler{
		usersRepo, quotesRepo,
	}
}

func (svc CreateQuoteHandler) Handle(cmd eda.Command) {
	switch cmd.(type) {
	case application.CreateQuoteCommand:
		svc.createQuote(cmd.(application.CreateQuoteCommand))
	default:
		return
	}
}

func (svc CreateQuoteHandler) createQuote(cmd application.CreateQuoteCommand) {
	user, err := svc.usersRepo.FindByID(cmd.CustomerID)
	if err != nil {
		//TODO Add log
		return nil, err
	}
	customer := entity.NewCustomerFromUser(*user)

	user, err = svc.usersRepo.FindByID(cmd.ServiceProviderID)
	if err != nil {
		//TODO Add log
		return nil, err
	}
	provider := entity.NewServiceProviderFromUser(*user)

	builder := entity.NewQuoteBuilder()

	builder.ID(cmd.ID)
	builder.Title(cmd.Title).Customer(*customer).ServiceProvider(*provider)

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
