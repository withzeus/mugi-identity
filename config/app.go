package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port        string
	DatabaseUrl string
	Mode        string
}

var AppEnv = loadEnvConfig()

func loadEnvConfig() Env {
	godotenv.Load()

	return Env{
		Port:        fmt.Sprintf(":%s", getEnv("PORT", "8000")),
		DatabaseUrl: getEnv("POSTGRES_URL", ""),
		Mode:        getEnv("Mode", "development"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
