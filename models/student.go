package models

type Student struct {
	ID         int    `db:"id, omitempty"`
	EnrollDate string `db:"enrolldate, omitempty"`
	Parent     int    `db:"parent, omitempty"`
	ClassGroup int    `db:"class_group, omitempty"`
}
