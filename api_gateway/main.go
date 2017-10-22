package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

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
	app.Options("/api/users/**", helpers.corsHandler(usersServiceProxy))
	app.Get("/api/users/**", helpers.authHandler(usersServiceProxy))
	app.Post("/api/users/**", helpers.authHandler(usersServiceProxy))
	app.Put("/api/users/**", helpers.authHandler(usersServiceProxy))
	app.Delete("/api/users/**", helpers.authHandler(usersServiceProxy))
	app.Patch("/api/users/**", helpers.authHandler(usersServiceProxy))

	// Auth Service
	app.Post("/api/authenticate/**", defaultHandler(authServiceProxy))
	app.Options("/api/authenticate/**", helpers.corsHandler(authServiceProxy))

	app.Run()
}

func defaultHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request, martini.Params) {
	return func(w http.ResponseWriter, r *http.Request, params martini.Params) {
		p.ServeHTTP(w, r)
	}
}
