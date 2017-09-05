package config

import "upper.io/db.v3/mysql"

var DBSettings = mysql.ConnectionURL{
	Database: `skolar_users`,
	Host:     `mysql:3306`,
	User:     `root`,
	Password: `root`,
}
