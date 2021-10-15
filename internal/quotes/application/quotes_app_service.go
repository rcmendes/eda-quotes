package application

import (
	"com.github.rcmendes/eda/quotes/internal/common/eda"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/service"
	"github.com/google/uuid"
)

type QuoteApplicationService struct {
	publisher     eda.CommandPublisher
	quotesService service.QuoteService
}

func NewQuoteApplicationService(publisher eda.CommandPublisher, quotesService service.QuoteService) *QuoteApplicationService {
	return &QuoteApplicationService{
		publisher,
		quotesService,
	}
}

func (svc QuoteApplicationService) CreateUser() {

}

func (svc QuoteApplicationService) UpdateUser(userID string) {

}

func (svc QuoteApplicationService) SubmitQuote(quoteID uuid.UUID) {

}
