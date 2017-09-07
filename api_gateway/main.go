package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"skolar-api/api_gateway/helpers"
	"skolar-api/api_gateway/services"

	"github.com/codegangsta/martini"
)

func main() {
	usersService, err := url.Parse("http://localhost:8081")
	authService, err := url.Parse("http://localhost:8082")

	if err != nil {
		panic(err)
	}

	usersServiceProxy := httputil.NewSingleHostReverseProxy(usersService)
	authServiceProxy := httputil.NewSingleHostReverseProxy(authService)

	app := martini.Classic()

	app.Get("/api/users/**", handler(usersServiceProxy))
	app.Post("/api/users/**", handler(usersServiceProxy))
	app.Put("/api/users/**", handler(usersServiceProxy))
	app.Delete("/api/users/**", handler(usersServiceProxy))
	app.Patch("/api/users/**", handler(usersServiceProxy))

	app.Get("/api/authenticate/**", handler(authServiceProxy))
	app.Post("/api/authenticate/**", handler(authServiceProxy))
	app.Put("/api/authenticate/**", handler(authServiceProxy))
	app.Delete("/api/authenticate/**", handler(authServiceProxy))
	app.Patch("/api/authenticate/**", handler(authServiceProxy))

	app.RunOnAddr(":8080")
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request, martini.Params) {
	return func(w http.ResponseWriter, r *http.Request, params martini.Params) {
		var jwtToken = r.Header["Auth_jwt_token"]

		if len(jwtToken) > 0 {
			var validToken, _, err = services.ParseToken(jwtToken[0])

			if err != nil {
				helpers.StatusUnauthorized(w)
			}

			if validToken {
				// Delegate request to the given handle
				p.ServeHTTP(w, r)
			} else {
				// Request Basic Authentication otherwise
				helpers.StatusUnauthorized(w)
			}
		} else {
			// Request Basic Authentication otherwise
			helpers.StatusUnauthorized(w)
		}
	}
}
