package actions

import (
	"log"

	"upper.io/db.v3/mysql"

	"github.com/petar-prog91/showreel-api/auth_service/config"
	"github.com/petar-prog91/showreel-api/auth_service/models"
)

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
