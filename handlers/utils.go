package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// generates an html response based by parsing files named by the filenames
// and adding the necessary data provided by the data params to the filenames
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string

	for _, filename := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", filename))
	}

	template := template.Must(template.ParseFiles(files...))
	template.ExecuteTemplate(w, "layout", data)
}
