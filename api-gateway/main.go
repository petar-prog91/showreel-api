package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/codegangsta/martini"
)

func main() {
	usersService, err := url.Parse("http://localhost:8081")
	if err != nil {
		panic(err)
	}

	usersServiceProxy := httputil.NewSingleHostReverseProxy(usersService)
	app := martini.Classic()

	app.Get("/api/users/**", handler(usersServiceProxy))
	app.Post("/api/users/**", handler(usersServiceProxy))
	app.Put("/api/users/**", handler(usersServiceProxy))
	app.Delete("/api/users/**", handler(usersServiceProxy))
	app.Patch("/api/users/**", handler(usersServiceProxy))

	app.RunOnAddr(":8080")
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request, martini.Params) {
	return func(w http.ResponseWriter, r *http.Request, params martini.Params) {
		fmt.Println(params)
		p.ServeHTTP(w, r)
	}
}
