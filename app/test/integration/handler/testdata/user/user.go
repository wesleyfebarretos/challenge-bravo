package user

import (
	"context"
	"time"

	"github.com/wesleyfebarretos/challenge-bravo/app/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/infra/repository/user_repository"
)

func Create() (entity.Currency, error) {
	user := entity.User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Active:    true,
		FirstName: "John",
		LastName:  "Doe",
		Password:  "integrationtest",
		Email:     "integrationtest@test.com",
		Role:      "user",
	}

	user, err := user_repository.New().Create(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateAdmin() (entity.Currency, error) {
	user := entity.User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Active:    true,
		FirstName: "John",
		LastName:  "Doe",
		Password:  "integrationtest",
		Email:     "integrationtest@test.com",
		Role:      "admin",
	}

	user, err := user_repository.New().Create(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
