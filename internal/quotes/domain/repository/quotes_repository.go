package repository

import (
	"fmt"

	"com.github.rcmendes/eda/quotes/internal/quotes/domain"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"github.com/google/uuid"
)

type QuotesRepository interface {
	Save(quote entity.Quote) error
	FindByID(quoteID uuid.UUID) (*entity.Quote, error)
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
