package tenant

import (
	"fmt"
	"log"
	"net/http"

	"github.com/withzeus/mugi-identity/core"
	"github.com/withzeus/mugi-identity/core/db"
	"github.com/withzeus/mugi-identity/core/lib"
)

type Handler struct {
	S Service
}

func NewHandler(pool db.IPgx, helper core.Helper) Handler {
	store := NewDatastore(pool)
	service := NewService(store, helper)
	return Handler{S: *service}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Handle(w, r)
	if err != nil {
		switch e := err.(type) {
		case lib.HttpError:
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

func (handler Handler) Handle(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodPost:
		log.Printf("HTTP %s %s - %s", r.Method, r.URL.Path, "Requested")
		return handler.CreateTenant(w, r)
	default:
		return lib.HttpStatusError{Code: http.StatusMethodNotAllowed, Err: fmt.Errorf("method not allowed")}
	}
}

func (handler Handler) CreateTenant(w http.ResponseWriter, r *http.Request) error {
	var md Model

	handler.S.h.FromIoReader(r.Body, &md)

	err := md.Validate()
	if err != nil {
		return lib.HttpStatusError{Code: 404, Err: err}
	}

	tenant, err := handler.S.Create(md)
	if err != nil {
		return lib.HttpStatusError{Code: 500, Err: err}
	}

	w.Header().Set("Content-Type", "application/json")

	log.Printf("HTTP 201 - %s", "Created")
	fmt.Fprintf(w, "%s", handler.S.h.ToJSON(tenant))

	return nil
}
