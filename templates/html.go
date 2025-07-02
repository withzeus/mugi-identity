package html

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed *.html
//go:embed views/*.html
var templatesFS embed.FS

type Template struct {
	templates *template.Template
}

func New() *Template {
	templates := template.Must(template.ParseFS(templatesFS, "layout.html", "views/*.html"))
	return &Template{
		templates: templates,
	}
}

func (t *Template) GetFS() *template.Template {
	return t.templates
}

func (t *Template) Render(w http.ResponseWriter, name string, data any) error {
	tmpl := template.Must(t.templates.Clone())

	view := fmt.Sprintf("views/%s.html", name)
	template, err := tmpl.ParseFS(templatesFS, view)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error:" "page_not_found"}`))
		return err
	}
	return template.ExecuteTemplate(w, "layout.html", data)
}
