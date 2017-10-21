package config

import "upper.io/db.v3/mysql"

var DBSettings = mysql.ConnectionURL{
	Database: `showreel`,
	Host:     `showreel_mysql_db:3306`,
	User:     `root`,
	Password: `root`,
}
