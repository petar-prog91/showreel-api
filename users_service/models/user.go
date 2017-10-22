package models

type User struct {
	Id       int    `db:"id, omitempty"`
	Username string `db:"username, omitempty"`
	Password string `db:"password, omitempty"`
	Sgroup   int    `db:"sgroup, omitempty"`
	Email    string `db:"email, omitempty"`
}
