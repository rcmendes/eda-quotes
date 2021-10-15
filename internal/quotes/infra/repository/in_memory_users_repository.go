package repository

import (
	"context"

	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/repository"
	"github.com/google/uuid"
)

type UserInMemoryDB struct {
	users map[uuid.UUID]*entity.User
}

func NewUserInMemoryDB() *UserInMemoryDB {
	return &UserInMemoryDB{}
}

func (repo *UserInMemoryDB) Save(ctx context.Context, user entity.User) error {
	repo.users[user.ID()] = &user
	return nil
}

func (repo *UserInMemoryDB) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user := repo.users[id]
	if user == nil {
		return nil, repository.ErrUserNotFound(id)
	}
	return user, nil
}
