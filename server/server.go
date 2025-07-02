package server

import (
	"io/fs"
	"net/http"

	"github.com/withzeus/mugi-identity/routes"
	"github.com/withzeus/mugi-identity/services/db"
)

type Server struct {
	*http.Server
}

type ServerConfig struct {
	Host           string
	Port           string
	TemplateAssets fs.FS
	StaticAssets   fs.FS
}

func New(cfg ServerConfig) *Server {
	db := db.NewORM()
	router := http.NewServeMux()

	staticsHandler := routes.NewStaticContentsRouter(cfg.StaticAssets)
	staticsHandler.RegisterStaticRoutes(router)

	oauth2Handler := routes.NewOAuthRouter(cfg.TemplateAssets, db)
	oauth2Handler.RegisterRoutes(router)

	return &Server{
		Server: &http.Server{
			Addr:    cfg.Port,
			Handler: router,
		}}
}
