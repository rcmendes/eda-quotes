package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"com.github.rcmendes/eda/quotes/internal/quotes/application"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"com.github.rcmendes/eda/quotes/internal/quotes/infra/queue"
	"com.github.rcmendes/eda/quotes/internal/quotes/infra/repository"
	"github.com/google/uuid"
)

var id1 = uuid.New()
var id2 = uuid.New()
var id3 = uuid.New()
var id4 = uuid.New()
var id5 = uuid.New()

type app struct {
	createQuote          application.CreateQuoteHandler
	listQuotesByCustomer application.ListAllQuotesByCustomer
}

func tearUp() *app {
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

	eventQueue := queue.NewInMemoryEventQueue()
	eventLoggerHandler := application.NewEventLoggerHandler()
	eventQueue.Register("quote-created-event", eventLoggerHandler)

	createQuoteHandler := application.NewCreateQuoteHandler(quotesRepo, eventQueue)
	listQuotesByCustomerHandler := application.NewListQuotesByCustomer(quotesRepo, usersRepo)

	return &app{
		createQuote:          createQuoteHandler,
		listQuotesByCustomer: listQuotesByCustomerHandler,
	}

}

func main() {
	app := tearUp()

	ctx := context.Background()

	description := "description 1"
	q := application.NewCreateQuoteCommand("title 1", id1, id3, &description)
	fmt.Println("\n\ncmd TYPE:", reflect.TypeOf(q))
	id, err := app.createQuote.Handle(ctx, *q)
	fmt.Printf("\nID:%s, err: %+v", id, err)

	description = "description 2"
	q = application.NewCreateQuoteCommand("title 2", id2, id4, &description)
	id, err = app.createQuote.Handle(ctx, *q)
	fmt.Printf("\nID:%s, err: %+v", id, err)

	description = "description 3"
	q = application.NewCreateQuoteCommand("title 3", id3, id5, &description)
	id, err = app.createQuote.Handle(ctx, *q)
	fmt.Printf("\nID:%s, err: %+v", id, err)

	fmt.Println("\nLIST OF QUOTES:")

	quotes, err := app.listQuotesByCustomer.Handle(ctx, id1)
	if err != nil {
		log.Fatalln(err)
	}
	for _, q := range *quotes {
		fmt.Printf("\n - %s", q)
	}

	quotes, err = app.listQuotesByCustomer.Handle(ctx, id2)
	if err != nil {
		log.Fatalln(err)
	}
	for _, q := range *quotes {
		fmt.Printf("\n - %s", q)
	}

	quotes, err = app.listQuotesByCustomer.Handle(ctx, id3)
	if err != nil {
		log.Fatalln(err)
	}
	for _, q := range *quotes {
		fmt.Printf("\n - %s", q)
	}

	time.Sleep(1 * time.Second)

}
