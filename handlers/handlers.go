package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index is main page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello, world!"))
}
