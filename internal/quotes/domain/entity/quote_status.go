package entity

type QuoteStatus interface {
	Name() string
}

type quoteStatus struct {
	name string
}

func (qs quoteStatus) Name() string {
	return qs.name
}

var DraftStatus QuoteStatus = &quoteStatus{"draft"}
var SubmittedStatus QuoteStatus = &quoteStatus{"submitted"}
var EstimatingStatus QuoteStatus = &quoteStatus{"estimating"}
var WaitingApprovalStatus QuoteStatus = &quoteStatus{"waiting-approval"}
var ApprovedStatus QuoteStatus = &quoteStatus{"approved"}
var RejectedStatus QuoteStatus = &quoteStatus{"rejected"}
var CanceledStatus QuoteStatus = &quoteStatus{"canceled"}
