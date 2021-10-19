package application

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"com.github.rcmendes/eda/quotes/internal/common/eda"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
)

type userModel struct {
	ID string `json:"id"`
}

type quoteModel struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Status          string    `json:"status"`
	Customer        userModel `json:"customer"`
	ServiceProvider userModel `json:"service_provider"`
}

//TODO Apply composition? (Generic struct with default methods and Aggregates with toJson method)
type quoteCreatedEvent struct {
	quote                entity.Quote
	createdAt            time.Time
	parseQuoteToJsonOnce sync.Once
	quoteAsJson          string
}

func NewQuoteCreatedEvent(quote entity.Quote, createdAt time.Time) eda.Event {
	return &quoteCreatedEvent{
		quote:     quote,
		createdAt: createdAt,
	}
}

func (evt quoteCreatedEvent) EventID() string {
	return "quote-created-event"
}

func (evt quoteCreatedEvent) CreatedAt() time.Time {
	return evt.createdAt
}

func (evt *quoteCreatedEvent) Payload() string {
	evt.parseQuoteToJsonOnce.Do(func() {
		q := evt.quote
		customerID := evt.quote.CustomerID()
		providerID := evt.quote.ServiceProviderID()
		qm := quoteModel{
			ID:          q.ID().String(),
			Title:       q.Title().String(),
			Description: q.Description().String(),
			Customer: userModel{
				ID: customerID.String(),
			},
			ServiceProvider: userModel{
				ID: providerID.String(),
			},
		}
		payload, _ := json.Marshal(qm)

		evt.quoteAsJson = string(payload)
	})
	return evt.quoteAsJson
}

type eventLoggerHandler struct{}

func NewEventLoggerHandler() eda.EventHandler {
	return &eventLoggerHandler{}
}

func (q *eventLoggerHandler) Handle(evt eda.Event) {
	fmt.Println("\n\ncmd TYPE:", reflect.TypeOf(evt))
	log.Default().Println("Published event:", evt.EventID(), "\n\t", evt.Payload())
}
