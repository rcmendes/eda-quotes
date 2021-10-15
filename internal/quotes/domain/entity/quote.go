package entity

import (
	"fmt"
	"strings"

	"com.github.rcmendes/eda/quotes/internal/quotes/domain"
	"github.com/google/uuid"
)

const QUOTE_TITLE_MIN_LENGTH = 5
const QUOTE_TITLE_MAX_LENGTH = 30
const QUOTE_DESCRIPTION_MAX_LENGTH = 4096

var ErrQuoteTitleLengthOutOfBounds = domain.NewDomainError(
	"quotes",
	"quote-title-length-out-of-bounds",
	fmt.Sprintf("Quote's title must have a length between %d and %d characters.",
		QUOTE_TITLE_MIN_LENGTH, QUOTE_TITLE_MAX_LENGTH))

var ErrQuoteDescriptionLengthOutOfBounds = domain.NewDomainError(
	"quotes",
	"quote-description-length-out-of-bounds",
	fmt.Sprintf("Quote's title must have a maximum length of %d characters.",
		QUOTE_DESCRIPTION_MAX_LENGTH))

type QuoteTitle struct {
	value string
}

func (qt QuoteTitle) String() string {
	return qt.value
}

type QuoteDescription struct {
	value string
}

func (qt QuoteDescription) String() string {
	return qt.value
}

func NewQuoteTitle(title string) (*QuoteTitle, error) {
	title_length := len(strings.TrimSpace(title))

	if title_length < QUOTE_TITLE_MIN_LENGTH || title_length > QUOTE_TITLE_MAX_LENGTH {
		return nil, ErrQuoteTitleLengthOutOfBounds
	}

	qt := &QuoteTitle{title}

	return qt, nil
}

func NewQuoteDescription(description string) (*QuoteDescription, error) {
	description_length := len(strings.TrimSpace(description))

	if description_length > QUOTE_TITLE_MAX_LENGTH {
		return nil, ErrQuoteDescriptionLengthOutOfBounds
	}

	qt := &QuoteDescription{description}

	return qt, nil
}

type Quote struct {
	id          uuid.UUID
	title       QuoteTitle
	description QuoteDescription
	status      QuoteStatus
	customer    Customer
	provider    ServiceProvider
	comments    []*Comments
}

func (qa Quote) ID() uuid.UUID {
	return qa.id
}

func (q Quote) Title() QuoteTitle {
	return q.title
}

func (q Quote) Description() QuoteDescription {
	return q.description
}

func (q Quote) Status() QuoteStatus {
	return q.status
}

func (qa Quote) Customer() Customer {
	return qa.customer
}

func (qa Quote) ServiceProvider() ServiceProvider {
	return qa.provider
}

func (qa Quote) Comments() []Comments {
	comments := make([]Comments, len(qa.comments))

	for _, comment := range qa.comments {
		comments = append(comments, *comment)
	}

	return comments
}

func (qa *Quote) ChangeTitle(title QuoteTitle) {

}

func (qa *Quote) ChangeDescription(description string) {

}

func (qa *Quote) AddComment(owner Reporter, description string) {

}

func (qa *Quote) RemoveComment(commentID uuid.UUID) {

}

func (qa *Quote) Submit() {

}

func (qa *Quote) Estimate() {

}

func (qa *Quote) Approve() {

}

func (qa *Quote) Reject() {

}

func (qa *Quote) Cancel() {

}
