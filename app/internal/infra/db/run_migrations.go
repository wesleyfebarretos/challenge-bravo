package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/config"
)

func RunMigrations(ctx context.Context) error {
	connString := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Envs.DB.Driver,
		config.Envs.DB.User,
		config.Envs.DB.Password,
		config.Envs.DB.Host,
		config.Envs.DB.Port,
		config.Envs.DB.Name,
	)

	migrate, err := migrate.New(
		"file://internal/migration",
		connString,
	)
	if err != nil {
		return err
	}

	if err = migrate.Up(); err != nil && !strings.Contains(err.Error(), "no change") {
		return err
	}

	if _, err := migrate.Close(); err != nil {
		return err
	}

	return nil
}
