package identity

import (
	"fmt"
	"log"
	"net/http"

	"github.com/withzeus/mugi-identity/core"
	"github.com/withzeus/mugi-identity/core/db/pgsql"
	"github.com/withzeus/mugi-identity/core/lib"
)

type Handler struct {
	S Service
}

func NewHandler(pool pgsql.IPgx, helper core.Helper) Handler {
	store := NewDatastore(pool)
	service := NewService(store, helper)
	return Handler{S: *service}
}

func (handler Handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var md Model

	handler.S.h.FromIoReader(r.Body, &md)

	err := md.Validate()
	if err != nil {
		return lib.NewHttpStatusCode(400, "")
	}

	user, err := handler.S.Create(md)
	if err != nil {
		return lib.NewHttpStatusCode(400, "")
	}

	w.Header().Set("Content-Type", "application/json")

	log.Printf("HTTP 201 - %s", "Created")
	fmt.Fprintf(w, "%s", handler.S.h.ToJSON(user))
	return nil
}
