package handler

import (
	"context"
	"fmt"

	"com.github.rcmendes/eda/quotes/internal/common/eda"
	"com.github.rcmendes/eda/quotes/internal/quotes/application"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/repository"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/service"
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
		svc.createQuote(cmd.(service.CreateQuoteInput))
	default:
		return
	}
}

func (svc CreateQuoteHandler) createQuote(cmd service.CreateQuoteInput) {
	ctx := context.Background()

	user, err := svc.usersRepo.FindByID(ctx, cmd.CustomerID())
	if err != nil {
		//TODO Add log
		//TODO publish in a queue?
		// return nil, err
	}
	customer := entity.NewCustomerFromUser(*user)

	user, err = svc.usersRepo.FindByID(ctx, cmd.ServiceProviderID())
	if err != nil {
		//TODO Add log
		//TODO publish in a queue?
		// return nil, err
	}
	provider := entity.NewServiceProviderFromUser(*user)

	builder := entity.NewQuoteBuilder()

	//TODO How to return that to the user?
	// builder.ID(cmd.ID())
	builder.Title(cmd.Title()).Customer(*customer).ServiceProvider(*provider)

	//TODO Builder should return a set of errors
	quote, err := builder.Build()
	if err != nil {
		//TODO Add log
		//TODO publish in a queue?
		// return nil, err
	}

	fmt.Println("\n\nProcessing (Handling) Quote:", quote)

	// err = svc.quotesRepo.Save(ctx, *quote)
	// if err != nil {
	//TODO Add log
	//TODO publish in a queue?
	// return nil, err
	// }

	//TODO publish a QuoteCreatedEvent

	// id := quote.ID()
	// return &id, nil
}
