package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Adapter wraps a httprouter.Handle with additional functionality
type Adapter func(httprouter.Handle) httprouter.Handle

// AccessLog is middleware function that prints
// request method, remote address, url, and rate limit
func AccessLog() Adapter {
	return func(h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			start := time.Now()
			defer func() {
				log.Printf("[%s] from %s to %s %s\n", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
			}()
			h(w, r, ps)
		}
	}
}

// Adapt applies all specified adapters to h
func Adapt(h httprouter.Handle, adapters ...Adapter) httprouter.Handle {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
