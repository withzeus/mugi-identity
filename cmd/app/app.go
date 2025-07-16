package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/withzeus/mugi-identity/core"
	"github.com/withzeus/mugi-identity/core/db/pgsql"
	"github.com/withzeus/mugi-identity/core/lib"
	"github.com/withzeus/mugi-identity/identity"
	"github.com/withzeus/mugi-identity/tenant"
)

type App struct {
	helper core.Helper
	mux    *lib.RestServeMux
}

func NewApp() *App {
	h := core.Helper{}
	m := lib.NewRestServeMux()
	return &App{helper: h, mux: m}
}

func (app *App) Run() {
	pool, close, err := pgsql.NewPgxPool(pgsql.DBConfig{
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

	port := app.helper.GetEnv("APP_PORT", "8000")

	idhandler := identity.NewHandler(pool, app.helper)
	app.mux.Post("/users", idhandler.CreateUser)

	tenantshandler := tenant.NewHandler(pool, app.helper)
	app.mux.Post("/tenants", tenantshandler.CreateTenant)

	log.Printf("HTTP - server started on :%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), app.mux))
}
