package actions

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"upper.io/db.v3/mysql"

	"github.com/petar-prog91/showreel-api/users_service/config"
	"github.com/petar-prog91/showreel-api/users_service/models"
)

func ReadAllUsers() []models.User {
	// Attemping to establish a connection to the database.
	dbSess, err := mysql.Open(config.DBSettings)

	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	defer dbSess.Close() // Remember to close the database dbSession.

	// Query all results and fill the users variable with them.
	var users []models.User

	err = dbSess.Collection("user").Find().All(&users)
	if err != nil {
		log.Printf("res.All(): %q\n", err)
		return []models.User{}
	}

	return users
}

func ReadUserById(id int) models.User {
	// Attemping to establish a connection to the database.
	dbSess, err := mysql.Open(config.DBSettings)

	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	defer dbSess.Close() // Remember to close the database dbSession.

	// Query all results and fill the users variable with them.
	var user models.User

	err = dbSess.Collection("user").Find("id", id).One(&user)
	if err != nil {
		log.Printf("res.All(): %q\n", err)
		return models.User{}
	}

	return user
}

func ReadUserByUsername(username string) models.User {
	// Attemping to establish a connection to the database.
	dbSess, err := mysql.Open(config.DBSettings)

	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	defer dbSess.Close() // Remember to close the database dbSession.

	// Query all results and fill the users variable with them.
	var user models.User

	err = dbSess.Collection("user").Find("username", username).One(&user)
	if err != nil {
		log.Printf("res.All(): %q\n", err)
		return models.User{}
	}

	return user
}

func CreateUser(userData models.User) models.User {
	// Attemping to establish a connection to the database.
	dbSess, err := mysql.Open(config.DBSettings)

	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	defer dbSess.Close() // Remember to close the database dbSession.

	userCollection := dbSess.Collection("user")

	// Query all results and fill the users variable with them.
	var user models.User

	var password = []byte(userData.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	userData.Password = string(hashedPassword)

	_, userErr := userCollection.Insert(models.User{
		Username: userData.Username,
		Password: userData.Password,
		Email:    userData.Email,
	})

	if userErr != nil {
		log.Printf("userCollection.Insert(): %q\n", userErr)
		return models.User{}
	}

	err = dbSess.Collection("user").Find("username", userData.Username).One(&user)

	if err != nil {
		log.Printf("res.All(): %q\n", err)
		return models.User{}
	}

	return user
}

func UpdateUser(userId int, userData map[string]interface{}) models.User {
	// Attemping to establish a connection to the database.
	dbSess, err := mysql.Open(config.DBSettings)

	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	defer dbSess.Close() // Remember to close the database dbSession.

	var user models.User

	q := dbSess.Update("user").Set(userData).Where("id = ?", userId)
	_, err = q.Exec()

	err = dbSess.Collection("user").Find("id", userId).One(&user)

	if err != nil {
		log.Printf("res.All(): %q\n", err)
		return models.User{}
	}

	return user
}

func DeleteUser(id int) bool {
	// Attemping to establish a connection to the database.
	dbSess, err := mysql.Open(config.DBSettings)

	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	defer dbSess.Close() // Remember to close the database dbSession.

	err = dbSess.Collection("user").Find("id", id).Delete()

	if err != nil {
		panic(err)

		return false
	}

	return true
}
