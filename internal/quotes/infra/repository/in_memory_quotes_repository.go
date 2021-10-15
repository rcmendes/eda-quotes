package repository

import (
	"context"

	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/repository"
	"github.com/google/uuid"
)

type inMemoryQuotesDB struct {
	quotes map[uuid.UUID]*entity.Quote
}

func NewInMemoryQuotesDB() repository.QuotesRepository {
	return &inMemoryQuotesDB{}
}

func (repo *inMemoryQuotesDB) Save(ctx context.Context, quote entity.Quote) error {
	repo.quotes[quote.ID()] = &quote
	return nil
}

func (repo *inMemoryQuotesDB) FindByID(ctx context.Context, id uuid.UUID) (*entity.Quote, error) {
	quote := repo.quotes[id]
	if quote == nil {
		return nil, repository.ErrQuoteNotFound(id)
	}

	return quote, nil
}
