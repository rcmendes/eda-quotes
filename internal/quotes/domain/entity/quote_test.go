package entity_test

import (
	"errors"
	"strings"
	"testing"

	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"github.com/google/uuid"
)

func TestShouldSuccessWhenCreatingQuoteTitle(t *testing.T) {
	var title string

	title = strings.Repeat("x", 5)
	if _, err := entity.NewQuoteTitle(title); err != nil {
		t.Errorf("[MIN:%d, MAX:%d] Expected NOT to fail due to title length (%d). Got: %v.",
			entity.QUOTE_TITLE_MIN_LENGTH,
			entity.QUOTE_TITLE_MAX_LENGTH,
			len(title),
			err)
	}

	title = strings.Repeat("x", 10)
	if _, err := entity.NewQuoteTitle(title); err != nil {
		t.Errorf("[MIN:%d, MAX:%d] Expected NOT to fail due to title length: %d.",
			entity.QUOTE_TITLE_MIN_LENGTH,
			entity.QUOTE_TITLE_MAX_LENGTH,
			len(title))
	}

	title = strings.Repeat("x", 30)
	if _, err := entity.NewQuoteTitle(title); err != nil {
		t.Errorf("[MIN:%d, MAX:%d] Expected NOT to fail due to title length: %d.",
			entity.QUOTE_TITLE_MIN_LENGTH,
			entity.QUOTE_TITLE_MAX_LENGTH,
			len(title))
	}
}

func TestShouldFailWhenCreatingQuoteTitle(t *testing.T) {
	title := strings.Repeat("x", 4)
	_, err := entity.NewQuoteTitle(title)
	if err == nil || !errors.Is(err, entity.ErrQuoteTitleLengthOutOfBounds) {
		t.Errorf("[MIN:%d, MAX:%d] Expected to fail due to title length: %d.",
			entity.QUOTE_TITLE_MIN_LENGTH,
			entity.QUOTE_TITLE_MAX_LENGTH,
			len(title))
	}

	title = strings.Repeat("x", 31)
	_, err = entity.NewQuoteTitle(title)
	if err == nil || !errors.Is(err, entity.ErrQuoteTitleLengthOutOfBounds) {
		t.Errorf("[MIN:%d, MAX:%d] Expected to fail due to title length: %d.",
			entity.QUOTE_TITLE_MIN_LENGTH,
			entity.QUOTE_TITLE_MAX_LENGTH,
			len(title))
	}
}

func TestShouldSuccessWhenCreatingQuoteDescription(t *testing.T) {
	var description string

	description = ""
	if _, err := entity.NewQuoteDescription(description); err != nil {
		t.Errorf("[MAX:%d] Expected NOT to fail due to description length (%d). Got: %v.",
			entity.QUOTE_DESCRIPTION_MAX_LENGTH,
			len(description),
			err)
	}

	description = strings.Repeat("x", 10)
	if _, err := entity.NewQuoteDescription(description); err != nil {
		t.Errorf("[MAX:%d] Expected NOT to fail due to description length: %d.",
			entity.QUOTE_DESCRIPTION_MAX_LENGTH,
			len(description))
	}

	description = strings.Repeat("x", 4096)
	if _, err := entity.NewQuoteDescription(description); err != nil {
		t.Errorf("[MAX:%d] Expected NOT to fail due to description length: %d.",
			entity.QUOTE_DESCRIPTION_MAX_LENGTH,
			len(description))
	}
}

func TestShouldFailWhenCreatingQuoteDescription(t *testing.T) {
	description := strings.Repeat("x", 4097)
	_, err := entity.NewQuoteDescription(description)
	if err == nil || !errors.Is(err, entity.ErrQuoteDescriptionLengthOutOfBounds) {
		t.Errorf("[MAX:%d] Expected to fail due to description length: %d.",
			entity.QUOTE_DESCRIPTION_MAX_LENGTH,
			len(description))
	}
}

func TestShouldSuccessWhenCreatingQuote(t *testing.T) {
	id := uuid.New()
	title := "Title"
	description := "Description"
	customer := entity.NewCustomer(uuid.New(), "customer@test.com", "Customer Name")
	provider := entity.NewServiceProvider(uuid.New(), "provider@test.com", "Service Provider Name")

	builder := entity.NewQuoteBuilder()
	quote, err := builder.
		ID(id).
		Title(title).
		Description(description).
		Customer(*customer).
		ServiceProvider(*provider).
		Build()

	if err != nil {
		t.Errorf("Expected no errors. Found %v", err)
	}

	if quote.ID() != id {
		t.Errorf("Found Quote ID (%s) != %s.", quote.ID(), id)
	}

	if quote.Title().String() != title {
		t.Errorf("Found Quote Title(%s) != %s.", quote.Title().String(), title)
	}

	if quote.Description().String() != description {
		t.Errorf("Found Quote Description(%s) != %s.", quote.Description().String(), description)
	}

	if !quote.Customer().Equals(customer) {
		t.Errorf("Found Quote Customer(%+v) != (%+v).", quote.Customer(), *customer)
	}

	if !quote.ServiceProvider().Equals(provider) {
		t.Errorf("Found Quote ServiceProvider(%+v) != %+v.", quote.ServiceProvider(), description)
	}

	if !quote.Status().Equals(entity.DraftStatus) {
		t.Errorf("Found Quote Status(%+v) != %+v.", quote.Status(), entity.DraftStatus)
	}

	commentsLength := len(quote.Comments())
	if commentsLength != 0 {
		t.Errorf("Found Quote Comments length different than zero: %d.", commentsLength)
	}
}

// func TestShouldFailWhenCreatingQuote(t *testing.T) {
// 	id := uuid.New()
// 	title := strings.Repeat("x", 5)
// 	description := "Description"
// 	customer := entity.NewCustomer(uuid.New(), "customer@test.com", "Customer Name")
// 	provider := entity.NewServiceProvider(uuid.New(), "provider@test.com", "Service Provider Name")

// 	builder := entity.NewQuoteBuilder()
// 	_, err := builder.
// 		ID(id).
// 		Title(title).
// 		Description(description).
// 		Customer(*customer).
// 		ServiceProvider(*provider).
// 		Build()

// 	if !errors.Is(err, entity.ErrQuoteTitleLengthOutOfBounds) {
// 		t.Errorf("Title length should raise an error. Got: %s", err)
// 	}

// 	title = strings.Repeat("x", 31)
// 	_, err = builder.
// 		ID(id).
// 		Title(title).
// 		Description(description).
// 		Customer(*customer).
// 		ServiceProvider(*provider).
// 		Build()

// 	if !errors.Is(err, entity.ErrQuoteTitleLengthOutOfBounds) {
// 		t.Errorf("Title length should raise an error. Got: %s", err)
// 	}

// 	title = strings.Repeat("x", 6)
// 	description = strings.Repeat("")
// 	_, err = builder.
// 		ID(id).
// 		Title(title).
// 		Description(description).
// 		Customer(*customer).
// 		ServiceProvider(*provider).
// 		Build()

// 	if errors.Is(err, entity.ErrQuoteTitleLengthOutOfBounds) {
// 		t.Errorf("Title length should raise an error. Got: %s", err)
// 	}

// }
