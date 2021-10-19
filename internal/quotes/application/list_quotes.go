package application

import (
	"context"

	"com.github.rcmendes/eda/quotes/internal/quotes/domain/repository"
	"github.com/google/uuid"
)

type UserModel struct {
	ID    uuid.UUID
	Email string
	Name  string
}

type QuoteModel struct {
	ID              uuid.UUID
	Title           string
	Description     string
	Status          string
	Customer        UserModel
	ServiceProvider UserModel
}

type ListAllQuotesByCustomer interface {
	Handle(ctx context.Context, customerID uuid.UUID) (*[]QuoteModel, error)
}

type listQuotesByCustomer struct {
	quotesRepo repository.QuotesRepository
	usersRepo  repository.UsersRepository
}

func NewListQuotesByCustomer(quotesRepo repository.QuotesRepository, usersRepo repository.UsersRepository) ListAllQuotesByCustomer {
	return &listQuotesByCustomer{
		quotesRepo,
		usersRepo,
	}
}

func (svc listQuotesByCustomer) Handle(
	ctx context.Context,
	customerID uuid.UUID) (*[]QuoteModel, error) {

	quotes, err := svc.quotesRepo.FindByCustomerID(ctx, customerID)
	if err != nil {
		//TODO Handle non Domain Errors
		return nil, err
	}

	list := make([]QuoteModel, len(*quotes))

	//FIXME Replace for a join if database
	for i, q := range *quotes {
		customer, err := svc.usersRepo.FindByID(ctx, q.CustomerID())
		if err != nil {
			return nil, repository.ErrUserNotFound(q.CustomerID())
		}

		provider, err := svc.usersRepo.FindByID(ctx, q.ServiceProviderID())
		if err != nil {
			return nil, repository.ErrUserNotFound(q.ServiceProviderID())
		}

		qm := QuoteModel{
			ID:          q.ID(),
			Title:       q.Title().String(),
			Description: q.Description().String(),
			Customer: UserModel{
				ID:    customer.ID(),
				Email: customer.Email(),
				Name:  customer.Name(),
			},
			ServiceProvider: UserModel{
				ID:    provider.ID(),
				Email: provider.Email(),
				Name:  provider.Name(),
			},
		}
		list[i] = qm
	}

	return &list, nil
}
