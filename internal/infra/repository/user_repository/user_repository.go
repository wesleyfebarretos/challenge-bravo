package user_repository

import (
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/db"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/repository/sqlc/user_connection"
)

var (
	once       sync.Once
	repository entity.UserRepository
)

type UserRepository struct {
	queries *user_connection.Queries
}

func (u UserRepository) WithTx(tx pgx.Tx) entity.UserRepository {
	return &UserRepository{
		queries: u.queries.WithTx(tx),
	}
}

func New() entity.UserRepository {
	once.Do(func() {
		repository = &UserRepository{
			queries: user_connection.New(db.Conn),
		}
	})
	return repository
}
