package middleware

import (
	"net/http"
	"time"

	"github.com/Adictes/food-chooser/alog"
	"github.com/julienschmidt/httprouter"
)

// AccessLog is middleware function that prints
// request method, remote address, url, and rate limit
// with accessLoger
func AccessLog(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		ac := alog.New()
		defer func() {
			ac.Logger.Printf("[%s] from %s to %s %s\n", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
		}()
		h(w, r, ps)
	}
}
