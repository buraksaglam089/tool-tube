package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	listenAdrr := os.Getenv("LISTEN_ADDR")
	db, err := SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(listenAdrr, db)
	server.Run()
}
