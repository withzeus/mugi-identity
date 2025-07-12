package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/withzeus/mugi-identity/core/db"
)

func main() {
	godotenv.Load()

	db, _, err := db.NewDBP(db.DBConfig{
		Username: "postgres",
		Password: "postgres",
		Hostname: "localhost",
		Port:     "5432",
		DBName:   "mugi_idp",
	})

	if err != nil {
		log.Fatalf("db: connection error %v \n", err)
	}

	defer db.Close()
}
