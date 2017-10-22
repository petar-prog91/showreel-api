package main

import (
	"net/http"
	"os"

	"showreel-api/auth_service/controllers"
	"showreel-api/helpers"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.POST("/api/authenticate/", controllers.Authenticate)

	http.ListenAndServe(":8082", helpers.corsHandler(handlers.LoggingHandler(os.Stdout, router)))
}
