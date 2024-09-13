package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/config"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/infra/db"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/infra/service"
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

	currencies, err := service.NewCurrencyUpdaterService(http.Client{}).GetCurrenciesExchangeRatesInUSD(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	err = service.NewCurrencyUpdaterService(http.Client{}).UpdateRates(context.Background(), currencies)
	if err != nil {
		fmt.Print(err)
	}

	routes := route.Init()

	if err := routes.Run(fmt.Sprintf(":%s", config.Envs.Port)); err != nil {
		log.Fatalf("Error on starting API: %v", err)
	}
}
