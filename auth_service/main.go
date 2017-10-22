package main

import (
	"net/http"
	"os"

	"github.com/petar-prog91/showreel-api/auth_service/controllers"
	"github.com/petar-prog91/showreel-api/helpers"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.POST("/api/authenticate/", controllers.Authenticate)

	http.ListenAndServe(":8082", helpers.CorsHandler(handlers.LoggingHandler(os.Stdout, router)))
}
