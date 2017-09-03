package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"

	"skolar/services/controllers"
	"skolar/services/helpers"
	"skolar/services/services"
)

func main() {
	router := httprouter.New()
	router.POST("/api/authenticate", controllers.Authenticate)

	// Regular Role: 1
	router.PUT("/api/users/:id", JwtAuth(controllers.UpdateUser, 1))

	// Admin Role: 3
	router.GET("/api/users", JwtAuth(controllers.GetUsers, 3))
	router.GET("/api/users/:id", JwtAuth(controllers.GetUser, 3))
	router.POST("/api/users", JwtAuth(controllers.CreateUser, 3))
	router.DELETE("/api/users/:id", JwtAuth(controllers.DeleteUser, 3))

	http.ListenAndServe(":8081", corsHandler(handlers.LoggingHandler(os.Stdout, router)))
}

func JwtAuth(h httprouter.Handle, reqUserRole int) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var jwtToken = r.Header["Auth_jwt_token"]

		if len(jwtToken) > 0 {
			var validToken, claims, err = services.ParseToken(jwtToken[0])

			if err != nil {
				helpers.StatusUnauthorized(w)
			}

			var jwtUserRole = claims.UserRole

			if validToken && jwtUserRole >= reqUserRole {
				// Delegate request to the given handle
				h(w, r, ps)
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

func corsHandler(h http.Handler) http.HandlerFunc {
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

		h.ServeHTTP(w, r)
	}
}
