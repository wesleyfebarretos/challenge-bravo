package entity

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/internal/enum"
)

type User struct {
	Active *bool
	Name   string
	Email  string
	Role   enum.Role
	ID     int
}

type UserRepository interface {
	Create(context.Context, User) (User, error)
	Update(context.Context, User) error
}
