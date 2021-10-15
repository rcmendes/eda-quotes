package entity

import (
	"github.com/google/uuid"
)

type User struct {
	id    uuid.UUID
	name  string
	email string
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Email() string {
	return u.email
}

type Customer struct {
	User
}

type Reporter struct {
	User
}

type ServiceProvider struct {
	User
}

func NewCustomerFromUser(user User) *Customer {
	return &Customer{
		user,
	}
}

func NewServiceProviderFromUser(user User) *ServiceProvider {
	return &ServiceProvider{
		user,
	}
}

func NewReporterFromUser(user User) *Reporter {
	return &Reporter{
		user,
	}
}
