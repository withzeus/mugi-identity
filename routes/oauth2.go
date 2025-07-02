package routes

import (
	"io/fs"
	"net/http"

	"github.com/withzeus/mugi-identity/services/db"
	"github.com/withzeus/mugi-identity/services/oauth2"
)

type OAuthRouter struct {
	handler *oauth2.AuthService
}

func NewOAuthRouter(viewsFS fs.FS, orm *db.ORM) *OAuthRouter {
	as := oauth2.NewAuthService(viewsFS, orm)
	return &OAuthRouter{handler: as}
}

func (a *OAuthRouter) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /client/register", a.handler.RegisterClientApp)
	router.HandleFunc("GET /authorize", a.handler.Authorize)
	router.HandleFunc("/login", a.handler.Login)
	router.HandleFunc("/sign-up", a.handler.SignUp)
}
