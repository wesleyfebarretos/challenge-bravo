package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/challenge-bravo/internal/config"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/db"
	"github.com/wesleyfebarretos/challenge-bravo/internal/route"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	config.Init()

	if err := db.Init(); err != nil {
		log.Fatalf("db connection error %v", err)
	}

	defer db.Conn.Close()

	routes := route.Init()

	if err := routes.Run(fmt.Sprintf(":%s", config.Envs.Port)); err != nil {
		log.Fatalf("Error on starting API: %v", err)
	}
}
