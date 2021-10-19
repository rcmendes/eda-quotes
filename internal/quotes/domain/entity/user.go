package entity

import (
	"fmt"

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

func (u User) Equals(other *User) bool {
	return u.id == other.id && u.email == other.email && u.name == other.name
}

func (u User) String() string {
	return fmt.Sprintf("{id=%s, email=%s, name=%s}", u.id, u.email, u.name)
}

type CustomerID = uuid.UUID
type ServiceProviderID = uuid.UUID
type CommentOwnerID = uuid.UUID

type Customer struct {
	User
}

type CommentOwner struct {
	User
}

type ServiceProvider struct {
	User
}

func NewUser(id uuid.UUID, email string, name string) *User {
	return &User{
		id:    id,
		email: email,
		name:  name,
	}
}

func NewCustomer(id uuid.UUID, email string, name string) *Customer {
	return &Customer{
		User: User{
			id:    id,
			email: email,
			name:  name,
		},
	}
}

func (c Customer) Equals(other *Customer) bool {
	return other != nil && c.User.Equals(&other.User)
}

func NewCustomerFromUser(User User) *Customer {
	return &Customer{
		User,
	}
}

func NewServiceProvider(id uuid.UUID, email string, name string) *ServiceProvider {
	return &ServiceProvider{
		User: User{
			id:    id,
			email: email,
			name:  name,
		},
	}
}

func (sp ServiceProvider) Equals(other *ServiceProvider) bool {
	return other != nil && sp.User.Equals(&other.User)
}

func NewServiceProviderFromUser(User User) *ServiceProvider {
	return &ServiceProvider{
		User,
	}
}

func NewCommentOwner(id uuid.UUID, email string, name string) *CommentOwner {
	return &CommentOwner{
		User: User{
			id:    id,
			email: email,
			name:  name,
		},
	}
}

func (co CommentOwner) Equals(other *CommentOwner) bool {
	return other != nil && co.User.Equals(&other.User)
}

func NewCommentOwnerFromUser(User User) *CommentOwner {
	return &CommentOwner{
		User,
	}
}
