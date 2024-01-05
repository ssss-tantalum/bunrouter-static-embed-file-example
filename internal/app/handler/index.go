package handler

import (
	"net/http"
	"text/template"

	"github.com/ssss-tantalum/bunrouter-static-embed-file-example/internal/app"
	"github.com/uptrace/bunrouter"
)

type IndexHandler struct {
	app  *app.App
	tmpl *template.Template
}

func NewIndexHandler(app *app.App) *IndexHandler {
	tmpl := template.Must(template.New("").ParseFiles("templates/index.html"))

	return &IndexHandler{
		app:  app,
		tmpl: tmpl,
	}
}

func (h *IndexHandler) Index(w http.ResponseWriter, req bunrouter.Request) error {
	return h.tmpl.ExecuteTemplate(w, "Base", "World")
}
