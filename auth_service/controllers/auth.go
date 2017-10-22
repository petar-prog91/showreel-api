package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"

	"github.com/petar-prog91/showreel-api/auth_service/actions"
	"github.com/petar-prog91/showreel-api/auth_service/models"
	"github.com/petar-prog91/showreel-api/helpers"
)

func Authenticate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	// Let's check if Username already exists
	foundUser := actions.ReadUserByUsername(user.Username)

	if foundUser.Id > 0 {

		// let's compare user password from DB and user.Password which is in data
		err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))

		if err != nil {
			fmt.Println("Password not good")
			helpers.StatusAuthFail(w)

			return
		}

		createdToken, err := helpers.GenerateNewToken(foundUser.Id, foundUser.Username)

		if err != nil {
			fmt.Println("Creating token failed")
			helpers.StatusAuthFail(w)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(createdToken); err != nil {
			helpers.StatusAuthFail(w)
		}

		return
	}

	// If we didn't authenticate or found an user, return 404
	helpers.StatusAuthFail(w)
}
