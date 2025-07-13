package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/withzeus/mugi-identity/core"
	"github.com/withzeus/mugi-identity/core/db"
	"github.com/withzeus/mugi-identity/identity"
)

type App struct {
	helper core.Helper
}

func NewApp() *App {
	h := core.Helper{}
	return &App{helper: h}
}

func (app *App) Run() {
	pool, close, err := db.NewPgxPool(db.DBConfig{
		Username: app.helper.GetEnv("PG_USERNAME", "postgres"),
		Password: app.helper.GetEnv("PG_PASSWORD", "postgres"),
		Hostname: app.helper.GetEnv("PG_HOSTNAME", "localhost"),
		Port:     app.helper.GetEnv("PG_PORT", "5432"),
		DBName:   app.helper.GetEnv("PG_DBNAME", "mugi_idp"),
	})

	if err != nil {
		log.Fatalf("db.NewPgxPool - %v\n", err)
	}

	defer close()

	userHandler := identity.NewHandler(pool, app.helper)

	port := app.helper.GetEnv("APP_PORT", "8000")

	http.Handle("/users", userHandler)

	log.Printf("http.server - started on :%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
