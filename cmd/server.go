package main

import (
	"com.github.rcmendes/eda/quotes/internal/quotes/application"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/handler"
	"com.github.rcmendes/eda/quotes/internal/quotes/infra/queue"
	"com.github.rcmendes/eda/quotes/internal/quotes/infra/repository"
)

func tearUp() *application.QuoteApplicationService {
	quotesRepo := repository.NewQuoteInMemoryDB()
	usersRepo := repository.NewUserInMemoryDB()

	publisher := queue.NewInMemoryCommandPublisher()

	createQuoteHandler := handler.NewCreateQuoteHandler(usersRepo, quotesRepo)
	publisher.Register("create-quote", createQuoteHandler)

	return application.NewQuoteApplicationService(publisher)

}
