package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"skolar/services/actions"
	"skolar/services/helpers"
	"skolar/services/models"
)

func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var userId int
	var err error

	if userId, err = strconv.Atoi(ps.ByName("id")); err != nil {
		panic(err)
	}

	user := actions.ReadUserById(userId)

	if user.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(user); err != nil {
			panic(err)
		}

		return
	}

	// If we didn't find it, 404
	helpers.StatusNotFound(w)
}

func GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users := actions.ReadAllUsers()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	// Let's check if Username already exists
	userExists := actions.ReadUserByUsername(user.Username)

	// If user exists, lets return JSON 409 - StatusConflict and message that user already exists
	if userExists.Id > 0 {
		helpers.StatusUsernameExists(w)

		return
	}

	// Let's create a new user in database
	newUser := actions.CreateUser(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	// If something goes wrong, lets return JSON 404 and message
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		helpers.StatusBadRequest(w)

		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user map[string]interface{}
	var userId int
	var err error

	if userId, err = strconv.Atoi(ps.ByName("id")); err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	// Let's check if Username already exists
	userExists := actions.ReadUserById(userId)

	// If user exists, lets update it
	if userExists.Id > 0 {
		newUser := actions.UpdateUser(userId, user)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(newUser); err != nil {
			helpers.StatusBadRequest(w)
		}

		return
	}

	helpers.StatusNotFound(w)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var userId int
	var err error

	if userId, err = strconv.Atoi(ps.ByName("id")); err != nil {
		panic(err)
	}

	// Let's check if Username already exists
	userExists := actions.DeleteUser(userId)

	// If user not deleted, lets return JSON 404 - StatusBadRequest and message that something went wrong.
	if !userExists {
		helpers.StatusBadRequest(w)

		return
	}

	// If user deleted, lets return JSON 200 - StatusOK
	helpers.StatusOK(w)
}
