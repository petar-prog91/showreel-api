package config

import "upper.io/db.v3/mysql"

var DBSettings = mysql.ConnectionURL{
	Database: `skolar`,
	Host:     `127.0.0.1`,
	User:     `root`,
	Password: `root`,
}
