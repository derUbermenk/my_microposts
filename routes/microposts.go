package routes

import (
	"html/template"
	"net/http"
)

type MicropostsHandler struct{}

var Microposts *MicropostsHandler = &MicropostsHandler{}

func (handler *MicropostsHandler) Index(w http.ResponseWriter, r *http.Request) {
	// get all microposts
	files := []string{
		"templates/application/head.html",
		"templates/microposts/index.html",
	}

	template := template.Must(template.ParseFiles(files...))
	template.ExecuteTemplate(w, "layout", "")
}

func (handler *MicropostsHandler) New(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/application/head.html",
		"templates/microposts/new.html",
	}

	template := template.Must(template.ParseFiles(files...))
	template.ExecuteTemplate(w, "layout", "")
}

func (handler *MicropostsHandler) Create(w http.ResponseWriter, r *http.Request) {
	// create micropost
}
