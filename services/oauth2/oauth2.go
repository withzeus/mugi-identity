package oauth2

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/withzeus/mugi-identity/config"
	"github.com/withzeus/mugi-identity/helpers"
	"github.com/withzeus/mugi-identity/services/db"
	html "github.com/withzeus/mugi-identity/templates"
	"gorm.io/gorm"
)

type AuthService struct {
	orm       *db.ORM
	templates *html.Template
	store     *db.SessionStore
}

func NewAuthService(tmpfs fs.FS, orm *db.ORM) *AuthService {
	templates := html.New()
	orm.AutoMigrate(&db.Client{}, &db.User{})

	sessionStore := db.NewSessionStore(orm)
	sessionCleanup := make(chan struct{})
	go sessionStore.PeriodicCleanup(1*time.Hour, sessionCleanup)

	return &AuthService{templates: templates, orm: orm, store: sessionStore}
}

type AuthorizeRequest struct {
	State        string `schema:"state" json:"state" validate:"required"`
	ResponseType string `schema:"response_type" json:"response_type" validate:"required"`
	ClientId     string `schema:"client_id" json:"client_id" validate:"required"`
	RedirectUri  string `schema:"redirect_uri" json:"redirect_uri" validate:"required"`
	Scope        string `schema:"scope" json:"scope" validate:"required"`
}

func (as *AuthService) Authorize(w http.ResponseWriter, r *http.Request) {
	var req AuthorizeRequest

	if err := helpers.ReadSchemaJSON(r.URL.Query(), &req); err != nil {
		helpers.SendBadRequestError(w)
		return
	}

	if err := helpers.Validate.Struct(req); err != nil {
		helpers.SendBadRequestError(w)
		return
	}

	session, err := as.store.Get(r, config.ServerLookup.SK)
	if err != nil {
		helpers.SendBadRequestError(w)
		return
	}

	state, err := helpers.ToJSON(req)
	if err != nil {
		helpers.SendBadRequestError(w)
		return
	}

	session.Values["state"] = string(state)
	err = session.Save(r, w)
	if err != nil {
		helpers.SendBadRequestError(w)
		return
	}

	uid, notFound := r.Cookie(config.ServerLookup.UID)
	if notFound == http.ErrNoCookie {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	}

	fmt.Printf("%v \n", uid)
	helpers.WriteJSON(w, &req)
}

type LoginRequest struct {
	Handle   string `json:"handle" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (as *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	session, err := as.store.Get(r, config.ServerLookup.SK)
	if err != nil {
		helpers.SendHttpError(w, http.StatusInternalServerError, "server error")
		return
	}

	if session.Values["state"] == nil {
		helpers.SendBadRequestError(w)
		return
	}

	var req AuthorizeRequest
	if err := helpers.FromJSON(session.Values["state"].(string), &req); err != nil {
		helpers.SendBadRequestError(w)
		return
	}

	if r.Method == http.MethodGet {
		if err := as.templates.Render(w, "sign-in", nil); err != nil {
			helpers.SendBadRequestError(w)
			return
		}
		return
	}

	if err := r.ParseForm(); err != nil {
		helpers.SendBadRequestError(w)
		return
	}

	var login LoginRequest
	if err := helpers.ReadSchemaJSON(r.PostForm, &login); err != nil {
		helpers.SendBadRequestError(w)
		return
	}

	lt, err := helpers.ClassifyUserHandle(login.Handle)
	if err != nil {
		as.returnWithErrors(w, "sign-in", "အကောင့်အချက်အလက်များ မှားယွင်းနေပါသည်။")
		return
	}

	if lt != "gmail" {
		as.mgLogin(w, login.Handle, login.Password)
		return
	}
}

func (as *AuthService) returnWithErrors(w http.ResponseWriter, content string, message string) {
	if err := as.templates.Render(w, "sign-in", map[string]any{
		"Error": message,
	}); err != nil {
		helpers.SendBadRequestError(w)
		return
	}
}

func (as *AuthService) mgLogin(w http.ResponseWriter, handle string, password string) {
	user := &db.User{
		Handle: handle,
	}
	result := user.QueryByHandle(as.orm)
	if result.Error != nil || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		as.returnWithErrors(w, "sign-in", "အကောင့်အချက်အလက်များ မှားယွင်းနေပါသည်။")
	}
}

func (as *AuthService) SignUp(w http.ResponseWriter, r *http.Request) {
}

type ClientRegisterRequest struct {
	ApplicationName string   `json:"application_name"`
	Website         string   `json:"website"`
	Logo            string   `json:"logo"`
	RedirectUri     []string `json:"redirect_uri"`
}

func (as *AuthService) RegisterClientApp(w http.ResponseWriter, r *http.Request) {
	var req ClientRegisterRequest
	if err := helpers.ReadJSON(r.Body, &req); err != nil {
		helpers.SendHttpError(w, http.StatusBadRequest, "invalid_request")
		return
	}

	if req.ApplicationName == "" {
		helpers.SendHttpError(w, http.StatusBadRequest, "invalid_request")
		return
	}

	if len(req.RedirectUri) == 0 {
		helpers.SendHttpError(w, http.StatusBadRequest, "invalid_request")
		return
	}

	client := &db.Client{
		Name:        req.ApplicationName,
		Website:     req.Website,
		Logo:        req.Logo,
		RedirectURI: strings.Join(req.RedirectUri, " "),
	}
	if err := client.Create(as.orm); err != nil {
		helpers.SendHttpError(w, http.StatusInternalServerError, "server_error")
		return
	}

	helpers.WriteJSON(w, &req)
}
