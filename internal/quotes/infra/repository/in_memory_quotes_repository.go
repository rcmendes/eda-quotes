package repository

import (
	"context"

	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/repository"
	"github.com/google/uuid"
)

type QuoteInMemoryDB struct {
	quotes map[uuid.UUID]*entity.Quote
}

func NewQuoteInMemoryDB() *QuoteInMemoryDB {
	return &QuoteInMemoryDB{}
}

func (repo *QuoteInMemoryDB) Save(ctx context.Context, quote entity.Quote) error {
	repo.quotes[quote.ID()] = &quote
	return nil
}

func (repo *QuoteInMemoryDB) FindByID(ctx context.Context, id uuid.UUID) (*entity.Quote, error) {
	quote := repo.quotes[id]
	if quote == nil {
		return nil, repository.ErrQuoteNotFound(id)
	}

	return quote, nil
}
