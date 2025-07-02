package routes

import (
	"io/fs"
	"log"
	"net/http"
)

type StaticContentsRouter struct {
	FS fs.FS
}

func NewStaticContentsRouter(assets fs.FS) *StaticContentsRouter {
	sf, err := fs.Sub(assets, "static")
	if err != nil {
		log.Fatal(err)
	}
	return &StaticContentsRouter{FS: sf}
}

func (s *StaticContentsRouter) RegisterStaticRoutes(router *http.ServeMux) {
	staticFileServer := http.FileServerFS(s.FS)
	router.Handle("GET /static/", http.StripPrefix("/static/", staticFileServer))
}
