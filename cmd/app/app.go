package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/withzeus/mugi-identity/core"
	"github.com/withzeus/mugi-identity/core/db"
	"github.com/withzeus/mugi-identity/identity"
	"github.com/withzeus/mugi-identity/tenant"
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

	identityHandler := identity.NewHandler(pool, app.helper)
	tenantHandler := tenant.NewHandler(pool, app.helper)

	port := app.helper.GetEnv("APP_PORT", "8000")

	log.Printf("HTTP - registered %s", "/users")
	http.Handle("/users", identityHandler)
	log.Printf("HTTP - registered %s", "/tenants")
	http.Handle("/tenants", tenantHandler)

	log.Printf("HTTP - server started on :%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
