package handlers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var t = template.Must(template.New("FC").ParseGlob("templates/*.html"))

// Index is main page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t.ExecuteTemplate(w, "index", nil)
}
