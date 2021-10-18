package repository

import (
	"context"
	"fmt"

	"com.github.rcmendes/eda/quotes/internal/quotes/domain"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"github.com/google/uuid"
)

type QuotesRepository interface {
	Save(ctx context.Context, quote entity.Quote) error
	FindByID(ctx context.Context, quoteID uuid.UUID) (*entity.Quote, error)
	FindByCustomerID(ctx context.Context, customerID uuid.UUID) (*[]entity.Quote, error)
}

func ErrQuoteNotFound(quoteID uuid.UUID) error {
	details := fmt.Sprintf("quote with id '%s' was not found", quoteID)
	return domain.NewDomainErrorWithDetails(
		"quotes",
		"quote-not-found",
		"quote was not found",
		details,
	)
}
