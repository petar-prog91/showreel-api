package helpers

import (
	"net/http"
	"net/http/httputil"
)

func corsHandler(p *httputil.ReverseProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}

		// Stop here for a Preflighted OPTIONS request.
		if r.Method == "OPTIONS" {
			return
		}

		p.ServeHTTP(w, r)
	}
}
