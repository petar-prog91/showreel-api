package config

import "upper.io/db.v3/mysql"

var DBSettings = mysql.ConnectionURL{
	Database: `skolar_users`,
	Host:     `127.0.0.1`,
	User:     `root`,
	Password: `root`,
}
