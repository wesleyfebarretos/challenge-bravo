package integration

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/config"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/enum"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/infra/db"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/route"
	testcontainer "github.com/wesleyfebarretos/challenge-bravo/pkg/test_container"
)

var (
	runningContainers = []*testcontainer.ContainerResult{}
	server            = &httptest.Server{}
	client            = &http.Client{}
)

func beforeAll() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error trying to find working dir: %v", err)
	}

	if err := godotenv.Load(fmt.Sprintf("%s/.env.test", wd)); err != nil {
		log.Fatal(err)
	}

	config.Init()

	// Moving default logger output to Stderr which is closed to not spam the tests
	config.Envs.Log.Output = os.Stderr

	pgContainer := testcontainer.SetupPG(testcontainer.PgConfig{
		Image:    "postgres:13-alpine",
		Database: config.Envs.DB.Name,
		User:     config.Envs.DB.User,
		Password: config.Envs.DB.Password,
	})

	runningContainers = append(runningContainers, pgContainer)

	config.Envs.DB.Port = fmt.Sprintf("%d", pgContainer.Port)

	if _, err := db.Init(); err != nil {
		log.Fatal(err)
	}

	if err = db.RunMigrations(context.TODO()); err != nil {
		log.Fatalf("setup migrations error %v", err)
	}

	db.BeginTestTxWrapper(context.TODO())

	// Set application to test enviroment
	config.Envs.AppEnv = enum.TEST_ENVIROMENT

	server = httptest.NewServer(route.Init())
}
