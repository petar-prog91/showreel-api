package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/petar-prog91/showreel-api/users_service/models"
)

func StatusUnauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)

	if err := json.NewEncoder(w).Encode(models.JsonErr{Code: http.StatusUnauthorized, Text: "Authorization failed."}); err != nil {
		panic(err)
	}
}

func StatusAuthFail(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusExpectationFailed)

	if err := json.NewEncoder(w).Encode(models.JsonErr{Code: http.StatusExpectationFailed, Text: "Authentification failed."}); err != nil {
		panic(err)
	}
}

func StatusNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	if err := json.NewEncoder(w).Encode(models.JsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func StatusUsernameExists(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusConflict)

	if err := json.NewEncoder(w).Encode(models.JsonErr{Code: http.StatusConflict, Text: "Username already exists. Please try other."}); err != nil {
		panic(err)
	}
}

func StatusBadRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(w).Encode(models.JsonErr{Code: http.StatusBadRequest, Text: "Something went wrong. Please try again"}); err != nil {
		panic(err)
	}
}

func StatusOK(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(models.JsonErr{Code: http.StatusOK}); err != nil {
		panic(err)
	}
}
