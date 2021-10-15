package repository

import (
	"fmt"

	"com.github.rcmendes/eda/quotes/internal/quotes/domain"
	"com.github.rcmendes/eda/quotes/internal/quotes/domain/entity"
	"github.com/google/uuid"
)

type UsersRepository interface {
	Save(user entity.User) error
	FindByID(userID uuid.UUID) (*entity.User, error)
}

func ErrUserNotFound(userID uuid.UUID) error {
	details := fmt.Sprintf("user with id '%s' was not found", userID)
	return domain.NewDomainErrorWithDetails(
		"quotes",
		"user-not-found",
		"user was not found",
		details,
	)
}
