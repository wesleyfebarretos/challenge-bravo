package testdata

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/app/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/infra/repository/user_repository"
)

func CreateUser() (entity.User, error) {
	user, err := user_repository.New().Create(context.TODO(), entity.User{
		FirstName: "John",
		LastName:  "Doe",
		Password:  "123",
		Email:     "johndoe@test.com",
		Role:      "user",
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateAdminUser() (entity.User, error) {
	user, err := user_repository.New().Create(context.TODO(), entity.User{
		FirstName: "John",
		LastName:  "Doe",
		Password:  "123",
		Email:     "johndoe@test.com",
		Role:      "admin",
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
