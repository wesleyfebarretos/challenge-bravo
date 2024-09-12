package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/config"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/infra/db"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/route"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/scheduler"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/task"
)

func init() {
	// Set root dir
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../")
	os.Chdir(dir)
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	config.Init()

	if err := db.Init(); err != nil {
		log.Fatalf("db connection error %v", err)
	}

	defer db.Conn.Close()

	scheduler := scheduler.New()

	task.NewCurrencyUpdater().Start()

	scheduler.Start()

	routes := route.Init()

	if err := routes.Run(fmt.Sprintf(":%s", config.Envs.Port)); err != nil {
		log.Fatalf("Error on starting API: %v", err)
	}
}
