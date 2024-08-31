package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/challenge-bravo/internal/config"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/db"
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
}
