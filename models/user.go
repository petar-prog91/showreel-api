package models

type User struct {
	Id        int    `db:"id, omitempty"`
	Username  string `db:"username, omitempty"`
	Password  string `db:"password, omitempty"`
	UserRole  int    `db:"user_role, omitempty"`
	FirstName string `db:"first_name, omitempty"`
	LastName  string `db:"last_name, omitempty"`
	Address   string `db:"address, omitempty"`
	City      string `db:"city, omitempty"`
	ZipCode   int    `db:"zipcode, omitempty"`
	Email     string `db:"email, omitempty"`
	Phone     string `db:"phone, omitempty"`
}

type UserRoles struct {
	Id        int    `db:"id"`
	Role_name string `db:"role_name"`
}
