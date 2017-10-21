package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"showreel-api/api_gateway/helpers"
	"showreel-api/api_gateway/services"

	"github.com/codegangsta/martini"
)

func main() {
	usersService, err := url.Parse("http://users_service:8081")
	authService, err := url.Parse("http://auth_service:8082")

	if err != nil {
		panic(err)
	}

	usersServiceProxy := httputil.NewSingleHostReverseProxy(usersService)
	authServiceProxy := httputil.NewSingleHostReverseProxy(authService)

	app := martini.Classic()

	app.Get("/api/users/**", authHandler(usersServiceProxy))
	app.Post("/api/users/**", authHandler(usersServiceProxy))
	app.Put("/api/users/**", authHandler(usersServiceProxy))
	app.Delete("/api/users/**", authHandler(usersServiceProxy))
	app.Patch("/api/users/**", authHandler(usersServiceProxy))

	app.Post("/api/authenticate/**", defaultHandler(authServiceProxy))

	app.Run()
}

func authHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request, martini.Params) {
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

func defaultHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request, martini.Params) {
	return func(w http.ResponseWriter, r *http.Request, params martini.Params) {
		p.ServeHTTP(w, r)
	}
}
