package config

import "upper.io/db.v3/mysql"

var DBSettings = mysql.ConnectionURL{
	Database: `showreel`,
	Host:     `showreel_mysql_db:3308`,
	User:     `root`,
	Password: `root`,
}
