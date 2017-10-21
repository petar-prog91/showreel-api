package models

type User struct {
	Id       int    `db:"id, omitempty"`
	Username string `db:"username, omitempty"`
	Password string `db:"password, omitempty"`
	SGroup   int    `db:"s_group, omitempty"`
	Email    string `db:"email, omitempty"`
}
