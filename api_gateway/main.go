package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/petar-prog91/showreel-api/helpers"

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

	// Users Service
	app.Options("/api/users/**", helpers.CorsHandler(usersServiceProxy))
	app.Get("/api/users/**", helpers.AuthHandler(usersServiceProxy))
	app.Post("/api/users/**", defaultHandler(usersServiceProxy))
	app.Put("/api/users/**", helpers.AuthHandler(usersServiceProxy))
	app.Delete("/api/users/**", helpers.AuthHandler(usersServiceProxy))
	app.Patch("/api/users/**", helpers.AuthHandler(usersServiceProxy))

	// Auth Service
	app.Post("/api/authenticate/**", defaultHandler(authServiceProxy))
	app.Options("/api/authenticate/**", helpers.CorsHandler(authServiceProxy))

	app.Run()
}

func defaultHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request, martini.Params) {
	return func(w http.ResponseWriter, r *http.Request, params martini.Params) {
		p.ServeHTTP(w, r)
	}
}
