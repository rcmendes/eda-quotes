package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"com.github.rcmendes/eda/quotes/internal/quotes/application"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/handler"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/service"
	"com.github.rcmendes/eda/quotes/internal/quotes/infra/queue"
	"com.github.rcmendes/eda/quotes/internal/quotes/infra/repository"
	"github.com/google/uuid"
)

var id1 = uuid.New()
var id2 = uuid.New()
var id3 = uuid.New()
var id4 = uuid.New()
var id5 = uuid.New()

func tearUp() *application.QuotesApplicationService {
	quotesRepo := repository.NewInMemoryQuotesDB()
	usersRepo := repository.NewInMemoryUsersDB()

	//Just for play
	// Users
	ctx := context.Background()
	u1 := entity.NewUser(id1, "user1@test.com", "User 1")
	usersRepo.Save(ctx, *u1)
	u2 := entity.NewUser(id2, "user2@test.com", "User 2")
	usersRepo.Save(ctx, *u2)
	u3 := entity.NewUser(id3, "user3@test.com", "User 3")
	usersRepo.Save(ctx, *u3)
	u4 := entity.NewUser(id4, "user4@test.com", "User 4")
	usersRepo.Save(ctx, *u4)
	u5 := entity.NewUser(id5, "user5@test.com", "User 5")
	usersRepo.Save(ctx, *u5)

	publisher := queue.NewInMemoryCommandPublisher()
	createQuoteHandler := handler.NewCreateQuoteHandler(usersRepo, quotesRepo)
	publisher.Register("create-quote", createQuoteHandler)

	quotesService := service.NewQuotesService(usersRepo, quotesRepo)

	return application.NewQuotesApplicationService(publisher, quotesService)

}

func main() {
	app := tearUp()

	ctx := context.Background()

	description := "description 1"
	id, err := app.CreateQuote(ctx, "title 1", &description, id1, id4)
	fmt.Printf("\nID:%s, err: %+v", id, err)

	description = "description 2"
	id, err = app.CreateQuote(ctx, "title 2", &description, id2, id5)
	fmt.Printf("\nID:%s, err: %+v", id, err)

	description = "description 3"
	id, err = app.CreateQuote(ctx, "title 3", &description, id3, id4)
	fmt.Printf("\nID:%s, err: %+v", id, err)

	fmt.Println("\nLIST OF QUOTES:")

	quotes, err := app.ListAllQuotesOfCustomer(ctx, id1)
	if err != nil {
		log.Fatalln(err)
	}
	for _, q := range *quotes {
		fmt.Printf("\n - %s", q)
	}

	quotes, err = app.ListAllQuotesOfCustomer(ctx, id2)
	if err != nil {
		log.Fatalln(err)
	}
	for _, q := range *quotes {
		fmt.Printf("\n - %s", q)
	}

	quotes, err = app.ListAllQuotesOfCustomer(ctx, id3)
	if err != nil {
		log.Fatalln(err)
	}
	for _, q := range *quotes {
		fmt.Printf("\n - %s", q)
	}

	time.Sleep(1 * time.Second)

}
