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

	if description_length > QUOTE_DESCRIPTION_MAX_LENGTH {
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

func (q Quote) ID() uuid.UUID {
	return q.id
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

func (q Quote) Customer() Customer {
	return q.customer
}

func (q Quote) ServiceProvider() ServiceProvider {
	return q.provider
}

func (q Quote) Comments() []Comments {
	comments := make([]Comments, len(q.comments))

	for _, comment := range q.comments {
		comments = append(comments, *comment)
	}

	return comments
}

func (q *Quote) ChangeTitle(title QuoteTitle) {

}

func (q *Quote) ChangeDescription(description string) {

}

func (q *Quote) AddComment(owner CommentOwner, description string) {

}

func (q *Quote) RemoveComment(commentID uuid.UUID) {

}

func (q *Quote) Submit() {

}

func (q *Quote) Estimate() {

}

func (q *Quote) Approve() {

}

func (q *Quote) Reject() {

}

func (q *Quote) Cancel() {

}

func (q Quote) String() string {
	return fmt.Sprintf("{id: '%s', title: '%s', description: '%s', customer: %s, serviceProvider: %s, status: '%s'}",
		q.id, q.title, q.description, q.customer, q.provider, q.status.Name())
}
