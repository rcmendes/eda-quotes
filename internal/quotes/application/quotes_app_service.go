package application

import (
	"com.github.rcmendes/eda/quotes/internal/common/eda"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/service"
	"github.com/google/uuid"
)

type QuotesApplicationService struct {
	publisher     eda.CommandPublisher
	quotesService service.QuotesService
}

func NewQuotesApplicationService(publisher eda.CommandPublisher, quotesService service.QuotesService) *QuotesApplicationService {
	return &QuotesApplicationService{
		publisher,
		quotesService,
	}
}

func (svc QuotesApplicationService) CreateUser() {

}

func (svc QuotesApplicationService) UpdateUser(userID string) {

}

func (svc QuotesApplicationService) SubmitQuote(quoteID uuid.UUID) {

}
