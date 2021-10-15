package main

import (
	"com.github.rcmendes/eda/quotes/internal/quotes/application"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/handler"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/service"
	"com.github.rcmendes/eda/quotes/internal/quotes/infra/queue"
	"com.github.rcmendes/eda/quotes/internal/quotes/infra/repository"
)

func tearUp() *application.QuotesApplicationService {
	quotesRepo := repository.NewInMemoryQuotesDB()
	usersRepo := repository.NewInMemoryUsersDB()

	publisher := queue.NewInMemoryCommandPublisher()
	createQuoteHandler := handler.NewCreateQuoteHandler(usersRepo, quotesRepo)
	publisher.Register("create-quote", createQuoteHandler)

	quotesService := service.NewQuotesService(usersRepo, quotesRepo)

	return application.NewQuotesApplicationService(publisher, quotesService)

}
