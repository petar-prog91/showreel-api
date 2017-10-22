package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"

	"github.com/petar-prog91/showreel-api/helpers"
	"github.com/petar-prog91/showreel-api/users_service/controllers"
)

func main() {
	router := httprouter.New()
	// Role 0: Admin
	// Role 1: Users

	router.POST("/api/users/", controllers.CreateUser)
	router.PUT("/api/users/:id", helpers.JwtAuth(controllers.UpdateUser, 1))
	router.PATCH("/api/users/:id", helpers.JwtAuth(controllers.UpdateUser, 1))

	router.GET("/api/users/", helpers.JwtAuth(controllers.GetUsers, 0))
	router.GET("/api/users/:id", helpers.JwtAuth(controllers.GetUser, 1))
	router.DELETE("/api/users/:id", helpers.JwtAuth(controllers.DeleteUser, 1))

	http.ListenAndServe(":8081", handlers.LoggingHandler(os.Stdout, helpers.CorsHandlerRouter(router)))
}
