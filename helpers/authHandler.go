package helpers

import (
	"net/http"
	"net/http/httputil"

	"github.com/codegangsta/martini"
)

func AuthHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request, martini.Params) {
	return func(w http.ResponseWriter, r *http.Request, params martini.Params) {
		var jwtToken = r.Header["Auth_jwt_token"]

		if len(jwtToken) > 0 {
			var validToken, _, err = ParseToken(jwtToken[0])

			if err != nil {
				StatusUnauthorized(w)
			}

			if validToken {
				// Delegate request to the given handle
				p.ServeHTTP(w, r)
			} else {
				// Request Basic Authentication otherwise
				StatusUnauthorized(w)
			}
		} else {
			// Request Basic Authentication otherwise
			StatusUnauthorized(w)
		}
	}
}
