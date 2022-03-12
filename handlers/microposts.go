package handlers

import (
	"fmt"
	"my_microposts/data"
	"net/http"
)

type MicropostsHandler struct{}

var Microposts *MicropostsHandler = &MicropostsHandler{}

func (handler *MicropostsHandler) Index(w http.ResponseWriter, r *http.Request) {
	filenames := []string{
		"application/head",
		"microposts/index",
	}

	// get all microposts
	microposts, err := data.Microposts()
	fmt.Println(microposts)

	if err != nil {
		panic("can't get microposts")
	}

	generateHTML(w, microposts, filenames...)
}

func (handler *MicropostsHandler) New(w http.ResponseWriter, r *http.Request) {
	// associated filenames needed for html
	filenames := []string{
		"application/head",
		"microposts/new",
	}

	generateHTML(w, nil, filenames...)
}

func (handler *MicropostsHandler) Create(w http.ResponseWriter, r *http.Request) {
	// parse form first
	if err := r.ParseForm(); err != nil {
		panic("cannot parse form")
	} else {

		_, err = data.CreateMicropost(r.PostFormValue("title"), r.PostFormValue("content"))

		if err != nil {
			panic(err)
		}

		// redirect to root path after create
		http.Redirect(w, r, "/", 302)
	}
}
