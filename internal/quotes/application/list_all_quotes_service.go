package application

import (
	"context"

	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"github.com/google/uuid"
)

func (svc QuotesApplicationService) ListAllQuotesOfCustomer(
	ctx context.Context,
	customerID uuid.UUID) (*[]entity.Quote, error) {

	list, err := svc.quotesService.ListQuotesByCustomerID(ctx, customerID)
	if err != nil {
		//TODO Handle non Domain Errors
		return nil, err
	}

	return list, nil
}
