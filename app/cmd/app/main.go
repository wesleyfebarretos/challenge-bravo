package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/config"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/infra/db"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/route"
	aredis "github.com/wesleyfebarretos/challenge-bravo/pkg/redis"
)

func init() {
	// Set root dir
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	os.Chdir(dir)
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	config.Init()

	err := aredis.Init(config.Envs.Redis.HostAndPort, config.Envs.Redis.Password)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.Init()
	if err != nil {
		log.Fatalf("db connection error %v", err)
	}

	defer conn.Close()

	ctx := context.Background()

	if err = db.RunMigrations(ctx); err != nil {
		log.Fatalf("setup migrations error %v", err)
	}

	routes := route.Init()

	if err := routes.Run(fmt.Sprintf(":%s", config.Envs.Port)); err != nil {
		log.Fatalf("Error on starting API: %v", err)
	}
}
